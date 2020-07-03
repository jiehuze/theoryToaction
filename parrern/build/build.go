package build

import "fmt"

/*
https://design-patterns.readthedocs.io/zh_CN/latest/creational_patterns/builder.html

建造者模式包含如下角色：
Builder：抽象建造者
ConcreteBuilder：具体建造者
Director：指挥者
Product：产品角色

模式分析：
抽象建造者类中定义了产品的创建方法和返回方法;

建造者模式的结构中还引入了一个指挥者类Director，该类的作用主要有两个：
	一方面它隔离了客户与生产过程；
	另一方面它负责控制产品的生成过程。
指挥者针对抽象建造者编程，客户端只需要知道具体建造者的类型，即可通过指挥者类调用建造者的相关方法，返回一个完整的产品对象

在客户端代码中，无须关心产品对象的具体组装过程，只需确定具体建造者的类型即可，建造者模式将复杂对象的构建与对象的表现分离开来，这样使得同样的构建过程可以创建出不同的表现。
*/

type Builder interface {
	buildPart1()
	buildPart2()
	buildPart3()
}

type Director struct {
	builder Builder
}

func NewDirector(build Builder) *Director {
	return &Director{builder: build}
}

//为具体的product 产品角色
func (d *Director) Construct() {
	d.builder.buildPart1()
	d.builder.buildPart2()
	d.builder.buildPart3()
}

type ConcreteBuilder struct {
}

func (b *ConcreteBuilder) buildPart1() {
	fmt.Print("!!!!!!build1 buildpart1\n")
}

func (b *ConcreteBuilder) buildPart2() {
	fmt.Print("!!!!!!build1 buildpart2\n")
}

func (b *ConcreteBuilder) buildPart3() {
	fmt.Print("!!!!!!build1 buildpart3\n")
}
