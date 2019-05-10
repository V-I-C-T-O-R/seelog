package seelog

import (
	"log"
	"time"
	"sync"
	"os"
)

//  启动seelog
func See(filePath string,port int,wg *sync.WaitGroup) bool {

	// 检查参数
	if !checkParam(filePath,port) || !fileChecker(filePath){
		wg.Done()
		return false
	}

	// 开启socket管理器
	go manager.start()
	// 监控文件
	go monitor(filePath)
	// 开启httpServer
	go server(port)

	//等待服务运行起再返回，否则可能导致开头的部分日志无法输出到网页
	time.Sleep(200 * time.Millisecond)
	return true
}

// 参数验证
func checkParam(filePath string,port int) bool {
	if filePath == "" {
		log.Println("filePath 不可为空")
		return false
	}

	if port == 0 {
		log.Println("port 不可为空")
		return false
	}
	return true
}

// 判断资源包是否存在
func fileChecker(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}else {
		log.Println("没有此日志文件:"+filename)
		return false
	}
}


