package utils

import (
	"io"
	"net/http"
	"os"
)

func Download(url string, target string) error {
	// download the resource from the URL and save it to the target path.
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
