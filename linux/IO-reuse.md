# Linux 中网络I/O复用并发模型

## 什么是阻塞，非阻塞，IO多路复用？
- 流
    - 可进行I/O操作的内核对象
        - 文字
        - 管道
        - 套接字
    - 流的入口：文件描述符(fd)

- 阻塞等待
不占用CPU时间片

    - 缺点
        - 不能处理多个I/O请求的问题
        - 同一个阻塞，同一时间只能处理一个流的阻塞监听
- 非阻塞（忙轮询）
占用资源不断去查看资源锁有没有释放

- 多路IO复用
既能阻塞等待，不浪费资源，又能同一时刻监听多个IO请求的状态

## IO复用解决的问题
如何解决大量IO读写问题
- 忙轮询
```
while true {
	for i in 流[] {
		if i has 数据 {
			读 或者 其他处理
		}
	}
}
```
同时和多个流访问。

- select
```
while true {
	select(流[]); //阻塞

  //有消息抵达
	for i in 流[] {
		if i has 数据 {
			读 或者 其他处理
		}
	}
}
```
select负责接收IO，但不能精准直到哪些IO可用，还需要遍历。监听的IO有限，一般1024个，与平台无关

- epoll
```
while true {
	可处理的流[] = epoll_wait(epoll_fd); //阻塞

  //有消息抵达，全部放在 “可处理的流[]”中
	for i in 可处理的流[] {
		读 或者 其他处理
	}
}
```
类似select,但能返回所有可以处理的流，不必再去遍历确认

## epoll是什么
- linux独有
- 一种I/O 多路复用技术
- 只关心活跃的链接，无需遍历所有的描述符集合。
- 能处理大量链接请求（系统可以打开的文件数目 `cat /proc/sys/fs/file-max`）

## API
1. 创建
```
/** 
 * @param size 告诉内核监听的数目 
 * 
 * @returns 返回一个epoll句柄（即一个文件描述符） 
 */
int epoll_create(int size);
```
2. 控制
```
/**
* @param epfd 用epoll_create所创建的epoll句柄
* @param op 表示对epoll监控描述符控制的动作
*
* EPOLL_CTL_ADD(注册新的fd到epfd)
* EPOLL_CTL_MOD(修改已经注册的fd的监听事件)
* EPOLL_CTL_DEL(epfd删除一个fd)
*
* @param fd 需要监听的文件描述符
* @param event 告诉内核需要监听的事件
*
* @returns 成功返回0，失败返回-1, errno查看错误信息
*/
int epoll_ctl(int epfd, int op, int fd,
struct epoll_event *event);


struct epoll_event {
	__uint32_t events; /* epoll 事件 */
	epoll_data_t data; /* 用户传递的数据 */
}

/*
 * events : {EPOLLIN, EPOLLOUT, EPOLLPRI,
						 EPOLLHUP, EPOLLET, EPOLLONESHOT}
 */
typedef union epoll_data {
	void *ptr;
	int fd;
	uint32_t u32;
	uint64_t u64;
} epoll_data_t;
```
3. 等待
