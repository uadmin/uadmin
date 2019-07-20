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
		syncABTests()
		time.Sleep(time.Second * 10)
	}
}
