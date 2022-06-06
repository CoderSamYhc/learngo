package day_2

import (
	"fmt"
)


type Animal interface {
	SetName(v string)
	GetName() string
}

type Bird interface {
	Animal
	Call()
}

type Catamount interface {
	Animal
	Eating(food string)
}

type Duck struct {
	Name string
}

func (d *Duck) SetName(name string)  {
	d.Name = name
}

func (d *Duck) GetName() string  {
	return d.Name
}

func (d *Duck) Call() {
	fmt.Println(d.Name, "嘎嘎嘎！！")
}

type Tiger struct {
	Name string
}

func (t *Tiger) SetName(name string)  {
	t.Name = name
}

func (t *Tiger) GetName() string  {
	return t.Name
}

func (t *Tiger) Eating(food string)  {
	fmt.Println(t.Name , "is eating", food)
}