# WebRTCDemo
基于golang提供web服务器的WebRTC demo


# Demo列表
## 1. local_media_demo
简单的本地摄像头调用,将摄像头的内容放到video标签中进行播放.

## 2. local_screen_replay
简单的本地屏幕抓取\录制\下载.

## 3. signal_server_demo
一个简单的信令服务器通信的demo

## 4.local_peerconnection_demo
本地端对端通信，除了进行信令服务器交换信息外，已具备webrtc实现一对一视频聊天的基础。

## 5. remote_chat
远程一对一视频聊天。由于webrtc不允许远程使用http调用摄像头等媒体设备，因此这里也将web服务器改成了https，使用openssl创建了临时证书，放在ssl目录下。使用时请注意忽略浏览器安全报警。
演示地址：`https://134.175.26.82:8082/static/demo/remote_chat.html`.
![image](https://github.com/243286065/pictures_markdown/blob/master/media/webrtc/cbf6e68b090305689fbfeeb8a9060fba.png?raw=true)

# 使用
本项目采用go作为web服务器实现,webrtc实现采用纯js,只依赖jquery.

使用时请先配置号golang开发环境(除非你已经编译成了二进制包),可以参考`https://studygolang.com/dl`.

然后运行:
```
go run webserver/main.go
```

web服务器就已经启动,默认端口是8081,你可以在`config/web_server.go`文件中进行修改.