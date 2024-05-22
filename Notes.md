How to parse a html file in golang

1. Use the `template.ParseFiles()` function this function reads the template file into a template set. If there is an error, we lof the detailed error message, use the `http.Error()` to send an Internal Server Error response to the user, and then return from the handler so no subsequent code is executed.

2. Then we use the `Execute()` method on the template set to write the template content as the repsonse body. the last parameter to `Excute()` represents any dynamic data that we eat to pass in,
   whic for now we will leave as nil

## Template Composition

As we add more pages to our web application, there will be some shared boilerplate, HTML marku on every page. To prevent duplication and svae typing its is a good idea to create a masteer template that contains the shared content, which we I can compose with page specific markup for the individual pages.

## Update the home handler so that it Parses both template files

1. Intialize a slice that contains the path to both files.

   1. The file containing tha base template must be the first file in the slice'

   go

   `files := []string{
  "./ui/html/base.tmpl.html"
  "./ui/html/pages/homr.tmpl.html
`}

   the `template.ParseFiles()` reads the files and store the templates in the template set.

   Then use the `ExecuteTemplate()` method to write the content of the `base` template as the reponse body

## Creating a File server in Golang

#### To create a file server that serves files out of the `"./ui/static"` directory

the path to the `http.Dir` function is realtive to the project directory root
`fileServer := http.FileServer(http.Dir("./ui/static"))`
