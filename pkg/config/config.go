package config

type AppConfig struct {
	InProduction  bool
	Session       *scs.SessionManager
	UseCache      bool
	TemplateCache map[string]*template.Template
}
