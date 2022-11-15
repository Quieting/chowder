package mall

import (
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"

	"github.com/Quieting/chowder/script/http"
	db "github.com/Quieting/chowder/script/sql"
	"github.com/Quieting/chowder/script/xerror"
)

const (
	PropsTableName = "`props`"
)

type PropListArg struct {
	Limit    int64
	Offset   int64
	Name     string // 道具名称
	Id       int64  // 道具id
	Cat      int64  // 道具分类
	ChildCat int64  // 道具子分类
	Scene    int64
}

func (arg *PropListArg) query(columns ...string) squirrel.SelectBuilder {
	builder := squirrel.Select(columns...).From(PropsTableName)
	if arg.Limit > 0 {
		builder = builder.Limit(uint64(arg.Limit))
	}
	if arg.Offset > 0 {
		builder = builder.Offset(uint64(arg.Limit))
	}
	if arg.Id != 0 {
		builder = builder.Where(squirrel.Eq{"id": arg.Id})
	}

	// 适用场景
	if arg.Scene > 0 {
		builder = builder.Where(squirrel.Expr("find_in_set(?, scene)", arg.Scene))
	}
	// 道具分类
	if arg.Cat != 0 {
		builder = builder.Where("cat = ?", arg.Cat)
	}

	// 上级分类
	if arg.ChildCat != 0 {
		builder = builder.Where(squirrel.Eq{"child_cat": arg.ChildCat})
	}

	// 名字
	if arg.Name != "" {
		builder = builder.Where(squirrel.Like{"name": fmt.Sprint("%", arg.Name, "%")})
	}

	return builder
}

type Prop struct {
	Id            int64     `db:"id"`             // 序号
	Name          string    `db:"name"`           // 道具名称
	Cat           int64     `db:"cat"`            // 道具分类 参考enum.PropsCategory
	ChildCat      int64     `db:"child_cat"`      // 子类，没有则默认0
	Scene         string    `db:"scene"`          // 适用场景，1:私聊 2:房间 3:游戏 多个用逗号分隔如 ,1,2,
	CharmValue    int64     `db:"charm_value"`    // 魅力值
	WealthValue   int64     `db:"wealth_value"`   // 财富值
	IntimateValue int64     `db:"intimate_value"` // 亲密度
	PicUrl        string    `db:"pic_url"`        // 图片地址
	TransPicUrl   string    `db:"trans_pic_url"`  // 无背景图地址
	AnimationUrl  string    `db:"animation_url"`  // 动画地址
	Desc          string    `db:"desc"`           // 道具描述
	Content       string    `db:"content"`        // 触发内容
	UpdateTime    time.Time `db:"update_time"`    // 更新时间
	CreateTime    time.Time `db:"create_time"`    // 创建时间
}

func PropList(p PropListArg) ([]Prop, error) {
	propsFieldNames := RawFieldNames(&Prop{})
	query := p.query(propsFieldNames...)

	str, args, err := query.ToSql()
	if err != nil {
		return nil, xerror.New(err, "组装sql语句失败")
	}

	list := make([]Prop, 0, p.Limit)

	err = db.QueryRows(&list, str, args...)
	if err != nil {
		return nil, xerror.New(err, "查询道具列表失败")
	}

	return list, nil
}

func OneProp(id int64) (*Prop, error) {
	p := PropListArg{Id: id}
	propsFieldNames := RawFieldNames(&Prop{})
	query := p.query(propsFieldNames...)

	str, args, err := query.ToSql()
	if err != nil {
		return nil, xerror.New(err, "组装sql语句失败")
	}

	item := new(Prop)

	err = db.QueryRow(item, str, args...)
	if err != nil {
		return nil, xerror.New(err, "查询道具列表失败")
	}

	return item, nil
}

type GoodsItem struct {
	Id            string          `json:"id"`             // 主键ID
	GoodsName     string          `json:"goods_name"`     // 商品名称
	Tag           string          `json:"tag"`            // 标签
	BuyType       int64           `json:"buy_type"`       // 0: 无限制 1: 今日限购 2:永久限购
	BuyNumber     int64           `json:"buy_number"`     // 已购买数量, max_buy_number 大于0时有意义
	MaxBuyNumber  int64           `json:"max_buy_number"` // 最大购买数量，0表示无限制
	EndTime       int64           `json:"end_time"`       // 结束出售时间
	SaleInfo      []GoodsSaleInfo `json:"sale_info"`      // 销售信息
	PropId        int64           `json:"prop_id"`        // 对应的道具ID(prop表的id)
	PropCat       int64           `json:"prop_cat"`       // 道具类型(便于查询)
	PropChildCat  int64           `json:"prop_child_cat"` // 道具子类型
	PicUrl        string          `json:"pic_url"`        // 图片地址
	TransPicUrl   string          `json:"trans_pic_url"`  // 图片地址（有背景）
	AnimationUrl  string          `json:"animation_url"`  // 动画地址
	Scene         []int64         `json:"scene"`          // 适用场景 1:私聊场景 2:房间内场景 3:游戏
	Desc          string          `json:"desc"`           // 描述
	IntimateValue int64           `json:"intimate_value"` // 亲密度
	CharmValue    int64           `json:"charm_value"`    // 魅力值
	WealthValue   int64           `json:"wealth_value"`   // 财富值
}

type GoodsSaleInfo struct {
	OriginPrice   int64 `json:"origin_price"`   // 原价
	Price         int64 `json:"price"`          // 实际售价
	Number        int64 `json:"number"`         // 数量
	EffectiveTime int64 `json:"effective_time"` // 有效时长，单位天，0表示永久有效
}

func GoodsList() ([]GoodsItem, error) {
	resp := &struct {
		HasMore bool        `json:"has_more"`
		List    []GoodsItem `json:"list"`
	}{}
	err := http.Get("http://localhost:50000/api/mall/goods/list?cat=3", nil, resp, "")
	if err != nil {
		return nil, err
	}

	return resp.List, nil
}
