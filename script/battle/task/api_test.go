package task

import (
	"reflect"
	"testing"
)

func TestMgrTaskList(t *testing.T) {
	type args struct {
		req *MgrTaskListReq
	}
	tests := []struct {
		name    string
		args    args
		want    []MgrTaskListItem
		want1   int64
		wantErr bool
	}{
		{
			name: "任务列表: 按前端分类过滤",
			args: args{
				req: &MgrTaskListReq{
					Class: "daily",
				},
			},
		},
		{
			name: "任务列表: 按前端分类过滤",
			args: args{
				req: &MgrTaskListReq{
					ChildClass: "newer_first",
				},
			},
		},
		{
			name: "任务列表: 任务类别过滤",
			args: args{
				req: &MgrTaskListReq{
					Category: "daily",
				},
			},
		},

		{
			name: "任务列表: 触发类型过滤",
			args: args{
				req: &MgrTaskListReq{
					TriggerType: "sign_in",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := MgrTaskList(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("MgrTaskList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			req := tt.args.req
			for _, val := range got {
				if req.Class != "" && val.Class != req.Class {
					t.Errorf("class 过滤条件未生效")
				}
				if req.ChildClass != "" && val.ChildClass != req.ChildClass {
					t.Errorf("child_class 过滤条件未生效")
				}

				if req.AttrCat != "" && val.AttrCat != req.AttrCat {
					t.Errorf("attr_cat 过滤条件未生效")
				}

				if req.Category != "" && val.Category != req.Category {
					t.Errorf("category 过滤条件未生效")
				}

				if req.TriggerType != "" && val.TriggerType != req.TriggerType {
					t.Errorf("trigger_type 过滤条件未生效")
				}

				if req.Name != "" && val.Name != req.Name {
					t.Errorf("name 条件未生效")
				}

				// t.Logf("item: %+v\n", val)
			}
			// t.Logf("total: %d\n", got1)
		})
	}
}

func TestMgrAddTask(t *testing.T) {
	type args struct {
		req *MgrTaskAddReq
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "示例1",
			args: args{
				req: &MgrTaskAddReq{
					Name:           "测试任务",
					Icon:           "",
					AttrCat:        "user_task",
					Class:          "daily",
					ChildClass:     "",
					Category:       "daily",
					TriggerType:    "open_app",
					FinishType:     "",
					Status:         0,
					Stage:          1,
					Weight:         10,
					StartTime:      100,
					EndTime:        100,
					JumpParams:     "home",
					JumpPath:       "home",
					GrantType:      0,
					Rewards:        "[]",
					RelationTaskId: 10,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := MgrAddTask(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("MgrAddTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := LastTask()
			if err != nil {
				t.Errorf("获取最近添加任务失败")
			}

			want := tt.args.req.TaskMod()
			want.Id = got.Id
			want.CreateTime = got.CreateTime
			want.UpdateTime = got.UpdateTime
			if !reflect.DeepEqual(got, want) {
				t.Errorf("添加数据和预期不一致,\ngot:%+v,\n want:%+v\n", got, want)
			}

			Delete(got.Id)

		})
	}
}
