package lab4

type CspStack struct{}

func NewCspStack() *CspStack {
	return &CspStack{}
}

func (cs *CspStack) Len() int {
	return 0
}

func (cs *CspStack) Push(value interface{}) {}

func (cs *CspStack) Pop() (value interface{}) {
	return nil
}

func (cs *CspStack) run() {
}
