package logger

import (
    "fmt"
    "log"
    "os"
    "sync"
    "strconv"
)

// internal type for our enums
type levels int

// emulate enum
const (
	FATAL levels = -1
	ERROR levels = 0
	WARNING levels = 1
	INFO levels = 2
	DEBUG levels = 3
	TRACE levels = 4
	UNSUPP levels = 99
)

// operator overload - cast enum to string
func (l levels) String() string {
	switch l {
	case FATAL:
		return "FATAL"
	case ERROR:
		return "ERROR"
	case WARNING:
		return "WARNING"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	default:
		return "UNSUPP"
	}
}

// interface to wrap our enum - it cannot be inherited / subclassed
type Level interface {
	Levels() levels
}

// operator overload - access the underlying enum value
func(l levels) Levels() levels {
	return l
}

// our custome logger structure
type local_logger struct {
    filename string
    *log.Logger
}

// instance singletons
var the_logger *local_logger
var once sync.Once

// start logging via singleton
func GetInstance() *local_logger {
    once.Do(func() {
        the_logger = createLogger("mylogger.log")
    })
    return the_logger
}

// internal function to access logger (eck)
func createLogger(fname string) *local_logger {
    file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

    return &local_logger{
        filename: fname,
				Logger:   log.New(file, "fedramp-templater: ", log.Lshortfile),
    }
}

// log and print with given stack and format
func internalPrint(stackLevels int, desired Level, format string, args...interface{}) bool {
	if( Check( desired ) ) {
		var formatted_str = fmt.Sprintf(format, args...)
		GetInstance().Output(stackLevels, fmt.Sprintf("%s %s\n", desired.Levels(), formatted_str))
		return true
	}
	return false
}

// should logging occur?
func Check(desired Level) bool {
	var env_debug_str = os.Getenv("DEBUG")
	env_debug_val, err := strconv.Atoi(env_debug_str)
	if err != nil { return false }
	var actual = UNSUPP
	switch env_debug_val {
		case int(FATAL):
			actual = FATAL
		case int(ERROR):
			actual = ERROR
		case int(WARNING):
			actual = WARNING
		case int(INFO):
			actual = INFO
		case int(DEBUG):
			actual = DEBUG
		case int(TRACE):
			actual = TRACE
		default:
			actual = UNSUPP
	}
	if( actual.Levels() >= UNSUPP ) { return false }
	return ( actual.Levels() >= desired.Levels() );
}

// log and print
func Print(desired Level, msg string) bool {
	return internalPrint(2, desired, "%s", msg)
}

// log and print
func Trace(msg string) bool {
	return internalPrint(3, TRACE, "%s", msg)
}

// log and print
func Debug(msg string) bool {
	return internalPrint(3, DEBUG, "%s", msg)
}

// log and print
func Info(msg string) bool {
	return internalPrint(3, DEBUG, "%s", msg)
}

// log and print
func Warning(msg string) bool {
	return internalPrint(3, WARNING, "%s", msg)
}

// log and print
func Error(msg string) bool {
	return internalPrint(3, ERROR, "%s", msg)
}

// log and print
func Fatal(msg string) bool {
	return internalPrint(3, FATAL, "%s", msg)
}

// log and print
func Printf(desired Level, format string, args ...interface{}) bool {
	return internalPrint(2, desired, format, args...)
}

// log and print
func Tracef(format string, args ...interface{}) bool {
	return internalPrint(3, TRACE, format, args...)
}

// log and print
func Debugf(format string, args ...interface{}) bool {
	return internalPrint(3, DEBUG, format, args...)
}

// log and print
func Infof(format string, args ...interface{}) bool {
	return internalPrint(3, DEBUG, format, args...)
}

// log and print
func Warningf(format string, args ...interface{}) bool {
	return internalPrint(3, WARNING, format, args...)
}

// log and print
func Errorf(format string, args ...interface{}) bool {
	return internalPrint(3, ERROR, format, args...)
}

// log and print
func Fatalf(format string, args ...interface{}) bool {
	return internalPrint(3, FATAL, format, args...)
}

