package uadmin

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"math/big"
	"reflect"
)

// generateByteArray generates a base64 string of length length
func generateByteArray(length int) []byte {
	base := new(big.Int)
	base.SetString("256", 10)

	tempKey := []byte{}
	var tempByte *big.Int
	for i := 0; i < length; i++ {
		tempByte, _ = rand.Int(rand.Reader, base)
		tempKey = append(tempKey, byte(tempByte.Uint64()))
	}
	return tempKey
}

func encrypt(key []byte, message string) (encmess string, err error) {
	plainText := []byte(message)

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	//returns to base64 encoded string
	encmess = base64.URLEncoding.EncodeToString(cipherText)
	return
}

func decrypt(key []byte, securemess string) (decodedmess string, err error) {
	cipherText, err := base64.URLEncoding.DecodeString(securemess)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("ciphertext block size is too short")
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	decodedmess = string(cipherText)
	return
}

func decryptArray(a interface{}) {
	if !registered {
		return
	}
	if schema, ok := getSchema(getModelName(a)); ok {
		for _, f := range schema.Fields {
			if f.Encrypt {
				// TODO: Decrypt
				allArray := reflect.ValueOf(a)
				for i := 0; i < allArray.Elem().Len(); i++ {
					encryptedValue := allArray.Elem().Index(i).FieldByName(f.Name).String()
					decryptedValue, _ := decrypt(EncryptKey, encryptedValue)
					allArray.Elem().Index(i).FieldByName(f.Name).Set(reflect.ValueOf(decryptedValue))
				}
			}
		}
	}
}

func encryptArray(a interface{}) {
	if !registered {
		return
	}
	if schema, ok := getSchema(getModelName(a)); ok {
		for _, f := range schema.Fields {
			if f.Encrypt {
				allArray := reflect.ValueOf(a)
				for i := 0; i < allArray.Elem().Len(); i++ {
					encryptedValue := allArray.Elem().Index(i).FieldByName(f.Name).String()
					decryptedValue, _ := encrypt(EncryptKey, encryptedValue)
					allArray.Elem().Index(i).FieldByName(f.Name).Set(reflect.ValueOf(decryptedValue))
				}
			}
		}
	}
}

func decryptRecord(a interface{}) {
	if !registered {
		return
	}
	if schema, ok := getSchema(getModelName(a)); ok {
		for _, f := range schema.Fields {
			if f.Encrypt {
				recordValue := reflect.ValueOf(a)
				encryptedValue := recordValue.Elem().FieldByName(f.Name).String()
				decryptedValue, _ := decrypt(EncryptKey, encryptedValue)
				recordValue.Elem().FieldByName(f.Name).Set(reflect.ValueOf(decryptedValue))
			}
		}
	}
}

func encryptRecord(a interface{}) {
	if !registered {
		return
	}
	if schema, ok := getSchema(getModelName(a)); ok {
		for _, f := range schema.Fields {
			if f.Encrypt {
				recordValue := reflect.ValueOf(a)
				encryptedValue := recordValue.Elem().FieldByName(f.Name).String()
				decryptedValue, _ := encrypt(EncryptKey, encryptedValue)
				recordValue.Elem().FieldByName(f.Name).Set(reflect.ValueOf(decryptedValue))
			}
		}
	}
}
