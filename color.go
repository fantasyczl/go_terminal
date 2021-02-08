package go_terminal

import "fmt"

const (
	ColorRed    = "\033[0;31m"
	ColorGreen  = "\033[0;32m"
	ColorYellow = "\033[0;33m"
	ColorNone   = "\033[0m" // No Color
)

func RedString(msg string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, msg, ColorNone)
}

func GreenString(msg string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, msg, ColorNone)
}

func YellowString(msg string) string {
	return fmt.Sprintf("%s%s%s", ColorYellow, msg, ColorNone)
}
