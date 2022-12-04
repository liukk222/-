package main

import (
	"container/list"
	"fmt"
	"reflect"
	"strconv"
)

// Employee 职员类
type Employee struct {
	Name         string
	Dept         string
	Salary       int
	Subordinates *list.List
}

// NewEmployee 实例化职员类
func NewEmployee(name, dept string, salary int) *Employee {
	sub := list.New()
	return &Employee{
		Name:         name,
		Dept:         dept,
		Salary:       salary,
		Subordinates: sub,
	}
}

// Add 添加职员的下属
func (e *Employee) Add(emp Employee) {
	e.Subordinates.PushBack(emp)
}

// Remove 删除职员的下属
func (e *Employee) Remove(emp Employee) {
	for i := e.Subordinates.Front(); i != nil; i = i.Next() {
		if reflect.DeepEqual(i.Value, emp) {
			e.Subordinates.Remove(i)
		}
	}
}

// GetSubordinates 获取职员下属列表
func (e *Employee) GetSubordinates() *list.List {
	return e.Subordinates
}

// ToString 获取职员的string信息
func (e *Employee) ToString() string {
	return "[ Name: " + e.Name + ", dept: " + e.Dept + ", Salary: " + strconv.Itoa(e.Salary) + " ]"
}

func main() {

	ceo := NewEmployee("老李", "ceo", 9999)

	pm := NewEmployee("ceo下属张三", "技术", 1000)
	programmer1 := NewEmployee("张三下属李四", "技术", 8000)
	programmer2 := NewEmployee("张三下属王五", "技术", 8000)

	minister := NewEmployee("ceo下属赵六", "财务", 5000)

	finance1 := NewEmployee("赵六下属陈七", "财务", 3000)
	finance2 := NewEmployee("赵六下属牛八", "财务", 2900)

	ceo.Add(*pm)
	ceo.Add(*minister)

	pm.Add(*programmer1)
	pm.Add(*programmer2)

	minister.Add(*finance1)
	minister.Add(*finance2)

	//打印所有职员
	fmt.Println(ceo.ToString())

	fmt.Println("####################")

	for i := ceo.Subordinates.Front(); i != nil; i = i.Next() {
		em := i.Value.(Employee)
		fmt.Println(em.ToString())
		fmt.Println("************************")
		for j := i.Value.(Employee).Subordinates.Front(); j != nil; j = j.Next() {
			em := j.Value.(Employee)
			fmt.Println(em.ToString())
		}
		fmt.Println("----------------")
	}
}
