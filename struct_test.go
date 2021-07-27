package _04_struct

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

var v *validator.Validate

/*
Ex.
ContactType == email -> ContactValue가 email 형식
ContactType == phone -> ContactValue가 전화번호 형식
*/

/*
1. RegisterStructValidation로 특정 struct에 validate 등록
2. 의도한 타입으로 형변환후 validate 로직 작성
3. invalid할 경우, sl.ReportError
*/

type Person struct {
	ContactType  string `validate:"oneof=email phone"`
	ContactValue string `validate:"required"`
}

func TestStruct(t *testing.T) {
	v = validator.New()
	v.RegisterStructValidation(StructLevelValidation, Person{})

	t.Run("[invalid] type: email, value: phone", func(t *testing.T) {
		invalid := Person{ContactType: "email", ContactValue: "0210102444"}
		err := v.Struct(invalid)
		assert.NotNil(t, err)
		// Key: 'Person.content_value' Error:Field validation for 'content_value' failed on the 'email' tag
	})

	t.Run("[invalid] type: phone, value: email", func(t *testing.T) {
		invalid := Person{ContactType: "phone", ContactValue: "hi@naver.com"}
		err := v.Struct(invalid)
		assert.NotNil(t, err)
		// Key: 'Person.content_value' Error:Field validation for 'content_value' failed on the 'e164' tag
	})

	t.Run("tag에 등록된 validate 먼저한다. (tag level에서 에러 -> struct level 검사 x)", func(t *testing.T) {
		invalid := Person{ContactType: "kakaotalk"}
		err := v.Struct(invalid)
		assert.NotNil(t, err)
		// Key: 'Person.ContactType' Error:Field validation for 'ContactType' failed on the 'oneof' tag
		// Key: 'Person.ContactValue' Error:Field validation for 'ContactValue' failed on the 'required' tag
	})
}

func StructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(Person)

	if user.ContactType == "email" &&
		(v.Var(user.ContactType, "email") != nil) {
		sl.ReportError(user.ContactValue, "content_value", "content_type", "email", "")
		// tag -> custom message로 연동된다.
	}

	if user.ContactType == "phone" &&
		(v.Var(user.ContactType, "e164") != nil) {
		sl.ReportError(user.ContactValue, "content_value", "content_type", "e164", "")
	}
}
