package http

var supportedFrameworks = map[string]FrameTemplate{
	"chi":        &ChiFrame{},
	"gin":        &GinFrame{},
	"fiber":      &FiberFrame{},
	"fasthttp":   &FasthttpFrame{},
	"huma":       &HumaFrame{},
	"echo":       &EchoFrame{},
	"mux":        &MuxFrame{},
	"httpRouter": &HttpRouterFrame{},
}

// FrameTemplate is the interface that wraps the Build method.
type FrameTemplate interface {
	Build() []byte
}

func GetAllSupportedHTTPFrameworks() []string {
	frameworks := make([]string, 0, len(supportedFrameworks))

	for k := range supportedFrameworks {
		frameworks = append(frameworks, k)
	}

	return frameworks
}
