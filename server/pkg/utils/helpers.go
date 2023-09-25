package utils

import "fmt"

func CalcFileSize(size int64) string {
	decimalPlaces := 2
	units := []string{"B", "KB", "MB", "GB", "TB"}
	
	index := 0
	for size >= 1024 && index < len(units) - 1 {
		size /= 1024;
		index++;
	}
	
	// return fmt.Sprintf("%d %s", size, units[index])
	return fmt.Sprintf(fmt.Sprintf("%%.%df %%s", decimalPlaces), size, units[index])
}