package lab4

type CspStack interface {
	Push(interface{})
	Pop() (interface{})
	Len() int
}

func (ss safeStack) Len() int {
	reply := make(chan interface{})
	ss <- commandData{action: length, result: reply}
	return (<-reply).(int)
}

func (ss safeStack) Push(element interface{}) {
	ss <- commandData{action: push, element: element}
}

func (ss safeStack) Pop() (value interface{}) {
	reply := make(chan interface{})
	ss <- commandData{action: pop, result: reply}
	result := <- reply
	return result
}

type safeStack chan commandData

type commandData struct {
	action commandAction
	element interface{}
	result chan<- interface{}
	data chan<- UnsafeStack
}

type commandAction int

const (
	end commandAction = iota
	pop
	push
	length
)

func NewCspStack() CspStack {
	ss := make(safeStack)
	go ss.run()
	return ss
}

func (ss safeStack) run() {
	stack := new(UnsafeStack)
	for command := range ss {
		switch command.action {
		case push:
			stack.Push(command.element)
		case pop:
			command.result <- stack.Pop()
		case length:
			command.result <- stack.Len()
		case end:
			close(ss)
			command.data <- *stack
		}
	}
}
