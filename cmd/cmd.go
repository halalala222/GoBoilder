package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/halalala222/GoBoilder/internal/model"
)

func Execute() {
	_, err := tea.NewProgram(model.NewModel()).Run()

	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}
}
