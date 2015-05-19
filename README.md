# gorender

Render go template on the fly

## Install

```
go get github.com/cgcgbcbc/gorender
```

## Usage

Use template from command line

```
gorender --string "{{.Count}}" Count=1
```

Use template file
```
# tmpl.txt
{{.Count}}
```
```
gorender --path tmpl.txt Count=1
```