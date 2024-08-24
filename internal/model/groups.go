package model

import (
	"github.com/charmbracelet/huh"

	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/http"
	"github.com/halalala222/GoBoilder/internal/logger"
	"github.com/halalala222/GoBoilder/internal/validation"
)

func newProjectNameInputGroup() *huh.Group {
	return huh.NewGroup(
		huh.NewInput().
			Key(constants.ProjectNameKey).
			Title(constants.ProjectName).
			Placeholder(constants.ProjectNamePlaceholder).
			Validate(func(s string) error {
				return validation.CheckProjectName(s)
			}).
			Description(constants.ProjectNameDescription),
		huh.NewInput().
			Key(constants.ModulePathKey).
			Title(constants.ModulePath).
			Placeholder(constants.ModulePathPlaceholder).
			Validate(func(s string) error {
				return validation.CheckModulePath(s)
			}).
			Description(constants.ModulePathDescription),
	)
}

func newFormHuhGroup() *huh.Group {
	return huh.NewGroup(
		huh.NewSelect[string]().
			Key(constants.LoggerKey).
			Options(huh.NewOptions(logger.GetAllSupportedLibraries()...)...).
			Title(constants.ChoiceLoggerTitle).
			Description(constants.ChoiceLoggerDescription),

		huh.NewSelect[string]().
			Key(constants.HTTPFrameKey).
			Options(huh.NewOptions(http.GetAllSupportedHTTPFrameworks()...)...).
			Title(constants.ChoiceHTTPFrameTitle).
			Description(constants.ChoiceHTTPFrameDescription),

		huh.NewConfirm().
			Key(constants.DoneKey).
			Title(constants.AllDone).
			Validate(func(v bool) error {
				if !v {
					return constants.ErrQuit
				}
				return nil
			}).
			Affirmative(constants.Yep).
			Negative(constants.Quit),
	)
}
