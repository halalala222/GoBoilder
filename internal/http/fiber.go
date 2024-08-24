package http

var _ FrameTemplate = &FiberFrame{}

type FiberFrame struct{}

func (f *FiberFrame) Build() []byte {
	return []byte(`package http`)
}
