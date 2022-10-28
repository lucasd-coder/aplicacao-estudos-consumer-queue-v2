package logger

import (
	"os"

	"github.com/lucasd-coder/aplicacao-estudos-consumer-queue-v2/config"

	log "github.com/sirupsen/logrus"
)

var Log = log.WithFields(log.Fields{
	"logName":  "aplicacao-estudos-consumer-queue-v2",
	"logIndex": "message",
})

func SetUpLog(cfg *config.Config) {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	logLevel, _ := log.ParseLevel(cfg.Level)
	log.SetLevel(logLevel)
}
