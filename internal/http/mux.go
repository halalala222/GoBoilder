package http

var _ FrameTemplate = &MuxFrame{}

type MuxFrame struct{}

func (m *MuxFrame) Build() []byte {
	return []byte(`package http`)
}
