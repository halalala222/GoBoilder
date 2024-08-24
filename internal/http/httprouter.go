package http

var _ FrameTemplate = &HttpRouterFrame{}

type HttpRouterFrame struct{}

func (h *HttpRouterFrame) Build() []byte {
	return []byte(`package http`)
}
