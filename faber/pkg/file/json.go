package file

import (
	"fmt"
	"os"
)

func ExportJsonFile(data []byte, folder string, file string) error {
	f, err := os.OpenFile(fmt.Sprintf("%s/%s", folder, file), os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	if nil != err {
		return err
	}
	_, err = f.Write(data)
	if nil != err {
		return err
	}
	return f.Sync()
}
