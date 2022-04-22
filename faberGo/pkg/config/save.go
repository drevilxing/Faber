package config

import (
	"encoding/json"
	"os"
)

func (that *GenerateConfig) SaveToPath(path string) error {
	data, err := json.Marshal(*that)
	file, err := os.OpenFile(path+"/"+that.Key+".json", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	if nil != err {
		return err
	}
	_, err = file.Write(data)
	if nil != err {
		return err
	}
	return file.Sync()
}

func Remove(path string, key string) error {
	return os.Remove(path + "/" + key + ".json")
}
