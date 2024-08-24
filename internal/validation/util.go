package validation

import (
	"os"

	"github.com/halalala222/GoBoilder/internal/constants"
)

func CheckProjectName(projectName string) error {
	var (
		err error
	)

	if projectName == "" {
		return constants.ErrProjectNameEmpty
	}

	if _, err = os.Stat(projectName); os.IsExist(err) {
		return constants.ErrProjectNameExists
	}

	if os.IsNotExist(err) {
		return nil
	}

	return err
}

func CheckModulePath(modulePath string) error {
	if modulePath == "" {
		return constants.ErrModulePathEmpty
	}

	return nil
}
