package random

import (
	"crypto/rand"
	"github.com/shura1014/common/goerr"
)

var (
	// 提前缓存一万个随机值
	bufferChan = make(chan []byte, 10000)
)

func init() {
	go generateRandomLoop()
}

func generateRandomLoop() {
	for {
		buffer := make([]byte, 1024)
		if n, err := rand.Read(buffer); err != nil {
			panic(goerr.Text("read rand error"))
		} else {
			// 为什么缓存4，因为一个int刚好是4个字节
			for i := 0; i <= n-4; i += 4 {
				// 达到阈值10000后阻塞
				bufferChan <- buffer[i : i+4]
			}
		}
	}
}
