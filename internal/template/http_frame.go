package template

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template/config/http"
)

var supportedFrameworks = map[string]*FileInfo{
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

// GetHTTPFrameFileTemplateInfo returns the FrameworkInfo for the given framework.
func GetHTTPFrameFileTemplateInfo(framework string) (*FileInfo, error) {
	if frameworkInfo, ok := supportedFrameworks[framework]; ok {
		return frameworkInfo, nil
	}

	return nil, constants.ErrUnsupportedHTTPFramework
}

// GetAllSupportedHTTPFrameworks returns the supported HTTP frameworks.
func GetAllSupportedHTTPFrameworks() []string {
	var frameworks = make([]string, 0, len(supportedFrameworks))

	for framework := range supportedFrameworks {
		frameworks = append(frameworks, framework)
	}

	return frameworks
}
