package helper

import "encoding/json"

func MustMarshal(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		// ここでパニックしてもいい（あり得ない想定）
		panic("failed to marshal: " + err.Error())
	}
	return b
}
