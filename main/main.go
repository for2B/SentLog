package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)


type LogMsg struct {
	SystemName string `json:"systemname"`
	Message string `json:"sysmessage"`
	Time string `json:"systime"`
	Level string `json:"level"`
}

func main(){
	//conn,err:=net.Dial("tcp", "52.81.54.233:30900")
	//if err!=nil{
	//	fmt.Println(err)
	//	panic("conn fail")
	//}

	var filename = "./syslog.log"
	var f *os.File
	var err error
	/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
	if checkFileIsExist(filename) { //如果文件存在
		f, err = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
	} else {
		f, err = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}

	if err!=nil{
		panic("open file fail")
	}


	for ;;{
		LogData := GetLogMsg()
		//_, err := conn.Write(LogData)
		_, err := io.WriteString(f, string(LogData)) //写入文件(字符串)
		if err!=nil{
			panic("write file fail")
		}
		if err!=nil{
			panic(err)
		}
		fmt.Println("success send"+string(LogData))
		time.Sleep(5*time.Second)
	}

}

var level = []string{
	"Info",
	"Debug",
	"Warn",
	"Error",
}

var message = []string{
	"GET /presentations/logstash-monitorama-2013/images/github-contributions.png HTTP/1.1",
	"GET /presentations/logstash-monitorama-2013/images/sad-medic.png HTTP/1.1",
	"GET /presentations/logstash-monitorama-2013/images/kibana-dashboard.png HTTP/1.1",
	"messageError",
}

func GetLogMsg() []byte{
	randName := rand.Intn(4)
	randLevel := rand.Intn(4)
	randMsg := rand.Intn(4)
	lm:=LogMsg{
		SystemName:"IVSYSTEM"+strconv.Itoa(randName),
		Message:message[randMsg],
		Time:time.Now().String(),
		Level:level[randLevel],
	}
	jsondata,err:=json.Marshal(lm)
	if err!=nil{
		panic("marshal fail")
	}
	str := string(jsondata)
	str = str+"\n"
	return []byte(str)
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}