package main

import "fmt"

//Context 状态保存 类
type Context struct {
	state State
}

//NewContext 实例化状态保存类
func NewContext() *Context {
	return &Context{
		state: nil,
	}
}

//SetState 设置状态保存类当前的状态
func (c *Context) SetState(s State) {
	c.state = s
}

//GetState 获取状态保存类当前的状态
func (c *Context) GetState() State {
	return c.state
}

//State 状态接口
type State interface {
	DoAction(context *Context)
	ToString() string
}

//StartState 开始状态类
type StartState struct{}

//NewStartState 实例化开始状态类
func NewStartState() *StartState {
	return &StartState{}
}

//DoAction 开始状态类的DoAction，实现State接口
func (start *StartState) DoAction(context *Context) {
	fmt.Println("现在是开始状态")
	context.state = start
}

//ToString 返回开始状态类名称
func (start *StartState) ToString() string {
	return "开始状态"
}

//StopState 停止状态类
type StopState struct{}

//NewStopState 实例化停止状态类
func NewStopState() *StopState {
	return &StopState{}
}

//DoAction 停止状态类方法，实现State接口
func (stop *StopState) DoAction(context *Context) {
	fmt.Println("现在是停止状态")
	context.state = stop
}

//ToString 返回停止状态类名称
func (stop *StopState) ToString() string {
	return "停止状态"
}

func main() {
	context := NewContext()

	startState := NewStartState()
	startState.DoAction(context)
	fmt.Println(context.GetState().ToString())
	fmt.Println("-------------")
	stopState := NewStopState()
	stopState.DoAction(context)
	fmt.Println(context.GetState().ToString())
}
