// ref1
package ref

import (
	"fmt"
	"reflect"
)

//type MyStruct struct {
//	name string
//}

//func (this *MyStruct) GetName() string {
//	return this.name
//}

func Ref1() {
	s := "this is string"
	fmt.Println(reflect.TypeOf(s))
	fmt.Println("-------------------")

	fmt.Println(reflect.ValueOf(s))
	var x float64 = 3.4
	fmt.Println(reflect.ValueOf(x))

	y := 33
	fmt.Println(reflect.ValueOf(y))
	fmt.Println("-------------------")

	a := new(MyStruct)
	a.name = "yejianfeng"
	typ := reflect.TypeOf(a)

	fmt.Println(typ.NumMethod())
	fmt.Println("-------------------")

	b := reflect.ValueOf(a).MethodByName("GetName").Call([]reflect.Value{})
	fmt.Println(b[0])

}
