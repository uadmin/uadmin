Best Practices
==============
In this article, we will learn what the coding standards and naming conventions are in uAdmin using Golang.

**Before we start creating a new project**

* Make it sure it follows the model structure as shown in `Tutorial Part 1`_.
* In every Go file, always remember the pattern:

.. code-block:: go

    // Package name (main if you are in the parent folder, folder_name if you are in the subfolders).

    // Import required libraries (every Go file has github.com/uadmin/uadmin).

    // Input code executions on this part.

.. _Tutorial Part 1: https://uadmin.readthedocs.io/en/latest/tutorial/part1.html

|


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

|

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