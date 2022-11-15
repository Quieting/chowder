package task

import (
	"testing"
)

func TestList(t *testing.T) {
	got, err := List(ListArg{
		Limit:       10,
		Offset:      0,
		Id:          0,
		TriggerType: "",
	})
	if err != nil {
		t.Errorf("List() error = %v", err)
		return
	}

	for _, task := range got {
		t.Logf("%+v\n", task)

	}
}
