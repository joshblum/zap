package benchmarks

import (
	"io/ioutil"

	"github.com/keybase/go-logging"
)

func newKBLogging() *logging.Logger {
	backend := logging.NewLogBackend(ioutil.Discard, "", 0)
	logging.SetBackend(backend)
	return logging.MustGetLogger("bench")
}
