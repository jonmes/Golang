package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Employee struct {
	Person
	EmployeeID int
	Department string
}

type engine struct {
	model      string
	engineType string
	gallon     uint8
	mpg        uint8
}

type gasEngine struct {
	mpg    uint8
	gallon uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type engineType interface {
	milesLeft() uint8
}

func main() {
	fmt.Println("Structs")
	var p1 Person = Person{"John", "Doe", 25}
	var e1 Employee = Employee{Person{"Abebe", "kebede", 44}, 1001, "IT"}
	type Person2 struct {
		firstname, address string
		age                int
	}
	fmt.Println("-------", p1.FirstName, e1.Age)
	tata := Person2{firstname: "abebe", address: "addis_ababa", age: 3}
	tutu := Person2{firstname: "abebe", age: 44, address: "addis_ababa"}
	fmt.Println(tata.firstname, tata.address, tata.age)
	fmt.Println(tutu.firstname, tutu.address, tutu.age)
	e := engine{model: "XYZ", engineType: "V8"}
	e.printEngine()

	var gas engineType = gasEngine{mpg: 20, gallon: 10}
	var electric engineType = electricEngine{mpkwh: 9, kwh: 10}
	canMakeit(gas, 100)
	canMakeit(electric, 100)
}

func (e engine) printEngine() {
	fmt.Println("from the fuction", e.model)
	fmt.Println("This is from the function ", e.engineType)
}

// Note that we can make two different methods with the same name but different receiver types

func (g gasEngine) milesLeft() uint8 {
	return g.gallon * g.mpg
}

func (g electricEngine) milesLeft() uint8 {
	return g.mpkwh * g.kwh
}

func canMakeit(g engineType, mile uint8) {
	if mile <= g.milesLeft() {
		fmt.Println("You can make it")
	} else {
		fmt.Println("You can't make it")
	}
}
