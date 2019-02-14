package common

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type gwLog struct {
	*logrus.Logger
}

var Log *gwLog

func InitLog(logfile string, level uint32) {
	Log = &gwLog{}
	Log.Logger = logrus.New()

	file, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		Log.Out = io.MultiWriter(os.Stdout, file)
	} else {
		Log.Info("Failed to log to file, using default stderr")
	}

	Log.Formatter = &logrus.TextFormatter{
		ForceColors: true,
		DisableTimestamp: false,
		FullTimestamp:true,
		TimestampFormat: time.StampMilli,
	}

	// Only log the warning severity or above.
	Log.SetLevel(logrus.Level(level))
}

func (l *gwLog) Trace(args ...interface{}) {
	args = l.getTraceInfo(args)
	l.Debug(args...)
}

func (l *gwLog) Tracef(format string, args ...interface{}) {
	l.Debug(format)
	l.Debug(args)

	args = l.getTraceInfo(args)

	l.Debug(format)
	l.Debug(args)

	l.Debugf("%s "+format, args...)
}

func (l *gwLog) getTraceInfo(args ...interface{}) []interface{} {
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fileName := filepath.Base(file)

	nameFull := f.Name()
	nameEnd := filepath.Ext(nameFull)
	funcName := strings.TrimPrefix(nameEnd, ".")

	return append([]interface{}{"trace func " + funcName + "@" + fileName + ":" + strconv.Itoa(line-1) + ":"}, args...)
}
