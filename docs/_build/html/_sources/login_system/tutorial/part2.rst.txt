Login System Tutorial Part 2 - Login Form
=========================================
In this, we will discuss about creating a login form in HTML.

Before you proceed, make sure you have at least the basic knowledge of HTML. If you are not familiar with HTML, we advise you to go over `W3Schools`_.

.. _W3Schools: https://www.w3schools.com/


First of all, create a new file in the views folder named "login.html".

.. image:: assets/loginfileviews.png

Inside login.html, create a login form containing the username, password, OTP Password input fields, and a submit button with the method of "POST".

.. code-block:: html

    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <title>Login Form</title>
    </head>
    <body>
        <form method="POST">
            <!-- The assigned name attribute is equivalent to r.FormValue in
            Golang while the assigned value attribute is the value of the
            r.FormValue. (e.g. r.FormValue("request") = "login") -->
            <input type="text" name="username" placeholder="Username"><br>
            <input type="password" name="password" placeholder="Password"><br>
            <input type="text" name="otp_pass" placeholder="OTP Password"><br><br>
            <button type="submit" name="request" value="login">Login</button><br>
        </form>
    </body>
    </html>

Now create a new file in the handlers folder named "login.go".

.. image:: assets/loginfilehandler.png

Inside login.go, let's build a UserContext struct containing the User field that has a data type of a pointer of uadmin.User that is a built-in system model, an OTP field that has a data type of boolean, and a Message field that has a data type of string. We will use that in the later chapter of this tutorial.

.. code-block:: go

    package handlers

    import (
        "net/http"
        "strings"

        "github.com/uadmin/uadmin"
    )

    // UserContext !
    type UserContext struct {
        User    *uadmin.User
        OTP     bool
        Message string
    }

Below the UserContext struct, create a LoginHandler function that creates a new URL path of "login" and passes the userContext data object to the login.html file.

.. code-block:: go

    // LoginHandler !
    func LoginHandler(w http.ResponseWriter, r *http.Request) {
        // r.URL.Path creates a new path called /login
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/login")

        // Initialize the UserContext struct that we have created
        userContext := UserContext{}

        // Pass the userContext data object to the HTML file
        uadmin.HTMLContext(w, userContext, "views/login.html")
        return
    }

Establish a connection in the main.go to the handlers by using http.HandleFunc. It should be placed before the StartServer.

.. code-block:: go

    import (
        "net/http"

        // Specify the username that you used inside github.com folder
        "github.com/username/project_name/handlers"

        "github.com/uadmin/uadmin"
    )

    func main() {
        // Some codes

        // Login Handler
        http.HandleFunc("/login/", handlers.LoginHandler)
    }

Now run your application. Go to the login path (e.g. http://0.0.0.0:8080/login/) and see what happens.

.. image:: assets/customloginform.png
   :align: center

|

In the `next part`_, we will talk about sending data from login form in HTML to the LoginHandler.

Click `here`_ to view the full source code in this part.

.. _next part: https://uadmin.readthedocs.io/en/latest/login_system/tutorial/part3.html

.. _here: https://uadmin.readthedocs.io/en/latest/login_system/tutorial/full_code/part2.html

.. toctree::
   :maxdepth: 1

   full_code/part2
