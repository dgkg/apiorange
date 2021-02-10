package util

import (
	"log"
	"testing"
)

func TestNewJWT(t *testing.T) {
	res, err := NewJWT("mySigningKey", "Bob")
	t.Log(res, err)
	log.Println(res)
}
