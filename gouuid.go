package gouuid

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

// Check if pointer of uuid is empty then return the default 00000000-0000-0000-0000-000000000000
// Use Case: let say we have ImageID : UUID sent by user via Rest API empty (nil) but database is required soem value
// So we need to validate the nil value before sending it to the database insert query.
func DefaultIfEmpty(val *uuid.UUID) uuid.UUID {
	if val == nil {
		return uuid.UUID{}
	}
	return *val
}

// Try to parse string as uuid and return default if error
func ParseToDefault(val string) uuid.UUID {
	valUUID, err := uuid.Parse(val)
	if err != nil {
		return uuid.UUID{}
	}

	return valUUID
}

// Find the index of UUID in a give Slice
func IndexOf(slice []uuid.UUID, val uuid.UUID) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// Convert Postgres array `{c4ba81a1-9e57-4e60-b811-2860136ab803,e0d6cbbb-e5b8-4f84-9f11-65b7e14e5817}` to UUID Slide uuid.UUID[]
func PgStringArrayToUUIDSlide(array string) []uuid.UUID {
	var (
		// unquoted array values must not contain: (" , \ { } whitespace NULL)
		// and must be at least one char
		unquotedChar  = `[^",\\{}\s(NULL)]`
		unquotedValue = fmt.Sprintf("(%s)+", unquotedChar)

		// quoted array values are surrounded by double quotes, can be any
		// character except " or \, which must be backslash escaped:
		quotedChar  = `[^"\\]|\\"|\\\\`
		quotedValue = fmt.Sprintf("\"(%s)*\"", quotedChar)

		// an array value may be either quoted or unquoted:
		arrayValue = fmt.Sprintf("(?P<value>(%s|%s))", unquotedValue, quotedValue)

		// Array values are separated with a comma IF there is more than one value:
		arrayExp = regexp.MustCompile(fmt.Sprintf("((%s)(,)?)", arrayValue))
	)

	var valueIndex int
	results := make([]uuid.UUID, 0)
	matches := arrayExp.FindAllStringSubmatch(array, -1)
	for _, match := range matches {
		s := match[valueIndex]
		// the string _might_ be wrapped in quotes, so trim them:
		s = strings.Trim(s, "\"")
		s = strings.Replace(s, ",", "", -1)

		val, err := uuid.Parse(s)
		if err != nil {
			results = append(results, val)
		}
	}
	return results
}

// Convert custom types to uuid.UUID[] e.g. type UUIDArray []uuid.UUID so we can pass
// UUIDArray as slice and will return uuid.UUID[]
func NamedTypeSliceToUUIDSlice(values []uuid.UUID) []uuid.UUID {
	var result []uuid.UUID
	result = append(result, values...)
	// for _, v := range values {
	// }

	return result
}

// Convert slide of UUIDs to String Slice
func ToStringSlice(values []uuid.UUID) []string {
	var result []string
	for _, v := range values {
		result = append(result, v.String())
	}

	return result
}

// Extract Slice of []uuid.UUID from a given any slice of struct
func ExtractIdsFromStructSlice(items interface{}, key string) []uuid.UUID {
	s := reflect.ValueOf(items)
	if s.Kind() != reflect.Slice {
		return nil
		// panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}
	var ids []uuid.UUID

	for i := 0; i < s.Len(); i++ {
		user := s.Index(i).Interface()
		// fmt.Println(reflect.TypeOf(user))
		// fmt.Println(reflect.ValueOf(user).Kind())
		val := reflect.ValueOf(&user).Elem().Elem()
		kind := reflect.ValueOf(user).Kind()

		switch kind {
		case reflect.Ptr:
			val = reflect.Indirect(val)
		}

		if val.FieldByName(key).IsValid() {
			n := val.FieldByName(key).Interface()

			nKind := val.FieldByName(key).Kind()

			switch nKind {
			case reflect.Ptr:
				n1 := n.(*uuid.UUID)
				if *n1 != uuid.Nil {
					ids = append(ids, *n1)
				}
			default:
				n1 := n.(uuid.UUID)
				if n1 != uuid.Nil {
					ids = append(ids, n1)
				}
			}
		}
	}

	return ids
}
