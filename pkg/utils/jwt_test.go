package utils

import "testing"

func TestGenToken(t *testing.T) {
	token, err := GenToken(uint32(123))
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("token: %v\n", token)
}

func TestVerifyToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MTIzLCJpc3MiOiJwYWxhY2UiLCJzdWIiOiJhdXRoIiwiZXhwIjoxNzEzNTMzMzMzLCJpYXQiOjE3MTM0NDY5MzN9.f3Qe7ij2vdjTI1kNV2DNj09V_70J5EU8zE0Znqb3EVc"
	claims, err := VerifyToken(token)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("%v\n", claims)
}
