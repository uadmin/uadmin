Login System Tutorial Part 3 - Sending Request (Full Source Code)
=================================================================

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

                // Print the result
                uadmin.Trail(uadmin.DEBUG, "Username: %s", username)

                // Check if the fetched record is existing in the User model
                if user.ID > 0 {
                    // Create the parameters of "password" and "otp_pass"
                    password := r.FormValue("password")
                    otpPass := r.FormValue("otp_pass")

                    // Print results
                    uadmin.Trail(uadmin.DEBUG, "Password: %s", password)
                    uadmin.Trail(uadmin.DEBUG, "OTP Password: %s", otpPass)
                }
            }
        }

        // Pass the userContext data object to the HTML file
        uadmin.HTMLContext(w, userContext, "views/login.html")
        return
    }
