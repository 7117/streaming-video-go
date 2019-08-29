package taskrunner

import (
	"testing"
	"log"
	"time"
	"errors"
)

func TestRunner(t *testing.T) {
	// 给datachannel进行赋值  谁给datachannel进行赋值 就是分配者  所以d就是一个分配者
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i;
			log.Printf("Dispatcher sent: %v", i)
		}

		return nil
	}

	// 给datachannel取出消费数据  谁给datachannel进行取出数据  就是执行者   所以e就是一个执行者
	e := func(dc dataChan) error {
		forloop:
			for {
				select {
				case a :=<- dc:
					log.Printf("Executor received: %v", a)
				default:
					break forloop
				}
			}

		return errors.New("Executor")
	}

	// 初始化  进行赋值属性
	// channel长度  是否清空垃圾  分配者  消费者
	runner := NewRunner(30, false, d, e)
	// 运行程序  开启了生产消费者模式
	go runner.StartAll()
	// 睡三秒
	time.Sleep(3 * time.Second)
}