package fileutil

import (
	"bufio"
	"io"
	"os"
)

// ReadFile 文件读取操作
// 按行读取文件
func ReadFile(file io.Reader) []string {
	var result []string
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()

		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}

		result = append(result, string(line))

	}
	return result
}

func Read(fileName string) []string {

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return ReadFile(file)
}

func ReadBytes(fileName string) []byte {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil
	}
	return data
}
