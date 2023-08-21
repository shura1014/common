package fileutil

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const Point = "."

var (
	Separator = string(filepath.Separator)
)

// Dir 目录
func Dir(path string) string {
	if strings.HasPrefix(path, Point) {
		// 相对路径转成绝对路径
		return filepath.Dir(AbsPath(path))
	}
	return filepath.Dir(path)
}

// AbsPath 绝对路径，但是不检查是否存在
func AbsPath(path string) string {
	p, err := filepath.Abs(path)
	if err != nil {
		return ""
	}
	return p
}

// RealPath 绝对路径
func RealPath(path string) string {
	p, err := filepath.Abs(path)
	if err != nil {
		return ""
	}
	if !Exists(p) {
		return ""
	}
	return p
}

func Exists(path string) bool {
	if stat, err := os.Stat(path); stat != nil && !os.IsNotExist(err) {
		return true
	}
	return false
}

func Join(paths ...string) string {
	var s string
	for _, path := range paths {
		if s != "" {
			s += Separator
		}
		s += strings.TrimRight(path, Separator)
	}
	return s
}

func JoinCurrent(paths ...string) string {
	pwd, err := os.Getwd()
	if err != nil {
		return ""
	}

	return Join(pwd, Join(paths...))
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

func ListFileName(path string) []string {
	fileNames := make([]string, 0)
	if IsDir(path) {
		files, err := os.ReadDir(path)
		if err != nil {
			return nil
		}
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames
}

func DirFunc(dir string, f func()) {
	pwd, _ := os.Getwd()
	defer func() {
		_ = os.Chdir(pwd)
	}()
	_ = os.Chdir(dir)
	f()
}

// CallDir 返回调用路径
func CallDir(dir string, skips ...int) string {
	if dir == "" {
		dir = Point
	}
	if strings.HasPrefix(dir, "/") {
		// 绝对路径
		return dir
	}

	// 	相对路径
	var (
		realPath string
		skip     = 2
	)

	if len(skips) > 0 {
		skip += skips[0]
	}

	_, file, _, _ := runtime.Caller(skip)
	if file == "" {
		return ""
	}
	DirFunc(Dir(file), func() {
		realPath = RealPath(dir)
	})
	return realPath
}
