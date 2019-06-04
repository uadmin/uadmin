Login System Tutorial Part 6 - Home Page
========================================
In this tutorial, we will discuss about redirecting a webpage after the user submits a form and passing the data object to the Home Page.

Before you proceed, make sure you have at least the basic knowledge of HTML. If you are not familiar with HTML, we advise you to go over `W3Schools`_.

.. _W3Schools: https://www.w3schools.com/

Create a new file in the views folder named "home.html" and apply the following codes below:

.. code-block:: html

    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <title>Home Page</title>
    </head>
    <body>
        <!-- .User is a field that came from the UserContext struct
        in Golang -->
        <h1>Login as {{.User}}
        <!-- Validate if the OTP is enabled in the user -->
        {{if eq .OTP true}} with {{else}} without {{end}}
        2FA Authentication</h1>
    </body>
    </html>

Now assign the value of the login, login2fa, and otp in the UserContext struct.

.. code-block:: go

    if otp == true && user.VerifyOTP(otpPass) {
        http.SetCookie(w, usersession)

        // ----------------------- PLACE IT HERE -----------------------
        // Assign the full name of the user and OTP boolean value to the
        // userContext
        userContext = UserContext{
            User: login2fa,
            OTP:  otp,
        }

        // Pass the userContext data object to the HTML file
        uadmin.HTMLContext(w, userContext, "views/home.html")
        return
    }

    if otp == false && otpPass == "" {
        http.SetCookie(w, usersession)

        // ----------------------- PLACE IT HERE -----------------------
        // Assign the full name of the user and OTP boolean value to the
        // userContext
        userContext = UserContext{
            User: login,
            OTP:  otp,
        }

        // Pass the userContext data object to the HTML file
        uadmin.HTMLContext(w, userContext, "views/home.html")
        return
    }

Run your application. Go to the login path in the address bar (e.g. http://0.0.0.0:8080/login/). Assign the username, password, and OTP password fetched from the 2FA image in /admin/profile/ path in the address bar or assigned on your terminal in the login form (e.g. admin, admin, 123456). Click Login button to submit.

.. image:: assets/adminloginformdatatest.png
   :align: center

|

Result

.. image:: assets/loginwith2faresult.png

|

Now go to the admin path in the address bar (e.g. http://0.0.0.0:8080/admin/). Inside the "USERS" model, disable the OTPRequired in the System Admin user.

.. image:: assets/otprequiredfalse.png

|

Go back to the login path in the address bar (e.g. http://0.0.0.0:8080/login/). Assign the username and password in the login form (e.g. admin, admin). Click Login button to submit.

.. image:: assets/adminusernamepassword.png
   :align: center

|

Result

.. image:: assets/loginwithout2faresult.png

|

In the `next part`_, we will talk about logging out a user account from the home page.

Click `here`_ to view the full source code in this part.

.. _next part: https://uadmin.readthedocs.io/en/latest/login_system/tutorial/part7.html

.. _here: https://uadmin.readthedocs.io/en/latest/login_system/tutorial/full_code/part6.html

.. toctree::
   :maxdepth: 1

   full_code/part6
