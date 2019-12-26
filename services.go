package uadmin

import (
	"time"
)

func init() {
	go func() {
		for !dbOK {
			time.Sleep(time.Second)
		}
		go abTestService()
	}()
}

func abTestService() {
	for {
		if abTestCount != 0 {
			syncABTests()
		}
		time.Sleep(time.Second * 10)
	}
}
