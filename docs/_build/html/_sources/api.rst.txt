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

There are 11 types of actions:

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

Go to the logs in the uAdmin dashboard. You can see the Action field inside it as shown below.

.. image:: assets/actionhighlighted.png

|

Go to the main.go. Let's return a value of each types of actions.

.. code-block:: go

    func main(){
        // Some codes contained in this part
    uadmin.Trail(uadmin.INFO, "Added = %v", uadmin.Action.Added(0))
    uadmin.Trail(uadmin.INFO, "Custom = %v", uadmin.Action.Custom(0))
    uadmin.Trail(uadmin.INFO, "Deleted = %v", uadmin.Action.Deleted(0))
    uadmin.Trail(uadmin.INFO, "LoginDenied = %v", uadmin.Action.LoginDenied(0))
    uadmin.Trail(uadmin.INFO, "LoginSuccessful = %v", uadmin.Action.LoginSuccessful(0))
    uadmin.Trail(uadmin.INFO, "Logout = %v", uadmin.Action.Logout(0))
    uadmin.Trail(uadmin.INFO, "Modified = %v", uadmin.Action.Modified(0))
    uadmin.Trail(uadmin.INFO, "PasswordResetDenied = %v", uadmin.Action.PasswordResetDenied(0))
    uadmin.Trail(uadmin.INFO, "PasswordResetRequest = %v", uadmin.Action.PasswordResetRequest(0))
    uadmin.Trail(uadmin.INFO, "PasswordResetSuccessful = %v", uadmin.Action.PasswordResetSuccessful(0))
    uadmin.Trail(uadmin.INFO, "Read = %v", uadmin.Action.Read(0))
    }

Check your terminal to see the result.

.. code-block:: go

    [  INFO  ]   Added = 2
    [  INFO  ]   Custom = 11
    [  INFO  ]   Deleted = 4
    [  INFO  ]   LoginDenied = 6
    [  INFO  ]   LoginSuccessful = 5
    [  INFO  ]   Logout = 7
    [  INFO  ]   Modified = 3
    [  INFO  ]   PasswordResetDenied = 9
    [  INFO  ]   PasswordResetRequest = 8
    [  INFO  ]   PasswordResetSuccessful = 10
    [  INFO  ]   Read = 1
    
**uadmin.AdminPage**
^^^^^^^^^^^^^^^^^^^^
AdminPage fetches records from the database with some standard rules such as sorting data, multiples of, and setting a limit that can be used in pagination.

Syntax:

.. code-block:: go

    AdminPage func(order string, asc bool, offset int, limit int, a interface{}, query interface{}, args ...interface{}) (err error)

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

    All func(a interface{}) (err error)

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
        // Some codes are contained in this line ... (ignore this part)
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

First of all, create a function with a parameter of interface{} and a pointer of User that returns an array of Choice which will be used that later below the main function in main.go.

.. code-block:: go

    func GetChoices(m interface{}, user *uadmin.User) []uadmin.Choice {
        // Build choices
        choices := []uadmin.Choice{
            uadmin.Choice{
                K: 0,
                V: "-",
            },
        }

        choices = append(choices, uadmin.Choice{
            V:        uadmin.GetString(m),
            K:        uadmin.GetID(reflect.ValueOf(m)),
            Selected: true,
        })

        return choices
    }

Now inside the main function, apply `uadmin.Schema`_ function that calls a model name of "todo", accesses "Choices" as the field name that uses the LimitChoices to then assign it to GetChoices which is your function name.

.. code-block:: go

    uadmin.Schema["todo"].FieldByName("Choices").LimitChoicesTo = GetChoices

Run your application, go to the Todo model and see what happens in the Choices field.

.. image:: assets/choicestrue.png

|

Well done! You have created one choice that gets from the Todo name itself. You can also add the list of choices manually. Put it in the GetChoices function between the first choice that you have created and the return value.

.. code-block:: go

	choices = append(choices, uadmin.Choice{
		V:        "Build a robot",
		K:        1,
		Selected: false,
	})
	choices = append(choices, uadmin.Choice{
		V:        "Washing the dishes",
		K:        2,
		Selected: false,
	})

Now rerun your application to see the result.

.. image:: assets/choicesfalse.png

|

Well done! You have a total of 3 choices in the list.

**uadmin.ClearDB**
^^^^^^^^^^^^^^^^^^
ClearDB clears the database object.

Syntax:

.. code-block:: go

    ClearDB func()

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
        // Some codes are contained in this part.
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

    CookieTimeout int

Let's apply this function in the main.go.

.. code-block:: go

    func main() {
        // Some codes are contained in this line ... (ignore this part)
        uadmin.CookieTimeout = 10 // <--  place it here
    }

.. WARNING::
   Use it at your own risk. Once the cookie expires in your user account, your account will be permanently deactivated. In this case, you must have an extra user account in the User database.

Login your account, wait for 10 seconds and see what happens.

.. image:: tutorial/assets/loginform.png

It will redirect you to the login form because your cookie has already been expired.

**uadmin.Count**
^^^^^^^^^^^^^^^^
Count return the count of records in a table based on a filter.

Syntax:

.. code-block:: go

    Count func(a interface{}, query interface{}, args ...interface{}) int

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

    CustomTranslation []string

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Some codes are contained in this part.
        uadmin.CustomTranslation = []string{"uadmin/system", "uadmin/user"}
        fmt.Println(uadmin.CustomTranslation)
    }

Result

.. code-block:: bash

    [uadmin/system uadmin/user]

**uadmin.DashboardMenu**
^^^^^^^^^^^^^^^^^^^^^^^^
DashboardMenu is a system in uAdmin used to add, modify, and delete the settings of a model.

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

Go to the main.go and apply the following codes below after the RegisterInlines section.

.. code-block:: go

    func main(){

        // Some codes are contained in this part.

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

**uadmin.Database**
^^^^^^^^^^^^^^^^^^^
Database is the active Database settings.

Syntax:

.. code-block:: go

    Database *DBSettings

See `uadmin.DBSettings`_ for the example.

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

.. image:: tutorial/assets/todolistdbhighlighted.png

**uadmin.DEBUG**
^^^^^^^^^^^^^^^^
DEBUG is the display tag under Trail. It is the process of identifying and removing errors.

Syntax:

.. code-block:: go

    const DEBUG int = 0

See `uadmin.Trail`_ for the example.

**uadmin.DebugDB**
^^^^^^^^^^^^^^^^^^
DebugDB prints all SQL statements going to DB.

Syntax:

.. code-block:: go

    DebugDB bool

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

    Delete func(a interface{}) (err error)

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

    DeleteList func(a interface{}, query interface{}, args ...interface{}) (err error)

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

    EmailFrom string

.. code-block:: go

    func main(){
        uadmin.EmailFrom = "myemail@integritynet.biz"
        uadmin.EmailUsername = "myemail@integritynet.biz"
        uadmin.EmailPassword = "abc123"
        uadmin.EmailSMTPServer = "smtp.integritynet.biz"
        uadmin.EmailSMTPServerPort = 587
        // Some codes are contained in this line ... (ignore this part)
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

    EmailPassword string

See `uadmin.EmailFrom`_ for the example.

**uadmin.EmailSMTPServer**
^^^^^^^^^^^^^^^^^^^^^^^^^^
EmailSMTPServer sets the name of the SMTP Server in an email.

Syntax:

.. code-block:: go

    EmailSMTPServer string

See `uadmin.EmailFrom`_ for the example.

**uadmin.EmailSMTPServerPort**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
EmailSMTPServerPort sets the port number of an SMTP Server in an email.

Syntax:

.. code-block:: go

    EmailSMTPServerPort int

See `uadmin.EmailFrom`_ for the example.

**uadmin.EmailUsername**
^^^^^^^^^^^^^^^^^^^^^^^^
EmailUsername sets the username of an email.

Syntax:

.. code-block:: go

    EmailUsername string

See `uadmin.EmailFrom`_ for the example.

**uadmin.ERROR**
^^^^^^^^^^^^^^^^
ERROR is a status to notify the user that there is a problem in an application.

Syntax:

.. code-block:: go

    const ERROR int = 5

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

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Some codes are contained in this part.
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

    Filter func(a interface{}, query interface{}, args ...interface{}) (err error)

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

    FilterBuilder func(params map[string]interface{}) (query string, args []interface{})

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

    GenerateBase32 func(length int) string

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

    GenerateBase64 func(length int) string

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

    Get func(a interface{}, query interface{}, args ...interface{}) (err error)

Parameters:

    **a interface{}:** Is the variable where the model name was initialized.

    **query interface{}:** Is an action that you want to perform with in your data list.

    **args ...interface{}:** Is the variable or container that can be used in execution process.

Suppose you have ten records in your Todo model.

.. image:: tutorial/assets/tendataintodomodel.png

|

Go to the main.go. Let's count how many todos do you have with a friend in your model.

.. code-block:: go

    func main(){
        // Some codes contained in this part

        // Initialized the Todo model in the todo variable
        todo := models.Todo{}

        // Initialized the Friend model in the todo variable
        friend := models.Friend{}

        // Fetches the first record from the database
        uadmin.Get(&friend, "id=?", todo.FriendID)

        // Returns the count of records in a table based on a Get function to  
        // be stored in the total variable
        total := uadmin.Count(&todo, "friend_id = ?", todo.FriendID)

        // Prints the result
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

    GetDB func() *gorm.DB

See `uadmin.ClearDB`_ for the example.

**uadmin.GetID**
^^^^^^^^^^^^^^^^
GetID returns an ID number of a field.

Syntax:

.. code-block:: go

    GetID func(m.reflectValue) uint

See `uadmin.Choice`_ for the example.

**uadmin.GetString**
^^^^^^^^^^^^^^^^^^^^
GetString returns string representation on an instance of a model.

Syntax:

.. code-block:: go

    GetString func(a interface{}) string

Parameters:

    **a interface{}:** Is the variable where the model name was initialized.

See `uadmin.Choice`_ for the example.

**uadmin.GetUserFromRequest**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
GetUserFromRequest returns a user from a request.

Syntax:

.. code-block:: go

    GetUserFromRequest func(r *http.Request) *User

Before we proceed to the example, read `Tutorial Part 7 - Introduction to API`_ to familiarize how API works in uAdmin.

Suppose that the admin account has logined.

.. image:: tutorial/assets/adminhighlighted.png

|

Create a file named info.go inside the api folder with the following codes below:

.. code-block:: go

    // InfoHandler !
    func InfoHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/info")

        res := map[string]interface{}{}

        // Place it here
        uadmin.Trail(uadmin.INFO, "GetUserFromRequest: %s", uadmin.GetUserFromRequest(r))

        res["status"] = "ok"
        uadmin.ReturnJSON(w, r, res)
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
   :align: center

|

Check your terminal for the result.

.. code-block:: bash

    [  INFO  ]   GetUserFromRequest: System Admin

The result is coming from the user in the dashboard.

.. image:: assets/getuserfromrequest.png

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

Suppose that Even Demata account is a part of the Front Desk User Group.

.. image:: assets/useraccountfrontdesk.png

|

Go to the main.go and apply the following codes below after the RegisterInlines section.

.. code-block:: go

    func main(){

        // Some codes are contained in this part.

        grouppermission := uadmin.GroupPermission{
            DashboardMenuID: 9,     // Todos
            UserGroupID:     1,     // Front Desk
            Read:            true,
            Add:             false,
            Edit:            false,
            Delete:          false,
        }

        // This will create a new group permission based on the information
        // assigned in the grouppermission variable.
        uadmin.Save(&grouppermission)
    }

Now run your application and see what happens.

.. image:: assets/grouppermissioncreated.png

|

Log out your System Admin account. This time login your username and password using the user account that has group permission. Afterwards, you will see that only the Todos model is shown in the dashboard because your user account is not an admin and has no remote access to it. Now click on TODOS model.

.. image:: assets/userpermissiondashboard.png

|

As you will see, your user account is restricted to add, edit, or delete a record in the Todo model. You can only read what is inside this model.

.. image:: assets/useraddeditdeleterestricted.png

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

    const INFO int = 2

See `uadmin.Trail`_ for the example.

**uadmin.IsAuthenticated**
^^^^^^^^^^^^^^^^^^^^^^^^^^
IsAuthenticated returns the session of the user.

Syntax:

.. code-block:: go

    IsAuthenticated func(r *http.Request) *Session

Before we proceed to the example, read `Tutorial Part 7 - Introduction to API`_ to familiarize how API works in uAdmin.

Suppose that the admin account has logined.

.. image:: tutorial/assets/adminhighlighted.png

|

Create a file named info.go inside the api folder with the following codes below:

.. code-block:: go

    // InfoHandler !
    func InfoHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/info")

        res := map[string]interface{}{}

        // Place it here
        uadmin.Trail(uadmin.INFO, "IsAuthenticated: %s", uadmin.IsAuthenticated(r))

        res["status"] = "ok"
        uadmin.ReturnJSON(w, r, res)
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
   :align: center

|

Check your terminal for the result.

.. code-block:: bash

    [  INFO  ]   IsAuthenticated: FbdwBVT30p-4a7Afrsp3gvM0

The result is coming from the session in the dashboard.

.. image:: assets/isauthenticated.png

**uadmin.JSONMarshal**
^^^^^^^^^^^^^^^^^^^^^^
JSONMarshal returns the JSON encoding of v.

Syntax:

.. code-block:: go

    JSONMarshal func(v interface{}, safeEncoding bool) ([]byte, error)

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
Language is a system in uAdmin used to add and modify the settings of a language.

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

Suppose the Tagalog language is not active and you want to set this to Active.

.. image:: assets/tagalognotactive.png

|

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Some codes are contained in this part.
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
        uadmin.Update(&language, "Active", true, "english_name = ?", language.EnglishName)
    }

Now run your application, refresh your browser and see what happens.

.. image:: assets/tagalogactive.png

|

As expected, the Tagalog language is now set to active.

**uadmin.Log**
^^^^^^^^^^^^^^
Log is a system in uAdmin used to add, modify, and delete the status of the user activities.

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

Go to the main.go and apply the following codes below after the RegisterInlines section.

.. code-block:: go

    func main(){

        // Some codes are contained in this part.

        log := uadmin.Log{
            Username:  "admin",
            Action:    uadmin.Action.Added(0),
            TableName: "Todo",
            TableID:   1,
            Activity:  "Manually added from uadmin.Log in the main function",
            RollBack:  "",
            CreatedAt: time.Now(),
        }

        // This will create a new log based on the information assigned in
        // the log variable.
        uadmin.Save(&log)
    }

Now run your application and see what happens.

.. image:: assets/logcreated.png

**uadmin.Login**
^^^^^^^^^^^^^^^^
Login returns the pointer of User and a bool for Is OTP Required.

Syntax:

.. code-block:: go

    Login func(r *http.Request, username string, password string) (*User, bool)

Before we proceed to the example, read `Tutorial Part 7 - Introduction to API`_ to familiarize how API works in uAdmin.

Create a file named info.go inside the api folder with the following codes below:

.. code-block:: go

    // InfoHandler !
    func InfoHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/info")

        res := map[string]interface{}{}

        fmt.Println(uadmin.Login(r, "admin", "admin")) // <-- place it here

        res["status"] = "ok"
        uadmin.ReturnJSON(w, r, res)
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
   :align: center

|

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

    Login2FA func(r *http.Request, username string, password string, otpPass string) *User

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

        res := map[string]interface{}{}

        // Place it here
        fmt.Println(uadmin.Login2FA(r, "admin", "admin", "445215"))

        res["status"] = "ok"
        uadmin.ReturnJSON(w, r, res)
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
   :align: center

|

Check your terminal for the result.

.. code-block:: bash

    System Admin

**uadmin.Logout**
^^^^^^^^^^^^^^^^^
Logout deactivates a session.

Syntax:

.. code-block:: go

    Logout func(r *http.Request)

Suppose that the admin account has logined.

.. image:: tutorial/assets/adminhighlighted.png

|

Create a file named logout.go inside the api folder with the following codes below:

.. code-block:: go

    // LogoutHandler !
    func LogoutHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/logout")

        res := map[string]interface{}{}

        uadmin.Logout(r)

        res["status"] = "ok"
        uadmin.ReturnJSON(w, r, res)
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
   :align: center

|

Refresh your browser and see what happens.

.. image:: tutorial/assets/loginform.png

|

Your account has been logged out automatically that redirects you to the login form.

**uadmin.MaxImageHeight**
^^^^^^^^^^^^^^^^^^^^^^^^^
MaxImageHeight sets the maximum height of an image.

Syntax:

.. code-block:: go

    MaxImageHeight int

See `uadmin.MaxImageWidth`_ for the example.

**uadmin.MaxImageWidth**
^^^^^^^^^^^^^^^^^^^^^^^^
MaxImageWidth sets the maximum width of an image.

Syntax:

.. code-block:: go

    MaxImageWidth int

Let's set the MaxImageWidth to 360 pixels and the MaxImageHeight to 240 pixels.

.. code-block:: go

    func main() {
        // Some codes are contained in this line ... (ignore this part)
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

    MaxUploadFileSize int64

Go to the main.go. Let's set the MaxUploadFileSize value to 1024. 1024 is equivalent to 1 MB.

.. code-block:: go

    func main() {
        // Some codes are contained in this line ... (ignore this part)
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

Before you proceed to this example, see `uadmin.F`_.

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Some codes are contained in this part.
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

    MongoDB *MongoSettings

**uadmin.MongoModel (Experimental)**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
MongoModel is a uAdmin function for interfacing with MongoDB databases.

Syntax:

.. code-block:: go

    type MongoModel struct {
	    ID bson.ObjectId `bson:"_id,omitempty"`
    }

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

    NewModel func(modelName string, pointer bool) (reflect.Value, bool)

Suppose I have three records in my Expressions model with an ID of 4, 5, 6.

.. image:: assets/expressionthreevalues.png

|

Now I want to fetch only the last record inside that model. Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){
    
        // Some codes are contained in this part.

        // Checks and fetches a record from the expression database with an
        // ID of 6. 
        if m, ok := uadmin.NewModel("expression", true); ok {
            uadmin.Get(m.Interface(), "id = ?", 6)
            fmt.Println(m.Interface())
        }
    }

Now run your application and check your terminal to see the result.

.. code-block:: bash

    &{{6 <nil>} Nice! 1}

**uadmin.NewModelArray**
^^^^^^^^^^^^^^^^^^^^^^^^
NewModelArray creates a new model array from a model name.

Syntax:

.. code-block:: go

    NewModelArray func(modelName string, pointer bool) (reflect.Value, bool)

Suppose I have three records in my Expressions model with an ID of 4, 5, 6.

.. image:: assets/expressionthreevalues.png

|

Now I want to fetch all records inside that model. Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){
    
        // Some codes are contained in this part.

        // Checks and fetches records from the expression database with an
        // ID greater than 1.
        if m, ok := uadmin.NewModelArray("expression", true); ok {
            uadmin.Filter(m.Interface(), "id > ?", 1)
            fmt.Println(m.Interface())
        }
    }

Now run your application and check your terminal to see the result.

.. code-block:: bash

    &[{{4 <nil>} Yes! 1} {{5 <nil>} Wow! 1} {{6 <nil>} Nice! 1}]

**uadmin.OK**
^^^^^^^^^^^^^
OK is the display tag under Trail. It is a status to show that the application is doing well.

Syntax:

.. code-block:: go

    const OK int = 3

See `uadmin.Trail`_ for the example.

**uadmin.OTPAlgorithm**
^^^^^^^^^^^^^^^^^^^^^^^
OTPAlgorithm is the hashing algorithm of OTP.

Syntax:

.. code-block:: go

    OTPAlgorithm string

There are 3 different algorithms:

* sha1 (default)
* sha256
* sha512

**uadmin.OTPDigits**
^^^^^^^^^^^^^^^^^^^^
OTPDigits is the number of digits for the OTP.

Syntax:

.. code-block:: go

    OTPDigits int

Go to the main.go and set the OTPDigits to 8.

.. code-block:: go

    func main() {
        // Some codes are contained in this line ... (ignore this part)
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

    OTPPeriod uint

Go to the main.go and set the OTPPeriod to 10 seconds.

.. code-block:: go

    func main() {
        // Some codes are contained in this line ... (ignore this part)
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

    OTPSkew uint

Go to the main.go and set the OTPSkew to 2 minutes.

.. code-block:: go

    func main() {
        // Some codes are contained in this line ... (ignore this part)
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

    PageLength int

Go to the main.go and apply the PageLength function.

.. code-block:: go

    func main() {
        // Some codes are contained in this line ... (ignore this part)
        uadmin.PageLength = 4  // <--  place it here
    }

Run your application, go to the Item model, inside it you have 6 total elements. The elements in the item model will display 4 elements per page.

.. image:: tutorial/assets/pagelength.png

**uadmin.Port**
^^^^^^^^^^^^^^^
Port is the port used for http or https server.

Syntax:

.. code-block:: go

    Port int

Go to the main.go in your Todo list project and apply **8000** as a port number.

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

    PublicMedia bool

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

    Register func(m ...interface{})

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

    ReportingLevel int

There are 6 different levels:

* DEBUG   = 0
* WORKING = 1
* INFO    = 2
* OK      = 3
* WARNING = 4
* ERROR   = 5

Let's set the ReportingLevel to 1 to show that the debugging process is working.

.. code-block:: go

    func main() {
        // Some codes are contained in this line ... (ignore this part)
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
        // Some codes are contained in this line ... (ignore this part)
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

    ReportTimeStamp bool

Go to the main.go and set the ReportTimeStamp value as true.

.. code-block:: go

    func main() {
        // Some codes are contained in this line ... (ignore this part)
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

    ReturnJSON func(w http.ResponseWriter, r *http.Request, v interface{})

See `Tutorial Part 7 - Introduction to API`_ for the example.

.. _Tutorial Part 7 - Introduction to API: https://uadmin.readthedocs.io/en/latest/tutorial/part7.html

**uadmin.RootURL**
^^^^^^^^^^^^^^^^^^
RootURL is where the listener is mapped to.

Syntax:

.. code-block:: go

    RootURL string

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

    Salt string

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

    Save func(a interface{}) (err error)

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

    Schema map[string]ModelSchema

Before you proceed to this example, see `uadmin.ModelSchema`_.

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Some codes are contained in this part.
        // uadmin.F codes here
        // uadmin.ModelSchema codes here

        // Sets the actual name in the field from a modelschema
        uadmin.Schema[modelschema.ModelName].FieldByName("Name").DisplayName = modelschema.DisplayName

        // Generates the converted string value of two fields combined
        uadmin.Schema[modelschema.ModelName].FieldByName("Name").DefaultValue = modelschema.Fields[0].Value.(string) + " " + modelschema.Fields[1].Value.(string)

        // Set the Name field of an Expression model as required
        uadmin.Schema[modelschema.ModelName].FieldByName("Name").Required = true
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

    SendEmail func(to, cc, bcc []string, subject, body string) (err error)

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){

        // Some codes are contained in this part.

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

Go to the main.go and apply the following codes below after the RegisterInlines section.

.. code-block:: go

    func main(){

        // Some codes are contained in this part.

        now := time.Now()
        then := now.AddDate(0, 0, 1)
        session := uadmin.Session{
            // Generates a random string dynamically
            Key:        uadmin.GenerateBase64(20),
            // UserID of System Admin account
            UserID:     1,
            LoginTime:  now,
            LastLogin:  now,
            Active:     true,
            IP:         "",
            PendingOTP: false,
            ExpiresOn:  &then,
        }

        // This will create a new session based on the information assigned in
        // the session variable.
        uadmin.Save(&session)
    }

Now run your application and see what happens.

.. image:: assets/sessioncreated.png

**uadmin.SiteName**
^^^^^^^^^^^^^^^^^^^
SiteName is the name of the website that shows on title and dashboard.

Syntax:

.. code-block:: go

    SiteName string

Go to the main.go and assign the SiteName value as **Todo List**.

.. code-block:: go

    func main() {
        // Some codes are contained in this line ... (ignore this part)
        uadmin.SiteName = "Todo List"
    }

Run your application and see the changes above the web browser.

.. image:: tutorial/assets/todolisttitle.png

**uadmin.StartSecureServer**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^
StartSecureServer is the process of activating a uAdmin server using a localhost IP or an apache with SSL certificate and a private key.

Syntax:

.. code-block:: go

    StartSecureServer func(certFile, keyFile string)

First of all, get your wildcard certificate using Let's Encrypt/Certbot `here`_.

.. _here: https://medium.com/@saurabh6790/generate-wildcard-ssl-certificate-using-lets-encrypt-certbot-273e432794d7

Once installed, move the **fullchain.pem** and **privkey.pem** to your project folder.

.. image:: assets/sslcertificate.png

|

Afterwards, go to the main.go and apply this function on the last section.

.. code-block:: go

    func main(){
        // Some codes are contained in this part.
        uadmin.StartSecureServer("fullchain.pem", "privkey.pem")
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

    StartServer func()

Go to the main.go and put **uadmin.StartServer()** inside the main function.

.. code-block:: go

    func main() {
        // Some codes are contained in this line ... (ignore this part)
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

    Tf func(path string, lang string, term string, args ...interface{}) string

Parameters:

    **path (string):** This is where to get the translation from. It is in the
    format of "GROUPNAME/FILENAME" for example: "models/Todo"

    **lang (string):** Is the language code. If empty string is passed we will use
    the default language.

    **term (string):** The term to translate.

    **args (...interface{}):** Is a list of args to fill the term with place holders.

|

Create a back-end validation function inside the todo.go.

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

Run your application and see what happens.

.. code-block:: bash

    [   OK   ]   Initializing DB: [13/13]
    [ WARNING]   Translation of tl at 1% [1/170]

uAdmin has found 1 word we automatically translated and is telling us we are at 1% translation for the Tagalog language.

Login your account and set your language as **Wikang Tagalog (Tagalog)**

.. image:: assets/loginformtagalog.png

|

Suppose you have one record in your Todo model.

.. image:: assets/todomodeloutput.png

|

Now create a duplicate record in Todo model and see what happens.

.. image:: assets/todotagalogtranslatedtf.png

|

Congrats, you know now how to translate your sentence using uadmin.Tf.

**uadmin.Theme**
^^^^^^^^^^^^^^^^
Theme is the name of the theme used in uAdmin.

Syntax:

.. code-block:: go

    Theme string

**uadmin.Trail**
^^^^^^^^^^^^^^^^
Trail prints to the log.

Syntax:

.. code-block:: go

    Trail func(level int, msg interface{}, i ...interface{})

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

    Translate func(raw string, lang string, args ...bool) string

Go to the item.go inside the models folder and apply the following codes below:

.. code-block:: go

    // Item model ...
    type Item struct {
        uadmin.Model
        Name        string `uadmin:"required"`
        Description string `uadmin:"multilingual"` // <-- set this tag
        Cost   int
        Rating int
    }

    // Save ...
    func (i *Item) Save() {
        // This function can translate any type of language
        uadmin.Translate(i.Description, "", true)

        uadmin.Save(i)
    }

Run your application. Suppose I want to translate my description from English to Tagalog. Go to the Item model, manually translate your description and store it in the tl field. X symbol means it is not yet translated.

.. image:: assets/tlnotyetranslated.png

|

Save it, log out your account then login again. Set your language to **Wikang Tagalog (Tagalog)**.

.. image:: assets/loginformtagalog.png

|

Now open your Item model. The item description is now translated to Tagalog language.

.. image:: assets/tltranslated.png

**uadmin.Update**
^^^^^^^^^^^^^^^^^
Update updates the field name and value of an interface.

Syntax:

.. code-block:: go

    Update func(a interface{}, fieldName string, value interface{}, query string, args ...interface{}) (err error)

Suppose you have one record in your Todo model.

.. image:: assets/todoreadabook.png

|

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Some codes are contained in this part.

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

    UploadImageHandler func(w http.ResponseWriter, r *http.Request, session *Session)

**uadmin.User**
^^^^^^^^^^^^^^^
User is a system in uAdmin used to check and modify the settings of a user.

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

Go to the main.go and apply the following codes below after the RegisterInlines section.

.. code-block:: go

    func main(){

        // Some codes are contained in this part.

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
            UserGroupID:  1,    // Front Desk
            Photo:        "/media/images/users.png",
            LastLogin:    &now,
            OTPRequired:  false,
        }

        // This will create a new user based on the information assigned in
        // the user variable.
        uadmin.Save(&user)
    }

Now run your application and see what happens.

.. image:: assets/usercreated.png

**uadmin.UserGroup**
^^^^^^^^^^^^^^^^^^^^
UserGroup is a system in uAdmin used to add, modify, and delete the group name. 

Syntax:

.. code-block:: go

    type UserGroup struct {
        Model
        GroupName string `uadmin:"filter"`
    }

Go to the main.go and apply the following codes below after the RegisterInlines section.

.. code-block:: go

    func main(){

        // Some codes are contained in this part.

        usergroup := uadmin.UserGroup{
            GroupName: "Front Desk",
        }

        // This will create a new user group based on the information assigned
        // in the usergroup variable.
        uadmin.Save(&usergroup)
    }

Now run your application and see what happens.

.. image:: assets/usergroupcreated.png

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

Go to the main.go and apply the following codes below after the RegisterInlines section.

.. code-block:: go

    func main(){

        // Some codes are contained in this part.

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

**uadmin.Version**
^^^^^^^^^^^^^^^^^^
Version number as per Semantic Versioning 2.0.0 (semver.org)

Syntax:

.. code-block:: go

    const Version string = "0.1.0-beta.4"

Let's check what version of uAdmin are we using.

.. code-block:: go

    func main() {
        // Some codes are contained in this line ... (ignore this part)
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

    const WARNING int = 4

See `uadmin.Trail`_ for the example.

**uadmin.WORKING**
^^^^^^^^^^^^^^^^^^
OK is the display tag under Trail. It is a status to show that the application is working.

Syntax:

.. code-block:: go

    const WORKING int = 1

See `uadmin.Trail`_ for the example.

Reference
---------
.. [#f1] Rouse, Margaret (2018). MongoDB. Retrieved from https://searchdatamanagement.techtarget.com/definition/MongoDB
.. [#f2] QuinStreet Inc. (2018). User Session. Retrieved from https://www.webopedia.com/TERM/U/user_session.html