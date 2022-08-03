package elegant

import (
	"encoding/json"
)

func dataToMap(data any) (map[string]interface{}, error) {
	m := map[string]interface{}{}

	j, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(j, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}
