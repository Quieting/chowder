package task

import (
	"github.com/Quieting/chowder/copier"
	"github.com/Quieting/chowder/script/http"
)

var urlPrefix = "http://localhost:50000"
var mgrToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Njc4OTc0OTUsImlhdCI6MTY2NTMwNTQ5NSwidXNlcklkIjoyfQ.jHG3b9CmGvWFVDGUD4XfOTu9_mCAtPygaC2P7qz3gBI"
var apiToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjgxMzcwMjIsImlhdCI6MTY2NTU0NTAyMiwidXNlcklkIjoxMDEyMjk1NX0.CfFi-RbWgYMdz5OezYgYZLeHyBMRW6tNZTqwDs6Jd0c"

type MgrTaskListReq struct {
	Page        int    `form:"page,optional,default=1"`  // 页码
	Size        int64  `form:"size,optional,default=10"` // 限制条数（默认为10）
	Name        string `form:"name,optional"`            // 任务名
	TriggerType string `form:"trigger_type,optional"`    // 触发类型
	AttrCat     string `form:"attr_cat,optional"`        // 任务属性类别 参考 enum.TaskAttributeCat
	Category    string `form:"category,optional"`        // 任务分类 参考 enum.TaskType
	Status      int64  `form:"status,optional"`          // 任务状态
	Class       string `form:"class,optional"`           // 前端分类
	ChildClass  string `form:"child_class,optional"`     // 前端子分类
}

func (req *MgrTaskListReq) path() string {
	vals := http.PathsValues(req)
	return vals.Encode()
}

type MgrTaskListResp struct {
	Total int64              `json:"total"` // 总数
	List  []*MgrTaskListItem `json:"list"`  // 数据
}

type MgrTaskListItem struct {
	Id              int64    `json:"id"`
	Name            string   `json:"name"`              // 任务名称
	Icon            string   `json:"icon"`              // 任务小图标
	AttrCat         string   `json:"attr_cat"`          // 任务属性类别 参考enum.TaskAttributeCat
	Category        string   `json:"category"`          // 任务分类 参考设计方案http://139.155.244.153:8090/pages/viewpage.action?pageId=8981563
	Class           string   `json:"class"`             // 任务前端分类
	ClassName       string   `json:"class_name"`        // 任务前端分类名字
	ChildClass      string   `json:"child_class"`       // 任务前端子分类
	ChildClassName  string   `json:"child_class_name"`  // 任务前端子分类
	TriggerType     string   `json:"trigger_type"`      // 触发类型 参考设计方案http://139.155.244.153:8090/pages/viewpage.action?pageId=8981563
	TriggerTypeName string   `json:"trigger_type_name"` // 触发类型名字
	FinishType      string   `json:"finish_type"`       // 完成类型 参考设计方案http://139.155.244.153:8090/pages/viewpage.action?pageId=8981563
	Status          int64    `json:"status"`            // 状态 0:正常 1:无效
	Stage           int64    `json:"stage"`             // 阶段(次数/次/个/天)
	Rewards         string   `json:"rewards"`           // 奖励配置 json,参考设计方案http://139.155.244.153:8090/pages/viewpage.action?pageId=8981575
	Stages          []Stages `json:"-"`
	StartTime       int64    `json:"start_time"`  // 开始时间 0 及时生效
	EndTime         int64    `json:"end_time"`    // 结束时间，0 是永久
	GrantType       int64    `json:"grant_type"`  // 奖励发放类型 0:领取发放 1:直接发放 2:邮箱发放
	Weight          int64    `json:"weight"`      // 权重，越大前端显示越靠前
	JumpParams      string   `json:"jump_params"` // 跳转参数
	JumpPath        string   `json:"jump_path"`   // 跳转路由
	UpdateTime      int64    `json:"update_time"` // 更新时间
	CreateTime      int64    `json:"create_time"` // 创建时间
}

func MgrTaskList(req *MgrTaskListReq) ([]MgrTaskListItem, int64, error) {
	url := urlPrefix + "/mgr/task/list"
	resp := &struct {
		List  []MgrTaskListItem `json:"list"`
		Total int64             `json:"total"`
	}{}
	err := http.Get(url, req, &resp, mgrToken)
	if err != nil {
		return nil, 0, err
	}
	return resp.List, resp.Total, nil
}

type MgrTaskAddReq struct {
	Name           string `json:"name"`                      // 任务名称
	Icon           string `json:"icon,optional"`             // 任务小图标
	AttrCat        string `json:"attr_cat,optional"`         // 任务属性类别 参考 enum.TaskAttributeCat
	Class          string `json:"class"`                     // 任务分类
	ChildClass     string `json:"child_class,optional"`      // 任务子分类
	Category       string `json:"category"`                  // 任务分类 参考设计方案http://139.155.244.153:8090/pages/viewpage.action?pageId=8981563
	TriggerType    string `json:"trigger_type"`              // 触发类型 参考设计方案http://139.155.244.153:8090/pages/viewpage.action?pageId=8981563
	FinishType     string `json:"finish_type"`               // 完成类型 参考设计方案http://139.155.244.153:8090/pages/viewpage.action?pageId=8981563
	Status         int64  `json:"status"`                    // 状态 0:正常 1:无效
	Stage          int64  `json:"stage"`                     // 阶段(次数/次/个/天)
	Weight         int64  `json:"weight,optional"`           // 权重，越大显示越靠前
	StartTime      int64  `json:"start_time,optional"`       // 开始时间 0 及时生效
	EndTime        int64  `json:"end_time,optional"`         // 结束时间，0 是永久
	JumpParams     string `json:"jump_params,optional"`      // 跳转参数
	JumpPath       string `json:"jump_path,optional"`        // 跳转路由
	GrantType      int64  `json:"grant_type"`                // 奖励发放类型 0:领取发放 1:直接发放 2:邮箱发放
	Rewards        string `json:"rewards"`                   // 奖励配置 json,参考设计方案http://139.155.244.153:8090/pages/viewpage.action?pageId=8981575
	RelationTaskId int64  `json:"relation_task_id,optional"` // 关联任务id
}

func (req *MgrTaskAddReq) TaskMod() *Task {
	mod := new(Task)
	_ = copier.Copy(req, mod)
	return mod
}

func MgrAddTask(req *MgrTaskAddReq) (int64, error) {
	url := urlPrefix + "/mgr/task/add"
	err := http.Post(url, req, nil, mgrToken)
	if err != nil {
		return 0, err
	}
	return 0, nil
}
