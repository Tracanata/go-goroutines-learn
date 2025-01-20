package gogorountineslearn

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var pool sync.Pool

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Fatra")
	pool.Put("Candra")
	pool.Put("Ardiwinata")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Selesai")
}
