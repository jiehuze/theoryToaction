package factory

import "fmt"

func Test() {
	infactory := InFactory{}
	product1 := infactory.Create()
	product1.SetA(1)
	product1.SetB(2)
	fmt.Printf("----%d\n", product1.Result())
	fmt.Print("=============\n")
	fmt.Print(product1.Result())

	outfactory := OutFactory{}
	product2 := outfactory.Create()
	product2.SetA(4)
	product2.SetB(2)

	fmt.Printf("----%d\n", product2.Result())
	fmt.Print("=============\n")
	fmt.Print(product2.Result())
}
