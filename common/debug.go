package common

import "runtime/debug"

func printStackTrace() {
	debug.PrintStack() // 打印调用堆栈信息
}
