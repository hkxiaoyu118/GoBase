package File

import (
	"archive/zip"
	"bytes"
	"errors"
	"github.com/hkxiaoyu/gobase/BaseString"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

// 判断文件是否存在
func FsIsFileExist(path string) (bool, error) {
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

// 读取指定路径文件的数据
func FsReadFile(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	defer FsCloseFile(f)
	return ioutil.ReadAll(f)
}

// 获取可执行文件的路径
func FsGetFilePath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	fileDir, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(fileDir, "/")
	if i < 0 {
		i = strings.LastIndex(fileDir, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(fileDir[0 : i+1]), nil
}

// 复制文件
func FsCopyFile(src string, des string) (int64, error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer FsCloseFile(srcFile)

	desFile, err := os.Create(des)
	if err != nil {
		return 0, err
	}
	defer FsCloseFile(desFile)
	return io.Copy(desFile, srcFile)
}

// 获取去除了后缀的文件名
func FsGetFileNameNoExt(filePath string) string {
	filenameWithSuffix := path.Base(filePath)                          //获取文件名带后缀
	fileSuffix := path.Ext(filenameWithSuffix)                         //获取文件后缀
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix) //获取文件名
	return filenameOnly
}

// 删除指定目录的所有文件
func FsRemoveAllFiles(path string) {
	_ = os.RemoveAll(path)
}

// 删除指定路径的文件
func FsRemoveFile(path string) {
	_ = os.Remove(path)
}

// 关闭文件句柄
func FsCloseFile(file *os.File) {
	_ = file.Close()
}

// 获取文件所在的文件目录
func FsGetDir(filePath string) string {
	return BaseString.StrSubString(filePath, 0, strings.LastIndex(filePath, "/"))
}

// 判断文件是否是zip压缩包
func IsZip(zipPath string) bool {
	fd, err := os.Open(zipPath)
	if err != nil {
		return false
	}
	defer FsCloseFile(fd)

	buf := make([]byte, 4)
	if n, err := fd.Read(buf); err != nil || n < 4 {
		return false
	}
	return bytes.Equal(buf, []byte("PK\x03\x04"))
}

// 解压文件
func UnZip(zipFile string, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}

	defer func() {
		_ = reader.Close()
	}()

	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}

		defer func() {
			_ = rc.Close()
		}()

		filename := dest + "/" + file.Name
		err = os.MkdirAll(getDir(filename), 0755)
		if err != nil {
			return err
		}

		w, err := os.Create(filename)
		if err != nil {
			return err
		}

		defer func() {
			_ = w.Close()
		}()

		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
	}
	return nil
}
