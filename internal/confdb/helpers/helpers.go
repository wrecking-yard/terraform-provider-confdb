package helpers

import "github.com/itchyny/gojq"

func GoJqQuery(query string, object map[string]any) (map[string]any, error) {
	_query, err := gojq.Parse(query)
	if err != nil {
		return map[string]any{}, err
	}

	// https://github.com/itchyny/gojq?tab=readme-ov-file#usage-as-a-library
	iter := _query.Run(object)
	result := map[string]any{}
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			if err, ok := err.(*gojq.HaltError); ok && err.Value() == nil {
				break
			}
			return map[string]any{}, err
		}
		result, _ = v.(map[string]any)
		break
	}
	return result, nil
}
