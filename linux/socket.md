# socket (套接字)
socket 是计算机网络 `应用层` 与 `TPC/IP`协议族通信的中间软件抽象层，是一组接口。

网络中的进程之间通信，首先需要标识进程在网络中的唯一标识。我们可以使用【ip:port】来实现。再加上使用的通信协议。我们最后可以使用【ip, 协议， port】来描述网络通信。而socket就是利用这个三元组来解决网络通信。

## socket的两种传输方式
- SOCKET_STREAM
    - 面向连接。数据准确无误的传输到另一端，如果数据出现损坏或者丢失，可以重新发送。HTTP协议就是用这个方式进行传输数据的。

- SOCKET_DGRAM
    - 无连接。计算机只管传输数据，不校验数据。数据错了也无法重传，效率比较高。视频聊天，语音聊天，直播等通信就是用这个方式传输数据。

## socket 常用的函数接口
- Server
    - `socket()`：创建socket
    - `bind()`：绑定socket与端口号
    - `listen()`：监听端口号
    - `accept()`：接收客户端请求
    - `recv`：从socket 中读取字符
    - `close()`：关闭socket
- Client
    - `socket()`：创建socket
    - `connect()`：连接指定计算机的端口
    - `send()`：向socket中写入信息
    - `close()`：关闭socket