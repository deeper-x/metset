# metset

# Go HTML template testing

### _Rationale_
Golang web application usually are made for template rendering: with this library you keep track of rendered variables inside template (tipically {{.SomeVar}}).
Idea is to define a slice of variables template contains, then calling a checkUp method, template is scanned searching for *all* of them. 


Prototype:
```go
// given a basket of defined vars ({{.var1}}, {{.var2}}, {{.var3}}) 
// we need to search for *all* of them inside template
basketVar := ["var1", "var2", "var3"]

if template.NotContains(basketVar){
    // test fails
}

```



### Clone project

    git clone git@github.com:deeper-x/metset.git

### Usage

    go get github.com/deeper-x/metset

```go
func TestContainsVar(t *testing.T) {
	f := metset.Open("./templates/index.html")
	defer f.Close()

	if !f.Contains("Var") {
		t.Error("{{.Var}} not found, even if is an existent variable")
	}
}
```