package logger

import (
	"io"
	"log"
	"os"
)

func SetupLogger() {
	f, err := os.OpenFile("social-media-data-mining.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
}
