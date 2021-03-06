package config

import (
	"html/template"
	"log"

	"github.com/AdiF1/solidity/bookings/internal/models"
	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config variables

type AppConfig struct {
	UseCache 		bool
	TemplateCache 	map[string]*template.Template
	InfoLog  		*log.Logger
	ErrorLog		*log.Logger
	InProduction 	bool
	Session 		*scs.SessionManager
	MailChan		chan models.MailData

}


// --------------------------------------------

