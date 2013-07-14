package basic

import (
    "fmt"
)

func DemoTypeCast() {
    var general interface{}
    general = 6.6
    type_cast(general)
    general = 2
    type_cast(general)
}


func type_cast(general interface{}) {
    switch general.(type) {
    case int :
        fmt.Println("the general type is int")
        newInt, ok := general.(int)
        check_convert(ok)
        fmt.Println("newInt 的值本来是", newInt)
        newInt += 2
        fmt.Println("加2后，结果是", newInt)
        newInt -= 6
        fmt.Println("接着减6后，结果是", newInt)
        newInt *= 4
        fmt.Println("然后乘4，结果是", newInt)
        newInt /= 3
        fmt.Println("最后除3，结果是", newInt)
        fmt.Println()
        fmt.Println()


    case float32:
        fmt.Println("the general type is float32")
        newFloat32, ok := general.(float32)
        check_convert(ok)
        fmt.Println("newFloat32 的值本来是", newFloat32)
        newFloat32 += 2.0
        fmt.Println("加2.0后，结果是", newFloat32)
        newFloat32 -= 6.0
        fmt.Println("接着减6.0后，结果是", newFloat32)
        newFloat32 *= 4.0
        fmt.Println("然后乘4.0，结果是", newFloat32)
        newFloat32 /= 3.0
        fmt.Println("最后除3.0，结果是", newFloat32)
        fmt.Println()
        fmt.Println()

    case float64:
        fmt.Println("the general type is float64")
        newFloat64, ok := general.(float64)
        check_convert(ok)
        fmt.Println("newFloat64 的值本来是", newFloat64)
        newFloat64 += 2.0
        fmt.Println("加2.0后，结果是", newFloat64)
        newFloat64 -= 6.0
        fmt.Println("接着减6.0后，结果是", newFloat64)
        newFloat64 *= 4.0
        fmt.Println("然后乘4.0，结果是", newFloat64)
        newFloat64 /= 3.0
        fmt.Println("最后除3.0，结果是", newFloat64)
        fmt.Println()
        fmt.Println()

    default:
        fmt.Println("unknown type")
    }
}

func check_convert(ok bool) {
    if false == ok {
        panic("type cast failed!")
    }
}
