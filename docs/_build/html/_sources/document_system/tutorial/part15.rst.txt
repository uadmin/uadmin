Document System Tutorial Part 15 - Schema Form Modifier
=======================================================
In this part, we will talk about schema form modifier based on the CreatedBy form filter that checks the admin status of the user and the CreatedBy is not an empty string. If the user is not an admin and the CreatedBy is an empty string, the CreatedBy field will set as Read Only that means it cannot be modified.

Run your application then login using "johndoe" account.

.. image:: assets/johndoelogin.png
   :align: center

|

Click "DOCUMENTS".

.. image:: assets/documentsaccessdashboard.png

|

From here, go to your terminal. You will notice that the debug output is 2. This is the UserID of the "johndoe" account. The result was originally came from custom AdminPage function which was discussed in Part 13 tutorial of this application. Now click any existing record that you have in the Document model (e.g. Computer).

.. image:: assets/documentthreerecords.png

|

In fact that "johndoe" is not an admin account, we want to set some limitations to the records that "johndoe" can do such as the CreatedBy field cannot be modified by the user.

.. image:: assets/createdbyadmineditable.png

|

In order to do that, we need to use Form Modifier. Form Modifier is a function that you could pass that will allow you to modify the schema when rendering a form. It will pass to you the a pointer to the schema so you could modify it and a copy of the Model that is being rendered and the user access it to be able to customize per user (or per user group).

.. code-block:: go

    func(*uadmin.ModelSchema, interface{}, *uadmin.User)

uadmin.ModelSchema has the following fields and their definitions:

* **Name** - The name of the Model
* **DisplayName** - A human readable version of the name of the Model
* **ModelName** - The same as the Name but in small letters.
* **ModelID** - **(Data)** A place holder to store the primary key of a single row for form processing
* **Inlines** - A list of associated inlines to this model
* **InlinesData** - **(Data)** A place holder to store the data of the inlines
* **Fields** - A list of uadmin.F type representing the fields of the model
* **IncludeFormJS** - A list of string where you could add URLs to javascript files that uAdmin will run when a form view of this model is rendered
* **IncludeListJS** - A list of string where you could add URLs to javascript files that uAdmin will run when a list view of this model is rendered
* **FormModifier** - A function that you could pass that will allow you to modify the schema when rendering a form. It will pass to you the a pointer to the schema so you could modify it and a copy of the Model that is being rendered and the user access it to be able to customize per user (or per user group).
* **ListModifier** - A function that you could pass that will allow you to modify the schema when rendering a list. It will pass to you the a pointer to the schema so you could modify it and the user access it to be able to customize per user (or per user group).

**interface{}** is the parameter used to cast or access the model to modify the fields.

uadmin.User has the following fields and their definitions:

* **Username** - The username that you can use in login process and CreatedBy which is a reserved word in uAdmin
* **FirstName** - The given name of the user
* **LastName** - The surname of the user
* **Password** - A secret word or phrase that must be used to gain admission to something. This field is automatically hashed for security protection.
* **Email** - A method of exchanging messages between people using electronic devices.
* **Active** - Checks whether the user is logged in
* **Admin** - Checks whether the user is authorized to access all features in the system
* **RemoteAccess** - Checks whether the user has access to remote devices
* **UserGroup** - Returns the GroupName
* **UserGroupID** - An ID to access the UserGroup
* **Photo** - Profile picture of the user
* **LastLogin** - The date when the user last logged in his account
* **ExpiresOn** - The date when the user account expires
* **OTPRequired** - Checks whether the OTP is Active
* **OTPSeed** - Private field for OTP

.. image:: assets/userfields.png

Exit your application. Inside the main function, create a Schema Form Modifier that calls the Document model. Place it after the RegisterInlines function.

.. code-block:: go

    // Initialize docS variable that calls the document model in the schema
    docS := uadmin.Schema["document"]

    // FormModifier makes CreatedBy read only if the user is not an admin
    // and the CreatedBy is not an empty string.
    docS.FormModifier = func(s *uadmin.ModelSchema, m interface{}, u *uadmin.User) {
        // Casts an interface to the Document model
        d, _ := m.(*models.Document)
        
        // Check whether the user is not an admin and the CreatedBy Field of
        // Document model is not an empty string
        if !u.Admin && d.CreatedBy != "" {
            // Set the CreatedBy Field to read only
            s.FieldByName("CreatedBy").ReadOnly = "true"
        }
    }

    // Pass back to the schema of document model
    uadmin.Schema["document"] = docS

Now run your application using "johndoe" account.

.. image:: assets/johndoelogin.png
   :align: center

|

Click "DOCUMENTS".

.. image:: assets/documentsaccessdashboard.png

|

Click any existing record that you have in the Document model (e.g. Computer).

.. image:: assets/documentthreerecords.png

|

In fact that we are using "johndoe" non-admin account, the CreatedBy field is now set as Read Only that means it cannot be modified.

.. image:: assets/createdbyadminreadonly.png

|

In the `next part`_, we will discuss about schema list modifier based on the document list filter that checks the admin status of the user. If it is not an admin, what are the models that user can access to.

.. _next part: https://uadmin.readthedocs.io/en/latest/document_system/tutorial/part16.html
