// Author: Turing Zhu
// Date: 2020/8/10 1:01 PM
// File: api.go

package ginSwagger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwag "github.com/swaggo/gin-swagger"
)

// @title Huobi API
// @version 0.1
// @description Huobi API
// @contact.name Turing Zhu
// @contact.email qishiwenjun@163.com
// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT
// @host api.testnet.huobi.pro
// @BasePath /api
func Run() {
	engine := gin.New()
	listening := fmt.Sprintf("%s:%d", "127.0.0.1", 8808)

	url := ginSwag.URL(fmt.Sprintf("http://%s/swagger/doc.json", listening))
	engine.GET("/swagger/*any", ginSwag.WrapHandler(swaggerFiles.Handler, url))
	err := engine.Run(listening)
	if err != nil {
		fmt.Println(err)
	}
}
