package gopool

import (
	"errors"
	"log"
)

// 单协程工作池
// 无超时处理，已不再使用
type SingleExecutor struct {
	taskQueue chan Task
	shutdownSign chan bool
	Running bool
}


func NewSingleExecutor() *SingleExecutor{
	single := &SingleExecutor{
		taskQueue:make(chan Task),
		shutdownSign:make(chan bool),
		Running:false,
	}
	return single
}

func (ex *SingleExecutor) Start() {
	if ex.Running{
		return
	}
	ex.Running = true
	go ex.work()
}

func (ex *SingleExecutor) Submit(task Task) error{
	if ex.Running{
		ex.taskQueue <- task
		return nil
	}
	return errors.New("executor now is shutdown")
}

func (ex *SingleExecutor) Shutdown() {
	if ex.Running{
		ex.Running = false
		ex.shutdownSign <- true
		close(ex.shutdownSign)
	}
}

func (ex *SingleExecutor)work(){
	for{
		select {
		case <- ex.shutdownSign:
			return
		case task ,ok := <- ex.taskQueue:
			if ok{
				if err := task.Run();err != nil{
					log.Println("running task error")
				}
			}
		}
	}
}