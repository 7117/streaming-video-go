package taskrunner

import (
	"testing"
	"log"
	"time"
	"errors"
)

func TestRunner(t *testing.T) {
	// 给datachan进行赋值
	// 然后给分配者分配数据
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i;
			log.Printf("Dispatcher sent: %v", i)
		}

		return nil
	}

	// 给消费者分配数据
	e := func(dc dataChan) error {
		forloop:
			for {
				select {
				case d :=<- dc:
					log.Printf("Executor received: %v", d)
				default:
					break forloop
				}
			}

		return errors.New("Executor")
	}

	// 初始化
	runner := NewRunner(30, false, d, e)
	// 运行
	go runner.StartAll()
	// 睡三秒
	time.Sleep(3 * time.Second)
}