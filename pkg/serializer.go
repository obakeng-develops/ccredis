package pkg

import (
	"fmt"
	"strconv"
)

const (
	CRLF                 = "\r\n"
	SIMPLE_STRING_PREFIX = "+"
	ERROR_PREFIX         = "-"
	INT_PREFIX           = ":"
	BULK_STRING_PREFIX   = "$"
	ARRAY_PREFIX         = "*"
)

func SerializeSimpleStrings(value string) string {
	return SIMPLE_STRING_PREFIX + value + CRLF
}

func SerializeErrors(value string) string {
	return ERROR_PREFIX + value + CRLF
}

func SerializeIntegers(value int) string {
	return INT_PREFIX + strconv.Itoa(value) + CRLF
}

func SerializeBulkStrings(value string) string {
	if value == "" {
		emptyBulkString := BULK_STRING_PREFIX + "0" + CRLF + CRLF
		return emptyBulkString
	}

	return BULK_STRING_PREFIX + strconv.Itoa(len(value)) + CRLF + value + CRLF
}

func SerializeArrays(s interface{}) string {
	emptyArray := ARRAY_PREFIX + "0" + CRLF

	switch values := s.(type) {
	case []string:
		arrayLength := len(values)
		finalArray := ARRAY_PREFIX + strconv.Itoa(arrayLength) + CRLF

		if arrayLength == 0 {
			return emptyArray
		}

		for _, value := range values {
			str := fmt.Sprintf("%v", value)
			finalArray += BULK_STRING_PREFIX + strconv.Itoa(len(str)) + CRLF + str + CRLF
		}

		return finalArray
	case []int:
		arrayLength := len(values)
		finalArray := ARRAY_PREFIX + strconv.Itoa(arrayLength) + CRLF

		if arrayLength == 0 {
			return emptyArray
		}

		for _, value := range values {
			str := fmt.Sprintf("%v", value)
			finalArray += INT_PREFIX + str + CRLF
		}

		return finalArray
	default:
		fmt.Println("Unsupported type")
		return ""
	}
}
