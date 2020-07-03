/**
 * @Author: jiehu
 * @Description:
 * @Project: test
 * @File:  strategy
 * @CopyRight: fuxi
 * @Date: 2020/7/2 3:16 下午
 */
package strategy

import (
	"fmt"
)

/**
 * @Date 2020/7/2 3:31 下午 jiehu
 * 策略模式
 * 定义一系列算法，让这些算法在运行时可以互换，使得分离算法，符合开闭原则。
 * 主要解决：在有多种算法相似的情况下，使用 if...else 所带来的复杂和难以维护。
 * 何时使用：一个系统有许多许多类，而区分它们的只是他们直接的行为。
 *
 **/
type StragegyPayment interface {
	Pay(conext PaymentContext)
}

type PaymentContext struct {
	Name, CardId string
	Money        int
}

type Payment struct {
	Conext   PaymentContext
	Strategy StragegyPayment
}

func NewPayment(name, cardId string, money int, strategy StragegyPayment) *Payment {
	return &Payment{
		Conext: PaymentContext{
			Name:   name,
			CardId: cardId,
			Money:  money,
		},
		Strategy: strategy,
	}
}

func (p *Payment) Pay(conext PaymentContext) {
	p.Strategy.Pay(conext)
}

type Cash struct {
}

func (c *Cash) Pay(context PaymentContext) {
	fmt.Printf("Cash name: %s, id: %s, money: %d\n", context.Name, context.CardId, context.Money)
}

type Bank struct {
}

func (c *Bank) Pay(context PaymentContext) {
	fmt.Printf("Bank name: %s, id: %s, money: %d\n", context.Name, context.CardId, context.Money)
}
