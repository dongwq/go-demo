

package ref

import (
	"fmt"
)

func aFunc(){
	fmt.Println("hello a")
}

func bfunc(a, b int){
	fmt.Println("hello a")
}


//func Call(m map[string]interface{}, name string, params ... interface{}) (result []reflect.Value, err error) {
//    f = reflect.ValueOf(m[name])
//    if len(params) != f.Type().NumIn() {
//        err = errors.New("The number of params is not adapted.")
//        return
//    }
//    in := make([]reflect.Value, len(params))
//    for k, param := range params {
//        in[k] = reflect.ValueOf(param)
//    }
//    result = f.Call(in)
//    return
//}


func DemoMapFunc(){
	
	//mapFuncs := map[string]interface{}{"aFunc":aFunc}
	
	fmt.Println("hello")
	////mapFuncs := map[string]func(a... int){"a":a, "b",b}
	//mapFuncs["aFunc"]()
	
}

