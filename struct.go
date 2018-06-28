package main

import "fmt"

type person struct {
	name string
	age  int
}

func Structure() {
	fmt.Println("------Struct Sections------")
	st := person{name: "Salehin Rafi", age: 1}
	fmt.Println(st.name)
	fmt.Println(st.age)
}
