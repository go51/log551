package log551
import (
	"os"
	"fmt"
	"time"
)

type logLevel int

const (
	DEBUG logLevel = iota
	INFORMATION
	WARNING
	ERROR
	CRITICAL
)

func (ll logLevel) String() string {
	switch ll {
	case DEBUG:
		return "DEBUG"
	case INFORMATION:
		return "INFORMATION"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case CRITICAL:
		return "CRITICAL"
	}
	return "NONE"
}

type Config struct {
	Debug       bool `json:"debug"`
	Information bool `json:"information"`
	Warning     bool `json:"warning"`
	Error       bool `json:"error"`
	Critical    bool `json:"critical"`
	Path        string `json:"path"`
}

type log551 struct {
	config *Config
	file   *os.File
}

func New(config *Config) *log551 {
	return &log551{
		config:config,
	}
}

func (l *log551) Open() {
	if l.file != nil {
		return
	}

	file, err := os.OpenFile(l.config.Path, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	l.file = file
}

func (l *log551) Close() {
	err := l.file.Close()
	if err != nil {
		panic(err)
	}

	l.file = nil
}

func (l *log551) Debug(a interface{}) {
	l.output(DEBUG, "%v", a)
}

func (l *log551) Debugf(format string, a ...interface{}) {
	l.output(DEBUG, format, a...)
}

func (l *log551) Debugln(a interface{}) {
	l.output(DEBUG, "%v", a)
}

func (l *log551) Information(a interface{}) {
	l.output(INFORMATION, "%v", a)
}

func (l *log551) Informationf(format string, a ...interface{}) {
	l.output(INFORMATION, format, a...)
}

func (l *log551) Informationln(a interface{}) {
	l.output(INFORMATION, "%v", a)
}

func (l *log551) Warning(a interface{}) {
	l.output(WARNING, "%v", a)
}

func (l *log551) Warningf(format string, a ...interface{}) {
	l.output(WARNING, format, a...)
}

func (l *log551) Warningln(a interface{}) {
	l.output(WARNING, "%v", a)
}

func (l *log551) Error(a interface{}) {
	l.output(ERROR, "%v", a)
}

func (l *log551) Errorf(format string, a ...interface{}) {
	l.output(ERROR, format, a...)
}

func (l *log551) Errorln(a interface{}) {
	l.output(ERROR, "%v", a)
}

func (l *log551) Critical(a interface{}) {
	l.output(CRITICAL, "%v", a)
}

func (l *log551) Criticalf(format string, a ...interface{}) {
	l.output(CRITICAL, format, a...)
}

func (l *log551) Criticalln(a interface{}) {
	l.output(CRITICAL, "%v", a)
}

func (l *log551) output(level logLevel, format string, a ...interface{}) {
	if ! l.isOutput(level) {
		return
	}

	param := []interface{}{
		time.Now().Format("2006/01/02 15:04:05.00000"),
		level.String(),
	}
	param = append(param, a...)

	fmt.Fprintf(l.file, "%s [%s] " + format + "\n", param...)

}

func (l *log551) isOutput(level logLevel) bool {
	if !l.config.Debug && level == DEBUG {
		return false
	}
	if !l.config.Information && level == INFORMATION {
		return false
	}
	if !l.config.Warning && level == WARNING {
		return false
	}
	if !l.config.Error && level == ERROR {
		return false
	}
	if !l.config.Critical && level == CRITICAL {
		return false
	}

	return true
}
