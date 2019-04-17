Document System Tutorial Part 11 - Document and Folder Permissions
==================================================================
In this part, we will talk about setting and getting permissions for Document and Folder records.

Go to document.go inside the models folder and create a function named **GetPermissions** that holds the user parameter with the data type uadmin.User and returns Read, Add, Edit, and Delete parameters with the data type bool. If the user is an admin, all permission levels should be given.

.. code-block:: go

    // GetPermissions !
    func (d Document) GetPermissions(user uadmin.User) (Read bool, Add bool, Edit bool, Delete bool) {
        // Check whether the user is an admin
        if user.Admin {
            // Set all permissions to true
            Read = true
            Add = true
            Edit = true
            Delete = true
        }
    }

Inside the GetPermissions folder, let's check the Folder permissions.

.. code-block:: go

    // Since Folder is a foreign key to the Document model, we need to check
    // whether there is a Folder specified in the Document model.
    // We will check for folder permissions first
    // Then we will check for document permissions after that
    if d.FolderID != 0 {
        // Initialize the FolderGroup model
        folderGroup := FolderGroup{}

        // Get data by GroupID and FolderID
        uadmin.Get(&folderGroup, "group_id = ? AND folder_id = ?", user.UserGroupID, d.FolderID)

        // Check whether there is a FolderGroup record
        if folderGroup.ID != 0 {
            // Assign FolderGroup permission values to the variables
            Read = folderGroup.Read
            Add = folderGroup.Add
            Edit = folderGroup.Edit
            Delete = folderGroup.Delete
        }

        // Initialize the FolderUser model
        folderUser := FolderUser{}

        // Get data by UserID and FolderID
        uadmin.Get(&folderUser, "user_id = ? AND folder_id = ?", user.ID, d.FolderID)

        // Check whether there is a FolderUser record
        if folderUser.ID != 0 {
            // Assign FolderUser permission values to the variables
            Read = folderUser.Read
            Add = folderUser.Add
            Edit = folderUser.Edit
            Delete = folderUser.Delete
        }
    }

Now we will check for Document permissions after validating the Folder specified in the Document model.

.. code-block:: go

    // Document Permissions
	// Initialize the DocumentGroup model
	documentGroup := DocumentGroup{}

	// Get data by GroupID and DocumentID
	uadmin.Get(&documentGroup, "group_id = ? AND document_id = ?", user.UserGroupID, d.ID)

	// Check whether there is a DocumentGroup record
	if documentGroup.ID != 0 {
		// Assign DocumentGroup permission values to the variables
		Read = documentGroup.Read
		Add = documentGroup.Add
		Edit = documentGroup.Edit
		Delete = documentGroup.Delete
	}

	// Initialize the DocumentUser model
	documentUser := DocumentUser{}

	// // Get data by UserID and DocumentID
	uadmin.Get(&documentUser, "user_id = ? AND document_id = ?", user.ID, d.ID)

	// Check whether there is a DocumentUser record
	if documentUser.ID != 0 {
		// Assign DocumentUser permission values to the variables
		Read = documentUser.Read
		Add = documentUser.Add
		Edit = documentUser.Edit
		Delete = documentUser.Delete
	}

	// Return Read, Add, Edit, and Delete values
	return

In the `next part`_, we will discuss about creating a custom Count function that checks the query and the UserID.

Click `here`_ to view the full source code in this part.

.. _next part: https://uadmin.readthedocs.io/en/latest/document_system/tutorial/part12.html

.. _here: https://uadmin.readthedocs.io/en/latest/document_system/tutorial/full_code/part11.html

.. toctree::
   :maxdepth: 1

   full_code/part11
