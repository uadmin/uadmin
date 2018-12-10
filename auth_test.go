package uadmin

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// TestGenerateBase64 is a unit testing function for GenerateBase64() function
func TestGenerateBase64(t *testing.T) {
	examples := []struct {
		length int
	}{
		{0},
		{1},
		{10},
	}

	for _, e := range examples {
		code := GenerateBase64(e.length)
		if len(code) != e.length {
			t.Errorf("lenght of GenerateBase64(%d) = %d != %d", e.length, len(code), e.length)
		}
	}
}

// TestGenerateBase32 is a unit testing function for GenerateBase32() function
func TestGenerateBase32(t *testing.T) {
	examples := []struct {
		length int
	}{
		{0},
		{1},
		{10},
	}

	for _, e := range examples {
		code := GenerateBase32(e.length)
		if len(code) != e.length {
			t.Errorf("lenght of GenerateBase32(%d) = %d != %d", e.length, len(code), e.length)
		}
	}
}

// TestHashPass is a unit testing function for hashPass() function
func TestHashPass(t *testing.T) {
	examples := []struct {
		pass string
	}{
		{"1234"},
		{"abc123"},
		{"password"},
		{"password1"},
		{"Password1"},
		{" Password1 "},
		{"Pass 123"},
		{"Pass 123!"},
		{"Pass 123! "},
		{"كلمة السر 123! "},
		{GenerateBase64(10)},
		{GenerateBase64(20)},
		{GenerateBase64(30)},
		{GenerateBase64(40)},
		{GenerateBase64(50)},
		{GenerateBase64(60)},
		{GenerateBase64(70)},
	}

	bcryptDiff = 5

	Salt = ""
	for _, e := range examples {
		code := hashPass(e.pass)
		if bcrypt.CompareHashAndPassword([]byte(code), []byte(e.pass+Salt)) != nil {
			t.Errorf("hashPass(\"%s\") invalid denied password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte("1"+e.pass+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte("a"+e.pass+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte(" "+e.pass+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte(e.pass+" "+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte("234"+code), []byte(e.pass+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
	}

	Salt = GenerateBase64(72)
	for _, e := range examples {
		code := hashPass(e.pass)
		if bcrypt.CompareHashAndPassword([]byte(code), []byte(e.pass+Salt)) != nil {
			t.Errorf("hashPass(\"%s\") invalid denied password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte(Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte("1"+e.pass+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte("a"+e.pass+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte(" "+e.pass+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte(e.pass+" "+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte(code), []byte(e.pass+" "+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
		if bcrypt.CompareHashAndPassword([]byte("234"+code), []byte(e.pass+Salt)) == nil {
			t.Errorf("hashPass(\"%s\") invalid accepted password with salt %s", e.pass, Salt)
		}
	}

	Salt = ""
}
