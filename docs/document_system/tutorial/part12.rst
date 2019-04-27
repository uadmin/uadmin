Document System Tutorial Part 12 - Custom Count function
========================================================
In this part, we will discuss about creating a custom Count function that checks the query and the UserID.

Go to document.go inside the models folder and create a function named **Count** that holds a, query, and args as an array of interface. args... means you can assign multiple values inside the function parameters.

.. code-block:: go

    // Count !
    func (d Document) Count(a interface{}, query interface{}, args ...interface{}) int {
        // Some codes
    }

Inside the function, convert a query parameter into a string. We will use that for some validation process later.

.. code-block:: go

    // Converts the query into a string
    Q := fmt.Sprint(query)

Let's check whether the string contains a query and a UserID.

.. code-block:: go

    if strings.Contains(Q, "user_id = ?") {
        // Some codes
    }

Inside the if statement, we have to split the query part by part for multiple queries.

.. code-block:: go

    // Split the query part by part
    qParts := strings.Split(Q, " AND ")

Initialize two variables for the query and an argument.

.. code-block:: go

    // Initialize tempArgs as an interface and tempQuery as a string
    tempArgs := []interface{}{}
    tempQuery := []string{}

Create a for loop statement that checks whether the specific query part is not equal to the UserID value. If it does, append the specific query part and arguments in a single variable.

.. code-block:: go

    // Loop the query every part
    for i := range qParts {
        // Checks whether the specific query part is not equal to the
        // UserID value
        if qParts[i] != "user_id = ?" {
            // Append the arguments into the tempArgs variable
            tempArgs = append(tempArgs, args[i])

            // Append the specific query part into the tempQuery variable
            tempQuery = append(tempQuery, qParts[i])
        }
    }

Now concatenate the tempQuery into the query variable to create a single string.

.. code-block:: go

    // Concatenate the query to create a single string
    query = strings.Join(tempQuery, " AND ")

Assign tempArgs object into the args variable.

.. code-block:: go

    // Assign tempArgs object into the args variable
    args = tempArgs

Outside the strings.Contains(Q, "user_id = ?") if statement, return the a, query, and args... inside the Count function parameters.

.. code-block:: go

    // Return the a, query, and args... inside the Count function parameters
    return uadmin.Count(a, query, args...)

In the `next part`_, we will talk about creating a custom AdminPage function that checks the query and the UserID.

Click `here`_ to view the full source code in this part.

.. _next part: https://uadmin.readthedocs.io/en/latest/document_system/tutorial/part13.html

.. _here: https://uadmin.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html

.. toctree::
   :maxdepth: 1

   full_code/part12
