package uadmin

import (
	"net/http"
	"strings"
	"time"
)

var rateLimitMap = map[string]int64{}

// CheckRateLimit checks if the request has remaining quota or not. If it returns false,
// the IP in the request has exceeded their quota
func CheckRateLimit(r *http.Request) bool {
	ip := r.RemoteAddr
	index := strings.LastIndex(ip, ":")
	now := time.Now().Unix() * RateLimit
	ip = ip[0:index]
	if val, ok := rateLimitMap[ip]; ok {
		if (val + RateLimitBurst) < now {
			rateLimitMap[ip] = now - RateLimitBurst
		}
	} else {
		rateLimitMap[ip] = now - RateLimit
	}

	rateLimitMap[ip]++
	if rateLimitMap[ip] > now {
		return false
	}
	return true
}
