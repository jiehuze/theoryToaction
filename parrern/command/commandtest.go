package command

func Test() {
	var commandStart *StartCommand
	mathod := &Mathod{}
	commandStart = NewStartCommand(mathod)

	var commandStop *StopCommand
	commandStop = NewStopCommand(mathod)

	box1 := NewBox(commandStart, commandStop)
	box1.PressButton1()
	box1.PressButton2()

	box2 := NewBox(commandStop, commandStart)
	box2.PressButton1()
	box2.PressButton2()

}
