package env

import (
	"os"
	"regexp"
	"strings"
)

/*
argsRegex 获取参数的正则表达式
如下：都应该是一个环境参数
main.go --a=b
main.go -a=b
main.go --a b
main.go -a b

# 下面的一个例子中应当拿到环境参数 a=b e=f t=h
go run os.go --a b c d --e f -t=h
# 使用 FindStringSubmatch 得到结果如下
[]
[--a a  ]
[]
[]
[]
[--e e  ]
[]
[-t=h t = h]
我们应该判断len > 2
*/
var (
	argsRegex = regexp.MustCompile(`^-{1,2}([\w?.\-]+)(=)?(.*)$`)
)

func init() {
	InitArgs(os.Args)
	initOSEnv()
}

func initOSEnv() {
	osEnvs := os.Environ()
	// 遍历环境变量并填充map
	for _, env := range osEnvs {
		// 拆分键值对
		parts := strings.SplitN(env, "=", 2)
		if len(parts) == 2 {
			key, value := parts[0], parts[1]
			// 将键值对存入map
			envs[key] = value
		}
	}
}

func InitArgs(args []string) {
	size := len(args)
	for i := 0; i < size; {
		// array要么为0要么为4
		array := argsRegex.FindStringSubmatch(args[i])
		if len(array) > 2 {
			if array[2] != "=" && i < size-1 && args[i+1][0] != '-' {
				envs[array[1]] = args[i+1]
				i += 2
				continue
			}
			envs[array[1]] = array[3]
		}
		i++
	}
}
