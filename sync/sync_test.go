package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {

	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})

}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter returns a new Counter
func NewCounter() *Counter {
	return &Counter{}
}

// Inc the count
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current count
func (c *Counter) Value() int {
	return c.value
}
// func TestCounter(t *testing.T) {
// 	t.Run("Increment counter 3 times leave 3", func(t *testing.T) {
// 		counter := Counter{}
// 		counter.Inc()
// 		counter.Inc()
// 		counter.Inc()
// 		assertCountertimes(t, 3, &counter)
// 	})
// 	t.Run("Concurrent Test", func(t *testing.T) {
// 		counter := Counter{}
// 		wantedtimes := 1000
// 		var wg sync.WaitGroup
// 		wg.Add(wantedtimes)
// 		for i := 0; i < wantedtimes; i++ {
// 			go func(w *sync.WaitGroup) {
// 				//counter.mu.Lock()
// 				//defer counter.mu.Unlock()
// 				counter.Inc()
// 				w.Done()
// 			}(&wg)
// 		}
// 		wg.Wait()
// 		assertCountertimes(t, wantedtimes, &counter)
// 	})

// }
// func assertCountertimes(t *testing.T, want int, c *Counter) {
// 	t.Helper()
// 	if want != c.Value() {
// 		t.Errorf("want %d but got %d", want, c.Value())
// 	}
// }

// type Counter struct {
// 	Val int
// 	mu  sync.Mutex
// }

// func (c *Counter) Inc() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.Val++
// }
// func (c *Counter) Value() int {
// 	return c.Val
// }
