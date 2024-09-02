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
	return []string{"1"}
}
