package build

type buildOptions struct {
	projectName   string
	loggerLibrary string
	modulePath    string
}

type Option func(*buildOptions)

func getDefaultBuildOption() *buildOptions {
	return &buildOptions{
		projectName: "default",
	}
}

func WithProjectName(projectName string) Option {
	return func(o *buildOptions) {
		o.projectName = projectName
	}
}

func WithLoggerLibrary(loggerLibrary string) Option {
	return func(o *buildOptions) {
		o.loggerLibrary = loggerLibrary
	}
}

func WithModulePath(modulePath string) Option {
	return func(o *buildOptions) {
		o.modulePath = modulePath
	}
}

func getOptions(opts []Option) *buildOptions {
	builtOptions := getDefaultBuildOption()

	for _, opt := range opts {
		opt(builtOptions)
	}

	return builtOptions
}
