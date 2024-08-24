package cmd

import (
	"context"
	"fmt"
	"os"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"

	"github.com/halalala222/GoBoilder/internal/build"
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
)

func NewExecutor() *Executor {
	return &Executor{
		errorChan: make(chan error, 10),
		hasErr:    false,
	}
}

func (e *Executor) Execute() {
	huhModel := model.NewModel()
	_, err := tea.NewProgram(huhModel).Run()

	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	builderList := build.GenerateAllBuilder(build.WithProjectName(huhModel.GetProjectName()))
	wg := sync.WaitGroup{}
	wg.Add(len(builderList))
	for _, builder := range builderList {
		go func() {
			defer wg.Done()
			if buildErr := builder.Build(); buildErr != nil {
				e.errorChan <- buildErr
				e.hasErr = true
			}
		}()
	}

	_ = spinner.New().Title("Building your project...").Action(func() {
		wg.Wait()
		close(e.errorChan)
	}).Run()

	for err = range e.errorChan {
		fmt.Printf(errorHeaderStyle.Render("Oh no! Error building your project : "))
		fmt.Printf(errorContentStyle.Render(err.Error()))
	}

	if !e.hasErr {
		fmt.Printf(successStyle.Render("Project built successfully!"))
	}
}
