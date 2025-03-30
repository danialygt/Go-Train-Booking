package config

import "github.com/alexedwards/scs"

type AppConfig struct {
	InProduction  bool
	Session       *scs.Session
	UseCache      bool
	TemplateCache map[string]*template.Template
}
