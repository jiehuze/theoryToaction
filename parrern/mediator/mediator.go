/**
 * @Author: jiehu
 * @Description:
 * @Project: test
 * @File:  mediator
 * @CopyRight: fuxi
 * @Date: 2020/7/2 5:20 下午
 */
package mediator

import "fmt"

/**
 * @Date 2020/7/2 5:26 下午 jiehu
 *
Mediator(抽象中介者)：它定义了一个接口，该接口用于与各同事对象之间进行通信。
ConcreteMediator(具体中介者)：它实现了接口，通过协调各个同事对象来实现协作行为，维持各个同事对象的引用
Colleague(抽象同事类)：它定义了各个同事类公有的方法，并声明了一些抽象方法来供子类实现，同时维持了一个对抽象中介者类的引用，其子类可以通过该引用来与中介者通信。
ConcreteColleague(具体同事类)：抽象同事类的子类，每一个同事对象需要和其他对象通信时，都需要先与中介者对象通信，通过中介者来间接完成与其他同事类的通信
 *
 **/
type Mediator interface {
	operation(nWho int, str string)
	registered(nWho int, aColleague *Colleague)
}

type Colleague interface {
	sendmsg(toWho int, str string)
	receivemsg(str string)
}

type ConcreteMediator struct {
	mediatorMap map[int]Colleague
}

func NewConcreteMediator() ConcreteMediator {
	return ConcreteMediator{
		mediatorMap: make(map[int]Colleague, 0),
	}
}

func (m *ConcreteMediator) operation(nWho int, str string) {
	m.mediatorMap[nWho].receivemsg(str)
}

func (m *ConcreteMediator) registered(nWho int, aColleague Colleague) {
	m.mediatorMap[nWho] = aColleague
}

type ConcreteColleague struct {
	cm ConcreteMediator
}

func (c ConcreteColleague) sendmsg(toWho int, str string) {
	c.cm.operation(toWho, str)
}

func (c ConcreteColleague) receivemsg(str string) {
	fmt.Printf("reveivemsg: %s\n", str)
}

type ConcreteColleague1 struct {
	cm ConcreteMediator
}

func (c ConcreteColleague1) sendmsg(toWho int, str string) {
	c.cm.operation(toWho, str)
}

func (c ConcreteColleague1) receivemsg(str string) {
	fmt.Printf("ConcreteColleague1 reveivemsg: %s\n", str)
}
