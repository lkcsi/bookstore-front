package controller

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lkcsi/bookstore-front/custerror"
)

func setViewError(context *gin.Context, err error) {
	context.Writer.Header().Add("HX-Retarget", "#errors")
	context.Writer.Header().Add("HX-Reswap", "innetHTML")
	errors := ""

	switch e := err.(type) {
	case validator.ValidationErrors:
		for _, elem := range e {
			errors = fmt.Sprintf("%s %s", errors, getDangerMessage(getValidationMsg(elem)))
		}
	case custerror.CustError:
		errors = getDangerMessage(err.Error())
	default:
		errors = getDangerMessage("Internal sever error")
	}

	fmt.Println("errors", errors)
	tmpl, _ := template.New("t").Parse(errors)
	tmpl.Execute(context.Writer, nil)
}

func getDangerMessage(msg string) string {
	return fmt.Sprintf("<p class='alert alert-danger'>%s</p>", msg)
}

func getValidationMsg(fe validator.FieldError) string {
	switch fe.ActualTag() {
	case "required":
		return fe.Field() + " field is mandatory"
	case "gte":
		return fe.Field() + " must be greater than or equals " + fe.Param()
	}
	return "unkown error"
}
