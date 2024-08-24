package constants

const (
	ZapLoggerLibrary  = "zap"
	SlogLoggerLibrary = "slog"
	DefaultChoiceItem = "default"

	NoneCurrentBuildInfo = "(none)"
	CurrentBuildTitle    = "Current build"

	ApplicationHeader = "GoBoiler Build"

	LoggerKey               = "logger"
	ChoiceLoggerTitle       = "Choose your logger"
	ChoiceLoggerDescription = "This will determine your logger"

	HTTPFrameKey               = "http_frame"
	ChoiceHTTPFrameTitle       = "Choose your http frame"
	ChoiceHTTPFrameDescription = "This will determine your http frame"

	DoneKey = "done"
	AllDone = "All done?"

	Yep  = "Yep"
	Quit = "Quit"

	QuitBody = "Application quit"

	ProjectNameKey         = "project_name"
	ProjectName            = "Enter project name"
	ProjectNamePlaceholder = "my_project"
	ProjectNameDescription = "This will be the name of your project"
)
