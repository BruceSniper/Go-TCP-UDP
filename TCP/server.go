package main

import (
	"fmt"
	"net"
)
/*
1.服务端处理理流程
	a.监听端口
	b.接受客户端的链接
	c.创建Goroutine，处理这个链接(⼀个服务端要链接多个客户端，所以使用Goroutine⾮常简单)
	要是用Java、C#服务这边每⼀个请求都开⼀个线程处理的话，顶多⼏千个，但是Goroutine就⾮常简单。
*/
func main() {
	//1.建立监听端口
	listener, err := net.Listen("tcp", "0.0.0.0:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	fmt.Println("listen Start...:")

	for {
		//2.接收客户端的链接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		//3.开启一个Goroutine，处理链接
		go process(conn)
	}
}

//处理请求，类型就是net.Conn
func process(conn net.Conn) {
	//处理结束后关闭链接
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v", err)
			break
		}
		fmt.Printf("recv from client, content:%v\n", string(buf[:n]))
	}
}
