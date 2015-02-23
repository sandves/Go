type Message struct {
	Sender net.IP
	Content string
}

func (msg *Message) CheckForError() error {
	return error
}