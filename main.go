package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	// 开启服务前默认创建ip_block 组
	preScript()

	g := gin.New()
	addRoute(g)
	if err := g.Run("127.0.0.1:9111"); err != nil {
		log.Fatal("set up app fail: ", err)
	}
}

func preScript() {
	var cmd = exec.Command("ipset create ip_block hash:ip")
	cmd.Run()
}

func addRoute(g *gin.Engine)  {
	g.GET("api/block/:name/:ip", func(context *gin.Context) {
		name := context.Param("name")
		ip := context.Param("ip")
		fmt.Print("params name: ", name, " ip: ", ip)
		if len(ip) <= 0 {
			context.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"msg": "ip 是空的",
			})
			return
		}
		timeoutStr := context.DefaultQuery("timeout", "300")
		timeout, err := strconv.Atoi(timeoutStr)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"msg": err.Error(),
			})
		}
		cmdStr := fmt.Sprintf("ipset -exist add %s %s timeout %d", name, ip, timeout)
		cmdList := strings.Split(cmdStr, " ")
		cmd := exec.Command(cmdList[0], cmdList[1:]...)
		if  err = cmd.Run(); err != nil {
			fmt.Println("exec cmd err: ", err)
			context.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"msg": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})
}
