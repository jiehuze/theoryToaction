/**
 * @Author: jiehu
 * @Description:
 * @Project: test
 * @File:  strategytest
 * @CopyRight: fuxi
 * @Date: 2020/7/2 3:31 下午
 */
package strategy

func Test() {

	cash := &Cash{}
	bank := &Bank{}

	cashStrategy := NewPayment("cash", "123445", 1000, cash)
	cashStrategy.Pay(cashStrategy.Conext)

	bankStrategy:= NewPayment("bank", "8765f", 9000, bank)
	bankStrategy.Pay(bankStrategy.Conext)
}
