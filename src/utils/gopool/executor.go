package gopool

type ExecutorService interface {
	Start()
	// 抵达task队列上线,通道关闭等情况抛出异常或等待
	Submit(task Task) error
	Shutdown()
}

type Task interface {
	Run() error
	Info() TaskInfo
}

type TaskInfo struct {
	Desc string
	ErrMsg string
}

