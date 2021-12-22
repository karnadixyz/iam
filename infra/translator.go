package infra

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
	"gitlab.com/odeo/admin-iam/utils"
)

type transClient struct {
	Translator *ut.Translator
}

var TransSrv = &transClient{}

func (t *transClient) Init() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		enLocale := en.New()
		uniTrans := ut.New(enLocale, enLocale)
		trans, found := uniTrans.GetTranslator("en")
		if !found {
			return errors.New("translator not found")
		}
		if err := enTranslation.RegisterDefaultTranslations(v, trans); err != nil {
			return err
		}
		t.Translator = &trans
	}
	return nil
}

func (t *transClient) TranslateValidationError(err error) (errs []string) {
	if err == nil {
		return nil
	}

	var j *json.UnmarshalTypeError

	if errors.As(err, &validator.ValidationErrors{}) {
		validatorError := err.(validator.ValidationErrors)
		for _, e := range validatorError {
			translateErr := e.Translate(*t.Translator)
			errs = append(errs, translateErr)
		}
	} else if errors.As(err, &j) {
		fieldRgx, _ := regexp.Compile(`[ ]+`)
		invalidField := fieldRgx.Split(err.Error(), -1)
		errs = append(errs, fmt.Sprintf("%s must be %s type", invalidField[8], invalidField[len(invalidField)-1]))
	} else if pe := (&time.ParseError{}); errors.As(err, &pe) {
		errs = append(errs, fmt.Sprintf("failed to format '%s' with '%s' layout", pe.Value, pe.Layout))
	} else {
		log.Printf("ERROR: %v", err)
		errs = append(errs, "some error when binding request")
	}

	return errs
}

func (t *transClient) ErrorValidationResponse(c *gin.Context, validationErr error) {
	utils.BuildErrorResponse(c, utils.ErrorCodeValidation, t.TranslateValidationError(validationErr))
}
