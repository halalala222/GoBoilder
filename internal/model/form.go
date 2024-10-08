package model

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"

	"github.com/halalala222/GoBoilder/internal/constants"
)

const maxWidth = 85

var (
	red    = lipgloss.AdaptiveColor{Light: "#FE5F86", Dark: "#FE5F86"}
	indigo = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	green  = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
	info   = &Info{}
)

type Styles struct {
	Base,
	HeaderText,
	Status,
	StatusHeader,
	Highlight,
	ErrorHeaderText,
	Help lipgloss.Style
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.Base = lg.NewStyle().
		Padding(1, 4, 0, 1)
	s.HeaderText = lg.NewStyle().
		Foreground(indigo).
		Bold(true).
		Padding(0, 1, 0, 2)
	s.Status = lg.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(indigo).
		PaddingLeft(1).
		MarginTop(1)
	s.StatusHeader = lg.NewStyle().
		Foreground(green).
		Bold(true)
	s.Highlight = lg.NewStyle().
		Foreground(lipgloss.Color("212"))
	s.ErrorHeaderText = s.HeaderText.
		Foreground(red)
	s.Help = lg.NewStyle().
		Foreground(lipgloss.Color("240"))
	return &s
}

type Model struct {
	lg     *lipgloss.Renderer
	styles *Styles
	form   *huh.Form
	width  int
}

func NewModel() Model {
	m := Model{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)

	m.form = huh.NewForm(
		newProjectNameInputGroup(),
		newDBFormHuhGroup(),
		newConfigHuhGroup(),
		newFormHuhGroup(),
	).
		WithWidth(50).
		WithShowHelp(false).
		WithShowErrors(false)
	return m
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			info.IsQuit = true
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd

	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		// Quit when the form is done.
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) currentProjectName(buildInfo string) string {
	var (
		projectName = m.form.GetString(constants.ProjectNameKey)
	)

	if projectName != "" {
		return fmt.Sprintf("Project Name: %s\n", projectName)
	}

	return buildInfo
}

func (m Model) currentModulePathPrefix(buildInfo string) string {
	var (
		modulePathPrefix = m.form.GetString(constants.ModulePathPrefixKey)
	)

	if modulePathPrefix != "" {
		return fmt.Sprintf("%sModule Path Prefix: %s\n", buildInfo, modulePathPrefix)
	}

	return buildInfo
}

func (m Model) currentDBBuildShow(buildInfo string) string {
	var (
		db = m.form.GetString(constants.DBKey)
	)

	if db != "" {
		return fmt.Sprintf("%sDB: %s\n", buildInfo, db)
	}

	return buildInfo
}

func (m Model) currentDBLibraryBuildShow(buildInfo string) string {
	var (
		dbLibrary = m.form.GetString(constants.DBLibraryKey)
	)

	if dbLibrary != "" {
		return fmt.Sprintf("%sDB Library: %s\n", buildInfo, dbLibrary)
	}

	return buildInfo
}

func (m Model) currentConfigFileTypeBuildShow(buildInfo string) string {
	var (
		configFileType = m.form.GetString(constants.ConfigFileTypeKey)
	)

	if configFileType != "" {
		return fmt.Sprintf("%sConfig File Type: %s\n", buildInfo, configFileType)
	}

	return buildInfo
}

func (m Model) currentLoggerBuildShow(buildInfo string) string {
	var (
		logger = m.form.GetString(constants.LoggerKey)
	)

	if logger != "" {
		return fmt.Sprintf("%sLogger: %s\n", buildInfo, logger)
	}

	return buildInfo
}

func (m Model) currentHTTPFrameBuildShow(buildInfo string) string {
	var (
		httpFrame = m.form.GetString(constants.HTTPFrameKey)
	)

	if httpFrame != "" {
		return fmt.Sprintf("%sHTTP Frame: %s\n", buildInfo, httpFrame)
	}

	return buildInfo
}

func (m Model) currentBuildShow(modelStyle *Styles, form string) string {
	var (
		buildInfo      = constants.NoneCurrentBuildInfo
		role           string
		jobDescription string
	)

	buildInfo = m.currentProjectName(buildInfo)
	buildInfo = m.currentModulePathPrefix(buildInfo)
	buildInfo = m.currentDBBuildShow(buildInfo)
	buildInfo = m.currentDBLibraryBuildShow(buildInfo)
	buildInfo = m.currentConfigFileTypeBuildShow(buildInfo)
	buildInfo = m.currentLoggerBuildShow(buildInfo)
	buildInfo = m.currentHTTPFrameBuildShow(buildInfo)

	const statusWidth = 28
	statusMarginLeft := m.width - statusWidth - lipgloss.Width(form) - modelStyle.Status.GetMarginRight()

	return modelStyle.Status.
		Height(lipgloss.Height(form)).
		Width(statusWidth).
		MarginLeft(statusMarginLeft).
		Render(modelStyle.StatusHeader.Render(constants.CurrentBuildTitle) + "\n" +
			buildInfo +
			role +
			jobDescription)
}

func (m Model) printCompletedProjectName(completedInfo *strings.Builder) {
	var (
		projectName = m.form.GetString(constants.ProjectNameKey)
	)

	if projectName != "" {
		completedInfo.Write([]byte(fmt.Sprintf("Project Name: %s\n", projectName)))
	}
}

func (m Model) printCompletedModulePath(completedInfo *strings.Builder) {
	var (
		modulePathPrefix = m.form.GetString(constants.ModulePathPrefixKey)
		projectName      = m.form.GetString(constants.ProjectNameKey)
	)

	if projectName != "" && modulePathPrefix != "" {
		completedInfo.Write([]byte(fmt.Sprintf("Module Path: %s/%s\n", modulePathPrefix, projectName)))
	}

	if projectName != "" && modulePathPrefix == "" {
		completedInfo.Write([]byte(fmt.Sprintf("Module Path: %s\n", projectName)))
	}
}

func (m Model) printCompletedLogger(completedInfo *strings.Builder) {
	var (
		logger = m.form.GetString(constants.LoggerKey)
	)

	if logger != "" {
		completedInfo.Write([]byte(fmt.Sprintf("Logger: %s\n", logger)))
	}
}

func (m Model) printCompletedHTTPFrame(completedInfo *strings.Builder) {
	var (
		httpFrame = m.form.GetString(constants.HTTPFrameKey)
	)

	if httpFrame != "" {
		completedInfo.Write([]byte(fmt.Sprintf("HTTP Frame: %s\n", httpFrame)))
	}
}

func (m Model) printCompletedDB(completedInfo *strings.Builder) {
	var (
		db = m.form.GetString(constants.DBKey)
	)

	if db != "" {
		completedInfo.Write([]byte(fmt.Sprintf("DB: %s\n", db)))
	}
}

func (m Model) printCompletedDBLibrary(completedInfo *strings.Builder) {
	var (
		dbLibrary = m.form.GetString(constants.DBLibraryKey)
	)

	if dbLibrary != "" {
		completedInfo.Write([]byte(fmt.Sprintf("DB Library: %s\n", dbLibrary)))
	}
}

func (m Model) printCompletedConfigFileType(completedInfo *strings.Builder) {
	var (
		configFileType = m.form.GetString(constants.ConfigFileTypeKey)
	)

	if configFileType != "" {
		completedInfo.Write([]byte(fmt.Sprintf("Config File Type: %s\n", configFileType)))
	}
}

func (m Model) completedShow(modelStyle *Styles) string {
	var (
		completedInfo = &strings.Builder{}
		errors        = m.form.Errors()
	)

	if len(errors) > 0 {
		return modelStyle.Status.Margin(0, 1).Padding(1, 2).Width(20).BorderForeground(red).Render(constants.QuitBody) + "\n\n"
	}

	completedInfo.Write([]byte("You have chosen the following:\n\n"))

	m.printCompletedProjectName(completedInfo)
	m.printCompletedModulePath(completedInfo)
	m.printCompletedDB(completedInfo)
	m.printCompletedDBLibrary(completedInfo)
	m.printCompletedConfigFileType(completedInfo)
	m.printCompletedLogger(completedInfo)
	m.printCompletedHTTPFrame(completedInfo)

	return modelStyle.Status.Margin(0, 1).Padding(1, 2).Width(48).Render(completedInfo.String()) + "\n\n"
}

func (m Model) applicationErrorShow(modelStyle *Styles, form string, currentBuild string) string {
	var (
		headerBuilder = strings.Builder{}
	)

	for _, err := range m.form.Errors() {
		headerBuilder.WriteString(m.appErrorBoundaryView(err.Error()))
		headerBuilder.WriteString("\n")
	}

	header := headerBuilder.String()

	body := lipgloss.JoinHorizontal(lipgloss.Top, form, currentBuild)

	footer := m.appErrorBoundaryView(m.form.Help().ShortHelpView(m.form.KeyBinds()))

	return modelStyle.Base.Render(header + "\n" + body + "\n\n" + footer)
}

func (m Model) applicationShow(modelStyle *Styles, form string, currentBuild string) string {
	if errors := m.form.Errors(); len(errors) > 0 {
		return m.applicationErrorShow(modelStyle, form, currentBuild)
	}

	header := m.appBoundaryView(constants.ApplicationHeader)

	body := lipgloss.JoinHorizontal(lipgloss.Top, form, currentBuild)

	footer := m.appBoundaryView(m.form.Help().ShortHelpView(m.form.KeyBinds()))

	return modelStyle.Base.Render(header + "\n" + body + "\n\n" + footer)
}

func (m Model) View() string {
	modelStyle := m.styles

	switch m.form.State {
	case huh.StateCompleted:
		return m.completedShow(modelStyle)
	default:
		// FormShow (left side)
		formShow := m.lg.NewStyle().Margin(1, 0).Render(strings.TrimSuffix(m.form.View(), "\n\n"))

		// currentBuildShow (right side)
		var currentBuildShow = m.currentBuildShow(modelStyle, formShow)

		return m.applicationShow(modelStyle, formShow, currentBuildShow)
	}
}

func (m Model) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.HeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(indigo),
	)
}

func (m Model) appErrorBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.ErrorHeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(red),
	)
}

type Info struct {
	IsQuit           bool
	ProjectName      string
	ModulePath       string
	ModulePathPrefix string
	LoggerLibrary    string
	HTTPFramework    string
	DB               string
	DBLibrary        string
	ConfigFileType   string
}

func (m Model) GetInfo() *Info {
	if len(info.ModulePathPrefix) != 0 {
		info.ModulePath = fmt.Sprintf("%s/%s", info.ModulePathPrefix, info.ProjectName)
	} else {
		info.ModulePath = info.ProjectName
	}

	return info
}

func (m Model) GetForm() *huh.Form {
	return m.form
}
