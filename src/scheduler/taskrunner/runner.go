package taskrunner

import (
	"fmt"
)

// runner里面跑一个常驻的任务 比如叫做startDispatcher 任务会长时间的等待channel
// channel分为两部分
// control channel:两者进行交换信息的----
// 处理者说我的任务已经处理完成，请分发者读取分发新的任务；
// 分发者说新的任务我已经读取分发了，请处理者进行处理。
// data channel：数据
// 分发者的作用就是给数据channel进行赋值
// 执行者的作用就是把数据channel进行读取
type Runner struct{
	// 控制channel
	Controller controlChan
	Error controlChan
	Data dataChan
	datasize int
	// 资源的是否回收
	longlived bool
	Dispatcher fn 
	Executor fn
}

func NewRunner(size int, longlived bool, d fn, e fn) *Runner {
	return &Runner {
		// 非堵塞的才可以
		// 如果不为1的话 ，只要有值就会堵塞
		// 设置为1之后   有一个值不会堵塞
		Controller: make(chan string,1),
		Error: make(chan string,1),
		Data: make(chan interface{}, size),
		datasize: size,
		longlived: longlived,
		// 函数
		Dispatcher: d,
		// 函数
		Executor: e,
	}
}

// 常驻任务
func (r *Runner) startDispatch() {
	defer func() {
		// 不是longlived
		if !r.longlived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
		// ()是实参
	}()
	
	// 常驻任务for循环
	// 只要r.Controller有信息，就会执行
	for {
		select {
		// 控制器channel里面有信息
		case c :=<- r.Controller:
			// d==d
			if c == READY_TO_DISPATCH {
				// 分发者说新的任务数据我已经读取分发了，请处理者进行处理。
				// Dispatcher就是d  d的作用给DataChannel进行赋值  就是一个分发者的过程
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					// 给执行器发送信息  说你执行吧
					r.Controller <- READY_TO_EXECUTE
				}
			}

			if c == READY_TO_EXECUTE {
				// 处理者说我的任务已经处理完成，请分发者读取分发新的任务；
				// 如果没有信息的任务，那就等待吧
				// Executor就是e   e的作用就是读取datachannel里面的数据  就是一个执行者的过程
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					// 执行完之后  给ctrolchannel说执行完了  再让分配者去执行吧
					r.Controller <- READY_TO_DISPATCH
				}
			}
		//错误退channel里面有信息
		case e :=<- r.Error:
			if e == CLOSE {
				return
			}
		//默认执行的 
		default:
			fmt.Println("default");
		}
	}
}

func (r *Runner) StartAll() {
	fmt.Println("堵塞");
	// 赋值  让分配者执行
	r.Controller <- READY_TO_DISPATCH
	fmt.Println("堵塞");
	r.startDispatch()
}