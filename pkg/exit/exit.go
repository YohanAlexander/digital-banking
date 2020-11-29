package exit

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/yohanalexander/desafio-banking-go/pkg/logger"
)

// Init callback function invocada para lidar com signals de exit na API
func Init(cb func()) {
	sigs := make(chan os.Signal, 1)
	terminate := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		logger.Error.Fatal("Exit reason: ", sig)
		close(terminate)
	}()

	<-terminate
	cb()
}
