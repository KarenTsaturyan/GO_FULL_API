package jwt_test

import (
	"http_5/pkg/jwt"
	"testing"
)

// unit
func TestJWTCreate(t *testing.T) {
	const email = "a@a.com"
	jwtService := jwt.NewJWT("/2+XktJGz2j5ehICI/5K9lL+UshtE1LxS7mfT+qar5w=")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is invalid")
	}

	if data.Email != email {
		t.Fatalf("Email %s not equal %s", data.Email, email)
	}
}
