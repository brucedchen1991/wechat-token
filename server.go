package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/devfeel/dotweb"
	"github.com/tidwall/buntdb"
)


//Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。
/* 声明变量，默认 map 是 nil */
//var map_variable map[key_data_type]value_data_type

/* 使用 make 函数 */
//map_variable := make(map[key_data_type]value_data_type)
//声明一个结构体
//结构体包含 Accounts DB Web WxToken
type App struct {
	Accounts map[string]string //map类型 key是string value是string
	DB       *buntdb.DB
	Web      *dotweb.DotWeb
	WxToken  *Token
}
//声明一个Account结构体
type Account struct {
	AppID  string `json:"appid"`//AppID是string类型的
	Secret string `json:"secret"`
}
//函数返回App结构体
func NewApp() *App {
	var a = &App{}
	a.Accounts = make(map[string]string)
	a.Web = dotweb.New()
	a.WxToken = new(Token)

	return a
}

// 读取配置文件中的appid和secret值到一个map中
func (a *App) SetAccounts(config *string) {
	accounts := make([]Account, 1)

	if _, err := os.Stat(*config); err != nil {
		fmt.Println("配置文件无法打开！")
		os.Exit(1)
	}

	raw, err := ioutil.ReadFile(*config)
	if err != nil {
		fmt.Println("配置文件读取失败！")
		os.Exit(1)
	}

	if err := json.Unmarshal(raw, &accounts); err != nil {
		fmt.Println("配置文件内容错误！")
		os.Exit(1)
	}

	for _, acc := range accounts {
		a.Accounts[acc.AppID] = acc.Secret
	}
}

func (a *App) Query(appid string, key string) string {
	var value string

	err := a.DB.View(func(tx *buntdb.Tx) error {
		v, err := tx.Get(appid + "_" + key)
		if err != nil {
			return err
		}
		value = v
		return nil
	})
	if err != nil {
		value = ""
	}

	return value
}

// 更新AppID上下文环境中的Access Token和到期时间
func (a *App) UpdateToken(appid string) {
	timestamp := time.Now().Unix()

	a.DB.Update(func(tx *buntdb.Tx) error {
		tx.Delete(appid + "_timestamp")
		tx.Delete(appid + "_access_token")
		tx.Delete(appid + "_expires_in")

		tx.Set(appid+"_timestamp", strconv.FormatInt(timestamp, 10), nil)
		tx.Set(appid+"_access_token", a.WxToken.AccessToken, nil)
		tx.Set(appid+"_expires_in", strconv.Itoa(a.WxToken.Expire), nil)
		return nil
	})
}
