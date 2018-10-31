uAdmin Tutorial Part 2 - Models
===============================
Here are the following subtopics to be discussed in this part:

    * `External models`_
    * `Tags`_
    * `Internal models`_
    * `Linking models`_
    * `Creating more models`_
    * `Applying more uAdmin tags`_
    * `Register inlines`_

External models
^^^^^^^^^^^^^^^^
External models are models outside of main.go and have their own .go file. Let’s add a category external model, create a file named category.go and add the following code:

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Category model ...
    type Category struct {
	    uadmin.Model
	    Name string
	    Icon string
    }

Category Model User Interface

.. image:: assets/categorymodeldesign.png

|

Now register the model on main.go where models is folder name and category is model/struct name:

Copy this code below

.. code-block:: go

    "github.com/username/todo/models" // put this code inside import
    models.Category{}, // put this code inside the func main()

To the main.go

.. code-block:: go

    package main

    import (
        "time"
        "github.com/username/todo/models" // <-- place it here
        "github.com/uadmin/uadmin"
    )

    // TODO model ...
    type TODO struct {
	    uadmin.Model
	    Name        string
	    Description string   `uadmin:"html"`
	    TargetDate  time.Time
	    Progress    int `uadmin:"progress_bar"`
    }

    func main() {
	    uadmin.Register(
		    TODO{},
		    models.Category{}, // <-- place it here
	    )
	    uadmin.StartServer()
    }

|

Let's run the code and see what happens:

.. code-block:: bash

    $ cd ~/go/src/github.com/your_name/todo
    $ go build; ./firstapp
    [   OK   ]   Initializing DB: [10/10]
    [   OK   ]   Server Started: http://127.0.0.1:8080

|

As expected, the category model is added in the uAdmin Dashboard.

.. image:: assets/categorymodelselected.png

|

Let's create a new data in the category model.

.. image:: assets/categorydata.png

|

Output

.. image:: assets/categorydataoutput.png

|

Tags
^^^^
uAdmin has a tag feature that allows a field to change to an appropriate type. Let’s tag the Name as “required” and Icon as “image” in our category model.

Tags are added beside the field names after the data type, like this:

.. code-block:: go

    Name string `uadmin:"required"`
    Icon string `uadmin:"image"`

To the category.go inside the models folder

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Category model ...
    type Category struct {
	    uadmin.Model
	    Name string `uadmin:"required"` // <-- place it here
	    Icon string `uadmin:"image"` // <-- place it here
    }

|

Let's run the code and see what happens.

.. image:: assets/categorywithtagapplied.png

As you can see, the Name field is now required indicated by the * symbol and the Icon field is now an image type.

|

Output

.. image:: assets/categorydataoutputwithtag.png

|

.. code-block:: go

    Icon string `uadmin:"image"`

uAdmin also allows you to crop your images. In order to that, click the image icon highlighted below.

.. image:: assets/iconhighlighted.png

|

Click the crop icon on the top left corner.

.. image:: assets/cropiconhighlighted.png

|

You are now set to edit mode. Click any points highlighted below then drag your mouse in order to crop/resize your image.

.. image:: assets/croppointshighlighted.png

.. image:: assets/croppedicon.png

|

Once you are done, click the Crop button below and refresh the webpage to save your progress.

.. image:: assets/croppediconoutput.png

Well done! The travel icon is now cropped in the model structure.

|

.. code-block:: go

    Name string `uadmin:"required"`

What if I set the name value as empty?

.. image:: assets/namefieldempty.png

A warning message "Please fill out this field." will display on your screen because the Name field has a "required" tag on it.

That is how the uAdmin tag works in this scenario. For more information about tags, click `here`_.

.. _here: file:///home/dev1/go/src/github.com/uadmin/uadmin/docs/_build/html/tags.html

Internal Models
^^^^^^^^^^^^^^^
Internal models are models inside your main.go and don’t have their .go file, they are useful if you want to make something quick but it is advisable to always you external models.

The code below is an example of internal model:

.. code-block:: go

    package main

    import (
	    "time"
	    "github.com/rn1hd/todo/models"
	    "github.com/uadmin/uadmin"
    )

    // TODO internal model ... 
    type TODO struct {
	    uadmin.Model
	    Name        string
	    Description string `uadmin:"html"`
	    TargetDate  time.Time
	    Progress    int `uadmin:"progress_bar"`
    }

    func main() {
	    uadmin.Register(
		    TODO{}, // register the TODO struct
		    models.Category{},
	    )
	    uadmin.StartServer()
    }

Linking models
^^^^^^^^^^^^^^
Linking a model to another model is as simple as creating a field. In the example below we linked the Category model into TODO model, now the TODO model will return its data as a field in the Category model.

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
	    Description string   `uadmin:"html"`
	    Category    Category // <-- Category Model
	    CategoryID  uint     // <-- CategoryID
	    TargetDate  time.Time
	    Progress    int `uadmin:"progress_bar"`
    }

|

Result:

.. image:: assets/categoryaddedintodo.png

|

Now let's add CreatedAt field in the TODO model, set the tag as "hidden". The "hidden" tag means the field is invisible in the editing section.

Copy this code below

.. code-block:: go

    CreatedAt   time.Time `uadmin:"hidden"`

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
	    Category    Category
	    CategoryID  uint
	    CreatedAt   time.Time `uadmin:"hidden"` // <-- place it here
	    TargetDate  time.Time
	    Progress    int `uadmin:"progress_bar"`
    }

|

Now let's create a new data in the Todo model. As you can see, the CreatedAt field cannot be seen in the editing section.

.. image:: assets/buildarobotdataintodo.png

|

But when you save it...

.. image:: assets/buildarobotdataintodooutput.png

Tada! The CreatedAt field is shown in the output of the Todo model.


Creating more models
^^^^^^^^^^^^^^^^^^^^
Create a file named friends.go inside your models folder, containing the following codes below.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Friends model ...
    type Friends struct {
	    uadmin.Model
	    Name     string `uadmin:"required"`
	    Email    string `uadmin:"email"`
	    Password string `uadmin:"password;list_exclude"`
    }

Friends Model User Interface

.. image:: assets/friendsmodeldesign.png

|

Now connect the friends model into the main.go by calling the models.Friends{} inside the uadmin.Register.

Copy this code below

.. code-block:: go

    models.Friends{}, // put this code inside the func main()

To the main.go

.. code-block:: go

    package main

    import (
	    "github.com/rn1hd/todo/models"
	    "github.com/uadmin/uadmin"
    )

    func main() {
	    uadmin.Register(
		    models.TODO{},
		    models.Category{},
		    models.Friends{}, // <-- place it here
	    )
	    uadmin.StartServer()
    }

|

Let's run the code and see what happens:

.. code-block:: bash

    $ cd ~/go/src/github.com/your_name/todo
    $ go build; ./firstapp
    [   OK   ]   Initializing DB: [11/11]
    [   OK   ]   Server Started: http://127.0.0.1:8080

|

As expected, the friends model is added in the uAdmin Dashboard.

.. image:: assets/friendsmodelselected.png

|

Let's create a new data in the friends model.

.. image:: assets/friendsdata.png

|

Output

.. image:: assets/friendsdataoutput.png

|

As you can see, the password field is not shown in the output. Why? If you go back to the friends model, the password field has the tag name "list_exclude". It means it will hide the field or column name in the model structure.

Let's create a relationship between the friends and todo models. In order to do that, call the struct name you wish to include on the first line and the ID with the data type on the second line in todo.go.

Copy this code below

.. code-block:: go

    Friends     Friends
    FriendsID   uint

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
	    Category    Category
	    CategoryID  uint
	    Friends     Friends   // <-- place it here
	    FriendsID   uint      // <-- place it here
	    CreatedAt   time.Time `uadmin:"hidden"`
	    TargetDate  time.Time
	    Progress    int `uadmin:"progress_bar"`
    }

|

Let's run the code again. Go back to your todo model and see what happens.

.. image:: assets/friendsaddedintodo.png

|

Output

.. image:: assets/friendsaddedintodooutput.png

The friends model is now connected into the todo model.

Create a file named items.go inside your models folder, containing the following codes below.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Items model ...
    type Items struct {
	    uadmin.Model
	    Name        string `uadmin:"required"`
	    Description string
	    Cost        int
	    Rating      int
    }

Item Model User Interface

.. image:: assets/itemsmodeldesign.png

|

Now connect the items model into the main.go by calling the models.Items{} inside the uadmin.Register.

Copy this code below

.. code-block:: go

    models.Items{}, // put this code inside the func main()

To the main.go

.. code-block:: go

    package main

    import (
	    "github.com/rn1hd/todo/models"
	    "github.com/uadmin/uadmin"
    )

    func main() {
	    uadmin.Register(
		    models.TODO{},
		    models.Category{},
		    models.Friends{},
		    models.Items{}, // <-- place it here
	    )
	    uadmin.StartServer()
    }

|

Let's run the code and see what happens:

.. code-block:: bash

    $ cd ~/go/src/github.com/your_name/todo
    $ go build; ./firstapp
    [   OK   ]   Initializing DB: [12/12]
    [   OK   ]   Server Started: http://127.0.0.1:8080

|

As expected, the items model is added in the uAdmin Dashboard.

.. image:: assets/itemsmodelselected.png

|

Let's create a new data in the items model.

.. image:: assets/itemsdata.png

|

Output

.. image:: assets/itemsdataoutput.png

|

Let's create a relationship between the items and todo models. In order to do that, call the struct name you wish to include on the first line and the ID with the data type on the second line in todo.go.

Copy this code below

.. code-block:: go

    Items       Items
    ItemsID     uint

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
	    Category    Category
	    CategoryID  uint
	    Friends     Friends
	    FriendsID   uint
	    Items       Items     // <-- place it here
	    ItemsID     uint      // <-- place it here
	    CreatedAt   time.Time `uadmin:"hidden"`
	    TargetDate  time.Time
	    Progress    int `uadmin:"progress_bar"`
    }

|

Let's run the code again. Go back to your todo model and see what happens.

.. image:: assets/itemsaddedintodo.png

|

Output

.. image:: assets/itemsaddedintodooutput.png

The items model is now connected into the todo model.

Applying more uAdmin tags
^^^^^^^^^^^^^^^^^^^^^^^^^
Now let's try something much cooler that we can apply in the items model by adding different types of tags. Before we proceed, add more data in your items model. Once you are done, let's add the "search" tag in the name field of items.go and see what happens.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Items model ...
    type Items struct {
	    uadmin.Model
	    Name        string `uadmin:"required;search"` // <-- place it here
	    Description string
	    Cost        int
	    Rating      int
    }

Output

.. image:: assets/searchtagapplied.png

|

Search the word "mini" and see what happens.

.. image:: assets/searchtagappliedoutput.png

|

Nice! Now go back to items.go and apply the tag categorical_filter and filter in the name field and see what happens.

.. code-block:: go

	Name string `uadmin:"required;search;categorical_filter;filter"` // <-- place it here

Click the filter button on the upper right.

Output

.. image:: assets/filtertagapplied.png

|

Now let's filter the word "iPad" and see what happens.

.. image:: assets/filtertagappliedoutput.png

|

We can also apply display_name tag with a given value such as "Product Name".

.. code-block:: go

    Name string `uadmin:"required;search;categorical_filter;filter;display_name:Product Name"` // <-- place it here

|

Output

.. image:: assets/displaynametagapplied.png

|

uAdmin has a default_value tag which will generate a value automatically in the field. Let's say "Computer".

.. code-block:: go

    Name string `uadmin:"required;search;categorical_filter;filter;display_name:Product Name;default_value:Computer"`

|

Output

.. image:: assets/defaultvaluetagapplied.png

|

You can also add multilingual tag in the Description field. This means you can use more than two languages for input.

.. code-block:: go

    Description string `uadmin:"multilingual"` // <-- place it here

|

Output

.. image:: assets/multilingualtagapplied.png

|

In the Cost field, set the "money" tag and see what happens.

.. code-block:: go

    Cost int `uadmin:"money"` // <-- place it here

|

Output

.. image:: assets/moneytagapplied.png

|

You can also set pattern and pattern_msg tag in the Cost field. This means the user must input numbers only. If he inputs otherwise, the pattern message will show up on the screen.

.. code-block:: go

    Cost int `uadmin:"money;pattern:^[0-9]*$;pattern_msg:Your input must be a number."` // <-- place it here

|

Output

.. image:: assets/patterntagapplied.png

|

To solve this case, we can use a help tag feature in order to give users a solution to the complex tasks encountered in the model.

.. code-block:: go

    Cost int `uadmin:"money;pattern:^[0-9]*$;pattern_msg:Your input must be a number.;help:Input numeric characters only in this field."` // <-- place it here

|

Output:

.. image:: assets/helptagapplied.png

|

We can also use min and max tags in the Rating field. Min tag means the minimum value that a user can input and the max one means the maximum value. Let's set the min value as 1 and the max value as 5.

.. code-block:: go

    Rating int `uadmin:"min:1;max:5"`

|

See what happens if the user inputs the value outside the range.

.. image:: assets/minmaxtagapplied.png

|

uAdmin also has a multiselection feature that allows you to select more than one element inside an input box field. In order to do that, let's add Category on the first line, use the array type, set as "m2m" and "list_exclude", and add CategoryList on the second line with the tag "read_only". This means it cannot be modified.

Copy this code below

.. code-block:: go

    Category     []Category `uadmin:"m2m;list_exclude"`
    CategoryList string     `uadmin:"read_only"`

To the items.go inside the models folder

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Items model ...
    type Items struct {
	    uadmin.Model
	    Name         string     `uadmin:"search;categorical_filter;filter;display_name:Product Name"`
	    Description  string     `uadmin:"multilingual"`
	    Category     []Category `uadmin:"m2m;list_exclude"`  // <-- place it here
	    CategoryList string     `uadmin:"read_only"`         // <-- place it here
	    Cost         int        `uadmin:"money;pattern:^[0-9]*$;pattern_msg:Your input must be a number."`
	    Rating       int        `uadmin:"min:1;max:5"`
    }

Copy this one as well and paste it below the items struct.

.. code-block:: go

    // CategorySave ...
    func (i *Items) CategorySave() {
	    catList := ""

	    for x, key := range i.Category {
		    catList += key.Name
		    if x != len(i.Category)-1 {
			    catList += ", "
		    }
	    }

	    i.CategoryList = catList
	    uadmin.Save(i)
    }

    // Save ...
    func (i *Items) Save() {
	    if i.ID == 0 {
		    i.CategorySave()
	    }
	
	    i.CategorySave()
    }

|

Let's run the application and see what happens.

.. image:: assets/m2mtagapplied.png

|

Output

.. image:: assets/m2mtagappliedoutput.png

Well done! You already know how to apply most of the tags available in our uAdmin framework that are functional in our Todo List project.

Register inlines
^^^^^^^^^^^^^^^^
Register inline allows you to merge a submodel to a parent model where the foreign key of the submodels are specified.

**Why do we use Register inlines?** We use them to show that the field of a model is related to another model as long as there is a foreign key specified.

Syntax:

.. code-block:: go

    uadmin.RegisterInlines(/folder_name/./struct_name of a submodel/{}, map[string]string{
		"/parent_model name/": "/sub_model name/ID",
	})

Now let's apply it in the main.go. Copy the codes below and paste it after the uadmin.Register function.

.. code-block:: go

    uadmin.RegisterInlines(models.Category{}, map[string]string{
        "TODO": "CategoryID",
    })
    uadmin.RegisterInlines(models.Friends{}, map[string]string{
        "TODO": "FriendsID",
    })
    uadmin.RegisterInlines(models.Items{}, map[string]string{
        "TODO": "ItemsID",
    })

Let's run the application and see what happens.

.. image:: assets/registerinlinetodo.png

Tada! The parent model TODO is now included in the Category submodel as shown above. You can go to Friends and Items models and it will display the same result.