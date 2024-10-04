package i18n

// helper function to flatten a map
// works like this => map[string]any{"a": map[string]any{"b": "c"}} => map[string]string{"a.b": "c"}
func flatten(nestedMap map[string]any, separator string) map[string]string {
	flatMap := make(map[string]string)

	flattenMap(nestedMap, "", separator, flatMap)
	return flatMap
}

func flattenMap(nestedMap map[string]any, parentKey string, separator string, flatMap map[string]string) {
	for key, value := range nestedMap {
		var newKey string
		if parentKey == "" {
			newKey = key
		} else {
			newKey = parentKey + separator + key
		}

		if _, ok := value.(map[string]any); ok {
			flattenMap(value.(map[string]any), newKey, separator, flatMap)
		} else {
			flatMap[newKey] = value.(string)
		}
	}
}
