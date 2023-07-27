package main

import "net"

type User struct {
	//名字即是地址
	Name string
	Addr string
	//管道，用来通信
	C chan string
	//连接
	conn net.Conn
}

//创建一个用户API
func NewUser(conn net.Conn) *User {
	//1.获取远端对象ip
	userAddr := conn.RemoteAddr().String()

	//2.创建一个User对象，维护信息
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}

	//3.启动监听状态，如果有消息就发送给这个用户
	go user.ListenMessage()

	return user
}

//监听当前User channel方法，一旦有消息，就直接发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}
