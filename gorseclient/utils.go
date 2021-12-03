package gorseclient

import (
	"encoding/json"
)

func (a *API) jsonUnmarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		a.debugError(err)
		return err
	}
	return nil
}

// func (a *API) cleanCursor(cursor string) string {
// 	return strings.ReplaceAll(cursor, "\"", "")
// }
