package taskrunner

import (

)

// runner里面跑一个常驻的任务 比如叫做startDispatcher 任务会长时间的等待channel
// channel分为两部分
// control channel:两者进行交换信息的----
// 处理者说我的任务已经处理完成，请分发者读取分发新的任务；
// 分发者说新的任务我已经读取分发了，请处理者进行处理。
// data channel：数据
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
		Controller: make(chan string, 1),
		Error: make(chan string, 1),
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
		case c :=<- r.Controller:
			// d==d
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_EXECUTE
				}
			}

			if c == READY_TO_EXECUTE {
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}
		case e :=<- r.Error:
			if e == CLOSE {
				return
			}
		default:

		}
	}
}

func (r *Runner) StartAll() {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}