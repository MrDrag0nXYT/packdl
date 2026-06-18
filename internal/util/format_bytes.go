package util

import "fmt"

const unit = 1024

func FormatBytes(bytes uint64) string {
	switch {
	case bytes < unit:
		return fmt.Sprintf("%v B", bytes)
	case bytes < unit*unit:
		return fmt.Sprintf("%.2f KB", float64(bytes)/(unit))
	case bytes < unit*unit*unit:
		return fmt.Sprintf("%.2f MB", float64(bytes)/(unit*unit))
	case bytes < unit*unit*unit*unit:
		return fmt.Sprintf("%.2f GB", float64(bytes)/(unit*unit*unit))
	case bytes < unit*unit*unit*unit*unit:
		return fmt.Sprintf("%.2f TB", float64(bytes)/(unit*unit*unit*unit))

	default:
		return fmt.Sprintf("%.2d", bytes)
	}
}
