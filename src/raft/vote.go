package raft

import (
	"fmt"
	"sync/atomic"
	"time"
)

func (rf *Raft) vote(term int64, i int) bool {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	if _, ok := rf.voteMap[term]; ok {
		return false
	}

	rf.voteMap[term] = i
	rf.Log(fmt.Sprintf("vote for %d", i))
	return true
}

func (rf *Raft) askForVotes() (voteCnt int32) {
	var cnt, done int32
	for idx, _ := range rf.peers {
		if idx == rf.me {
			continue
		}
		cnt++
		go func(idx int) {
			defer atomic.AddInt32(&done, 1)
			var resp RequestVoteReply
			RunWithTimeOut(func() {
				rf.sendRequestVote(idx, &RequestVoteArgs{Term: atomic.LoadInt64(&rf.term), Idx: rf.me}, &resp)
			}, 1000*time.Millisecond)
			if resp.Success {
				atomic.AddInt32(&voteCnt, 1)
			}
		}(idx)
	}
	for {
		if atomic.LoadInt32(&voteCnt) > int32(len(rf.peers)/2) {
			return atomic.LoadInt32(&voteCnt)
		}
		if atomic.LoadInt32(&done) == cnt {
			return voteCnt
		}
	}
}
