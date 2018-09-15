package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

func main() {
	os.Setenv("test", "test")
	var settings Settings
	currentPath, _ := os.Executable()
	currentPath = filepath.Dir(currentPath)
	_, err := toml.DecodeFile(filepath.Join(currentPath, "settings.toml"), &settings)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := resourceCheck(settings, currentPath); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := os.Stat(settings.InstantHome); os.IsNotExist(err) {
		err = os.MkdirAll(settings.InstantHome, 0755)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		err = os.RemoveAll(filepath.Join(settings.InstantHome))
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.MkdirAll(settings.InstantHome, 0755)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	for _, p := range settings.ResourcePaths(currentPath) {
		err = unZip(p, settings.ResourceHome)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = create(settings)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func resourceCheck(settings Settings, currentPath string) error {
	for _, path := range settings.ResourcePaths(currentPath) {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return fmt.Errorf("resource file not found: %s", path)
		}
	}
	return nil
}

func unZip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
