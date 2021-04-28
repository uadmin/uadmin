package uadmin

import (
	"net"
	"net/http"
	"sync"
	"time"
)

var rateLimitMap = map[string]int64{}
var rateLimitLock = sync.Mutex{}

// CheckRateLimit checks if the request has remaining quota or not. If it returns false,
// the IP in the request has exceeded their quota
func CheckRateLimit(r *http.Request) bool {
	rateLimitLock.Lock()
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	now := time.Now().Unix() * RateLimit
	if val, ok := rateLimitMap[ip]; ok {
		if (val + RateLimitBurst) < now {
			rateLimitMap[ip] = now - RateLimitBurst
		}
	} else {
		rateLimitMap[ip] = now - RateLimit
	}

	rateLimitMap[ip]++
	rateLimitLock.Unlock()
	return rateLimitMap[ip] <= now
}
