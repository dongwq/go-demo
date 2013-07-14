/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-1-10，上午10:00
 * @version 0.1
 */
package usego

import (
	"fmt"
	"reflect"
)

type MvcBase struct{

}
func (b * MvcBase) Show(){
	fmt.Println("show this ")
}
func (b * MvcBase) Edit(){
	fmt.Println("show this ")
}
func (b  MvcBase) Put(){
	fmt.Println("show this ")
}
func (b  MvcBase) Get(){
	fmt.Println("show this ")
}

func TestReflect() {
	var a MvcBase;

	mvcBase := reflect.ValueOf(a)
	numMethods := mvcBase.NumMethod()

	fmt.Println("num = %v" , numMethods)

//	for i:=0;i < numMethods;i++{
//		m := mvcBase.Method(i).Type()
//		fmt.Println("method name=%s",m.Name)
//	}

}

