# seelog [![License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://github.com/xmge/seelog/blob/master/LICENSE)

### 项目介绍
* 项目fork于[seelog](https://github.com/xmge/seelog)
* 细微修改了部分为了linux部署
* 借鉴并学习

### 使用方式
* git clone https://github.com/V-I-C-T-O-R/seelog
* 修改main/main.go中的端口和日志位置
* go bulid -o name main/main.go
* 上传对应的编译文件'name'到对应的服务器对应的地址
* 执行nohup /path/name &
* 在浏览器中访问 *http://host:port

### 项目展示
![image](https://github.com/V-I-C-T-O-R/seelog/blob/master/demo.gif)

## 项目demo
* 检测单个日志文件：[https://github.com/V-I-C-T-O-R/seelog/blob/master/example/e1.go](https://github.com/V-I-C-T-O-R/seelog/blob/master/example/e1.go)
