package utils

import (
	"errors"
	"reflect"
	custom_color "watchtower/custom_Color"
)

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

func StructToMap(obj any, filters ...string) map[string]any {
	result := make(map[string]any)
	v := reflect.ValueOf(obj)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		canInclude := true
		// Skip unexported fields
		if t.Field(i).IsExported() {
			for _, v := range filters {
				if v == t.Field(i).Name {
					canInclude = false
					break
				}
			}
			if canInclude {

				result[t.Field(i).Name] = field.Interface()
			}
		}
	}
	return result
}
