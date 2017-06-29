package main

import (
	_ "beepro/routers"
	"fmt"
	"log"
	"os"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.Listen.EnableAdmin = true
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix(fmt.Sprintf("PID.%d ", os.Getpid()))
	beego.Run()
}
