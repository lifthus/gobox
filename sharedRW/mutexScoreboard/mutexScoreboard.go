package mutexScoreboard

import (
	"sync"
)

type MutexScoreboardManager struct {
	l          sync.RWMutex
	scoreboard map[string]int
}

func NewMutexScoreboardManager() *MutexScoreboardManager {
	return &MutexScoreboardManager{
		scoreboard: map[string]int{},
	}
}

func (msm *MutexScoreboardManager) Update(name string, val int) {
	msm.l.Lock()
	defer msm.l.Unlock()
	msm.scoreboard[name] = val
}

func (msm *MutexScoreboardManager) Read(name string) (int, bool) {
	msm.l.RLock()
	defer msm.l.RUnlock()
	val, ok := msm.scoreboard[name]
	return val, ok
}

func (msm *MutexScoreboardManager) UpdateDangerously(name string, val int) {
	msm.scoreboard[name] = val
}

func (msm *MutexScoreboardManager) ReadDangerously(name string) (int, bool) {
	val, ok := msm.scoreboard[name]
	return val, ok
}
