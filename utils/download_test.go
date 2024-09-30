package utils

import "testing"

func TestDownload(t *testing.T) {

	url := ""
	target := "./"
	err := Download(url, target)
	if err != nil {
		t.Error(err)
	}
}
