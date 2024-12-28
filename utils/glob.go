package utils

import "path/filepath"

type GlobFunc interface {
	Glob(pattern string) ([]string, error)
}

type GlobFuncImpl struct{}

func NewGlobFunc() GlobFunc {
	return &GlobFuncImpl{}
}

func (gf *GlobFuncImpl) Glob(pattern string) ([]string, error) {
	return filepath.Glob(pattern)
}
