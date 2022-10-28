package str

import (
	"encoding/json"
)

// PrettyJSON return JSON formatted string
func PrettyJSON(data interface{}) string {
	b, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		return "{}"
	}

	return string(b)
}
