/**
 * @Author: jiehu
 * @Description:
 * @Project: theoryToaction
 * @File:  statetest
 * @CopyRight: fuxi
 * @Date: 2020/7/2 2:26 下午
 */
package state

import "fmt"

func Test() {
	envclass := NewEnvClass()

	/**
	 * @Date 2020/7/2 2:45 下午 jiehu
	 * 通过不同health值来使用不同的行为，做到了状态值与行为的分离
	 *
	 **/
	envclass.SetValue(1)
	envclass.State.Comment()
	envclass.State.Post()
	envclass.State.View()
	fmt.Print("-------------\n")

	envclass.SetValue(14)
	envclass.State.Comment()
	envclass.State.Post()
	envclass.State.View()

	fmt.Print("-------------\n")
	envclass.SetValue(30)
	envclass.State.Comment()
	envclass.State.Post()
	envclass.State.View()
}
