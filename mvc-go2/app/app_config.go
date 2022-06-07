package app

import (
	"os"

	log "github.com/sirupsen/logrus" //libreria de un tercero
)

func init() { // funcion init por defecto siempre está
	log.SetOutput(os.Stdout)
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.Info("Starting logger system")
}
