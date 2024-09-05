package pkg

import (
	"strconv"
	"strings"
)

func DeserializeSimpleStrings(value string) string {
	value = strings.TrimPrefix(value, SIMPLE_STRING_PREFIX)
	value = strings.TrimSuffix(value, CRLF)

	return value
}

func DeserializeErrors(value string) string {
	value = strings.TrimPrefix(value, ERROR_PREFIX)
	value = strings.TrimSuffix(value, CRLF)

	return value
}

func DeserializeIntegers(value string) int {
	value = strings.TrimPrefix(value, INT_PREFIX)
	value = strings.TrimSuffix(value, CRLF)

	convertedValue, _ := strconv.Atoi(value)

	return convertedValue
}

func DeserializeBulkStrings(value string) string {
	value = strings.TrimPrefix(value, BULK_STRING_PREFIX)
	value = strings.TrimSuffix(value, CRLF)

	return value[3:] // removes the length in the final string (the number after $)
}

func DeserializeArrays(value string) interface{} {
	lines := strings.Split(value, CRLF)
	if len(lines) < 2 {
		return nil
	}

	if !strings.HasPrefix(lines[0], ARRAY_PREFIX) {
		return nil
	}

	arrayLength, err := strconv.Atoi(strings.TrimPrefix(lines[0], ARRAY_PREFIX))
	if err != nil {
		return nil
	}

	var result []interface{}
	index := 1

	for i := 0; i < arrayLength; i++ {
		if index >= len(lines) {
			return nil
		}

		switch {
		case strings.HasPrefix(lines[index], INT_PREFIX):
			intVal, err := strconv.Atoi(strings.TrimPrefix(lines[index], INT_PREFIX))
			if err != nil {
				return nil
			}
			result = append(result, intVal)
			index++

		case strings.HasPrefix(lines[index], BULK_STRING_PREFIX):
			length, err := strconv.Atoi(strings.TrimPrefix(lines[index], BULK_STRING_PREFIX))
			if err != nil || index+1 >= len(lines) {
				return nil
			}
			strVal := lines[index+1]
			if len(strVal) != length {
				return nil
			}
			result = append(result, strVal)
			index += 2

		default:
			return nil
		}
	}

	if allInts(result) {
		var intResult []int
		for _, v := range result {
			intResult = append(intResult, v.(int))
		}
		return intResult
	} else if allStrings(result) {
		var stringResult []string
		for _, v := range result {
			stringResult = append(stringResult, v.(string))
		}
		return stringResult
	}

	return result
}

func allInts(arr []interface{}) bool {
	for _, v := range arr {
		if _, ok := v.(int); !ok {
			return false
		}
	}
	return true
}

func allStrings(arr []interface{}) bool {
	for _, v := range arr {
		if _, ok := v.(string); !ok {
			return false
		}
	}
	return true
}
