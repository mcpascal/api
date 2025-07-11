package main

import (
	"api/bootstraps"
	"fmt"
	"log"
	"os"
	"runtime"
)

// AppOption 定义应用配置选项
type AppOption func(*AppConfig)

// AppConfig 应用配置结构
type AppConfig struct {
	Environment string
	Host        string
	Port        int
	Debug       bool
	EnablePprof bool
	PprofPort   int
}

type Application struct {
}

type Option func(*Application)

func Default() *Application {
	return &Application{
		config: DefaultAppConfig(),
		app:    bootstraps.NewApp(),
	}
}

// DefaultAppConfig 返回默认配置
func DefaultAppConfig() *AppConfig {
	return &AppConfig{
		Environment: "development",
		Host:        "127.0.0.1",
		Port:        8080,
		Debug:       true,
		EnablePprof: true,
		PprofPort:   6060,
	}
}

// WithEnvironment 设置环境
func WithEnvironment(env string) AppOption {
	return func(c *AppConfig) {
		c.Environment = env
	}
}

// WithHost 设置主机地址
func WithHost(host string) AppOption {
	return func(c *AppConfig) {
		c.Host = host
	}
}

// WithPort 设置端口
func WithPort(port int) AppOption {
	return func(c *AppConfig) {
		c.Port = port
	}
}

// WithDebug 设置调试模式
func WithDebug(debug bool) AppOption {
	return func(c *AppConfig) {
		c.Debug = debug
	}
}

// WithPprof 设置性能分析
func WithPprof(enable bool, port int) AppOption {
	return func(c *AppConfig) {
		c.EnablePprof = enable
		c.PprofPort = port
	}
}

// Application 应用结构
type Application struct {
	config *AppConfig
	app    interface{} // 使用interface{}来避免类型问题
}

// NewApplication 创建新的应用实例
func NewApplication(opts ...AppOption) *Application {
	config := DefaultAppConfig()

	// 应用所有选项
	for _, opt := range opts {
		opt(config)
	}

	// 设置环境变量
	os.Setenv("ENV_NAME", config.Environment)

	return &Application{
		config: config,
		app:    bootstraps.NewApp(),
	}
}

// Start 启动应用
func (a *Application) Start() error {
	// 打印启动信息
	a.printStartupInfo()

	// 启动应用
	if app, ok := a.app.(interface{ Start() }); ok {
		app.Start()
	}
	return nil
}

// printStartupInfo 打印启动信息
func (a *Application) printStartupInfo() {
	fmt.Printf("🚀 Starting Youthank API Server...\n")
	fmt.Printf("📋 Environment: %s\n", a.config.Environment)
	fmt.Printf("🌐 Host: %s:%d\n", a.config.Host, a.config.Port)
	fmt.Printf("🐛 Debug Mode: %t\n", a.config.Debug)
	fmt.Printf("📊 Pprof Enabled: %t (port: %d)\n", a.config.EnablePprof, a.config.PprofPort)
	fmt.Printf("💻 Go Version: %s\n", runtime.Version())
	fmt.Printf("🏗️  Architecture: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("📦 Module: %s\n", "api")
}

// 便捷的构建器方法
func (a *Application) SetEnvironment(env string) *Application {
	a.config.Environment = env
	os.Setenv("ENV_NAME", env)
	return a
}

func (a *Application) SetHost(host string) *Application {
	a.config.Host = host
	return a
}

func (a *Application) SetPort(port int) *Application {
	a.config.Port = port
	return a
}

func (a *Application) SetDebug(debug bool) *Application {
	a.config.Debug = debug
	return a
}

func main() {
	// 从环境变量获取配置
	env := os.Getenv("ENV_NAME")
	if env == "" {
		env = "development"
	}

	// 使用Option模式创建应用
	app := NewApplication(
		WithEnvironment(env),
		WithHost("127.0.0.1"),
		WithPort(8080),
		WithDebug(env != "production"),
		WithPprof(env == "development", 6060),
	)

	log.Println("✅ Application initialized successfully")

	// 启动应用
	if err := app.Start(); err != nil {
		log.Fatalf("❌ Failed to start application: %v", err)
	}
}
