package hello

import "github.com/sirupsen/logrus"

func Hello() {
	logrus.Info("Hello, World")
}

func HelloWarn() {
	logrus.Warning("Hello, World")
}

func HelloErr() {
	logrus.Error("Hello, World")
}
