package simplefactory

import "fmt"

/*
 * 简单工厂模式，通过factory类，添加所有产品实现类
 *
*/
//产品接口
type Product interface {
	Create()
}

// 产品1，的接口实现及定义
type ProductPen struct {
}

func (p ProductPen) Create() {
	fmt.Print("factory: product pen!\n")
}

// 产品2，的接口实现及定义
type ProductBook struct {
}

func (p ProductBook) Create() {
	fmt.Print("factory: product book !\n")
}

//工厂结构体
type Factory struct {
}

func (f Factory) Generate(name string) Product {
	switch name {
	case "pen":
		return ProductPen{}
	case "book":
		return ProductBook{}
	default:
		return nil
	}
}
