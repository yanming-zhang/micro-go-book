package main

import (
	"fmt"
	"log"
	"micro-go-book/ch8-config/conf"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	http.HandleFunc("/resumes", func(w http.ResponseWriter, req *http.Request) {
		//q := events.goreq.URL.Query().Get("q")
		_, _ = fmt.Fprintf(w, "个人信息：\n")
		_, _ = fmt.Fprintf(w, "姓名：%s，\n性别：%s，\n年龄 %d!", viper.GetString("resume.name"), conf.Resume.Sex, conf.Resume.Age) //这个写入到w的是输出到客户端的
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
