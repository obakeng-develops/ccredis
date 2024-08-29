package pkg

import "strconv"

func SerializeSimpleStrings(value string) string {
	return "+" + value + "\r\n"
}

func SerializeErrors(value string) string {
	return "-" + value + "\r\n"
}

func SerializeIntegers(value int) string {
	return ":" + strconv.Itoa(value) + "\r\n"
}

func SerializeBulkStrings(value string) string {
	if value == "" {
		return "$-1\r\n"
	}

	return "$" + strconv.Itoa(len(value)) + "\r\n" + value + "\r\n"
}

func SerializeArrays(values []string) string {
	return "nil"
}
