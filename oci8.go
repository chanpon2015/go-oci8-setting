package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	Oci8Mac = "prefixdir=[prefix]¥nlibdir=${prefixdir}¥nincludedir=${prefixdir}/sdk/include¥n¥nName: OCI¥nDescription: Oracle database driver¥nVersion: 12.2¥nLibs: -L${libdir} -lclntsh¥nCflags: -I${includedir}"
)

func create(settings Settings) error {
	if _, err := os.Stat(settings.PkgConfigPath); os.IsNotExist(err) {
		os.MkdirAll(settings.PkgConfigPath, 0755)
	}
	var oci8 string
	switch runtime.GOOS {
	case "darwin":
		oci8 = Oci8Mac
	}
	f, err := os.Create(filepath.Join(settings.PkgConfigPath, "oci8.pc"))
	if err != nil {
		return err
	}
	defer f.Close()
	oci8 = strings.Replace(oci8, "[prefix]", settings.ResourceHome, -1)
	f.Write(([]byte)(oci8))
	return nil
}
