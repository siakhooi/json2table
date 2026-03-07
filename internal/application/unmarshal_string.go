/*
Package application run the application
*/
package application

import "encoding/json"

// UnmarshalAsString attempts to unmarshal JSON data as a string
func UnmarshalAsString(data []byte) (string, error) {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return "", err
	}
	return str, nil
}

// UnmarshalAsStringArray attempts to unmarshal JSON data as an array of strings
func UnmarshalAsStringArray(data []byte) ([]string, error) {
	var arr []string
	if err := json.Unmarshal(data, &arr); err != nil {
		return nil, err
	}
	return arr, nil
}
