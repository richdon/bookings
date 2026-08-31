package models

// TemplateData holds data to be injected into a template from handlers
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]any
	CSRFToken string
	Flash     string
	Warning   error
	Error     string
}