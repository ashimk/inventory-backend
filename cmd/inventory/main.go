package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
	"inventory/common"
	"inventory/dm"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	flag.Parse()
	config, err := common.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	url := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DbUser, config.DbPass, config.DbUrl, config.DbName)
	DB, err := gorm.Open(mysql.Open(url))
	DB.Logger.LogMode(logger.Info)
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB, _ := DB.DB()
	defer mysqlDB.Close()

	router := gin.Default()
	router.Use(CORS())
	h := dm.NewHandler(DB)
	router.POST("/create", h.CreateHandle)
	router.GET("/view", h.ViewAllHandler)
	router.POST("/update", h.UpdateHandle)
	router.POST("/search", h.SearchHandle)
	router.DELETE("/delete", h.DeleteHandle)

	tlsConfig, err := getTLSConfig()
	if err != nil {
		fmt.Println("Error loading TLS config:", err)
		return
	}

	server := &http.Server{
		Addr:      config.ServerPort, // HTTPS port
		Handler:   router,
		TLSConfig: tlsConfig,
	}

	fmt.Printf("Starting HTTPS server on port ...", config.ServerPort)
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatalf("Error starting HTTPS server:", err)
	}

	//err = router.Run(":" + config.ServerPort)
	//if err != nil {
	//	log.Fatalf("error in starting server , err:%v", err)
	//}
}

func getTLSConfig() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
	}, nil
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Origin, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
