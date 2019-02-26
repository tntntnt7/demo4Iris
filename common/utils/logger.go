package utils

import (
	"log"
	"os"
)

type Log struct {
	L *log.Logger
}

var Logger *Log

func InitLogger() {
	file, err := initLogFile()
	if err != nil {
		log.Fatalln("fail to create demo4Iris.log file!")
	}
	Logger = new(Log)
	Logger.L = log.New(file, "", log.LstdFlags | log.Llongfile)
	Logger.L.SetFlags(log.LstdFlags | log.Lshortfile)
}

func initLogFile() (*os.File, error) {
	workPath, _ := os.Getwd()
	logPath := workPath + "/log"

	_, err := os.Stat(logPath)
	if err != nil {
		os.Mkdir(logPath, os.ModePerm)
	}
	return os.OpenFile("./log/demo4Iris.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
}

func (l *Log) Info(s string) {
	l.L.SetPrefix("INFO    ")
	l.L.Println(s)
}

func (l Log) Error(s string) {
	l.L.SetPrefix("ERROR    ")
	l.L.Println(s)
}

func (l Log) Fatal(s string) {
	l.L.SetPrefix("FATAL    ")
	l.L.Fatal(s)
}

func (l Log) Panic(s string) {
	l.L.SetPrefix("PANIC    ")
	l.L.Panic(s)
}
