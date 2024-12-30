package auth

import (
	"fanxan/web-service-gin/infra/response"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthEntity(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "fannan@gmail.com",
			Password: "123456",
		}
		err := authEntity.Validate()
		require.Nil(t, err)
	})
	t.Run("email is required", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "",
			Password: "123456",
		}
		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailRequired, err)
	})
	t.Run("email is invalid", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "fannan.id",
			Password: "123456",
		}
		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailInvalid, err)
	})

	t.Run("password is required", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "fannan@gmail.com",
			Password: "",
		}
		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordRequired, err)
	})
	t.Run("password must have minimum 6 character", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "fannan.id@gmail.com",
			Password: "12345",
		}
		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordInvalidLength, err)
	})
}

func TestEncryptPassword(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "fan@gmail.com",
			Password: "rahasia",
		}

		err := authEntity.EncryptPassword(bcrypt.DefaultCost)
		require.Nil(t, err)
		fmt.Printf("%+v\n", authEntity)
	})
}
