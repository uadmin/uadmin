Login System Tutorial Part 1 - Build A Project
==============================================
In this part, we will cover on building and preparing a project from scratch.

First of all, let's create a folder for your project and prepare it.

.. code-block:: bash

    $ mkdir -p ~/go/src/github.com/your_name/login_system
    $ cd ~/go/src/github.com/your_name/login_system
    $ uadmin prepare
    [   OK   ]   Created: /home/pc_name/go/src/github.com/your_name/login_system/models
    [   OK   ]   Created: /home/pc_name/go/src/github.com/your_name/login_system/api
    [   OK   ]   Created: /home/pc_name/go/src/github.com/your_name/login_system/views
    [   OK   ]   Created: /home/pc_name/go/src/github.com/your_name/login_system/media
    [   OK   ]   Created: /home/pc_name/go/src/github.com/your_name/login_system/handlers
    [   OK   ]   Created: /home/pc_name/go/src/github.com/your_name/login_system/static
    [   OK   ]   Created: /home/pc_name/go/src/github.com/your_name/login_system/templates

Use your favorite editor to create “main.go” inside that path. Put the following code in “main.go”.

.. code-block:: go

    package main

    import (
        "github.com/uadmin/uadmin"
    )

    func main() {
        // The listener is mapped to /admin/ path in the URL.
        uadmin.RootURL = "/admin/"

        // Sets the name of the website that shows on title and dashboard
        uadmin.SiteName = "Login System"

        // Run the server
        uadmin.StartServer()
    }

.. IMPORTANT::
   In Windows, you need to use localhost in order to run your application (e.g. http://localhost:8080). Another way is to set your loopback Internet protocol (IP) address by using uadmin.BindIP to establish an IP connection to the same machine or computer being used by the end-user.

Sample:

.. code-block:: go

    func main(){
        // Put this code before uadmin.StartServer
        uadmin.BindIP = "127.0.0.1"
    }

Now to run your code (Linux and Apple macOS):

.. code-block:: bash

    $ go build; ./login_system
    [   OK   ]   Initializing DB: [8/8]
    [   OK   ]   Initializing Languages: [185/185]
    [  INFO  ]   Auto generated admin user. Username: admin, Password: admin.
    [   OK   ]   Server Started: http://0.0.0.0:8080
             ___       __          _
      __  __/   | ____/ /___ ___  (_)___
     / / / / /| |/ __  / __  __ \/ / __ \
    / /_/ / ___ / /_/ / / / / / / / / / /
    \__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/

In Windows:

.. code-block:: bash

    > go build && login_system.exe
    [   OK   ]   Initializing DB: [8/8]
    [   OK   ]   Initializing Languages: [185/185]
    [  INFO  ]   Auto generated admin user. Username: admin, Password: admin.
    [   OK   ]   Server Started: http://0.0.0.0:8080
             ___       __          _
      __  __/   | ____/ /___ ___  (_)___
     / / / / /| |/ __  / __  __ \/ / __ \
    / /_/ / ___ / /_/ / / / / / / / / / /
    \__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/

Open your browser and type the IP address above including the path that you have assigned in RootURL (e.g. http://0.0.0.0:8080/admin/). Then login using “admin” as username and password.

.. image:: assets/loginform.png

|

You will be greeted by the Login System dashboard that contains the system models built in uAdmin.

.. image:: assets/loginsystemdashboard.png

|

In the `next part`_, we will discuss about creating a login form in HTML.

.. _next part: https://uadmin.readthedocs.io/en/latest/login_system/tutorial/part2.html