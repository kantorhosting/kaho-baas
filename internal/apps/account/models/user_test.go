package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_JSONOmitsPassword(t *testing.T) {
	user := User{
		Email:    "john@demo.com",
		Password: "$2a$14$hashedpassword",
	}

	body, err := json.Marshal(user)

	assert.Nil(t, err)
	assert.NotContains(t, string(body), `"password":`)
	assert.NotContains(t, string(body), user.Password)
}
