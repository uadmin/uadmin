Login System Tutorial Part 7 - Logout
=====================================
In this tutorial, we will talk about logging out a user account from the home page.

First of all, go to home.html in the views folder and add logout button below the user login status.

.. code-block:: html

    <body>
        <h1>Login as {{.User}}
        {{if eq .OTP true}} with {{else}} without {{end}}
        2FA Authentication</h1>

        <!-- ADD THIS PIECE OF CODE TO CREATE LOGOUT BUTTON -->
        <form method="POST">
            <button type="submit" name="request" value="logout">Logout</button>
        </form>
    </body>

Go to login.html and add this piece of code to notify the user that he logged out his account.

.. code-block:: html

    <body>
        <!-- ADD THIS PIECE OF CODE TO CREATE LOGOUT NOTIFICATION -->
        <p>{{.Message}}</p>

        <form method="POST">
            <!-- Some input fields -->
        </form>
    </body>

Now go to login.go in handlers folder and apply the following codes below to delete the cookie when the user logged out his account:

.. code-block:: go

    if r.FormValue("request") == "login" {
        // Some codes
    }

    // Check if the request submitted is logout
    if r.FormValue("request") == "logout" {
        // Assign the message to the Message field of userContext
        userContext.Message = "User has logged out."

        // Logout the user in uAdmin
        uadmin.Logout(r)

        // Deletes the cookie
        usersession := &http.Cookie{
            Name:   "user_session",
            Value:  "",
            MaxAge: -1,
        }
        http.SetCookie(w, usersession)

        // Pass the userContext data object to the HTML file
        uadmin.HTMLContext(w, userContext, "views/login.html")
        return
    }

Run your application. Go to the login path in the address bar (e.g. http://0.0.0.0:8080/login/). Assign the username and password in the login form (e.g. admin, admin). Click Login button to submit.

.. image:: assets/adminusernamepassword.png
   :align: center

|

As expected, the logout button has been created in the form. Click Logout button and see what happens.

.. image:: assets/logoutbuttoncreated.png

|

Result

.. image:: assets/logoutnotification.png

|

Now check the user_session cookie to ensure that it was deleted.

.. image:: assets/usersessioncookiedeleted.png
   :align: center

|

In the `next part`_, we will discuss about reading a cookie and getting the user from the model based on the value of the cookie to ensure that the user is active.

Click `here`_ to view the full source code in this part.

.. _next part: https://uadmin.readthedocs.io/en/latest/login_system/tutorial/part8.html

.. _here: https://uadmin.readthedocs.io/en/latest/login_system/tutorial/full_code/part7.html

.. toctree::
   :maxdepth: 1

   full_code/part7
