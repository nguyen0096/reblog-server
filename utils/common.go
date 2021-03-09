package utils

import "reflect"

// IsZeroValue reports whether value is the zero value for its type
func IsZeroValue(value interface{}) bool {
	return reflect.ValueOf(value).IsZero()
}

func ToMap(obj interface{}) map[string]interface{} {
	var m map[string]interface{}

	v := reflect.ValueOf(obj).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		// Print it out if needed
		// fmt.Printf("Field: %v - Value: %v - Tag: %s\n", t.Field(i).Name, v.Field(i), t.Field(i).Tag.Get("json"))
		m[t.Field(i).Name] = v.Field(i).Interface()
	}

	return m
}
