package raft

import "sync/atomic"

func (rf *Raft) nextTerm() {
	atomic.AddInt64(&rf.term, 1)
}
