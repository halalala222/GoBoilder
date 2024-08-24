package build

type buildOptions struct {
	projectName string
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

func getOptions(opts []Option) *buildOptions {
	builtOptions := getDefaultBuildOption()

	for _, opt := range opts {
		opt(builtOptions)
	}

	return builtOptions
}
