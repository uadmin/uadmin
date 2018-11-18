uAdmin Tutorial Part 6 - Back-end Validation
============================================
For more advanced validation, sometimes you need to implement some validation from the back-end. This is the case for validation that required access to the database to check for duplicate entries or check some permissions like “You are not allowed to assign this task to people outside your department”. Regardless of the case this is how to implement back-end validation.

Let’s say we don’t want people to add duplicate entries for todo. The way we will do that is check the database and see if there is another todo record with the same name. If we find another record, we can return a message that tells the user that the todo item has been added to the system already.

Open /models/todo.go and add a new method called Validate to your Todo struct.

.. code-block:: go

    // Todo model ...
    type Todo struct {
        uadmin.Model
        Name        string
        Description string `uadmin:"html"`
        Category    Category
        CategoryID  uint
        Friend      Friend `uadmin:"help:Who will be a part of your activity?"`
        FriendID    uint
        Item        Item `uadmin:"help:What are the requirements needed in order to accomplish your activity?"`
        ItemID      uint
        TargetDate  time.Time
        Progress    int `uadmin:"progress_bar"`
    }

    // Save model ...
    func (t *Todo) Save() {
        // Save the model to DB
        uadmin.Save(t)
        // Some other business Logic ...
    }

    // Validate function ...
    func (t Todo) Validate() (errMsg map[string]string) {
        // Initialize the error messages
        errMsg = map[string]string{}
        // Get any records from the database that maches the name of
        // this record and make sure the record is not the record we are
        // editing right now
        todo := Todo{}
        if uadmin.Count(&todo, "name = ? AND id <> ?", t.Name, t.ID) != 0 {
            errMsg["Name"] = "This todo name is already in the system"
        }
        return
    }

Notice that the receiver for Validate() is not a pointer but the struct type. Also notice that the return is a map where the key is the field name and the value is the error message.

If you try now to add a new record with an existing todo record’s name, it will show me this error:

.. image:: assets/todobackendvalidate.png
   :align: center

|

You may also do the same process of applying validate function in the other models that you have, this time with using different variables related to the model and different error messages as part of your challenge. Once you master them, congrats! You are now ready to proceed with `configuring APIs`_.

.. _configuring APIs: https://uadmin.readthedocs.io/en/latest/tutorial/part7.html