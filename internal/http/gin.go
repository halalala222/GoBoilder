package http

var _ FrameTemplate = &GinFrame{}

type GinFrame struct{}

func (g *GinFrame) Build() []byte {
	return []byte(`package http`)
}
