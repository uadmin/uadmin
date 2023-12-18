package uadmin

import (
	"net/http"
	"os"
	"regexp"
)

// Constants

// cPOST post
const cPOST = "POST"

// cID true
const cID = "id"

// cTRUE true
const cTRUE = "true"

// cJPG jpg
const cJPG = "jpg"

// cJPEG jpeg
const cJPEG = "jpeg"

// cPNG png
const cPNG = "png"

// cGIF gif
const cGIF = "gif"

// cSTRING !
const cSTRING = "string"

// cNUMBER !
const cNUMBER = "number"

// cDATE !
const cDATE = "date"

// cBOOL !
const cBOOL = "bool"

// cLIST !
const cLIST = "list"

// cIMAGE !
const cIMAGE = "image"

// cFK !
const cFK = "fk"

// cLINK !
const cLINK = "link"

// cMONRY !
const cMONEY = "money"

// cCODE !
const cCODE = "code"

// cHTML !
const cHTML = "html"

// cMULTILINGUAL !
const cMULTILINGUAL = "multilingual"

// cPROGRESSBAR !
const cPROGRESSBAR = "progress_bar"

// cPASSWORD !
const cPASSWORD = "password"

// cFILE !
const cFILE = "file"

// cEMAIL !
const cEMAIL = "email"

// cM2M !
const cM2M = "m2m"

// Version number as per Semantic Versioning 2.0.0 (semver.org)
const Version = "0.10.1"

// VersionCodeName is the cool name we give to versions with significant changes.
// This name should always be a bug's name starting from A-Z them revolving back.
// This started at version 0.5.0 (Atlas Moth)
// 0.6.0 Beetle
// 0.7.0 Catterpiller
// 0.8.0 Dragonfly
// 0.9.0 Gnat
// 0.10.0 Gnat
const VersionCodeName = "Housefly"

// Public Global Variables

// DefaultLang is the default language of the system.
var defaultLang Language

// Theme is the name of the theme used in uAdmin.
var Theme = "default"

// SiteName is the name of the website that shows on title and dashboard.
var SiteName = "uAdmin"

// ReportingLevel is the standard reporting level.
var ReportingLevel = DEBUG

// ReportTimeStamp set this to true to have a time stamp in your logs.
var ReportTimeStamp = false

// DebugDB prints all SQL statements going to DB.
var DebugDB = false

// Schema is the global schema of the system.
var Schema map[string]ModelSchema

// PageLength is the list view max number of records.
var PageLength = 100

// MaxImageHeight sets the maximum height of an image.
var MaxImageHeight = 600

// MaxImageWidth sets the maximum width of an image.
var MaxImageWidth = 800

// MaxUploadFileSize is the maximum upload file size in bytes.
var MaxUploadFileSize = int64(25 * 1024 * 1024)

// BindIP is the IP the application listens to.
var BindIP = ""

// Port is the port used for http or https server.
var Port = 8080

// EmailFrom identifies where the email is coming from.
var EmailFrom string

// EmailUsername sets the username of an email.
var EmailUsername string

// EmailPassword sets the password of an email.
var EmailPassword string

// EmailSMTPServer sets the name of the SMTP Server in an email.
var EmailSMTPServer string

// EmailSMTPServerPort sets the port number of an SMTP Server in an email.
var EmailSMTPServerPort int

// RootURL is where the listener is mapped to.
var RootURL = "/"

// OTPAlgorithm is the hashing algorithm of OTP.
var OTPAlgorithm = "sha1"

// OTPDigits is the number of digits for the OTP.
var OTPDigits = 6

// OTPPeriod is the number of seconds for the OTP to change.
var OTPPeriod = uint(30)

// OTPSkew is the number of minutes to search around the OTP.
var OTPSkew = uint(5)

// PublicMedia allows public access to media handler without authentication.
var PublicMedia = false

// EncryptKey is a key for encryption and decryption of data in the DB.
var EncryptKey = []byte{}

// LogDelete adds a log when a record is deleted.
var LogDelete = true

// LogAdd adds a log when a record is added.
var LogAdd = true

// LogEdit adds a log when a record is edited.
var LogEdit = true

// LogRead adds a log when a record is read.
var LogRead = false

// CacheTranslation allows a translation to store data in a cache memory.
var CacheTranslation = false

// DefaultMediaPermission is the default permission applied to to files uploaded to the system
var DefaultMediaPermission = os.FileMode(0644)

// ErrorHandleFunc is a function that will be called everytime Trail is called. It receives
// one parameter for error level, one for error message and one for runtime stack trace
var ErrorHandleFunc func(int, string, string)

// AllowedIPs is a list of allowed IPs to access uAdmin interfrace in one of the following formats:
// - "*" = Allow all
// - "" = Allow none
// - "192.168.1.1" Only allow this IP
// - "192.168.1.0/24" Allow all IPs from 192.168.1.1 to 192.168.1.254
// You can also create a list of the above formats using comma to separate them.
// For example: "192.168.1.1,192.168.1.2,192.168.0.0/24"
var AllowedIPs = "*"

// BlockedIPs is a list of blocked IPs from accessing uAdmin interfrace in one of the following formats:
// - "*" = Block all
// - "" = Block none
// - "192.168.1.1" Only block this IP
// - "192.168.1.0/24" Block all IPs from 192.168.1.1 to 192.168.1.254
// You can also create a list of the above formats using comma to separate them.
// For example: "192.168.1.1,192.168.1.2,192.168.0.0/24"
var BlockedIPs = ""

// RestrictSessionIP is to block access of a user if their IP changes from their original IP during login
var RestrictSessionIP = false

// RetainMediaVersions is to allow the system to keep files uploaded even after they are changed.
// This allows the system to "Roll Back" to an older version of the file.
var RetainMediaVersions = true

// ApprovalHandleFunc is a function that could be called during the save process of each approval
var ApprovalHandleFunc func(*Approval) bool

// RateLimit is the maximum number of requests/second for any unique IP
var RateLimit int64 = 3

// RateLimitBurst is the maximum number of requests for an idle user
var RateLimitBurst int64 = 3

// OptimizeSQLQuery selects columns during rendering a form a list to visible fields.
// This means during the filtering of a form the select statement will not include
// any field with `hidden` tag. For list it will not select any field with `list_exclude`
var OptimizeSQLQuery = false

// APILogRead controls the data API's logging for read commands.
var APILogRead = false

// APILogEdit controls the data API's logging for edit commands.
var APILogEdit = true

// APILogAdd controls the data API's logging for add commands.
var APILogAdd = true

// APILogDelete controls the data API's logging for delete commands.
var APILogDelete = true

// APILogSchema controls the data API's logging for schema commands.
var APILogSchema = true

// APIPublicRead controls the data API’s public for add commands.
var APIPublicRead = false

// APIPublicEdit controls the data API's public for edit commands.
var APIPublicEdit = false

// APIPublicAdd controls the data API's public for add commands.
var APIPublicAdd = false

// APIPublicDelete controls the data API's public for delete commands.
var APIPublicDelete = false

// APIPublicSchema controls the data API's public for schema commands.
var APIPublicSchema = false

// APIDisabledRead controls the data API’s disabled for add commands.
var APIDisabledRead = false

// APIDisabledEdit controls the data API's disabled for edit commands.
var APIDisabledEdit = false

// APIDisabledAdd controls the data API's disabled for add commands.
var APIDisabledAdd = false

// APIDisabledDelete controls the data API's disabled for delete commands.
var APIDisabledDelete = false

// APIDisabledSchema controls the data API's disabled for schema commands.
var APIDisabledSchema = false

// APIPreQueryRead controls the data API’s pre query for add commands.
var APIPreQueryRead = false

// APIPreQueryEdit controls the data API's pre query for edit commands.
var APIPreQueryEdit = false

// APIPreQueryAdd controls the data API's pre query for add commands.
var APIPreQueryAdd = false

// APIPreQueryDelete controls the data API's pre query for delete commands.
var APIPreQueryDelete = false

// APIPreQuerySchema controls the data API's pre query for schema commands.
var APIPreQuerySchema = false

// APIPostQueryRead controls the data API’s post query for add commands.
var APIPostQueryRead = false

// APIPostQueryEdit controls the data API's post query for edit commands.
var APIPostQueryEdit = false

// APIPostQueryAdd controls the data API's post query for add commands.
var APIPostQueryAdd = false

// APIPostQueryDelete controls the data API's post query for delete commands.
var APIPostQueryDelete = false

// APIPostQuerySchema controls the data API's post query for schema commands.
var APIPostQuerySchema = false

// LogHTTPRequests logs http requests to syslog
var LogHTTPRequests = true

/*
HTTPLogFormat is the format used to log HTTP access
%a: Client IP address
%A: Server hostname/IP
%{local}p: Server port
%U: Path
%c: All cookies
%{NAME}c: Cookie named 'NAME'
%{GET}f: GET request parameters
%{POST}f: POST request parameters
%B: Response length
%>s: Response code
%D: Time taken in microseconds
%T: Time taken in seconds
%I: Request length
*/
var HTTPLogFormat = "%a %>s %B %U %D"

// LogTrail stores Trail logs to syslog
var LogTrail = false

// TrailLoggingLevel is the minimum level to be logged into syslog
var TrailLoggingLevel = INFO

// SystemMetrics enables uAdmin system metrics to be recorded
var SystemMetrics = false

// UserMetrics enables the user metrics to be recorded
var UserMetrics = false

// CacheSessions allows uAdmin to store sessions data in memory
var CacheSessions = true

// CachePermissions allows uAdmin to store permissions data in memory
var CachePermissions = true

// PasswordAttempts is the maximum number of invalid password attempts before
// the IP address is blocked for some time from using the system
var PasswordAttempts = 5

// PasswordTimeout is the amount of time in minutes the IP will be blocked for after
// reaching the the maximum invalid password attempts
var PasswordTimeout = 15

// AllowedHosts is a comma separated list of allowed hosts for the server to work. The
// default value is only for development. Production domain should be added before
// deployment
var AllowedHosts = "0.0.0.0,127.0.0.1,localhost,::1"

// Logo is the main logo that shows on uAdmin UI
var Logo = "/static/uadmin/logo.png"

// FavIcon is the fav icon that shows on uAdmin UI
var FavIcon = "/static/uadmin/favicon.ico"

// DisableAdminUI disables access to the UI and all related APIs
var DisableAdminUI = false

// TimeZone specifies the the default time zone for the system. Valid values include
// IANA Time Zones like America/New_York, Asia/Riyadh and UTC. To get local time zone
// use "local"
var TimeZone = "local"

// TrailCacheSize is the number of bytes to keep in memory of trail logs
var TrailCacheSize = 65536

// DisableDAPIAuth enables or disables access to dAPI auth API
var DisableDAPIAuth = true

// AllowDAPISignup allows unauthenticated users to sign up
var AllowDAPISignup = false

// DAPISignupGroupID is the default user group id new users get when they sign up
// to leave new signed up users without a group, use value 0 for this variable
var DAPISignupGroupID = 0

// DAPISignupActive controls if new signed up users are activate automatically
var DAPISignupActive = true

// DAPISignupAllowRemote controls if new signed up users are can login over the internet
var DAPISignupAllowRemote = true

// SignupValidationHandler can be used to validate or customize new
// signed up users. Note that the password in the password field
// is passed in plain text. Do not plain text passwords anywhere.
var SignupValidationHandler func(user *User) error

// CustomResetPasswordLink is the link sent to the user's email to reset their password
// the string may include the following place holder:
// "{SCHEMA}://{HOST}/resetpassword?u={USER_ID}&key={OTP}"
var CustomResetPasswordLink = ""

// ResetPasswordMessage is a message that can be sent to the user email when
// a password reset request sends an email. This message may include the
// following place holders:
// {NAME}: user real name
// {WEBSITE}: website name
// {URL}: link to the password reset page
var ResetPasswordMessage = `Dear {NAME},

Have you forgotten your password to access {WEBSITE}. Don't worry we got your back. Please follow the link below to reset your password.

If you want to reset your password, click this link:
<a href="{URL}">{URL}</a>

If you didn't request a password reset, you can ignore this message.

Regards,
{WEBSITE} Support
`

// CustomEmailHandler allows to customize or even override the default email sending method.
// The return of of the function is a boolean and an error. All parameters are of pointers to
// allow customization of the parameters that will be passed to the default email sender.
//   - The boolean will determine whether to proceed or not. If true, the process will proceed with
//     the default method of sending the email
//   - The error will be reported to Trail as type uadmin.ERROR
var CustomEmailHandler func(to, cc, bcc *[]string, subject, body *string, attachments ...*string) (bool, error)

// CustomDAPIHeaders are extra handlers that would be added to dAPI responses
var CustomDAPIHeaders = map[string]string{}

// EnableDAPICORS controller whether dAPI is uses CORS protocol to allow cross-origan requests
var EnableDAPICORS bool

// AllowedCORSOrigins is a list of allowed CORS origins
var AllowedCORSOrigins []string

// CustomizeJSON is a function to allow customization of JSON returns
var CustomizeJSON func(http.ResponseWriter, *http.Request, interface{}, []byte) []byte

// CustomDAPILoginHandler is a function that can provide extra information
// in the login return
var CustomDAPILoginHandler func(*http.Request, *User, map[string]interface{}) map[string]interface{}

// CustomDAPISignupHandler
var CustomDAPISignupHandler func(*http.Request, *User) error

// FullMediaURL allows uAdmin to send you full path URL instead on relative
// path for dAPI read requests
var FullMediaURL = false

// MaskPasswordInAPI will replace any password fields with an asterisk mask
var MaskPasswordInAPI = true

// APIPreQueryReadHandler is a function that runs before all dAPI read requests
var APIPreQueryReadHandler func(http.ResponseWriter, *http.Request) bool

// APIPostQueryReadHandler is a function that runs after all dAPI read requests
var APIPostQueryReadHandler func(http.ResponseWriter, *http.Request, map[string]interface{}) bool

// APIPreQueryAddHandler is a function that runs before all dAPI add requests
var APIPreQueryAddHandler func(http.ResponseWriter, *http.Request) bool

// APIPostQueryAddHandler is a function that runs after all dAPI add requests
var APIPostQueryAddHandler func(http.ResponseWriter, *http.Request, map[string]interface{}) bool

// APIPreQueryEditHandler is a function that runs before all dAPI edit requests
var APIPreQueryEditHandler func(http.ResponseWriter, *http.Request) bool

// APIPostQueryEditHandler is a function that runs after all dAPI edit requests
var APIPostQueryEditHandler func(http.ResponseWriter, *http.Request, map[string]interface{}) bool

// APIPreQueryDeleteHandler is a function that runs before all dAPI delete requests
var APIPreQueryDeleteHandler func(http.ResponseWriter, *http.Request) bool

// APIPostQueryDeleteHandler is a function that runs after all dAPI delete requests
var APIPostQueryDeleteHandler func(http.ResponseWriter, *http.Request, map[string]interface{}) bool

// PreLoginHandler is a function that runs after all dAPI delete requests
var PreLoginHandler func(r *http.Request, username string, password string)

// PostHandler is a function that runs after the files are uploaded
var PostUploadHandler func(filePath string, modelName string, f *F) string

// CompressJSON is a variable that allows the user to reduce the size of JSON responses
var CompressJSON = false

// CompressJSON is a variable that allows the user to reduce the size of JSON responses
var RemoveZeroValueJSON = false

// SSOURL enables SSO using OpenID Connect
var SSOURL = ""

// Private Global Variables
// Regex
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// Global active languages
var activeLangs []Language

// langMapCache is a computer memory used for storage of frequently or recently used translations.
var langMapCache = map[string][]byte{}

// Models is where we keep all registered models
var models map[string]interface{}

// Inlines is where we keep all registered models' inlines
var inlines map[string][]interface{}

// ForeignKeys is the link between models' and their inlines
var foreignKeys map[string]map[string]string

var registered = false

var handlersRegistered = false

var defaultProgressBarColor = "#07c"

var settingsSynched = false
