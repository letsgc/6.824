package raft

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

func (rf *Raft) startElection() {
	rf.Log(fmt.Sprintf("start a election"))

	rf.asCandidate()
	rf.nextTerm()

	var voteCnt int32 = 0

	if rf.vote(atomic.LoadInt64(&rf.term), rf.me) {
		voteCnt = 1
	}

	voteCnt += rf.askForVotes()

	if voteCnt > int32(len(rf.peers)/2) {
		rf.asLeader()
		return
	}

	s := time.Now().UnixMilli()
	timeLimit := int64(500 + rand.Uint64()%500)
	for {
		if rf.isFollower() {
			// detect a new leader
			return
		}
		if time.Now().UnixMilli()-s > timeLimit {
			rf.Log("wait timeout, start new election")
			rf.startElection()
			return
		}
	}
}
