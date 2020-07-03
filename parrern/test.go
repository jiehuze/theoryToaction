package main

import (
	"theoryToaction/parrern/build"
	"theoryToaction/parrern/command"
	"theoryToaction/parrern/factory/abstractfactory"
	"theoryToaction/parrern/factory/factory"
	"theoryToaction/parrern/factory/simplefactory"
	"theoryToaction/parrern/mediator"
	"theoryToaction/parrern/observer"
	"theoryToaction/parrern/state"
	"theoryToaction/parrern/strategy"
)

func main() {
	//简单工厂模式测试程序
	print("-------simplefactory--------\n")
	simplefactory.Test()

	//工厂模式
	print("-------factory--------\n")
	factory.Test()

	//抽象工厂模式
	print("-------abstractfactory--------\n")
	abstractfactory.Test()

	//build模式
	print("-------build--------\n")
	build.Test()

	//build模式
	print("-------observer--------\n")
	observer.Test()

	//command模式
	print("-------command--------\n")
	command.Test()

	//状态模式
	print("-------state--------\n")
	state.Test()

	//策略模式
	print("-------strategy--------\n")
	strategy.Test()

	//中介模式
	print("-------mediator--------\n")
	mediator.Test()

}
