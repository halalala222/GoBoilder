package config

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template/config/http"
)

type FrameworkInfo struct {
	Template []byte
	FileName string
}

var supportedFrameworks = map[string]*FrameworkInfo{
	constants.HTTPFrameworkGin: {
		Template: http.GinConfigTemplate,
		FileName: constants.HTTPFrameGinFileName,
	},
	constants.HTTPFrameworkEcho: {
		Template: http.EchoConfigTemplate,
		FileName: constants.HTTPFrameEchoFileName,
	},
	constants.HTTPFrameworkFiber: {
		Template: http.FiberConfigTemplate,
		FileName: constants.HTTPFrameFiberFileName,
	},
	constants.HTTPFrameworkChi: {
		Template: http.ChiConfigTemplate,
		FileName: constants.HTTPFrameChiFileName,
	},
}

// FrameTemplate is the interface that wraps the Build method.
type FrameTemplate interface {
	Build() []byte
}

func GetAllSupportedHTTPFrameworks() []string {
	frameworks := make([]string, 0, len(supportedFrameworks))

	for supportedFramework := range supportedFrameworks {
		frameworks = append(frameworks, supportedFramework)
	}

	return frameworks
}

func GetFrameworkInfo(framework string) (*FrameworkInfo, error) {
	var (
		frameworkInfo *FrameworkInfo
		ok            bool
	)

	if frameworkInfo, ok = supportedFrameworks[framework]; !ok {
		return nil, constants.ErrUnsupportedHTTPFramework
	}

	return frameworkInfo, nil
}
