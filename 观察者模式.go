package main

import (
	"fmt"
)

// 读者(观察者)
type Reader interface {
	// 更新方法
	Update()
}

// 读者1
type Reader1 struct {
	subject *Subject
}

func (b *Reader1) Reader1Observer(subject *Subject) {
	b.subject = subject
	b.subject.Attach(b)
}

func (b *Reader1) Update() {
	fmt.Println("读者1: ", b.subject.GetState())
}

// 读者2
type Reader2 struct {
	subject *Subject
}

func (o *Reader2) Reader2Observer(subject *Subject) {
	o.subject = subject
	o.subject.Attach(o)
}

func (o *Reader2) Update() {
	fmt.Println("读者2:", o.subject.GetState())
}

// 读者3
type Reader3 struct {
	subject *Subject
}

func (h *Reader3) Reader3Observer(subject *Subject) {
	h.subject = subject
	h.subject.Attach(h)
}

func (h *Reader3) Update() {
	fmt.Println("读者3:", h.subject.GetState())
}

// 主题
type Subject struct {
	// 观察者列表
	readers []Reader
	state   int
}

func NewSubject() *Subject {
	return &Subject{
		state:   0,
		readers: make([]Reader, 0),
	}
}

// 获得状态
func (s *Subject) GetState() int {
	return s.state
}

// 设置状态
func (s *Subject) SetState(state int) {
	s.state = state
	s.NotifyAllObservers()
}

// 添加观察者
func (s *Subject) Attach(observer Reader) {
	s.readers = append(s.readers, observer)
}

// 通知所有观察者
func (s *Subject) NotifyAllObservers() {
	for _, observer := range s.readers {
		observer.Update()
	}
}

func main() {
	subject := NewSubject()
	reader1 := Reader1{subject: subject}
	reader1.Reader1Observer(subject)
	a := subject.readers
	for _, v := range a {
		v.Update()
	}
	reader2 := Reader2{subject: subject}
	reader3 := Reader3{subject: subject}
	subject.Attach(&reader2)
	subject.Attach(&reader3)

	subject.SetState(1)
	subject.SetState(2)
}
