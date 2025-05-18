package cmd

// import (
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"user_service/internal/app"
// 	"user_service/internal/config"
// 	"user_service/log"
// 	"fmt"
// )



// func main() {
// 	logger,file,err := logger.NewLogger()
// 	if err != nil {
// 		panic(err)
// 	}
// 	cfg,err := config.NewConfig()
// 	if err != nil {
// 		logger.Console.Err(err)
// 	}

// 	application := app.NewApp(cfg,logger)
// 	go application.GRPCServer.Run()
// 	logger.Console.Warn().Msg("to stop the server, press CTRL + C")
// 	stop  :=  make(chan os.Signal,1)
// 	signal.Notify(stop,syscall.SIGTERM,syscall.SIGINT)
// 	sig  := <-stop

// 	logger.Console.Info().Msg(fmt.Sprintf("received shutdown signal %s",sig.String()))
// 	application.GRPCServer.Stop()
// 	logger.Console.Info().Msg("shuttin down server")
// 	defer file.Close()
// }



