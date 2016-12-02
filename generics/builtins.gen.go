package generics


import (
	"fmt"
	"go/token"
	"reflect"
	"strings"
)

func To_string(x interface{}) string {
	return fmt.Sprint(x)
}

func Equal(x, y interface{}) bool {
	return reflect.DeepEqual(x, y)
}

func NotEqual(x, y interface{}) bool {
	return !reflect.DeepEqual(x, y)
}
// BinaryOpTokens are all the binary operations
var BinaryOpTokens = map[string]token.Token{
	"Equal": token.EQL,
	"NotEqual": token.NEQ,
	"Add": token.ADD,
	"Subtract": token.SUB,
	"Multiply": token.MUL,
	"Divide": token.QUO,
	"Remainder": token.REM,
	"BitwiseAnd": token.AND,
	"BitwiseOr": token.OR,
	"BitwiseXor": token.XOR,
	"BitClear": token.AND_NOT,
	"LeftShift": token.SHL,
	"RightShift": token.SHR,
}

// Add calls the '+' operator
func Add(x, y interface{}) interface{} {
    switch t := x.(type) {

        case complex128:
            return t + To_complex128(y)

        case complex64:
            return t + To_complex64(y)

        case float32:
            return t + To_float32(y)

        case float64:
            return t + To_float64(y)

        case int:
            return t + To_int(y)

        case int16:
            return t + To_int16(y)

        case int32:
            return t + To_int32(y)

        case int8:
            return t + To_int8(y)

        case string:
            return t + To_string(y)

        case uint:
            return t + To_uint(y)

        case uint16:
            return t + To_uint16(y)

        case uint64:
            return t + To_uint64(y)

        case uint8:
            return t + To_uint8(y)

        case uintptr:
            return t + To_uintptr(y)

    }
    panic(fmt.Sprintf("+ not implemented for %T", x))
}

// Subtract calls the '-' operator
func Subtract(x, y interface{}) interface{} {
    switch t := x.(type) {

        case complex128:
            return t - To_complex128(y)

        case complex64:
            return t - To_complex64(y)

        case float32:
            return t - To_float32(y)

        case float64:
            return t - To_float64(y)

        case int:
            return t - To_int(y)

        case int16:
            return t - To_int16(y)

        case int32:
            return t - To_int32(y)

        case int8:
            return t - To_int8(y)

        case uint:
            return t - To_uint(y)

        case uint16:
            return t - To_uint16(y)

        case uint64:
            return t - To_uint64(y)

        case uint8:
            return t - To_uint8(y)

        case uintptr:
            return t - To_uintptr(y)

    }
    panic(fmt.Sprintf("- not implemented for %T", x))
}

// Multiply calls the '*' operator
func Multiply(x, y interface{}) interface{} {
    switch t := x.(type) {

        case complex128:
            return t * To_complex128(y)

        case complex64:
            return t * To_complex64(y)

        case float32:
            return t * To_float32(y)

        case float64:
            return t * To_float64(y)

        case int:
            return t * To_int(y)

        case int16:
            return t * To_int16(y)

        case int32:
            return t * To_int32(y)

        case int8:
            return t * To_int8(y)

        case uint:
            return t * To_uint(y)

        case uint16:
            return t * To_uint16(y)

        case uint64:
            return t * To_uint64(y)

        case uint8:
            return t * To_uint8(y)

        case uintptr:
            return t * To_uintptr(y)

    }
    panic(fmt.Sprintf("* not implemented for %T", x))
}

// Divide calls the '/' operator
func Divide(x, y interface{}) interface{} {
    switch t := x.(type) {

        case complex128:
            return t / To_complex128(y)

        case complex64:
            return t / To_complex64(y)

        case float32:
            return t / To_float32(y)

        case float64:
            return t / To_float64(y)

        case int:
            return t / To_int(y)

        case int16:
            return t / To_int16(y)

        case int32:
            return t / To_int32(y)

        case int8:
            return t / To_int8(y)

        case uint:
            return t / To_uint(y)

        case uint16:
            return t / To_uint16(y)

        case uint64:
            return t / To_uint64(y)

        case uint8:
            return t / To_uint8(y)

        case uintptr:
            return t / To_uintptr(y)

    }
    panic(fmt.Sprintf("/ not implemented for %T", x))
}

// Remainder calls the '%' operator
func Remainder(x, y interface{}) interface{} {
    switch t := x.(type) {

        case int:
            return t % To_int(y)

        case int16:
            return t % To_int16(y)

        case int32:
            return t % To_int32(y)

        case int8:
            return t % To_int8(y)

        case uint:
            return t % To_uint(y)

        case uint16:
            return t % To_uint16(y)

        case uint64:
            return t % To_uint64(y)

        case uint8:
            return t % To_uint8(y)

        case uintptr:
            return t % To_uintptr(y)

    }
    panic(fmt.Sprintf("% not implemented for %T", x))
}

// BitwiseAnd calls the '&' operator
func BitwiseAnd(x, y interface{}) interface{} {
    switch t := x.(type) {

        case int:
            return t & To_int(y)

        case int16:
            return t & To_int16(y)

        case int32:
            return t & To_int32(y)

        case int8:
            return t & To_int8(y)

        case uint:
            return t & To_uint(y)

        case uint16:
            return t & To_uint16(y)

        case uint64:
            return t & To_uint64(y)

        case uint8:
            return t & To_uint8(y)

        case uintptr:
            return t & To_uintptr(y)

    }
    panic(fmt.Sprintf("& not implemented for %T", x))
}

// BitwiseOr calls the '|' operator
func BitwiseOr(x, y interface{}) interface{} {
    switch t := x.(type) {

        case int:
            return t | To_int(y)

        case int16:
            return t | To_int16(y)

        case int32:
            return t | To_int32(y)

        case int8:
            return t | To_int8(y)

        case uint:
            return t | To_uint(y)

        case uint16:
            return t | To_uint16(y)

        case uint64:
            return t | To_uint64(y)

        case uint8:
            return t | To_uint8(y)

        case uintptr:
            return t | To_uintptr(y)

    }
    panic(fmt.Sprintf("| not implemented for %T", x))
}

// BitwiseXor calls the '^' operator
func BitwiseXor(x, y interface{}) interface{} {
    switch t := x.(type) {

        case int:
            return t ^ To_int(y)

        case int16:
            return t ^ To_int16(y)

        case int32:
            return t ^ To_int32(y)

        case int8:
            return t ^ To_int8(y)

        case uint:
            return t ^ To_uint(y)

        case uint16:
            return t ^ To_uint16(y)

        case uint64:
            return t ^ To_uint64(y)

        case uint8:
            return t ^ To_uint8(y)

        case uintptr:
            return t ^ To_uintptr(y)

    }
    panic(fmt.Sprintf("^ not implemented for %T", x))
}

// BitClear calls the '%^' operator
func BitClear(x, y interface{}) interface{} {
    switch t := x.(type) {

        case int:
            return t %^ To_int(y)

        case int16:
            return t %^ To_int16(y)

        case int32:
            return t %^ To_int32(y)

        case int8:
            return t %^ To_int8(y)

        case uint:
            return t %^ To_uint(y)

        case uint16:
            return t %^ To_uint16(y)

        case uint64:
            return t %^ To_uint64(y)

        case uint8:
            return t %^ To_uint8(y)

        case uintptr:
            return t %^ To_uintptr(y)

    }
    panic(fmt.Sprintf("%^ not implemented for %T", x))
}

// LeftShift calls the '<<' operator
func LeftShift(x, y interface{}) interface{} {
    switch t := x.(type) {

        case int:
            return t << To_uint(y)

        case int16:
            return t << To_uint(y)

        case int32:
            return t << To_uint(y)

        case int8:
            return t << To_uint(y)

        case uint:
            return t << To_uint(y)

        case uint16:
            return t << To_uint(y)

        case uint64:
            return t << To_uint(y)

        case uint8:
            return t << To_uint(y)

        case uintptr:
            return t << To_uint(y)

    }
    panic(fmt.Sprintf("<< not implemented for %T", x))
}

// RightShift calls the '>>' operator
func RightShift(x, y interface{}) interface{} {
    switch t := x.(type) {

        case int:
            return t >> To_uint(y)

        case int16:
            return t >> To_uint(y)

        case int32:
            return t >> To_uint(y)

        case int8:
            return t >> To_uint(y)

        case uint:
            return t >> To_uint(y)

        case uint16:
            return t >> To_uint(y)

        case uint64:
            return t >> To_uint(y)

        case uint8:
            return t >> To_uint(y)

        case uintptr:
            return t >> To_uint(y)

    }
    panic(fmt.Sprintf(">> not implemented for %T", x))
}



// To_complex128 converts anything to a complex128
func To_complex128(x interface{}) complex128 {
	switch t := x.(type) {

		case complex128:
			
				
					return t
				
			

		case complex64:
			
				
					return complex(float64(real(t)), float64(imag(t)))
				
			

		case float32:
			
				
					return complex(To_float64(t), 0)
				
			

		case float64:
			
				
					return complex(To_float64(t), 0)
				
			

		case int:
			
				
					return complex(To_float64(t), 0)
				
			

		case int16:
			
				
					return complex(To_float64(t), 0)
				
			

		case int32:
			
				
					return complex(To_float64(t), 0)
				
			

		case int8:
			
				
					return complex(To_float64(t), 0)
				
			

		case uint:
			
				
					return complex(To_float64(t), 0)
				
			

		case uint16:
			
				
					return complex(To_float64(t), 0)
				
			

		case uint64:
			
				
					return complex(To_float64(t), 0)
				
			

		case uint8:
			
				
					return complex(To_float64(t), 0)
				
			

		case uintptr:
			
				
					return complex(To_float64(t), 0)
				
			

	}
	var o complex128
	fmt.Fscanln(strings.NewReader(fmt.Sprint(x)), &o)
	return o
}

// To_complex64 converts anything to a complex64
func To_complex64(x interface{}) complex64 {
	switch t := x.(type) {

		case complex128:
			
				
					return complex(float32(real(t)), float32(imag(t)))
				
			

		case complex64:
			
				
					return t
				
			

		case float32:
			
				
					return complex(To_float32(t), 0)
				
			

		case float64:
			
				
					return complex(To_float32(t), 0)
				
			

		case int:
			
				
					return complex(To_float32(t), 0)
				
			

		case int16:
			
				
					return complex(To_float32(t), 0)
				
			

		case int32:
			
				
					return complex(To_float32(t), 0)
				
			

		case int8:
			
				
					return complex(To_float32(t), 0)
				
			

		case uint:
			
				
					return complex(To_float32(t), 0)
				
			

		case uint16:
			
				
					return complex(To_float32(t), 0)
				
			

		case uint64:
			
				
					return complex(To_float32(t), 0)
				
			

		case uint8:
			
				
					return complex(To_float32(t), 0)
				
			

		case uintptr:
			
				
					return complex(To_float32(t), 0)
				
			

	}
	var o complex64
	fmt.Fscanln(strings.NewReader(fmt.Sprint(x)), &o)
	return o
}

// To_float32 converts anything to a float32
func To_float32(x interface{}) float32 {
	switch t := x.(type) {

		case complex128:
			
				
					return float32(real(t))
				
			

		case complex64:
			
				
					return float32(real(t))
				
			

		case float32:
			
				
					return float32(t)
				
			

		case float64:
			
				
					return float32(t)
				
			

		case int:
			
				
					return float32(t)
				
			

		case int16:
			
				
					return float32(t)
				
			

		case int32:
			
				
					return float32(t)
				
			

		case int8:
			
				
					return float32(t)
				
			

		case uint:
			
				
					return float32(t)
				
			

		case uint16:
			
				
					return float32(t)
				
			

		case uint64:
			
				
					return float32(t)
				
			

		case uint8:
			
				
					return float32(t)
				
			

		case uintptr:
			
				
					return float32(t)
				
			

	}
	var o float32
	fmt.Fscanln(strings.NewReader(fmt.Sprint(x)), &o)
	return o
}

// To_float64 converts anything to a float64
func To_float64(x interface{}) float64 {
	switch t := x.(type) {

		case complex128:
			
				
					return float64(real(t))
				
			

		case complex64:
			
				
					return float64(real(t))
				
			

		case float32:
			
				
					return float64(t)
				
			

		case float64:
			
				
					return float64(t)
				
			

		case int:
			
				
					return float64(t)
				
			

		case int16:
			
				
					return float64(t)
				
			

		case int32:
			
				
					return float64(t)
				
			

		case int8:
			
				
					return float64(t)
				
			

		case uint:
			
				
					return float64(t)
				
			

		case uint16:
			
				
					return float64(t)
				
			

		case uint64:
			
				
					return float64(t)
				
			

		case uint8:
			
				
					return float64(t)
				
			

		case uintptr:
			
				
					return float64(t)
				
			

	}
	var o float64
	fmt.Fscanln(strings.NewReader(fmt.Sprint(x)), &o)
	return o
}

// To_int converts anything to a int
func To_int(x interface{}) int {
	switch t := x.(type) {

		case complex128:
			
				
					return int(real(t))
				
			

		case complex64:
			
				
					return int(real(t))
				
			

		case float32:
			
				
					return int(t)
				
			

		case float64:
			
				
					return int(t)
				
			

		case int:
			
				
					return int(t)
				
			

		case int16:
			
				
					return int(t)
				
			

		case int32:
			
				
					return int(t)
				
			

		case int8:
			
				
					return int(t)
				
			

		case uint:
			
				
					return int(t)
				
			

		case uint16:
			
				
					return int(t)
				
			

		case uint64:
			
				
					return int(t)
				
			

		case uint8:
			
				
					return int(t)
				
			

		case uintptr:
			
				
					return int(t)
				
			

	}
	var o int
	fmt.Fscanln(strings.NewReader(fmt.Sprint(x)), &o)
	return o
}

// To_int16 converts anything to a int16
func To_int16(x interface{}) int16 {
	switch t := x.(type) {

		case complex128:
			
				
					return int16(real(t))
				
			

		case complex64:
			
				
					return int16(real(t))
				
			

		case float32:
			
				
					return int16(t)
				
			

		case float64:
			
				
					return int16(t)
				
			

		case int:
			
				
					return int16(t)
				
			

		case int16:
			
				
					return int16(t)
				
			

		case int32:
			
				
					return int16(t)
				
			

		case int8:
			
				
					return int16(t)
				
			

		case uint:
			
				
					return int16(t)
				
			

		case uint16:
			
				
					return int16(t)
				
			

		case uint64:
			
				
					return int16(t)
				
			

		case uint8:
			
				
					return int16(t)
				
			

		case uintptr:
			
				
					return int16(t)
				
			

	}
	var o int16
	fmt.Fscanln(strings.NewReader(fmt.Sprint(x)), &o)
	return o
}

// To_int32 converts anything to a int32
func To_int32(x interface{}) int32 {
	switch t := x.(type) {

		case complex128:
			
				
					return int32(real(t))
				
			

		case complex64:
			
				
					return int32(real(t))
				
			

		case float32:
			
				
					return int32(t)
				
			

		case float64:
			
				
					return int32(t)
				
			

		case int:
			
				
					return int32(t)
				
			

		case int16:
			
				
					return int32(t)
				
			

		case int32:
			
				
					return int32(t)
				
			

		case int8:
			
				
					return int32(t)
				
			

		case uint:
			
				
					return int32(t)
				
			

		case uint16:
			
				
					return int32(t)
				
			

		case uint64:
			
				
					return int32(t)
				
			

		case uint8:
			
				
					return int32(t)
				
			

		case uintptr:
			
				
					return int32(t)
				
			

	}
	var o int32
	fmt.Fscanln(strings.NewReader(fmt.Sprint(x)), &o)
	return o
}

// To_int8 converts anything to a int8
func To_int8(x interface{}) int8 {
	switch t := x.(type) {

		case complex128:
			
				
					return int8(real(t))
				
			

		case complex64:
			
				
					return int8(real(t))
				
			

		case float32:
			
				
					return int8(t)
				
			

		case float64:
			
				
					return int8(t)
				
			

		case int:
			
				
					return int8(t)
				
			

		case int16:
			
				
					return int8(t)
				
			

		case int32:
			
				
					return int8(t)
				
			

		case int8:
			
				
					return int8(t)
				
			

		case uint:
			
				
					return int8(t)
				
			

		case uint16:
			
				
					return int8(t)
				
			

		case uint64:
			
				
					return int8(t)
				
			

		case uint8:
			
				
					return int8(t)
				
			

		case uintptr:
			
				
					return int8(t)
				
			

	}
	var o int8
	fmt.Fscanln(strings.NewReader(fmt.Sprint(x)), &o)
	return o
}

// To_uint converts anything to a uint
func To_uint(x interface{}) uint {
	switch t := x.(type) {

		case complex128:
			
				
					return uint(real(t))
				
			

		case complex64:
			
				
					return uint(real(t))
				
			

		case float32:
			
				
					return uint(t)
				
			

		case float64:
			
				
					return uint(t)
				
			

		case int:
			
				
					return uint(t)
				
			

		case int16:
			
				
					return uint(t)
				
			

		case int32:
			
				
					return uint(t)
				
			

		case int8:
			
				
					return uint(t)
				
			

		case uint:
			
				
					return uint(t)
				
			

		case uint16:
			
				
					return uint(t)
				
			

		case uint64:
			
				
					return uint(t)
				
			

		case uint8:
			
				
					return uint(t)
				
			

		case uintptr:
			
				
					return uint(t)
				
			

	}
	var o uint
	fmt.Fscanln(strings.NewReader(fmt.Sprint(x)), &o)
	return o
}

// To_uint16 converts anything to a uint16
func To_uint16(x interface{}) uint16 {
	switch t := x.(type) {

		case complex128:
			
				
					return uint16(real(t))
				
			

		case complex64:
			
				
					return uint16(real(t))
				
			

		case float32:
			
				
					return uint16(t)
				
			

		case float64:
			
				
					return uint16(t)
				
			

		case int:
			
				
					return uint16(t)
				
			

		case int16:
			
				
					return uint16(t)
				
			

		case int32:
			
				
					return uint16(t)
				
			

		case int8:
			
				
					return uint16(t)
				
			

		case uint:
			
				
					return uint16(t)
				
			

		case uint16:
			
				
					return uint16(t)
				
			

		case uint64:
			
				
					return uint16(t)
				
			

		case uint8:
			
				
					return uint16(t)
				
			

		case uintptr:
			
				
					return uint16(t)
				
			

	}
	var o uint16
	fmt.Fscanln(strings.NewReader(fmt.Sprint(x)), &o)
	return o
}

// To_uint64 converts anything to a uint64
func To_uint64(x interface{}) uint64 {
	switch t := x.(type) {

		case complex128:
			
				
					return uint64(real(t))
				
			

		case complex64:
			
				
					return uint64(real(t))
				
			

		case float32:
			
				
					return uint64(t)
				
			

		case float64:
			
				
					return uint64(t)
				
			

		case int:
			
				
					return uint64(t)
				
			

		case int16:
			
				
					return uint64(t)
				
			

		case int32:
			
				
					return uint64(t)
				
			

		case int8:
			
				
					return uint64(t)
				
			

		case uint:
			
				
					return uint64(t)
				
			

		case uint16:
			
				
					return uint64(t)
				
			

		case uint64:
			
				
					return uint64(t)
				
			

		case uint8:
			
				
					return uint64(t)
				
			

		case uintptr:
			
				
					return uint64(t)
				
			

	}
	var o uint64
	fmt.Fscanln(strings.NewReader(fmt.Sprint(x)), &o)
	return o
}

// To_uint8 converts anything to a uint8
func To_uint8(x interface{}) uint8 {
	switch t := x.(type) {

		case complex128:
			
				
					return uint8(real(t))
				
			

		case complex64:
			
				
					return uint8(real(t))
				
			

		case float32:
			
				
					return uint8(t)
				
			

		case float64:
			
				
					return uint8(t)
				
			

		case int:
			
				
					return uint8(t)
				
			

		case int16:
			
				
					return uint8(t)
				
			

		case int32:
			
				
					return uint8(t)
				
			

		case int8:
			
				
					return uint8(t)
				
			

		case uint:
			
				
					return uint8(t)
				
			

		case uint16:
			
				
					return uint8(t)
				
			

		case uint64:
			
				
					return uint8(t)
				
			

		case uint8:
			
				
					return uint8(t)
				
			

		case uintptr:
			
				
					return uint8(t)
				
			

	}
	var o uint8
	fmt.Fscanln(strings.NewReader(fmt.Sprint(x)), &o)
	return o
}

// To_uintptr converts anything to a uintptr
func To_uintptr(x interface{}) uintptr {
	switch t := x.(type) {

		case complex128:
			
				
					return uintptr(real(t))
				
			

		case complex64:
			
				
					return uintptr(real(t))
				
			

		case float32:
			
				
					return uintptr(t)
				
			

		case float64:
			
				
					return uintptr(t)
				
			

		case int:
			
				
					return uintptr(t)
				
			

		case int16:
			
				
					return uintptr(t)
				
			

		case int32:
			
				
					return uintptr(t)
				
			

		case int8:
			
				
					return uintptr(t)
				
			

		case uint:
			
				
					return uintptr(t)
				
			

		case uint16:
			
				
					return uintptr(t)
				
			

		case uint64:
			
				
					return uintptr(t)
				
			

		case uint8:
			
				
					return uintptr(t)
				
			

		case uintptr:
			
				
					return uintptr(t)
				
			

	}
	var o uintptr
	fmt.Fscanln(strings.NewReader(fmt.Sprint(x)), &o)
	return o
}


