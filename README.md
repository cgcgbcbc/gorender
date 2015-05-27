# gorender

[![Build Status](https://travis-ci.org/cgcgbcbc/gorender.svg?branch=master)](https://travis-ci.org/cgcgbcbc/gorender)
[![Coverage Status](https://coveralls.io/repos/cgcgbcbc/gorender/badge.svg?branch=master)](https://coveralls.io/r/cgcgbcbc/gorender?branch=master)

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
