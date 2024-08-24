package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"

	"github.com/halalala222/GoBoilder/internal/build"
	"github.com/halalala222/GoBoilder/internal/model"
)

type Executor struct {
	errorChan chan error
	cancel    context.CancelFunc
}

var (
	errorHeaderStyle  = lipgloss.NewStyle().Bold(true)
	errorContentStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#D86A64"))
)

func (e *Executor) prepareBurger(projectName string) func() {
	return func() {
		projectBuilder := build.NewProjectBuilder(projectName)

		if err := projectBuilder.Build(); err != nil {
			e.errorChan <- err
			close(e.errorChan)
			e.cancel()
			return
		}
	}
}

func NewExecutor() *Executor {
	return &Executor{
		errorChan: make(chan error, 1),
	}
}

func (e *Executor) Execute() {
	huhModel := model.NewModel()
	_, err := tea.NewProgram(huhModel).Run()

	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	go func() {
		e.prepareBurger(huhModel.GetProjectName())()
	}()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	e.cancel = cancel
	_ = spinner.New().Title("Building your project...").Context(ctx).Run()

	for err = range e.errorChan {
		fmt.Printf(errorHeaderStyle.Render("Oh no! Error building your project : "))
		//fmt.Printf(errorContentStyle.Render(err.Error()))
	}
}
