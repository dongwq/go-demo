package main

import(
	"fmt"
	"reflect"
)

type User struct{
	Id int
	Name string
	Age int
}


func SetV(){
	x := 123
	v := reflect.ValueOf(&x)
	
	v.Elem().SetInt(789)
	
	fmt.Println(x)
}



func SetU(){
	u := User{2,"ww",30}
	ChangeU(&u)
	fmt.Println("%v", u)
}

func ChangeU(o interface{}){
	
}


func main(){
	fmt.Println("reflect demo")
	//SetV()
	SetU()	
	
}