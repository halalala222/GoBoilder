package validation

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/halalala222/GoBoilder/internal/constants"
)

func CheckProjectName(projectName string) error {
	var (
		err error
	)

	if projectName == "" {
		return constants.ErrProjectNameEmpty
	}

	if err = validateProjectName(projectName); err != nil {
		return err
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

	return validateModulePathPrefix(modulePath)
}

// isAllowedCharacter checks if the character is an allowed ASCII character
// ASCII letters (a-z, A-Z), digits (0-9), and special characters (-, ., _, ~) are allowed
func isAllowedCharacter(char rune) bool {
	// Check if it's an ASCII letter (a-z, A-Z)
	if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
		return true
	}
	// Check if it's a digit (0-9)
	if char >= '0' && char <= '9' {
		return true
	}
	// Check if it's one of the special characters (-, ., _, ~)
	return char == '-' || char == '.' || char == '_' || char == '~'
}

// validateSpecifiedASCII validates the specified ASCII characters
func validateSpecifiedASCII(needToValidateString string) error {
	for _, char := range needToValidateString {
		if char >= 0x80 {
			return errors.Join(errors.New(fmt.Sprintf("char: %c", char)), constants.ErrInvalidASCIICharacters)
		}

		if !isAllowedCharacter(char) {
			return errors.Join(errors.New(fmt.Sprintf("char: %c", char)), constants.ErrInvalidASCIICharacters)
		}
	}

	return nil
}

/*
*	from https://go.p2hp.com/ref/mod#module-path
*	A module path must satisfy the following requirements:
*
*	1.The path must consist of one or more path elements separated by slashes (/, U+002F). It must not begin or end with a slash.
*	2.Each path element is a non-empty string made of up ASCII letters, ASCII digits, and limited ASCII punctuation (-, ., _, and ~).
*	3.A path element may not begin or end with a dot (., U+002E).
*	4.The element prefix up to the first dot must not be a reserved file name on Windows, regardless of case (CON, com1, NuL, and so on).
*	5.The element prefix up to the first dot must not end with a tilde followed by one or more digits (like EXAMPL~1.COM).
*
*   for project name, it must satisfy the following requirements:
*	1. The name must consist of one or more path elements separated by slashes (/). It must not begin or end with a slash.
*	2. Each path element is a non-empty string made of up ASCII letters, ASCII digits, and limited ASCII punctuation (-, ., _, and ~).
*	3. A path element may not begin or end with a dot (.).
 */
func validateProjectName(projectName string) error {
	// for requirement 2
	if strings.HasSuffix(projectName, ".") || strings.HasPrefix(projectName, ".") {
		return constants.ErrProjectNameStartOrEndWithDot
	}

	// for requirement 1 and 3
	return validateSpecifiedASCII(projectName)
}

// windowsReservedFileNames is a list of reserved file names on Windows
var windowsReservedFileNames = []string{
	"CON", "PRN", "AUX", "NUL",
	"COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9",
	"LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9",
}

func validateModulePathPrefixBeforeFirstDotWithWindowReservedFileNames(firstSplit string) error {
	// for requirement 4
	for _, reservedFileName := range windowsReservedFileNames {
		if strings.EqualFold(firstSplit, reservedFileName) {
			return errors.Join(errors.New(fmt.Sprintf("element: %s", firstSplit)), constants.ErrProjectModulePathContainsWindowsReservedFileName)
		}
	}

	return nil
}

func validateModulePathPrefixBeforeFirstDotWithTildeFollowedByDigits(firstSplit string) error {
	// for requirement 5
	lastTildeIndex := strings.LastIndex(firstSplit, "~")
	if lastTildeIndex == -1 {
		return nil
	}

	for i := len(firstSplit) - 1; i > lastTildeIndex; i-- {
		if firstSplit[i] < '0' || firstSplit[i] > '9' {
			return nil
		}
	}

	return errors.Join(errors.New(fmt.Sprintf("element: %s", firstSplit)), constants.ErrProjectModulePathEndWithTildeFollowedByDigits)
}

func validateModulePathPrefixBeforeFirstDot(modulePathPrefix string) error {
	var (
		modulePathPrefixSplit = strings.Split(modulePathPrefix, ".")
	)

	if len(modulePathPrefixSplit) == 0 {
		return nil
	}

	firstSplit := modulePathPrefixSplit[0]

	if err := validateModulePathPrefixBeforeFirstDotWithWindowReservedFileNames(firstSplit); err != nil {
		return err
	}

	return validateModulePathPrefixBeforeFirstDotWithTildeFollowedByDigits(firstSplit)
}

func validateModulePathPrefixSplit(modulePathPrefix string) error {
	var (
		modulePathPrefixElements = strings.Split(modulePathPrefix, "/")
	)

	for _, element := range modulePathPrefixElements {
		// for requirement 2
		if err := validateSpecifiedASCII(element); err != nil {
			return errors.Join(errors.New(fmt.Sprintf("element: %s", element)), err)
		}
		// for requirement 3
		if strings.HasPrefix(element, ".") || strings.HasSuffix(element, ".") {
			return errors.Join(errors.New(fmt.Sprintf("element: %s", element)), constants.ErrProjectNameStartOrEndWithDot)
		}
	}

	return nil
}

func validateModulePathPrefix(modulePathPrefix string) error {
	// for requirement 1
	if strings.HasSuffix(modulePathPrefix, "/") || strings.HasPrefix(modulePathPrefix, "/") {
		return constants.ErrProjectModulePathPrefixStartOrEndWithSlash
	}

	// for requirement 1
	if strings.Contains(modulePathPrefix, "//") {
		return constants.ErrProjectModulePathPrefixContainsDoubleSlash
	}

	if err := validateModulePathPrefixSplit(modulePathPrefix); err != nil {
		return err
	}

	return validateModulePathPrefixBeforeFirstDot(modulePathPrefix)
}
