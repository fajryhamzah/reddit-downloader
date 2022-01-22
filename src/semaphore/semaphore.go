package semaphore

import "sync"

var waitGroup *sync.WaitGroup
var once sync.Once

func GetWaitGroup() *sync.WaitGroup {
	once.Do(func() {
		waitGroup = &sync.WaitGroup{}
	})

	return waitGroup
}
