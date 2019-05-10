package main

import (
	"sync"
	"seelog"
	"strconv"
)

func main() {
	var logFile = "../test.log"
	var port = 8002

	var wg sync.WaitGroup
	wg.Add(1)
	var status = seelog.See(logFile, port,&wg)
	if status {
		println("日志文件:"+logFile+"\n网页端口:"+strconv.Itoa(port)+"\n开始监控日志...")
	}
	wg.Wait()
}