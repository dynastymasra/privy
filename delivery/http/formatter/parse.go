package formatter

import (
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	DefaultFrom = 0
	DefaultSize = 20
)

func BuildPaginationFrom(str string) int {
	if len(str) < 1 {
		return DefaultFrom
	}

	parse, err := strconv.Atoi(str)
	if err != nil {
		logrus.WithField("value", str).WithError(err).Errorln("Failed parse string to int")

		return DefaultFrom
	}
	return parse
}

func BuildPaginationSize(str string) int {
	if len(str) < 1 {
		return DefaultSize
	}

	parse, err := strconv.Atoi(str)
	if err != nil {
		logrus.WithField("value", str).WithError(err).Errorln("Failed parse string to int")

		return DefaultSize
	}
	return parse
}
