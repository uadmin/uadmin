Login System Tutorial Part 4 - Login Access Debugging
=====================================================
In this tutorial, we will discuss about checking the status of the user login with and without OTP validation.

Go to login.go in handlers folder and call the login function that passes the HTTP request, username, and password. If all three parameters are valid, it will return the full name of the User and the boolean value for IsOTPRequired.

.. code-block:: go

    if user.ID > 0 {
        password := r.FormValue("password")

        // Comment otpPass for now
        // otpPass := r.FormValue("otp_pass")

        // ----------------------- PLACE IT HERE -----------------------
        // Pass the requested username and password in Login function to
        // return the full name of the User and the boolean value for
        // IsOTPRequired
        login, otp := uadmin.Login(r, username, password)

        // Print results
        uadmin.Trail(uadmin.DEBUG, "Login as: %s", login)
        uadmin.Trail(uadmin.DEBUG, "OTP: %t", otp)
    }

Now run your application and go to the login path in the address bar (e.g. http://0.0.0.0:8080/login/). Assign the username and password in the login form (e.g. admin, admin). Click Login button to submit.

.. image:: assets/adminusernamepassword.png
   :align: center

|

Check your terminal for the result.

.. code-block:: bash

    [  DEBUG ]   Login as: System Admin
    [  DEBUG ]   OTP: false

By default, the OTPRequired for that user is disabled.

Exit your application. Go to login.go in handlers folder and call the login2fa function that passes the HTTP request, username, password, and OTP Password. If all four parameters are valid, it will return the full name of the User.

.. code-block:: go

    // Check if the fetched record is existing in the User model
    if user.ID > 0 {
        password := r.FormValue("password")

        // Uncomment this part
        otpPass := r.FormValue("otp_pass")

        // Pass the requested username and password in Login function to
        // return the full name of the User and the boolean value for
        // IsOTPRequired
        login, otp := uadmin.Login(r, username, password)

        // ----------------------- PLACE IT HERE -----------------------
        // Initialize Login2FA that returns the User
        login2fa := &uadmin.User{}

        // Check whether the OTP value from Login function is true
        // and the OTP Password is valid
        if otp == true && user.VerifyOTP(otpPass) {
            // Pass the requested username, password, and OTP Password in
            // Login2FA function to return the full name of the User
            login2fa = uadmin.Login2FA(r, username, password, otpPass)

            // Print the result
            uadmin.Trail(uadmin.DEBUG, "Login with 2FA as: %s", login2fa)
        }
    }

Run your application and go to the admin path in the address bar (e.g. http://0.0.0.0:8080/admin/). Login using “admin” as username and password.

.. image:: assets/loginform.png

|

Click on "USERS".

.. image:: assets/usermodelhighlighted.png

|

Click System Admin.

.. image:: assets/systemadminhighlighted.png

|

Scroll down the form then activate OTP Required on that user.

.. image:: assets/activateotprequired.png

|

Result

.. image:: assets/otprequiredtrue.png

|

Click the blue person icon on the top right corner then select admin in order to visit the profile page.

.. image:: assets/adminhighlightedprofile.png
   :align: center

|

Scroll down the form. There is a 2FA image to fetch the QR code which is typically used for storing URLs or other information for reading by the camera on a smartphone. In order to do that, you can use Google Authenticator (`Android`_, `iOS`_). It is a software-based authenticator that implements two-step verification services using the Time-based One-time Password Algorithm and HMAC-based One-time Password algorithm, for authenticating users of mobile applications by Google. [#f1]_

.. image:: assets/2faimage.png
   :align: center

.. _Android: https://play.google.com/store/apps/details?id=com.google.android.apps.authenticator2&hl=en
.. _iOS: https://itunes.apple.com/ph/app/google-authenticator/id388497605?mt=8

|

If there is a problem, you may go to your terminal and check the OTP verification code for login.

Now go to the login path in the address bar (e.g. http://0.0.0.0:8080/login/). Assign the username, password, and OTP password that you fetched from the 2FA image in the login form (e.g. admin, admin, 123456). Click Login button to submit.

.. image:: assets/adminloginformdatatest.png
   :align: center

|

Check your terminal for the result.

.. code-block:: bash

    [  DEBUG ]   Login with 2FA as: System Admin
    [  DEBUG ]   Login as: System Admin
    [  DEBUG ]   OTP: true

In the `next part`_, we will talk about getting the session key if the user login is valid and setting an HTTP cookie for the user session.

Click `here`_ to view the full source code in this part.

.. _next part: https://uadmin.readthedocs.io/en/latest/login_system/tutorial/part5.html

.. _here: https://uadmin.readthedocs.io/en/latest/login_system/tutorial/full_code/part4.html

.. toctree::
   :maxdepth: 1

   full_code/part4

Reference
---------
.. [#f1] No author (28 May 2019). Google Authenticator. Retrieved from https://en.wikipedia.org/wiki/Google_Authenticator
