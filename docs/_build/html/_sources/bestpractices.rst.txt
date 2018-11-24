Best Practices
==============
In this section, we will learn what are the coding standards and naming conventions are in uAdmin using Golang.

**Before we start creating a new project**

* Make it sure it follows the model structure as shown in `Tutorial Part 1`_.
* In every Go file, always remember the pattern:

.. code-block:: go

    // Package name (main if you are in the parent folder, folder_name if you are in the subfolders).

    // Import required libraries (every Go file has github.com/uadmin/uadmin).

    // Input code executions on this part.

.. _Tutorial Part 1: https://uadmin.readthedocs.io/en/latest/tutorial/part1.html

* Inside the function of main.go, follow this pattern:

.. code-block:: go

    // uAdmin Global Configuration

    // Register

    // RegisterInlines

    // API Handlers

    // StartServer

* Never ever tamper the files inside the static folder as it may cause a widespread problem to your application.
* In terms of publishing, static files are not included.
* Suppose that you have created your own layout. Our server cannot access static files such as HTML/CSS/JS/Images. If you go to that path, it only reads the plain text. In order to serve your static files into your layout, establish a handler in main.go by using http.Handle to access them with the syntax as shown below:

.. code-block:: go

    http.Handle("/assets/folder_name/", http.StripPrefix("/assets/folder_name/", http.FileServer(http.Dir("./assets/folder_name/"))))

**Before we start creating a model**

* The first letter of a model name must be in uppercase format.
* Make sure the model name in your project is in singular form because if you run the application, the uAdmin Dashboard will automatically generate a model name in plural form.
* It is advisable to always put the model struct externally instead of joining it in the main.go. It can be called by using uadmin.Register.
* In every struct, uadmin.Model must always come first before creating a field.

.. code-block:: go

    type (struct_name) struct{
        uadmin.Model
        // Some codes here
    }

**Before we start creating an inline**

* Inline must be created in main.go after the uadmin.Register.
* If you have more than one inline in your project, always match the declared element in a single line inside the uadmin.RegisterInlines for easier debugging process as shown below.

.. code-block:: go

    uadmin.RegisterInlines(/folder_name/./struct_name of a parent model/{}, map[string]string{
        "/sub_model name/": "/parent_model name/ID", // first element
        "/sub_model name/": "/parent_model name/ID", // second element
        "/sub_model name/": "/parent_model name/ID", // third element
    })

Example:

.. code-block:: go

    uadmin.RegisterInlines(models.TODOID{}, map[string]string{
        "Category": "TODOID",
        "Friend": "TODOID",
        "Item": "TODOID",
    })

**For the login process**

* Avoid using common passwords such as "123456" and "password". Use a password that contains an uppercase and lowercase letters, numbers, and special symbols for strong security.
* Enable two factor authentication in your user account. 2FA adds an extra layer of security that makes it harder for an attacker to access your data.
* Always set an email address in the user account just in case if he forgots his password.
* Getting the User through `IsAuthenticated`_ function
* For every password field in the model, apply "`encrypt`_" tag to protect the user from security attacks.
* You can also apply `uadmin.GenerateBase32`_, `uadmin.GenerateBase64`_, or `uadmin.Salt`_ as an alternative way to secure the user's password.

.. _IsAuthenticated: https://uadmin.readthedocs.io/en/latest/api.html#uadmin-isauthenticated
.. _encrypt: https://uadmin.readthedocs.io/en/latest/tags.html#encrypt
.. _uadmin.GenerateBase32: https://uadmin.readthedocs.io/en/latest/api.html#uadmin-generatebase32
.. _uadmin.GenerateBase64: https://uadmin.readthedocs.io/en/latest/api.html#uadmin-generatebase64
.. _uadmin.Salt: https://uadmin.readthedocs.io/en/latest/api.html#uadmin-salt

**For using the logs**

* Edit and Delete logs will allow you to “Undo” them or “Roll Back” your changes. It is a good feature for the user who accidentally made changes to the record in the database.
* When you access to any records you have in your system, there is a "History" button which redirects you to the Log on the top left corner.
* You can use "Filter" to narrow down what you are looking for. This is useful if your log has too many records in your system.