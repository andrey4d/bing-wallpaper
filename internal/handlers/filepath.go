/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"io/fs"
	"os"
	"path/filepath"
)

func GetAbsPath(path string) (string, error) {

	if filepath.IsAbs(path) {
		return path, nil
	}

	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(pwd, path), err
}

func MakeDirAllIfNotExists(name string, chmod fs.FileMode) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		if err := os.MkdirAll(name, chmod); err != nil {
			return err
		}
	}
	return nil
}

func MakeDirIfNotExists(name string, chmod fs.FileMode) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		if err := os.Mkdir(name, chmod); err != nil {
			return err
		}
	}
	return nil
}
