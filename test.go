package main

import "fmt"

// Greeting function types
type Greeting func(name string) string//先把func(name string) string这样的函数声明成Greeting类型

func (g Greeting) say(n string) {
	fmt.Println(g(n))
}

func english(name string) string {
	return "Hello, " + name
}

func main() {
	g := Greeting(english)
	g.say("World")//变量g调用Greeting类型的say()方法
}

