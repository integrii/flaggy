package main

import "github.com/integrii/flaggy"

// The custom help message template.
// For rendering text/template will be used: https://godoc.org/text/template
// Object propperties can be looked up here: https://github.com/integrii/flaggy/blob/master/helpValues.go
const helpTemplate = `{{.CommandName}}{{if .Description}} - {{.Description}}{{end}}{{if .PrependMessage}}
{{.PrependMessage}}{{end}}
{{if .UsageString}}
  Usage:
    {{.UsageString}}{{end}}{{if .Positionals}}

  Positional Variables: {{range .Positionals}}
    {{.Name}}  {{.Spacer}}{{if .Description}} {{.Description}}{{end}}{{if .DefaultValue}} (default: {{.DefaultValue}}){{else}}{{if .Required}} (Required){{end}}{{end}}{{end}}{{end}}{{if .Subcommands}}

  Subcommands: {{range .Subcommands}}
    {{.LongName}}{{if .ShortName}} ({{.ShortName}}){{end}}{{if .Position}}{{if gt .Position 1}}  (position {{.Position}}){{end}}{{end}}{{if .Description}}   {{.Spacer}}{{.Description}}{{end}}{{end}}
{{end}}{{if (gt (len .Flags) 0)}}
  Flags: {{if .Flags}}{{range .Flags}}
    {{if .ShortName}}-{{.ShortName}} {{else}}   {{end}}{{if .LongName}}--{{.LongName}}{{end}}{{if .Description}}   {{.Spacer}}{{.Description}}{{if .DefaultValue}} (default: {{.DefaultValue}}){{end}}{{end}}{{end}}{{end}}
{{end}}{{if .AppendMessage}}{{.AppendMessage}}
{{end}}{{if .Message}}
{{.Message}}{{end}}
`

func main() {
	// Declare variables and their defaults
	var stringFlag = "defaultValue"

	// Add a flag
	flaggy.String(&stringFlag, "f", "flag", "A test string flag")

	// Set the help template
	flaggy.DefaultParser.SetHelpTemplate(helpTemplate)

	// Parse the flag
	flaggy.Parse()
}
