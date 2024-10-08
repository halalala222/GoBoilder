package model

import (
	"github.com/charmbracelet/huh"

	"github.com/halalala222/GoBoilder/internal/constants"
	templateConfig "github.com/halalala222/GoBoilder/internal/template/config"
	"github.com/halalala222/GoBoilder/internal/template/config/db"
	"github.com/halalala222/GoBoilder/internal/template/config/http"
	"github.com/halalala222/GoBoilder/internal/template/logger"
	"github.com/halalala222/GoBoilder/internal/validation"
)

func newProjectNameInputGroup() *huh.Group {
	return huh.NewGroup(
		huh.NewInput().
			Key(constants.ProjectNameKey).
			Value(&info.ProjectName).
			Title(constants.ProjectName).
			Placeholder(constants.ProjectNamePlaceholder).
			Validate(func(s string) error {
				return validation.CheckProjectName(s)
			}).
			Description(constants.ProjectNameDescription),
		huh.NewInput().
			Key(constants.ModulePathPrefixKey).
			Value(&info.ModulePathPrefix).
			Title(constants.ModulePathPrefix).
			Placeholder(constants.ModulePathPrefixPlaceholder).
			Validate(func(s string) error {
				return validation.CheckModulePath(s)
			}).
			Description(constants.ModulePathPrefixDescription),
	)
}

func newDBFormHuhGroup() *huh.Group {
	return huh.NewGroup(
		huh.NewSelect[string]().
			Value(&info.DB).
			Key(constants.DBKey).
			Options(huh.NewOptions(db.GetSupportedDatabases()...)...).
			Title(constants.ChoiceDBTitle).
			Description(constants.ChoiceDBDescription),
		huh.NewSelect[string]().
			Value(&info.DBLibrary).
			Key(constants.DBLibraryKey).
			OptionsFunc(func() []huh.Option[string] {
				return huh.NewOptions(db.GetDBLibraries(info.DB)...)
			}, &info.DB).
			Title(constants.ChoiceDBLibraryTitle).
			Description(constants.ChoiceDBLibraryDescription),
	)
}

func newConfigHuhGroup() *huh.Group {
	return huh.NewGroup(
		huh.NewSelect[string]().
			Key(constants.ConfigFileTypeKey).
			Value(&info.ConfigFileType).
			Options(huh.NewOptions(templateConfig.GetSupportedConfigFileTypes()...)...).
			Title(constants.ChoiceConfigFileTypeTitle).
			Description(constants.ChoiceConfigFileTypeDescription),
	)
}

func newFormHuhGroup() *huh.Group {
	return huh.NewGroup(
		huh.NewSelect[string]().
			Key(constants.LoggerKey).
			Value(&info.LoggerLibrary).
			Options(huh.NewOptions(logger.GetSupportedLoggerLibraries()...)...).
			Title(constants.ChoiceLoggerTitle).
			Description(constants.ChoiceLoggerDescription),

		huh.NewSelect[string]().
			Key(constants.HTTPFrameKey).
			Value(&info.HTTPFramework).
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
