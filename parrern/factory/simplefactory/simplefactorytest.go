package simplefactory

func Test() {
	var factory = new(Factory)

	ppen := factory.Generate("pen")
	pbook := factory.Generate("book")

	ppen.Create()
	pbook.Create()
}
