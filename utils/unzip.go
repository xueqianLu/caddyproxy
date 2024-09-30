package utils

import (
	"archive/zip"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// unzip unzip zip file to dest directory.
func Unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return errors.New(fmt.Sprintf("open zip file failed: %s", err))
	}
	defer r.Close()
	// if dest directory not exists, create it.
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		if err := os.MkdirAll(dest, os.ModePerm); err != nil {
			return errors.New(fmt.Sprintf("create dest directory failed: %s", err))
		}
	}

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		// 防止 ZipSlip 漏洞
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			log.WithField("fpath", fpath).Error("invalid file path, ignore to unzip")
			continue
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return errors.New(fmt.Sprintf("create file directory failed: %s", err))
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}
