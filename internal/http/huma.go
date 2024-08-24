package http

var _ FrameTemplate = &HumaFrame{}

type HumaFrame struct{}

func (h *HumaFrame) Build() []byte {
	return []byte(`package http`)
}
