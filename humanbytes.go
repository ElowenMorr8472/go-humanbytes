// Package humanbytes formats byte counts, dependency-free.
package humanbytes

import "fmt"

var units = []string{"B", "KB", "MB", "GB", "TB", "PB"}

func Format(n float64) string {
	for i, u := range units {
		if n < 1024 || i == len(units)-1 {
			if u == "B" {
				return fmt.Sprintf("%d%s", int(n), u)
			}
			return fmt.Sprintf("%.1f%s", n, u)
		}
		n /= 1024
	}
	return ""
}
