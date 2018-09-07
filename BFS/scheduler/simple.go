package scheduler

import "studygo/BFS/drive"

type SimpleScheduler struct {
	//向work传送Request
	WorkChan chan drive.Request
}

func (s *SimpleScheduler) ConfigureWorkChan(c chan drive.Request) {
	s.WorkChan = c
}

func (s *SimpleScheduler) Submit(r drive.Request) {
	go func() { s.WorkChan <- r }()
}
