Tag Reference
=============

What is a tag?
--------------
A tag is a generic term for a language element descriptor. The set of tags for a document or other unit of information is sometimes referred to as markup, a term that dates to pre-computer days when writers and copy editors marked up document elements with copy editing symbols or shorthand. [#f1]_

In uAdmin, there are two different types of tags: `Meta Tags`_ and `Type Tags`_.

Meta Tags vs. Type Tags
-----------------------
**Meta tags** provide metadata about the uAdmin document that describes some aspect of the contents of a model structure.

Meta tags are used to add extra features on the fields initialized in the model structure. Each field can have multiple meta tags.

Example:

.. code::
   
   type (model_name) struct {
       uadmin.Model
       Name string `uadmin:"required;filter"`
   }

As shown above, required and filter are used meta tags.

**Type tags** are used to specify what type of component should be displayed.

Type tags are used to implement the type of component on the fields initialized in the model structure. Unlike in meta tags, type tags can be called only once.

Example:

.. code::
   
   type (model_name) struct {
       uadmin.Model
       Icon string `uadmin:"image"`
   }

|

.. list-table:: **LIST OF UADMIN TAGS**
   :widths: 15 30 15
   :align: center
   :header-rows: 1

   * - Meta Tags
     -
     - Type Tags
   * - * `categorical_filter`_
     -
     - * `code`_
   * - * `default_value`_
     -
     - * `email`_
   * - * `display_name`_
     -
     - * `file`_
   * - * `encrypt`_
     -
     - * `html`_
   * - * `filter`_
     -
     - * `image`_
   * - * `format`_
     -
     - * `link`_
   * - * `help`_
     -
     - * `money`_
   * - * `hidden`_
     -
     - * `multilingual`_
   * - * `list_exclude`_
     -
     - * `password`_
   * - * `max`_
     -
     - * `progress_bar`_
   * - * `min`_
     -
     -
   * - * `pattern`_
     - 
     -
   * - * `pattern_msg`_
     - 
     -
   * - * `read_only`_
     - 
     -
   * - * `required`_
     - 
     -
   * - * `search`_
     - 
     -

Meta Tags
---------

**categorical_filter**
^^^^^^^^^^^^^^^^^^^^^^
A section of code that is designed to process user input and output request to produce a new data structure containing exactly those elements of the original data structure in the form of combo box

Syntax:

.. code-block:: go

    `uadmin:"categorical_filter"`

Open your Todo List project, go to the items.go and set the categorical_filter tag in the Name field.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Item model ...
    type Item struct {
        uadmin.Model
        Name        string `uadmin:"categorical_filter"`
        Description string
        Cost        int
        Rating      int
    }

Let's run the application to see the output.

.. image:: assets/categoricalfilteroutput.png

**default_value**
^^^^^^^^^^^^^^^^^
Mainly used in the input field on which value you want to initialize

Syntax:

.. code-block:: go

    `uadmin:"default_value"`

Open your Todo List project, go to the items.go and set the default_value tag in the Name field. Let's say "Computer".

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Item model ...
    type Item struct {
        uadmin.Model
        Name        string `uadmin:"default_value:Computer"` // <-- place it here
        Description string
        Cost        int
        Rating      int
    }

Let's run the application to see the output.

.. image:: assets/defaultvaluetagapplied.png

**display_name**
^^^^^^^^^^^^^^^^
A feature to set the actual name in the field

Syntax:

.. code-block:: go

    `uadmin:"display_name"`

Open your Todo List project, go to the items.go and set the display_name tag in the Name field. Let's say "Product Name".

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Item model ...
    type Item struct {
        uadmin.Model
        Name        string `uadmin:"display_name:Product Name"` // <-- place it here
        Description string
        Cost        int
        Rating      int
    }

Let's run the application to see the output.

.. image:: assets/displaynametagapplied.png

**encrypt**
^^^^^^^^^^^
This meta tag encrypts the input field in the record. It was released in version 0.1.0-beta.3.

Syntax:

.. code-block:: go

    `uadmin:"encrypt"`

Add a record in the Friend model. Notice that the password you have inputed is 123456.

.. image:: assets/addrecordinfriendmodel.png

|

Go to the Friend model and apply the tag as "encrypt" in the Password field.

.. code-block:: go

    // Friend model ...
    type Friend struct {
    uadmin.Model
        Name     string 
        Email    string 
        Password string `uadmin:"encrypt"` // <- place it here
    }

Now rerun your application, refresh your browser and see what happens.

.. image:: assets/passwordgone.png

|

The password is invisible now. Go to your project folder, open uadmin.db file, go to Browse Data tab, and you will notice that the password field is encrypted.

.. image:: assets/sqlitepasswordencrypt.png

|

Remove the encrypt tag in the Friend model, rerun your application and see what happens.

.. image:: assets/addrecordinfriendmodel.png

|

The password is shown again which means it is decrypted.

**filter**
^^^^^^^^^^
A section of code that is designed to process user input and output request to produce a new data structure containing exactly those elements of the original data structure in the form of fill-up text

Syntax:

.. code-block:: go

    `uadmin:"filter"`

Open your Todo List project, go to the item.go and set the filter tag in the Name field.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Item model ...
    type Item struct {
        uadmin.Model
        Name        string `uadmin:"filter"` // <-- place it here
        Description string
        Cost        int
        Rating      int
    }

Run your application. Click the filter button on the upper right.

.. image:: tutorial/assets/filtertagapplied.png

|

Now let's filter the word "iPad" and see what happens.

.. image:: tutorial/assets/filtertagappliedoutput.png

**format**
^^^^^^^^^^
A feature to set the syntax rule to follow by the user

Syntax:

.. code-block:: go

    `uadmin:"format"`

**help**
^^^^^^^^
A feature that will give a solution to solve advanced tasks

Syntax:

.. code-block:: go

    `uadmin:"help"`

Open your Todo List project, go to the item.go and set the help tag in the Name field. Let's say "Input numeric characters only in this field.".

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Item model ...
    type Item struct {
        uadmin.Model
        Name        string
        Description string
        Cost        int `uadmin:"help:Input numeric characters only in this field."` // <-- place it here
        Rating      int
    }

Let's run the application to see the output.

.. image:: assets/helptagapplied.png

**hidden**
^^^^^^^^^^
A feature to hide the component in the editing section of the data

Syntax:

.. code-block:: go

    `uadmin:"hidden"`

Open your Todo List project, go to the todo.go and set the hidden tag in the CreatedAt field.

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
        Description string
        CreatedAt   time.Time `uadmin:"hidden"` // <-- place it here
        TargetDate  time.Time
        Progress    int
    }

Let's run the application to see the output.

.. image:: assets/hiddentagapplied.png

CreatedAt does not show up in the editing section of the data because it is set as "hidden".

**list_exclude**
^^^^^^^^^^^^^^^^
A feature that will hide the field or column name in the model structure

Syntax:

.. code-block:: go

    `uadmin:"list_exclude"`

Open your Todo List project, go to the friend.go and set the list_exclude tag in the Password field.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Friend model ...
    type Friend struct {
        uadmin.Model
        Name     string
        Email    string
        Password string `uadmin:"list_exclude"` // <-- place it here
    }

Let's run the application to see the output.

.. image:: assets/listexcludetagapplied.png

Password does not show up in the model structure because it is set as "list_exclude".

**max**
^^^^^^^
Mainly used in the input field to set the maximum value

Syntax:

.. code-block:: go

    `uadmin:"max"`

Open your Todo List project, go to the item.go and set the max tag in the Rating field. Let's say 5.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Item model ...
    type Item struct {
        uadmin.Model
        Name        string
        Description string
        Cost        int
        Rating      int `uadmin:"max:5"` // <-- place it here
    }

Let's run the application to see the output.

.. image:: assets/maxtagapplied.png

It returns an error because the value is greater than 5 which is the maximum value allowed.

**min**
^^^^^^^
Mainly used in the input field to set the minimum value

Syntax:

.. code-block:: go

    `uadmin:"min"`

Open your Todo List project, go to the item.go and set the min tag in the Rating field. Let's say 1.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Item model ...
    type Item struct {
        uadmin.Model
        Name        string
        Description string
        Cost        int
        Rating      int `uadmin:"min:1"` // <-- place it here
    }

Let's run the application to see the output.

.. image:: assets/mintagapplied.png

It returns an error because the value is lesser than 1 which is the minimum value allowed.

**pattern**
^^^^^^^^^^^
Equivalent to regular expression that describes a pattern of characters

Syntax:

.. code-block:: go

    `uadmin:"pattern:(regexp)"`

Open your Todo List project, go to the item.go and set the pattern tag in the Cost field. Let's say ^[0-9]*$. This accepts numeric characters only.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Item model ...
    type Item struct {
        uadmin.Model
        Name        string
        Description string
        Cost        int `uadmin:"pattern:^[0-9]*$"` // <-- place it here
        Rating      int
    }

Let's run the application and see what happens.

.. image:: assets/patterntagapplied.png

|

Output

.. image:: assets/patterntagappliedoutput.png

**pattern_msg**
^^^^^^^^^^^^^^^
Notifies the user once the input has been done following the given pattern

Syntax:

.. code-block:: go

    `uadmin:"pattern_msg:(message)"`

Open your Todo List project, go to the item.go and set the pattern tag in the Cost field. Let's say "Your input must be a number.". This accepts numeric characters only.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Item model ...
    type Item struct {
        uadmin.Model
        Name        string
        Description string
        Cost        string `uadmin:"pattern:^[0-9]*$;pattern_msg:Your input must be a number."` // <-- place it here
        Rating      int
    }

Let's run the application and see what happens.

.. image:: assets/patternmsgtagapplied.png

It returns an error because the input value has letters and special symbols.

**read_only**
^^^^^^^^^^^^^
A feature that cannot be modified

Syntax:

.. code-block:: go

    `uadmin:"read_only"`

Open your Todo List project, go to the todo.go and set the read_only tag in the CreatedAt field.

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
        Description string 
        CreatedAt   time.Time `uadmin:"read_only"` // <-- place it here
        TargetDate  time.Time
        Progress    int
    }

Let's run the application to see the output.

.. image:: assets/readonlytagapplied.png

**required**
^^^^^^^^^^^^
A section of code that the user must perform the given tasks. It cannot be skipped or left empty.

Syntax:

.. code-block:: go

    `uadmin:"required"`

Open your Todo List project, go to the category.go and set the required tag in the Name field.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Category model ...
    type Category struct {
        uadmin.Model
        Name string `uadmin:"required"` // <-- place it here
        Icon string
    }

Let's run the application to see the output.

.. image:: assets/requiredtagapplied.png

It returns an error because the input value is empty. * symbol indicates that the Name field is required to fill up.

**search**
^^^^^^^^^^
A feature that allows the user to search for a field or column name

Syntax:

.. code-block:: go

    `uadmin:"search"`

Before we proceed, add more data in your items model. Once you are done, let's add the "search" tag in the name field of items.go and see what happens.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Items model ...
    type Items struct {
	    uadmin.Model
	    Name        string `uadmin:"search"` // <-- place it here
	    Description string
	    Cost        int
	    Rating      int
    }

Result

.. image:: tutorial/assets/searchtagapplied.png

|

Search the word "mini" and see what happens.

.. image:: tutorial/assets/searchtagappliedoutput.png

Type Tags
---------

**code**
^^^^^^^^
A set of instructions that will be executed by a computer

Syntax:

.. code-block:: go

    `uadmin:"code"`

Go to the friend.go and apply the following codes below:

.. code-block:: go

    // Friend model ...
    type Friend struct {
        uadmin.Model
        Name     string `uadmin:"required"`
        Email    string `uadmin:"email"`
        Password string `uadmin:"password;list_exclude"`
        Message  string `uadmin:"code"`     // <-- place it here
    }

    // Save !
    func (f *Friend) Save() {
        // Initialize two variables
        x := 5
        y := 3

        // Execution code. strconv.Itoa means converting from int to string.
        f.Message = "Hi, I'm " + f.Name + ". Can you solve " + strconv.Itoa(x) + " + " + strconv.Itoa(y) + " for me? The answer is " + strconv.Itoa(x+y) + "."

        // Override save
        uadmin.Save(f)
    }

Now let's run the application, go to the Friend model, create a record, save then let's see the result.

.. image:: assets/codetagapplied.png

Well done! The execution code has performed successfully in the message field.

**email**
^^^^^^^^^
It identifies an email box to which email messages are delivered. It follows the syntax as follows: (name)@(domain).

e.g. abc123@gmail.com

Syntax:

.. code-block:: go

    `uadmin:"email"`

Open your Todo List project, go to the friend.go and set the email tag in the Email field.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Friend model ...
    type Friend struct {
	    uadmin.Model
	    Name     string
	    Email    string `uadmin:"email"` // <-- place it here
	    Password string
    }

Let's run the application to see the output.

.. image:: assets/emailtagapplied.png

It returns an error because the input value does not follow the email format.

**file**
^^^^^^^^
A tag that enables the user to upload files/attachments in the model

Syntax:

.. code-block:: go

    `uadmin:"file"`

Go to the category.go and apply the following codes below:

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Category model ...
    type Category struct {
        uadmin.Model
        Name string `uadmin:"required"`
        Icon string `uadmin:"image"`
        File string `uadmin:"file"` // <-- place it here
    }

Now run your application. Go to the Category model. In File field, you can upload any type of files in the model.

.. image:: assets/filetagapplied.png

|

Now click the filename and see what happens.

.. image:: assets/filetagappliedoutput.png

|

Result

.. image:: assets/filetagappliedresult.png

**html**
^^^^^^^^
A tag that allows the user to modify text in HTML format

Syntax:

.. code-block:: go

    `uadmin:"html"`

Open your Todo List project, go to the todo.go and set the html tag in the Description field.

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
        Description string `uadmin:"html"` // <-- place it here
        TargetDate  time.Time
        Progress    int
    }

Let's run the application to see the output.

.. image:: assets/htmlpic.png

|

HTML has a source code feature that allows you to modify your own code through the application itself.

.. image:: assets/sourcecodehighlighted.png

|

Add this piece of code in the source code editor. This will create a bulleted unordered list.

.. image:: assets/addedulhighlighted.png

Result

.. image:: assets/addeduloutput.png

**image**
^^^^^^^^^
A tag to mark a field as an image

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

Let's open the category model.

.. image:: tutorial/assets/categorymodelselected.png

|

Create a new data in the category model. Press Save button below afterwards.

.. image:: tutorial/assets/categorywithtagapplied.png

|

Result

.. image:: tutorial/assets/categorydataoutputwithtag.png

|

uAdmin also allows you to crop your images.

.. image:: tutorial/assets/cropiconhighlighted.png

.. image:: tutorial/assets/croppedicon.png

Once you are done, click the Crop button below and refresh the webpage to save your progress.

**link**
^^^^^^^^
This type will display a button in the model.

Syntax:

.. code-block:: go

    `uadmin:"link"`

Let's add an Invite field in the friend.go that will direct you to his website. In order to do that, set the field name as "Invite" with the tag "link".

.. code-block:: go

    // Friend model ...
    type Friend struct {
        uadmin.Model
        Name        string 
        Email       string 
        Password    string 
        Nationality string
        Invite      string `uadmin:"link"` // <-- place it here
    }

To make it functional, add the overriding save function after the Friend struct.

.. code-block:: go

    // Save !
    func (f *Friend) Save() {
        f.Invite = "https://uadmin.io/"
        uadmin.Save(f)
    }

Run your application, go to the Friends model and update the elements inside. Afterwards, click the Invite button on the output structure and see what happens.

.. image:: tutorial/assets/invitebuttonhighlighted.png

|

Result

.. image:: tutorial/assets/uadminwebsitescreen.png

**money**
^^^^^^^^^
This will set the type of currency.

Syntax:

.. code-block:: go

    `uadmin:"money"`

Open your Todo List project, go to the item.go and set the money tag in the Cost field.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Item model ...
    type Item struct {
        uadmin.Model
        Name        string
        Description string
        Cost        int `uadmin:"money"` // <-- place it here
        Rating      int
    }

Let's run the application and see what happens.

.. image:: assets/moneytagapplied.png

**multilingual**
^^^^^^^^^^^^^^^^
A tag that allows the user to use more than two languages for input

Syntax:

.. code-block:: go

    `uadmin:"multilingual"`

Open your Todo List project, go to the item.go and set the multilingual tag in the Description field.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Item model ...
    type Item struct {
        uadmin.Model
        Name        string
        Description string `uadmin:"multilingual"` // <-- place it here
        Cost        int
        Rating      int
    }

Let's run the application and see what happens.

.. image:: assets/multilingualtagapplied.png

|

If you want to add more languages in your model, go to the Languages in the uAdmin dashboard.

.. image:: tutorial/assets/languageshighlighted.png

|

Let's say I want to add Chinese and Tagalog in the Item model. In order to do that, set the Active as enabled.

.. image:: tutorial/assets/activehighlighted.png

|

Now go back to the Item model and see what happens.

.. image:: tutorial/assets/multilingualtagappliedmultiple.png

As expected, Chinese and Tagalog languages were added in the Description field.

To customize your own languages, click `here`_ for the instructions.

.. _here: https://medium.com/@twistedhardware/uadmin-the-golang-web-framework-4-customizing-dashboard-d96d90792a07

**password**
^^^^^^^^^^^^
A string of characters that hides the input data for security

Syntax:

.. code-block:: go

    `uadmin:"password"`

Open your Todo List project, go to the friend.go and set the password tag in the Password field.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Friend model ...
    type Friend struct {
        uadmin.Model
        Name     string
        Email    string
        Password string `uadmin:"password"` // <-- place it here
    }

Let's run the application to see the output.

.. image:: assets/passwordtagapplied.png

In this case, the string of characters will hide every time you input something in the password field. If you want to show your input, click the eye icon button on the right side highlighted above.

**progress_bar**
^^^^^^^^^^^^^^^^
A feature used for testing the data to check whether the instructions will execute or not

Syntax (default):

.. code-block:: go

    `uadmin:"progress_bar"` // Any number from 0 to 100 will display blue as the default color.

Syntax (one parameter):

.. code-block:: go

    `uadmin:"progress_bar:100:orange"` // Any number from 0 to 100 will display orange color.

Syntax (multiple parameters):

.. code-block:: go

    `uadmin:"progress_bar:40:red,70:yellow,100:green"` // Any number from 0 to 40 will display red color; 41 to 70 will display yellow color; 71 and above will display green color.

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

To run your code:

.. code-block:: bash

    $ cd ~/go/src/github.com/your_name/todo
    $ go build; ./todo
    [   OK   ]   Initializing DB: [9/9]
    [   OK   ]   Server Started: http://127.0.0.1:8000

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

Well done! You have mastered the concepts of creating and modifying the progress bar in the model.

Reference
---------
.. [#f1] Rouse, Margaret (2005, April). Tag. Retrieved from https://searchmicroservices.techtarget.com/definition/tag