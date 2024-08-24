package http

var _ FrameTemplate = &FasthttpFrame{}

type FasthttpFrame struct{}

func (f *FasthttpFrame) Build() []byte {
	return []byte(`package http`)
}
