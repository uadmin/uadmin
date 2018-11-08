uAdmin the Golang Web Framwork
==============================

uAdmin is easy to use, blazing fast and secure. It is a simple yet powerful web framework for building web applications.

Installation
^^^^^^^^^^^^

.. code-block:: bash

    $ go get -u github.com/uadmin/uadmin/...

To test if your installation is fine, run the **uadmin** command line:

.. code-block:: bash

    $ uadmin
    Usage: uadmin COMMAND [-e email] [-d domain]
    This tools allows you to publish your project online

    Commands:
    publish         This publishes your project online
    prepare         Generates folders and prepares static and templates

    Arguments:
    -e, --email     Your email. This is required for you to be able to maintain your project.
    -d, --domain    You can choose your domain name which will customize your URL

    Get full documentation online:
    https://uadmin.io/docs/

Your First App
^^^^^^^^^^^^^^^
Let's build your first app which is a Todo list. First, we will create a folder for your project and prepare it.

.. code-block:: bash

    $ mkdir -p ~/go/src/github.com/your_name/todo
    $ cd ~/go/src/github.com/your_name/todo
    $ uadmin prepare
    [   OK   ]   Created: /home/abdullah/go/src/github.com/uadmin/todo1/models
    [   OK   ]   Created: /home/abdullah/go/src/github.com/uadmin/todo1/api
    [   OK   ]   Created: /home/abdullah/go/src/github.com/uadmin/todo1/views
    [   OK   ]   Created: /home/abdullah/go/src/github.com/uadmin/todo1/media
    [   OK   ]   Created: /home/abdullah/go/src/github.com/uadmin/todo1/static
    [   OK   ]   Created: /home/abdullah/go/src/github.com/uadmin/todo1/templates

Use your favorite editor to create "main.go" inside that path. Put the
following code in "main.go".

.. code-block:: go

    package main

    import (
        "github.com/uadmin/uadmin"
        "time"
    )

    // Todo model ...
    type Todo struct {
	    uadmin.Model
	    Name        string
	    Description string `uadmin:"html"`
	    TargetDate  time.Time
	    Progress    int `uadmin:"progress_bar"`
    }

    func main() {
	    uadmin.Register(Todo{})
	    uadmin.StartServer()
    }


Now to run your code:

.. code-block:: bash

    $ go build; ./todo
    [   OK   ]   Initializing DB: [9/9]
    [   OK   ]   Initializing Languages: [185/185]
    [  INFO  ]   Auto generated admin user. Username: admin, Password: admin.
    [   OK   ]   Server Started: http://0.0.0.0:8080
             ___       __          _
      __  __/   | ____/ /___ ___  (_)___
     / / / / /| |/ __  / __  __ \/ / __ \
    / /_/ / ___ / /_/ / / / / / / / / / /
    \__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/

Open your browser and type the IP address above. Then login using “admin” as username and password.

.. image:: tutorial/assets/loginform.png

|

You will be greeted by the uAdmin dashboard. System models are built in to uAdmin, and the rest are the ones we created, in this case TODO model.

.. image:: assets/uadmindashboard.png

|

Open the TODO model and add a new TODO.

.. image:: assets/todomodel.png

|

Fill up the fields like in the example below:

.. image:: assets/todomodelcreate.png

|

Save it and new data will be added to your model.

.. image:: assets/todomodeloutput.png

Well done! You have created your first application.

.. toctree::
   :maxdepth: 1

   getting_started
   tutorial/part1
   tutorial/part2
   tutorial/part3
   tutorial/part4
   tutorial/part5
   tutorial/part6
   tutorial/part7
   api
   quick_reference
   tags
   bestpractices
   profile
   about
   license
   roadmap