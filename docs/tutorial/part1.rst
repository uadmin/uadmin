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
        views/      # Automatically Generated - Custom UI http handlers
            view.go
            some_view.go
        media/      # Automatically Generated - User uploads
            files           # This is where your uploaded files are stored.
            images          # This is where your uploaded images are stored.
            otp             # This is where your OTPs are stored in PNG format.
        static/     # Automatically Generated - Static files (images/js/css...)
            i18n            # This is where the JSON files for translation are stored.
            uadmin          # Built-in files for uadmin
        templates/  # Automatically Generated - html templates
        main.go

The first time you run your project, these folders are automatically generated for you.

In the `next part`_ we will talk about creating a model and how to crop an image.

.. _next part: https://uadmin.readthedocs.io/en/latest/tutorial/part2.html