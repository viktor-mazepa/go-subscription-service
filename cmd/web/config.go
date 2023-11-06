package main

import (
	"database/sql"
	"github.com/alexedwards/scs/v2"
	"go-subscription-service/data"
	"log"
	"sync"
)

type Config struct {
	Session       *scs.SessionManager
	DB            *sql.DB
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Wait          *sync.WaitGroup
	Models        data.Models
	Mailer        Mail
	ErrorChan     chan error
	ErrorChanDone chan bool
}

func (app *Config) createMail() Mail {
	errorChan := make(chan error)
	mailerChan := make(chan Message)
	doneChan := make(chan bool)
	return Mail{
		Domain:      "localhost",
		Host:        "localhost",
		Port:        1025,
		Encryption:  "none",
		FromName:    "MazaSoft",
		FromAddress: "info@mazasoft.com",
		WaitGroup:   app.Wait,
		ErrorChan:   errorChan,
		MailerChan:  mailerChan,
		DoneChan:    doneChan,
	}

}
