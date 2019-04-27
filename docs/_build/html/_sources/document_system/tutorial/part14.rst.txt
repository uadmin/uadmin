Document System Tutorial Part 14 - Permissions Form
===================================================
In this part, we will discuss about creating the Permissions field and displaying the permission status for each document records.

First of all, run your application. We need to get the ID of each users. In order to do that, go to the Document System Dashboard then click "USERS".

.. image:: assets/usershighlighted.png

|

Click the "System Admin" user.

.. image:: assets/systemadminrecord.png

|

In the address bar, there is a number 1 in the last part of the hyperlink. In short, the User ID of the System Admin is 1.

.. image:: assets/userid1.png
   :align: center

|

Go back then click the "John Doe" user.

.. image:: assets/johndoerecord.png

|

In the address bar, there is a number 2 in the last part of the hyperlink. In short, the User ID of the John Doe is 2.

.. image:: assets/userid2.png
   :align: center

|

Here is the summary:

    * System Admin - 1
    * John Doe - 2

Now exit your application. Go to the document.go and create a new function named Permissions__Form() that returns a string. This returns the Read, Add, Edit, and Delete permissions based on an assigned user ID.

.. code-block:: go

    // Permissions__Form creates a new field named Permissions !
    func (d Document) Permissions__Form() string {
        // Initialize u variable that calls the User model
        u := uadmin.User{}

        // Get the user record based on an assigned ID
        uadmin.Get(&u, "id = ?", 1)

        // Initialize read, add, edit and delete that gets the permission for a
        // specific user based on an assigned ID
        r, a, e, del := d.GetPermissions(u)

        // Returns the permission status
        return fmt.Sprintf("Read: %v Add: %v, Edit: %v, Delete: %v", r, a, e, del)
    }

In fact that our assigned user ID is 1, we are getting the System Admin user record then we are passing the Read, Add, Edit, and Delete permission values to that user.

Now run your application. Go to the Document System Dashboard then click "DOCUMENT USERS".

.. image:: assets/documentusershighlighted.png

|

Click the first record that has a document of "Computer".

.. image:: assets/documentuserfirstrecord.png

|

Because all existing records are using John Doe as the User, let's change the User of this record to System Admin. Remove Read and Add permissions to this one as well so that no one can access to this record except the System Admin itself.

.. image:: assets/documentuserfirstrecordmodify.png
   :align: center

|

Result

.. image:: assets/documentuserfirstrecordresult.png

|

Now go back to the Document System Dashboard then click "DOCUMENTS".

.. image:: assets/documentshighlighted.png

|

Click the "Computer" document.

.. image:: assets/documentcomputer.png

|

Scroll down the form then you will see the new field named "Permissions" where all permission levels are set to false.

.. image:: assets/permissionfieldfirstrecord.png

|

In the `next part`_, we will talk about schema form modifier based on the CreatedBy form filter that checks the admin status of the user and the CreatedBy is not an empty string. If the user is not an admin and the CreatedBy is an empty string, the CreatedBy field will set as Read Only that means it cannot be modified.

.. _next part: https://uadmin.readthedocs.io/en/latest/document_system/tutorial/part15.html
