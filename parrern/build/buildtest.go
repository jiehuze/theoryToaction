package build

func Test() {
	var build = new(ConcreteBuilder)

	director := NewDirector(build)
	//具体的生产者
	director.Construct()
}
