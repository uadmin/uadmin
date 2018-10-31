Simple Web Framework for Golang
===============================

uAdmin is a simple yet powerful web framework for building web applications.

Installation
^^^^^^^^^^^^

Install uAdmin:

.. code-block:: bash

    $ go get github.com/uadmin/uadmin/...

Check if the installation went well.

.. code-block:: bash

    $ uadmin

Expected Result

.. image:: assets/uadmin.png

Your First Project
^^^^^^^^^^^^^^^^^^

Once you have uAdmin installed, let's start a project.

Note: the last directory is your project name, in this case we named it todo.

.. code-block:: bash

    $ mkdir -p ~/go/src/github.com/your_name/todo
    $ cd ~/go/src/github.com/your_name/todo
    $ uadmin prepare

Expected output

.. image:: assets/uadminprepareoutput.png

Use your favorite editor to create "main.go" inside that path. Put the
following code in "main.go".

.. code-block:: go

    package main

    import (
	    "time"
	    "github.com/uadmin/uadmin"
    )

    // TODO model ...
    type TODO struct {
	    uadmin.Model
	    Name        string
	    Description string `uadmin:"html"`
	    TargetDate  time.Time
	    Progress    int `uadmin:"progress_bar"`
    }

    func main() {
	    uadmin.Register(TODO{})
	    uadmin.StartServer()
    }


To run your code:

.. code-block:: bash

    $ cd ~/go/src/github.com/your_name/todo
    $ go build; ./todo
    [   OK   ]   Initializing DB: [9/9]
    [   OK   ]   Server Started: http://127.0.0.1:8080

Open your browser and type the IP address above. Then login using “admin” as username and password.

.. image:: assets/loginform.png

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
   api
   about
   roadmap
   license
   tags
