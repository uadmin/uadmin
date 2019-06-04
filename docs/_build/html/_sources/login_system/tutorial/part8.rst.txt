Login System Tutorial Part 8 - Webpage Manipulation
===================================================
In this tutorial, we will discuss about reading a cookie and getting the user from the model based on the value of the cookie to ensure that the user is active.

First of all, run your application. Go to the login path in the address bar (e.g. http://0.0.0.0:8080/login/). Assign the username and password in the login form (e.g. admin, admin). Click Login button to submit.

.. image:: assets/adminusernamepassword.png
   :align: center

|

Result

.. image:: assets/logoutbuttoncreated.png

|

Now refresh your webpage and see what happens.

.. image:: assets/customloginform.png
   :align: center

|

It went back to the login form but if you take a look on the user_session cookie, it is active because we have not implemented the handler yet that redirects to the home page.

In order to do that, first we need to read the cookie of the user_session.

.. code-block:: go

    if r.Method == "POST" {
        // Some codes
    }

    // Read the cookie of "user_session"
    cookie, _ := r.Cookie("user_session")

    // Print the result
    uadmin.Trail(uadmin.DEBUG, "Cookie: %v", cookie)

Run your application. Go to the login path in the address bar (e.g. http://0.0.0.0:8080/login/). Check your terminal for the result.

.. code-block:: bash

    [  DEBUG ]   Cookie: user_session=1

Exit your application. Now create a handler that fetches the user record based on the value of the cookie, assign it to the userContext data object and pass that object to home.html.

.. code-block:: go

    // Check if the fetched cookie is existing
    if cookie != nil {
        // Get the user record based on the value of the cookie
        uadmin.Get(&user, "id = ?", cookie.Value)

        // Assign the full name of the user and OTP boolean value to the
        // userContext
        userContext = UserContext{
            User: &user,
            OTP:  user.OTPRequired,
        }

        // Pass the userContext data object to the HTML file
        uadmin.HTMLContext(w, userContext, "views/home.html")
        return
    }

Now run your application. Go to the login path in the address bar (e.g. http://0.0.0.0:8080/login/). Assign the username and password in the login form (e.g. admin, admin). Click Login button to submit.

.. image:: assets/adminusernamepassword.png
   :align: center

|

Result

.. image:: assets/logoutbuttoncreated.png

|

Now refresh your webpage and see what happens.

.. image:: assets/logoutbuttoncreated.png

|

Click on Logout button then check the result.

.. image:: assets/logoutresult.png
   :align: center

|

Refresh your webpage once again and see what happens.

.. image:: assets/customloginform.png
   :align: center

|

Congrats, now you know how to do the following in the entire series:

* Preparing uAdmin files in the project folder
* Build an application from scratch
* Change the dashboard title
* Create custom login form in HTML
* Sending request from front-end to back-end
* Getting the session key based on the user login status
* Setting an HTTP cookie
* OTP Scanning
* Logout user
* Deleting the cookie
* Maintaining the webpage based on the user login status

Click `here`_ to view the full source code in this part.

.. _here: https://uadmin.readthedocs.io/en/latest/login_system/tutorial/full_code/part8.html

.. toctree::
   :maxdepth: 1

   full_code/part8

If you want to learn more and discover about the concepts of uAdmin, you may go to these references with examples:

* `API Reference`_
* `Quick Reference`_
* `System Reference`_
* `Tag Reference`_

.. _API Reference: https://uadmin.readthedocs.io/en/latest/api.html
.. _Quick Reference: https://uadmin.readthedocs.io/en/latest/quick_reference.html
.. _System Reference: https://uadmin.readthedocs.io/en/latest/system_reference.html
.. _Tag Reference: https://uadmin.readthedocs.io/en/latest/tags.html
