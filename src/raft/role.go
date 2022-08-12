package raft

import "sync/atomic"

func (rf *Raft) asCandidate() {
	atomic.StoreInt32(&rf.role, 1)
}

func (rf *Raft) asLeader() {
	atomic.StoreInt32(&rf.role, 2)
	go rf.continuousHB()
}

func (rf *Raft) asFollower() {
	atomic.StoreInt32(&rf.role, 0)
}

func (rf *Raft) isLeader() bool {
	return atomic.LoadInt32(&rf.role) == 2
}
func (rf *Raft) isFollower() bool {
	return atomic.LoadInt32(&rf.role) == 0
}
