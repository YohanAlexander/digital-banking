package exit

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

// Init lida com os signals de exit na API
func Init(cb func()) {
	sigs := make(chan os.Signal, 1)
	terminate := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		logrus.Fatal("Exit reason: ", sig)
		terminate <- true
	}()

	<-terminate
	cb()
	logrus.Panic("Exiting banking server")
}
