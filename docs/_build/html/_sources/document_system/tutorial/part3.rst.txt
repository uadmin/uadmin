Document System Tutorial Part 3 - Linking Models (Part 1)
=========================================================
Linking a model to another model is as simple as creating a field using a foreign key. Foreign Key is the key used to link two models together.

**What is the purpose of the foreign key?** The purpose of the foreign key is to ensure referential integrity of the data. In other words, only values that are supposed to appear in the database are permitted.

Let’s create a new file in the models folder named “folder_user.go” with the following codes below:

.. code-block:: go

    package models

    import (
        "github.com/uadmin/uadmin"
    )

    // FolderUser !
    type FolderUser struct {
        uadmin.Model
        User     uadmin.User
        UserID   uint
        Folder   Folder
        FolderID uint
        Read     bool
        Add      bool
        Edit     bool
        Delete   bool
    }

    // FolderUser function that returns string value
    func (f *FolderUser) String() string {

        // Gives access to the fields in another model
        uadmin.Preload(f)

        // Returns the full name from the User model
        return f.User.String()
    }

In the example above, we declared the User field that calls the uadmin.User. This is a built-in system model where we can access the returning data inside which is the full name. UserID field was initialized because this is where we can fetch the returning data to be stored in the User field.

We also declared a Folder field that calls the model name "Folder" together with FolderID as a uint data type. This is our created model that was discussed in the previous chapter.

Let's add Read, Add, Edit, and Delete permissions to the user with the data type as bool (True or False). This is important especially if the folder contains confidential information. In this way, we cannot give access to all users who can read, add, edit, and delete the contents of the specific folder. We can give all access to the administrators. We can give access to some users but limited to what administrators can do. For instance, the administrators can read, add, edit and delete that specific folder. For some users, they can only read the folder but cannot add, edit and delete it because they have no access into it. That is how user permissions work.

Let’s create another file in the models folder named “folder_group.go” with the following codes below:

.. code-block:: go

    package models

    import (
        "github.com/uadmin/uadmin"
    )

    // FolderGroup !
    type FolderGroup struct {
        uadmin.Model
        Group    uadmin.UserGroup
        GroupID  uint
        Folder   Folder
        FolderID uint
        Read     bool
        Add      bool
        Edit     bool
        Delete   bool
    }

    // FolderGroup function that returns string value
    func (f *FolderGroup) String() string {
        
        // Gives access to the fields in another model
        uadmin.Preload(f)

        // Returns the GroupName from the Group model
        return f.Group.GroupName
    }

In the example above, we declared the Group field that calls the uadmin.UserGroup. This is a built-in system model where we can access the returning data inside which is the GroupName. GroupID field was initialized because this is where we can fetch the returning data to be stored in the Group field.

We also declared a Folder field that calls the model name "Folder" together with FolderID as a uint data type. This is our created model that was discussed in the previous chapter.

Like in FolderUser, we can also create permissions to the group as well. For instance, the folder contains a movie that has a rating system of 18+. User A belongs to the childhood group (age 3 to 11) and User B belongs to the adulthood one (age 18 to 55). In fact that the rating system of that movie is 18+, User A is unable to watch that movie because his age is less than 18 years old. User B has access to watch that movie because he is at least 18 years old.

Now go to main.go and register the models that we have created.

.. code-block:: go

    func main() {
        // Register models to uAdmin
        uadmin.Register(
            models.Folder{},
            models.FolderGroup{}, // place it here
            models.FolderUser{}, // place it here
        )

        // Some codes
    }

Run your application. As expected, FolderGroup and FolderUser models are added in the uAdmin Dashboard.

.. image:: assets/folderusergroup.png

|

In the `next part`_, we will discuss about folder concepts and how to create records in an application.

.. _next part: https://uadmin.readthedocs.io/en/latest/document_system/tutorial/part4.html
