/**
 * @Author: jiehu
 * @Description:
 * @Project: theoryToaction
 * @File:  state
 * @CopyRight: fuxi
 * @Date: 2020/7/2 1:18 下午
 */
package state

import "fmt"

/**
 * @Date 2020/7/2 2:06 下午 jiehu
 * 状态模式用于分离状态和行为。
 **/

//状态接口
type ActionState interface {
	View()
	Comment()
	Post()
}

//环境类
type EnvClass struct {
	State       ActionState
	HealthValue int
}

func NewEnvClass() *EnvClass {
	return &EnvClass{}
}

func (e *EnvClass) SetValue(v int) {
	e.HealthValue = v
	e.ChangeState(v)
}

//状态与行为分离
func (e *EnvClass) ChangeState(v int) {
	if v < 10 {
		e.State = &NormalState{}
	} else if v >= 10 && v < 20 {
		e.State = &Retricted{}
	} else {
		e.State = &CloseState{}
	}
}

// state 三种状态
type NormalState struct{}
type Retricted struct{}
type CloseState struct{}

// 封装state 三种状态的不同行为
func (c *CloseState) View() {
	fmt.Println("无法查看")
}

func (c *CloseState) Comment() {
	fmt.Println("不能评论")
}

func (c *CloseState) Post() {
	fmt.Println("不能发布")
}

func (r *Retricted) View() {
	fmt.Println("正常")
}

func (r *Retricted) Comment() {
	fmt.Println("正常")
}

func (r *Retricted) Post() {
	fmt.Println("不能发布")
}

func (r *NormalState) View() {
	fmt.Println("正常")
}

func (r *NormalState) Comment() {
	fmt.Println("正常")
}

func (r *NormalState) Post() {
	fmt.Println("正常")
}
