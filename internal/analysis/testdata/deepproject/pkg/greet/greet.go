package greet

import "fmt"

// Hello returns a greeting string.
func Hello(name string) string {
	return Format(fmt.Sprintf("Hello, %s!", name))
}

// Goodbye returns a farewell string.
func Goodbye(name string) string {
	return Format(fmt.Sprintf("Goodbye, %s!", name))
}

// Format uppercases a greeting.
func Format(s string) string {
	return s
}

// UnusedExport is an exported function nobody calls from outside.
func UnusedExport() string {
	return "nobody calls me"
}

// internalHelper is unexported and called only within the package.
func internalHelper() string {
	return Format("internal")
}
