package wyzvalidator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"gopkg.in/go-playground/validator.v8"
)

// Below is borrowed from some very kind stranger.
// sauce: https://github.com/gin-gonic/gin/issues/430#issuecomment-446113460
var (
	ErrorInternalError = errors.New("whoops something went wrong")
)

func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func LcFirst(str string) string {
	return strings.ToLower(str)
}

func Split(src string) string {
	// don't split invalid utf8
	if !utf8.ValidString(src) {
		return src
	}
	var entries []string
	var runes [][]rune
	lastClass := 0
	class := 0
	// split into fields based on class of unicode character
	for _, r := range src {
		switch true {
		case unicode.IsLower(r):
			class = 1
		case unicode.IsUpper(r):
			class = 2
		case unicode.IsDigit(r):
			class = 3
		default:
			class = 4
		}
		if class == lastClass {
			runes[len(runes)-1] = append(runes[len(runes)-1], r)
		} else {
			runes = append(runes, []rune{r})
		}
		lastClass = class
	}

	for i := 0; i < len(runes)-1; i++ {
		if unicode.IsUpper(runes[i][0]) && unicode.IsLower(runes[i+1][0]) {
			runes[i+1] = append([]rune{runes[i][len(runes[i])-1]}, runes[i+1]...)
			runes[i] = runes[i][:len(runes[i])-1]
		}
	}
	// construct []string from results
	for _, s := range runes {
		if len(s) > 0 {
			entries = append(entries, string(s))
		}
	}

	for index, word := range entries {
		if index == 0 {
			entries[index] = UcFirst(word)
		} else {
			entries[index] = LcFirst(word)
		}
	}
	justString := strings.Join(entries, " ")
	return justString
}

// Unit will take in the field being validated and return
// the appropriate string to use when describing the desired
// amount of whatever is being validated.
func Unit(e *validator.FieldError) string {
	switch e.Kind.String() {
	case "slice":
		i, _ := strconv.Atoi(e.Param)
		if i < 2 {
			return "entry"
		}
		return "entries"
	case "string":
		i, _ := strconv.Atoi(e.Param)
		if i < 2 {
			return "character"
		}
		return "characters"
	default:
		return "unknown"
	}
}

// ValidationErrorToText will take a field error and return the
// appropriate readable version of the error
func ValidationErrorToText(e *validator.FieldError) string {
	// NOTE: A case needs to be added to this for each tag
	//       you implement - this is probably the best and
	//       most obvious reason for consistency.
	//
	// EXPLANATIONS:
	// - e.Field => this is the field being validated
	// - word => this parses the field into readable form (ex: OldPassword -> Old password)
	// - e.Param => this will be the value the field field is "supposed" to be, or the limit, etc

	word := Split(e.Field)
	switch e.Tag {
	case "required":
		return fmt.Sprintf("%s is required", word)
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", word, e.Param)
	case "min":
		return fmt.Sprintf("%s must be longer than %s", word, e.Param)
	case "email":
		return fmt.Sprintf("Invalid email format")
	case "len":
		return fmt.Sprintf("%s must be %s characters long", word, e.Param)
	case "lte":
		return fmt.Sprintf("%s must contain no more than %s %s", word, e.Param, Unit(e))
	case "gte":
		return fmt.Sprintf("%s must contain at least %s %s", word, e.Param, Unit(e))
	case "alphanum":
		return fmt.Sprintf("%s must be alphanumeric", word)
	case "nefield":
		return fmt.Sprintf("%s must not be the same as %s", word, Split(e.Param))
	case "excludes":
		return fmt.Sprintf("%s must not be '%s'", word, e.Param)
	case "excludesrune":
		return fmt.Sprintf("%s must not contain '%s'", word, e.Param)
	case "eqfield":
		return fmt.Sprintf("%s must match %s", word, e.Param)
	}
	return fmt.Sprintf("%s is not valid", word)
}
