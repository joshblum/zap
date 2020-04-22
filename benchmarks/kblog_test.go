package benchmarks

import (
	"io/ioutil"
	"log"

	"github.com/keybase/client/go/logger"
	"github.com/keybase/go-logging"
)

func newKBLogging() *logging.Logger {
	backend := logging.NewLogBackend(ioutil.Discard, "", 0)
	logging.SetBackend(backend)
	return logging.MustGetLogger("bench")
}

func newKBProdLogging() *logger.InternalLogger {
	logger.EnableBufferedLogging()
	return logger.NewInternalLogger(log.New(ioutil.Discard, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile))
}
