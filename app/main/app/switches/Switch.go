package switches

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("Hello world")
		getCurrentTime()
		time.Sleep(time.Second)
	}
}

func GetKindOfOS() string {
	result := runtime.GOOS
	return result
}

func getCurrentTime() {
	fmt.Println(time.Now().Hour(), ":", time.Now().Minute(), ":", time.Now().Second())
}
