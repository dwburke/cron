package cron

import (
	"github.com/dwburke/atexit"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Cron *cron.Cron

func init() {
	viper.SetDefault("cron.enabled", true)
}

func Run() {
	if viper.GetBool("cron.enabled") == false {
		log.Info("cron.enabled == false; not starting")
		return
	}

	log.Info("cron.enabled == true; starting")

	Cron = cron.New()

	Cron.Start()

	atexit.Add(Shutdown)
}

func Shutdown() {
	if Cron != nil {
		log.Info("cron: shutting down")
		Cron.Stop()
	}
}
