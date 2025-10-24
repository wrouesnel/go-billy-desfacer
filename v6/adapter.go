package desfacer

import (
	"io/fs"
	"os"
)

type dentryShim struct {
	os.FileInfo
}

func (d *dentryShim) Type() fs.FileMode {
	return d.Mode()
}

func (d *dentryShim) Info() (fs.FileInfo, error) {
	return d.FileInfo, nil
}
