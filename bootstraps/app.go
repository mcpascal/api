package bootstraps

import (
	"api/configs"
	"api/pkg/cache"
	"api/pkg/logger"
	"api/pkg/mysql"
	"api/pkg/redis"
	"api/pkg/validator"
	"api/routers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"

	_ "net/http/pprof"
)

type application struct {
	Env    string
	Engine *gin.Engine
}

func NewApp() *application {
	undo, err := maxprocs.Set(maxprocs.Logger(log.Printf))
	defer undo()
	if err != nil {
		panic(err)
	}

	err = godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	env := os.Getenv("ENV_NAME")
	if env == "" {
		env = "local"
	}

	// Initialize the Gin engine
	engine := gin.New()
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	app := &application{
		Env:    env,
		Engine: engine,
	}
	return app
}

func (app *application) Start() {
	// 开启 pprof 端点，监听在 localhost:6060
	if app.Env == "development" {
		go func() {
			log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
		}()
	}
	// Setup components
	app.Setup(app.Env)

	// register router
	routers.RegisterRouter(app.Engine)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", configs.App.Host, configs.App.Port),
		Handler: app.Engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Panic("** service start failed", zap.String("error", "启动错误!"))
			return
		}
	}()
	color.Cyan("** service start success **")

	defer app.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}

func (app *application) Stop() {
	logger.Close()
	mysql.Close()
	redis.Close()
}

func (app *application) Setup(env string) {
	configs.Setup(env)
	logger.Setup()
	validator.Setup("zh")
	mysql.Setup()
	redis.Setup()
	cache.Setup()
}
