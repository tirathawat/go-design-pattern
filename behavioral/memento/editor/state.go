package editor

import "encoding/json"

var (
	StateManager stateManager
)

type state string

type stateManager struct {
	histories []state
	current   int
}

func init() {
	StateManager.current = -1
	StateManager.histories = make([]state, 0)
}

func (s *stateManager) Push(memento state) {
	s.histories = append(s.histories, memento)
	s.current++
}

func (s *stateManager) Undo() state {
	if s.current > 0 {
		s.current--
	}
	return s.histories[s.current]
}

func (s *stateManager) Redo() state {
	if s.current < len(s.histories)-1 {
		s.current++
	}
	return s.histories[s.current]
}

func setState(s interface{}) state {
	serializedState, _ := json.Marshal(s)
	return state(string(serializedState))
}

func (m state) getState(s interface{}) {
	json.Unmarshal([]byte(m), s)
}
