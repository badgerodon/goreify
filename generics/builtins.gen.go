package generics


import (
	"fmt"
	"go/token"
)
// BinaryOpTokens are all the binary operations
var BinaryOpTokens = map[string]token.Token{
"Sum": token.ADD,
"Difference": token.SUB,
"Product": token.MUL,
}

// Sum calls the '+' operator
func Sum(x, y interface{}) interface{} {
    switch x.(type) {

        case complex128:
            return x.(complex128) + y.(complex128)

        case complex64:
            return x.(complex64) + y.(complex64)

        case float32:
            return x.(float32) + y.(float32)

        case float64:
            return x.(float64) + y.(float64)

        case int:
            return x.(int) + y.(int)

        case int16:
            return x.(int16) + y.(int16)

        case int32:
            return x.(int32) + y.(int32)

        case int8:
            return x.(int8) + y.(int8)

        case string:
            return x.(string) + y.(string)

        case uint:
            return x.(uint) + y.(uint)

        case uint16:
            return x.(uint16) + y.(uint16)

        case uint64:
            return x.(uint64) + y.(uint64)

        case uint8:
            return x.(uint8) + y.(uint8)

        case uintptr:
            return x.(uintptr) + y.(uintptr)

    }
    panic(fmt.Sprintf("+ not implemented for %T", x))
}

// Difference calls the '-' operator
func Difference(x, y interface{}) interface{} {
    switch x.(type) {

        case complex128:
            return x.(complex128) - y.(complex128)

        case complex64:
            return x.(complex64) - y.(complex64)

        case float32:
            return x.(float32) - y.(float32)

        case float64:
            return x.(float64) - y.(float64)

        case int:
            return x.(int) - y.(int)

        case int16:
            return x.(int16) - y.(int16)

        case int32:
            return x.(int32) - y.(int32)

        case int8:
            return x.(int8) - y.(int8)

        case uint:
            return x.(uint) - y.(uint)

        case uint16:
            return x.(uint16) - y.(uint16)

        case uint64:
            return x.(uint64) - y.(uint64)

        case uint8:
            return x.(uint8) - y.(uint8)

        case uintptr:
            return x.(uintptr) - y.(uintptr)

    }
    panic(fmt.Sprintf("- not implemented for %T", x))
}

// Product calls the '*' operator
func Product(x, y interface{}) interface{} {
    switch x.(type) {

        case complex128:
            return x.(complex128) * y.(complex128)

        case complex64:
            return x.(complex64) * y.(complex64)

        case float32:
            return x.(float32) * y.(float32)

        case float64:
            return x.(float64) * y.(float64)

        case int:
            return x.(int) * y.(int)

        case int16:
            return x.(int16) * y.(int16)

        case int32:
            return x.(int32) * y.(int32)

        case int8:
            return x.(int8) * y.(int8)

        case uint:
            return x.(uint) * y.(uint)

        case uint16:
            return x.(uint16) * y.(uint16)

        case uint64:
            return x.(uint64) * y.(uint64)

        case uint8:
            return x.(uint8) * y.(uint8)

        case uintptr:
            return x.(uintptr) * y.(uintptr)

    }
    panic(fmt.Sprintf("* not implemented for %T", x))
}


