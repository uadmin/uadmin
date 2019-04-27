package uadmin

import (
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
const Version = "0.1.0"

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

// EncryptKey is a key for encyption and decryption of data in the DB.
var EncryptKey = []byte{}

// LogDelete adds a log when a record is deleted.
var LogDelete = true

// LogAdd adds a log when a record is added.
var LogAdd = true

// LogEdit adds a log when a record is edited.
var LogEdit = true

// LogRead adds a log when a record is read.
var LogRead = false

// LangMapCache is a computer memory used for storage of frequently or recently used translations.
var LangMapCache = map[string][]byte{}

// CacheTranslation allows a translation to store data in a cache memory.
var CacheTranslation = false

// Private Global Variables
// Regex
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// Global active languages
var activeLangs []Language

// Models is where we keep all registered models
var models map[string]interface{}

// Inlines is where we keep all registered models' inlines
var inlines map[string][]interface{}

// ForeignKeys is the link between models' and their inlines
var foreignKeys map[string]map[string]string

// Menu ?
var menu []interface{}

var registered = false

var handlersRegistered = false

var defaultProgressBarColor = "#07c"
