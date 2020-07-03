package command

import "fmt"

/*
命令模式本质是把某个对象的方法调用封装到对象中，方便传递、存储、调用。

示例中把主板单中的启动(start)方法和重启(reboot)方法封装为命令对象，再传递到主机(box)对象中。于两个按钮进行绑定：

第一个机箱(box1)设置按钮1(button1) 为开机按钮2(button2)为重启。
第二个机箱(box1)设置按钮2(button2) 为开机按钮1(button1)为重启。
从而得到配置灵活性。

除了配置灵活外，使用命令模式还可以用作：

批处理
任务队列
undo, redo
等把具体命令封装到对象中使用的场合
*/

type Mathod struct {
}

/**
 * @Description //TODO
 * @Date 1:33 下午 2020/7/2
 * @Param
 * @return
 **/
func (m *Mathod) Start() {
	fmt.Print("print start \n")
}

func (m *Mathod) Stop() {
	fmt.Print("print stop \n")
}

type Command interface {
	Execute()
}

type StartCommand struct {
	mathod *Mathod
}

func NewStartCommand(m *Mathod) *StartCommand {
	return &StartCommand{
		mathod: m,
	}
}

func (s *StartCommand) Execute() {
	s.mathod.Start()
}

type StopCommand struct {
	mathod *Mathod
}

func NewStopCommand(m *Mathod) *StopCommand {
	return &StopCommand{
		mathod: m,
	}
}

func (s *StopCommand) Execute() {
	s.mathod.Stop()
}

type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}

func (b *Box) PressButton2() {
	b.button2.Execute()
}
