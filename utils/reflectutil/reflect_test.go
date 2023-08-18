package reflectutil

import (
	"reflect"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestIndirectType(t *testing.T) {
	v := Indirect(&User{})
	k := IndirectKind(&User{})
	ty := IndirectType(&User{})
	t.Logf("%+v\n", v) // {Name: Age:0}
	t.Logf("%+v\n", k) // struct
	t.Logf("%+v", ty)  // reflectutil.User
}
func TestNewData(t *testing.T) {
	ty := IndirectType(&User{})
	data := NewData(ty)
	t.Logf("%+v", data) // &{Name: Age:0}
}

func TestGetName(t *testing.T) {
	name := GetName(&User{})
	t.Log(name) // User
}

func TestIsPointer(t *testing.T) {
	t.Log(IsPointer(&User{})) // true
	t.Log(IsPointer(User{}))  // false
}

func TestConvertType(t *testing.T) {
	u := [2]any{"hello", 18}
	v1 := reflect.ValueOf(u)
	v2 := reflect.ValueOf(&User{})

	for i := range u {
		convertType := ConvertType(v1, i, v2, i)
		v2.Elem().Field(i).Set(convertType)
	}

	t.Log(v2) // &{hello 18}
}
