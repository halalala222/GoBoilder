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
	isQuit = false
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
			isQuit = true
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

func (m Model) currentLoggerBuildShow(buildInfo string) string {
	var (
		logger = m.form.GetString(constants.LoggerKey)
	)

	if logger != "" {
		return fmt.Sprintf("%s\nLogger: %s\n", buildInfo, logger)
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

func (m Model) completedShow(modelStyle *Styles) string {
	var (
		logger           = m.form.GetString(constants.LoggerKey)
		httpFrame        = m.form.GetString(constants.HTTPFrameKey)
		projectName      = m.form.GetString(constants.ProjectNameKey)
		modulePathPrefix = m.form.GetString(constants.ModulePathPrefixKey)
		completedInfo    strings.Builder
		errors           = m.form.Errors()
	)

	if len(errors) > 0 {
		return modelStyle.Status.Margin(0, 1).Padding(1, 2).Width(20).BorderForeground(red).Render(constants.QuitBody) + "\n\n"
	}

	completedInfo.Write([]byte("You have chosen the following:\n\n"))

	if projectName != "" {
		completedInfo.Write([]byte(fmt.Sprintf("Project Name: %s\n", projectName)))
	}

	if modulePathPrefix != "" {
		completedInfo.Write([]byte(fmt.Sprintf("Module Path Prefix: %s\n", modulePathPrefix)))
	}

	if logger != "" {
		completedInfo.Write([]byte(fmt.Sprintf("Logger: %s\n", logger)))
	}

	if httpFrame != "" {
		completedInfo.Write([]byte(fmt.Sprintf("HTTP Frame: %s\n", httpFrame)))
	}

	return modelStyle.Status.Margin(0, 1).Padding(1, 2).Width(48).Render(completedInfo.String()) + "\n\n"
}

func (m Model) applicationShow(modelStyle *Styles, form string, currentBuild string) string {
	errors := m.form.Errors()
	header := m.appBoundaryView(constants.ApplicationHeader)
	if len(errors) > 0 {
		m.form.State = huh.StateCompleted
	}

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
	IsQuit        bool
	ProjectName   string
	ModulePath    string
	LoggerLibrary string
}

func (m Model) GetInfo() *Info {
	info := &Info{
		IsQuit:        isQuit,
		ProjectName:   m.form.GetString(constants.ProjectNameKey),
		LoggerLibrary: m.form.GetString(constants.LoggerKey),
	}

	info.ModulePath = fmt.Sprintf("%s/%s", m.form.GetString(constants.ModulePathPrefixKey), info.ProjectName)

	return info
}

func (m Model) GetForm() *huh.Form {
	return m.form
}
