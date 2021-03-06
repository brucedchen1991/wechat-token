package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tidwall/buntdb"
)

var app = NewApp()

func main() {
	var err error //声明一个error接口变量

	//在 go 标准库中提供了一个包：flag，方便进行命令行解析。
	var (
		version = flag.Bool("version", false, "version v0.1")
		config  = flag.String("config", "account.json", "config file.")
		port    = flag.Int("port", 8000, "listen port.")
	)

	flag.Parse()

	if *version {
		fmt.Println("v0.1")
		os.Exit(0)
	}

	app.SetAccounts(config)
	app.DB, err = buntdb.Open("wechat.db")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer app.DB.Close()

	InitRoute(app.Web.HttpServer)
	log.Println("Start AccessToken Server on ", *port)
	app.Web.StartServer(*port)
}
