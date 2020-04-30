# WebRTCDemo
基于golang提供web服务器的WebRTC demo


# Demo列表
## 1. local_media_demo
简单的本地摄像头调用,将摄像头的内容放到video标签中进行播放.

## 2. local_screen_replay
简单的本地屏幕抓取\录制\下载.

# 使用
本项目采用go作为web服务器实现,webrtc实现采用纯js,只依赖jquery.

使用时请先配置号golang开发环境(除非你已经编译成了二进制包),然后运行:
```
go run webserver/main.go
```

web服务器就已经启动,默认端口是8081,你可以在`config/web_server.go`文件中进行修改.