package http

var _ FrameTemplate = &EchoFrame{}

type EchoFrame struct{}

func (e *EchoFrame) Build() []byte {
	return []byte(`package http`)
}
