package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subodhqss/testing/models"
)

func TestGetUser(t *testing.T) {
	expResult := &models.User{
		ID:   1,
		Name: "Alam",
	}

	arg := []int{}

	// for empty arr as an arg
	res := GetUser(arg)
	assert.Nil(t, res)

	arg1 := []int{1, 2, 3}
	// for tesing result
	res1 := GetUser(arg1)
	assert.Equal(t, expResult, res1)

}
