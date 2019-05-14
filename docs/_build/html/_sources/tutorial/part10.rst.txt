uAdmin Tutorial Part 10 - Introduction to HTML Template
=======================================================
In this part, we will discuss about designing a table in HTML and setting up a template file.

Before you proceed, make sure you have at least the basic knowledge of HTML. If you are not familiar with HTML, we advise you to go over `W3Schools`_.

.. _W3Schools: https://www.w3schools.com/

In this tutorial, we will use Bootstrap 4. For the tutorials, click `here`_.

.. _here: https://www.w3schools.com/bootstrap4/default.asp

First of all, go to your project folder and select views.

.. image:: assets/viewsfolderhighlighted.png
   :align: center

|

Inside the views folder, create a new file named **todo.html**.

.. image:: assets/todohtmlcreate.png
   :align: center

|

Inside the todo.html, create an HTML5 structure following the codes below and change the title from Document to Todo List.

.. code-block:: html

    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">

        <!-- Latest compiled and minified CSS -->
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css">

        <!-- Change the title from Document to Todo List -->
        <title>Todo List</title> 
    </head>
    <body>
        
    </body>
    </html>

Save the file. Run your application in the browser and see what happens.

.. image:: assets/todolisthtmltitle.png

|

The title bar is named as Todo List. Now inside the <body>, create a table header following the code structure below. You can choose which class of Bootstrap table that you want to display in your application. In this tutorial, we will use table-striped.

.. code-block:: html

    <div class="container-fluid">
        <table class="table table-striped">
            <!-- Todo Fields -->
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Description</th>
                    <th>Category</th>
                    <th>Friend</th>
                    <th>Item</th>
                    <th>Target Date</th>
                    <th>Progress</th>
                </tr>
            </thead>
            <tbody>

            </tbody>
        </table>
    </div>

Save the file. Run your application in the browser and see what happens.

.. image:: assets/todolisthtmlheader.png

|

Nice! Now go back to your project folder then select handlers.

.. image:: assets/handlersfolderhighlighted.png
   :align: center

|

Inside the handlers folder, create a new file named **handler.go**.

.. image:: assets/handlergofile.png

|

In the `next part`_, we will talk about establishing a connection to the HTTP Handler, setting the URL path name, and executing an HTML file.

.. _next part: https://uadmin.readthedocs.io/en/latest/tutorial/part11.html

