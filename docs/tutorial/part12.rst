uAdmin Tutorial Part 12 - Storing the data to HTML
==================================================
In this part, we will discuss about fetching the records in the API and migrating the data from API to HTML that will display the records using Go template.

Go to todo_handler.go inside the handlers folder with the following codes below:

.. code-block:: go

    package handlers

    import (
        "html/template"
        "net/http"
        "strings"

        // Specify the username that you used inside github.com folder and
        // import this library
        "github.com/username/todo/models"
        "github.com/uadmin/uadmin"
    )

    // TodoHandler !
    func TodoHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/todo")

        type Context struct {
            TodoList []map[string]interface{}
        }
        c := Context{}

        // ------------------ ADD THIS CODE ------------------
        // Fetch Data from DB
        todo := []models.Todo{}
        uadmin.All(&todo)

        for i := range todo {
            // Accesses and fetches the record of the linking models in Todo
            uadmin.Preload(&todo[i])

            // Assigns the string of interface in each Todo fields
            c.TodoList = append(c.TodoList, map[string]interface{}{
                "ID":   todo[i].ID,
                "Name": todo[i].Name,
                // In fact that description has an html type tag in uAdmin,
                // we have to convert this field from text to HTML so that
                // the HTML tags from models will be applied to HTML file.
                "Description": template.HTML(todo[i].Description),
                "Category":    todo[i].Category.Name,
                "Friend":      todo[i].Friend.Name,
                "Item":        todo[i].Item.Name,
                "TargetDate":  todo[i].TargetDate,
                "Progress":    todo[i].Progress,
            })
        }
        // ----------------------------------------------------

        // Some codes

    }

Now go to views/todo.html. After the <tbody> tag, add the following codes shown below:

.. code-block:: html

    {{range .TodoList}}
    <tr>
        <td>{{.Name}}</td>
        <td>{{.Description}}</td>
        <td>{{.Category}}</td>
        <td>{{.Friend}}</td>
        <td>{{.Item}}</th>
        <td>{{.TargetDate}}</td>
        <td>{{.Progress}}</td>
    </tr>
    {{end}}

In Go programming language, **range** is equivalent to **for** loop.

The double brackets **{{ }}** are Golang delimiter.

**.TodoList** is the assigned field inside the Context struct.

**.Name**, **Description**, **.Category**, **.Friend**, **.Item**, **.TargetDate**, **.Progress** are the fields assigned in c.TodoList.

Now run your application, go to http_handler/todo path and see what happens.

.. image:: assets/todohtmlresult.png

|

Congrats, now you know how to set up a handler file in an organized manner, access the HTML in localhost and store the data from API to HTML using Go templates.

In the `next part`_, we will talk about generating a self-signed SSL certificate using the **openssl** command and implementing two factor authentication (2FA).

.. _next part: https://uadmin.readthedocs.io/en/latest/tutorial/part13.html

