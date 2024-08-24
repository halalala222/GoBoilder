package http

var _ FrameTemplate = &ChiFrame{}

type ChiFrame struct{}

func (c *ChiFrame) Build() []byte {
	return []byte(`package http`)
}
