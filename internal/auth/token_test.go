package auth

import "testing"

func TestTokenHashStable(t *testing.T) {
	h := HashToken("abc")
	if h == "abc" || h != HashToken("abc") {
		t.Fatal("hash")
	}
}
func TestPasswordVerification(t *testing.T) {
	h := HashPassword("p")
	if !VerifyPassword(h, "p") || VerifyPassword(h, "q") {
		t.Fatal("password")
	}
}
func TestTokenRandom(t *testing.T) {
	a, _ := NewToken()
	b, _ := NewToken()
	if a == b || len(a) != 64 {
		t.Fatal("token")
	}
}
