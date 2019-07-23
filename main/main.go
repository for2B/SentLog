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
	t:=time.Now()
	var filename0 =  fmt.Sprintf("./SYSTEM0-%v-%v-%v.log",t.Year(),t.Month().String(),t.Day())
	var filename1 =  fmt.Sprintf("./SYSTEM1-%v-%v-%v.log",t.Year(),t.Month().String(),t.Day())
	var filename2 =  fmt.Sprintf("./SYSTEM2-%v-%v-%v.log",t.Year(),t.Month().String(),t.Day())
	var filename3 =  fmt.Sprintf("./SYSTEM3-%v-%v-%v.log",t.Year(),t.Month().String(),t.Day())

	var f0,f1,f2,f3 *os.File
	var err error

	if checkFileIsExist(filename0) { //如果文件存在
		f0, err = os.OpenFile(filename0, os.O_APPEND|os.O_RDWR, 0777) //打开文件
	} else {
		f0, err = os.Create(filename0) //创建文件
		fmt.Println("文件不存在")
	}

	if checkFileIsExist(filename1) { //如果文件存在
		f1, err = os.OpenFile(filename1, os.O_APPEND|os.O_RDWR, 0777) //打开文件
	} else {
		f1, err = os.Create(filename1) //创建文件
		fmt.Println("文件不存在")
	}

	if checkFileIsExist(filename2) { //如果文件存在
		f2, err = os.OpenFile(filename2, os.O_APPEND|os.O_RDWR, 0777) //打开文件
	} else {
		f2, err = os.Create(filename2) //创建文件
		fmt.Println("文件不存在")
	}

	if checkFileIsExist(filename3) { //如果文件存在
		f3, err = os.OpenFile(filename3, os.O_APPEND|os.O_RDWR, 0777) //打开文件
	} else {
		f3, err = os.Create(filename3) //创建文件
		fmt.Println("文件不存在")
	}

	if err!=nil{
		panic("open file fail")
	}


	for ;;{
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

		var f *os.File
		switch randName {
		case 0 :
			f = f0
		case 1:
			f = f1
		case 2:
			f = f2
		case 3:
			f = f3
		default:
			f = f0
		}

		_, err = io.WriteString(f, str) //写入文件(字符串)
		if err!=nil{
			panic("write file fail")
		}
		if err!=nil{
			panic(err)
		}
		fmt.Println("success send"+str)
		time.Sleep(2*time.Second)
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