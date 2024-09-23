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

import "time"

const FreeTierTimeLimit int64 = 10

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	if !u.IsPremium {
		if u.TimeUsed >= FreeTierTimeLimit {
			return false
		}
	}

	processDuration := measureProcessDuration(process)

	if !u.IsPremium {
		u.TimeUsed += processDuration
		if u.TimeUsed > FreeTierTimeLimit {
			return false
		}
	}

	return true
}

func measureProcessDuration(process func()) int64 {
	start := time.Now()
	process()
	elapsed := time.Since(start)
	return int64(elapsed.Seconds())
}

func main() {
	RunMockServer()
}
