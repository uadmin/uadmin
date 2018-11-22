API Reference
=============
Here are all public functions in the uAdmin, their syntax, and how to use them in the project.

* `uadmin.Action`_
* `uadmin.AdminPage`_
* `uadmin.All`_
* `uadmin.BindIP`_
* `uadmin.Choice`_
* `uadmin.ClearDB`_
* `uadmin.CookieTimeout`_
* `uadmin.Count`_
* `uadmin.CustomTranslation`_
* `uadmin.DashboardMenu`_
* `uadmin.Database`_
* `uadmin.DBSettings`_
* `uadmin.DEBUG`_
* `uadmin.DebugDB`_
* `uadmin.Delete`_
* `uadmin.DeleteList`_
* `uadmin.EmailFrom`_
* `uadmin.EmailPassword`_
* `uadmin.EmailSMTPServer`_
* `uadmin.EmailSMTPServerPort`_
* `uadmin.EmailUsername`_
* `uadmin.ERROR`_
* `uadmin.F`_
* `uadmin.Filter`_
* `uadmin.FilterBuilder`_
* `uadmin.GenerateBase32`_
* `uadmin.GenerateBase64`_
* `uadmin.Get`_
* `uadmin.GetDB`_
* `uadmin.GetID`_
* `uadmin.GetString`_
* `uadmin.GetUserFromRequest`_
* `uadmin.GroupPermission`_
* `uadmin.HideInDashboarder`_
* `uadmin.INFO`_
* `uadmin.IsAuthenticated`_
* `uadmin.JSONMarshal`_
* `uadmin.Language`_
* `uadmin.Log`_
* `uadmin.Login`_
* `uadmin.Login2FA`_
* `uadmin.Logout`_
* `uadmin.MaxImageHeight`_
* `uadmin.MaxImageWidth`_
* `uadmin.MaxUploadFileSize`_
* `uadmin.Model`_
* `uadmin.ModelSchema`_
* `uadmin.MongoDB (Experimental)`_
* `uadmin.MongoModel (Experimental)`_
* `uadmin.MongoSettings (Experimental)`_
* `uadmin.NewModel`_
* `uadmin.NewModelArray`_
* `uadmin.OK`_
* `uadmin.OTPAlgorithm`_
* `uadmin.OTPDigits`_
* `uadmin.OTPPeriod`_
* `uadmin.OTPSkew`_
* `uadmin.PageLength`_
* `uadmin.Port`_
* `uadmin.Preload`_
* `uadmin.PublicMedia`_
* `uadmin.Register`_
* `uadmin.RegisterInlines`_
* `uadmin.ReportingLevel`_
* `uadmin.ReportTimeStamp`_
* `uadmin.ReturnJSON`_
* `uadmin.RootURL`_
* `uadmin.Salt`_
* `uadmin.Save`_
* `uadmin.Schema`_
* `uadmin.SendEmail`_
* `uadmin.Session`_
* `uadmin.SiteName`_
* `uadmin.StartSecureServer`_
* `uadmin.StartServer`_
* `uadmin.Tf`_
* `uadmin.Theme`_
* `uadmin.Trail`_
* `uadmin.Translate`_
* `uadmin.Update`_
* `uadmin.UploadImageHandler`_
* `uadmin.User`_
* `uadmin.UserGroup`_
* `uadmin.UserPermission`_
* `uadmin.Version`_
* `uadmin.WARNING`_
* `uadmin.WORKING`_

Functions
---------

**uadmin.Action**
^^^^^^^^^^^^^^^^^
Action is the process of doing something where you can check the status of your activities in the uAdmin project.

Syntax:

.. code-block:: go

    type Action int

There are 11 methods of actions:

* **Added** - Saved a new record
* **Custom** - For any other action that you would like to log
* **Deleted** - Deleted a record
* **LoginDenied** - User invalid login
* **LoginSuccessful** - User login
* **Logout** - User logout
* **Modified** - Save an existing record
* **PasswordResetDenied** - A password reset attempt was rejected
* **PasswordResetRequest** - A password reset was received
* **PasswordResetSuccessful** - A password was reset
* **Read** - Opened a record

Open "LOGS" in the uAdmin dashboard. You can see the Action field inside it as shown below.

.. image:: assets/actionhighlighted.png

|

Now go to the main.go. Let's add each methods of actions in the log.

.. code-block:: go

    func main(){
        // Some codes
        for i := 0; i < 11; i++ {
            // Initialize the log model
            log := uadmin.Log{}

            // Call each methods of action based on the specific loop count
            switch i {
            case 0:
                log.Action = uadmin.Action.Added(0)
            case 1:
                log.Action = uadmin.Action.Custom(0)
            case 2:
                log.Action = uadmin.Action.Deleted(0)
            case 3:
                log.Action = uadmin.Action.LoginDenied(0)
            case 4:
                log.Action = uadmin.Action.LoginSuccessful(0)
            case 5:
                log.Action = uadmin.Action.Logout(0)
            case 6:
                log.Action = uadmin.Action.Modified(0)
            case 7:
                log.Action = uadmin.Action.PasswordResetDenied(0)
            case 8:
                log.Action = uadmin.Action.PasswordResetRequest(0)
            case 9:
                log.Action = uadmin.Action.PasswordResetSuccessful(0)
            default:
                log.Action = uadmin.Action.Read(0)
            }

            // Add the method to the logs
            log.Save()
        }
    }

Once you are done, rebuild your application. Check your "LOGS" again to see the result.

.. image:: assets/actionlist.png

|

As expected, all types of actions were added in the logs. Good job man!
    
**uadmin.AdminPage**
^^^^^^^^^^^^^^^^^^^^
AdminPage fetches records from the database with some standard rules such as sorting data, multiples of, and setting a limit that can be used in pagination.

Syntax:

.. code-block:: go

    func(order string, asc bool, offset int, limit int, a interface{}, query interface{}, args ...interface{}) (err error)

Parameters:

    **order string:** Is the field you want to specify in the database.

    **asc bool:** true in ascending order, false in descending order.

    **offset int:** Is the starting point of your list.

    **limit int:** Is until where an element should be taken in your list from database.

    **a interface{}:** Is the variable where the model name was initialized.

    **query interface{}:** Is an action that you want to perform with in your data list.

    **args ...interface{}:** Is the variable or container that can be used in execution process.

See `Tutorial Part 8 - Customizing your API Handler`_ for the example.

.. _Tutorial Part 8 - Customizing your API Handler: https://uadmin.readthedocs.io/en/latest/tutorial/part8.html

**uadmin.All**
^^^^^^^^^^^^^^
All fetches all object in the database.

Syntax:

.. code-block:: go

    func(a interface{}) (err error)

Parameters:

    **a interface{}:** Is the variable where the model name was initialized.

Before we proceed to the example, read `Tutorial Part 7 - Introduction to API`_ to familiarize how API works in uAdmin.

.. _Tutorial Part 7 - Introduction to API: https://uadmin.readthedocs.io/en/latest/tutorial/part7.html

Create a file named friend_list.go inside the api folder with the following codes below:

.. code-block:: go

    // FriendListHandler !
    func FriendListHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/friend_list")

        res := map[string]interface{}{}

        friend := []models.Friend{}
        uadmin.All(&friend) // <-- place it here

        res["status"] = "ok"
        res["todo"] = friend
        uadmin.ReturnJSON(w, r, res)
    }

Establish a connection in the main.go to the API by using http.HandleFunc. It should be placed after the uadmin.Register and before the StartServer.

.. code-block:: go

    func main() {
        // Some codes

        // FilterListHandler
        http.HandleFunc("/friend_list/", api.FriendListHandler) // <-- place it here
    }

api is the folder name while FriendListHandler is the name of the function inside friend_list.go.

Run your application and see what happens.

.. image:: assets/friendlistapi.png
   :align: center

**uadmin.BindIP**
^^^^^^^^^^^^^^^^^
BindIP is the IP the application listens to.

Syntax:

.. code-block:: go

    BindIP string

Go to the main.go. Connect to the server using a private IP e.g. (10.x.x.x,192.168.x.x, 127.x.x.x or ::1). Let's say **127.0.0.2**

.. code-block:: go

    func main() {
        // Some codes
        uadmin.BindIP = "127.0.0.2" // <--  place it here
    }

If you run your code,

.. code-block:: bash

    [   OK   ]   Initializing DB: [12/12]
    [   OK   ]   Server Started: http://127.0.0.2:8080
             ___       __          _
      __  __/   | ____/ /___ ___  (_)___
     / / / / /| |/ __  / __  __ \/ / __ \
    / /_/ / ___ / /_/ / / / / / / / / / /
    \__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/

In the Server Started, it will redirect you to the IP address of **127.0.0.2**.

But if you connect to other private IP addresses, it will not work as shown below (User connects to 127.0.0.3).

.. image:: tutorial/assets/bindiphighlighted.png

**uadmin.Choice**
^^^^^^^^^^^^^^^^^
Choice is a struct for the list of choices.

Syntax:

.. code-block:: go

    type Choice struct {
        V        string
        K        uint
        Selected bool
    }

Suppose I have four records in my Category model.

* Education ID = 4
* Family ID = 3
* Work ID = 2
* Travel ID = 1

.. image:: assets/categorylist.png

Create a function with a parameter of interface{} and a pointer of User that returns an array of Choice which will be used that later below the main function in main.go.

.. code-block:: go

    func GetChoices(m interface{}, user *uadmin.User) []uadmin.Choice {
        // Initialize the Category model
        categorylist := models.Category{}

        // Get the ID of the category
        uadmin.Get(&categorylist, "id = 4")

        // Build choices
        choices := []uadmin.Choice{}

        // Append by getting the ID and string of categorylist
        choices = append(choices, uadmin.Choice{
            V:        uadmin.GetString(categorylist),
            K:        uadmin.GetID(reflect.ValueOf(categorylist)),
            Selected: true,
        })

        return choices
    }

Now inside the main function, apply `uadmin.Schema`_ function that calls a model name of "todo", accesses "Choices" as the field name that uses the LimitChoicesTo then assign it to GetChoices which is your function name.

.. code-block:: go

    uadmin.Schema["todo"].FieldByName("Choices").LimitChoicesTo = GetChoices

Run your application, go to the Todo model and see what happens in the Choices field.

.. image:: assets/choicesid4.png

|

When you notice, the Education is automatically selected. This function has the ability to search whatever you want in the drop down list.

You can also add or replace the list of choices manually. This time, set the value of the Selected to false.

.. code-block:: go

    func GetChoices(m interface{}, user *uadmin.User) []uadmin.Choice {
        // Initialize the Category model
        categorylist := models.Category{}

        // Build choices
        choices := []uadmin.Choice{}

        // Append by getting the ID and string of categorylist
        choices = append(choices, uadmin.Choice{
            V:        uadmin.GetString(categorylist),
            K:        uadmin.GetID(reflect.ValueOf(categorylist)),
            Selected: true,
        })

        // Create the list of choices manually
        choices = append(choices, uadmin.Choice{
            V:        "Tour",
            K:        1,
            Selected: false,
        })
        choices = append(choices, uadmin.Choice{
            V:        "Employment",
            K:        2,
            Selected: false,
        })
        choices = append(choices, uadmin.Choice{
            V:        "Clan",
            K:        3,
            Selected: false,
        })
        choices = append(choices, uadmin.Choice{
            V:        "Learning",
            K:        4,
            Selected: false,
        })

        return choices
    }

Now rerun your application to see the result.

.. image:: assets/manualchoiceslist.png

|

When you notice, the values of the Category field were replaced in the choices list. You can also type whatever you want to search in the choices list above. For this example, let's choose "Learning".

Once you are done, save the record and see what happens.

.. image:: assets/choicesid4manualoutput.png

It returns Education because choices is based on the GetString of categorylist.

**uadmin.ClearDB**
^^^^^^^^^^^^^^^^^^
ClearDB clears the database object.

Syntax:

.. code-block:: go

    func()

Suppose I have two databases in my project folder.

.. image:: assets/twodatabases.png

|

And I set the Name to **uadmin.db** on Database Settings in main.go.

.. code-block:: go

    func main(){
        uadmin.Database = &uadmin.DBSettings{
            Type: "sqlite",
            Name: "uadmin.db",
        }
        // Some codes
    }

Let's create a new file in the models folder named "expression.go" with the following codes below:

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // ---------------- DROP DOWN LIST ----------------
    // Status ...
    type Status int

    // Keep ...
    func (s Status) Keep() Status {
        return 1
    }

    // ClearDatabase ...
    func (s Status) ClearDatabase() Status {
        return 2
    }
    // -----------------------------------------------

    // Expression model ...
    type Expression struct {
        uadmin.Model
        Name   string `uadmin:"required"`
        Status Status `uadmin:"required"`
    }
    
    // Save ...
    func (e *Expression) Save() {
        // If Status is equal to ClearDatabase(), the database
        // will reset and open a new one which is todolist.db.
        if e.Status == e.Status.ClearDatabase() {
            db := uadmin.GetDB()    // <-- Returns a pointer to the DB
            uadmin.ClearDB()        // <-- Place it here

            // Database configurations
            uadmin.Database = &uadmin.DBSettings{
                Type: "sqlite",
                Name: "todolist.db",
            }

            // Instantiate
            db2 := uadmin.GetDB()
            
            // Close the old ones
            db.Close()

            // Open the new ones
            db2.Begin()
        }

        // Override save
        uadmin.Save(e)
    }

Register your Expression model in the main function.

.. code-block:: go

    func main() {

        // Some codes contained in this part

        uadmin.Register(
            // Some registered models
            models.Expression{}, // <-- place it here
        )

        // Some codes contained in this part
    }

Run the application. Go to the Expressions model and add at least 3 interjections, all Status set to "Keep".

.. image:: assets/expressionkeep.png

|

Now create another data, this time set the Status as "Clear Database" and see what happens.

.. image:: assets/cleardatabase.png

|

Your account will automatically logout in the application. Login your account again, go to the Expressions model and see what happens.

.. image:: assets/cleardatabasesecondmodel.png

|

As expected, all previous records were gone in the model. It does not mean that they were deleted. It's just that you have opened a new database called "todolist.db". Check out the other models that you have. You may notice that something has changed in your database.

**uadmin.CookieTimeout**
^^^^^^^^^^^^^^^^^^^^^^^^
CookieTimeout is the timeout of a login cookie in seconds.

Syntax:

.. code-block:: go

    int

Let's apply this function in the main.go.

.. code-block:: go

    func main() {
        // Some codes
        uadmin.CookieTimeout = 10 // <--  place it here
    }

.. WARNING::
   Use it at your own risk. Once the cookie expires, your account will be permanently deactivated. In this case, you must have an extra user account in the User database.

Login your account, wait for 10 seconds and see what happens.

.. image:: tutorial/assets/loginform.png

It will redirect you to the login form because your cookie has already been expired.

**uadmin.Count**
^^^^^^^^^^^^^^^^
Count return the count of records in a table based on a filter.

Syntax:

.. code-block:: go

    func(a interface{}, query interface{}, args ...interface{}) int

Parameters:

    **a interface{}:** Is the variable where the model name was initialized.

    **query interface{}:** Is an action that you want to perform with in your data list.

    **args ...interface{}:** Is the variable or container that can be used in execution process.

See `uadmin.Get`_ for the example.

**uadmin.CustomTranslation**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^
CustomTranslation allows a user to customize any languages in the uAdmin system.

Syntax:

.. code-block:: go

    []string

Suppose that English is the only active language in your application. Go to the main.go and apply the following codes below. It should be placed before uadmin.Register.

.. code-block:: go

    func main(){
        // Place it here
        uadmin.CustomTranslation = []string{"models/custom", "models/todo_custom"}

        uadmin.Register(
            // Some codes
        )
    }

From your project folder, go to static/i18n/models. You will notice that two JSON files are created in the models folder.

.. image:: assets/customtranslationcreate.png

Every JSON file is per language. In other words, if you have 2 languages available in your application, there will be a total of 4 created JSON files.

**uadmin.DashboardMenu**
^^^^^^^^^^^^^^^^^^^^^^^^
DashboardMenu is a system in uAdmin that is used to add, modify and delete the elements of a model.

Syntax:

.. code-block:: go

    type DashboardMenu struct {
        Model
        MenuName string `uadmin:"required;list_exclude;multilingual;filter"`
        URL      string `uadmin:"required"`
        ToolTip  string
        Icon     string `uadmin:"image"`
        Cat      string `uadmin:"filter"`
        Hidden   bool   `uadmin:"filter"`
    }

There is a function that you can use in DashboardMenu:

* **String()** - returns the MenuName

Go to the main.go and apply the following codes below after the RegisterInlines section.

.. code-block:: go

    func main(){

        // Some codes

        dashboardmenu := uadmin.DashboardMenu{
            MenuName: "Expressions",
            URL:      "expression",
            ToolTip:  "",
            Icon:     "/media/images/expression.png",
            Cat:      "Yeah!",
            Hidden:   false,
        }

        // This will create a new model based on the information assigned in
        // the dashboardmenu variable.
        uadmin.Save(&dashboardmenu)
    }

Now run your application and see what happens.

.. image:: assets/expressionmodelcreated.png

|

You can also apply a String() function in uadmin.DashboardMenu which returns a MenuName. Go to the main.go and apply the following codes below.

.. code-block:: go

    func main(){
        // Some codes
        dashboardmenu := uadmin.DashboardMenu{
            MenuName: "Model",
        }
        uadmin.Trail(uadmin.INFO, "String() returns %s", dashboardmenu.String())
    }

Result

.. code-block:: bash

    [  INFO  ]   String() returns Model

**uadmin.Database**
^^^^^^^^^^^^^^^^^^^
Database is the active Database settings.

Syntax:

.. code-block:: go

    *uadmin.DBSettings

There are 6 fields that you can use in this function:

* **Host** - returns a string
* **Name** - returns a string. This will generate a database file in your project folder.
* **Password** - returns a string
* **Port** - returns an int. It is the port used for http or https server.
* **Type** - returns a string. There are 2 types: SQLLite and MySQL.
* **User** - returns a string

Go to the main.go in your Todo list project. Add the codes below above the uadmin.Register.

.. code-block:: go

    func main(){
        database := uadmin.Database
        database.Host = "192.168.149.108"
        database.Name = "todolist.db"
        database.Password = "admin"
        database.Port = 8000
        database.Type = "sqlite"
        database.User = "admin"
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

.. image:: tutorial/assets/todolistdbhighlighted.png

|

See `uadmin.DBSettings`_ for the process of configuring your database in MySQL.

**uadmin.DBSettings**
^^^^^^^^^^^^^^^^^^^^^
DBSettings is a feature that allows a user to configure the settings of a database.

Syntax:

.. code-block:: go

    type DBSettings struct {
        Type     string // SQLLite, MySQL
        Name     string // File/DB name
        User     string
        Password string
        Host     string
        Port     int
    }

Go to the main.go in your Todo list project. Add the codes below above the uadmin.Register.

.. code-block:: go

    func main() {
        uadmin.Database = &uadmin.DBSettings{
            Type:      "sqlite",
            Name:      "todolist.db",
            User:      "admin",
            Password:  "admin",
            Host:      "192.168.149.108",
            Port:      8000,
        }
        // Some codes
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

.. image:: tutorial/assets/todolistdbhighlighted.png

|

You can also migrate your application into the MySQL database server. In order to do that, you must have the `MySQL Workbench`_ application installed on your computer. Open your MySQL Workbench and set up your Connection Name (example below is uadmin). Hostname, Port and Username are automatically provided for you but you can change the values there if you wish to. For this example, let's apply the following information below.

.. _MySQL Workbench: https://dev.mysql.com/downloads/workbench/

.. image:: assets/mysqlsetup.png

|

Click Test Connection to see if the connection is working properly.

.. image:: assets/mysqlprompt.png
   :align: center

|

Result

.. image:: assets/testconnectionresult.png
   :align: center

|

Once you are done with the connection testing, click OK on the bottom right corner. You will see the interface of the application. Let's create a new schema by right clicking the area on the bottom left corner highlighted below then select "Create Schema".

.. image:: assets/rightclickarea.png

|

Input the value of the schema name as "todo" then click Apply.

.. image:: assets/schemasetuptodo.png

|

You will see the Apply SQL Script to the Database form. Leave it as it is and click Apply.

.. image:: assets/applysqlscriptform.png

|

Your todo schema has been created in the MySQL. Congrats!

.. image:: assets/todocreatedmysql.png
   :align: center

|

Now go back to your todo list project. Open main.go and apply the following codes below:

.. code-block:: go

    uadmin.Database = &uadmin.DBSettings{
        Type:     "mysql",
        Name:     "todo",
        User:     "root",
        Password: "todolist",
        Host:     "127.0.0.1",
        Port:     3306,
    }

The information above is well-based on the database configuration settings in MySQL Workbench.

Once you are done, run your application and see what happens.

.. code-block:: bash

    [   OK   ]   Initializing Languages: [185/185]
    [  INFO  ]   Auto generated admin user. Username:admin, Password:admin.
    [   OK   ]   Server Started: http://0.0.0.0:8080

Open your browser and type the IP address above. Then login using “admin” as username and password.

.. image:: tutorial/assets/loginform.png

|

You will be greeted by the uAdmin dashboard. System models are built in to uAdmin, and the rest are the ones we created, in this case TODOS model.

.. image:: assets/uadmindashboard.png

|

Now open your MySQL Workbench. On todo database in the schema panel, the tables are automatically generated from your uAdmin dashboard.

.. image:: assets/mysqluadminmodelslist.png
   :align: center

Congrats, now you know how to configure your database settings in both SQLite and MySQL.

**uadmin.DEBUG**
^^^^^^^^^^^^^^^^
DEBUG is the display tag under Trail. It is the process of identifying and removing errors.

Syntax:

.. code-block:: go

    untyped int

See `uadmin.Trail`_ for the example.

**uadmin.DebugDB**
^^^^^^^^^^^^^^^^^^
DebugDB prints all SQL statements going to DB.

Syntax:

.. code-block:: go

    bool

Go to the main.go. Set this function as true.

.. code-block:: go

    func main(){
        uadmin.DebugDB = true
        // Some codes contained in this part
    }

Check your terminal to see the result.

.. code-block:: bash

    [   OK   ]   Initializing DB: [13/13]

    (/home/dev1/go/src/github.com/uadmin/uadmin/db.go:428) 
    [2018-11-10 12:43:07]  [0.09ms]  SELECT count(*) FROM "languages"  WHERE "languages"."deleted_at" IS NULL  
    [0 rows affected or returned ] 

    (/home/dev1/go/src/github.com/uadmin/uadmin/db.go:298) 
    [2018-11-10 12:43:07]  [0.17ms]  SELECT * FROM "languages"  WHERE "languages"."deleted_at" IS NULL AND ((active = 'true'))  
    [1 rows affected or returned ] 

    (/home/dev1/go/src/github.com/uadmin/uadmin/db.go:238) 
    [2018-11-10 12:43:07]  [0.16ms]  SELECT * FROM "languages"  WHERE "languages"."deleted_at" IS NULL AND ((`default` = 'true')) ORDER BY "languages"."id" ASC LIMIT 1  
    [1 rows affected or returned ] 

    (/home/dev1/go/src/github.com/uadmin/uadmin/db.go:162) 
    [2018-11-10 12:43:07]  [0.32ms]  SELECT * FROM "dashboard_menus"  WHERE "dashboard_menus"."deleted_at" IS NULL  
    [13 rows affected or returned ] 

    (/home/dev1/go/src/github.com/uadmin/uadmin/db.go:428) 
    [2018-11-10 12:43:07]  [0.07ms]  SELECT count(*) FROM "users"  WHERE "users"."deleted_at" IS NULL  
    [0 rows affected or returned ] 

**uadmin.Delete**
^^^^^^^^^^^^^^^^^
Delete records from database.

Syntax:

.. code-block:: go

    func(a interface{}) (err error)

Parameters:

    **a interface{}:** Is the variable where the model name was initialized.

Let's create a new file in the models folder named "expression.go" with the following codes below:

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // ---------------- DROP DOWN LIST ----------------
    // Status ...
    type Status int

    // Keep ...
    func (s Status) Keep() Status {
        return 1
    }

    // DeletePrevious ...
    func (s Status) DeletePrevious() Status {
        return 2
    }
    // -----------------------------------------------

    // Expression model ...
    type Expression struct {
        uadmin.Model
        Name   string `uadmin:"required"`
        Status Status `uadmin:"required"`
    }

    // Save ...
    func (e *Expression) Save() {
        // If Status is equal to DeletePrevious(), it will delete
        // the previous data in the list.
        if e.Status == e.Status.DeletePrevious() {
            uadmin.Delete(e) // <-- place it here
        }

        uadmin.Save(e)
    }

Register your Expression model in the main function.

.. code-block:: go

    func main() {

        // Some codes contained in this part

        uadmin.Register(
            // Some registered models
            models.Expression{}, // <-- place it here
        )

        // Some codes contained in this part
    }

Run the application. Go to the Expressions model and add at least 3 interjections, all Status set to "Keep".

.. image:: assets/expressionkeep.png

|

Now create another data, this time set the Status as "Delete Previous" and see what happens.

.. image:: assets/deleteprevious.png

|

Result

.. image:: assets/deletepreviousresult.png

|

All previous records are deleted from the database.

**uadmin.DeleteList**
^^^^^^^^^^^^^^^^^^^^^
Delete the list of records from database.

Syntax:

.. code-block:: go

    func(a interface{}, query interface{}, args ...interface{}) (err error)

Parameters:

    **a interface{}:** Is the variable where the model name was initialized.

    **query interface{}:** Is an action that you want to perform with in your data list.

    **args ...interface{}:** Is the variable or container that can be used in execution process.

Let's create a new file in the models folder named "expression.go" with the following codes below:

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // ---------------- DROP DOWN LIST ----------------
    // Status ...
    type Status int

    // Keep ...
    func (s Status) Keep() Status {
        return 1
    }

    // Custom ...
    func (s Status) Custom() Status {
        return 2
    }

    // DeleteCustom ...
    func (s Status) DeleteCustom() Status {
        return 3
    }
    // -----------------------------------------------

    // Expression model ...
    type Expression struct {
        uadmin.Model
        Name   string `uadmin:"required"`
        Status Status `uadmin:"required"`
    }

    // Save ...
    func (e *Expression) Save() {
        // If Status is equal to DeleteCustom(), it will delete the
        // list of data that contains Custom as the status.
        if e.Status == e.Status.DeleteCustom() {
            uadmin.DeleteList(&e, "status = ?", 2)
        }

        uadmin.Save(e)
    }

Register your Expression model in the main function.

.. code-block:: go

    func main() {

        // Some codes contained in this part

        uadmin.Register(
            // Some registered models
            models.Expression{}, // <-- place it here
        )

        // Some codes contained in this part
    }

Run the application. Go to the Expressions model and add at least 3 interjections, one is set to "Keep" and the other two is set to "Custom".

.. image:: assets/expressionkeepcustom.png

|

Now create another data, this time set the Status as "Delete Custom" and see what happens.

.. image:: assets/deletecustom.png

|

Result

.. image:: assets/deletecustomresult.png

|

All custom records are deleted from the database.

**uadmin.EmailFrom**
^^^^^^^^^^^^^^^^^^^^
EmailFrom identifies where the email is coming from.

Syntax:

.. code-block:: go

    string

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){
        uadmin.EmailFrom = "myemail@integritynet.biz"
        uadmin.EmailUsername = "myemail@integritynet.biz"
        uadmin.EmailPassword = "abc123"
        uadmin.EmailSMTPServer = "smtp.integritynet.biz"
        uadmin.EmailSMTPServerPort = 587
        // Some codes
    }

Let's go back to the uAdmin dashboard, go to Users model, create your own user account and set the email address based on your assigned EmailFrom in the code above.

.. image:: tutorial/assets/useremailhighlighted.png

|

Log out your account. At the moment, you suddenly forgot your password. How can we retrieve our account? Click Forgot Password at the bottom of the login form.

.. image:: tutorial/assets/forgotpasswordhighlighted.png

|

Input your email address based on the user account you wish to retrieve it back.

.. image:: tutorial/assets/forgotpasswordinputemail.png

|

Once you are done, open your email account. You will receive a password reset notification from the Todo List support. To reset your password, click the link highlighted below.

.. image:: tutorial/assets/passwordresetnotification.png

|

You will be greeted by the reset password form. Input the following information in order to create a new password for you.

.. image:: tutorial/assets/resetpasswordform.png

Once you are done, you can now access your account using your new password.

**uadmin.EmailPassword**
^^^^^^^^^^^^^^^^^^^^^^^^
EmailPassword sets the password of an email.

Syntax:

.. code-block:: go

    string

See `uadmin.EmailFrom`_ for the example.

**uadmin.EmailSMTPServer**
^^^^^^^^^^^^^^^^^^^^^^^^^^
EmailSMTPServer sets the name of the SMTP Server in an email.

Syntax:

.. code-block:: go

    string

See `uadmin.EmailFrom`_ for the example.

**uadmin.EmailSMTPServerPort**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
EmailSMTPServerPort sets the port number of an SMTP Server in an email.

Syntax:

.. code-block:: go

    int

See `uadmin.EmailFrom`_ for the example.

**uadmin.EmailUsername**
^^^^^^^^^^^^^^^^^^^^^^^^
EmailUsername sets the username of an email.

Syntax:

.. code-block:: go

    string

See `uadmin.EmailFrom`_ for the example.

**uadmin.ERROR**
^^^^^^^^^^^^^^^^
ERROR is a status to notify the user that there is a problem in an application.

Syntax:

.. code-block:: go

    untyped int

See `uadmin.Trail`_ for the example.

**uadmin.F**
^^^^^^^^^^^^
F is a field.

Syntax:

.. code-block:: go

    type F struct {
        Name              string
        DisplayName       string
        Type              string
        Value             interface{}
        Help              string
        Max               interface{}
        Min               interface{}
        Format            string
        DefaultValue      string
        Required          bool
        Pattern           string
        PatternMsg        string
        Hidden            bool
        ReadOnly          string
        Searchable        bool
        Filter            bool
        ListDisplay       bool
        FormDisplay       bool
        CategoricalFilter bool
        Translations      []translation
        Choices           []Choice
        IsMethod          bool
        ErrMsg            string
        ProgressBar       map[float64]string
        LimitChoicesTo    func(interface{}, *User) []Choice
        UploadTo          string
        Encrypt           bool
    }

There are 2 ways you can do for initialization process using this function: one-by-one and by group.

One-by-one initialization:

.. code-block:: go

    func main(){
        // Some codes
        f := uadmin.F{}
        f.Name = "Name"
        f.DisplayName = "DisplayName"
        f.Type = "Type"
        f.Value = "Value"
    }

By group initialization:

.. code-block:: go

    func main(){
        // Some codes
        f := uadmin.F{
            Name:        "Name",
            DisplayName: "DisplayName",
            Type:        "Type",
            Value:       "Value",
        }
    }

In this example, we will use "by group" initialization process.

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Some codes
        f1 := uadmin.F{
            Name:        "Name",
            DisplayName: "Reaction",
            Type:        "string",
            Value:       "Wow!",
        }
        f2 := uadmin.F{
            Name:        "Reason",
            DisplayName: "Reason",
            Type:        "string",
            Value:       "My friend's performance is amazing.",
        }
    }

The code above shows the two initialized F structs using the Name, DisplayName, Type, and Value fields.

See `uadmin.ModelSchema`_ for the continuation of this example.

**uadmin.Filter**
^^^^^^^^^^^^^^^^^
Filter fetches records from the database.

Syntax:

.. code-block:: go

    func(a interface{}, query interface{}, args ...interface{}) (err error)

Parameters:

    **a interface{}:** Is the variable where the model name was initialized.

    **query interface{}:** Is an action that you want to perform with in your data list.

    **args ...interface{}:** Is the variable or container that can be used in execution process.

Before we proceed to the example, read `Tutorial Part 7 - Introduction to API`_ to familiarize how API works in uAdmin.

.. _Tutorial Part 7 - Introduction to API: https://uadmin.readthedocs.io/en/latest/tutorial/part7.html

Create a file named filter_list.go inside the api folder with the following codes below:

.. code-block:: go

    package api

    import (
        "net/http"
        "strings"

        "github.com/username/todo/models"
        "github.com/uadmin/uadmin"
    )

    // FilterListHandler !
    func FilterListHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/filter_list")

        res := map[string]interface{}{}

        filterList := []string{}
        valueList := []interface{}{}
        if r.URL.Query().Get("todo_id") != "" {
            filterList = append(filterList, "todo_id = ?")
            valueList = append(valueList, r.URL.Query().Get("todo_id"))
        }
        filter := strings.Join(filterList, " AND ")

        todo := []models.Todo{}
        results := []map[string]interface{}{}

        uadmin.Filter(&todo, filter, valueList) // <-- place it here

        // This loop returns only the name of your todo list.
        for i := range todo {
            results = append(results, map[string]interface{}{
                "Name": todo[i].Name,
            })
        }

        res["status"] = "ok"
        res["todo"] = results
        uadmin.ReturnJSON(w, r, res)
    }

Establish a connection in the main.go to the API by using http.HandleFunc. It should be placed after the uadmin.Register and before the StartServer.

.. code-block:: go

    func main() {
        // Some codes

        // FilterListHandler
        http.HandleFunc("/filter_list/", api.FilterListHandler) // <-- place it here
    }

api is the folder name while FilterListHandler is the name of the function inside filter_list.go.

Run your application and see what happens.

.. image:: assets/filterlistapi.png
   :align: center

See `uadmin.Preload`_ for more examples of using this function.

**uadmin.FilterBuilder**
^^^^^^^^^^^^^^^^^^^^^^^^
FilterBuilder changes a map filter into a query.

Syntax:

.. code-block:: go

    func(params map[string]interface{}) (query string, args []interface{})

Before we proceed to the example, read `Tutorial Part 7 - Introduction to API`_ to familiarize how API works in uAdmin.

.. _Tutorial Part 7 - Introduction to API: https://uadmin.readthedocs.io/en/latest/tutorial/part7.html

Suppose you have ten records in your Todo model.

.. image:: tutorial/assets/tendataintodomodel.png

|

Create a file named filterbuilder.go inside the api folder with the following codes below:

.. code-block:: go

    package api

    import (
        "net/http"
        "strings"

        "github.com/rn1hd/todo/models"
        "github.com/uadmin/uadmin"
    )

    // FilterBuilderHandler !
    func FilterBuilderHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/filterbuilder")

        res := map[string]interface{}{}

        filterList := []string{}
        valueList := []interface{}{}
        if r.URL.Query().Get("todo_id") != "" {
            filterList = append(filterList, "todo_id = ?")
            valueList = append(valueList, r.URL.Query().Get("todo_id"))
        }

        todo := []models.TODO{}

        query, args := uadmin.FilterBuilder(res) // <-- place it here
        uadmin.Filter(&todo, query, args)
        for t := range todo {
            todo[t].Preload()
        }

        res["status"] = "ok"
        res["todo"] = todo
        uadmin.ReturnJSON(w, r, res)
    }

Establish a connection in the main.go to the API by using http.HandleFunc. It should be placed after the uadmin.Register and before the StartServer.

.. code-block:: go

    func main() {
        // Some codes

        // FilterListHandler
        http.HandleFunc("/filterbuilder/", api.FilterBuilderHandler) // <-- place it here
    }

api is the folder name while FilterBuilderHandler is the name of the function inside filterbuilder.go.

Run your application and see what happens.

.. image:: assets/filterbuilderapi.png
   :align: center

**uadmin.GenerateBase32**
^^^^^^^^^^^^^^^^^^^^^^^^^
GenerateBase32 generates a base32 string of length.

Syntax:

.. code-block:: go

    func(length int) string

Go to the friend.go and initialize the Base32 field inside the struct. Set the tag as "read_only".

.. code-block:: go

    // Friend model ...
    type Friend struct {
        uadmin.Model
        Name     string `uadmin:"required"`
        Email    string `uadmin:"email"`
        Password string `uadmin:"password;list_exclude"`
        Base32   string `uadmin:"read_only"` // <-- place it here
    }

Apply overriding save function. Use this function to the Base32 field and set the integer value as 40.

.. code-block:: go

    // Save !
    func (f *Friend) Save() {
        f.Base32 = uadmin.GenerateBase32(40) // <-- place it here
        uadmin.Save(f)
    }

Now run your application. Go to the Friend model and save any element to see the changes.

.. image:: assets/friendbase32.png

|

Result

.. image:: assets/friendbase32output.png

As you notice, the Base32 value changed automatically.

**uadmin.GenerateBase64**
^^^^^^^^^^^^^^^^^^^^^^^^^
GenerateBase64 generates a base64 string of length.

Syntax:

.. code-block:: go

    func(length int) string

Go to the friend.go and initialize the Base64 field inside the struct. Set the tag as "read_only".

.. code-block:: go

    // Friend model ...
    type Friend struct {
        uadmin.Model
        Name     string `uadmin:"required"`
        Email    string `uadmin:"email"`
        Password string `uadmin:"password;list_exclude"`
        Base64   string `uadmin:"read_only"` // <-- place it here
    }

Apply overriding save function. Use this function to the Base64 field and set the integer value as 75.

.. code-block:: go

    // Save !
    func (f *Friend) Save() {
        f.Base64 = uadmin.GenerateBase64(75) // <-- place it here
        uadmin.Save(f)
    }

Now run your application. Go to the Friend model and save any element to see the changes.

.. image:: assets/friendbase64.png

|

Result

.. image:: assets/friendbase64output.png

As you notice, the Base64 value changed automatically.

**uadmin.Get**
^^^^^^^^^^^^^^
Get fetches the first record from the database.

Syntax:

.. code-block:: go

    func(a interface{}, query interface{}, args ...interface{}) (err error)

Parameters:

    **a interface{}:** Is the variable where the model name was initialized.

    **query interface{}:** Is an action that you want to perform with in your data list.

    **args ...interface{}:** Is the variable or container that can be used in execution process.

Suppose you have ten records in your Todo model.

.. image:: tutorial/assets/tendataintodomodel.png

Go to the main.go. Let's count how many todos do you have with a friend in your model.

.. code-block:: go

    func main(){
        // Some codes contained in this part

        // Initialize the Todo model in the todo variable
        todo := models.Todo{}

        // Initialize the Friend model in the todo variable
        friend := models.Friend{}

        // Fetch the first record from the database
        uadmin.Get(&friend, "id=?", todo.FriendID)

        // Return the count of records in a table based on a Get function to  
        // be stored in the total variable
        total := uadmin.Count(&todo, "friend_id = ?", todo.FriendID)

        // Print the result
        uadmin.Trail(uadmin.INFO, "You have %v todos with a friend in your list.", total)
    }

Check your terminal to see the result.

.. code-block:: bash

    [  INFO  ]   You have 5 todos with a friend in your list.

**uadmin.GetDB**
^^^^^^^^^^^^^^^^
GetDB returns a pointer to the DB.

Syntax:

.. code-block:: go

    func() *gorm.DB

Before we proceed to the example, read `Tutorial Part 7 - Introduction to API`_ to familiarize how API works in uAdmin.

.. _Tutorial Part 7 - Introduction to API: https://uadmin.readthedocs.io/en/latest/tutorial/part7.html

Suppose I have one record in the Todo model.

.. image:: assets/todomodeloutput.png

Create a file named custom_todo.go inside the api folder with the following codes below:

.. code-block:: go

    // CustomTodoHandler !
    func CustomTodoHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/custom_todo")

        res := map[string]interface{}{}

        // Initialize the Todo model
        todolist := []models.Todo{}

        // Create a query in the sql variable to select all records in todos
        sql := `SELECT * FROM todos`

        // Place it here
        db := uadmin.GetDB()

        // Store the query inside the Raw function in order to scan value to
        // the Todo model
        db.Raw(sql).Scan(&todolist)

        // Print the result in JSON format
        res["status"] = "ok"
        res["todo"] = todolist
        uadmin.ReturnJSON(w, r, res)
    }

Establish a connection in the main.go to the API by using http.HandleFunc. It should be placed after the uadmin.Register and before the StartServer.

.. code-block:: go

    func main() {
        // Some codes

        // CustomTodoHandler
        http.HandleFunc("/custom_todo/", api.CustomTodoHandler) // <-- place it here
    }

api is the folder name while CustomTodoHandler is the name of the function inside custom_todo.go.

Run your application and see what happens.

.. image:: assets/getdbjson.png

**uadmin.GetID**
^^^^^^^^^^^^^^^^
GetID returns an ID number of a field.

Syntax:

.. code-block:: go

    func(m.reflectValue) uint

Suppose I have four records in my Category model.

* Education ID = 4
* Family ID = 3
* Work ID = 2
* Travel ID = 1

.. image:: assets/categorylist.png

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){

        // Some codes

        // Initialize the Category model
        categorylist := models.Category{}

        // Get the value of the name in the categorylist
        uadmin.Get(&categorylist, "name = 'Family'")

        // Get the ID of the name "Family"
        getid := uadmin.GetID(reflect.ValueOf(categorylist))

        // Print the result
        uadmin.Trail(uadmin.INFO, "GetID is %d.", getid)
    }

Run your application and check the terminal to see the result.

.. code-block:: bash

    [  INFO  ]   GetID is 3.

**uadmin.GetString**
^^^^^^^^^^^^^^^^^^^^
GetString returns string representation on an instance of a model.

Syntax:

.. code-block:: go

    func(a interface{}) string

Parameters:

    **a interface{}:** Is the variable where the model name was initialized.

Suppose I have four records in my Category model.

* Education ID = 4
* Family ID = 3
* Work ID = 2
* Travel ID = 1

.. image:: assets/categorylist.png

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){

        // Some codes

        // Initialize the Category model
        categorylist := models.Category{}

        // Get the ID in the categorylist
        uadmin.Get(&categorylist, "id = 3")

        // Get the name of the ID 3
        getstring := uadmin.GetString(categorylist)

        // Print the result
        uadmin.Trail(uadmin.INFO, "GetString is %s.", getstring)
    }

Run your application and check the terminal to see the result.

.. code-block:: bash

    [  INFO  ]   GetString is Family.

**uadmin.GetUserFromRequest**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
GetUserFromRequest returns a user from a request.

Syntax:

.. code-block:: go

    func(r *http.Request) *uadmin.User

Before we proceed to the example, read `Tutorial Part 7 - Introduction to API`_ to familiarize how API works in uAdmin.

Suppose that the admin account has logined.

.. image:: tutorial/assets/adminhighlighted.png

|

Create a file named info.go inside the api folder with the following codes below:

.. code-block:: go

    // InfoHandler !
    func InfoHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/info")

        // Place it here
        uadmin.Trail(uadmin.INFO, "GetUserFromRequest: %s", uadmin.GetUserFromRequest(r))
    }

Establish a connection in the main.go to the API by using http.HandleFunc. It should be placed after the uadmin.Register and before the StartServer.

.. code-block:: go

    func main() {
        // Some codes

        // InfoHandler
        http.HandleFunc("/info/", api.InfoHandler) // <-- place it here
    }

api is the folder name while InfoHandler is the name of the function inside info.go.

Run your application and see what happens.

.. image:: assets/infoapi.png

Check your terminal for the result.

.. code-block:: bash

    [  INFO  ]   GetUserFromRequest: System Admin

The result is coming from the user in the dashboard.

.. image:: assets/getuserfromrequest.png

|

There is another way of using this function:

.. code-block:: go

    // InfoHandler !
    func InfoHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/info")

        getuser := uadmin.GetUserFromRequest(r)
        getuser.XXXX
    }

XXXX contains user fields and functions that you can use. See `uadmin.User`_ for the list and examples.

Go to the info.go in API folder containing the following codes below:

.. code-block:: go

    // InfoHandler !
    func InfoHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/info")

        // Get the User that returns the first and last name
        getuser := uadmin.GetUserFromRequest(r)

        // Print the result using Golang fmt
        fmt.Println("GetActiveSession() is", getuser.GetActiveSession())
        fmt.Println("GetDashboardMenu() is", getuser.GetDashboardMenu())

        // Print the result using Trail
        uadmin.Trail(uadmin.INFO, "GetOTP() is %s.", getuser.GetOTP())
        uadmin.Trail(uadmin.INFO, "String() is %s.", getuser.String())
    }

Run your application and see what happens.

.. image:: assets/infoapi.png

Check your terminal for the result.

.. code-block:: bash

    GetActiveSession() is Pfr7edaO7bBjv9zL9j1Yi01I
    GetDashboardMenu() is [Dashboard Menus Users User Groups Sessions User Permissions Group Permissions Languages Logs Todos Categorys Friends Items]
    [  INFO  ]   GetOTP() is 363669.
    [  INFO  ]   String() is System Admin.

**uadmin.GroupPermission**
^^^^^^^^^^^^^^^^^^^^^^^^^^
GroupPermission sets the permission of a user group handled by an administrator.

Syntax:

.. code-block:: go

    type GroupPermission struct {
        Model
        DashboardMenu   DashboardMenu `gorm:"ForeignKey:DashboardMenuID" required:"true" filter:"true"`
        DashboardMenuID uint          `fk:"true" displayName:"DashboardMenu"`
        UserGroup       UserGroup     `gorm:"ForeignKey:UserGroupID" required:"true" filter:"true"`
        UserGroupID     uint          `fk:"true" displayName:"UserGroup"`
        Read            bool
        Add             bool
        Edit            bool
        Delete          bool
    }

There are 2 functions that you can use in GroupPermission:

* **HideInDashboard()** - Return true and auto hide this from dashboard
* **String()** - Returns the GroupPermission ID

There are 2 ways you can do for initialization process using this function: one-by-one and by group.

One-by-one initialization:

.. code-block:: go

    func main(){
        // Some codes
        grouppermission := uadmin.GroupPermission{}
        grouppermission.DashboardMenu = dashboardmenu
        grouppermission.DashboardMenuID = 1
        grouppermission.UserGroup = usergroup
        grouppermission.UserGroupID = 1
    }

By group initialization:

.. code-block:: go

    func main(){
        // Some codes
        grouppermission := uadmin.GroupPermission{
            DashboardMenu: dashboardmenu,
            DashboardMenuID: 1,
            UserGroup: usergroup,
            UserGroupID: 1,
        }
    }

In this example, we will use "by group" initialization process.

Suppose that Even Demata is a part of the Front Desk group.

.. image:: assets/useraccountfrontdesk.png

|

Go to the main.go and apply the following codes below after the RegisterInlines section.

.. code-block:: go

    func main(){

        // Some codes

        grouppermission := uadmin.GroupPermission{
            DashboardMenuID: 9, // Todos
            UserGroupID:     1, // Front Desk
            Read:            true,
            Add:             false,
            Edit:            false,
            Delete:          false,
        }

        // This will create a new group permission based on the information
        // assigned in the grouppermission variable.
        uadmin.Save(&grouppermission)

        // Returns the GroupPermissionID
        uadmin.Trail(uadmin.INFO, "String() returns %s.", grouppermission.String())
    }

Now run your application and see what happens.

**Terminal**

.. code-block:: bash

    [  INFO  ]   String() returns 1.

.. image:: assets/grouppermissioncreated.png

|

Log out your System Admin account. This time login your username and password using the user account that has group permission. Afterwards, you will see that only the Todos model is shown in the dashboard because your user account is not an admin and has no remote access to it. Now click on TODOS model.

.. image:: assets/userpermissiondashboard.png

|

As you will see, your user account is restricted to add, edit, or delete a record in the Todo model. You can only read what is inside this model.

.. image:: assets/useraddeditdeleterestricted.png

|

If you want to hide the Todo model in your dashboard, first of all, create a HideInDashboard() function in your todo.go inside the models folder and set the return value to "true".

.. code-block:: go

    // HideInDashboard !
    func (t Todo) HideInDashboard() bool {
        return true
    }

Now you can do something like this in main.go:

.. code-block:: go

    func main(){

        // Some codes

        // Initializes the DashboardMenu
        dashboardmenu := uadmin.DashboardMenu{}

        // Assign the grouppermission, call the HideInDashboard() function
        // from todo.go, store it to the Hidden field of the dashboardmenu
        dashboardmenu.Hidden = grouppermission.HideInDashboard()

        // Checks the Dashboard Menu ID number from the grouppermission. If it
        // matches, it will update the value of the Hidden field.
        uadmin.Update(&dashboardmenu, "Hidden", dashboardmenu.Hidden, "id = ?", grouppermission.DashboardMenuID)
    }

Now rerun your application using the Even Demata account and see what happens.

.. image:: assets/dashboardmenuempty.png

|

The Todo model is now hidden from the dashboard. If you login your System Admin account, you will see in the Dashboard menu that the hidden field of the Todo model is set to true.

.. image:: assets/todomodelhidden.png

**uadmin.HideInDashboarder**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^
HideInDashboarder is used to check if a model should be hidden in the dashboard.

Syntax:

.. code-block:: go

    type HideInDashboarder interface{
        HideInDashboard() bool
    }

Suppose I have five models in my dashboard: Todos, Categorys, Items, Friends, and Expressions. I want Friends and Expressions models to be hidden in the dashboard. In order to do that, go to the friend.go and expression.go inside the models folder and apply the HideInDashboard() function. Set the return value to **true** inside it.

**friend.go**

.. code-block:: go

    func (f Friend) HideInDashboard() bool {
        return true
    }

**expression.go**

.. code-block:: go

    func (e Expression) HideInDashboard() bool {
        return true
    }

Now go to the main.go and apply the following codes below inside the main function:

.. code-block:: go

    // Initialize the Expression and Friend models inside the modelList with
    // the array type of interface
    modelList := []interface{}{
        models.Expression{},
        models.Friend{},
    }
    
    // Loop the execution process based on the modelList count
    for i := range modelList {

        // Returns the reflection type that represents the dynamic type of i
        t := reflect.TypeOf(modelList[i])

        // Calls the HideInDashboarder function to access the HideInDashboard()
        hideItem := modelList[i].(uadmin.HideInDashboarder).HideInDashboard()

        // Initializes the hidethismodel variable to assign the DashboardMenu
        hidethismodel := uadmin.DashboardMenu{

            // Returns the name of the model based on reflection
            MenuName: strings.Join(helper.SplitCamelCase(t.Name()), " "),

            // Returns the boolean value based on the assigned return in the
            // HideInDashboard()
            Hidden:   hideItem,
        }

        // Prints the information of the hidethismodel
        uadmin.Trail(uadmin.INFO, "MenuName: %s,  Hidden: %t", hidethismodel.MenuName, hidethismodel.Hidden)
    }

Go back to your application. Open the DashboardMenu then delete the Expressions and Friends model.

.. image:: assets/deletetwomodels.png

|

Now rerun your application and see what happens.

.. code-block:: bash

    [  INFO  ]   MenuName: Expression,  Hidden: true
    [  INFO  ]   MenuName: Friend,  Hidden: true

.. image:: assets/twomodelshidden.png

|

As expected, Friends and Expressions models are now hidden in the dashboard. If you go to the Dashboard Menus, you will see that they are checked in the Hidden field.

.. image:: assets/twomodelshiddenchecked.png

**uadmin.INFO**
^^^^^^^^^^^^^^^
INFO is the display tag under Trail. It is a data that is presented within a context that gives it meaning and relevance.

Syntax:

.. code-block:: go

    untyped int

See `uadmin.Trail`_ for the example.

**uadmin.IsAuthenticated**
^^^^^^^^^^^^^^^^^^^^^^^^^^
IsAuthenticated returns the session of the user.

Syntax:

.. code-block:: go

    func(r *http.Request) *uadmin.Session

See `uadmin.Session`_ for the list of fields and functions that you can use in IsAuthenticated.

Before we proceed to the example, read `Tutorial Part 7 - Introduction to API`_ to familiarize how API works in uAdmin.

Suppose that the admin account has logined.

.. image:: tutorial/assets/adminhighlighted.png

|

Create a file named custom_todo.go inside the api folder with the following codes below:

.. code-block:: go

    // CustomTodoHandler !
    func CustomTodoHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/custom_todo")

        // Get the session or key
        session := uadmin.IsAuthenticated(r)

        // If there is no value in the session, it will return the
        // LoginHandler.
        if session == nil {
            LoginHandler(w, r)
            return
        }

        // Fetch the values from a User model using session IsAuthenticated
        user := session.User
        userid := session.UserID
        username := session.User.Username
        active := session.User.Active

        // Print the result
        uadmin.Trail(uadmin.INFO, "Session / Key: %s", session)
        uadmin.Trail(uadmin.INFO, "User: %s", user)
        uadmin.Trail(uadmin.INFO, "UserID: %d", userid)
        uadmin.Trail(uadmin.INFO, "Username: %s", username)
        uadmin.Trail(uadmin.INFO, "Active: %v", active)
    }

Establish a connection in the main.go to the API by using http.HandleFunc. It should be placed after the uadmin.Register and before the StartServer.

.. code-block:: go

    func main() {
        // Some codes

        // CustomTodoHandler
        http.HandleFunc("/custom_todo/", api.CustomTodoHandler) // <-- place it here
    }

api is the folder name while CustomTodoHandler is the name of the function inside custom_todo.go.

Run your application and see what happens.

.. image:: assets/customtodoapi.png

Check your terminal for the result.

.. code-block:: bash

    [  INFO  ]   Session / Key: Pfr7edaO7bBjv9zL9j1Yi01I
    [  INFO  ]   Username: System Admin
    [  INFO  ]   UserID: 1
    [  INFO  ]   Username: admin
    [  INFO  ]   Active: true

The result is coming from the session in the dashboard.

.. image:: assets/isauthenticated.png

|

And the values in the User model by calling the User, UserID, Username, and Active fields.

.. image:: assets/usersession.png

**uadmin.JSONMarshal**
^^^^^^^^^^^^^^^^^^^^^^
JSONMarshal returns the JSON encoding of v.

Syntax:

.. code-block:: go

    func(v interface{}, safeEncoding bool) ([]byte, error)

Before we proceed to the example, read `Tutorial Part 7 - Introduction to API`_ to familiarize how API works in uAdmin.

.. _Tutorial Part 7 - Introduction to API: https://uadmin.readthedocs.io/en/latest/tutorial/part7.html

Create a file named friend_list.go inside the api folder with the following codes below:

.. code-block:: go

    // FriendListHandler !
    func FriendListHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/friend_list")

        res := map[string]interface{}{}

        filterList := []string{}
        valueList := []interface{}{}
        if r.URL.Query().Get("friend_id") != "" {
            filterList = append(filterList, "friend_id = ?")
            valueList = append(valueList, r.URL.Query().Get("friend_id"))
        }
        filter := strings.Join(filterList, " AND ")

        // Fetch Data from DB
        friend := []models.Friend{}
        uadmin.Filter(&friend, filter, valueList...)

        // Place it here
        output, err := uadmin.JSONMarshal(&friend, true)
        if err != nil {
            log.Fatal(output)
        }

        // Prints the output to the terminal in JSON format
        os.Stdout.Write(output)

        // Unmarshal parses the JSON-encoded data and stores the result in the
        // value pointed to by v.
        json.Unmarshal(output, &friend)

        // Prints the JSON format in the API webpage
        res["status"] = "ok"
        res["todo"] = friend
        uadmin.ReturnJSON(w, r, res)
    }

Establish a connection in the main.go to the API by using http.HandleFunc. It should be placed after the uadmin.Register and before the StartServer.

.. code-block:: go

    func main() {
        // Some codes

        // FilterListHandler
        http.HandleFunc("/friend_list/", api.FriendListHandler) // <-- place it here
    }

api is the folder name while FriendListHandler is the name of the function inside friend_list.go.

Run your application and see what happens.

**Terminal**

.. code-block:: bash

    [
        {
        "ID": 1,
        "DeletedAt": null,
        "Name": "Even Demata",
        "Email": "test@gmail.com",
        "Password": "$2a$12$p3yNEVq9JR4W4ac6x7JM0u1c6rQq7w10ID7Y9yjKLWFd9wbp2PMLq",
        }
    ]

**API**

.. image:: assets/friendlistjsonmarshal.png
   :align: center

**uadmin.Language**
^^^^^^^^^^^^^^^^^^^
Language is a system in uAdmin that is used to add, modify and delete the elements of a language.

Syntax:

.. code-block:: go

    type Language struct {
        Model
        EnglishName    string `uadmin:"required;read_only;filter;search"`
        Name           string `uadmin:"required;read_only;filter;search"`
        Flag           string `uadmin:"image;list_exclude"`
        Code           string `uadmin:"filter;read_only;list_exclude"`
        RTL            bool   `uadmin:"list_exclude"`
        Default        bool   `uadmin:"help:Set as the default language;list_exclude"`
        Active         bool   `uadmin:"help:To show this in available languages;filter"`
        AvailableInGui bool   `uadmin:"help:The App is available in this language;read_only"`
    }

There are 2 functions that you can use in Language:

* **Save()** - Saves the object in the database
* **String()** - Returns the Code of the language

There are 2 ways you can do for initialization process using this function: one-by-one and by group.

One-by-one initialization:

.. code-block:: go

    func main(){
        // Some codes
        language := uadmin.Language{}
        language.EnglishName = "English Name"
        language.Name = "Name"
    }

By group initialization:

.. code-block:: go

    func main(){
        // Some codes
        language := uadmin.Language{
            EnglishName: "English Name",
            Name: "Name",
        }
    }

In this example, we will use "by group" initialization process.

Suppose the Tagalog language is not active and you want to set this to Active.

.. image:: assets/tagalognotactive.png

|

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){

        // Some codes

        // Language configurations
        language := uadmin.Language{
            EnglishName:    "Tagalog",
            Name:           "Wikang Tagalog",
            Flag:           "",
            Code:           "tl",
            RTL:            false,
            Default:        false,
            Active:         false,
            AvailableInGui: false,
        }

        // Checks the English name from the language. If it matches, it will
        // update the value of the Active field.
        uadmin.Update(&language, "Active", true, "english_name = ?", language.EnglishName)

        // Returns the Code of the language
        uadmin.Trail(uadmin.INFO, "String() returns %s.", language.String())
    }

Now run your application, refresh your browser and see what happens.

**Terminal**

.. code-block:: bash

    [  INFO  ]   String() returns tl.

.. image:: assets/tagalogactive.png

|

As expected, the Tagalog language is now set to active.

**uadmin.Log**
^^^^^^^^^^^^^^
Log is a system in uAdmin that is used to add, modify, and delete the status of the user activities.

Syntax:

.. code-block:: go

    type Log struct {
        Model
        Username  string    `uadmin:"filter;read_only"`
        Action    Action    `uadmin:"filter;read_only"`
        TableName string    `uadmin:"filter;read_only"`
        TableID   int       `uadmin:"filter;read_only"`
        Activity  string    `uadmin:"code;read_only" gorm:"type:longtext"`
        RollBack  string    `uadmin:"link;"`
        CreatedAt time.Time `uadmin:"filter;read_only"`
    }

There are 5 functions that you can use in Log:

**ParseRecord** - Uses this syntax as shown below:

.. code-block:: go

    func(a reflect.Value, modelName string, ID uint, user *User, action Action, r *http.Request) (err error)

**PasswordReset** - Uses this syntax as shown below:

.. code-block:: go

    func(user string, action Action, r *http.Request) (err error)

**Save()** - Saves the object in the database

**SignIn** - Uses this syntax as shown below:

.. code-block:: go

    func(user string, action Action, r *http.Request) (err error)

**String()** - Returns the Log ID

Go to the main.go and apply the following codes below after the RegisterInlines section.

.. code-block:: go

    func main(){

        // Some codes

        log := uadmin.Log{
            Username:  "admin",
            Action:    uadmin.Action.Custom(0),
            TableName: "Todo",
            TableID:   1,
            Activity:  "Custom Add from the source code",
            RollBack:  "",
            CreatedAt: time.Now(),
        }

        // This will create a new log based on the information assigned in
        // the log variable.
        log.Save()

        // Returns the Log ID
        uadmin.Trail(uadmin.INFO, "String() returns %s.", log.String())
    }

Now run your application and see what happens.

**Terminal**

.. code-block:: bash

    [  INFO  ]   String() returns 1.

.. image:: assets/logcreated.png

**uadmin.Login**
^^^^^^^^^^^^^^^^
Login returns the pointer of User and a bool for Is OTP Required.

Syntax:

.. code-block:: go

    func(r *http.Request, username string, password string) (*uadmin.User, bool)

Before we proceed to the example, read `Tutorial Part 7 - Introduction to API`_ to familiarize how API works in uAdmin.

Create a file named info.go inside the api folder with the following codes below:

.. code-block:: go

    // InfoHandler !
    func InfoHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/info")
        fmt.Println(uadmin.Login(r, "admin", "admin")) // <-- place it here
    }

Establish a connection in the main.go to the API by using http.HandleFunc. It should be placed after the uadmin.Register and before the StartServer.

.. code-block:: go

    func main() {
        // Some codes

        // InfoHandler
        http.HandleFunc("/info/", api.InfoHandler) // <-- place it here
    }

api is the folder name while InfoHandler is the name of the function inside info.go.

Run your application and see what happens.

.. image:: assets/infoapi.png

Check your terminal for the result.

.. code-block:: bash

    System Admin false

The result is coming from the user in the dashboard.

.. image:: assets/systemadminotphighlighted.png

**uadmin.Login2FA**
^^^^^^^^^^^^^^^^^^^
Login2FA returns the pointer of User with a two-factor authentication.

Syntax:

.. code-block:: go

   func(r *http.Request, username string, password string, otpPass string) *uadmin.User

Before we proceed to the example, read `Tutorial Part 7 - Introduction to API`_ to familiarize how API works in uAdmin.

First of all, activate the OTP Required in your System Admin account.

.. image:: assets/otprequired.png

|

Afterwards, logout your account then login again to get the OTP verification code in your terminal.

.. image:: assets/loginformwithotp.png

.. code-block:: bash

    [  INFO  ]   User: admin OTP: 445215

Now create a file named info.go inside the api folder with the following codes below:

.. code-block:: go

    package api

    import (
        "fmt"
        "net/http"
        "strings"

        "github.com/uadmin/uadmin"
    )

    // InfoHandler !
    func InfoHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/info")

        // Place it here
        fmt.Println(uadmin.Login2FA(r, "admin", "admin", "445215"))
    }

Establish a connection in the main.go to the API by using http.HandleFunc. It should be placed after the uadmin.Register and before the StartServer.

.. code-block:: go

    func main() {
        // Some codes

        // InfoHandler
        http.HandleFunc("/info/", api.InfoHandler) // <-- place it here
    }

api is the folder name while InfoHandler is the name of the function inside info.go.

Run your application and see what happens.

.. image:: assets/infoapi.png

|

Check your terminal for the result.

.. code-block:: bash

    System Admin

**uadmin.Logout**
^^^^^^^^^^^^^^^^^
Logout deactivates a session.

Syntax:

.. code-block:: go

    func(r *http.Request)

Suppose that the admin account has logined.

.. image:: tutorial/assets/adminhighlighted.png

|

Create a file named logout.go inside the api folder with the following codes below:

.. code-block:: go

    // LogoutHandler !
    func LogoutHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/logout")
        uadmin.Logout(r) // <-- place it here
    }

Establish a connection in the main.go to the API by using http.HandleFunc. It should be placed after the uadmin.Register and before the StartServer.

.. code-block:: go

    func main() {
        // Some codes

        // LogoutHandler
        http.HandleFunc("/logout/", api.LogoutHandler)) // <-- place it here
    }

api is the folder name while LogoutHandler is the name of the function inside logout.go.

Run your application and see what happens.

.. image:: assets/logoutapi.png

Refresh your browser and see what happens.

.. image:: tutorial/assets/loginform.png

|

Your account has been logged out automatically that redirects you to the login form.

**uadmin.MaxImageHeight**
^^^^^^^^^^^^^^^^^^^^^^^^^
MaxImageHeight sets the maximum height of an image.

Syntax:

.. code-block:: go

    int

See `uadmin.MaxImageWidth`_ for the example.

**uadmin.MaxImageWidth**
^^^^^^^^^^^^^^^^^^^^^^^^
MaxImageWidth sets the maximum width of an image.

Syntax:

.. code-block:: go

    int

Let's set the MaxImageWidth to 360 pixels and the MaxImageHeight to 240 pixels.

.. code-block:: go

    func main() {
        // Some codes
        uadmin.MaxImageWidth = 360      // <--  place it here
        uadmin.MaxImageHeight = 240     // <--  place it here
    }

uAdmin has a feature that allows you to customize your own profile. In order to do that, click the profile icon on the top right corner then select admin as highlighted below.

.. image:: tutorial/assets/adminhighlighted.png

|

By default, there is no profile photo inserted on the top left corner. If you want to add it in your profile, click the Choose File button to browse the image on your computer.

.. image:: tutorial/assets/choosefilephotohighlighted.png

|

Let's pick a photo that surpasses the MaxImageWidth and MaxImageHeight values.

.. image:: tutorial/assets/widthheightbackground.png
   :align: center

|

Once you are done, click Save Changes on the left corner and refresh the webpage to see the output.

.. image:: tutorial/assets/profilepicadded.png

As expected, the profile pic will be uploaded to the user profile that automatically resizes to 360x240 pixels.

**uadmin.MaxUploadFileSize**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^
MaxUploadFileSize is the maximum upload file size in bytes.

Syntax:

.. code-block:: go

    int64

Go to the main.go. Let's set the MaxUploadFileSize value to 1024. 1024 is equivalent to 1 MB.

.. code-block:: go

    func main() {
        // Some codes
        uadmin.MaxUploadFileSize = 1024     // <--  place it here
    }

Run the application, go to your profile and upload an image that exceeds the MaxUploadFileSize limit. If you click Save changes...

.. image:: tutorial/assets/noprofilepic.png

The profile picture has failed to upload in the user profile because the file size is larger than the limit.

**uadmin.Model**
^^^^^^^^^^^^^^^^
Model is the standard struct to be embedded in any other struct to make it a model for uAdmin.

Syntax:

.. code-block:: go

    type Model struct {
        ID        uint       `gorm:"primary_key"`
        DeletedAt *time.Time `sql:"index"`
    }

In every struct, uadmin.Model must always come first before creating a field.

.. code-block:: go

    type (struct_name) struct{
        uadmin.Model // <-- place it here
        // Some codes here
    }

**uadmin.ModelSchema**
^^^^^^^^^^^^^^^^^^^^^^
ModelSchema is a representation of a plan or theory in the form of an outline or model.

Syntax:

.. code-block:: go

    type ModelSchema struct {
        Name          string // Name of the Model
        DisplayName   string // Display Name of the model
        ModelName     string // URL
        ModelID       uint
        Inlines       []*ModelSchema
        InlinesData   []listData
        Fields        []F
        IncludeFormJS []string
        IncludeListJS []string
    }

There is a function that you can use in ModelSchema:

* **FieldByName** - Uses this syntax as shown below:

.. code-block:: go

    func(a string) *uadmin.F

Syntax:

.. code-block:: go

    modelschema.FieldByName("Name").XXXX = Value

XXXX has many things: See `uadmin.F`_ syntax for the list. It is an alternative way of changing the feature of the field rather than using Tags. For more information, see `Tag Reference`_.

.. _Tag Reference: https://uadmin.readthedocs.io/en/latest/tags.html

There are 2 ways you can do for initialization process using this function: one-by-one and by group.

One-by-one initialization:

.. code-block:: go

    func main(){
        // Some codes
        modelschema := uadmin.ModelSchema{}
        modelschema.Name = "Name"
        modelschema.DisplayName = "Display Name"
    }

By group initialization:

.. code-block:: go

    func main(){
        // Some codes
        modelschema := uadmin.ModelSchema{
            Name: "Name",
            DisplayName: "Display Name",
        }
    }

In this example, we will use "by group" initialization process.

Before you proceed to this example, see `uadmin.F`_.

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Some codes
        // uadmin.F codes here
        modelschema := uadmin.ModelSchema{
            Name:        "Expressions",
            DisplayName: "What's on your mind?",
            ModelName:   "expression",
            ModelID:     13,

            // f1 and f2 are initialized variables in uadmin.F
            Fields:      []uadmin.F{f1, f2},
        }
    }

The code above shows an initialized modelschema struct using the Name, DisplayName, ModelName, ModelID, and Fields.

See `uadmin.Schema`_ for the continuation of this example.

**uadmin.MongoDB (Experimental)**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
MongoDB is an open source database management system (DBMS) that uses a document-oriented database model which supports various forms of data. [#f1]_ It is the active Mongo settings.

Syntax:

.. code-block:: go

    *uadmin.MongoSettings

There are 3 fields that you can use in MongoDB:

* **Debug** - returns a boolean value
* **IP** - returns a string
* **Name** - returns a string

**uadmin.MongoModel (Experimental)**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
MongoModel is a uAdmin function for interfacing with MongoDB databases.

Syntax:

.. code-block:: go

    type MongoModel struct {
	    ID bson.ObjectId `bson:"_id,omitempty"`
    }

There are 8 functions that you can use in MongoModel:

**All** - Fetches all objects in the database. It uses this syntax as shown below:

.. code-block:: go

    func(a interface{}, ColNameExtra string) error

**Count** - Return the count of records in a table based on a filter. It uses this syntax as shown below:

.. code-block:: go

    func(filter interface{}, a interface{}, ColNameExtra string) int

**Delete** - Delete records from the database. It uses this syntax as shown below:

.. code-block:: go

    func(a interface{}, ColNameExtra string) error

**Filter** - Fetches records from the database. It uses this syntax as shown below:

.. code-block:: go

    func(filter interface{}, a interface{}, ColNameExtra string) error

**Get** - Fetches the first record from the database. It uses this syntax as shown below:

.. code-block:: go

    func(filter interface{}, a interface{}, ColNameExtra string) error

**GetCol** - Uses this syntax as shown below:

.. code-block:: go

    func(a interface{}, ColNameExtra string) (*mgo.Collection, error)

**Query** - Uses this syntax as shown below:

.. code-block:: go

    func(filter interface{}, a interface{}, ColNameExtra string) *mgo.Query

**Save** - Saves the object in the database. It uses this syntax as shown below:

.. code-block:: go

    func(a interface{}, ColNameExtra string)

**uadmin.MongoSettings (Experimental)**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
MongoSettings is a feature that allows a user to configure the settings of a Mongo.

Syntax:

.. code-block:: go

    type MongoSettings struct {
        Name  string
        IP    string
        Debug bool
    }

**uadmin.NewModel**
^^^^^^^^^^^^^^^^^^^
NewModel creates a new model from a model name.

Syntax:

.. code-block:: go

    func(modelName string, pointer bool) (reflect.Value, bool)

Suppose I have four records in my Category model.

* Education ID = 4
* Family ID = 3
* Work ID = 2
* Travel ID = 1

.. image:: assets/categorylist.png

Create a file named custom_todo.go inside the api folder with the following codes below:

.. code-block:: go

    // CustomTodoHandler !
    func CustomTodoHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/custom_todo")

        res := map[string]interface{}{}

        // Call the category model and set the pointer to true
        m, _ := uadmin.NewModel("category", true)

        // Fetch the records of the category model
        uadmin.Get(m.Interface(), "id = ?", 3)

        // Assign the m.Interface() to the newmode
        newmodel := m.Interface()

        // Print the result in JSON format
        res["status"] = "ok"
        res["category"] = newmodel
        uadmin.ReturnJSON(w, r, res)
    }

Establish a connection in the main.go to the API by using http.HandleFunc. It should be placed after the uadmin.Register and before the StartServer.

.. code-block:: go

    func main() {
        // Some codes

        // CustomTodoHandler
        http.HandleFunc("/custom_todo/", api.CustomTodoHandler) // <-- place it here
    }

api is the folder name while CustomTodoHandler is the name of the function inside custom_todo.go.

Run your application and see what happens.

.. image:: assets/newmodeljson.png

**uadmin.NewModelArray**
^^^^^^^^^^^^^^^^^^^^^^^^
NewModelArray creates a new model array from a model name.

Syntax:

.. code-block:: go

    func(modelName string, pointer bool) (reflect.Value, bool)

Suppose I have four records in my Category model.

.. image:: assets/categorylist.png

Create a file named custom_todo.go inside the api folder with the following codes below:

.. code-block:: go

    // CustomTodoHandler !
    func CustomTodoHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/custom_todo")

        res := map[string]interface{}{}

        // Call the category model and set the pointer to true
        m, _ := uadmin.NewModelArray("category", true)

        // Fetch the records of the category model
        uadmin.Filter(m.Interface(), "id >= ?", 1)

        // Assign the m.Interface() to the newmodelarray
        newmodelarray := m.Interface()

        // Print the result in JSON format
        res["status"] = "ok"
        res["category"] = newmodelarray
        uadmin.ReturnJSON(w, r, res)
    }

Establish a connection in the main.go to the API by using http.HandleFunc. It should be placed after the uadmin.Register and before the StartServer.

.. code-block:: go

    func main() {
        // Some codes

        // CustomTodoHandler
        http.HandleFunc("/custom_todo/", api.CustomTodoHandler) // <-- place it here
    }

api is the folder name while CustomTodoHandler is the name of the function inside custom_todo.go.

Run your application and see what happens.

.. image:: assets/newmodelarrayjson.png

**uadmin.OK**
^^^^^^^^^^^^^
OK is the display tag under Trail. It is a status to show that the application is doing well.

Syntax:

.. code-block:: go

    untyped int

See `uadmin.Trail`_ for the example.

**uadmin.OTPAlgorithm**
^^^^^^^^^^^^^^^^^^^^^^^
OTPAlgorithm is the hashing algorithm of OTP.

Syntax:

.. code-block:: go

    string

There are 3 different algorithms:

* sha1 (default)
* sha256
* sha512

**uadmin.OTPDigits**
^^^^^^^^^^^^^^^^^^^^
OTPDigits is the number of digits for the OTP.

Syntax:

.. code-block:: go

    int

Go to the main.go and set the OTPDigits to 8.

.. code-block:: go

    func main() {
        // Some codes
        uadmin.OTPDigits = 8 // <--  place it here
    }

Run your application, login your account, and check your terminal afterwards to see the OTP verification code assigned by your system.

.. code-block:: bash

    [  INFO  ]   User: admin OTP: 90401068

As shown above, it has 8 OTP digits.

**uadmin.OTPPeriod**
^^^^^^^^^^^^^^^^^^^^
OTPPeriod is the number of seconds for the OTP to change.

Syntax:

.. code-block:: go

    uint

Go to the main.go and set the OTPPeriod to 10 seconds.

.. code-block:: go

    func main() {
        // Some codes
        uadmin.OTPPeriod = uint(10) // <--  place it here
    }

Run your application, login your account, and check your terminal afterwards to see how the OTP code changes every 10 seconds by refreshing your browser.

.. code-block:: bash

    // Before refreshing your browser
    [  INFO  ]   User: admin OTP: 433452

    // After refreshing your browser in more than 10 seconds
    [  INFO  ]   User: admin OTP: 185157

**uadmin.OTPSkew**
^^^^^^^^^^^^^^^^^^
OTPSkew is the number of minutes to search around the OTP.

Syntax:

.. code-block:: go

    uint

Go to the main.go and set the OTPSkew to 2 minutes.

.. code-block:: go

    func main() {
        // Some codes
        uadmin.OTPSkew = uint(2) // <--  place it here
    }

Run your application, login your account, and check your terminal afterwards to see the OTP verification code assigned by your system. Wait for more than two minutes and check if the OTP code is still valid.

After waiting for more than two minutes,

.. image:: assets/loginformwithotp.png

It redirects to the same webpage which means your OTP code is no longer valid.

**uadmin.PageLength**
^^^^^^^^^^^^^^^^^^^^^
PageLength is the list view max number of records.

Syntax:

.. code-block:: go

    int

Go to the main.go and apply the PageLength function.

.. code-block:: go

    func main() {
        // Some codes
        uadmin.PageLength = 4  // <--  place it here
    }

Run your application, go to the Item model, inside it you have 6 total elements. The elements in the item model will display 4 elements per page.

.. image:: tutorial/assets/pagelength.png

**uadmin.Port**
^^^^^^^^^^^^^^^
Port is the port used for http or https server.

Syntax:

.. code-block:: go

    int

Go to the main.go in your Todo list project and apply **8000** as a port number.

.. code-block:: go

    func main() {
        // Some codes
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

**uadmin.Preload**
^^^^^^^^^^^^^^^^^^
Preload accesses the information of the fields in another model.

Syntax:

.. code-block:: go

    func(a interface{}, preload ...string) (err error)

Go to the friend.go and add the Points field inside the struct.

.. code-block:: go

    // Friend model ...
    type Friend struct {
        uadmin.Model
        Name     string `uadmin:"required"`
        Email    string `uadmin:"email"`
        Password string `uadmin:"password;list_exclude"`
        TotalPoints int // <-- place it here
    }

Now go to the todo.go and apply some business logic that will get the total points of each friend in the todo list. Let's apply overriding save function and put it below the Todo struct.

.. code-block:: go

    // Save ...
    func (t *Todo) Save() {
        // Save the model to DB
        uadmin.Save(t)

        // Get a list of other todo items that share the same
        // FriendID. Notice that in the filter we use friend_id which
        // is the way this is created in the DB
        todoList := []Todo{}
        uadmin.Filter(&todoList, "friend_id = ?", t.FriendID)
        progressSum := 0

        // Sum up the progress of all todos
        for _, todo := range todoList {
            progressSum += todo.Progress
        }

        // Preload the todo model to get the related points
        uadmin.Preload(t) // <-- place it here

        // Calculate the total progress
        t.Friend.TotalPoints = progressSum

        // Finally save the Friend
        uadmin.Save(&t.Friend)
    }

Suppose you have ten records in your Todo model.

.. image:: tutorial/assets/tendataintodomodel.png

|

Now go to the Friend model and see what happens.

.. image:: assets/friendpoints.png

|

In my list, Willie Revillame wins 85 points and Even Demata wins 130 points.

**uadmin.PublicMedia**
^^^^^^^^^^^^^^^^^^^^^^
PublicMedia allows public access to media handler without authentication.

Syntax:

.. code-block:: go

    bool

For instance, my account was not signed in.

.. image:: tutorial/assets/loginform.png

|

And you want to access **travel.png** inside your media folder.

.. image:: assets/mediapath.png

|

Go to the main.go and apply this function as "true". Put it above the uadmin.Register.

.. code-block:: go

    func main() {
        uadmin.PublicMedia = true // <-- place it here
        uadmin.Register(
            // Some codes
        )
    }

Result

.. image:: assets/publicmediaimage.png

**uadmin.Register**
^^^^^^^^^^^^^^^^^^^
Register is used to register models to uAdmin.

Syntax:

.. code-block:: go

    func(m ...interface{})

Create an internal Todo model inside the main.go. Afterwards, call the Todo{} inside the uadmin.Register so that the application will identify the Todo model to be added in the dashboard.

.. code-block:: go

    // Todo model ...
    type Todo struct {
        uadmin.Model
        Name        string
        Description string `uadmin:"html"`
        TargetDate  time.Time
        Progress    int `uadmin:"progress_bar"`
    }

    func main() {
        uadmin.Register(Todo{}) // <-- place it here
    }

Output

.. image:: assets/uadmindashboard.png

If you click the Todos model, it will display this result as shown below.

.. image:: assets/todomodel.png

**uadmin.RegisterInlines**
^^^^^^^^^^^^^^^^^^^^^^^^^^
RegisterInlines is a function to register a model as an inline for another model

Syntax:

.. code-block:: go

    func RegisterInlines(model interface{}, fk map[string]string)

Parameters:

    **model (struct instance):** Is the model that you want to add inlines to.

    **fk (map[interface{}]string):** This is a map of the inlines to be added to the model. The map's key is the name of the model of the inline and the value of the map is the foreign key field's name.

Example:

.. code-block:: go

    type Person struct {
        uadmin.Model
        Name string
    }

    type Card struct {
        uadmin.Model
        PersonID uint
        Person   Person
    }

    func main() {
        // ...
        uadmin.RegisterInlines(Person{}, map[string]string{
            "Card": "PersonID",
        })
        // ...
    }

**uadmin.ReportingLevel**
^^^^^^^^^^^^^^^^^^^^^^^^^
ReportingLevel is the standard reporting level.

Syntax:

.. code-block:: go

    int

There are 6 different levels:

* DEBUG
* WORKING
* INFO
* OK
* WARNING
* ERROR

Let's set the ReportingLevel to 1 to show that the debugging process is working.

.. code-block:: go

    func main() {
        // Some codes
        uadmin.ReportingLevel = 1 // <--  place it here
    }

Result

.. code-block:: bash

    [   OK   ]   Initializing DB: [12/12]
    [   OK   ]   Server Started: http://0.0.0.0:8080
             ___       __          _
      __  __/   | ____/ /___ ___  (_)___
     / / / / /| |/ __  / __  __ \/ / __ \
    / /_/ / ___ / /_/ / / / / / / / / / /
    \__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/

What if I set the value to 5?

.. code-block:: go

    func main() {
        // Some codes
        uadmin.ReportingLevel = 5 // <--  place it here
    }

Result

.. code-block:: bash

    [   OK   ]   Initializing DB: [12/12]
             ___       __          _
      __  __/   | ____/ /___ ___  (_)___
     / / / / /| |/ __  / __  __ \/ / __ \
    / /_/ / ___ / /_/ / / / / / / / / / /
    \__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/

The database was initialized. However, the server did not start because the status of the ReportingLevel is ERROR.

**uadmin.ReportTimeStamp**
^^^^^^^^^^^^^^^^^^^^^^^^^^
ReportTimeStamp set this to true to have a time stamp in your logs.

Syntax:

.. code-block:: go

    bool

Go to the main.go and set the ReportTimeStamp value as true.

.. code-block:: go

    func main() {
        // Some codes
        uadmin.ReportTimeStamp = true // <--  place it here
    }

If you run your code,

.. code-block:: bash

    [   OK   ]   Initializing DB: [12/12]
    2018/11/07 08:52:14 [   OK   ]   Server Started: http://0.0.0.0:8080
             ___       __          _
      __  __/   | ____/ /___ ___  (_)___
     / / / / /| |/ __  / __  __ \/ / __ \
    / /_/ / ___ / /_/ / / / / / / / / / /
    \__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/

**uadmin.ReturnJSON**
^^^^^^^^^^^^^^^^^^^^^
ReturnJSON returns JSON to the client.

Syntax:

.. code-block:: go

    func(w http.ResponseWriter, r *http.Request, v interface{})

See `Tutorial Part 7 - Introduction to API`_ for the example.

.. _Tutorial Part 7 - Introduction to API: https://uadmin.readthedocs.io/en/latest/tutorial/part7.html

**uadmin.RootURL**
^^^^^^^^^^^^^^^^^^
RootURL is where the listener is mapped to.

Syntax:

.. code-block:: go

    string

Go to the main.go and apply this function as "/admin/". Put it above the uadmin.Register.

.. code-block:: go

    func main() {
        uadmin.RootURL = "/admin/" // <-- place it here
        uadmin.Register(
            // Some codes
        )
    }

Result

.. image:: assets/rooturladmin.png

**uadmin.Salt**
^^^^^^^^^^^^^^^
Salt is extra salt added to password hashing.

Syntax:

.. code-block:: go

    string

Go to the friend.go and apply the following codes below:

.. code-block:: go

    // This function hashes a password with a salt.
    func hashPass(pass string) string {
        // Generates a random string
        uadmin.Salt = uadmin.GenerateBase64(20)

        // Combine salt and password
        password := []byte(uadmin.Salt + pass)

        // Returns the bcrypt hash of the password at the given cost
        hash, err := bcrypt.GenerateFromPassword(password, 12)
        if err != nil {
            log.Fatal(err)
        }

        // Returns the string of hash value
        return string(hash)
    }

    // Save !
    func (f *Friend) Save() {

        // Calls the function of hashPass to store the value in the password
        // field.
        f.Password = hashPass(f.Password)
        
        // Override save
        uadmin.Save(f)
    }

Now go to the Friend model and put the password as 123456. Save it and check the result.

.. image:: assets/passwordwithsalt.png

**uadmin.Save**
^^^^^^^^^^^^^^^
Save saves the object in the database.

Syntax:

.. code-block:: go

    func(a interface{}) (err error)

Parameters:

    **a interface{}:** Is the variable where the model name was initialized.

Let's add an Invite field in the friend.go that will direct you to his website. In order to do that, set the field name as "Invite" with the tag "link".

.. code-block:: go

    // Friend model ...
    type Friend struct {
        uadmin.Model
        Name        string 
        Email       string 
        Password    string 
        Nationality string
        Invite      string `uadmin:"link"`
    }

To make it functional, add the overriding save function after the Friend struct.

.. code-block:: go

    // Save !
    func (f *Friend) Save() {
        f.Invite = "https://uadmin.io/"
        uadmin.Save(f) // <-- place it here
    }

Run your application, go to the Friends model and update the elements inside. Afterwards, click the Invite button on the output structure and see what happens.

.. image:: tutorial/assets/invitebuttonhighlighted.png

|

Result

.. image:: tutorial/assets/uadminwebsitescreen.png

**uadmin.Schema**
^^^^^^^^^^^^^^^^^
Schema is the global schema of the system.

Syntax:

.. code-block:: go

    map[string]uadmin.ModelSchema

Before you proceed to this example, see `uadmin.ModelSchema`_.

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Some codes
        // uadmin.F codes here
        // uadmin.ModelSchema codes here

        // Sets the actual name in the field from a modelschema
        uadmin.Schema[modelschema.ModelName].FieldByName("Name").DisplayName = modelschema.DisplayName

        // Generates the converted string value of two fields combined
        uadmin.Schema[modelschema.ModelName].FieldByName("Name").DefaultValue = modelschema.Fields[0].Value.(string) + " " + modelschema.Fields[1].Value.(string)

        // Set the Name field of an Expression model as required
        uadmin.Schema[modelschema.ModelName].FieldByName("Name").Required = true
    }

Alternative/shortcut way:

.. code-block:: go

    func main(){
        // Sets the actual name in the field from a modelschema
        modelschema.FieldByName("Name").DisplayName = modelschema.DisplayName

        // Generates the converted string value of two fields combined
        modelschema.FieldByName("Name").DefaultValue = modelschema.Fields[0].Value.(string) + " " + modelschema.Fields[1].Value.(string)

        // Set the Name field of an Expression model as required
        modelschema.FieldByName("Name").Required = true
    }

Now run your application, go to the Expression model and see what happens.

The name of the field has changed to "What's on your mind?"

.. image:: assets/expressiondisplayname.png

|

Click Add New Expression button at the top right corner and see what happens.

.. image:: assets/expressionrequireddefault.png

|

Well done! The Name field is now set to required and the value has automatically generated using the Schema function.

**uadmin.SendEmail**
^^^^^^^^^^^^^^^^^^^^
SendEmail sends email using system configured variables.

Syntax:

.. code-block:: go

    func(to, cc, bcc []string, subject, body string) (err error)

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){

        // Some codes

        // Email configurations
        uadmin.EmailFrom = "myemail@integritynet.biz"
        uadmin.EmailUsername = "myemail@integritynet.biz"
        uadmin.EmailPassword = "abc123"
        uadmin.EmailSMTPServer = "smtp.integritynet.biz"
        uadmin.EmailSMTPServerPort = 587

        // Place it here
        uadmin.SendEmail([]string{"myemail@integritynet.biz"}, []string{}, []string{}, "Todo List", "Here are the tasks that I should have done today.")
    }

Once you are done, open your email account. You will receive an email from a sender.

.. image:: assets/sendemailnotification.png

**uadmin.Session**
^^^^^^^^^^^^^^^^^^
Session is an activity that a user with a unique IP address spends on a Web site during a specified period of time. [#f2]_

Syntax:

.. code-block:: go

    type Session struct {
        Model
        Key        string
        User       User `gorm:"ForeignKey:UserID" uadmin:"filter"`
        UserID     uint `fk:"true" displayName:"User"`
        LoginTime  time.Time
        LastLogin  time.Time
        Active     bool   `uadmin:"filter"`
        IP         string `uadmin:"filter"`
        PendingOTP bool   `uadmin:"filter"`
        ExpiresOn  *time.Time
    }

There are 5 functions that you can use in Session:

* **GenerateKey()** - Automatically generates a random string of characters for you
* **HideInDashboard()** - Return true and auto hide this from dashboard
* **Logout()** - Deactivates a session
* **Save()** - Saves the object in the database
* **String()** - Returns the value of the Key

There are 2 ways you can do for initialization process using this function: one-by-one and by group.

One-by-one initialization:

.. code-block:: go

    func main(){
        // Some codes
        session := uadmin.Session{}
        session.Key = "Key"
        session.UserID = 1
    }

By group initialization:

.. code-block:: go

    func main(){
        // Some codes
        session := uadmin.Session{
            Key:    "Key",
            UserID: 1,
        }
    }

In this example, we will use "by group" initialization process.

Go to the main.go and apply the following codes below after the RegisterInlines section.

.. code-block:: go

    func main(){

        // Some codes

        now := time.Now()
        then := now.AddDate(0, 0, 1)
        session := uadmin.Session{

            // Generates a random string dynamically
            Key: uadmin.GenerateBase64(20),

            // UserID of System Admin account
            UserID: 1,

            LoginTime:  now,
            LastLogin:  now,
            Active:     true,
            IP:         "",
            PendingOTP: false,
            ExpiresOn:  &then,
        }

        // This will create a new session based on the information assigned in
        // the session variable.
        session.Save()

        // Returns the value of the key
        uadmin.Trail(uadmin.INFO, "String() returns %s", session.String())
    }

Now run your application and see what happens.

**Terminal**

.. code-block:: bash

    [  INFO  ]   String() returns 0G81O_LecZLru3CTm_Qz

.. image:: assets/sessioncreated.png

The other way around is you can use **GenerateKey()** function instead of initializing the Key field inside the uadmin.Session. Omit the session.Save() as well because session.GenerateKey() has the ability to save it.

.. code-block:: go

    func main(){
        now := time.Now()
        then := now.AddDate(0, 0, 1)
        session := uadmin.Session{

            // ------------ KEY FIELD REMOVED ------------ 

            // UserID of System Admin account
            UserID: 1,

            LoginTime:  now,
            LastLogin:  now,
            Active:     true,
            IP:         "",
            PendingOTP: false,
            ExpiresOn:  &then,
        }

        // Automatically generates a random string of characters for you
        session.GenerateKey()

        // Deactivates a session
        session.Logout()

        // ------------ SESSION.SAVE() REMOVED ------------ 

        // Returns the value of the key
        uadmin.Trail(uadmin.INFO, "String() returns %s", session.String())
    }

Now run your application and see what happens.

**Terminal**

.. code-block:: bash

    [  INFO  ]   String() returns 8dDjMOvX8onCVuRUJstZ1Jrl

.. image:: assets/sessioncreated2.png

|

Suppose that "SESSIONS" model is visible in the dashboard.

.. image:: assets/sessionshighlighteddashboard.png

|

In order to hide it, you can use **HideInDashboard()** built-in function from uadmin.Session. Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Initialize the session and dashboardmenu
        session := uadmin.Session{}
        dashboardmenu := uadmin.DashboardMenu{}

        // Checks the url from the dashboardmenu. If it matches, it will
        // update the value of the Hidden field.
        uadmin.Update(&dashboardmenu, "Hidden", session.HideInDashboard(), "url = ?", "session")
    }

Now run your application, go to "DASHBOARD MENUS" and you will notice that Sessions is now hidden.

.. image:: assets/sessionshidden.png

**uadmin.SiteName**
^^^^^^^^^^^^^^^^^^^
SiteName is the name of the website that shows on title and dashboard.

Syntax:

.. code-block:: go

    string

Go to the main.go and assign the SiteName value as **Todo List**.

.. code-block:: go

    func main() {
        // Some codes
        uadmin.SiteName = "Todo List"
    }

Run your application and see the changes above the web browser.

.. image:: tutorial/assets/todolisttitle.png

**uadmin.StartSecureServer**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^
StartSecureServer is the process of activating a uAdmin server using a localhost IP or an apache with SSL security.

Syntax:

.. code-block:: go

    func(certFile, keyFile string)

To enable SSL for your project, you need an SSL certificate. This is a two parts system with a public key and a private key. The public key is used for encryption and the private key is used for decryption. To get an SSL certificate, you can generate one using openssl which is a tool for generating self-signed SSL certificate.

.. code-block:: bash

    openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout priv.pem -out pub.pem

It will ask you for several certificate parameters but you can just press “Enter” and skip filling them for development.

You can change the key size by changing 2048 to a higher value like 4096. For production, you would want to get a certificate that is not self-signed to avoid the SSL error message on the browser. For that, you can buy one from any SSL vendor or you can get a free one from `letsencrypt.org`_ or follow the instructions `here`_.

.. _letsencrypt.org: https://letsencrypt.org/
.. _here: https://medium.com/@saurabh6790/generate-wildcard-ssl-certificate-using-lets-encrypt-certbot-273e432794d7

Once installed, move the **pub.pem** and **priv.pem** to your project folder.

.. image:: assets/sslcertificate.png

|

Afterwards, go to the main.go and apply this function on the last section.

.. code-block:: go

    func main(){
        // Some codes
        uadmin.StartSecureServer("pub.pem", "priv.pem")
    }

Go to https://uadmin.io/ as an example of a secure server. Click the padlock icon at the top left section then click Certificate (Valid).

.. image:: assets/uadminiosecure.png

|

You will see the following information in the certificate viewer.

.. image:: assets/certificateinfo.png

**uadmin.StartServer**
^^^^^^^^^^^^^^^^^^^^^^
StartServer is the process of activating a uAdmin server using a localhost IP or an apache.

Syntax:

.. code-block:: go

    func()

Go to the main.go and put **uadmin.StartServer()** inside the main function.

.. code-block:: go

    func main() {
        // Some codes
        uadmin.StartServer() // <-- place it here
    }

Now to run your code:

.. code-block:: bash

    $ go build; ./todo
    [   OK   ]   Initializing DB: [9/9]
    [   OK   ]   Initializing Languages: [185/185]
    [  INFO  ]   Auto generated admin user. Username: admin, Password: admin.
    [   OK   ]   Server Started: http://0.0.0.0:8080
             ___       __          _
      __  __/   | ____/ /___ ___  (_)___
     / / / / /| |/ __  / __  __ \/ / __ \
    / /_/ / ___ / /_/ / / / / / / / / / /
    \__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/

**uadmin.Tf**
^^^^^^^^^^^^^
Tf is a function for translating strings into any given language.

Syntax:

.. code-block:: go

    func(path string, lang string, term string, args ...interface{}) string

Parameters:

    **path (string):** This is where to get the translation from. It is in the
    format of "GROUPNAME/FILENAME" for example: "models/Todo"

    **lang (string):** Is the language code. If empty string is passed we will use
    the default language.

    **term (string):** The term to translate.

    **args (...interface{}):** Is a list of args to fill the term with place holders.

|

First of all, create a back-end validation function inside the todo.go.

.. code-block:: go

    // Validate !
    func (t Todo) Validate() (errMsg map[string]string) {
        // Initialize the error messages
        errMsg = map[string]string{}

        // Get any records from the database that maches the name of
        // this record and make sure the record is not the record we are
        // editing right now
        todo := Todo{}
        system := "system"
        if uadmin.Count(&todo, "name = ? AND id <> ?", t.Name, t.ID) != 0 {
            errMsg["Name"] = uadmin.Tf("models/Todo/Name/errMsg", "", fmt.Sprintf("This todo name is already in the %s", system))
        }
        return
    }

Run your application and login using “admin” as username and password.

.. image:: assets/loginformadmin.png

|

Open "LANGUAGES" model.

.. image:: assets/languageshighlighted.png

|

Search whatever languages you want to be available in your application. For this example, let's choose Tagalog and set it to Active.

.. image:: assets/tagalogactive.png

|

Open "TODOS" model and create at least one record inside it.

.. image:: assets/todomodeloutput.png

|

Logout your account and login again. Set your language to **Wikang Tagalog (Tagalog)**.

.. image:: assets/loginformtagalog.png

|

Open "TODOS" model, create a duplicate record, save it and let's see what happens.

.. image:: assets/duplicaterecord.png
   :align: center

|

The error message appears. Now rebuild your application and see what happens.

.. code-block:: go

    [   OK   ]   Initializing DB: [9/9]
    [ WARNING]   Translation of tl at 0% [0/134]

It says tl is 0% which means we have not translated yet. 

From your project folder, go to static/i18n/models/todo.tl.json. Inside it, you will see a bunch of data in JSON format that says Translate Me. This is where you put your translated text. For this example, let's translate the err_msg value in Tagalog language then save it.

.. image:: assets/errmsgtagalog.png

|

Once you are done, go back to your application, refresh your browser and see what happens.

.. image:: assets/todotagalogtranslatedtf.png
   :align: center

|

And if you rebuild your application, you will notice that uAdmin has found 1 word we automatically translated and is telling us we are at 1% translation for the Tagalog language.

.. code-block:: bash

    [   OK   ]   Initializing DB: [13/13]
    [ WARNING]   Translation of tl at 1% [1/134]

Congrats, now you know how to translate your sentence using uadmin.Tf.

**uadmin.Theme**
^^^^^^^^^^^^^^^^
Theme is the name of the theme used in uAdmin.

Syntax:

.. code-block:: go

    string

**uadmin.Trail**
^^^^^^^^^^^^^^^^
Trail prints to the log.

Syntax:

.. code-block:: go

    func(level int, msg interface{}, i ...interface{})

Parameters:

    **level int:** This is where we apply Trail tags.

    **msg interface{}:** Is the string of characters used for output.

    **i ...interface{}:** A variable or container that can be used to store a value in the msg interface{}.

Trail has 6 different tags:

* DEBUG
* WORKING
* INFO
* OK
* WARNING
* ERROR

Let's apply them in the overriding save function under the friend.go.

.. code-block:: go

    // Save !
    func (f *Friend) Save() {
        f.Invite = "https://uadmin.io/"
        temp := "saved"                                                  // declare temp variable
        uadmin.Trail(uadmin.DEBUG, "Your friend has been %s.", temp)     // used DEBUG tag
        uadmin.Trail(uadmin.WORKING, "Your friend has been %s.", temp)   // used WORKING tag
        uadmin.Trail(uadmin.INFO, "Your friend has been %s.", temp)      // used INFO tag
        uadmin.Trail(uadmin.OK, "Your friend has been %s.", temp)        // used OK tag
        uadmin.Trail(uadmin.WARNING, "Someone %s your friend.", temp)    // used WARNING tag
        uadmin.Trail(uadmin.ERROR, "Your friend has not been %s.", temp) // used ERROR tag
        uadmin.Save(f)
    }

Run your application, go to the Friend model and save any of the elements inside it. Check your terminal afterwards to see the result.

.. image:: tutorial/assets/trailtagsoutput.png
   :align: center

The output shows the different colors per tag.

**uadmin.Translate**
^^^^^^^^^^^^^^^^^^^^
Translate is used to get a translation from a multilingual fields.

Syntax:

.. code-block:: go

    func(raw string, lang string, args ...bool) string

Before we proceed to the example, read `Tutorial Part 7 - Introduction to API`_ to familiarize how API works in uAdmin.

.. _Tutorial Part 7 - Introduction to API: https://uadmin.readthedocs.io/en/latest/tutorial/part7.html

Suppose I have two multilingual fields in my Item record.

.. image:: assets/itementl.png

Create a file named custom_todo.go inside the api folder with the following codes below:

.. code-block:: go

    // CustomTodoHandler !
    func CustomTodoHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/custom_todo")

        res := map[string]interface{}{}

        item := models.Item{}

        results := []map[string]interface{}{}

        uadmin.Get(&item, "id = 1")

        results = append(results, map[string]interface{}{
            "Description (en)": uadmin.Translate(item.Description, "en"),
            "Description (tl)": uadmin.Translate(item.Description, "tl"),
        })

        res["status"] = "ok"
        res["item"] = results
        uadmin.ReturnJSON(w, r, res)
    }

Establish a connection in the main.go to the API by using http.HandleFunc. It should be placed after the uadmin.Register and before the StartServer.

.. code-block:: go

    func main() {
        // Some codes

        // CustomTodoHandler
        http.HandleFunc("/custom_todo/", api.CustomTodoHandler) // <-- place it here
    }

api is the folder name while CustomTodoHandler is the name of the function inside custom_todo.go.

Run your application and see what happens.

.. image:: assets/translatejson.png

**uadmin.Update**
^^^^^^^^^^^^^^^^^
Update updates the field name and value of an interface.

Syntax:

.. code-block:: go

    func(a interface{}, fieldName string, value interface{}, query string, args ...interface{}) (err error)

Suppose you have one record in your Todo model.

.. image:: assets/todoreadabook.png

|

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Some codes

        // Initialize todo and id
        todo := models.TODO{}
        id := 1

        // Updates the Todo name
        uadmin.Update(&todo, "Name", "Read a magazine", "id = ?", id)
    }

Now run your application, go to the Todo model and see what happens.

.. image:: assets/todoreadamagazine.png

|

The Todo name has updated from "Read a book" to "Read a magazine".

**uadmin.UploadImageHandler**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
UploadImageHandler handles the uploading process of an image.

Syntax:

.. code-block:: go

    func(w http.ResponseWriter, r *http.Request, session *uadmin.Session)

**uadmin.User**
^^^^^^^^^^^^^^^
User is a system in uAdmin that is used to add, modify and delete the elements of the user.

Syntax:

.. code-block:: go

    type User struct {
        Model
        Username     string    `uadmin:"required;filter"`
        FirstName    string    `uadmin:"filter"`
        LastName     string    `uadmin:"filter"`
        Password     string    `uadmin:"required;password;help:To reset password, clear the field and type a new password.;list_exclude"`
        Email        string    `uadmin:"email"`
        Active       bool      `uadmin:"filter"`
        Admin        bool      `uadmin:"filter"`
        RemoteAccess bool      `uadmin:"filter"`
        UserGroup    UserGroup `uadmin:"filter"`
        UserGroupID  uint
        Photo        string `uadmin:"image"`
        LastLogin   *time.Time `uadmin:"read_only"`
        ExpiresOn   *time.Time
        OTPRequired bool
        OTPSeed     string `uadmin:"list_exclude;hidden;read_only"`
    }

There are 9 functions that you can use in User:

* **GetActiveSession()** - returns a pointer of `uadmin.Session`_
* **GetDashboardMenu()** - returns (menus []uadmin.DashboardMenu)
* **GetOTP()** - returns a string
* **HasAccess** - searches for the url in the modelName. Uses this syntax as shown below:

.. code-block:: go

    func(modelName string) UserPermission

**Login** - Returns the pointer of User and a bool for Is OTP Required. It uses this syntax as shown below:

.. code-block:: go

    func(pass string, otp string) *uadmin.Session

* **Save()** - Saves the object in the database
* **String()** - Returns the first name and the last name
* **Validate()** - Validate user when saving from uadmin. It returns (ret map[string]string).
* **VerifyOTP** - Verifies the OTP of the user. It uses this syntax as shown below:

.. code-block:: go

    func(pass string) bool

There are 2 ways you can do for initialization process using this function: one-by-one and by group.

One-by-one initialization:

.. code-block:: go

    func main(){
        // Some codes
        user := uadmin.User{}
        user.Username = "Username"
        user.FirstName = "First Name"
        user.LastName = "Last Name"
    }

By group initialization:

.. code-block:: go

    func main(){
        // Some codes
        user := uadmin.User{
            Username: "Username",
            FirstName: "First Name",
            LastName: "Last Name",
        }
    }

In this example, we will use "by group" initialization process.

Go to the main.go and apply the following codes below after the RegisterInlines section.

.. code-block:: go

    func main(){

        // Some codes

        now := time.Now()
        user := uadmin.User{
            Username:     "even",
            FirstName:    "Even",
            LastName:     "Demata",
            Password:     "123456",
            Email:        "evendemata@gmail.com",
            Active:       true,
            Admin:        false,
            RemoteAccess: false,
            UserGroupID:  1, // Front Desk
            Photo:        "/media/images/users.png",
            LastLogin:    &now,
            OTPRequired:  false,
        }

        // This will create a new user based on the information assigned in
        // the user variable.
        user.Save()

        // Returns the first name and the last name
        uadmin.Trail(uadmin.INFO, "String() returns %s.", user.String())
    }

Now run your application and see what happens.

**Terminal**

.. code-block:: bash

    [  INFO  ]   String() returns Even Demata.

.. image:: assets/usercreated.png

|

Select "Even Demata" account in the list.

.. image:: assets/evendematahighlighted.png

|

Go to the User Permission tab. Afterwards, click Add New User Permission button at the right side.

.. image:: assets/addnewuserpermission.png

|

Set the Dashboard Menu to "Todos" model, User linked to "Even Demata", and activate the "Read" only. It means Even Demata user account has restricted access to adding, editing and deleting a record in the Todos model.

.. image:: assets/userpermissionevendemata.png

|

Result

.. image:: assets/userpermissionevendemataoutput.png

|

Log out your System Admin account. This time login your username and password using the user account that has user permission. Afterwards, you will see that only the Todos model is shown in the dashboard because your user account is not an admin and has no remote access to it.

.. image:: assets/userpermissiondashboard.png

|

Now go back to the main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Initialize the User function
        user := uadmin.User{}

        // Fetch the username record as "even" from the user
        uadmin.Get(&user, "username = ?", "even")

        // Print the results
        fmt.Println("GetActiveSession() is", user.GetActiveSession())
        fmt.Println("GetDashboardMenu() is", user.GetDashboardMenu())
        fmt.Println("GetOTP() is", user.GetOTP())
        fmt.Println("HasAccess is", user.HasAccess("todo"))
    }

Run your application and check your terminal to see the results.

.. code-block:: bash

    GetActiveSession() is GOzo21lIBCIaj3YkXJsCZXnj
    GetDashboardMenu() is [Todos]
    GetOTP() is 251553
    HasAccess is 1

Take note the value of the GetOTP(). Go to the main.go again and apply the following codes below:

.. code-block:: go

    func main(){
        user := uadmin.User{}
        uadmin.Get(&user, "username = ?", "even")

        // First parameter is password and second parameter is the value from
        // GetOTP()
        fmt.Println("Login is", user.Login("123456", "251553"))

        // The parameter is the value from GetOTP()
        fmt.Println("VerifyOTP is", user.VerifyOTP("251553"))
    }

Run your application and check your terminal to see the results.

.. code-block:: bash

    Login is GOzo21lIBCIaj3YkXJsCZXnj
    VerifyOTP is true

If your Login does not return anything and VerifyOTP is false, it means your OTP is no longer valid. You need to use GetOTP() again to get a new one. OTP usually takes 5 minutes of validity by default.

**Validate()** function allows you to search if the user already exists. For instance, the username is "even" and all of the contents about him are there which was already included in the User model. Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Some codes
        now := time.Now()
        user := uadmin.User{
            Username:     "even",
            FirstName:    "Even",
            LastName:     "Demata",
            Password:     "123456",
            Email:        "evendemata@gmail.com",
            Active:       true,
            Admin:        false,
            RemoteAccess: false,
            UserGroupID:  1, // Front Desk
            Photo:        "/media/images/users.png",
            LastLogin:    &now,
            OTPRequired:  false,
        }

        fmt.Println("Validate is", user.Validate())
    }

Run your application and check your terminal to see the results.

.. code-block:: bash

    Validate is map[Username:Username is already Taken.]

Congrats, now you know how to configure the User fields, fetching the username record and applying the functions of the User.

**uadmin.UserGroup**
^^^^^^^^^^^^^^^^^^^^
UserGroup is a system in uAdmin used to add, modify, and delete the group name. 

Syntax:

.. code-block:: go

    type UserGroup struct {
        Model
        GroupName string `uadmin:"filter"`
    }

There are 2 functions that you can use in UserGroup:

**HasAccess()** - Uses this syntax as shown below:

.. code-block:: go

    func(modelName string) uadmin.GroupPermission

**String()** - Returns the GroupName

There are 2 ways you can do for initialization process using this function: one-by-one and by group.

One-by-one initialization:

.. code-block:: go

    func main(){
        // Some codes
        usergroup := uadmin.UserGroup{}
        user.GroupName = "Group Name"
    }

By group initialization:

.. code-block:: go

    func main(){
        // Some codes
        usergroup := uadmin.UserGroup{
            GroupName: "Group Name",
        }
    }

In this example, we will use "by group" initialization process.

Go to the main.go and apply the following codes below after the RegisterInlines section.

.. code-block:: go

    func main(){

        // Some codes

        usergroup := uadmin.UserGroup{
            GroupName: "Front Desk",
        }

        // This will create a new user group based on the information assigned
        // in the usergroup variable.
        uadmin.Save(&usergroup)

        // Returns the GroupName
        uadmin.Trail(uadmin.INFO, "String() returns %s.", usergroup.String())
    }

Now run your application and see what happens.

**Terminal**

.. code-block:: bash

    [  INFO  ]   String() returns Front Desk.

.. image:: assets/usergroupcreated.png

|

Link your created user group to any of your existing accounts (example below is Even Demata).

.. image:: assets/useraccountfrontdesklinked.png

|

Afterwards, click the Front Desk highlighted below.

.. image:: assets/frontdeskhighlighted.png

|

Go to the Group Permission tab. Afterwards, click Add New Group Permission button at the right side.

.. image:: assets/addnewgrouppermission.png

|

Set the Dashboard Menu to "Todos" model, User linked to "Even Demata", and activate the "Read" only. It means Front Desk User Group has restricted access to adding, editing and deleting a record in the Todos model.

.. image:: assets/grouppermissionadd.png

|

Result

.. image:: assets/grouppermissionaddoutput.png

|

Log out your System Admin account. This time login your username and password using the user account that has group permission.

.. image:: assets/userpermissiondashboard.png

|

Now go back to the main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Initializes the UserGroup function
        usergroup := uadmin.UserGroup{}

        // Fetches the Group Permission ID from the user
        uadmin.Get(&usergroup, "id = ?", 1)

        // Prints the HasAccess result
        fmt.Println("HasAccess is", usergroup.HasAccess("todo"))
    }

Run your application and check your terminal to see the result.

.. code-block:: bash

    HasAccess is 1

Congrats, now you know how to add the UserGroup from code, fetching the record from ID and applying the functions of the UserGroup.

**uadmin.UserPermission**
^^^^^^^^^^^^^^^^^^^^^^^^^
UserPermission sets the permission of a user handled by an administrator.

Syntax:

.. code-block:: go

    type UserPermission struct {
        Model
        DashboardMenu   DashboardMenu `gorm:"ForeignKey:DashboardMenuID" required:"true" filter:"true" uadmin:"filter"`
        DashboardMenuID uint          `fk:"true" displayName:"DashboardMenu"`
        User            User          `gorm:"ForeignKey:UserID" required:"true" filter:"true" uadmin:"filter"`
        UserID          uint          `fk:"true" displayName:"User"`
        Read            bool          `uadmin:"filter"`
        Add             bool          `uadmin:"filter"`
        Edit            bool          `uadmin:"filter"`
        Delete          bool          `uadmin:"filter"`
    }

There are 2 functions that you can use in GroupPermission:

* **HideInDashboard()** - Return true and auto hide this from dashboard
* **String()** - Returns the UserPermission ID

There are 2 ways you can do for initialization process using this function: one-by-one and by group.

One-by-one initialization:

.. code-block:: go

    func main(){
        // Some codes
        userpermission := uadmin.UserPermission{}
        userpermission.DashboardMenu = dashboardmenu
        userpermission.DashboardMenuID = 1
        userpermission.User = user
        userpermission.UserID = 1
    }

By group initialization:

.. code-block:: go

    func main(){
        // Some codes
        userpermission := uadmin.UserPermission{
            DashboardMenu: dashboardmenu,
            DashboardMenuID: 1,
            User: user,
            UserID: 1,
        }
    }

In this example, we will use "by group" initialization process.

Go to the main.go and apply the following codes below after the RegisterInlines section.

.. code-block:: go

    func main(){

        // Some codes

        userpermission := uadmin.UserPermission{
            DashboardMenuID: 9,     // Todos
            UserID:          2,     // Even Demata
            Read:            true,
            Add:             false,
            Edit:            false,
            Delete:          false,
        }

        // This will create a new user permission based on the information
        // assigned in the userpermission variable.
        uadmin.Save(&userpermission)
    }

Now run your application and see what happens.

.. image:: assets/userpermissioncreated.png

|

Log out your System Admin account. This time login your username and password using the user account that has user permission. Afterwards, you will see that only the Todos model is shown in the dashboard because your user account is not an admin and has no remote access to it. Now click on TODOS model.

.. image:: assets/userpermissiondashboard.png

|

As you will see, your user account is restricted to add, edit, or delete a record in the Todo model. You can only read what is inside this model.

.. image:: assets/useraddeditdeleterestricted.png

|

If you want to hide the Todo model in your dashboard, first of all, create a HideInDashboard() function in your todo.go inside the models folder and set the return value to “true”.

.. code-block:: go

    // HideInDashboard !
    func (t Todo) HideInDashboard() bool {
        return true
    }

Now you can do something like this in main.go:

.. code-block:: go

    func main(){

        // Some codes

        // Initializes the DashboardMenu
        dashboardmenu := uadmin.DashboardMenu{}

        // Assign the userpermission, call the HideInDashboard() function
        // from todo.go, store it to the Hidden field of the dashboardmenu
        dashboardmenu.Hidden = userpermission.HideInDashboard()

        // Checks the Dashboard Menu ID number from the userpermission. If it
        // matches, it will update the value of the Hidden field.
        uadmin.Update(&dashboardmenu, "Hidden", dashboardmenu.Hidden, "id = ?", userpermission.DashboardMenuID)
    }

Now rerun your application using the Even Demata account and see what happens.

.. image:: assets/dashboardmenuempty.png

|

The Todo model is now hidden from the dashboard. If you login your System Admin account, you will see in the Dashboard menu that the hidden field of the Todo model is set to true.

.. image:: assets/todomodelhidden.png

**uadmin.Version**
^^^^^^^^^^^^^^^^^^
Version number as per Semantic Versioning 2.0.0 (semver.org)

Syntax:

.. code-block:: go

    untyped string

Let's check what version of uAdmin are we using.

.. code-block:: go

    func main() {
        // Some codes
        uadmin.Trail(uadmin.INFO, uadmin.Version)
    }

Result

.. code-block:: bash

    [   OK   ]   Initializing DB: [9/9]
    [  INFO  ]   0.1.0-beta.4
    [   OK   ]   Server Started: http://0.0.0.0:8080
             ___       __          _
      __  __/   | ____/ /___ ___  (_)___
     / / / / /| |/ __  / __  __ \/ / __ \
    / /_/ / ___ / /_/ / / / / / / / / / /
    \__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/

You can also directly check it by typing **uadmin version** in your terminal.

.. code-block:: bash

    $ uadmin version
    [  INFO  ]   0.1.0-beta.4

**uadmin.WARNING**
^^^^^^^^^^^^^^^^^^
WARNING is the display tag under Trail. It is the statement or event that indicates a possible problems occurring in an application.

Syntax:

.. code-block:: go

    untyped int

See `uadmin.Trail`_ for the example.

**uadmin.WORKING**
^^^^^^^^^^^^^^^^^^
OK is the display tag under Trail. It is a status to show that the application is working.

Syntax:

.. code-block:: go

    untyped int

See `uadmin.Trail`_ for the example.

Reference
---------
.. [#f1] Rouse, Margaret (2018). MongoDB. Retrieved from https://searchdatamanagement.techtarget.com/definition/MongoDB
.. [#f2] QuinStreet Inc. (2018). User Session. Retrieved from https://www.webopedia.com/TERM/U/user_session.html