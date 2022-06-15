package day_3

import "sync"

type Total struct {
	sync.Mutex
	Value int
}

func (t *Total) Worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		t.Lock()
		t.Value += 1
		t.Unlock()
	}
}
