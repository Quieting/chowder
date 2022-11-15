package mall

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPropList(t *testing.T) {
	got, err := PropList(PropListArg{
		Limit:  1,
		Offset: 2,
	})
	if err != nil {
		t.Errorf("%v\n", err)
	}
	t.Logf("list: %+v", got)
}

func TestOneProp(t *testing.T) {
	got, err := OneProp(117)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	t.Logf("propinfo: %+v", got)
}

func TestGoodsList(t *testing.T) {
	list, err := GoodsList()
	if err != nil {
		t.Errorf("%v\n", err)
	}
	t.Logf("propinfo: %+v", list)
}

func TestAny(t *testing.T) {
	var i *int
	i = new(int)
	*i = 2

	t.Logf("type: %s\n", reflect.ValueOf(i).Kind())

	v := new(struct {
		i int
	})

	val := reflect.ValueOf(v).Elem().Field(0)
	t.Logf("type: %s\n", val.Addr().Kind())

	var arr1 []int
	a0 := reflect.ValueOf(&arr1).Elem()

	e0 := make([]reflect.Value, 0)
	e0 = append(e0, reflect.ValueOf(100))
	valArr1 := reflect.Append(a0, e0...)

	a0.Set(valArr1)
	fmt.Println(valArr1)

	var sli []int
	val0 := reflect.ValueOf(&sli).Elem()

	val0 = reflect.Append(val0, reflect.ValueOf(100))
	t.Logf("val0: %v\n", val0)
	t.Logf("sli: %v\n", sli)

}
