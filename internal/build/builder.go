package build

type Builder interface {
	Build() error
}

func GenerateAllBuilder(options ...Option) []Builder {
	opts := getOptions(options)

	return []Builder{
		NewProjectBuilder(opts.projectName),
	}
}
