package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"./eval"
)

func main() {
	http.HandleFunc("/", calcHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type content struct {
	Expr   string
	Result float64
	Err    string
	X, Y   float64
}

func calcHandler(w http.ResponseWriter, req *http.Request) {
	c := content{}
	errs := []error{}
	env := eval.Env{}

	exprStr := req.URL.Query().Get("expression")
	xStr := req.URL.Query().Get("x")
	yStr := req.URL.Query().Get("y")
	c.Expr = exprStr

	if len(xStr) > 0 {
		x, err := strconv.ParseFloat(xStr, 64)
		if err == nil {
			env["x"] = x
			c.X = x
		} else {
			errs = append(errs, err)
		}
	}

	if len(yStr) > 0 {
		y, err := strconv.ParseFloat(yStr, 64)
		if err == nil {
			env["y"] = y
			c.Y = y
		} else {
			errs = append(errs, err)
		}
	}

	if len(exprStr) > 0 {
		expr, err := eval.Parse(exprStr)
		if err == nil {
			c.Result = expr.Eval(env)
		} else {
			errs = append(errs, err)
		}
	}

	buf := new(bytes.Buffer)
	for i, err := range errs {
		if i > 0 {
			buf.WriteString("; ")
		}
		buf.WriteString(err.Error())
	}
	c.Err = buf.String()

	if len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
	}
	if err := htmlTemplate.Execute(w, &c); err != nil {
		log.Fatal(err)
	}
}

var htmlTemplate = template.Must(template.New("htmlTemplate").Parse(`
<div>
{{.Err}}
<form action="">
expression: <input type="text" name="expression" value="{{.Expr}}"/>
<br/>
variables:
<br/>
x = <input type="text" name="x" value="{{.X}}"/>
<br/>
y = <input type="text" name="y" value="{{.Y}}"/>
<br/>
<input type="submit" value="Submit">
<br/>
result: {{.Result}}
</form>
</div>
`))
