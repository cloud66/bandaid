package bandaid

import (
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DirExists(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}

	return true
}

func ChangeExt(filename string, newExt string) string {
	ext := path.Ext(filename)
	return filename[0:len(filename)-len(ext)] + "." + newExt
}

func HomeDir() string {
	dir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func ConfigDir() string {
	return path.Join(HomeDir(), ".cloud66")
}

func ExpandTilde(path string) string {
	dir := HomeDir()
	if path == "~" {
		path = dir
	} else if strings.HasPrefix(path, "~/") {
		path = filepath.Join(dir, path[2:])
	}

	return path
}

func TrimExt(path string) string {
	return strings.TrimSuffix(path, filepath.Ext(path))
}

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func DeleteFile(path string) error {
	if !FileExists(path) {
		return nil
	}

	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}
