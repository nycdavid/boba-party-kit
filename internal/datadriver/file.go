package datadriver

import "github.com/nycdavid/boba-party-kit/internal/config"

type (
	File struct {
		cfg config.File
	}
)

func NewFile(cfg config.File) *File {
	return &File{cfg: cfg}
}

func (f *File) Fetch() ([]byte, error) {
	return []byte{}, nil
}
