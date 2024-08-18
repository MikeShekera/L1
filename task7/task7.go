package task7

import "sync"

// Для обеспечения безопасности при конкурентной записи в Map необходимо делать lock через Mutex. Можно использовать как обычный, так и RW
func concurrentMapWrite(m map[string]string, key string, val string, mu *sync.RWMutex) {
	mu.Lock()
	m[key] = val
	mu.Unlock()
}
