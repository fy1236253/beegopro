package main

import (
	_ "beegopro/routers"
	"fmt"
	"log"
	"os"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.Listen.EnableAdmin = true
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix(fmt.Sprintf("PID.%d ", os.Getpid()))
	f, _ := os.OpenFile("logs/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	log.SetOutput(f)
	beego.SetLogger("file", `{"filename":"logs/log.log"}`)
	beego.Run()
}
