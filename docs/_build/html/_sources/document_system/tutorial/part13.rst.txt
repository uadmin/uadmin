Document System Tutorial Part 13 - Custom AdminPage function
============================================================
In this part, we will talk about creating a custom AdminPage function that checks the query and the UserID.

First of all, what does AdminPage function do? It fetches records from the database with some standard rules such as sorting data, multiples of, and setting a limit that can be used in pagination.

.. code-block:: go

    func(order string, asc bool, offset int, limit int, a interface{}, query interface{}, args ...interface{}) (err error)

Parameters:

    **order string:** Is the field you want to specify in the database.

    **asc bool:** true in ascending order, false in descending order.

    **offset int:** Is the starting point of your list.

    **limit int:** Is the number of records that you want to display in your application.

    **a interface{}:** Is the variable where the model was initialized

    **query interface{}:** Is an action that you want to perform with in your data list

    **args ...interface{}:** Is the series of arguments for query input

Here is an example:

    **Model:** Ocean

    **Field to order:** Name

    **Sort:** Ascending

    **Index to start:** 0

    **How many records:** 10

    **Query:** Where ID is ?

    **Argument:** 5

The answer is uadmin.AdminPage("name", true, 0, 10, &Ocean, "id = ?", 5).

Go to document.go inside the models folder and create a function named **AdminPage** that holds order with the data type string, asc with the data type bool, offset and limit with the type int, and a, query, and args... with the data type interface{}. args... means you can assign multiple values inside the function parameters.

.. code-block:: go

    // AdminPage !
    func (d Document) AdminPage(order string, asc bool, offset int, limit int, a interface{}, query interface{}, args ...interface{}) (err error) {
        // Checks whether the starting point is less than 0
        if offset < 0 {
            offset = 0
        }

        // Converts the userID into uint because SQL Database reads the model ID
        // as uint
        userID := uint(0)

        // Converts the query into a string
        Q := fmt.Sprint(query)

        // Checks whether the string contains a query and a UserID
        if strings.Contains(Q, "user_id = ?") {
            // Prints the result for debugging
            uadmin.Trail(uadmin.DEBUG, "1")

            // Split the query part by part
            qParts := strings.Split(Q, " AND ")

            // Initialize tempArgs as an interface and tempQuery as a string
            tempArgs := []interface{}{}
            tempQuery := []string{}

            // Loop the query every part
            for i := range qParts {
                // Checks whether the specific query part is not equal to the
                // UserID value
                if qParts[i] != "user_id = ?" {
                    // Append the arguments into the tempArgs variable
                    tempArgs = append(tempArgs, args[i])

                    // Append the specific query part into the tempQuery variable
                    tempQuery = append(tempQuery, qParts[i])
                } else {
                    // Prints the result for debugging
                    uadmin.Trail(uadmin.DEBUG, "UserID: %d", args[i])
                    
                    // A type assertion that provides access to an interface
                    // value's (args[i]) underlying concrete value (uint).
                    userID, _ = (args[i]).(uint)
                }
            }
            // Concatenate the query to create a single string
            query = strings.Join(tempQuery, " AND ")

            // Assign tempArgs object into the args variable
            args = tempArgs
        }

        // Checks whether the userID is equal to 0
        if userID == 0 {
            // Prints the result for debugging
            uadmin.Trail(uadmin.DEBUG, "2")

            // Fetch the error by using AdminPage function
            err = uadmin.AdminPage(order, asc, offset, limit, a, query, args...)

            // Returns an error
            return err
        }

        // Initialize the user variable that calls the User model
        user := uadmin.User{}

        // Gets the ID of the user
        uadmin.Get(&user, "id = ?", userID)

        // Initialize docList and tempList that calls the Document model
        docList := []Document{}
        tempList := []Document{}

        // Loop execution
        for {
            // Fetch the error by using AdminPage function
            err = uadmin.AdminPage(order, asc, offset, limit, &tempList, query, args)
            uadmin.Trail(uadmin.DEBUG, "8: offset:%d, limit:%d", offset, limit)

            // Checks whether an error is not equal to nil
            if err != nil {
                // Prints the result for debugging
                uadmin.Trail(uadmin.DEBUG, "3")

                // Cast a model of interface as an array of Document then assigns
                // the docList object
                *a.(*[]Document) = docList

                // Return an error
                return err
            }

            // Checks whether the length of tempList is equal to 0
            if len(tempList) == 0 {
                // Prints the result for debugging
                uadmin.Trail(uadmin.DEBUG, "4")

                // Cast a model of interface as an array of Document then assigns
                // the docList object
                *a.(*[]Document) = docList

                // Prints the result for debugging
                uadmin.Trail(uadmin.DEBUG, "a: %#v", a)

                // Returns nothing
                return nil
            }

            // Loop the tempList values
            for i := range tempList {
                // Initialize p variable as Read permission
                p, _, _, _ := tempList[i].GetPermissions(user)

                // Checks whether the Document has read permission access
                if p {
                    // Prints the result for debugging
                    uadmin.Trail(uadmin.DEBUG, "5")

                    // Append the tempList (Document) object to the docList
                    // variable
                    docList = append(docList, tempList[i])
                }

                // Checks whether the length of docList is equal to the limit
                if len(docList) == limit {
                    // Prints the result for debugging
                    uadmin.Trail(uadmin.DEBUG, "6")

                    // Cast a model of interface as an array of Document then
                    // assigns the docList object
                    *a.(*[]Document) = docList

                    // Returns nothing
                    return nil
                }
            }

            // Add limit values to the offset variable
            offset += limit
        }
        // Cast a model of interface as an array of Document then assigns the
        // docList object
        *a.(*[]Document) = docList

        // Prints the result for debugging
        uadmin.Trail(uadmin.DEBUG, "7")

        // Returns nothing
        return nil
    }

In the `next part`_, we will discuss about displaying the permission status for each document records.

.. _next part: https://uadmin.readthedocs.io/en/latest/document_system/tutorial/part14.html
