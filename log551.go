package log551
import (
	"os"
	"fmt"
	"time"
	"runtime"
	"path"
	"regexp"
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
	Debug       ConfigDetail `json:"debug"`
	Information ConfigDetail `json:"information"`
	Warning     ConfigDetail `json:"warning"`
	Error       ConfigDetail `json:"error"`
	Critical    ConfigDetail `json:"critical"`
	Path        string `json:"path"`
}

type ConfigDetail struct {
	EachTime bool `json:"each_time"`
	Cache    bool `json:"cache"`
}

type Log551 struct {
	config *Config
	file   *os.File
	caches []*cache
}

type cache struct {
	datetime    time.Time
	packageName string
	file        string
	line        int
	level       logLevel
	format      string
	a           []interface{}
}

func New(config *Config) *Log551 {
	return &Log551{
		config:config,
		caches:make([]*cache, 0),
	}

}

func (l *Log551) Open() {
	if l.file != nil {
		return
	}

	file, err := os.OpenFile(l.config.Path, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	l.file = file
}

func (l *Log551) Close() {
	l.outputCacheLog()

	err := l.file.Close()
	if err != nil {
		panic(err)
	}

	l.caches = nil
	l.file = nil
}

func (l *Log551) Debug(a interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	packageName := l.getPackageName(runtime.FuncForPC(pc).Name())
	l.output(DEBUG, packageName, file, line, "%v", a)
	l.addCache(DEBUG, packageName, file, line, "%v", a)
}

func (l *Log551) Debugf(format string, a ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	packageName := l.getPackageName(runtime.FuncForPC(pc).Name())
	l.output(DEBUG, packageName, file, line, format, a...)
	l.addCache(DEBUG, packageName, file, line, format, a...)
}

func (l *Log551) Debugln(a interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	packageName := l.getPackageName(runtime.FuncForPC(pc).Name())
	l.output(DEBUG, packageName, file, line, "%v", a)
	l.addCache(DEBUG, packageName, file, line, "%v", a)
}

func (l *Log551) Information(a interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	packageName := l.getPackageName(runtime.FuncForPC(pc).Name())
	l.output(INFORMATION, packageName, file, line, "%v", a)
	l.addCache(INFORMATION, packageName, file, line, "%v", a)
}

func (l *Log551) Informationf(format string, a ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	packageName := l.getPackageName(runtime.FuncForPC(pc).Name())
	l.output(INFORMATION, packageName, file, line, format, a...)
	l.addCache(INFORMATION, packageName, file, line, format, a...)
}

func (l *Log551) Informationln(a interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	packageName := l.getPackageName(runtime.FuncForPC(pc).Name())
	l.output(INFORMATION, packageName, file, line, "%v", a)
	l.addCache(INFORMATION, packageName, file, line, "%v", a)
}

func (l *Log551) Warning(a interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	packageName := l.getPackageName(runtime.FuncForPC(pc).Name())
	l.output(WARNING, packageName, file, line, "%v", a)
	l.addCache(WARNING, packageName, file, line, "%v", a)
}

func (l *Log551) Warningf(format string, a ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	packageName := l.getPackageName(runtime.FuncForPC(pc).Name())
	l.output(WARNING, packageName, file, line, format, a...)
	l.addCache(WARNING, packageName, file, line, format, a...)
}

func (l *Log551) Warningln(a interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	packageName := l.getPackageName(runtime.FuncForPC(pc).Name())
	l.output(WARNING, packageName, file, line, "%v", a)
	l.addCache(WARNING, packageName, file, line, "%v", a)
}

func (l *Log551) Error(a interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	packageName := l.getPackageName(runtime.FuncForPC(pc).Name())
	l.output(ERROR, packageName, file, line, "%v", a)
	l.addCache(ERROR, packageName, file, line, "%v", a)
}

func (l *Log551) Errorf(format string, a ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	packageName := l.getPackageName(runtime.FuncForPC(pc).Name())
	l.output(ERROR, packageName, file, line, format, a...)
	l.addCache(ERROR, packageName, file, line, format, a...)
}

func (l *Log551) Errorln(a interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	packageName := l.getPackageName(runtime.FuncForPC(pc).Name())
	l.output(ERROR, packageName, file, line, "%v", a)
	l.addCache(ERROR, packageName, file, line, "%v", a)
}

func (l *Log551) Critical(a interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	packageName := l.getPackageName(runtime.FuncForPC(pc).Name())
	l.output(CRITICAL, packageName, file, line, "%v", a)
	l.addCache(CRITICAL, packageName, file, line, "%v", a)
}

func (l *Log551) Criticalf(format string, a ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	packageName := l.getPackageName(runtime.FuncForPC(pc).Name())
	l.output(CRITICAL, packageName, file, line, format, a...)
	l.addCache(CRITICAL, packageName, file, line, format, a...)
}

func (l *Log551) Criticalln(a interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	packageName := l.getPackageName(runtime.FuncForPC(pc).Name())
	l.output(CRITICAL, packageName, file, line, "%v", a)
	l.addCache(CRITICAL, packageName, file, line, "%v", a)
}

func (l *Log551) output(level logLevel, packageName, file string, line int, format string, a ...interface{}) {
	if ! l.isOutput(level) {
		return
	}

	param := []interface{}{
		time.Now().Format("2006/01/02 15:04:05.00000"),
		packageName,
		path.Base(file),
		line,
		level.String(),
	}
	param = append(param, a...)

	fmt.Fprintf(l.file, "%s %s/%s %d [%s] " + format + "\n", param...)

}

func (l *Log551) isOutput(level logLevel) bool {
	if !l.config.Debug.EachTime && level == DEBUG {
		return false
	}
	if !l.config.Information.EachTime && level == INFORMATION {
		return false
	}
	if !l.config.Warning.EachTime && level == WARNING {
		return false
	}
	if !l.config.Error.EachTime && level == ERROR {
		return false
	}
	if !l.config.Critical.EachTime && level == CRITICAL {
		return false
	}

	return true
}

func (l *Log551) addCache(level logLevel, packageName, file string, line int, format string, a ...interface{}) {
	if ! l.isCache(level) {
		return
	}

	cache := &cache{
		datetime: time.Now(),
		packageName:packageName,
		file: file,
		line: line,
		level: level,
		format: format,
		a: a,
	}

	l.caches = append(l.caches, cache)

}

func (l *Log551) isCache(level logLevel) bool {
	if !l.config.Debug.Cache && level == DEBUG {
		return false
	}
	if !l.config.Information.Cache && level == INFORMATION {
		return false
	}
	if !l.config.Warning.Cache && level == WARNING {
		return false
	}
	if !l.config.Error.Cache && level == ERROR {
		return false
	}
	if !l.config.Critical.Cache && level == CRITICAL {
		return false
	}

	return true
}

func (l *Log551) outputCacheLog() {
	if len(l.caches) == 0 {
		return
	}

	for _, v := range l.caches {
		param := []interface{}{
			v.datetime.Format("2006/01/02 15:04:05.00000"),
			v.packageName,
			path.Base(v.file),
			v.line,
			v.level.String(),
		}
		param = append(param, v.a...)
		fmt.Fprintf(l.file, "%s %s/%s %d [%s] " + v.format + "\n", param...)
	}
}

var regexpControllerPattern *regexp.Regexp

func (r *Log551) getPackageName(path string) string {
	if regexpControllerPattern == nil {
		regexpControllerPattern = regexp.MustCompile(`\/\w+\.`)
	}

	packageName := regexpControllerPattern.FindString(path)

	return packageName[1 : len(packageName)-1]
}
