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
const Version = "0.1.0-beta.2"

// Public Global Variables

// DefaultLang is the default language of the system
var defaultLang Language

// Theme is the name of the theme used in uAdmin
var Theme = "default"

// SiteName is the name of the website that shows on title and dashboard
var SiteName = "uAdmin"

// ReportingLevel is the standard reporting level
var ReportingLevel = DEBUG

// ReportTimeStamp set this to true to hav a time stamp in your logs
var ReportTimeStamp = false

// DebugDB prints all SQL statements going to DB
var DebugDB = false

// Schema is the gblobal schema of the system
var Schema map[string]ModelSchema

// PageLength is the list view max number of records
var PageLength = 100

// MaxImageHeight !
var MaxImageHeight = 600

// MaxImageWidth !
var MaxImageWidth = 800

// MaxUploadFileSize is the maximum upload file size in bytes
var MaxUploadFileSize = int64(25 * 1024 * 1024)

// BindIP is the IP the application listens to
var BindIP = ""

// Port is the port used for http or https server
var Port = 8080

// EmailFrom email from
var EmailFrom string

// EmailUsername !
var EmailUsername string

// EmailPassword !
var EmailPassword string

// EmailSMTPServer !
var EmailSMTPServer string

// EmailSMTPServerPort !
var EmailSMTPServerPort int

// RootURL is where the listener is mapped to
var RootURL = "/"

// OTPAlgorithm is the hashing algorithm of OTP
var OTPAlgorithm = "sha1"

// OTPDigits is the number of degits for the OTP
var OTPDigits = 6

// OTPPeriod is the number of seconds for the OTP to change
var OTPPeriod = uint(30)

// OTPSkew is the number of minutes to search around the OTP
var OTPSkew = uint(5)

// PublicMedia allows public access to media handler without authentication
var PublicMedia = false

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
