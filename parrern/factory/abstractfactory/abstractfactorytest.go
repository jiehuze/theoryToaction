package abstractfactory

import "fmt"

func Test() {
	//var absf AbstractFactoryInterface
	absf := AbstractFactory{}
	product1 := absf.CreateP1()
	product2 := absf.CreateP2()

	fmt.Printf("product1: %d\n", product1.sum(1, 2))
	fmt.Printf("product2:%d\n", product2.increase(4, 2))
}
