package gopool

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// 固定数量的协程池
// 增加了一个超时处理
type FixedExecutor struct {
	poolSize uint8
	chanBuf uint8
	taskQueue chan Task
	shutdownSign chan bool
	Running bool
	TaskTimeout time.Duration
}

func NewFixedExecutor(poolSize uint8 ,chanBuf uint8 ,taskTimeout time.Duration) *FixedExecutor{
	fixed := &FixedExecutor{
		chanBuf:chanBuf,
		poolSize:poolSize,
		taskQueue:make(chan Task ,chanBuf),
		shutdownSign:make(chan bool ,poolSize),
		TaskTimeout:taskTimeout,
		Running:false,
	}
	return fixed
}

func (executor *FixedExecutor) Start() {
	if executor.Running{
		return
	}
	executor.Running = true
	for i := 0;i<int(executor.poolSize);i++  {
		go executor.work()
	}
}

func (executor *FixedExecutor) Submit(task Task) error {
	if executor.Running{
		executor.taskQueue <- task
		return nil
	}
	return errors.New("executor now is shutdown")
}

func (executor *FixedExecutor) Shutdown() {
	if executor.Running{
		executor.Running = false
		executor.shutdownSign <- true
		close(executor.shutdownSign)
	}
}

func (executor *FixedExecutor)work(){
	for{
		select {
		case <- executor.shutdownSign:
			return
		case task ,ok := <- executor.taskQueue:
			timer := time.NewTimer(executor.TaskTimeout)
			ch := make(chan error)
			go func() {
				ch <- task.Run()
			}()
			if ok{
				select {
				case err := <- ch:
					if err != nil{
						log.Println(err)
					}
				case <- timer.C:
					//case <- time.After(executor.TaskTimeout):
					msg := fmt.Sprintf("goroutine time out after %dms" ,executor.TaskTimeout/time.Millisecond)
					log.Println(errors.New(msg))
				}
				timer.Stop()
			}
		}
	}
}

