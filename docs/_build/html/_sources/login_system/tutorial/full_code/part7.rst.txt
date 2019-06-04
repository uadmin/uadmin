Login System Tutorial Part 7 - Logout (Full Source Code)
========================================================

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
                    }

                    // Pass the requested password and OTP code to return the
                    // session key
                    session := user.Login(password, otpPass)

                    // Check if the session is fetched from the Login function
                    if session != nil {
                        // Create a cookie named "user_session" with the value of
                        // UserID
                        usersession := &http.Cookie{
                            Name:  "user_session",
                            Value: fmt.Sprint(user.ID),
                        }

                        if otp == true && user.VerifyOTP(otpPass) {
                            // Set the "user_session" cookie to the IP Address
                            http.SetCookie(w, usersession)

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
                            // Set the "user_session" cookie to the IP Address
                            http.SetCookie(w, usersession)

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
                    }
                }
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
        }

        // Pass the userContext data object to the HTML file
        uadmin.HTMLContext(w, userContext, "views/login.html")
        return
    }
