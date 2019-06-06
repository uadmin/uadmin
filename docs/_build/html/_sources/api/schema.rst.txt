uadmin.Schema
=============
Schema is the global schema of the system.

Structure:

.. code-block:: go

    map[string]uadmin.ModelSchema

Examples:

* `Choices`_
* `DefaultValue`_
* `DisplayName`_
* `Encrypt`_
* `ErrMsg`_
* `FormDisplay`_
* `Hidden`_
* `ListDisplay`_
* `Max`_
* `Min`_
* `Pattern`_
* `PatternMsg`_
* `ProgressBar`_
* `ReadOnly`_
* `Required`_
* `Type`_
* `UploadTo`_

**Choices**
^^^^^^^^^^^
A struct for the list of choices

Type:

.. code-block:: go

    []uadmin.Choice

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Choices

Suppose you have the given source code in friend.go where Nationality is the type of the drop down list:

.. code-block:: go

    // Nationality ...
    type Nationality int

    // Chinese ...
    func (Nationality) Chinese() Nationality {
        return 1
    }

    // Filipino ...
    func (Nationality) Filipino() Nationality {
        return 2
    }

    // Others ...
    func (Nationality) Others() Nationality {
        return 3
    }

Let’s build a choice that includes Chinese and Filipino and excludes Others. In order to do that, create a schema function of Friend model where the field name is Nationality then access Choices.

.. code-block:: go

    func main(){
        // Some codes
        
        // friend - Model Name
        // Nationality - Field Name
        uadmin.Schema["friend"].FieldByName("Nationality").Choices = []uadmin.Choice{
            // K is the ID of the choice.
            // V is the value of the choice.
            {K: 0, V: " - "},
            {K: 1, V: "Chinese"},
            {K: 2, V: "Filipino"},
        }
    }

Run your application, go to the Friend model and click Add New Friend button on the top right corner of the screen. As expected, Chinese and Filipino choices are included in the list.

.. image:: assets/friendnationalitychoices.png

**DefaultValue**
^^^^^^^^^^^^^^^^
A value assigned automatically if you want to add a new record

Type:

.. code-block:: go

    string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").DefaultValue

Let's set a feature that assigns a value automatically when creating a new record. In order to do that, create a schema function of Friend model where the field name is Nationality then access DefaultValue.

.. code-block:: go

    func main(){
        // Some codes
        
        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").DefaultValue = "Type here"
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, “Type here” value has assigned automatically in the Name field.

.. image:: assets/categorydefaultvalue.png
   :align: center

**DisplayName**
^^^^^^^^^^^^^^^
The name that you want to display in the model. It is an alias name.

Type:

.. code-block:: go

    string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").DisplayName

Let’s replace the actual field name. In order to do that, create a schema function of Category model where the field name is Name then access DisplayName.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").DisplayName = "Display Name"
    }

Run your application and go to Category model. As expected, the name has changed to “CATEGORY NAME”.

.. image:: assets/categorydisplayname.png

**Encrypt**
^^^^^^^^^^^
A feature used to encrypt the value in the database

Type:

.. code-block:: go

    bool

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Encrypt

Suppose you have two records in the Category model as shown below:

.. image:: assets/categorynametworecords.png

|

Let's encrypt the value of the Name field in the Category Model. In order to do that, create a schema function of Category model where the field name is Name then access Encrypt.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").Encrypt = true
    }

Run your application. From your project folder, open uadmin.db with DB Browser for SQLite.

.. image:: assets/uadmindbsqlite.png
   :align: center

|

Click on Execute SQL.

.. image:: assets/executesqlhighlighted.png
   :align: center

|

Get all records by typing this command: **SELECT \* FROM categories** then click the right arrow icon to execute your SQL command.

.. image:: assets/selectfromcategories.png
   :align: center

|

As expected, the Name value is encrypted in the database.

.. image:: assets/categorynameencrypt.png
   :align: center


**ErrMsg**
^^^^^^^^^^
An error message displayed beneath the input field

Type:

.. code-block:: go

    string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").ErrMsg

Let's set a feature where an error message will be displayed beneath the input Name field. In order to do that, create a schema function of Category model where the field name is Name then access ErrMsg.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").ErrMsg = "This field cannot be modified."
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, the error message was displayed beneath the input Name field.

.. image:: assets/categorynameerrmsg.png
   :align: center

**FormDisplay**
^^^^^^^^^^^^^^^
A feature that will hide the field in the editing section of the model if the value returns false

Type:

.. code-block:: go

    bool

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").FormDisplay

Let's set a feature that will hide the field in the editing section of the Category model. In order to do that, create a schema function of Category model where the field name is Name then access FormDisplay.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").FormDisplay = false
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, the Name Field is now invisible in the Category model.

.. image:: assets/categorynameformdisplay.png
   :align: center

**Hidden**
^^^^^^^^^^
A feature to hide the component in the editing section of the form

Type:

.. code-block:: go

    bool

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Hidden

Unlike in FormDisplay, the field will hide if the value is true. In order to hide the Name field in the Category model, create a schema function of Category model where the field name is Name then access Hidden.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").Hidden = true
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, the Name Field is now invisible in the Category model.

.. image:: assets/categorynameformdisplay.png

**ListDisplay**
^^^^^^^^^^^^^^^
A feature that will hide the field in the viewing section of the model if the value returns false

Type:

.. code-block:: go

    bool

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").ListDisplay

Let's set a feature that will hide the field or column name in the viewing section of the Category model. In order to hide the Name field in the Category model, create a schema function of Category model where the field name is Name then access ListDisplay.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").ListDisplay = false
    }

Run your application and go to the Category model. As expected, the Name Field in Category Model is now invisible in the list.

.. image:: assets/categorynamelistdisplay.png
   :align: center

**Max**
^^^^^^^
The maximum value the user can assign. It is applicable for numeric characters.

Type:

.. code-block:: go

    interface{}

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Max

Let's set a limitation where the user can assign a value up to 100. In order to do that, create a schema function of Todo model where the field name is Progress then access Max.

.. code-block:: go

    func main(){
        // Some codes

        // todo - Model Name
        // Progress - Field Name
        uadmin.Schema["todo"].FieldByName("Progress").Max = "100"
    }

Run your application and go to the Todo model. Let's put a numeric value beyond the maximum limit in the Progress field and see what happens.

.. image:: assets/todoprogressmax.png

**Min**
^^^^^^^
The minimum value the user can assign. It is applicable for numeric characters.

Type:

.. code-block:: go

    interface{}

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Min

Let's set a limitation where the user can assign a value at least 0. In order to do that, create a schema function of Todo model where the field name is Progress then access Min.

.. code-block:: go

    func main(){
        // Some codes

        // todo - Model Name
        // Progress - Field Name
        uadmin.Schema["todo"].FieldByName("Progress").Min = "0"
    }

Run your application and go to the Todo model. Let's put a numeric value beyond the minimum limit in the Progress field and see what happens.

.. image:: assets/todoprogressmin.png

**Pattern**
^^^^^^^^^^^
A regular expression

Type:

.. code-block:: go

    string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Pattern

Let's set a feature where the user can assign letters only in the Name field. In order to do that, create a schema function of Category model where the field name is Name then access Pattern for regular expression and PatternMsg for an error message if the user did not match the requested format.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").Pattern = "^[a-zA-Z _]*$"
        uadmin.Schema["category"].FieldByName("Name").PatternMsg = "Please match the requested format."
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. Let's assign a numeric value in the Name field. If you click Save, the system will prompt the user the the value of the Name field must assign letters only.

.. image:: assets/categorynamepattern.png
   :align: center

**PatternMsg**
^^^^^^^^^^^^^^
An error message if the user assigns a value that did not match the requested format

Type:

.. code-block:: go

    string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").PatternMsg

See `Pattern`_ for an example.

**ProgressBar**
^^^^^^^^^^^^^^^
A feature used to measure the progress of the activity

Type:

.. code-block:: go

    map[float64]string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").ProgressBar

Let's assign the value and the color of the progress bar. In order to do that, create a schema function of Todo model where the field name is Progress then access ProgressBar.

.. code-block:: go

    func main(){
        // Some codes

        // todo - Model Name
        // Progress - Field Name
        // 100.0 - maximum value
        // #07c - blue color
        uadmin.Schema["todo"].FieldByName("Progress").ProgressBar = map[float64]string{100.0: "#07c"}
    }

Run your application and go to the Todo model. As expected, the assigned values were applied to the progress bar.

.. image:: assets/todoprogressbar.png

**ReadOnly**
^^^^^^^^^^^^
A field that cannot be modified

Type:

.. code-block:: go

    string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").ReadOnly

Let's set a feature where the user cannot modify a Name field in the Category model. In order to do that, create a schema function of Category model where the field name is Name then access ReadOnly.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").ReadOnly = "true"
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, the Name field is now Read Only that means the value cannot be modified.

.. image:: assets/categorynamereadonly.png
   :align: center

**Required**
^^^^^^^^^^^^
A field that user must perform the given task(s). It cannot be skipped or left empty.

Type:

.. code-block:: go

    bool

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Required

Let's set a feature where the user needs to fill up the Name field. If the value is empty, the user will prompt the user that the value of the Name field should be assigned. In order to do that, create a schema function of Category model where the field name is Name then access Required.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").Required = true
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. If you notice, there is an asterisk (\*) symbol located on the top right after the "Name:". Let's leave the Name field value as it is. If you click Save, the system will prompt the user that the Name must be filled out.

.. image:: assets/categorynamerequired.png
   :align: center

**Type**
^^^^^^^^
The field type (e.g. file, list, progress_bar)

Type:

.. code-block:: go

    string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Type

Suppose you have this field in the Todo model as shown below:

.. image:: assets/todoprogressdefault.png

|

Let's convert the input type to the progress bar. In order to do that, create a schema function of Todo model where the field name is Progress then access Type.

.. code-block:: go

    func main(){
        // Some codes

        // todo - Model Name
        // Progress - Field Name
        uadmin.Schema["todo"].FieldByName("Progress").Type = "progress_bar"
    }

Run your application and go to the Todo model. As expected, the field type has changed from regular to a progress bar. However, the appearance does not look good because we have not assigned the value and color of the progress bar yet.

.. image:: assets/todoprogresstype.png

|

Let's improvise the appearance by assigning the value and the color of the progress bar. In order to do that, create a schema function of Todo model where the field name is Progress then access ProgressBar.

.. code-block:: go

    func main(){
        // Some codes

        // todo - Model Name
        // Progress - Field Name
        // 100.0 - maximum value
        // #07c - blue color
        uadmin.Schema["todo"].FieldByName("Progress").ProgressBar = map[float64]string{100.0: "#07c"}
    }

Run your application and go to the Todo model. As expected, the appearance of the progress bar is now good enough.

.. image:: assets/todoprogressbar.png

**UploadTo**
^^^^^^^^^^^^
A path where to save the uploaded files

Type:

.. code-block:: go

    string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").UploadTo

Let's set a feature where the uploaded file will save in the specified path on your project folder. In order to do that, create a schema function of Category model where the field name is File then access UploadTo.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // File - Field Name
        uadmin.Schema["category"].FieldByName("File").UploadTo = "/media/files/"
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. Let's add a new record that includes the uploaded file from your computer (e.g. Windows Installation.pdf).

.. image:: assets/categoryinstallationrecord.png
   :align: center

|

Result:

.. image:: assets/categoryinstallationrecordresult.png

|

From your project folder, go to /media/files/(generated_folder_name)/. As expected, the "Windows Installation.pdf" file was saved on that path.

.. image:: assets/categoryfileuploadto.png
   :align: center
