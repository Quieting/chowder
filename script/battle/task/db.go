package task

import (
	"encoding/json"
	"time"

	"github.com/Masterminds/squirrel"

	db "github.com/Quieting/chowder/script/sql"
	"github.com/Quieting/chowder/script/xerror"
)

const taskTableName = "task"

type Task struct {
	Id          int64     `db:"id"`
	Name        string    `db:"name"`         // 任务名称
	Icon        string    `db:"icon"`         // 任务小图标
	AttrCat     string    `db:"attr_cat"`     // 任务属性类别 参考enum.TaskAttributeCat
	Category    string    `db:"category"`     // 任务分类 参考enum.TaskType
	TriggerType string    `db:"trigger_type"` // 触发类型 参考enum.TaskTriggerType
	FinishType  string    `db:"finish_type"`  // 完成类型 参考enum.TaskFinishType
	Status      int64     `db:"status"`       // 状态 0:正常 1:无效
	Stage       int64     `db:"stage"`        // 阶段(次数/次/个/天)
	Rewards     string    `db:"rewards"`      // 奖励配置
	Stages      []Stages  `db:"-"`
	EndTime     int64     `db:"end_time"`    // 结束时间，0 是永久
	StartTime   int64     `db:"start_time"`  // 开始时间 0 及时生效
	GrantType   int64     `db:"grant_type"`  // 奖励发放类型 0:领取发放 1:直接发放 2:邮箱发放
	UpdateTime  time.Time `db:"update_time"` // 更新时间
	CreateTime  time.Time `db:"create_time"` // 创建时间
}

type Stages struct {
	Stage   int64     `json:"stage"` // 阶段，签到任务有效值，1：星期一 2：星期二，以此类推
	Name    string    `json:"name"`
	Rewards []*Reward `json:"rewards"` // 奖励
}

type Reward struct {
	GoodsType         int64  `json:"goods_type"`          // 物品类型 1:道具 2:钻石 3:礼包
	GoodsId           int64  `json:"goods_id"`            // 物品id，物品类型为道具的时候有效
	GoodsName         string `json:"goods_name"`          // 物品名称
	GoodsPicUrl       string `json:"goods_pic_url"`       // 物品图片
	GoodsAnimationUrl string `json:"goods_animation_url"` // 物品动画地址
	GoodsDesc         string `json:"goods_desc"`          // 物品描述
	GoodsNum          int64  `json:"goods_num"`           // 数量
	Expired           int64  `json:"expired"`             // 过期时间，物品类型为道具的时候有效（单位是天）
	ExpiredHour       int64  `json:"expired_hour"`        // 单位小时，含义同Expired（为了兼容老版本,不能直接修改Expired，故添加该字段）
}

type ListArg struct {
	Limit       int64
	Offset      int64
	Id          int64  // id
	TriggerType string // 触发类型
	Category    string // 任务类型
	Class       string
	ChildClass  string
}

func (arg *ListArg) query(columns ...string) squirrel.SelectBuilder {
	builder := squirrel.Select(columns...).From(taskTableName)
	if arg.Limit > 0 {
		builder = builder.Limit(uint64(arg.Limit))
	}
	if arg.Offset > 0 {
		builder = builder.Offset(uint64(arg.Offset))
	}
	if arg.Id != 0 {
		builder = builder.Where(squirrel.Eq{"id": arg.Id})
	}

	if arg.TriggerType != "" {
		builder = builder.Where(squirrel.Eq{"trigger_type": arg.TriggerType})
	}

	if arg.Class != "" {
		builder = builder.Where(squirrel.Eq{"class": arg.Class})
	}

	if arg.ChildClass != "" {
		builder = builder.Where(squirrel.Eq{"child_class": arg.ChildClass})
	}

	if arg.Category != "" {
		builder = builder.Where(squirrel.Eq{"category": arg.Category})
	}

	return builder
}

func List(p ListArg) ([]Task, error) {
	propsFieldNames := db.RawFieldNames(&Task{})
	query := p.query(propsFieldNames...)

	str, args, err := query.ToSql()
	if err != nil {
		return nil, xerror.New(err, "组装sql语句失败")
	}

	list := make([]Task, 0, p.Limit)

	err = db.QueryRows(&list, str, args...)
	if err != nil {
		return nil, xerror.New(err, "查询任务列表失败")
	}

	for i, task := range list {
		stages := make([]Stages, 0)
		_ = json.Unmarshal([]byte(task.Rewards), &stages)
		list[i].Stages = stages
	}

	return list, nil
}

func LastTask() (*Task, error) {
	columns := db.RawFieldNames(&Task{})
	query := squirrel.Select(columns...).From(taskTableName).OrderBy("id DESC").Limit(1)
	str, args, err := query.ToSql()
	if err != nil {
		return nil, xerror.New(err, "组装sql语句失败")
	}

	mod := new(Task)

	err = db.QueryRow(mod, str, args...)
	if err != nil {
		return nil, xerror.New(err, "查询最近添加任务失败")
	}
	return mod, nil
}

func Delete(id int64) error {
	exec := squirrel.Delete(taskTableName).Where(squirrel.Eq{"id": id})
	str, args, err := exec.ToSql()
	if err != nil {
		return xerror.New(err, "删除任务：组装sql语句失败")
	}

	_, err = db.Exec(str, args...)
	if err != nil {
		return xerror.New(err, "删除任务：执行sql语句失败")
	}
	return nil
}
