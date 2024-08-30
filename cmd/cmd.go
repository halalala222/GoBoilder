package cmd

import (
	"context"
	"errors"
	"fmt"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"

	"github.com/halalala222/GoBoilder/internal/build"
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/model"
)

type Executor struct {
	errorChan chan error
	cancel    context.CancelFunc
	hasErr    bool
}

var (
	errorHeaderStyle  = lipgloss.NewStyle().Bold(true)
	errorContentStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#D86A64"))
	successStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#6CA76C"))
	quitStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#6CA76C"))
	logoStyle         = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#6CA76C"))
)

func errorPrint(err error) {
	fmt.Printf(errorHeaderStyle.Render("Oh no! Error creating directories : "))
	fmt.Println(errorContentStyle.Render(err.Error()))
}

func NewExecutor() *Executor {
	return &Executor{
		errorChan: make(chan error, 10),
		hasErr:    false,
	}
}

func (e *Executor) Execute() {
	fmt.Println(logoStyle.Render(constants.Logo))

	huhModel := model.NewModel()

	if _, err := tea.NewProgram(huhModel).Run(); err != nil {
		errorPrint(err)
		return
	}

	huhModelInfo := huhModel.GetInfo()

	if huhModelInfo.IsQuit {
		fmt.Println(quitStyle.Render("Quit!Bye bye!"))
		return
	}

	if err := build.AllDir(huhModelInfo.ProjectName); err != nil {
		errorPrint(errors.New("Build all dir error: " + err.Error()))
		return
	}

	if err := build.GoModInit(huhModelInfo.ProjectName, huhModelInfo.ModulePath); err != nil {
		errorPrint(errors.New("Go mod init error: " + err.Error()))
		return
	}

	builderList := build.GenerateAllBuilder(
		build.WithProjectName(huhModelInfo.ProjectName),
		build.WithLoggerLibrary(huhModelInfo.LoggerLibrary),
		build.WithModulePath(huhModelInfo.ModulePath),
		build.WithDB(huhModelInfo.DB),
		build.WithDBLibrary(huhModelInfo.DBLibrary),
		build.WithConfigFileType(huhModelInfo.ConfigFileType),
		build.WithHTTPFramework(huhModelInfo.HTTPFramework),
	)

	wg := sync.WaitGroup{}
	wg.Add(len(builderList))
	for _, builder := range builderList {
		go func() {
			defer wg.Done()
			if buildErr := builder.Build(); buildErr != nil {
				e.errorChan <- errors.New(builder.String() + " : " + buildErr.Error())
				e.hasErr = true
			}
		}()
	}

	_ = spinner.New().Title("Building your project...").Action(func() {
		wg.Wait()
		if err := build.GoModTidy(huhModelInfo.ProjectName); err != nil {
			e.errorChan <- errors.New("Go mod tidy error: " + err.Error())
			e.hasErr = true
		}
		close(e.errorChan)
	}).Run()

	for err := range e.errorChan {
		errorPrint(err)
	}

	if !e.hasErr {
		fmt.Println(successStyle.Render("Project built successfully!"))
	}
}
