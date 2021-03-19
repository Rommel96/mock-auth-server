package config

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12/middleware/accesslog"
)

type envVars struct {
	portServer string
	dbType     string
	dbUser     string
	dbPass     string
	dbHost     string
	dbPort     string
	dbName     string
}

var config envVars

const newLine = '\n'
const colorRed = "\033[31m"
const colorGreen = "\033[97;32m"
const colorYellow = "\033[33m"
const colorBlue = "\033[97;34m"
const colorMagenta = "\033[97;35m"
const colorCyan = "\033[97;36m"
const colorReset = "\033[0m"

type customFormatter struct {
	w       io.Writer
	bufPool *sync.Pool

	delim byte
	blank string
}

func init() {
	io.WriteString(os.Stdout, colorBlue)
}

func LoadVars() {
	valuesEnv := godotenv.Load()
	if valuesEnv != nil {
		log.Fatal("Error loading environment variables")
	}

	config = envVars{
		portServer: os.Getenv("PORT"),
		dbType:     os.Getenv("DB_TYPE"),
		dbUser:     os.Getenv("DB_USER"),
		dbPass:     os.Getenv("DB_PASS"),
		dbHost:     os.Getenv("DB_HOST"),
		dbPort:     os.Getenv("DB_PORT"),
		dbName:     os.Getenv("DB_NAME"),
	}

}

func GetPortServer() string {
	return ":" + config.portServer
}

func GetConfigDB() (string, string) {
	connString := config.dbUser + ":" + config.dbPass + "@(" + config.dbHost + ":" + config.dbPort + ")/" + config.dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return config.dbType, connString
}

var _ accesslog.Formatter = (*customFormatter)(nil)

func NewCustomFormatter(delim byte, blank string) *customFormatter {
	return &customFormatter{delim: delim, blank: blank}
}

func (f *customFormatter) SetOutput(dest io.Writer) {
	f.w = dest
	f.bufPool = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}

	if f.delim == 0 {
		f.delim = ' '
	}
}

func (f *customFormatter) Format(log *accesslog.Log) (bool, error) {
	buf := f.bufPool.Get().(*bytes.Buffer)
	buf.WriteString(colorReset)
	buf.WriteString(log.Now.Format(log.TimeFormat))
	buf.WriteByte(f.delim)

	buf.WriteString(colorCyan)

	// reqid := log.Fields.GetString("reqid")
	// f.writeTextOrBlank(buf, reqid)
	buf.WriteString(uniformDuration(log.Latency))
	buf.WriteByte(f.delim)

	buf.WriteString(log.IP)
	buf.WriteByte(f.delim)
	if log.Code >= 200 && log.Code < 400 {
		buf.WriteString(colorGreen)
	} else {
		buf.WriteString(colorYellow)
	}
	buf.WriteString(strconv.Itoa(log.Code))
	buf.WriteByte(f.delim)

	buf.WriteString(log.Method)
	buf.WriteByte(f.delim)

	buf.WriteString(log.Path)
	buf.WriteString(colorReset)
	buf.WriteByte(newLine)

	// _, err := buf.WriteTo(f.w)
	// or (to make sure that it resets on errors too):
	_, err := f.w.Write(buf.Bytes())
	buf.Reset()
	f.bufPool.Put(buf)

	return true, err
}

func (f *customFormatter) writeTextOrBlank(buf *bytes.Buffer, s string) {
	if len(s) == 0 {
		if len(f.blank) == 0 {
			return
		} else {
			buf.WriteString(f.blank)
		}

	} else {
		buf.WriteString(s)
	}

	buf.WriteByte(f.delim)
}

func uniformDuration(t time.Duration) string {
	return fmt.Sprintf("%*s", len(t.String()), t.String())
}
