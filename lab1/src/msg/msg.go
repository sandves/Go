package msg

import (
	"encoding/gob"
	"fmt"
)

func init() {
	gob.Register(StrMsg{})
	gob.Register(ErrMsg{})
}

type StrMsg struct {
	Sender  string
	Content string
}

type ErrMsg struct {
	Sender string
	Error  string
}

func (sm StrMsg) String() string {
	return fmt.Sprintf("Message from %s with content %q", sm.Sender, sm.Content)
}

func (em ErrMsg) String() string {
	return fmt.Sprintf("Error from %s with content %q", em.Sender, em.Error)
}
