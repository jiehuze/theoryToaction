package abstractfactory

// 生产1接口
type Product1Interface interface {
	sum(a, b int) int
}

//生产2接口
type Product2Interface interface {
	increase(a, b int) int
}

type AbstractFactoryInterface interface {
	CreateP1() Product1Interface
	CreateP2() Product2Interface
}

type Product1 struct {
}

func (p *Product1) sum(a, b int) int {
	return a + b
}

type Product2 struct {
}

func (p *Product2) increase(a, b int) int {
	return a - b
}

type AbstractFactory struct {
}

func (a *AbstractFactory) CreateP1() Product1Interface {
	return &Product1{}
}

func (a *AbstractFactory) CreateP2() Product2Interface {
	return &Product2{}
}
