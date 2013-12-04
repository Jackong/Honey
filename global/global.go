/**
 * User: jackong
 * Date: 11/5/13
 * Time: 4:50 PM
 */
package global

import (
	"os"
	"fmt"
	"time"
	"github.com/Jackong/db"
	_ "github.com/Jackong/db/mysql"
	"github.com/Jackong/Honey/config"
	"github.com/Jackong/log"
	"github.com/Jackong/log/writer"
)

var (
	GoPath string
	Now func() string
	Time func() time.Time
	Today func() string
	Project config.Config
	DB db.Database
	Log log.Logger
)

func init() {
	fmt.Println("init env...")
	baseEnv()

	fmt.Println("loading config...")
	loadConfig()

	fmt.Println("opening db...")
	openDb()

	fmt.Println("init log...")
	initLog()
}

func initLog() {
	today := Today()
	fmt.Println("getting mail log...")
	mailLog := newDateLog(today, mailLog)

	fileLevel := int(Project.Get("log", "file", "level").(float64))

	fmt.Println("getting action log...")
	actionLog := newDateLog(today, func(date string) log.Logger {
			return fileLog("action", date, fileLevel)
		})

	Log = log.MultiLogger(actionLog, mailLog)
}

func fileLog(dir, date string, level int) log.Logger {
	dir = GoPath + "/" + Project.String("log", "dir") + "/" +  dir
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}
	file, err := os.OpenFile(dir + "/" + date + ".log", os.O_RDWR | os.O_CREATE | os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		ShutDown()
	}
	logger := log.NewLogger(file, Project.String("server", "name"), level)
	return logger
}

func mailLog(date string) log.Logger {
	mail := &writer.Email{
		User: Project.String("log", "email", "user"),
		Password: Project.String("log", "email", "password"),
		Host : Project.String("log", "email", "host"),
		To : Project.String("log", "email", "to"),
		Subject: Project.String("log", "email", "subject"),
	}

	server := Project.String("server", "name")
	mailLevel := int(Project.Get("log", "email", "level").(float64))
	check := Project.Get("formal").(bool)
	if check {
		if err := mail.SendMail("starting server " + server + "..."); err != nil {
			fmt.Println(err)
			ShutDown()
		}
		return log.NewLogger(&asyncMail{mail}, server, mailLevel)
	}
	return fileLog("email", date, mailLevel)
}

func baseEnv() {
	Time = func() time.Time {
		return time.Now()
	}

	Now = func() string {
		return Time().Format(FORMAT_DATE_TIME)
	}

	Today = func() string {
		return Time().Format(FORMAT_DATE)
	}

	GoPath = os.Getenv("GOPATH")
}

func loadConfig() {
	Project = config.NewConfig(GoPath  + "/src/github.com/Jackong/Honey/config/project.json")
}

func openDb() {
	settings := db.DataSource{
		Host:	Project.String("db", "host"),
		Port:	int(Project.Get("db", "port").(float64)),
		Database: Project.String("db", "database"),
		User:     Project.String("db", "user"),
		Password: Project.String("db", "password"),
	}

	var err error
	if DB, err = db.Open("mysql", settings); err != nil {
		fmt.Println(err)
		ShutDown()
	}

	OnShutDown(func() {
		DB.Close()
	})
}
