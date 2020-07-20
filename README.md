# golang-web-101

## Setting up Sublime Text 3 with Golang support

https://github.com/golang/sublime-build/blob/master/docs/configuration.md 

## Pass variable value to template

Use `{{.}}` in the template, which will be substituded by the pass-in param when executing the template.
For example, `temp.ExecuteTemplate(os.Stdout, "basic_template.gohtml", "Bruce")` will supply value "Bruce" to the occurrence of `{{.}}` in the template

Note, can only pass in one single piece of data into the template via the variable. 