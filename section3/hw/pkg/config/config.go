package config

import "text/template"

// AppConfig holds the application configuation
type AppConfig struct {
	UseTemplateCache bool
	TemplateCache    map[string]*template.Template
}
