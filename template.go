package main

var pkgTemplate = `{{with .PDoc}}{{if $.IsMain}}
# {{ base .ImportPath }}

{{comment_md .Doc}}
{{else}}
# {{ .Name }}
    import "{{.ImportPath}}"

{{comment_md .Doc}}
{{example_html $ ""}}

{{with .Consts}}
## Constants
{{range .}}
<pre>{{node $ .Decl}}</pre>
{{comment_md .Doc}}
{{end}}
{{end}}
{{with .Vars}}
## Variables
{{range .}}
<pre>{{node $ .Decl}}</pre>
{{comment_md .Doc}}
{{end}}
{{end}}
{{range .Funcs}}
{{/* Name is a string - no need for FSet */}}
{{$name_html := html .Name}}
## func {{$name_html}}
<pre>{{node $ .Decl}}</pre>
{{comment_md .Doc}}
{{example_html $ .Name}}
{{end}}
{{range .Types}}
{{$tname := .Name}}
{{$tname_html := html .Name}}
## type {{$tname_html}}
<pre>{{node $ .Decl}}</pre>
{{comment_md .Doc}}

{{range .Consts}}
<pre>{{node $ .Decl}}</pre>
{{comment_md .Doc}}
{{end}}

{{range .Vars}}
<pre>{{node $ .Decl}}</pre>
{{comment_md .Doc}}
{{end}}

{{example_html $ $tname}}

{{range .Funcs}}
{{$name_html := html .Name}}
### func {{$name_html}}

    {{node $ .Decl}}

{{comment_md .Doc}}
{{example_html $ .Name}}
{{end}}

{{range .Methods}}
{{$name_html := html .Name}}
### func ({{md .Recv}}) {{$name_html}}

    {{node $ .Decl}}

{{comment_md .Doc}}
{{$name := printf "%s_%s" $tname .Name}}
{{example_html $ $name}}
{{end}}
{{end}}
{{end}}

{{with $.Notes}}
{{range $marker, $content := .}}
## {{noteTitle $marker | html}}s
<ul style="list-style: none; padding: 0;">
{{range .}}
<li><a href="{{posLink_url $ .}}">&#x261e;</a> {{html .Body}}</li>
{{end}}
</ul>
{{end}}
{{end}}
{{end}}

{{with .PAst}}
<pre>{{node_html $ . false}}</pre>
{{end}}
- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)`
