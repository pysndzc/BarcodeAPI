package main

import (
db "BarcodeApi/database"
"BarcodeApi/initRouter"
"fmt"
"github.com/gin-gonic/gin"
"gopkg.in/ini.v1"
"io"
"os"
)
func main() {
	cfg, err := ini.Load("./conf/config.ini")
	if err != nil{
		fmt.Println("error: %d\n",err)
		return
	}
	f, _ := os.Create("gin_log")
	gin.DefaultWriter = io.MultiWriter(f)
	defer db.SqlDB.Close()
	router := initRouter.SetupRouter()
	_ = router.Run(cfg.Section("Concent").Key("Address").String())
}