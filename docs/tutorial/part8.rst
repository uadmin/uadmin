uAdmin Tutorial Part 8 - Customizing your API Handler
=====================================================
Before we start this tutorial, let's create at least 10 todos in the Todo model.

.. image:: assets/tentodos.png

|

For the case scenario, our client requests a data that returns only the last 5 activities sorted in descending order. In order to do that, use the public function called **uadmin.AdminPage**. AdminPage fetches records from the database with some standard rules such as sorting data, multiples of, and setting a limit that can be used in pagination. He also requests that the linking models should return only the name, not the other details within that model. Let's create another API file named "custom_list.go" containing the following codes below:

.. code-block:: go

    package api

    import (
        "net/http"
        "strings"

        "github.com/username/todo/models"
        "github.com/uadmin/uadmin"
    )

    // CustomListHandler !
    func CustomListHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/custom_list")
        res := map[string]interface{}{}
        if r.URL.Path == "" || r.URL.Path[0] != '.' {
            res["status"] = "ERROR"
            res["err_msg"] = "No data type was specified"
            uadmin.ReturnJSON(w, r, res)
            return
        }
        filterList := []string{}
        valueList := []interface{}{}
        if r.URL.Query().Get("todo_id") != "" {
            filterList = append(filterList, "todo_id = ?")
            valueList = append(valueList, r.URL.Query().Get("todo_id"))
        }
        filter := strings.Join(filterList, " AND ")

        // Fetch Data from DB
        todo := []models.TODO{}

        // Assigns a map as a string of interface to store any types of values
        results := []map[string]interface{}{}

        // Fetches the ID of todo in the first parameter, second parameter as 
        // false to sort in descending order, offset to 0 as a starting index
        // point in the third parameter, set the limit value to 5 to return 
        // five data in the fourth parameter, calls the model in the fifth
        // parameter, query interface is filter in the sixth parameter, and 
        // valueList is the argument called that can be used in the execution 
        // process as the last parameter.
        uadmin.AdminPage("id", false, 0, 5, &todo, filter, valueList)

        // Loop to fetch the record of todo
        for i := range todo {
            // Accesses and fetches the record of the linking models in Todo
            todo[i].Preload()

            // Assigns the string of interface in each Todo fields
            results = append(results, map[string]interface{}{
                "ID":          todo[i].ID,
                "Name":        todo[i].Name,
                "Description": todo[i].Description,
                // This returns only the name of the Category model, not the
                // other fields
                "Category":    todo[i].Category.Name,
                // This returns only the name of the Friend model, not the
                // other fields
                "Friend":      todo[i].Friend.Name,
                // This returns only the name of the Item model, not the other
                // fields
                "Item":        todo[i].Item.Name,
                "TargetDate":  todo[i].TargetDate,
                "Progress":    todo[i].Progress,
            })
        }

        // Prints the results in JSON format
        res["status"] = "ok"
        res["todo"] = results
        uadmin.ReturnJSON(w, r, res)
    }

Finally, add the following pieces of code in the api.go shown below. This will establish a communication between the CustomListHandler and the APIHandler.

.. code-block:: go

    const API_HELP = `TODO API HELP
    For more assistance please contact Integritynet:
    support@integritynet.biz

    - todo:
        ============
            # method     : todo_list
            # Parameters:  
            # Return    : json object that returns the list of your todo activities

        // ------------------ ADD THIS CODE ------------------
        ============
            # method     : custom_list
            # Parameters:  
            # Return    : json object that returns the list your last 5 todo activities sorted in descending order
        // ------------------ ADD THIS CODE ------------------
    `

    // APIHandler !
    func APIHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")
        if r.URL.Path == "/" {
            fmt.Fprintf(w, API_HELP)
        }
        if strings.HasPrefix(r.URL.Path, "/todo_list") {
            TodoListHandler(w, r)
            return
        }
        // ------------------ ADD THIS CODE ------------------
        if strings.HasPrefix(r.URL.Path, "/custom_list") {
            CustomListHandler(w, r)
            return
        }
        // ------------------ ADD THIS CODE ------------------
    }

Now run your application. If you go to /api/custom_list.json, you will see the list of your last 5 activities sorted in descending order in a more powerful way using JSON format.

.. image:: assets/todoapicustomjson.png

|

Congrats, you know now how to customize your own API by returning the data based on the limit, sorting the data in descending order, and assigning a value to the submodel that returns only one field.

In the `next part`_, we will discuss about inserting the data to the models through the API that fetches multiple values.

.. _next part: https://uadmin.readthedocs.io/en/latest/tutorial/part9.html