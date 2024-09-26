//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"sync"
	"time"
)

const FreeTierTimeLimit int64 = 10

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool

	TimeUsedmu sync.Mutex
	TimeUsed   int64 // in seconds
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	if u.IsPremium {
		process()
		return true
	}

	u.TimeUsedmu.Lock()
	defer u.TimeUsedmu.Unlock()

	if u.TimeUsed >= FreeTierTimeLimit {
		return false
	}

	start := time.Now()

	finishChan := make(chan bool)
	go func() {
		defer close(finishChan)
		process()
		finishChan <- true
	}()

	for {
		select {
		case <-finishChan:
			u.TimeUsed += int64(time.Since(start).Seconds())
			return u.TimeUsed <= FreeTierTimeLimit
		case <-time.After(time.Duration(FreeTierTimeLimit) * time.Second):
			u.TimeUsed = FreeTierTimeLimit
			return false
		}
	}
}

func main() {
	RunMockServer()
}
