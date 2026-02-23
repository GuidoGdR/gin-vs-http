package errorBody

import "github.com/go-playground/validator/v10"

type errorBody struct {
	Status  string              `json:"status"`
	Message string              `json:"message"`
	Fields  map[string][]string `json:"fields,omitempty"`
}

func Generic(status string, message string, fields map[string][]string) *errorBody {
	return &errorBody{}
}
func InternalServerError() *errorBody {
	return Generic("Internal server Error", "Try again later", nil)
}
func BadRequestFields(message string, fields map[string][]string) *errorBody {
	return Generic("Bad Request", message, fields)
}
func BadRequest(message string) *errorBody {
	return Generic("Bad Request", message, nil)
}
func BadRequestValidationErrors(fields validator.ValidationErrors) *errorBody {
	return BadRequestFields("Validation error", FormatValidationError(fields))
}
func BadRequestFormat() *errorBody {
	return BadRequest("Format error")
}
func MethodNotAllowed(message string) *errorBody {
	return Generic("Method not allowed", message, nil)
}
func Unauthorized() *errorBody {
	return Generic("Unauthorized", "Invalid credentials", nil)
}

func FormatValidationError(err validator.ValidationErrors) map[string][]string {
	eMap := make(map[string][]string)
	for _, f := range err {
		eMap[f.Field()] = append(eMap[f.Field()], f.Tag())
	}
	return eMap
}
