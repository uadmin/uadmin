uAdmin Tutorial Part 1 - Build A Project
========================================

In this part, we will build upon the Todo list from home.


Folder Structure
^^^^^^^^^^^^^^^^

There is no required folder structure, but from experience we found the following structure 
to work and scale well for a uAdmin projects:

.. code-block:: go

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
        static/     # Automatically Generated - Static files (images/js/css...)
        templates/  # Automatically Generated - html templates
        main.go

The first time you run your project, these folders are automatically generated for you.