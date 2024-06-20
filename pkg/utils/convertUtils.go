package utils

import "fmt"

// convert interface to string values
func ConvertToString(data interface{}) []string {
	var result []string

	// Check if data is a slice of interface{}
	if slice, ok := data.([]interface{}); ok {
		// Iterate over each element in the slice
		for _, item := range slice {
			// Convert each element to a string
			if str, ok := item.(string); ok {
				// Append the string to the result slice
				result = append(result, str)
			}
		}
	}

	return result
}

// Function to convert the interface value to a slice of strings
func ConvertInterfaceToSliceOfStrings(interfaceValue interface{}) ([]string, error) {
	interfaceSlice, ok := interfaceValue.([]interface{})
	if !ok {
		return nil, fmt.Errorf("interface value is not a slice of interfaces")
	}

	var result []string

	for _, item := range interfaceSlice {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("item is not a map[string]interface{}")
		}
		text, ok := itemMap["text"].(string)
		if !ok {
			return nil, fmt.Errorf("text field is not a string")
		}

		result = append(result, text)
	}

	return result, nil
}
