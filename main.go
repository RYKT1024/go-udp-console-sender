package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 检查参数数量
	if len(os.Args) != 4 {
		fmt.Println("Usage: ./main <targetIP> <targetPort> <data>")
		os.Exit(1)
	}

	// 从命令行参数获取目标IP、目标端口和数据内容
	targetIP := os.Args[1]
	targetPort := os.Args[2]
	data := os.Args[3]

	// 构建目标地址
	targetAddr, err := net.ResolveUDPAddr("udp", targetIP+":"+targetPort)
	if err != nil {
		fmt.Println("Error resolving target address:", err)
		os.Exit(1)
	}

	// 创建UDP连接
	conn, err := net.DialUDP("udp", nil, targetAddr)
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 将字符串转换为字节数组，并发送UDP数据包
	byteData := []byte(data)
	_, err = conn.Write(byteData)
	if err != nil {
		fmt.Println("Error sending UDP packet:", err)
		os.Exit(1)
	}

	fmt.Println("UDP packet sent successfully.")
}
