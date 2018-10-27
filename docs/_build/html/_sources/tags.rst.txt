Tags
====

What is a tag?
--------------
[#f1]_ A tag is a generic term for a language element descriptor. The set of tags for a document or other unit of information is sometimes referred to as markup, a term that dates to pre-computer days when writers and copy editors marked up document elements with copy editing symbols or shorthand.

In uAdmin, there are two different types of tags: Meta Tags and Type Tags.

What are Meta Tags?
-------------------
Meta tags provide metadata about the uAdmin document that describes some aspect of the contents of a model structure.

There are several kinds of meta tags:

* `categorical_filter`_
* `default_value`_
* `display_name`_
* `filter`_
* `format`_
* `help`_
* `hidden`_
* `limit_choices_to`_
* `list_exclude`_
* `max`_
* `min`_
* `pattern`_
* `pattern_msg`_
* `read_only`_
* `required`_
* `search`_

**categorical_filter**
^^^^^^^^^^^^^^^^^^^^^^
A section of code that is designed to process user input and output request to produce a new data structure containing exactly those elements of the original data structure in the form of combo box.

Syntax:

.. code-block:: go

    `uadmin:"categorical_filter"`

**default_value**
^^^^^^^^^^^^^^^^^
Mainly used in the progress bar on which value you want to initialize.

Syntax:

.. code-block:: go

    `uadmin:"default_value"`

**display_name**
^^^^^^^^^^^^^^^^
A feature to display the data from another model.

Syntax:

.. code-block:: go

    `uadmin:"display_name"`

**filter**
^^^^^^^^^^
A section of code that is designed to process user input and output request to produce a new data structure containing exactly those elements of the original data structure in the form of fill-up text.

Syntax:

.. code-block:: go

    `uadmin:"filter"`

**format**
^^^^^^^^^^
A feature to set the syntax rule to follow by the user.

Syntax:

.. code-block:: go

    `uadmin:"format"`

**help**
^^^^^^^^
A feature that will give a solution to solve advanced tasks.

Syntax:

.. code-block:: go

    `uadmin:"help"`

**hidden**
^^^^^^^^^^
A feature to hide the component in the model structure.

Syntax:

.. code-block:: go

    `uadmin:"hidden"`

**limit_choices_to**
^^^^^^^^^^^^^^^^^^^^
This meta tag has not yet been implemented.

Syntax:

.. code-block:: go

    `uadmin:"limit_choices_to"`

**list_exclude**
^^^^^^^^^^^^^^^^
A feature that will hide the field or column name in the model structure.

Syntax:

.. code-block:: go

    `uadmin:"list_exclude"`

**max**
^^^^^^^
Mainly used in the progress bar to set the maximum value.

Syntax:

.. code-block:: go

    `uadmin:"max"`

**min**
^^^^^^^
Mainly used in the progress bar to set the minimum value.

Syntax:

.. code-block:: go

    `uadmin:"min"`

**pattern**
^^^^^^^^^^^
Equivalent to regular expression that describes a pattern of characters.

Syntax:

.. code-block:: go

    `uadmin:"pattern:(regexp)"`

**pattern_msg**
^^^^^^^^^^^^^^^
Notifies the user once the input has been done following the given pattern.

Syntax:

.. code-block:: go

    `uadmin:"pattern_msg:(message)"`

**read_only**
^^^^^^^^^^^^^
A feature that cannot be modified.

Syntax:

.. code-block:: go

    `uadmin:"read_only"`

**required**
^^^^^^^^^^^^
A section of code that the user must perform the given tasks. It cannot be skipped or left empty.

Syntax:

.. code-block:: go

    `uadmin:"required"`

**search**
^^^^^^^^^^
A feature that allows the user to search for a field or column name.

Syntax:

.. code-block:: go

    `uadmin:"search"`


Where do we use Meta Tags?
--------------------------

Meta tags are used to add extra features on the fields initialized in the model struct. It can be called several times.

Example:

.. code::
   
   type (model_name) struct {
       uadmin.Model
       Name string `uadmin:"required;filter"`
   }

As shown above, required and filter are used meta tags.

What are Type Tags?
-------------------
Type tags are used to specify what type of component should be displayed.

There are several kinds of type tags:

* `code`_
* `email`_
* `file`_
* `html`_
* `image`_
* `link`_
* `m2m`_
* `money`_
* `multilingual`_
* `password`_
* `progress_bar`_

**code**
^^^^^^^^
A set of instructions that will be executed by a computer.

Syntax:

.. code-block:: go

    `uadmin:"code"`

**email**
^^^^^^^^^
It identifies an email box to which email messages are delivered. It follows the syntax as follows: (name)@(domain)

e.g. abc123@gmail.com

Syntax:

.. code-block:: go

    `uadmin:"email"`

**file**
^^^^^^^^
A tag that enables the user to upload files/attachments in the model.

Syntax:

.. code-block:: go

    `uadmin:"file"`

**html**
^^^^^^^^
A tag that allows the user to modify text in HTML format.

Syntax:

.. code-block:: go

    `uadmin:"html"`

.. image:: assets/htmlpic.png

**image**
^^^^^^^^^
A tag to mark a field as an image.

Syntax:

.. code-block:: go

    `uadmin:"image"`

Open your Todo project. Go to your category.go in the models folder and let's use the **`uadmin:"image"`** in the Icon field.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Category model ...
    type Category struct {
	    uadmin.Model
	    Name string `uadmin:"required"`
	    Icon string `uadmin:"image"` // <-- place it here
    }

To run your code:

.. code-block:: bash

    $ cd ~/go/src/github.com/your_name/todo
    $ go build; ./todo
    [   OK   ]   Initializing DB: [10/10]
    [   OK   ]   Server Started: http://127.0.0.1:8000

|

Let's open the category model.

.. image:: tutorial/assets/categorymodelselected.png

|

Create a new data in the category model. Press Save button below afterwards.

.. image:: tutorial/assets/categorywithtagapplied.png

|

Output

.. image:: tutorial/assets/categorydataoutputwithtag.png

|

Now let's do something even cooler. In uAdmin, the image feature will not only just upload your image file but also allows you to crop your own picture through the model itself. In order to that, click the image icon highlighted below.

.. image:: tutorial/assets/iconhighlighted.png

|

Click the crop icon on the top left corner.

.. image:: tutorial/assets/cropiconhighlighted.png

|

You are now set to edit mode. Click any points highlighted below then drag your mouse in order to crop/resize your image.

.. image:: tutorial/assets/croppointshighlighted.png

.. image:: tutorial/assets/croppedicon.png

|

Once you are done, click the Crop button below and refresh the webpage to save your progress.

.. image:: tutorial/assets/croppediconoutput.png

|

Well done! You have mastered the concepts of creating and modifying the image in the model.

**link**
^^^^^^^^
This will set the text in hyperlink format.

Syntax:

.. code-block:: go

    `uadmin:"link"`

**m2m**
^^^^^^^
Many-to-many relationship between two entities

**money**
^^^^^^^^^
This will set the type of currency.

Syntax:

.. code-block:: go

    `uadmin:"money"`

**multilingual**
^^^^^^^^^^^^^^^^
A tag that allows the user to use more than two languages for input.

Syntax:

.. code-block:: go

    `uadmin:"multilingual"`

**password**
^^^^^^^^^^^^
A string of characters that hides the input data for security.

Syntax:

.. code-block:: go

    `uadmin:"password"`

**progress_bar**
^^^^^^^^^^^^^^^^
A feature used for testing the data to check whether the instructions will execute or not.

Syntax (default):

.. code-block:: go

    `uadmin:"progress_bar"` // Any number from 0 to 100 will display blue as the default color.

Syntax (one parameter):

.. code-block:: go

    `uadmin:"progress_bar:100:orange"` // Any number from 0 to 100 will display orange color.

Syntax (multiple parameters):

.. code-block:: go

    `uadmin:"progress_bar:40:red,70:yellow,100:green"` // Any number from 0 to 40 will display red color; 41 to 70 will display yellow color; 71 and above will display green color.

|

Open your Todo project. Go to your main.go and let's use the default tag of the Progress field to **`uadmin:"progress_bar"`** inside the TODO struct.

Copy this code below:

.. code-block:: go

    Progress    int `uadmin:"progress_bar"`

To the todo.go inside the models folder

.. code-block:: go

    package models

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
	    Progress    int `uadmin:"progress_bar"` // <-- place the tag here
    }

|

To run your code:

.. code-block:: bash

    $ cd ~/go/src/github.com/your_name/todo
    $ go build; ./todo
    [   OK   ]   Initializing DB: [9/9]
    [   OK   ]   Server Started: http://127.0.0.1:8000

|

Let's open the Todos model.

.. image:: assets/uadmindashboard.png

|

On the right side, click Add New Todo.

.. image:: assets/todomodel.png

|

Input the progress value to 50 then let's see what happens.

.. image:: assets/todomodelcreate.png

|

Tada! The progress bar is set to 50% with the blue color as the default one.

.. image:: assets/todomodeloutput.png

|

If you want to change the color of the progress bar, let's set a parameter and the value inside the tag. Go back to your main.go again. Let's say I want to display an orange color between the range of 0 to 100. Add this piece of code after the progress_bar tag: **:100:orange** (100 is the value and orange is the parameter)

.. code-block:: go

    // TODO model ...
    type TODO struct {
	    uadmin.Model
	    Name        string
	    Description string `uadmin:"html"`
	    TargetDate  time.Time
	    Progress    int `uadmin:"progress_bar:100:orange"` // <-- place the tag here
    }

|

Run your code again, go to the Todos model in the uAdmin dashboard then replace the value of the progress bar to something like 30.

.. image:: assets/progress30.png

.. image:: assets/progress30output.png

|

If you want some conditions on your progress bar, let's set multiple parameters inside the tag. Let's say I want to display a red color between the range of 0 to 40, yellow color between 41 to 70, and green color between 71 to 100. Add this piece of code after the progress_bar tag: **:40:red,70:yellow,100:green**

.. code-block:: go

    // TODO model ...
    type TODO struct {
	    uadmin.Model
	    Name        string
	    Description string `uadmin:"html"`
	    TargetDate  time.Time
	    Progress    int `uadmin:"progress_bar:40:red,70:yellow,100:green"` // <-- place the tag here
    }

Run your code again, go to the Todos model in the uAdmin dashboard then replace the value of the progress bar to something like 20.

.. image:: assets/progress20.png

.. image:: assets/progress20output.png

|

What if I set the value in the progress bar to 60?

.. image:: assets/progress60.png

.. image:: assets/progress60output.png

|

How about 90?

.. image:: assets/progress90.png

.. image:: assets/progress90output.png

|

Well done! You have mastered the concepts of creating and modifying the progress bar in the model.


Where do we use Type Tags?
--------------------------
Type tags are used to implement the type of component on the fields initialized in the model struct. Unlike in meta tags, type tags can be called only once.

Example:

.. code::
   
   type (model_name) struct {
       uadmin.Model
       Icon string `uadmin:"image"`
   }


References
----------

.. [#f1] Rouse, Margaret (2005, April). Tag. Retrieved from https://searchmicroservices.techtarget.com/definition/tag
