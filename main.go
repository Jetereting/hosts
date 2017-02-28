package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"runtime"
	"fmt"
	"os"
	"io/ioutil"
)

func ExampleScrape() string{
	result:="";
	doc, err := goquery.NewDocument("https://github.com/racaljk/hosts/blob/master/hosts")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("td").Each(func(i int, s *goquery.Selection) {
		result+=s.Text()+"\r\n"
	})
	return result
}

func main() {
	result:=[]byte(ExampleScrape())
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
	}
	err := ioutil.WriteFile(hostsUir,result, 0777)
	if err != nil {
		fmt.Println("ERROR:",err)
	}
}