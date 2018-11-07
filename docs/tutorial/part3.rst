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

You may also set your own IP address within the range of 127.0.0.1 - 127.255.255.254 by using BindIP. BindIP is the IP the application listens to. It means you can connect only to that IP you have assigned to. Let's say **127.0.0.2**

.. code-block:: go

    func main() {
        // Some codes are contained in this line ... (ignore this part)
        uadmin.BindIP = "127.0.0.2" // <--  place it here
        uadmin.Port = 8000
    }

If you run your code,

.. code-block:: bash

    [   OK   ]   Initializing DB: [12/12]
    [   OK   ]   Server Started: http://127.0.0.2:8000
            ___       __          _
    __  __/   | ____/ /___ ___  (_)___
    / / / / /| |/ __  / __  __ \/ / __ \
    / /_/ / ___ / /_/ / / / / / / / / / /
    \__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/

In the Server Started, it will redirect you to the IP address of **127.0.0.2**.

But if you connect to other IP address within the range of 127.0.0.1 - 127.255.255.254 it will not work as shown below (User connects to 127.0.0.3).

.. image:: assets/bindiphighlighted.png

|

uAdmin has a feature that allows a user to set his own site name by using uadmin.SiteName. Let's say **Todo List**.

.. code-block:: go

    func main() {
        // Some codes are contained in this line ... (ignore this part)
        uadmin.SiteName = "Todo List" // <--  place it here
        uadmin.BindIP = "127.0.0.2"
        uadmin.Port = 8000
    }

Run your application and see the changes above the web browser.

.. image:: assets/todolisttitle.png

|

You can use ReportTimeStamp to have a time stamp in your logs.

.. code-block:: go

    func main() {
        // Some codes are contained in this line ... (ignore this part)
        uadmin.ReportTimeStamp = true // <--  place it here
        uadmin.SiteName = "Todo List"
        uadmin.BindIP = "127.0.0.2"
        uadmin.Port = 8000
    }

If you run your code,

.. code-block:: bash

    [   OK   ]   Initializing DB: [12/12]
    2018/11/07 08:52:14 [   OK   ]   Server Started: http://127.0.0.2:8000
            ___       __          _
    __  __/   | ____/ /___ ___  (_)___
    / / / / /| |/ __  / __  __ \/ / __ \
    / /_/ / ___ / /_/ / / / / / / / / / /
    \__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/

|

You can set a ReportingLevel to change the status of the debugging process.

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
        uadmin.ReportTimeStamp = true 
        uadmin.SiteName = "Todo List"
        uadmin.BindIP = "127.0.0.2"
        uadmin.Port = 8000
    }

Result

.. code-block:: bash

    [   OK   ]   Initializing DB: [12/12]
    2018/11/07 08:52:14 [   OK   ]   Server Started: http://127.0.0.2:8000
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
        uadmin.ReportTimeStamp = true 
        uadmin.SiteName = "Todo List"
        uadmin.BindIP = "127.0.0.2"
        uadmin.Port = 8000
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

uAdmin has built-in debugging functions that uses **uadmin.Trail**. Trail is necessary if you want to test something in your code.

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

.. image:: assets/trailtagsoutput.png

The output shows the different colors per tag.

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

We can also configure an email using uadmin public functions. 

**EmailFrom** identifies where the email is coming from.
    
**EmailUsername** sets the username of an email.
    
**EmailPassword** sets the password of an email.
    
**EmailSMTPServer** sets the name of the SMTP Server in an email.

**EmailSMTPServerPort** sets the port number of an SMTP Server in an email.

.. code-block:: go

    func main(){
        uadmin.EmailFrom = "rmamisay@integritynet.biz"
        uadmin.EmailUsername = "rmamisay@integritynet.biz"
        uadmin.EmailPassword = "abc123"
        uadmin.EmailSMTPServer = "smtp.integritynet.biz"
        uadmin.EmailSMTPServerPort = 587
        // Some codes are contained in this line ... (ignore this part)
    }

Let's go back to the uAdmin dashboard, go to Users model, create your own user account and set the email address based on your assigned EmailFrom in the code above.

.. image:: assets/useremailhighlighted.png

|

Log out your account. At the moment, you suddenly forgot your password. How can we retrieve our account? Click Forgot Password at the bottom of the login form.

.. image:: assets/forgotpasswordhighlighted.png

|

Input your email address based on the user account you wish to retrieve it back.

.. image:: assets/forgotpasswordinputemail.png

|

Once you are done, open your email account. You will receive a password reset notification from the Todo List support. To reset your password, click the link highlighted below.

.. image:: assets/passwordresetnotification.png

|

You will be greeted by the reset password form. Input the following information in order to create a new password for you.

.. image:: assets/resetpasswordform.png

Once you are done, you can now access your account using your new password.

You can use CookieTimeout function to set the timeout of a login cookie in seconds. Let's say 10 seconds.

.. code-block:: go

    func main() {
        // Some codes are contained in this line ... (ignore this part)
        uadmin.CookieTimeout = 10 // <--  place it here
        uadmin.ReportingLevel = 1
        uadmin.ReportTimeStamp = true 
        uadmin.SiteName = "Todo List"
        uadmin.BindIP = "127.0.0.2"
        uadmin.Port = 8000
    }

.. WARNING::
   Use it at your own risk. Once the cookie expires in your user account, your account will be permanently deactivated. In this case, you must have an extra user account in the User database.

Login your account, wait for 10 seconds and see what happens.

.. image:: assets/loginform.png

It will redirect you to the login form because your cookie has already been expired.