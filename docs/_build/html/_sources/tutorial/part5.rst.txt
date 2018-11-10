uAdmin Tutorial Part 5 - Applying uAdmin Tags
=============================================
Create a file named item.go inside your models folder, containing the following codes below.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Item model ...
    type Item struct {
        uadmin.Model
        Name        string `uadmin:"required"`
        Description string
        Cost        int
        Rating      int
    }

Now register the model on main.go where models is folder name and Item is model/struct name. Apply the inlines as well.

.. code-block:: go

    func main() {
        uadmin.Register(
            models.Todo{},
            models.Category{},
            models.Friend{},
            models.Item{},  //  <-- place it here
        )

        // Some codes contained in this part...

        // ----------- ADD THIS CODE -----------
        uadmin.RegisterInlines(models.Item{}, map[string]string{
            "Todo": "ItemID",
        })
        // ----------- ADD THIS CODE -----------
        uadmin.StartServer()
    }

Set the foreign key of an Item model to the Todo model and apply the tag "help" to inform the user waht are the requirements needed in order to accomplish his activity.

.. code-block:: go

    // Todo model ...
    type Todo struct {
        uadmin.Model
        Name        string
        Description string   `uadmin:"html"`
        Category    Category
        CategoryID  uint
        Friend      Friend `uadmin:"help:Who will be a part of your activity?"`
        FriendID    uint
        Item    Item    `uadmin:"help:What are the requirements needed in order to accomplish your activity?"`
        ItemID  uint
        TargetDate  time.Time
        Progress    int `uadmin:"progress_bar"`
    }

Now let's try something much cooler that we can apply in the Item model by adding different types of tags. Before we proceed, add more data in your Item model. Once you are done, let's add the "search" tag in the name field of item.go and see what happens.

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Item model ...
    type Item struct {
        uadmin.Model
        Name        string `uadmin:"required;search"` // <-- place it here
        Description string
        Cost        int
        Rating      int
    }

Result

.. image:: assets/searchtagapplied.png

|

Search the word "mini" and see what happens.

.. image:: assets/searchtagappliedoutput.png

|

Nice! Now go back to item.go and apply the tag categorical_filter and filter in the Name field and see what happens.

.. code-block:: go

	Name string `uadmin:"required;search;categorical_filter;filter"`

Click the filter button on the upper right.

Result

.. image:: assets/filtertagapplied.png

|

Now let's filter the word "iPad" and see what happens.

.. image:: assets/filtertagappliedoutput.png

|

We can also apply display_name tag with a given value such as "Product Name".

.. code-block:: go

    Name string `uadmin:"required;search;categorical_filter;filter;display_name:Product Name"`

|

Result

.. image:: assets/displaynametagapplied.png

|

uAdmin has a default_value tag which will generate a value automatically in the field. Let's say "Computer".

.. code-block:: go

    Name string `uadmin:"required;search;categorical_filter;filter;display_name:Product Name;default_value:Computer"`

|

Result

.. image:: assets/defaultvaluetagapplied.png

|

You can also add multilingual tag in the Description field. This means you can use more than two languages for input.

.. code-block:: go

    Description string `uadmin:"multilingual"`

|

Result

.. image:: assets/multilingualtagapplied.png

|

If you want to add more languages in your model, go to the Languages in the uAdmin dashboard.

.. image:: assets/languageshighlighted.png

|

Let's say I want to add Chinese and Tagalog in the Items model. In order to do that, set the Active as enabled.

.. image:: assets/activehighlighted.png

|

Now go back to the Items model and see what happens.

.. image:: assets/multilingualtagappliedmultiple.png

To customize your own languages, click `here`_ for the instructions.

.. _here: https://medium.com/@twistedhardware/uadmin-the-golang-web-framework-4-customizing-dashboard-d96d90792a07

|

In the Cost field, set the "money" tag and see what happens.

.. code-block:: go

    Cost int `uadmin:"money"`

|

Result

.. image:: assets/moneytagapplied.png

|

You can also set pattern and pattern_msg tag in the Cost field. This means the user must input numbers only. If he inputs otherwise, the pattern message will show up on the screen.

.. code-block:: go

    Cost int `uadmin:"money;pattern:^[0-9]*$;pattern_msg:Your input must be a number."`

|

Result

.. image:: assets/patterntagapplied.png

|

To solve this case, we can use a help tag feature in order to give users a solution to the complex tasks encountered in the model.

.. code-block:: go

    Cost int `uadmin:"money;pattern:^[0-9]*$;pattern_msg:Your input must be a number.;help:Input numeric characters only in this field."`

|

Result

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

To the item.go inside the models folder

.. code-block:: go

    package models

    import "github.com/uadmin/uadmin"

    // Item model ...
    type Item struct {
        uadmin.Model
        Name         string     `uadmin:"search;categorical_filter;filter;display_name:Product Name"`
        Description  string     `uadmin:"multilingual"`
        Category     []Category `uadmin:"m2m;list_exclude"`  // <-- place it here
        CategoryList string     `uadmin:"read_only"`         // <-- place it here
        Cost         int        `uadmin:"money;pattern:^[0-9]*$;pattern_msg:Your input must be a number."`
        Rating       int        `uadmin:"min:1;max:5"`
    }

Copy this one as well and paste it below the Item struct.

.. code-block:: go

    // CategorySave ...
    func (i *Item) CategorySave() {
        // Initializes the catList as empty string
        catList := ""

        // This process will get the name of the category, store into the catList and if the index value is not equal to the number of category, it will insert the comma symbol at the end of the word.
        for x, key := range i.Category {
            catList += key.Name
            if x != len(i.Category)-1 {
                catList += ", "
            }
        }

        // Store the catList variable to the CategoryList field in the Item model
        i.CategoryList = catList

        // Override save
        uadmin.Save(i)
    }

    // Save ...
    func (i *Item) Save() {
        if i.ID == 0 {
            i.CategorySave()
        }

        i.CategorySave()
    }

|

Let's run the application and see what happens.

.. image:: assets/m2mtagapplied.png

|

Result

.. image:: assets/m2mtagappliedoutput.png

Well done! You already know how to apply most of the tags available in our uAdmin framework that are functional in our Todo List project.

In the `next part`_, we will discuss on how to apply validation in the back-end.

.. _next part: https://uadmin.readthedocs.io/en/latest/tutorial/part6.html