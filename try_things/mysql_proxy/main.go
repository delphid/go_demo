package main

import (
	"database/sql"
	"fmt"
	"net"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/proxy"
)

func main() {
	// 设置socks5代理
	dialer, err := proxy.SOCKS5("tcp", "localhost:7891", nil, proxy.Direct)
	if err != nil {
		panic(err)
	}

	// 通过代理连接mysql8.0
	//conn, err := dialer.Dial("tcp", "rm-k2js0l9aug8l04v13.mysql.zhangbei.rds.aliyuncs.com:3306")
	conn, err := dialer.Dial("tcp", "127.0.0.1:3307")
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	// 转换为tcp连接
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		panic("connection error")
	}

	// 连接mysql数据库
	port := strconv.Itoa(tcpConn.LocalAddr().(*net.TCPAddr).Port)
	fmt.Println(tcpConn.RemoteAddr().(*net.TCPAddr).IP)
	fmt.Println(port)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/miata?charset=utf8", "root", "123456", tcpConn.RemoteAddr().(*net.TCPAddr).IP.String(), port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 执行查询语句
	rows, err := db.Query("SHOW TABLES;")
	if err != nil {
		panic(err.Error())
	}

	// 遍历结果集
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			panic(err.Error())
		}
		fmt.Println(id, name)
	}
}
