package repl

import (
	"strings"
)

func CleanInput(text string) []string {
	//checks for invalid input
	if text == "" {
		return nil
	}
	/*  removed for now
	workingString := strings.ToLower(text)
	//adaptable special character check that can be updated if anomalies are encountered
	specialCharacters := []string{
		",", ".", ":", "@", "!", "-", "_",
	}
	for _, sc := range specialCharacters {
		workingString = strings.ReplaceAll(workingString, sc, "")
	}
	*/
	cleanedInput := strings.Fields(strings.ToLower(text))

	return cleanedInput
}
