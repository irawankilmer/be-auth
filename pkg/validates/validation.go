package validates

import (
	"be-blog/pkg/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type Validates struct {
	validator *validator.Validate
}

func NewValidates(v *validator.Validate) *Validates {
	return &Validates{validator: v}
}

type Sanitizable interface {
	Sanitize() map[string]any
}

func (vu *Validates) ValidateJSON(c *gin.Context, input any) bool {
	if err := c.ShouldBindJSON(input); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			errorsMap := buildValidationErrorMap(input, verr)

			var old any
			if s, ok := input.(Sanitizable); ok {
				old = s.Sanitize()
			} else {
				old = input
			}

			response.BadRequest(c, gin.H{
				"errors": errorsMap,
				"old":    old,
			}, "Validasi gagal")
			return false
		}
		response.BadRequest(c, gin.H{
			"error": "Format data tidak valid. Pastikan format JSON benar.",
		}, "Permintaan tidak valid")
		return false
	}
	return true
}

func getJSONFieldName(input any, field string) string {
	t := reflect.TypeOf(input)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() == reflect.Slice {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return field
	}

	if f, ok := t.FieldByName(field); ok {
		jsonTag := f.Tag.Get("json")
		return strings.Split(jsonTag, ",")[0]
	}

	return field
}

func buildValidationErrorMap(input any, verr validator.ValidationErrors) map[string]string {
	errorMessages := make(map[string]string)
	for _, e := range verr {
		field := getJSONFieldName(input, e.Field())
		msg := GetValidationMessage(field, e.Tag(), e.Param())
		errorMessages[field] = msg
	}

	return errorMessages
}
