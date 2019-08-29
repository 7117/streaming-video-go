package taskrunner

import (
	"os"
	"errors"
	"log"
	"sync"
	"scheduler/dbops"
)

// 真的删除
func deleteVideo(vid string) error {
	err := os.Remove(VIDEO_PATH + vid)

	if err != nil && !os.IsNotExist(err) {
		log.Printf("Deleting video error: %v", err)
		return err
	}

	return nil
}

// 收集者：收集数据给datachannel
func VideoClearDispatcher(dc dataChan) error {
	// 每次读取三条
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All tasks finished")
	}

	for _, id := range res {
		//写入到datachannel里面
		dc <- id
	}

	return nil
}

// 生产者：
func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error

	forloop:
		for {
			select {
			case vid :=<- dc:
				// 希望对于video_id删除过程是并发的
				// 可能会存在重复读写

				// 因为消费者的时候是协程的
				// 所以我们采用把错误写入到map里面  
				// 我们并不知道哪个快哪个慢  
				go func(id interface{}) {
					if err := deleteVideo(id.(string)); err != nil {
						errMap.Store(id, err)
						return
					}
					if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
						errMap.Store(id, err)
						return 
					}
				}(vid)
			default:
				break forloop
			}
		}

	// 最后进行输出map里面的内容 
	errMap.Range(func(k, v interface{}) bool {
		// 只要有一个错误  就进行返回
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})
	// 返回一个错误即可
	return err
}
