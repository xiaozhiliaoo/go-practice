package main

import (
	"fmt"
	"time"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func EmployeeByID(id int) *Employee {
	employ := Employee{ID: 1, Name: "lili", Address: "beijing", Position: "111", Salary: 10, ManagerID: 111}
	return &employ
}

func EmployeeByID2(id int) Employee {
	employ := Employee{ID: 1, Name: "lili", Address: "beijing", Position: "111", Salary: 1000, ManagerID: 111}
	return employ
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}

func AwardAnnualRaise2(e Employee) {
	e.Salary = e.Salary * 105 / 100
}

type Point struct {
	X, Y int
}

type address struct {
	hostname string
	port     int
}

func main() {

	hits := make(map[address]int)
	hits[address{"golang.org", 433}]++

	pp := Point{
		X: 1,
		Y: 2,
	}
	qp := Point{
		X: 2,
		Y: 1,
	}
	fmt.Println(pp.X == qp.X && pp.Y == qp.Y)
	fmt.Println(pp == qp)

	ppp := new(Point)
	*ppp = Point{1, 2}
	pppp := &Point{1, 2}
	fmt.Println("pppp is", *pppp)
	fmt.Println("pppp is", pppp)

	employ := Employee{ID: 1, Name: "lili", Address: "beijing", Position: "111", Salary: 1000, ManagerID: 111}
	fmt.Println(employ)
	fmt.Println(EmployeeByID(employ.ManagerID).Position)
	fmt.Println(EmployeeByID2(employ.ManagerID).Position)
	//会改变结构体的值
	AwardAnnualRaise(&employ)
	fmt.Println(employ)

	employ2 := Employee{ID: 1, Name: "lili", Address: "beijing", Position: "111", Salary: 1000, ManagerID: 111}
	//不会改变结构体的值
	AwardAnnualRaise2(employ2)
	fmt.Println(employ2)

	//var ages2 map[string]int
	//ages2["lili"] = 4

	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	fmt.Println(ages["alice"])
	fmt.Println(ages["bob"])
	age, ok := ages["bob"]
	if !ok {
		//bob不在ages里面
		fmt.Println("age1:", age)
	}

	if age, ok := ages["bob"]; !ok {
		//bob不在ages里面
		fmt.Println("age2:", age)
	}

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...) // append the slice x
	fmt.Println(x)

	var runes []rune
	for _, r := range "hello" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
	fmt.Printf("%q\n", []rune("Hello, 世界"))

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(symbol[0:2])
	fmt.Println(RMB, symbol[RMB])

	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q)
	fmt.Println(r[2])

	q2 := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q2) // "[3]int"

	a1 := [2]int{1, 2}
	b1 := [...]int{1, 2}
	c1 := [2]int{1, 3}
	fmt.Println(a1 == b1, a1 == c1, b1 == c1)
	//d1:=[3]int{1,2}
	//fmt.Println(a1==d1)
}
