uAdmin Tutorial Part 3 - API
============================

In this part, we will apply public uAdmin functions in our Todo list project.

Applying API Configurations
^^^^^^^^^^^^^^^^^^^^^^^^^^^
Let's go back to the main.go and apply **uadmin.Port** inside the main function. It assigns a port number to be used for http or https server. Let's say port number **8000**.

.. code-block:: go

    func main() {
        // Some codes are contained in this line ... (ignore this part)
        uadmin.Port = 8000
    }

If you run your code,

.. code-block:: bash

    [   OK   ]   Initializing DB: [12/12]
    [   OK   ]   Server Started: http://0.0.0.0:8000
            ___       __          _
    __  __/   | ____/ /___ ___  (_)___
    / / / / /| |/ __  / __  __ \/ / __ \
    / /_/ / ___ / /_/ / / / / / / / / / /
    \__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/

In the Server Started, it will redirect you to port number **8000**.

You can also set your own database settings in the main function. Add it above the uadmin.Register.

.. code-block:: go

    func main() {
        uadmin.Database = &uadmin.DBSettings{
            Type: "sqlite",
            Name: "todolist.db",
        }
        // Some codes are contained in this line ... (ignore this part)
    }

If you run your code,

.. code-block:: bash

    [   OK   ]   Initializing DB: [12/12]
    [   OK   ]   Initializing Languages: [185/185]
    [  INFO  ]   Auto generated admin user. Username: admin, Password: admin.
    [   OK   ]   Server Started: http://0.0.0.0:8000
            ___       __          _
    __  __/   | ____/ /___ ___  (_)___
    / / / / /| |/ __  / __  __ \/ / __ \
    / /_/ / ___ / /_/ / / / / / / / / / /
    \__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/

The todolist.db file is automatically created in your main project folder.

.. image:: assets/todolistdbhighlighted.png

|

However, if you go back to a specific model on your application, there is no any data inside it.

.. image:: assets/todoemptyagain.png

|

If you wish to revert it, go back to the main.go, change the **todolist.db** to **uadmin.db** in the Name field inside the uadmin.Database so that your application will access that database.

.. code-block:: go

    func main() {
        uadmin.Database = &uadmin.DBSettings{
            Type: "sqlite",
            Name: "uadmin.db",  // Replaced from todolist.db to uadmin.db
        }
        // Some codes are contained in this line ... (ignore this part)
    }

Output

.. image:: assets/todooutputback.png

|

uAdmin has a feature that allows a user to set his own site name by using uadmin.SiteName. Let's say **Todo List**.

.. code-block:: go

    func main() {
        // Some codes are contained in this line ... (ignore this part)
        uadmin.SiteName = "Todo List"
    }

Run your application and see the changes above the web browser.

.. image:: assets/todolisttitle.png

