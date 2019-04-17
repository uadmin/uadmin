Document System Tutorial Part 11 - Document and Folder Permissions (Full Source Code)
=====================================================================================

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

        // Since Folder is a foreign key to the Document model, we need to check
        // whether there is a Folder specified in the Document model.
        // We will check for folder permissions first
        // Then we will check for document permissions after that
        if d.FolderID != 0 {
            // Initialize the FolderGroup model
            folderGroup := FolderGroup{}

            // Get data by GroupID and FolderID
            uadmin.Get(&folderGroup, "group_id = ? AND folder_id = ?", user.UserGroupID, d.FolderID)

            // Check whether there is a FolderGroup recird
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
    }
