package raft

import (
	"sync/atomic"
	"time"
)

func (rf *Raft) receiveHb(idx int, term int64) bool {
	if term < rf.term {
		return true
	}

	if term > rf.term {
		atomic.StoreInt64(&rf.term, term)
	}

	rf.asFollower()

	rf.Log("rec hb")

	atomic.StoreInt64(&rf.lasthb, time.Now().UnixMilli())
	return false

}
