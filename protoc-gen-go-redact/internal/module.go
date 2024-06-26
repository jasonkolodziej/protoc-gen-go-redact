package internal

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsGo "github.com/lyft/protoc-gen-star/lang/go"
)

// Redactor returns the implementation of the protoc-gen-redact plugin
// to generate redaction file
func Redactor() pgs.Module { return &Module{ModuleBase: &pgs.ModuleBase{}} }

// Module implements the pgs.Module interface for protoc-gen-redact plugin
type Module struct {
	*pgs.ModuleBase
	ctx  pgsGo.Context
	tmpl *template.Template
}

// Name...
func (*Module) Name() string { return "redactor" }

// InitContext satisfies the pgs.Module interface and helps build the module
func (m *Module) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsGo.InitContext(c.Parameters())
	tmpl, err := template.ParseFiles("template.tmpl")
	if err != nil {
		m.tmpl = template.Must(template.New("redactTemplate").Parse(templateConst))
		return
	}
	m.tmpl = template.Must(tmpl, err)
}

// Execute satisfies the pgs.Module interface & generates the redactor file
// for the targeted files
func (m *Module) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
	// process all the target files
	for _, file := range targets {
		m.Process(file)
	}
	return m.Artifacts()
}
