package usecase

import (
	"strings"
	"testing"
)

// Test_generateRandomToken : generateRandomTokenのテスト、指定された文字列の長さと規定の文字列のみを含むかを確認する
func Test_generateRandomToken(t *testing.T) {
	token, err := generateRandomToken(6)
	if err != nil {
		t.Error(err)
	}
	if len(token) != 6 {
		t.Error("token length is invalid")
	}

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for _, v := range token {
		if !strings.Contains(letters, string(v)) {
			t.Error("token contains invalid character")
		}
	}
}

func Test_CreateRoom(t *testing.T) {

}

func Test_FindRoom(t *testing.T) {}

func Test_JoinRoom(t *testing.T) {}
