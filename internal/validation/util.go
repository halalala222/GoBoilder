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
