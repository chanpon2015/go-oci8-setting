package main

import "path/filepath"

type Settings struct {
	ResourceHome  string
	InstantHome   string
	PkgConfigPath string
	Path          ResourcePath
}

type ResourcePath struct {
	InstantClientBasic   string
	InstantClientSqlPlus string
	InstantClientSdk     string
}

func (s *Settings) ResourcePaths(currentPath string) []string {
	return []string{
		filepath.Join(currentPath, s.Path.InstantClientBasic),
		filepath.Join(currentPath, s.Path.InstantClientSdk),
		filepath.Join(currentPath, s.Path.InstantClientSqlPlus),
	}
}
