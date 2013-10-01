package lab4

import "sync"

type SafeStack struct {
	top  *Element
	size int
	stackMu sync.Mutex
}

func (ss *SafeStack) Len() int {
	return ss.size
}

func (ss *SafeStack) Push(value interface{}) {
	ss.stackMu.Lock()
	defer ss.stackMu.Unlock()
	ss.top = &Element{value, ss.top}
	ss.size++
}

func (ss *SafeStack) Pop() (value interface{}) {
	if ss.size > 0 {
		ss.stackMu.Lock()
		defer ss.stackMu.Unlock()
		value, ss.top = ss.top.value, ss.top.next
		ss.size--
		return
	}

	return nil
}
