package helper

import (
	"fmt"
	"strings"
)

func KeysString(m map[string]interface{}) string {
	arr := make([]string, 0, len(m))
	for k, v := range m {
		switch v.(type) {
		case string:
			arr = append(arr, fmt.Sprintf(`"%v":"%v"`, k, v))
		case int:
			arr = append(arr, fmt.Sprintf(`"%v":%v`, k, v))
		case float64:
			arr = append(arr, fmt.Sprintf(`"%v":%v`, k, v))
		case bool:
			arr = append(arr, fmt.Sprintf(`"%v":%v`, k, v))
		case nil:
			arr = append(arr, fmt.Sprintf(`"%v":null`, k))
		default:
			arr = append(arr, fmt.Sprintf(`"%v":"%v"`, k, v))
		}
	}
	return "{" + strings.Join(arr, ", ") + "}"
}
