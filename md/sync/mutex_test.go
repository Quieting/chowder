package sync

import (
	"sync"
	"testing"
)

// 疑惑 sync.Mutex.sema 运作流程
// 猜想：控制多协程竞争锁时协程的阻塞和唤醒

// 简单模式
func TestMutexOneGr(t *testing.T) {
	var m sync.Mutex
	var incr int64
	m.Lock()
	incr++
	m.Unlock()
	t.Log(incr)
}
func TestMutex(t *testing.T) {
	var m sync.Mutex
	var incr int64
	quit := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			m.Lock()
			incr++
			quit <- struct{}{}
		}()
	}
	for i := 0; i < 10; i++ {
		<-quit
		m.Unlock()
	}
}
