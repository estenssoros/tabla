package main

import (
	"github.com/estenssoros/tabla/cmd"
	"github.com/sirupsen/logrus"
)

func init() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logrus.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true
}

func main() {
	if err := cmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
