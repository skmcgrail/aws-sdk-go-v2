package gomod

import (
	"bytes"
	"fmt"
	"golang.org/x/mod/modfile"
	"io"
	"os"
	"path/filepath"
)

func GetModulePath(file *modfile.File) (string, error) {
	if file.Module == nil {
		return "", fmt.Errorf("module directive not present")
	}
	return file.Module.Mod.Path, nil
}

func LoadModuleFile(path string, fix modfile.VersionFixer) (*modfile.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ReadModule(path, f, fix)
}

func ReadModule(path string, f io.Reader, fix modfile.VersionFixer) (*modfile.File, error) {
	fBytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	parse, err := modfile.Parse(path, fBytes, fix)
	if err != nil {
		return nil, err
	}

	return parse, nil
}

func WriteModuleFile(path string, file *modfile.File) (err error) {
	modPath := filepath.Join(path, "go.mod")

	var mf *os.File
	mf, err = os.OpenFile(modPath, os.O_CREATE|os.O_TRUNC, 644)
	if err != nil {
		return err
	}
	defer func() {
		fErr := mf.Close()
		if fErr != nil && err == nil {
			err = fErr
		}
	}()

	var fb []byte
	fb, err = file.Format()
	if err != nil {
		return err
	}

	_, err = io.Copy(mf, bytes.NewReader(fb))

	return err
}
