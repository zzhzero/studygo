package drive

type ConcurentDrive struct {
	Scheduler Scheduler
	WorkCount int
	ItemChan  chan interface{}
}
type Scheduler interface {
	//调度器，获取Request，配置WorkChan
	Submit(Request)
	ConfigureWorkChan(chan Request)
}

func (e *ConcurentDrive) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseRequest)

	e.Scheduler.ConfigureWorkChan(in)

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for i := 0; i < e.WorkCount; i++ {
		creatework(in, out)
	}

	for {
		result := <-out
		for _, ite := range result.Item {
			go func() { e.ItemChan <- ite }()
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}
func creatework(
	in chan Request, out chan ParseRequest) {

	go func() {
		for {
			request := <-in
			resualt, err := Worker(request)
			if err != nil {
				continue
			}
			out <- resualt
		}
	}()

}
