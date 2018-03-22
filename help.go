package flaggy

// defaultHelpTemplate is the help template used by default
const defaultHelpTemplate = `{{.CommandName}}{{if .Description}} - {{.Description}}{{end}}
{{if .PrependMessage}}{{.PrependMessage}}
{{end}}{{if .Positionals}}{{end}}
  Positional Variables:
{{range .Positionals}}    {{.Name}} (Position {{.Position}}){{if .Required}} (Required){{end}}{{if .Description}} {{.Description}}{{end}}
{{end}}
{{ if .Subcommands}}  Subcommands:{{end}}
{{range .Subcommands}}    {{.LongName}}{{if .ShortName}} ({{.ShortName}}){{end}}{{if .Position}} (Position {{.Position}}){{end}}{{if .Description}} {{.Description}}{{end}}
{{end}}
{{if (gt (len .StringFlags) 0) | (gt (len .IntFlags) 0) | (gt (len .BoolFlags) 0)}} Flags:
{{range .StringFlags}}    --{{.LongName}}{{if .ShortName}} (-{{.ShortName}}){{end}}{{if .Description}} {{.Description}}{{end}}{{end}}
{{range .IntFlags}}    --{{.LongName}}{{if .ShortName}} (-{{.ShortName}}){{end}}{{if .Description}} {{.Description}}{{end}}{{end}}
{{range .BoolFlags}}    --{{.LongName}}{{if .ShortName}} (-{{.ShortName}}){{end}}{{if .Description}} {{.Description}}{{end}}{{end}}{{end}}
{{if (gt (len .AppendMessage) 0) | (gt (len .Message) 0)}}
{{end}}{{if .AppendMessage}}{{.AppendMessage}}{{end}}{{if .Message}}
{{.Message}}
{{end}}`
