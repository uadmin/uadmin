API Documentation
=================
Here are all public functions in the uAdmin, their syntax, and how to use them in the project.

* `uadmin.Action`_
* `uadmin.AdminPage`_
* `uadmin.All`_
* `uadmin.BindIP`_
* `uadmin.Choice`_
* `uadmin.ClearDB`_
* `uadmin.CookieTimeout`_
* `uadmin.Count`_
* `uadmin.CustomTranslation`_
* `uadmin.DashboardMenu`_
* `uadmin.Database`_
* `uadmin.DBSettings`_
* `uadmin.DEBUG`_
* `uadmin.DebugDB`_
* `uadmin.Delete`_
* `uadmin.DeleteList`_
* `uadmin.EmailFrom`_
* `uadmin.EmailPassword`_
* `uadmin.EmailSMTPServer`_
* `uadmin.EmailSMTPServerPort`_
* `uadmin.EmailUsername`_
* `uadmin.ERROR`_
* `uadmin.F`_
* `uadmin.Filter`_
* `uadmin.FilterBuilder`_
* `uadmin.GenerateBase32`_
* `uadmin.GenerateBase64`_
* `uadmin.Get`_
* `uadmin.GetDB`_
* `uadmin.GetID`_
* `uadmin.GetString`_
* `uadmin.GetUserFromRequest`_
* `uadmin.GroupPermission`_
* `uadmin.HideInDashboarder`_
* `uadmin.INFO`_
* `uadmin.IsAuthenticated`_
* `uadmin.JSONMarshal`_
* `uadmin.Language`_
* `uadmin.Log`_
* `uadmin.Login`_
* `uadmin.Login2FA`_
* `uadmin.Logout`_
* `uadmin.MaxImageHeight`_
* `uadmin.MaxImageWidth`_
* `uadmin.MaxUploadFileSize`_
* `uadmin.Model`_
* `uadmin.ModelSchema`_
* `uadmin.MongoDB`_
* `uadmin.MongoModel`_
* `uadmin.MongoSettings`_
* `uadmin.NewModel`_
* `uadmin.NewModelArray`_
* `uadmin.OK`_
* `uadmin.OTPAlgorithm`_
* `uadmin.OTPDigits`_
* `uadmin.OTPPeriod`_
* `uadmin.OTPSkew`_
* `uadmin.PageLength`_
* `uadmin.Port`_
* `uadmin.Preload`_
* `uadmin.PublicMedia`_
* `uadmin.Register`_
* `uadmin.RegisterInlines`_
* `uadmin.ReportingLevel`_
* `uadmin.ReportTimeStamp`_
* `uadmin.ReturnJSON`_
* `uadmin.RootURL`_
* `uadmin.Salt`_
* `uadmin.Save`_
* `uadmin.Schema`_
* `uadmin.SendEmail`_
* `uadmin.Session`_
* `uadmin.SiteName`_
* `uadmin.StartSecureServer`_
* `uadmin.StartServer`_
* `uadmin.Tf`_
* `uadmin.Theme`_
* `uadmin.Trail`_
* `uadmin.Translate`_
* `uadmin.Update`_
* `uadmin.UploadImageHandler`_
* `uadmin.User`_
* `uadmin.UserGroup`_
* `uadmin.UserPermission`_
* `uadmin.Version`_
* `uadmin.WARNING`_
* `uadmin.WORKING`_

Functions
---------

**uadmin.Action**
^^^^^^^^^^^^^^^^^
Action is the process of doing something where you can check the status of your activities in the uAdmin project.

Syntax:

.. code-block:: go

    type Action int
    
**uadmin.AdminPage**
^^^^^^^^^^^^^^^^^^^^
AdminPage fetches records from the database with some standard rules such as sorting data, multiples of, and setting a limit that can be used in pagination.

Syntax:

.. code-block:: go

    AdminPage func(order string, asc bool, offset int, limit int, a interface{}, query interface{}, args ...interface{}) (err error)

**uadmin.All**
^^^^^^^^^^^^^^
All fetches all object in the database.

Syntax:

.. code-block:: go

    All func(a interface{}) (err error)

**uadmin.BindIP**
^^^^^^^^^^^^^^^^^
BindIP is the IP the application listens to.

Syntax:

.. code-block:: go

    BindIP string

**uadmin.Choice**
^^^^^^^^^^^^^^^^^
Choice is a struct for list choices.

Syntax:

.. code-block:: go

    type Choice struct {
	    V        string
	    K        uint
	    Selected bool
    }

**uadmin.ClearDB**
^^^^^^^^^^^^^^^^^^
ClearDB clears the database object.

Syntax:

.. code-block:: go

    ClearDB func()

**uadmin.CookieTimeout**
^^^^^^^^^^^^^^^^^^^^^^^^
CookieTimeout is the timeout of a login cookie in seconds.

Syntax:

.. code-block:: go

    CookieTimeout int

**uadmin.Count**
^^^^^^^^^^^^^^^^
Count return the count of records in a table based on a filter.

Syntax:

.. code-block:: go

    Count func(a interface{}, query interface{}, args ...interface{}) int

**uadmin.CustomTranslation**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^
CustomTranslation allows a user to customize any languages in the uAdmin system.

Syntax:

.. code-block:: go

    CustomTranslation []string

**uadmin.DashboardMenu**
^^^^^^^^^^^^^^^^^^^^^^^^
DashboardMenu is a system in uAdmin used to check and modify the settings of a model.

Syntax:

.. code-block:: go

    type DashboardMenu struct {
	    Model
	    MenuName string `uadmin:"required;list_exclude;multilingual;filter"`
	    URL      string `uadmin:"required"`
	    ToolTip  string
	    Icon     string `uadmin:"image"`
	    Cat      string `uadmin:"filter"`
	    Hidden   bool   `uadmin:"filter"`
    }

**uadmin.Database**
^^^^^^^^^^^^^^^^^^^
Database is the active Database settings.

Syntax:

.. code-block:: go

    Database *DBSettings

**uadmin.DBSettings**
^^^^^^^^^^^^^^^^^^^^^
DBSettings is a feature that allows a user to configure the settings of a database.

Syntax:

.. code-block:: go

    type DBSettings struct {
	    Type     string // SQLLite, MySQL
	    Name     string // File/DB name
	    User     string
	    Password string
	    Host     string
	    Port     int
    }

**uadmin.DEBUG**
^^^^^^^^^^^^^^^^
DEBUG is the process of identifying and removing errors.

Syntax:

.. code-block:: go

    const DEBUG int = 0

**uadmin.DebugDB**
^^^^^^^^^^^^^^^^^^
DebugDB prints all SQL statements going to DB.

Syntax:

.. code-block:: go

    DebugDB bool

**uadmin.Delete**
^^^^^^^^^^^^^^^^^
Delete records from database.

Syntax:

.. code-block:: go

    Delete func(a interface{}) (err error)

**uadmin.DeleteList**
^^^^^^^^^^^^^^^^^^^^^
Delete the list of records from database.

Syntax:

.. code-block:: go

    DeleteList func(a interface{}, query interface{}, args ...interface{}) (err error)

**uadmin.EmailFrom**
^^^^^^^^^^^^^^^^^^^^
EmailFrom identifies where the email is coming from.

Syntax:

.. code-block:: go

    EmailFrom string

**uadmin.EmailPassword**
^^^^^^^^^^^^^^^^^^^^^^^^
EmailPassword sets the password of an email.

Syntax:

.. code-block:: go

    EmailPassword string

**uadmin.EmailSMTPServer**
^^^^^^^^^^^^^^^^^^^^^^^^^^
EmailSMTPServer sets the name of the SMTP Server in an email.
Syntax:

.. code-block:: go

    EmailSMTPServer string

**uadmin.EmailSMTPServerPort**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
EmailSMTPServerPort sets the port number of an SMTP Server in an email.

Syntax:

.. code-block:: go

    EmailSMTPServerPort int

**uadmin.EmailUsername**
^^^^^^^^^^^^^^^^^^^^^^^^
EmailUsername sets the username of an email.

Syntax:

.. code-block:: go

    EmailUsername string

**uadmin.ERROR**
^^^^^^^^^^^^^^^^
ERROR is a status to notify the user that there is a problem in an application.

Syntax:

.. code-block:: go

    const ERROR int = 5

**uadmin.F**
^^^^^^^^^^^^
F is a field.

Syntax:

.. code-block:: go

    type F struct {
        Name              string
        DisplayName       string
        Type              string
        Value             interface{}
        Help              string
        Max               interface{}
        Min               interface{}
        Format            string
        DefaultValue      string
        Required          bool
        Pattern           string
        PatternMsg        string
        Hidden            bool
        ReadOnly          string
        Searchable        bool
        Filter            bool
        ListDisplay       bool
        FormDisplay       bool
        CategoricalFilter bool
        Translations      []translation
        Choices           []Choice
        IsMethod          bool
        ErrMsg            string
        ProgressBar       map[float64]string
        LimitChoicesTo    func(interface{}, *User) []Choice
        UploadTo          string
    }

**uadmin.Filter**
^^^^^^^^^^^^^^^^^
Filter fetches records from the database.

Syntax:

.. code-block:: go

    Filter func(a interface{}, query interface{}, args ...interface{}) (err error)

**uadmin.FilterBuilder**
^^^^^^^^^^^^^^^^^^^^^^^^
FilterBuilder changes a map filter into a query.

Syntax:

.. code-block:: go

    FilterBuilder func(params map[string]interface{}) (query string, args []interface{})

**uadmin.GenerateBase32**
^^^^^^^^^^^^^^^^^^^^^^^^^
GenerateBase32 generates a base32 string of length.

Syntax:

.. code-block:: go

    GenerateBase32 func(length int) string

**uadmin.GenerateBase64**
^^^^^^^^^^^^^^^^^^^^^^^^^
GenerateBase64 generates a base64 string of length.

Syntax:

.. code-block:: go

    GenerateBase64 func(length int) string

**uadmin.Get**
^^^^^^^^^^^^^^
Get fetches the first record from the database.

Syntax:

.. code-block:: go

    Get func(a interface{}, query interface{}, args ...interface{}) (err error)

**uadmin.GetDB**
^^^^^^^^^^^^^^^^
GetDB returns a pointer to the DB.

Syntax:

.. code-block:: go

    GetDB func() *gorm.DB

**uadmin.GetID**
^^^^^^^^^^^^^^^^
GetID returns an ID number of a field.

Syntax:

.. code-block:: go

    GetID func(m.reflectValue) uint

**uadmin.GetString**
^^^^^^^^^^^^^^^^^^^^
GetString returns string representation on an instance of a model.

Syntax:

.. code-block:: go

    GetString func(a interface{}) string

**uadmin.GetUserFromRequest**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
GetUserFromRequest returns a user from a request.

Syntax:

.. code-block:: go

    GetUserFromRequest func(r *http.Request) *User

**uadmin.GroupPermission**
^^^^^^^^^^^^^^^^^^^^^^^^^^
GroupPermission sets the permission of a user group handled by an administrator.

Syntax:

.. code-block:: go

    type GroupPermission struct {
        Model
        DashboardMenu   DashboardMenu `gorm:"ForeignKey:DashboardMenuID" required:"true" filter:"true"`
        DashboardMenuID uint          `fk:"true" displayName:"DashboardMenu"`
        UserGroup       UserGroup     `gorm:"ForeignKey:UserGroupID" required:"true" filter:"true"`
        UserGroupID     uint          `fk:"true" displayName:"UserGroup"`
        Read            bool
        Add             bool
        Edit            bool
        Delete          bool
    }

**uadmin.HideInDashboarder**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^
HideInDashboarder is used to check if a model should be hidden in the dashboard.

Syntax:

.. code-block:: go

    type HideInDashboarder interface{
        HideInDashboard() bool
    }

**uadmin.INFO**
^^^^^^^^^^^^^^^
INFO is a data that is presented within a context that gives it meaning and relevance.

Syntax:

.. code-block:: go

    const INFO int = 2

**uadmin.IsAuthenticated**
^^^^^^^^^^^^^^^^^^^^^^^^^^
IsAuthenticated returns if the http.Request is authenticated or not.

Syntax:

.. code-block:: go

    IsAuthenticated func(r *http.Request) *Session

**uadmin.JSONMarshal**
^^^^^^^^^^^^^^^^^^^^^^
JSONMarshal returns the JSON encoding of v.

Syntax:

.. code-block:: go

    JSONMarshal func(v interface{}, safeEncoding bool) ([]byte, error)

**uadmin.Language**
^^^^^^^^^^^^^^^^^^^
Language is a system in uAdmin used to check and modify the settings of a language.

Syntax:

.. code-block:: go

    type Language struct {
        Model
        EnglishName    string `uadmin:"required;read_only;filter;search"`
        Name           string `uadmin:"required;read_only;filter;search"`
        Flag           string `uadmin:"image;list_exclude"`
        Code           string `uadmin:"filter;read_only;list_exclude"`
        RTL            bool   `uadmin:"list_exclude"`
        Default        bool   `uadmin:"help:Set as the default language;list_exclude"`
        Active         bool   `uadmin:"help:To show this in available languages;filter"`
        AvailableInGui bool   `uadmin:"help:The App is available in this language;read_only"`
    }

**uadmin.Log**
^^^^^^^^^^^^^^
Log is a system in uAdmin used to check the status of the user activities.

Syntax:

.. code-block:: go

    type Log struct {
        Model
        Username  string    `uadmin:"filter;read_only"`
        Action    Action    `uadmin:"filter;read_only"`
        TableName string    `uadmin:"filter;read_only"`
        TableID   int       `uadmin:"filter;read_only"`
        Activity  string    `uadmin:"code;read_only" gorm:"type:longtext"`
        RollBack  string    `uadmin:"link;"`
        CreatedAt time.Time `uadmin:"filter;read_only"`
    }

**uadmin.Login**
^^^^^^^^^^^^^^^^
Login returns the pointer of User and a bool for Is OTP Required.

Syntax:

.. code-block:: go

    Login func(r *http.Request, username string, password string) (*User, bool)

**uadmin.Login2FA**
^^^^^^^^^^^^^^^^^^^
Login2FA returns the pointer of User with a two-factor authentication.

Syntax:

.. code-block:: go

    Login2FA func(r *http.Request, username string, password string, otpPass string) *User

**uadmin.Logout**
^^^^^^^^^^^^^^^^^
Logout deactivates a session.

Syntax:

.. code-block:: go

    Logout func(r *http.Request)

**uadmin.MaxImageHeight**
^^^^^^^^^^^^^^^^^^^^^^^^^
MaxImageHeight sets the maximum height of an image.

Syntax:

.. code-block:: go

    MaxImageHeight int

**uadmin.MaxImageWidth**
^^^^^^^^^^^^^^^^^^^^^^^^
MaxImageWidth sets the maximum width of an image.

Syntax:

.. code-block:: go

    MaxImageWidth int

**uadmin.MaxUploadFileSize**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^
MaxUploadFileSize is the maximum upload file size in bytes.

Syntax:

.. code-block:: go

    MaxUploadFileSize int64

**uadmin.Model**
^^^^^^^^^^^^^^^^
Model is the standard struct to be embedded in any other struct to make it a model for uAdmin.

Syntax:

.. code-block:: go

    type Model struct {
	    ID        uint       `gorm:"primary_key"`
	    DeletedAt *time.Time `sql:"index"`
    }

**uadmin.ModelSchema**
^^^^^^^^^^^^^^^^^^^^^^
ModelSchema is a representation of a plan or theory in the form of an outline or model.

Syntax:

.. code-block:: go

    type ModelSchema struct {
        Name          string // Name of the Model
        DisplayName   string // Display Name of the model
        ModelName     string // URL
        ModelID       uint
        Inlines       []*ModelSchema
        InlinesData   []listData
        Fields        []F
        IncludeFormJS []string
        IncludeListJS []string
    }

**uadmin.MongoDB**
^^^^^^^^^^^^^^^^^^
MongoDB is an open source database management system (DBMS) that uses a document-oriented database model which supports various forms of data. [#f1]_ It is the active Mongo settings.

Syntax:

.. code-block:: go

    MongoDB *MongoSettings

**uadmin.MongoModel**
^^^^^^^^^^^^^^^^^^^^^
MongoModel is a uAdmin function for interfacing with MongoDB databases.

Syntax:

.. code-block:: go

    type MongoModel struct {
	    ID bson.ObjectId `bson:"_id,omitempty"`
    }

**uadmin.MongoSettings**
^^^^^^^^^^^^^^^^^^^^^^^^
MongoSettings is a feature that allows a user to configure the settings of a Mongo.

Syntax:

.. code-block:: go

    type MongoSettings struct {
        Name  string
        IP    string
        Debug bool
    }

**uadmin.NewModel**
^^^^^^^^^^^^^^^^^^^
NewModel creates a new model from a model name.

Syntax:

.. code-block:: go

    NewModel func(modelName string, pointer bool) (reflect.Value, bool)

**uadmin.NewModelArray**
^^^^^^^^^^^^^^^^^^^^^^^^
NewModelArray creates a new model from a model name.

Syntax:

.. code-block:: go

    NewModelArray func(modelName string, pointer bool) (reflect.Value, bool)

**uadmin.OK**
^^^^^^^^^^^^^
OK is a status to show that the application is doing well.

Syntax:

.. code-block:: go

    const OK int = 3

**uadmin.OTPAlgorithm**
^^^^^^^^^^^^^^^^^^^^^^^
OTPAlgorithm is the hashing algorithm of OTP.

Syntax:

.. code-block:: go

    OTPAlgorithm string

**uadmin.OTPDigits**
^^^^^^^^^^^^^^^^^^^^
OTPDigits is the number of degits for the OTP.

Syntax:

.. code-block:: go

    OTPDigits int

**uadmin.OTPPeriod**
^^^^^^^^^^^^^^^^^^^^
OTPPeriod is the number of seconds for the OTP to change.

Syntax:

.. code-block:: go

    OTPPeriod uint

**uadmin.OTPSkew**
^^^^^^^^^^^^^^^^^^
OTPSkew is the number of minutes to search around the OTP.

Syntax:

.. code-block:: go

    OTPSkew uint

**uadmin.PageLength**
^^^^^^^^^^^^^^^^^^^^^
PageLength is the list view max number of records.

Syntax:

.. code-block:: go

    PageLength int

**uadmin.Port**
^^^^^^^^^^^^^^^
Port is the port used for http or https server.

Syntax:

.. code-block:: go

    Port int

**uadmin.Preload**
^^^^^^^^^^^^^^^^^^
Preload accesses the information of the fields in another model.

Syntax:

.. code-block:: go

    func(a interface{}, preload ...string) (err error)

**uadmin.PublicMedia**
^^^^^^^^^^^^^^^^^^^^^^
PublicMedia allows public access to media handler without authentication.

Syntax:

.. code-block:: go

    PublicMedia bool

**uadmin.Register**
^^^^^^^^^^^^^^^^^^^
Register is used to register models to uAdmin.

Syntax:

.. code-block:: go

    Register func(m ...interface{})

**uadmin.RegisterInlines**
^^^^^^^^^^^^^^^^^^^^^^^^^^
RegisterInlines is a function to register a model as an inline for another model

Syntax:

.. code-block:: go

    func RegisterInlines(model interface{}, fk map[string]string)

Parameters:

    **model (struct instance):** Is the model that you want to add inlines to.

    **fk (map[interface{}]string):** This is a map of the inlines to be added to the model. The map's key is the name of the model of the inline and the value of the map is the foreign key field's name.

Example:

.. code-block:: go

    type Person struct {
        uadmin.Model
        Name string
    }

    type Card struct {
        uadmin.Model
        PersonID uint
        Person   Person
    }

    func main() {
        // ...
        uadmin.RegisterInlines(Person{}, map[string]string{
            "Card": "PersonID",
        })
        // ...
    }

**uadmin.ReportingLevel**
^^^^^^^^^^^^^^^^^^^^^^^^^
ReportingLevel is the standard reporting level.

Syntax:

.. code-block:: go

    ReportingLevel int

**uadmin.ReportTimeStamp**
^^^^^^^^^^^^^^^^^^^^^^^^^^
ReportTimeStamp set this to true to hav a time stamp in your logs.

Syntax:

.. code-block:: go

    ReportTimeStamp bool

**uadmin.ReturnJSON**
^^^^^^^^^^^^^^^^^^^^^
ReturnJSON returns JSON to the client.

Syntax:

.. code-block:: go

    ReturnJSON func(w http.ResponseWriter, r *http.Request, v interface{})

**uadmin.RootURL**
^^^^^^^^^^^^^^^^^^
RootURL is where the listener is mapped to.

Syntax:

.. code-block:: go

    RootURL string

**uadmin.Salt**
^^^^^^^^^^^^^^^
Salt is extra salt added to password hashing.

Syntax:

.. code-block:: go

    Salt string

**uadmin.Save**
^^^^^^^^^^^^^^^
Save saves the object in the database.

Syntax:

.. code-block:: go

    Save func(a interface{}) (err error)

**uadmin.Schema**
^^^^^^^^^^^^^^^^^
Schema is the gblobal schema of the system.

Syntax:

.. code-block:: go

    Schema map[string]ModelSchema

**uadmin.SendEmail**
^^^^^^^^^^^^^^^^^^^^
SendEmail sends email using system configured variables.

Syntax:

.. code-block:: go

    SendEmail func(to, cc, bcc []string, subject, body string) (err error)

**uadmin.Session**
^^^^^^^^^^^^^^^^^^
Session is an activity that a user with a unique IP address spends on a Web site during a specified period of time. [#f2]_

Syntax:

.. code-block:: go

    type Session struct {
        Model
        Key        string
        User       User `gorm:"ForeignKey:UserID" uadmin:"filter"`
        UserID     uint `fk:"true" displayName:"User"`
        LoginTime  time.Time
        LastLogin  time.Time
        Active     bool   `uadmin:"filter"`
        IP         string `uadmin:"filter"`
        PendingOTP bool   `uadmin:"filter"`
        ExpiresOn  *time.Time
    }

**uadmin.SiteName**
^^^^^^^^^^^^^^^^^^^
SiteName is the name of the website that shows on title and dashboard.

Syntax:

.. code-block:: go

    SiteName string

**uadmin.StartSecureServer**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^
StartServer is the process of activating a uAdmin server using a localhost IP or an apache with SSL certificate and a private key.

Syntax:

.. code-block:: go

    StartSecureServer func(certFile, keyFile string)

**uadmin.StartServer**
^^^^^^^^^^^^^^^^^^^^^^
StartServer is the process of activating a uAdmin server using a localhost IP or an apache.

Syntax:

.. code-block:: go

    StartServer func()

**uadmin.Tf**
^^^^^^^^^^^^^
Tf is a function for translating strings into any given language.

Syntax:

.. code-block:: go

    Tf func(path string, lang string, term string, args ...interface{}) string

Parameters:

    **path (string):** This is where to get the translation from. It is in the
    format of "GROUPNAME/FILENAME" for example: "uadmin/system"

    **lang (string):** Is the language code. If empty string is passed we will use
    the default language.

    **term (string):** The term to translate.

    **args (...interface{}):** Is a list of args to fill the term with place holders

**uadmin.Theme**
^^^^^^^^^^^^^^^^
Theme is the name of the theme used in uAdmin.

Syntax:

.. code-block:: go

    Theme string

**uadmin.Trail**
^^^^^^^^^^^^^^^^
Trail prints to the log.

Syntax:

.. code-block:: go

    Trail func(level int, msg interface{}, i ...interface{})

**uadmin.Translate**
^^^^^^^^^^^^^^^^^^^^
Translate is used to get a translation from a multilingual fields.

Syntax:

.. code-block:: go

    Translate func(raw string, lang string, args ...bool) string

**uadmin.Update**
^^^^^^^^^^^^^^^^^
Update updates the field name and value of an interface.

Syntax:

.. code-block:: go

    Update func(a interface{}, fieldName string, value interface{}, query string, args ...interface{}) (err error)

**uadmin.UploadImageHandler**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
UploadImageHandler handles the uploading process of an image.

Syntax:

.. code-block:: go

    UploadImageHandler func(w http.ResponseWriter, r *http.Request, session *Session)

**uadmin.User**
^^^^^^^^^^^^^^^
User is a system in uAdmin used to check and modify the settings of a user.

Syntax:

.. code-block:: go

    type User struct {
        Model
        Username     string    `uadmin:"required;filter"`
        FirstName    string    `uadmin:"filter"`
        LastName     string    `uadmin:"filter"`
        Password     string    `uadmin:"required;password;help:To reset password, clear the field and type a new password.;list_exclude"`
        Email        string    `uadmin:"email"`
        Active       bool      `uadmin:"filter"`
        Admin        bool      `uadmin:"filter"`
        RemoteAccess bool      `uadmin:"filter"`
        UserGroup    UserGroup `uadmin:"filter"`
        UserGroupID  uint
        Photo        string `uadmin:"image"`
        LastLogin   *time.Time `uadmin:"read_only"`
        ExpiresOn   *time.Time
        OTPRequired bool
        OTPSeed     string `uadmin:"list_exclude;hidden;read_only"`
    }

**uadmin.UserGroup**
^^^^^^^^^^^^^^^^^^^^
UserGroup is a system in uADmin used to add, modify, and delete the group name. 

Syntax:

.. code-block:: go

    type UserGroup struct {
        Model
        GroupName string `uadmin:"filter"`
    }

**uadmin.UserPermission**
^^^^^^^^^^^^^^^^^^^^^^^^^
UserPermission sets the permission of a user handled by an administrator.

Syntax:

.. code-block:: go

    type UserPermission struct {
        Model
        DashboardMenu   DashboardMenu `gorm:"ForeignKey:DashboardMenuID" required:"true" filter:"true" uadmin:"filter"`
        DashboardMenuID uint          `fk:"true" displayName:"DashboardMenu"`
        User            User          `gorm:"ForeignKey:UserID" required:"true" filter:"true" uadmin:"filter"`
        UserID          uint          `fk:"true" displayName:"User"`
        Read            bool          `uadmin:"filter"`
        Add             bool          `uadmin:"filter"`
        Edit            bool          `uadmin:"filter"`
        Delete          bool          `uadmin:"filter"`
    }

**uadmin.Version**
^^^^^^^^^^^^^^^^^^
Version number as per Semantic Versioning 2.0.0 (semver.org)

Syntax:

.. code-block:: go

    const Version string = "0.1.0-alpha"

**uadmin.WARNING**
^^^^^^^^^^^^^^^^^^
WARNING is a statement or event that indicates a possible problems occurring in an application.

Syntax:

.. code-block:: go

    const WARNING int = 4

**uadmin.WORKING**
^^^^^^^^^^^^^^^^^^
OK is a status to show that the application is working.

Syntax:

.. code-block:: go

    const WORKING int = 1


Reference
---------
.. [#f1] Rouse, Margaret (2018). MongoDB. Retrieved from https://searchdatamanagement.techtarget.com/definition/MongoDB
.. [#f2] QuinStreet Inc. (2018). User Session. Retrieved from https://www.webopedia.com/TERM/U/user_session.html