package _03_error

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	Name string `validate:"required"`
	Age  int    `validate:"required"`
}

func TestError(t *testing.T) {
	v := validator.New()

	p := Person{}
	err := v.Struct(p)

	t.Run("valiator의 에러", func(t *testing.T) {
		validationErrors := err.(validator.ValidationErrors) // 타입 변환하여 사용
		assert.Equal(t, 2, len(validationErrors))   // 배열로 발생한 에러 모두 저장하고 있음
		assert.NotNil(t, validationErrors[0].Tag())
		assert.NotNil(t, validationErrors[0].Value())
	})
}
