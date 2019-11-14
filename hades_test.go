package hades

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDash(t *testing.T) {
	assert := assert.New(t)

	config := GetConfig("test.yaml", []string{"", "testConfig"})
	fmt.Println(config.Data)
	test := config.Map("test")
	assert.NotZero(test.Int("testInt"))
	assert.NotZero(test.Bool("testBool"))
	assert.NotZero(test.Float("testFloat"))
	assert.NotZero(test.Str("testString"))
	assert.NotEmpty(test.List("testList"))
	assert.NotZero(test.Interface("testInt"))

	assert.Panics(func() {
		getPath(
			"nopathg",
			[]string{"nopath"},
		)
	})
}
