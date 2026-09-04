package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"reflect"
	"regexp"
	"strings"
	custom_color "watchtower/custom_Color"
)

func wildcardToRegex(pattern string) string {
	escaped := regexp.QuoteMeta(pattern)
	escaped = strings.ReplaceAll(escaped, `\*`, `.*`)
	return "^" + escaped + "$"
}
func FailOnError(err error, msg string, fn func()) bool {

	if err != nil {
		custom_color.Error()(msg + "\n")
		custom_color.Error()("\n[-] %v\n", err)
		if fn != nil {

			fn()
		}
		return true
	}
	return false
}

func FailOnErrorPanic(err error, msg string, fn func()) {

	if err != nil {
		custom_color.Error()(msg)
		custom_color.Error()("\n[-] %v\n", err)
		if fn != nil {

			fn()
		}
		panic("")

	}
}

func GetStringMap(data map[string]any, key string) (string, error) {

	if di, ok := data[key]; ok {
		if r, ok := di.(string); ok {
			return r, nil
		}
	}
	return "", errors.New("Not Found")
}

func GetArrayStringMap(data map[string]any, key string) ([]string, error) {
	result := []string{}
	if di, ok := data[key]; ok {
		if arr, ok := di.([]interface{}); ok {
			for _, d := range arr {
				if str, ok := d.(string); ok {
					result = append(result, str)
				}
			}
		} else {
			return nil, errors.New("Not Found")
		}
	} else {
		return nil, errors.New("Not Found")
	}
	return result, nil
}
func HashObject(obj interface{}) (string, error) {
	// Convert object to JSON
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	// Create SHA256 hash
	hash := sha256.Sum256(jsonData)
	return hex.EncodeToString(hash[:]), nil
}
func StructToMapWithTags(obj any, filters ...string) map[string]any {
	result := make(map[string]any)
	v := reflect.ValueOf(obj)

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return result
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return result
	}

	t := v.Type()
	filterMap := make(map[string]bool, len(filters))
	for _, f := range filters {
		filterMap[f] = true
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		// Use json tag if available, fallback to field name
		tag := fieldType.Tag.Get("json")
		key := fieldType.Name
		if tag != "" && tag != "-" {
			// Handle options like "omitempty"
			if commaIdx := strings.Index(tag, ","); commaIdx != -1 {
				key = tag[:commaIdx]
			} else {
				key = tag
			}
		}

		if filterMap[key] || filterMap[fieldType.Name] {
			continue
		}

		result[key] = field.Interface()
	}

	return result
}
