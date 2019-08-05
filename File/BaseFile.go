package File

import (
	"errors"
	"os"
)

// 判断文件是否存在
func IsFileExist(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err == nil {
		if stat.Mode()&os.ModeType == 0 {
			return true, nil
		}
		return false, errors.New(path + " exist but is not regular file")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
