package scheduler

import "myProject/Single-crawler-go/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkReady(chan engine.Request) {

}

func (s *SimpleScheduler) Submit(r engine.Request) {
	//send request down to work chan
	//request发送到work的通道

	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) Run(){
	s.workerChan = make(chan engine.Request)
}

