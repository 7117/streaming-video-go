package taskrunner

const(
	// 预支controlChan的消息
	// controlChan发送READY_TO_DISPATCH被分配者DISPATCH收到后，分配者DISPATCH会进行下发数据到datachan里面
	// 下发完之后，给controlchan传递执行信息，到执行者运转了
	READY_TO_DISPATCH ="d"
	// 数据下发完之后，通过controlchan发消息值消费者EXECUTE，消费者EXECUTE就会去读datachan的消息，进行操作
	// 执行完之后，给controlchan传递分配信息，到分配者运转了
	READY_TO_EXECUTE = "e"
	// 所以controlChan是命令的传输者
	// 所以datachan是数据的池子
	CLOSE = "c"

	VIDEO_PATH = "./videos/"

)
// 预定义
type controlChan chan string
// 泛型
type dataChan chan interface{}
// fn函数类型的  
// 入参为datachan  出参为error
type fn func(dc dataChan) error

