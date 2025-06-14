package jwt_test

import (
	"go/adv-demo/pkg/jwt"
	"testing"
)

func TestJwtSign(t *testing.T) {
	const email = "asan02"
	jwtService := jwt.NewJWT("12345")
	res, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})

	if err != nil {
		t.Fatal(err)
	}

	isValid, data := jwtService.Parse(res)

	if !isValid {
		t.Fatal("invalid jwt")
	}

	if data.Email != email {
		t.Fatal("email not equal")
	}

}
