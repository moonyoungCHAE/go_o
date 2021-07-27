package _01_basic

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
1. validator 생성하여 사용 (struct 정보 캐싱 되기 때문에 하나의 객체로 사용)
2. tag 기반 사용 (제공하는 tag, custom tag)
3. v.Struct()를 호출할 때 검증된다.
*/

type Person struct {
	Name string `validate:"required"`
	Age  int    `validate:"required"`
}

func TestBasic(t *testing.T) {
	v := validator.New()

	t.Run("valid", func(t *testing.T) {
		validPerson := Person{
			Name: "tommy", Age: 26,
		}

		err := v.Struct(validPerson)
		assert.Nil(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		validPerson := Person{
			Name: "tommy",
		}

		err := v.Struct(validPerson)
		assert.NotNil(t, err)
		// Key: 'Person.Age' Error:Field validation for 'Age' failed on the 'required' tag
	})

	t.Run("단일 변수를 검사하고 싶을 때", func(t *testing.T) {
		err := v.Var("hi@", "email")
		assert.NotNil(t, err)
		// Key: '' Error:Field validation for '' failed on the 'email' tag
	})
}
