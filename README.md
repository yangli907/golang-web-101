# golang-web-101

## Setting up Sublime Text 3 with Golang support

https://github.com/golang/sublime-build/blob/master/docs/configuration.md 

## Pass variable value to template

Use `{{.}}` in the template, which will be substituded by the pass-in param when executing the template.
For example, `temp.ExecuteTemplate(os.Stdout, "basic_template.gohtml", "Bruce")` will supply value "Bruce" to the occurrence of `{{.}}` in the template

Note, can only pass in one single piece of data into the template via the variable. Further, in order for template to access the struct field or method, the naming of the field or method needs to be **Capitalized**.

## Parse the templates just once, with pointer variable

**Pointers** in Go programming language or Golang is a variable which is used to store the memory address of another variable. Pointers in Golang is also termed as the special variables. The variables are used to store some data at a particular memory address in the system.

To avoid parsing templates over and over, we can define the `init()` function to parse the whole template folder as a glob, and store in the pointer variable. See the updated `main` function in this repo.

## Skip rendering comment section in template

See https://discourse.gohugo.io/t/how-to-add-comments-in-a-template/75/2 , wrap the comment block with `{{/*   */}}`

## Using function in template

To use function in template, **if the method does not have receiver type**, we need to define a function map, and attach it to template memory space **before** parsing the template (otherwise it would've ended up with "variable of function not defined" error, due to the sequencing of operation). The `init()` operation code is as below:

```
temp = template.Must(template.New("")).Funcs(funcMap).ParseGlob("./*.gohtml")
// instead of template.Must(template.ParseGlob("./*.gohtml")).Func(funcMap)
```

However, if the method being used does have a "receiver type" defined, then there's no need to declare funcMap for this method, it is because when we call `ParseFiles()` or `ParseGlob()`, such method is already loaded as part of the type itself. For example,

```
type Player struct {
	Fname string
	Lname string
	Position string
}

func (p Player) IsPF() bool {
	return p.Position == "PF"
}
```

Then we can directly call the `IsPF()` method as below, in the template:

```
		<ul>
			{{range .Person}}
			<li> {{.Lname}}, {{.Fname}} plays {{.Position}} </li>
			<li> {{.IsPF}}
			{{end}}
		</ul>
```