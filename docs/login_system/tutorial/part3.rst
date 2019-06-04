Login System Tutorial Part 3 - Sending Request
==============================================
In this tutorial, we will talk about sending data from login form in HTML to the LoginHandler.

Go to login.go in handlers folder and initialize the User model inside the LoginHandler function.

.. code-block:: go

    // LoginHandler !
    func LoginHandler(w http.ResponseWriter, r *http.Request) {
        // Some codes

        // Initialize the User model from uAdmin
        user := uadmin.User{}
    }

Let's create a validation whether the user submits requests in the HTML form and if the value of the request is "login".

.. code-block:: go

    // Check if the user submits request in HTML form
    if r.Method == "POST" {
        // Check if the value of the request is login
        if r.FormValue("request") == "login" {

        }
    }

Inside the r.FormValue("request") condition, check if the username from the HTML form contains a value.

.. code-block:: go

    // Create the parameter of "username"
    username := r.FormValue("username")

    // Get the user record where username is the assigned value
    uadmin.Get(&user, "username=?", username)

    // Print the result
    uadmin.Trail(uadmin.DEBUG, "Username: %s", username)

Now run your application and go to the login path in the address bar (e.g. http://0.0.0.0:8080/login/). Assign the username in the login form (e.g. admin). Click Login button to submit.

.. image:: assets/adminusername.png
   :align: center

|

Check your terminal for the result.

.. code-block:: bash

    [  DEBUG ]   Username: admin

Exit your application. Go to the login.go in the handlers folder. Create a validation whether the fetched record from the User model is existing. Inside the validation, assign the parameters then check if the password and OTP password contains a value.

.. code-block:: go

    // Check if the fetched record is existing in the User model
    if user.ID > 0 {
        // Create the parameters of "password" and "otp_pass"
        password := r.FormValue("password")
        otpPass := r.FormValue("otp_pass")

        // Print results
        uadmin.Trail(uadmin.DEBUG, "Password: %s", password)
        uadmin.Trail(uadmin.DEBUG, "OTP Password: %s", otpPass)
    }

Now run your application and go to the login path in the address bar (e.g. http://0.0.0.0:8080/login/). Assign the username, password, and OTP password in the login form (e.g. admin, admin, 123456). Click Login button to submit.

.. image:: assets/adminloginformdatatest.png
   :align: center

|

Check your terminal for the result.

.. code-block:: bash

    [  DEBUG ]   Username: admin
    [  DEBUG ]   Password: admin
    [  DEBUG ]   OTP Password: 123456

In the `next part`_, we will discuss about checking the status of the user login with and without OTP validation.

Click `here`_ to view the full source code in this part.

.. _next part: https://uadmin.readthedocs.io/en/latest/login_system/tutorial/part4.html

.. _here: https://uadmin.readthedocs.io/en/latest/login_system/tutorial/full_code/part3.html

.. toctree::
   :maxdepth: 1

   full_code/part3
