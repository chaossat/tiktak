package middleware

import (
	"fmt"
	"testing"
)

func TestIsFollow(t *testing.T) {
	token, err := CreateToken(int64(2))
	fmt.Println(token, "\n", err)
	res, err := CheckToken(token)
	fmt.Println(res, "\n", err)
}
