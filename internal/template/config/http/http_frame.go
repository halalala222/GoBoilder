package http

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
)

var supportedFrameworks = map[string]*template.FileInfo{
	constants.HTTPFrameworkGin: {
		Template: GinConfigTemplate,
		FileName: constants.HTTPFrameGinFileName,
	},
	constants.HTTPFrameworkEcho: {
		Template: EchoConfigTemplate,
		FileName: constants.HTTPFrameEchoFileName,
	},
	constants.HTTPFrameworkFiber: {
		Template: FiberConfigTemplate,
		FileName: constants.HTTPFrameFiberFileName,
	},
	constants.HTTPFrameworkChi: {
		Template: ChiConfigTemplate,
		FileName: constants.HTTPFrameChiFileName,
	},
}

// GetHTTPFrameFileTemplateInfo returns the FrameworkInfo for the given framework.
func GetHTTPFrameFileTemplateInfo(framework string) (*template.FileInfo, error) {
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
