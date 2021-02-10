package util

import (
	"testing"
)

func TestPasswordHashing(t *testing.T) {
	const pass = "QSLDJFNEjvnlshf123"
	if res := PasswordHashing(pass); res == pass {
		t.Errorf("for pass %s got %s", pass, res)
	}
}
