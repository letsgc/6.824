package raft

import (
	"sync"
	"sync/atomic"
	"time"
)

func (rf *Raft) sendHB() {
	var wg sync.WaitGroup
	var followerCnt int32

	for idx, _ := range rf.peers {
		if idx == rf.me {
			continue
		}
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			var resp AppendEntriesResp
			var ok bool
			RunWithTimeOut(func() {
				ok = rf.sendAppendEntries(idx, &AppendEntriesReq{rf.term, rf.me}, &resp)
			}, 1000*time.Millisecond)
			if resp.Illegal {
				rf.asFollower()
			}
			if ok {
				atomic.AddInt32(&followerCnt, 1)
			}
		}(idx)
	}
	wg.Wait()
	if followerCnt <= int32(len(rf.peers)/2) {
		rf.asFollower()
	}
}
