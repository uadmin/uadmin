uAdmin Tutorial Part 1 - Build A Project
========================================

In this part, we will build upon the Todo list from home.


Folder Structure
^^^^^^^^^^^^^^^^

There is no required folder structure, but from experience we found the following structure 
to work and scale well for a uAdmin projects:

.. code-block:: bash

    ~/go/src/github.com/your_name/project/
        models/     # Automatically Generated - DB models
            a.go
            b.go
        api/        # Automatically Generated - Custom API
            api.go
            some_handler.go
        views/      # Automatically Generated - HTML Files
            index.html      # Home page
            view.html       
            some_view.html
        media/      # Automatically Generated - User uploads
            files           # This is where your uploaded files are stored.
            images          # This is where your uploaded images are stored.
            otp             # This is where your OTPs are stored in PNG format.
        handlers/   # Automatically Generated - Custom UI HTTP Handlers
            handler.go
            some_handler.go
        static/     # Automatically Generated - Static files (images/js/css...)
            i18n            # This is where the JSON files for translation are stored.
            uadmin          # Built-in files for uadmin
        templates/  # Automatically Generated - HTML Templates
        main.go

The first time you run your project, these folders are automatically generated for you.

**Models** is where your external models are located. In order to access those models in the dashboard, `uadmin.Register`_ function is necessary to be done in main.go which will be discussed in the next part of this tutorial. 

.. _uadmin.Register: https://uadmin.readthedocs.io/en/latest/api.html#uadmin-register

**API** is where back-end and front-end will communicate. Either the records stored in uAdmin database or the data that was stored in AJAX function from Javascript will pass to the JSON based on the query by matching the fields from a specific model and their values. AJAX call will perform the tasks.

**Views** is where your HTML files are located. In order to get the data from model(s) to HTML, you need to use Golang delimiters which will be discussed in the later part of this tutorial.

**Media** is where your uploaded multimedia files are located. It can be files, images, OTPs, sounds, and many others.

**Handlers** is where back-end and front-end will communicate. Unlike in API, handlers does not use a query or JSON to store data. Instead, it creates a replicate model struct based on uAdmin model(s) to store values on each field then pass the context variable into the ExecuteTemplate function which will be discussed in the later part of this tutorial.

**Static** is where the built-in files for uAdmin such as JSON for translation are located.

**Templates** is where the themes for uAdmin are located.

In the `next part`_ we will talk about creating and moving a model as well as cropping an image.

.. _next part: https://uadmin.readthedocs.io/en/latest/tutorial/part2.html

