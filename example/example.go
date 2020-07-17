package main

import (
	"fmt"
	"time"

	"github.com/mafangtao/log"
	"github.com/mafangtao/log/lager"
)

func main() {
	log.InitWithFile("log.yaml")

	for i := 0; i < 100; i++ {
		log.Infof("Hi %s, system is starting up ...", "paas-bot")
		log.Info("check-info", lager.Data{
			"info": "something",
		})

		log.Debug("check-info", lager.Data{
			"info": "something",
		})

		log.Warn("failed-to-do-somthing", lager.Data{
			"info": "something",
		})

		err := fmt.Errorf("This is an error")
		log.Error("failed-to-do-somthing", err)

		log.Info("shutting-down")

		time.Sleep(10*time.Millisecond)
	}
}
