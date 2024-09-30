package utils

import "testing"

func TestUnzip(t *testing.T) {

	f := "a.zip"
	d := "a"
	err := Unzip(f, d)
	if err != nil {
		t.Error(err)
	}
}
