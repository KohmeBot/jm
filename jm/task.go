package jm

import (
	"github.com/wdvxdr1123/ZeroBot/message"
	"sync"
	"time"
)

type TaskDuration struct {
	dur    time.Duration
	rw     sync.RWMutex
	tMp    map[int64]time.Time
	lastId map[int64]message.MessageID
}

func NewTaskDuration(dur time.Duration) *TaskDuration {
	return &TaskDuration{dur: dur, tMp: map[int64]time.Time{}, lastId: make(map[int64]message.MessageID)}
}

func (t *TaskDuration) AddTask(id int64) (bool, message.MessageID) {
	now := time.Now()
	t.rw.Lock()
	defer t.rw.Unlock()
	last, ok := t.tMp[id]
	if !ok || now.Sub(last) > t.dur {
		t.tMp[id] = now
		return true, message.MessageID{}
	}
	return false, t.lastId[id]
}

func (t *TaskDuration) Done(id int64, mid message.MessageID) {
	t.rw.Lock()
	defer t.rw.Unlock()
	t.lastId[id] = mid
}

func (t *TaskDuration) Fail(id int64) {
	t.rw.Lock()
	defer t.rw.Unlock()
	delete(t.lastId, id)
	delete(t.tMp, id)
}
