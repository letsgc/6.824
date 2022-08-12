package raft

import (
	"fmt"
	"log"
	"time"
)

// Debugging
const Debug = false

func DPrintf(format string, a ...interface{}) (n int, err error) {
	if Debug {
		log.Printf(format, a...)
	}
	return
}

func (rf *Raft) Log(pattern string, i ...interface{}) {
	return
	content := fmt.Sprintf(pattern, i...)
	var s = "                     "
	var blank string
	for i := 0; i < rf.me; i++ {
		blank += s
	}
	fmt.Printf("%srf%d %s\n", blank, rf.me, content)
}

func RunWithTimeOut(f func(), duration time.Duration) {
	timeoutCh := time.After(duration)
	returnCh := make(chan struct{}, 1)

	go func() {
		defer func() { returnCh <- struct{}{} }()
		f()
	}()

	for {
		select {
		case _ = <-timeoutCh:
			return
		case _ = <-returnCh:
			return
		}
	}
}
