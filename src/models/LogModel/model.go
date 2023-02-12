package LogModel

import "time"

type LogImpl struct {
	LogId   int
	LogName string
	LogDate time.Time
}

func NewLogImpl(logName string, logDate time.Time) *LogImpl {
	return &LogImpl{LogName: logName, LogDate: logDate}
}

func (*LogImpl) TableName() string {
	return "logs"
}
