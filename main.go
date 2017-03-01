package main

import (
	"runtime"
	"fmt"
	"os"
	"io/ioutil"
	"github.com/astaxie/beego/httplib"
)


func main() {
	tempResult,_:=httplib.Get("https://raw.githubusercontent.com/racaljk/hosts/master/hosts").String()
	result:=[]byte(tempResult)
	hostsUir:="";
	switch os := runtime.GOOS; os {
	case "windows":
		hostsUir="C:/Windows/System32/drivers/etc/hosts"
	default:
		hostsUir="/etc/hosts"
	}
	osErr:=os.Chmod(hostsUir,777)
	if osErr!=nil{
		fmt.Println("ERROR:",osErr)
		return
	}
	err := ioutil.WriteFile(hostsUir,result, 0777)
	if err != nil {
		fmt.Println("ERROR:",err)
		return
	}
	fmt.Println("SUCCESS!!!!!!!!!!!!!!!!!!!!")
}