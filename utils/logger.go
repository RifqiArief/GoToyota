package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var Logging *log.Logger

func Logger() error {
	err := godotenv.Load("config/log.env")
	if err != nil {
		log.Println("utils/logger/line:19")
		return err
	}

	t := time.Now()
	logName := fmt.Sprintf("%s/%s.log", os.Getenv("path"), string(t.Format("02-Jan-2006")))

	err = os.MkdirAll(os.Getenv("path"), os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	Logging = log.New(file, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	mw := io.MultiWriter(os.Stdout, file)
	Logging.SetOutput(mw)

	return nil
}
