package utils

func FindNestedField(data map[string]interface{}, field string) (interface{}, bool) {
	if value, found := data[field]; found {
		return value, true
	}

	for _, v := range data {
		switch v := v.(type) {
		case map[string]interface{}:
			if value, found := FindNestedField(v, field); found {
				return value, true
			}
		case []interface{}:
			for _, item := range v {
				if itemMap, ok := item.(map[string]interface{}); ok {
					if value, found := FindNestedField(itemMap, field); found {
						return value, true
					}
				}
			}
		}
	}

	return nil, false
}
