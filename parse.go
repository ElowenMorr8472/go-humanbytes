package humanbytes

import (
	"fmt"
	"strconv"
	"strings"
)

// Parse turns a value produced by Format ("1.5MB", "512B") back into bytes.
func Parse(s string) (float64, error) {
	s = strings.TrimSpace(s)
	i := 0
	for i < len(s) && (s[i] == '.' || s[i] == '-' || (s[i] >= '0' && s[i] <= '9')) {
		i++
	}
	if i == 0 {
		return 0, fmt.Errorf("no number in %q", s)
	}
	n, err := strconv.ParseFloat(s[:i], 64)
	if err != nil {
		return 0, err
	}
	unit := strings.ToUpper(strings.TrimSpace(s[i:]))
	if unit == "" {
		unit = "B"
	}
	for exp, u := range units {
		if u == unit {
			for j := 0; j < exp; j++ {
				n *= 1024
			}
			return n, nil
		}
	}
	return 0, fmt.Errorf("unknown unit %q", unit)
}
