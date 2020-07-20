# golang-web-101

## Setting up Sublime Text 3 with Golang support

https://github.com/golang/sublime-build/blob/master/docs/configuration.md 

## Pass variable value to template

Use `{{.}}` in the template, which will be substituded by the pass-in param when executing the template.
For example, `temp.ExecuteTemplate(os.Stdout, "basic_template.gohtml", "Bruce")` will supply value "Bruce" to the occurrence of `{{.}}` in the template

Note, can only pass in one single piece of data into the template via the variable. 

## Parse the templates just once, with pointer variable

**Pointers** in Go programming language or Golang is a variable which is used to store the memory address of another variable. Pointers in Golang is also termed as the special variables. The variables are used to store some data at a particular memory address in the system.

To avoid parsing templates over and over, we can define the `init()` function to parse the whole template folder as a glob, and store in the pointer variable. See the updated `main` function in this repo.

## Skip rendering comment section in template

See https://discourse.gohugo.io/t/how-to-add-comments-in-a-template/75/2 , wrap the comment block with `{{/*   */}}`