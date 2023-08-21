package fileutil

import (
	"path/filepath"
	"strings"
)

// ExtName 对文件名的操作
// 扩展名
func ExtName(path string) string {
	return strings.TrimLeft(Ext(path), Point)
}

// Ext 扩展
func Ext(path string) string {
	ext := filepath.Ext(path)
	if p := strings.IndexByte(ext, '?'); p != -1 {
		ext = ext[0:p]
	}
	return ext
}

// FileName 返回文件名
// /opt/temp.txt -> temp.txt
func FileName(path string) string {
	return filepath.Base(path)
}
