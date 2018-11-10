uAdmin Tutorial Part 7 - Introduction to API
============================================
In this part, we will discuss about establishing a connection to the API, setting the path name, and getting the todo list data in the API Handler using JSON.

Create a file named api.go inside the api folder with the following codes below:

.. code-block:: go

    package api

    import (
        "fmt"
        "net/http"
        "strings"
    )

    // This part of code is the API HELP to be printed out in the body of the
    // web page.
    const API_HELP = `TODO API HELP
    For more assistance please contact Integritynet:
    support@integritynet.biz

    - todo:
        # method     : todo_list
        # Parameters:  
        # Return    : json object that returns the list of your todo activities
    `

    // APIHandler !
    func APIHandler(w http.ResponseWriter, r *http.Request) {
        // r.URL.Path creates a new path called /api
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")

        // If there is no subsequent method, it will call the API_HELP
        // variable to display the message.
        if r.URL.Path == "/" {
            fmt.Fprintf(w, API_HELP)
        }
    }

As shown above, we have to call the variable named "API_HELP" to inform the user what are the methods to visit in the api path. To make the API function, we create a handler named "APIHandler" that handles two parameters which are **http.ResponseWriter** that assembles the HTTP server's response; by writing to it, we send data to the HTTP client; and **http.Request** which is a data structure that represents the client HTTP request. **r.URL.Path** is the path component of the request URL. In this case, we call /api. If there is no subsequent method, it will call the API_HELP variable to display the message.

Go back to the main.go and apply **uadmin.RootURL** as "/admin/" to make the /api functional. Put it above the uadmin.Register.

.. code-block:: go

    func main() {
        uadmin.RootURL = "/admin/" // <-- place it here
        uadmin.Register(
            // Some codes
        )
    }

Establish a connection in the main.go to the API by using http.HandleFunc. It should be placed after the uadmin.Register and before the StartServer.

.. code-block:: go

    func main() {
        // Some codes

        // API Handler
        http.HandleFunc("/api/", api.APIHandler) // <-- place it here
    }

api is the folder name while APIHandler is the name of the function inside api.go.

Go the todo.go inside the models folder. Create a Preload() function to call the ID of other models to fetch the first record from the database. Put it under the Todo struct.

.. code-block:: go

    // Todo model ...
    type Todo struct {
        // Some codes
    }

    // Preload ...
    func (t *Todo) Preload() {
        if t.Category.ID != t.CategoryID {
            uadmin.Get(&t.Category, "id = ?", t.CategoryID)
        }
        if t.Friend.ID != t.FriendID {
            uadmin.Get(&t.Friend, "id = ?", t.FriendID)
        }
        if t.Item.ID != t.ItemID {
            uadmin.Get(&t.Item, "id = ?", t.ItemID)
        }
    }

Now let's create another file inside the api folder named todo_list.go. This will return the list of your todo activities in JSON format.

.. code-block:: go

    package api

    import (
        "net/http"
        "strings"

        "github.com/username/todo/models"
        "github.com/uadmin/uadmin"
    )

    // TodoListHandler !
    func TodoListHandler(w http.ResponseWriter, r *http.Request) {
        // r.URL.Path creates a new path called /todo_list
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/todo_list")

        // Initializes res as a map[string]interface{}{} where you can put 
        // anything inside it.
        res := map[string]interface{}{}
        
        // If r.URL.Path has no .json, it will display this error message in
        // JSON format.
        if r.URL.Path == "" || r.URL.Path[0] != '.' {
            res["status"] = "ERROR"
            res["err_msg"] = "No data type was specified"
            uadmin.ReturnJSON(w, r, res)
            return
        }

        // Initializes filterList as an array of string and valueList as an 
        // array of interface
        filterList := []string{}
        valueList := []interface{}{}

        // Gets the ID of the todo model, append to the filterList and
        // valueList
        if r.URL.Query().Get("todo_id") != "" {
            filterList = append(filterList, "todo_id = ?")
            valueList = append(valueList, r.URL.Query().Get("todo_id"))
        }

        // Concatenates filterList by AND to store all the data in the filter 
        // variable
        filter := strings.Join(filterList, " AND ")

        // Fetch Data from DB
        todo := []models.TODO{}
        uadmin.Filter(&todo, filter, valueList...)

        // Accesses and fetches data from another model
        for t := range todo {
            todo[t].Preload()
        }

        // Prints the todo in JSON format
        res["status"] = "ok"
        res["todo"] = todo
        uadmin.ReturnJSON(w, r, res)
    }

Finally, add this piece of code in the api.go shown below. This will establish a communication between the TodoListHandler and the APIHandler.

.. code-block:: go

    // APIHandler !
    func APIHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")
        if r.URL.Path == "/" {
            fmt.Fprintf(w, API_HELP)
        }
        // ------------------ ADD THIS CODE ------------------
        if strings.HasPrefix(r.URL.Path, "/todo_list") {
            TodoListHandler(w, r)
            return
        }
        // ------------------ ADD THIS CODE ------------------
    }

Now run your application. Suppose you have two data in your Todo model.

.. image:: assets/todomodeltwodata.png

|

If you go to /api/todo_list.json, you will see the list of each data in a more powerful way using JSON format.

.. image:: assets/todoapijson.png

|

Congrats, you know now how to do the following:

* Establishing a connection to the API
* Setting the path name using r.URL.Path
* How to use API Handlers
* Fetches data in another model

In the `next part`_, we will discuss about customizing your own API handler such as sorting the record in ascending or descending order, the starting point of execution process start until the assigned limit, and the action you want to perform in your database.

.. _next part: https://uadmin.readthedocs.io/en/latest/tutorial/part8.html