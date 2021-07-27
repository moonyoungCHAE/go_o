package _02_custom_tag

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
1. customTag 등록: RegisterValidation
2. Reflection으로 변수 값을 본다.
*/

type Person struct {
	Name    string `validate:"required"`
	Age     int    `validate:"required"`
	Contact string `validate:"contact"`
}

func TestCustomTag(t *testing.T) {
	v := validator.New()

	v.RegisterValidation("contact", func(fl validator.FieldLevel) bool {
		// fl.Field().Interface().(string)
		return fl.Field().String() == "email" || fl.Field().String() == "phone"
	})

	t.Run("(정상) 연락 수단은 이메일이나 전화번호이다.", func(t *testing.T) {
		validPerson := Person{"Tommy", 26, "email"}
		err := v.Struct(validPerson)
		assert.Nil(t, err)
	})

	t.Run("(예외) 연락 수단이 카카오톡이다.", func(t *testing.T) {
		validPerson := Person{"Tommy", 26, "kakaotalk"}
		err := v.Struct(validPerson)
		assert.NotNil(t, err)
	})
}

// cf. 위의 예는 아래 oneof로도 구현 가능
type Request struct {
	Country string `validate:"oneof=us uk fr mv es de be"`
}

func TestOneof(t *testing.T) {
	v := validator.New()

	t.Run("valid", func(t *testing.T) {
		r := Request{Country: "us"}
		err := v.Struct(r)
		assert.Nil(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		r := Request{Country: "kor"}
		err := v.Struct(r)
		assert.NotNil(t, err)
	})
}
