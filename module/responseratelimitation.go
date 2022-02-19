package module

import (
	"sync"
	"time"
)

type ResponseRateLimitation struct {
	responseRateLimit int64
	history map[interface{}]int64
	lock sync.RWMutex
}

func(l *ResponseRateLimitation) Check(key interface{}) bool {
	l.lock.RLock()
	if historyTime, ok := l.history[key]; ok {
		l.lock.RUnlock()
		if time.Now().Unix() - historyTime > l.responseRateLimit {
			l.update(key)
			return true
		} else {
			return false
		}
	} else {
		l.lock.RUnlock()
		l.update(key)
		return true
	}
}

func (l *ResponseRateLimitation) update(key interface{}) {
	l.lock.Lock()
	l.history[key] = time.Now().Unix()
	l.lock.Unlock()
}

func NewResponseRateLimitation() *ResponseRateLimitation {
	return &ResponseRateLimitation{responseRateLimit: BotConfiguration.ResponseRateLimit, history: make(map[interface{}]int64)}
}
