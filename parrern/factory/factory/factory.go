package factory

import "fmt"

/*
 * 工厂模式
 * 描述： 可以通过工厂模式，创建不同的工厂
 */

//生产基础接口，主要做计算
type Product interface {
	SetA(a int)
	SetB(b int)
	Result() int
}

//工厂基础接口，创建新的工厂类，返回各个工厂类的生产者，所有工厂类都实现product中的接口
type Factory interface {
	Create() Product
}

//生产基础类
type ProductBase struct {
	a int
	b int
}

func (p *ProductBase) SetA(a int) {
	p.a = a
}

func (p *ProductBase) SetB(b int) {
	p.b = b
}

//生产1 继承生产基础类
type Product1 struct {
	*ProductBase
}

func (p Product1) Result() int {
	fmt.Printf("++++a: %d, b: %d\n", p.a, p.b)
	return p.a + p.b
}

//第一工厂类
type InFactory struct {
}

func (in InFactory) Create() Product {
	return &Product1{
		ProductBase: &ProductBase{},
	}
}

//生产2 继承生产基础类
type Product2 struct {
	*ProductBase
}

func (p Product2) Result() int {
	fmt.Printf("++++a: %d, b: %d\n", p.a, p.b)
	return p.a - p.b
}

//第二工厂类
type OutFactory struct {
}

func (out OutFactory) Create() Product {
	return &Product2{
		ProductBase: &ProductBase{},
	}
}
