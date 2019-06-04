Login System Tutorial Part 4 - Login Access Debugging (Full Source Code)
========================================================================

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

    // LoginHandler !
    func LoginHandler(w http.ResponseWriter, r *http.Request) {
        // r.URL.Path creates a new path called /login
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/login")

        // Initialize the User model from uAdmin
        user := uadmin.User{}

        // Initialize the UserContext struct that we have created
        userContext := UserContext{}

        // Check if the user submits request in HTML form
        if r.Method == "POST" {
            // Check if the value of the request is login
            if r.FormValue("request") == "login" {
                // Create the parameter of "username"
                username := r.FormValue("username")

                // Get the user record where username is the assigned value
                uadmin.Get(&user, "username=?", username)

                // Check if the fetched record is existing in the User model
                if user.ID > 0 {
                    // Create the parameters of "password" and "otp_pass"
                    password := r.FormValue("password")
                    otpPass := r.FormValue("otp_pass")

                    // Pass the requested username and password in Login function to
                    // return the full name of the User and the boolean value for
                    // IsOTPRequired
                    login, otp := uadmin.Login(r, username, password)

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

                    // Print results
                    uadmin.Trail(uadmin.DEBUG, "Login as: %s", login)
                    uadmin.Trail(uadmin.DEBUG, "OTP: %t", otp)
                }
            }
        }

        // Pass the userContext data object to the HTML file
        uadmin.HTMLContext(w, userContext, "views/login.html")
        return
    }
