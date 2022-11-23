package validate

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

func New(local string) (ut.Translator, *validator.Validate, error) {
	var (
		trans ut.Translator
		err   error
		has   bool
	)
	validate := validator.New()
	// 注册一个获取json tag的自定义方法
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	zhT := zh.New() //chinese
	enT := en.New() //english
	uni := ut.New(enT, zhT, enT)

	trans, has = uni.GetTranslator(local)
	if !has {
		return trans, validate, fmt.Errorf("uni.GetTranslator(%s) failed", local)
	}
	//register translate
	switch local {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(validate, trans)
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(validate, trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(validate, trans)
	}
	return trans, validate, err
}
