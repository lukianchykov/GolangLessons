package main

import (
	"fmt"
	"unsafe"
)

// Переменная - хранилище наших значений
// Типы (примитивные)
// Объявление переменных: var
// Нулевые значения
// Объявление с инициализацией, short syntax
// Области видимости переменных
// Константы
// Преобразование типов(простые)
// Вывод переменных в консоль

// MAIN TYPES
// string
// bool
// int
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte - alias for uint8
// rune - alias for int32
// float32 float64
// complex64 complex128

var name string

const constName = "Андрей"

func main() {
	fmt.Println(name)
	fmt.Printf("Type: %T Value: %v\n", name, name)

	name := "hello"
	fmt.Println(name)

	fmt.Println(constName)

	var hello string
	fmt.Println(hello)

	hello = fmt.Sprintf("hello %s", constName)
	fmt.Println(hello)
	fmt.Printf("Type: %T Value: %v\n", hello, hello)

	fmt.Println(true)
	fmt.Printf("Type: %T Value: %v\n", true, true)

	var num1 int64 = 15
	var num2 int = 15

	fmt.Println(num1 + int64(num2))

	var num3 uint8 = 1
	var num4 uint64 = 1

	fmt.Println(unsafe.Sizeof(num3))
	fmt.Println(unsafe.Sizeof(num4))
}
