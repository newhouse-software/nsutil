package nsutil

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Hello() string {
	return "Hello, world !"
}

func ExtractKeyValue(str, keyName string) string {
	// find the quoted value by the key=
	// example path: \\.\root\cimv2:Win32_Account.Domain="ZJIC-8G5T0S2",Name="SYSTEM"
	// keyname ".Domain" -> value "ZJIC-8G5T0S2"
	needle := keyName + `="`
	i := strings.Index(str, needle)
	if i < 0 {
		return ""
	}
	i += len(needle)
	j := strings.Index(str[i:], `"`)
	if j < 0 {
		return ""
	}
	return str[i : i+j]
}

func TrimPrefixIgnoreCase(str, prefix string) string {
	if strings.HasPrefix(strings.ToLower(str), strings.ToLower(prefix)) {
		return str[len(prefix):]
	}
	return str
}

func BytesToGB(bytes uint64) uint16 {
	const bytesInGB = 1024 * 1024 * 1024
	gb := float64(bytes) / float64(bytesInGB)
	return uint16(math.Round(gb))
}

func CsvToIntSlice(s string) ([]int, error) {
	parts := strings.Split(s, ",")
	result := make([]int, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		n, err := strconv.Atoi(p)
		if err != nil {
			return result, fmt.Errorf("invalid integer %q: %w", p, err)
		}
		result = append(result, n)
	}
	return result, nil
}
