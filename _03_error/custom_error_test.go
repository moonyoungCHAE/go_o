package _03_error

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"testing"

	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	v *validator.Validate
)

type Student struct {
	Name string `validate:"required"`
	Age  int    `validate:"required"`
}

/*
1. translator 생성
2. 각 tag마다 custom error message 등록
3. error translate해야 custom error message 확인 할 수 있음
 */

func TestStudent(t *testing.T) {
	v = validator.New()

	t.Run("invalid required -> custom error 생성", func(t *testing.T) {
		// translation set up
		en := en.New()
		uni = ut.New(en, en)

		trans, _ := uni.GetTranslator("en") // Accept-Lang 한국어 가능

		en_translations.RegisterDefaultTranslations(v, trans)

		// register custom error
		RegisterCustomError(trans)

		// validate
		invalidStudent := Student{}
		err := v.Struct(invalidStudent)

		/*
		translate 안했을 때
			Key: 'Student.Name' Error:Field validation for 'Name' failed on the 'required' tag
			Key: 'Student.Age' Error:Field validation for 'Age' failed on the 'required' tag
		*/
		fmt.Println(err.Error())

		/*
		translate 했을 때
			[My Custom Error Msg] Name must have a value!
			[My Custom Error Msg] Age must have a value!
		 */
		if err != nil {

			errs := err.(validator.ValidationErrors)

			for _, e := range errs {
				fmt.Println(e.Translate(trans))
			}
		}
	})
}


func RegisterCustomError(trans ut.Translator) {
	// todo
	v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "[My Custom Error Msg] {0} must have a value!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
}
