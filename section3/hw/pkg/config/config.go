package config

import "text/template"

//AppConfig holds the application configuation
type AppConfig struct {
	TemplateCache map[string]*template.Template
}