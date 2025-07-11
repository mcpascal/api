package main

import (
	"api/bootstraps"
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	// 设置环境变量
	if os.Getenv("ENV_NAME") == "" {
		os.Setenv("ENV_NAME", "development")
	}

	// 打印启动信息
	fmt.Printf("🚀 Starting Youthank API Server...\n")
	fmt.Printf("📋 Environment: %s\n", os.Getenv("ENV_NAME"))
	fmt.Printf("💻 Go Version: %s\n", runtime.Version())
	fmt.Printf("🏗️  Architecture: %s/%s\n", runtime.GOOS, runtime.GOARCH)

	// 创建并启动应用
	app := bootstraps.NewApp()

	log.Println("✅ Application initialized successfully")
	app.Start()
}
