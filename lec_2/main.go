package main

import "fmt"

func main() {
	//Простейший вывод Вывод аргумнта плюс перенос
	//fmt.Println("Hello world")
	//fmt.Println("second line")

	//fmt.Print("1")
	//fmt.Print("12")
	//fmt.Print("123")
	//форматированный вывод printf
	//fmt.Printf("\nHello my name is %s\n", "Marat")
	//fmt.Printf("My name is %s\nMy age is %d\n", "Marat", 100)

	//	var age int
	//fmt.Printf("My age is %d\n", age)
	//age = 32
	//fmt.Printf("My age is %d\n", age)

	//var height int = 180
	//fmt.Println("My height is ", height)

	//var uid = 12345
	//fmt.Print(uid)
	//fmt.Printf("%T:%d \n", uid, uid)

	// Множественное объявление переменных
	//var first, second, third int = 20, 30, 40
	//fmt.Print(first, second, third)

	// Блок декларации

	//var (
	//name       string = "bob"
	//person_age int    = 20
	//person_uid int
	//person_aid = 10
	//)

	//fmt.Print(name, person_age, person_uid, person_aid)

	//Странно
	//var a, b = 20, "Vova"
	//fmt.Print(a, b)
	// Хорошо
	//var a = 200 будет ошибка при повторной инициализации
	var age int
	var name string

	//fmt.Scan(&age)
	//fmt.Scan(&name)
	fmt.Scan(&age, &name)

	fmt.Printf("My name is %s my age is %d", name, age)
}
