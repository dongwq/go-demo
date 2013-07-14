// 
package ref

import (
	"fmt"
	"reflect"
)

type User struct {
	Id    int
	Name  string
	Age   int
	Title string
}

func (this User) Test(x int) {
	println("User.Test:", this.Name, x)
}

func test(i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	// Type Kind
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println(k)
		return
	}
	// Type Name
	println("Type:", t.Name())
	// Fields
	println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		fmt.Printf(" %6s: %v = %v\n", field.Name, field.Type, value)
	}
	// Methods
	println("Methods:")
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf(" %s: %v\n", method.Name, method.Type)
	}
}

func RefAction() {
	u := User{1, "Tom", 30, "Programmer"}
	test(u)
}
