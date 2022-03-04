package bandaid

import "strconv"

func GetIntf(value string, defaultFunc func() int) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultFunc()
	}

	return i
}

func GetInt(value string, defaultValue int) int {
	return GetIntf(value, func() int { return defaultValue })
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func TruncateString(str string, num int) string {
	stored := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		stored = str[0:num] + "..."
	}
	return stored
}
