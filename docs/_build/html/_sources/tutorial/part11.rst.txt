uAdmin Tutorial Part 11 - Accessing an HTML file
================================================
In this part, we will talk about establishing a connection to the template, setting the URL path name, and executing an HTML file.

Go to template.go inside the templates/custom path with the following codes below:

.. code-block:: go

    package templates

    import (
        "net/http"
        "strings"
    )

    // TemplateHandler !
    func TemplateHandler(w http.ResponseWriter, r *http.Request) {
        // r.URL.Path creates a new path called /template
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/template")
    }

Establish a connection in the main.go to the template by using http.HandleFunc. It should be placed after the uadmin.Register and before the StartServer.

.. code-block:: go

    import (
        "net/http"

        // Specify the username that you used inside github.com folder
        "github.com/username/todo/api"
        "github.com/username/todo/models"

        // Import this library
        "github.com/username/todo/templates/custom"

        "github.com/uadmin/uadmin"
    )

    func main() {
        // Some codes

        // Template Handler
        http.HandleFunc("/template/", templates.TemplateHandler)
    }

Create a file named todo_template.go inside the templates/custom path with the following codes below:

.. code-block:: go

    package templates

    import (
        "html/template"
        "net/http"
        "strings"

        "github.com/uadmin/uadmin"
    )

    // TodoTemplateHandler !
    func TodoTemplateHandler(w http.ResponseWriter, r *http.Request) {
        // r.URL.Path creates a new path called /todo_html
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/todo_html")

        // Creates a new template and parses the template definitions
        // from todo.html
        t, err := template.New("").ParseFiles("./views/todo.html")
        if err != nil {
            uadmin.Trail(uadmin.ERROR, "Cannot open todo list", err.Error())
        }

        // Applies the template associated with t that has todo.html to
        // the specified object and writes the output to w variable that
        // is the http.ResponseWriter
        err = t.ExecuteTemplate(w, "todo.html", nil)
    }

Finally, add this piece of code in the template.go shown below. This will establish a communication between the TodoTemplateHandler and the TemplateHandler.

.. code-block:: go

    // TemplateHandler !
    func TemplateHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/template")

        // ------------------ ADD THIS CODE ------------------
        if strings.HasPrefix(r.URL.Path, "/todo_html") {
            TodoTemplateHandler(w, r)
            return
        }
        // ------------------ ADD THIS CODE ------------------
    }

Now run your application, go to template/todo_html path and see what happens.

.. image:: assets/todohtmlaccess.png

|

In the `next part`_, we will discuss about fetching the records in the API and migrating the data from API to HTML that will display the records using Go template.

.. _next part: https://uadmin.readthedocs.io/en/latest/tutorial/part12.html
