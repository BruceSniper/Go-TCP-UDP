package main

import (
	"fmt"
	"io"
	"net"
)

/*
使用Golang创建⼀个HTTP连接
	A.HTTP协议是基于TCP协议之上的⽂本协议。
	B.每行⽂本使⽤\r\n结尾，当连续两个\r\n时，表示整个数据包结束。
*/
func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Printf("dial failed, err:%v\n", err)
		return
	}

	data := "GET / HTTP/1.1\r\n"
	data += "HOST: www.baidu.com\r\n"
	data += "connection: close\r\n"
	data += "\r\n\r\n"

	//写入数据
	_, err = io.WriteString(conn, data)
	if err != nil {
		fmt.Printf("wirte string failed, err:%v\n", err)
		return
	}

	var buf [1024]byte
	for {
		//读取返回的数据
		n, err := conn.Read(buf[:])
		if err != nil || n == 0 {
			break
		}

		fmt.Println(string(buf[:n]))
	}
}
