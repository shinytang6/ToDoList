package main

import (
	_ "quickstart/routers"
	"github.com/astaxie/beego"
  
    
)


func main() {
   
    beego.BConfig.Listen.HTTPPort = 8000
	beego.Run()
}

