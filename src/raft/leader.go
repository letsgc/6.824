package raft

import (
	"sync"
	"time"
)

func (rf *Raft) sendHB() {
	var wg sync.WaitGroup

	for idx, _ := range rf.peers {
		if idx == rf.me {
			continue
		}
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			var resp AppendEntriesResp
			RunWithTimeOut(func() {
				rf.sendAppendEntries(idx, &AppendEntriesReq{rf.term, rf.me}, &resp)
			}, 1000*time.Millisecond)
			if resp.Illegal {
				rf.asFollower()
			}
		}(idx)
	}

	wg.Wait()
}
