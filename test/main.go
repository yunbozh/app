package main

import (
	"fmt"
	"net"
	"os"
	"reflect"
	"time"
)

type Student struct {
	name string
	age  uint32
}

type A interface {
	getName() string
}

type B interface {
	getAge() uint32
}

type People struct {
	Student
}

func (self People) getName() string {
	return self.name
}

func (self People) getAge() uint32 {
	return self.age
}

func main() {

	people := People{Student{name: "yunbozh", age: 1000}}

	var a A = people

	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)

	fmt.Println(t.Field(0).Name, t.Field(0).Type, v.Field(0).Interface())

	return

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Uage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("Ths address is:", addr.String())
	}

	os.Exit(0)

	//ch1 := make(chan int, 3)
	//
	//go test1(ch1)
	//
	//time.Sleep(time.Second)
	//
	//println("结束")
}

func test1(ch chan int) {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		println("test1", i)
	}

	ch <- 1000
	ch <- 2000
	ch <- 3000
	ch <- 4000
	println("继续")
}

func test2(ch1 <-chan int, ch2 chan<- int) {
	println("读取")
	value := <-ch1
	ch2 <- value
	for i := 1; i <= 5; i++ {
		time.Sleep(200 * time.Millisecond)
		println("test2", i)
	}
}
