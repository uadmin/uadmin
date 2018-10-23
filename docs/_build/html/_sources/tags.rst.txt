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

categorical_filter
^^^^^^^^^^^^^^^^^^
A section of code that is designed to process user input and output request to produce a new data structure containing exactly those elements of the original data structure in the form of combo box.

default_value
^^^^^^^^^^^^^
Mainly used in the progress bar on which value you want to initialize.

display_name
^^^^^^^^^^^^
A feature to display the data from another model.

filter
^^^^^^
A section of code that is designed to process user input and output request to produce a new data structure containing exactly those elements of the original data structure in the form of fill-up text.

format
^^^^^^
A feature to set the syntax rule to follow by the user. 

help
^^^^
A feature that will give a solution to solve advanced tasks.

hidden
^^^^^^
A feature to hide the component in the model structure.

limit_choices_to
^^^^^^^^^^^^^^^^
This meta tag has not yet been implemented.

list_exclude
^^^^^^^^^^^^
A feature that will hide the field or column name in the model structure.

max
^^^
Mainly used in the progress bar to set the maximum value.

min
^^^
Mainly used in the progress bar to set the minimum value.

pattern
^^^^^^^
Equivalent to regular expression that describes a pattern of characters.

pattern_msg
^^^^^^^^^^^
Notifies the user once the input has been done following the given pattern.

read_only
^^^^^^^^^
A feature that cannot be modified.

required
^^^^^^^^
A section of code that the user must perform the given tasks. It cannot be skipped or left empty.

search
^^^^^^
A feature that allows the user to search for a field or column name.

Where do we use Meta Tags?
--------------------------

Meta tags are used to add extra features on the fields initialized in the model struct. It can be called several times.

Example:

.. code::
   
   type (model_name) struct {
       uadmin.Model
       Name string `uadmin:"required";"filter"`
   }

As shown above, required and filter are used meta tags.

What are Type Tags?
-------------------
Type tags are used to specify what type of component should be displayed.

There are several kinds of type tags:

* `code`_
* `email`_
* `file`_
* `fk`_
* `html`_
* `image`_
* `link`_
* `m2m`_
* `money`_
* `multilingual`_
* `password`_
* `progress_bar`_

code
^^^^
A set of instructions that will be executed by a computer.

email
^^^^^
It identifies an email box to which email messages are delivered. It follows the syntax as follows: (name)@(domain)

e.g. abc123@gmail.com

file
^^^^
A tag that enables the user to upload files/attachments in the model.

fk 
^^^
A foreign key used to link two tables together.

html
^^^^
A tag that allows the user to modify text in HTML format.

image
^^^^^
A tag that enables the user to upload and modify images such as cropping.

link
^^^^
This will set the text in hyperlink format.

m2m
^^^
Many-to-many relationship between two entities

money
^^^^^
This will set the type of currency.

multilingual
^^^^^^^^^^^^
A tag that allows the user to use more than two languages for input.

password
^^^^^^^^
A string of characters that hides the input data for security.

progress_bar
^^^^^^^^^^^^
A feature used for testing the data to check whether the instructions will execute or not.


Where do we use Type Tags?
--------------------------
Type tags are used to implement the type of component on the fields initialized in the model struct. Unlike in meta tags, type tags can be called only once.

Example:

.. code::
   
   type (model_name) struct {
       uadmin.Model
       ProfilePic string `uadmin:"image"`
   }


Citations
---------

.. [#f1] Rouse, Margaret (2005, April). Tag. Retrieved from https://searchmicroservices.techtarget.com/definition/tag
