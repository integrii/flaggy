package flaggy

// defaultHelpTemplate is the help template used by default
// {{if (or (or (gt (len .StringFlags) 0) (gt (len .IntFlags) 0)) (gt (len .BoolFlags) 0))}}
// {{if (or (gt (len .StringFlags) 0) (gt (len .BoolFlags) 0))}}
const defaultHelpTemplate = `{{.CommandName}}{{if .Description}} - {{.Description}}{{end}}{{if .PrependMessage}}
{{.PrependMessage}}{{end}}{{if .Positionals}}

  Positional Variables:
{{range .Positionals}}
    {{.Name}} (Position {{.Position}}){{if .Required}} (Required){{end}} -{{if .Description}} {{.Description}}{{end}}{{end}}{{end}}{{if .Subcommands}}

  Subcommands:
{{range .Subcommands}}
    {{.LongName}}{{if .ShortName}} ({{.ShortName}}){{end}} -{{if .Position}}{{if gt .Position 1}} (Position {{.Position}}){{end}}{{end}}{{if .Description}} {{.Description}}{{end}}{{end}}{{end}}{{if (or (or (gt (len .StringFlags) 0) (gt (len .IntFlags) 0)) (gt (len .BoolFlags) 0))}}

  Flags:
  {{if .StringFlags}}{{range .StringFlags}}
    {{if .LongName}}--{{.LongName}} {{end}}{{if .ShortName}}(-{{.ShortName}}){{end}}{{if .Description}} {{.Description}}{{end}}{{end}}{{end}}{{if .IntFlags}}{{range .IntFlags}}
    {{if .LongName}}--{{.LongName}} {{end}}{{if .ShortName}}(-{{.ShortName}}){{end}}{{if .Description}} {{.Description}}{{end}}{{end}}{{end}}{{if .BoolFlags}}{{range .BoolFlags}}
    {{if .LongName}}--{{.LongName}} {{end}}{{if .ShortName}}(-{{.ShortName}}){{end}}{{if .Description}} {{.Description}}{{end}}{{end}}{{end}}
{{end}}
{{if .AppendMessage}}
{{.AppendMessage}}{{end}}{{if .Message}}
{{.Message}}
{{end}}`
