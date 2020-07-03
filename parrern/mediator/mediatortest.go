/**
 * @Author: jiehu
 * @Description:
 * @Project: test
 * @File:  mediatortest
 * @CopyRight: fuxi
 * @Date: 2020/7/2 5:56 下午
 */
package mediator

func Test() {

	mediator := NewConcreteMediator()

	//实现类通过中介进行调用，减少依赖
	colleague := &ConcreteColleague{cm: mediator}
	colleague1 := &ConcreteColleague1{cm: mediator}

	mediator.registered(1, colleague)
	mediator.registered(2, colleague1)

	colleague.sendmsg(1, "你好")

	colleague1.sendmsg(2, "很好")
}
