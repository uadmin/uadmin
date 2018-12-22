package uadmin

import (
	"testing"
)

// TestGenerateByteArray is a unit testing function for generateByteArray() function
func TestGenerateByteArray(t *testing.T) {
	examples := []struct {
		length int
	}{
		{0},
		{10},
		{20},
		{100},
	}
	for _, e := range examples {
		code := generateByteArray(e.length)
		if len(code) != e.length {
			t.Errorf("length of generateByteArray(%d) = %d != %d", e.length, len(code), e.length)
		}
	}
}

// TestEncrypt is a unit testing function for encrypt() function
func TestEncrypt(t *testing.T) {
	examples := []struct {
		str string
	}{
		{""},
		{"a"},
		{"abc"},
		{" abc"},
		{" abc "},
		{"abcdefghijklmnopqrstuvwxyz"},
		{"a b c d e f g h i j k l m n o p q r s t u v w x y z"},
		{GenerateBase64(100)},
		{GenerateBase64(200)},
		{GenerateBase64(500)},
		{GenerateBase64(1000)},
		{GenerateBase64(10000)},
	}

	key := generateByteArray(32)

	for _, e := range examples {
		encrypted, err := encrypt(key, e.str)
		if err != nil {
			t.Errorf("error in running encrypt(\"%s\") %s", e.str, err)
			continue
		}
		decrypted, err := decrypt(key, encrypted)
		if err != nil {
			t.Errorf("error in running decrypt(\"%s\") %s", e.str, err)
			continue
		}
		if e.str != decrypted {
			t.Errorf("original text does't match decrypted text for (%s)", e.str)
		}
	}
}

// TestencryptRecord is a unit testing function for encryptRecord() function
func TestEncryptRecord(t *testing.T) {
	type TestModel struct {
		Model
		Name string `uadmin:"encrypt"`
	}
	//Schema = map[string]ModelSchema{}
	Schema["testmodel"], _ = getSchema(TestModel{})
	examples := []struct {
		str string
	}{
		{""},
		{"a"},
		{"abc"},
		{" abc"},
		{" abc "},
		{"abcdefghijklmnopqrstuvwxyz"},
		{"a b c d e f g h i j k l m n o p q r s t u v w x y z"},
		{GenerateBase64(100)},
		{GenerateBase64(200)},
		{GenerateBase64(500)},
		{GenerateBase64(1000)},
		{GenerateBase64(10000)},
	}
	registered = false
	for _, e := range examples {
		m := TestModel{Name: e.str}
		encryptRecord(&m)
		if m.Name != e.str {
			t.Errorf("TestEncryptRecord is encrypting without being registered")
		}
	}

	registered = true
	for _, e := range examples {
		m := TestModel{Name: e.str}
		encryptRecord(&m)
		decryptRecord(&m)
		if m.Name != e.str {
			t.Errorf("original text does't match decrypted text for using TestEncryptRecord (%s)", e.str)
		}
	}

	registered = false
	//Schema = map[string]ModelSchema{}
}

// TestencryptRecord is a unit testing function for encryptRecord() function
func TestEncryptArray(t *testing.T) {
	type TestModel struct {
		Model
		Name string `uadmin:"encrypt"`
	}
	//Schema = map[string]ModelSchema{}
	Schema["testmodel"], _ = getSchema(TestModel{})
	examples := []struct {
		str string
	}{
		{""},
		{"a"},
		{"abc"},
		{" abc"},
		{" abc "},
		{"abcdefghijklmnopqrstuvwxyz"},
		{"a b c d e f g h i j k l m n o p q r s t u v w x y z"},
		{GenerateBase64(100)},
		{GenerateBase64(200)},
		{GenerateBase64(500)},
		{GenerateBase64(1000)},
		{GenerateBase64(10000)},
	}
	registered = false
	for _, e := range examples {
		m := []TestModel{
			{Name: e.str},
			{Name: e.str},
			{Name: e.str},
			{Name: e.str},
			{Name: e.str},
		}

		for i := range m {
			encryptArray(&m[i])
			if m[i].Name != e.str {
				t.Errorf("TestEncryptRecord is encrypting without being registered")
			}
		}

	}

	registered = true
	for _, e := range examples {
		m := []TestModel{
			{Name: e.str},
			{Name: e.str},
			{Name: e.str},
			{Name: e.str},
			{Name: e.str},
		}

		for i := range m {
			encryptArray(&m)
			decryptArray(&m)
			if m[i].Name != e.str {
				t.Errorf("original text does't match decrypted text for using TestEncryptRecord (%s)", e.str)
			}
		}
	}

	registered = false
	//Schema = map[string]ModelSchema{}
}
