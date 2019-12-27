package uadmin

import (
	"fmt"
	"github.com/uadmin/uadmin/colors"
	"strconv"
	"strings"
	"time"
)

// DataType is a list of data types used for settings
type DataType int

// String is a type
func (DataType) String() DataType {
	return 1
}

// Integer is a type
func (DataType) Integer() DataType {
	return 2
}

// Float is a type
func (DataType) Float() DataType {
	return 3
}

// Boolean is a type
func (DataType) Boolean() DataType {
	return 4
}

// File is a type
func (DataType) File() DataType {
	return 5
}

// Image is a type
func (DataType) Image() DataType {
	return 6
}

// DateTime is a type
func (DataType) DateTime() DataType {
	return 7
}

// Setting model stored system settings
type Setting struct {
	Model
	Name         string `uadmin:"required;filter;search"`
	DefaultValue string
	DataType     DataType `uadmin:"required;filter"`
	Value        string
	Help         string          `uadmin:"search" sql:"type:text;"`
	Category     SettingCategory `uadmin:"required;filter"`
	CategoryID   uint
	Code         string `uadmin:"read_only;search"`
}

// Save overides save
func (s *Setting) Save() {
	Preload(s)
	s.Code = strings.Replace(s.Category.Name, " ", "", -1) + "." + strings.Replace(s.Name, " ", "", -1)
	s.ApplyValue()
	Save(s)
}

// ParseFormValue takes the value of a setting from an HTTP request and saves in the instance of setting
func (s *Setting) ParseFormValue(v []string) {
	switch s.DataType {
	case s.DataType.Boolean():
		tempV := len(v) == 1 && v[0] == "on"
		if tempV {
			s.Value = "1"
		} else {
			s.Value = "0"
		}
	case s.DataType.DateTime():
		if len(v) == 1 && v[0] != "" {
			s.Value = v[0] + ":00"
		} else {
			s.Value = ""
		}
	default:
		if len(v) == 1 && v[0] != "" {
			s.Value = v[0]
		} else {
			s.Value = ""
		}
	}
}

// GetValue returns an interface representing the value of the setting
func (s *Setting) GetValue() interface{} {
	var err error
	var v interface{}

	switch s.DataType {
	case s.DataType.String():
		if s.Value == "" {
			v = s.DefaultValue
		} else {
			v = s.Value
		}
	case s.DataType.Integer():
		if s.Value != "" {
			v, err = strconv.ParseInt(s.Value, 10, 64)
			v = int(v.(int64))
		}
		if err != nil {
			v, err = strconv.ParseInt(s.DefaultValue, 10, 64)
		}
		if err != nil {
			v = 0
		}
	case s.DataType.Float():
		if s.Value != "" {
			v, err = strconv.ParseFloat(s.Value, 64)
		}
		if err != nil {
			v, err = strconv.ParseFloat(s.DefaultValue, 64)
		}
		if err != nil {
			v = 0.0
		}
	case s.DataType.Boolean():
		if s.Value != "" {
			v = s.Value == "1"
		}
		if v == nil {
			v = s.DefaultValue == "1"
		}
	case s.DataType.File():
		if s.Value == "" {
			v = s.DefaultValue
		} else {
			v = s.Value
		}
	case s.DataType.Image():
		if s.Value == "" {
			v = s.DefaultValue
		} else {
			v = s.Value
		}
	case s.DataType.DateTime():
		if s.Value != "" {
			v, err = time.Parse("2006-01-02 15:04:05", s.Value)
		}
		if err != nil {
			v, err = time.Parse("2006-01-02 15:04:05", s.DefaultValue)
		}
		if err != nil {
			v = time.Now()
		}
	}
	return v
}

// ApplyValue changes uAdmin global variables' value based in the setting value
func (s *Setting) ApplyValue() {
	v := s.GetValue()

	switch s.Code {
	case "uAdmin.Theme":
		Theme = v.(string)
	case "uAdmin.SiteName":
		SiteName = v.(string)
	case "uAdmin.ReportingLevel":
		ReportingLevel = v.(int)
	case "uAdmin.ReportTimeStamp":
		ReportTimeStamp = v.(bool)
	case "uAdmin.DebugDB":
		if DebugDB != v.(bool) {
			DebugDB = v.(bool)
			db.LogMode(DebugDB)
		}
	case "uAdmin.PageLength":
		PageLength = v.(int)
	case "uAdmin.MaxImageHeight":
		MaxImageHeight = v.(int)
	case "uAdmin.MaxImageWidth":
		MaxImageWidth = v.(int)
	case "uAdmin.MaxUploadFileSize":
		MaxUploadFileSize = int64(v.(int))
	case "uAdmin.Port":
		Port = v.(int)
	case "uAdmin.EmailFrom":
		EmailFrom = v.(string)
	case "uAdmin.EmailUsername":
		EmailUsername = v.(string)
	case "uAdmin.EmailPassword":
		EmailPassword = v.(string)
	case "uAdmin.EmailSMTPServer":
		EmailSMTPServer = v.(string)
	case "uAdmin.EmailSMTPServerPort":
		EmailSMTPServerPort = v.(int)
	case "uAdmin.RootURL":
		RootURL = v.(string)
	case "uAdmin.OTPAlgorithm":
		OTPAlgorithm = v.(string)
	case "uAdmin.OTPDigits":
		OTPDigits = v.(int)
	case "uAdmin.OTPPeriod":
		OTPPeriod = uint(v.(int))
	case "uAdmin.OTPSkew":
		OTPSkew = uint(v.(int))
	case "uAdmin.PublicMedia":
		PublicMedia = v.(bool)
	case "uAdmin.LogDelete":
		LogDelete = v.(bool)
	case "uAdmin.LogAdd":
		LogAdd = v.(bool)
	case "uAdmin.LogEdit":
		LogEdit = v.(bool)
	case "uAdmin.LogRead":
		LogRead = v.(bool)
	case "uAdmin.CacheTranslation":
		CacheTranslation = v.(bool)
	case "uAdmin.AllowedIPs":
		AllowedIPs = v.(string)
	case "uAdmin.BlockedIPs":
		BlockedIPs = v.(string)
	case "uAdmin.RestrictSessionIP":
		RestrictSessionIP = v.(bool)
	case "uAdmin.RetainMediaVersions":
		RetainMediaVersions = v.(bool)
	case "uAdmin.RateLimit":
		if RateLimit != int64(v.(int)) {
			RateLimit = int64(v.(int))
			rateLimitMap = map[string]int64{}
		}
	case "uAdmin.RateLimitBurst":
		RateLimitBurst = int64(v.(int))
	case "uAdmin.OptimizeSQLQuery":
		OptimizeSQLQuery = v.(bool)
	case "uAdmin.APILogRead":
		APILogRead = v.(bool)
	case "uAdmin.APILogEdit":
		APILogEdit = v.(bool)
	case "uAdmin.APILogAdd":
		APILogAdd = v.(bool)
	case "uAdmin.APILogDelete":
		APILogDelete = v.(bool)
	case "uAdmin.APILogSchema":
		APILogSchema = v.(bool)
	case "uAdmin.LogHTTPRequests":
		LogHTTPRequests = v.(bool)
	case "uAdmin.HTTPLogFormat":
		HTTPLogFormat = v.(string)
	case "uAdmin.LogTrail":
		LogTrail = v.(bool)
	case "uAdmin.TrailLoggingLevel":
		TrailLoggingLevel = v.(int)
	case "uAdmin.SystemMetrics":
		SystemMetrics = v.(bool)
	case "uAdmin.UserMetrics":
		UserMetrics = v.(bool)
	case "uAdmin.CacheSessions":
		CacheSessions = v.(bool)
		if CacheSessions {
			loadSessions()
		}
	case "uAdmin.CachePermissions":
		CachePermissions = v.(bool)
		if CachePermissions {
			loadPermissions()
		}
	}
}

// GetSetting return the value of a setting based on its code
func GetSetting(code string) interface{} {
	s := Setting{}
	Get(&s, "code = ?", code)

	if s.ID == 0 {
		return nil
	}
	return s.GetValue()
}

func syncSystemSettings() {
	// Check if the uAdmin category is not there and add it
	cat := SettingCategory{}
	Get(&cat, "Name = ?", "uAdmin")
	if cat.ID == 0 {
		cat = SettingCategory{Name: "uAdmin"}
		Save(&cat)
	}

	t := DataType(0)

	settings := []Setting{
		{
			Name:         "Theme",
			Value:        Theme,
			DefaultValue: "default",
			DataType:     t.String(),
			Help:         "is the name of the theme used in uAdmin",
		},
		{
			Name:         "Site Name",
			Value:        SiteName,
			DefaultValue: "uAdmin",
			DataType:     t.String(),
			Help:         "is the name of the website that shows on title and dashboard",
		},
		{
			Name:         "Reporting Level",
			Value:        fmt.Sprint(ReportingLevel),
			DefaultValue: "0",
			DataType:     t.Integer(),
			Help:         "Reporting level. DEBUG=0, WORKING=1, INFO=2, OK=3, WARNING=4, ERROR=5",
		},
		{
			Name:         "Report Time Stamp",
			Value:        fmt.Sprint(ReportTimeStamp),
			DefaultValue: "0",
			DataType:     t.Boolean(),
			Help:         "set this to true to have a time stamp in your logs",
		},
		{
			Name: "Debug DB",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(DebugDB),
			DefaultValue: "0",
			DataType:     t.Boolean(),
			Help:         "prints all SQL statements going to DB",
		},
		{
			Name:         "Page Length",
			Value:        fmt.Sprint(PageLength),
			DefaultValue: "100",
			DataType:     t.Integer(),
			Help:         "is the list view max number of records",
		},
		{
			Name:         "Max Image Height",
			Value:        fmt.Sprint(MaxImageHeight),
			DefaultValue: "600",
			DataType:     t.Integer(),
			Help:         "sets the maximum height of an Image",
		},
		{
			Name:         "Max Image Width",
			Value:        fmt.Sprint(MaxImageWidth),
			DefaultValue: "800",
			DataType:     t.Integer(),
			Help:         "sets the maximum width of an image",
		},
		{
			Name:         "Max Upload File Size",
			Value:        fmt.Sprint(MaxUploadFileSize),
			DefaultValue: "26214400",
			DataType:     t.Integer(),
			Help:         "is the maximum upload file size in bytes. 1MB = 1024 * 1024",
		},
		{
			Name:         "Port",
			Value:        fmt.Sprint(Port),
			DefaultValue: "8080",
			DataType:     t.Integer(),
			Help:         "is the port used for http or https server",
		},
		{
			Name:         "Email From",
			Value:        EmailFrom,
			DefaultValue: "",
			DataType:     t.String(),
			Help:         "identifies where the email is coming from",
		},
		{
			Name:         "Email Username",
			Value:        EmailUsername,
			DefaultValue: "",
			DataType:     t.String(),
			Help:         "sets the username of an email",
		},
		{
			Name:         "Email Password",
			Value:        EmailPassword,
			DefaultValue: "",
			DataType:     t.String(),
			Help:         "sets the password of an email",
		},
		{
			Name:         "Email SMTP Server",
			Value:        EmailSMTPServer,
			DefaultValue: "",
			DataType:     t.String(),
			Help:         "sets the name of the SMTP Server in an email",
		},
		{
			Name:         "Email SMTP Server Port",
			Value:        fmt.Sprint(EmailSMTPServerPort),
			DefaultValue: "0",
			DataType:     t.Integer(),
			Help:         "sets the port number of an SMTP Server in an email",
		},
		{
			Name:         "Root URL",
			Value:        RootURL,
			DefaultValue: "/",
			DataType:     t.String(),
			Help:         "is where the listener is mapped to",
		},
		{
			Name:         "OTP Algorithm",
			Value:        OTPAlgorithm,
			DefaultValue: "sha1",
			DataType:     t.String(),
			Help:         "is the hashing algorithm of OTP. Other options are sha256 and sha512",
		},
		{
			Name:         "OTP Digits",
			Value:        fmt.Sprint(OTPDigits),
			DefaultValue: "6",
			DataType:     t.Integer(),
			Help:         "is the number of digits for the OTP",
		},
		{
			Name:         "OTP Period",
			Value:        fmt.Sprint(OTPPeriod),
			DefaultValue: "30",
			DataType:     t.Integer(),
			Help:         "the number of seconds for the OTP to change",
		},
		{
			Name:         "OTP Skew",
			Value:        fmt.Sprint(OTPSkew),
			DefaultValue: "5",
			DataType:     t.Integer(),
			Help:         "is the number of minutes to search around the OTP",
		},
		{
			Name: "Public Media",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(PublicMedia),
			DefaultValue: "0",
			DataType:     t.Boolean(),
			Help:         "allows public access to media handler without authentication",
		},
		{
			Name: "Log Delete",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(LogDelete),
			DefaultValue: "1",
			DataType:     t.Boolean(),
			Help:         "adds a log when a record is deleted",
		},
		{
			Name: "Log Add",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(LogAdd),
			DefaultValue: "1",
			DataType:     t.Boolean(),
			Help:         "adds a log when a record is added",
		},
		{
			Name: "Log Edit",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(LogEdit),
			DefaultValue: "1",
			DataType:     t.Boolean(),
			Help:         "adds a log when a record is edited",
		},
		{
			Name: "Log Read",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(LogRead),
			DefaultValue: "0",
			DataType:     t.Boolean(),
			Help:         "adds a log when a record is read",
		},
		{
			Name: "Cache Translation",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(CacheTranslation),
			DefaultValue: "0",
			DataType:     t.Boolean(),
			Help:         "allows a translation to store data in a cache memory",
		},
		{
			Name:         "Allowed IPs",
			Value:        AllowedIPs,
			DefaultValue: "*",
			DataType:     t.String(),
			Help: `is a list of allowed IPs to access uAdmin interfrace in one of the following formats:
										- * = Allow all
										- "" = Allow none
							 			- "192.168.1.1" Only allow this IP
										- "192.168.1.0/24" Allow all IPs from 192.168.1.1 to 192.168.1.254
											You can also create a list of the above formats using comma to separate them.
											For example: "192.168.1.1,192.168.1.2,192.168.0.0/24`,
		},
		{
			Name:         "Blocked IPs",
			Value:        BlockedIPs,
			DefaultValue: "",
			DataType:     t.String(),
			Help: `is a list of blocked IPs from accessing uAdmin interfrace in one of the following formats:
										 - "*" = Block all
										 - "" = Block none
										 - "192.168.1.1" Only block this IP
										 - "192.168.1.0/24" Block all IPs from 192.168.1.1 to 192.168.1.254
										 		You can also create a list of the above formats using comma to separate them.
												For example: "192.168.1.1,192.168.1.2,192.168.0.0/24`,
		},
		{
			Name: "Restrict Session IP",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(RestrictSessionIP),
			DefaultValue: "0",
			DataType:     t.Boolean(),
			Help:         "is to block access of a user if their IP changes from their original IP during login",
		},
		{
			Name: "Retain Media Versions",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(RetainMediaVersions),
			DefaultValue: "1",
			DataType:     t.Boolean(),
			Help:         "is to allow the system to keep files uploaded even after they are changed. This allows the system to \"Roll Back\" to an older version of the file",
		},
		{
			Name:         "Rate Limit",
			Value:        fmt.Sprint(RateLimit),
			DefaultValue: "3",
			DataType:     t.Integer(),
			Help:         "is the maximum number of requests/second for any unique IP",
		},
		{
			Name:         "Rate Limit Burst",
			Value:        fmt.Sprint(RateLimitBurst),
			DefaultValue: "3",
			DataType:     t.Integer(),
			Help:         "is the maximum number of requests for an idle user",
		},
		{
			Name: "Optimize SQL Query",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(OptimizeSQLQuery),
			DefaultValue: "1",
			DataType:     t.Boolean(),
			Help:         "OptimizeSQLQuery selects columns during rendering a form a list to visible fields.",
		},
		{
			Name: "API Log Read",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(APILogRead),
			DefaultValue: "0",
			DataType:     t.Boolean(),
			Help:         "APILogRead controls the data API's logging for read commands.",
		},
		{
			Name: "API Log Edit",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(APILogEdit),
			DefaultValue: "1",
			DataType:     t.Boolean(),
			Help:         "APILogEdit controls the data API's logging for edit commands.",
		},
		{
			Name: "API Log Add",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(APILogAdd),
			DefaultValue: "1",
			DataType:     t.Boolean(),
			Help:         "APILogAdd controls the data API's logging for add commands.",
		},
		{
			Name: "API Log Delete",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(APILogDelete),
			DefaultValue: "1",
			DataType:     t.Boolean(),
			Help:         "APILogDelete controls the data API's logging for delete commands.",
		},
		{
			Name: "API Log Schema",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(APILogSchema),
			DefaultValue: "1",
			DataType:     t.Boolean(),
			Help:         "APILogSchema controls the data API's logging for schema commands.",
		},
		{
			Name: "Log HTTP Requests",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(LogHTTPRequests),
			DefaultValue: "1",
			DataType:     t.Boolean(),
			Help:         "Logs http requests to syslog",
		},
		{
			Name:         "HTTP Log Format",
			Value:        HTTPLogFormat,
			DefaultValue: "",
			DataType:     t.String(),
			Help: `Is the format used to log HTTP access
									%a: Client IP address
									%{remote}p: Client port
									%A: Server hostname/IP
									%{local}p: Server port
									%U: Path
									%c: All coockies
									%{NAME}c: Cookie named 'NAME'
									%{GET}f: GET request parameters
									%{POST}f: POST request parameters
									%B: Response length
									%>s: Response code
									%D: Time taken in microseconds
									%T: Time taken in seconds
									%I: Request length`,
		},
		{
			Name: "Log Trail",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(LogTrail),
			DefaultValue: "0",
			DataType:     t.Boolean(),
			Help:         "Stores Trail logs to syslog",
		},
		{
			Name:         "Trail Logging Level",
			Value:        fmt.Sprint(TrailLoggingLevel),
			DefaultValue: "2",
			DataType:     t.Integer(),
			Help:         "Is the minimum level to be logged into syslog.",
		},
		{
			Name: "System Metrics",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(SystemMetrics),
			DefaultValue: "0",
			DataType:     t.Boolean(),
			Help:         "Enables uAdmin system metrics to be recorded",
		},
		{
			Name: "User Metrics",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(UserMetrics),
			DefaultValue: "0",
			DataType:     t.Boolean(),
			Help:         "Enables the user metrics to be recorded",
		},
		{
			Name: "Cache Sessions",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(CacheSessions),
			DefaultValue: "1",
			DataType:     t.Boolean(),
			Help:         "Allows uAdmin to store sessions data in memory",
		},
		{
			Name: "Cache Permissions",
			Value: func(v bool) string {
				n := 0
				if v {
					n = 1
				}
				return fmt.Sprint(n)
			}(CachePermissions),
			DefaultValue: "1",
			DataType:     t.Boolean(),
			Help:         "Allows uAdmin to store permissions data in memory",
		},
	}

	// Prepare uAdmin Settings
	for i := range settings {
		settings[i].CategoryID = cat.ID
		settings[i].Code = "uAdmin." + strings.Replace(settings[i].Name, " ", "", -1)
	}

	// Check if the settings exist in the DB
	var s Setting
	sList := []Setting{}
	Filter(&sList, "category_id = ?", cat.ID)
	tx := db.Begin()
	for i, setting := range settings {
		Trail(WORKING, "Synching System Settings: [%s%d/%d%s]", colors.FGGreenB, i+1, len(settings), colors.FGNormal)
		s = Setting{}
		for c := range sList {
			if sList[c].Code == setting.Code {
				s = sList[c]
			}
		}
		if s.ID == 0 {
			tx.Create(&setting)
			//setting.Save()
		} else {
			if s.DefaultValue != setting.DefaultValue || s.Help != setting.Help {
				if s.Help != setting.Help {
					s.Help = setting.Help
				}
				if s.Value == s.DefaultValue {
					s.Value = setting.DefaultValue
				}
				s.DefaultValue = setting.DefaultValue
				tx.Save(s)
				//s.Save()
			}
		}
	}
	tx.Commit()
	Trail(OK, "Synching System Settings: [%s%d/%d%s]", colors.FGGreenB, len(settings), len(settings), colors.FGNormal)
	applySystemSettings()
	settingsSynched = true
}

func applySystemSettings() {
	cat := SettingCategory{}
	settings := []Setting{}

	Get(&cat, "name = ?", "uAdmin")
	Filter(&settings, "category_id = ?", cat.ID)

	for _, setting := range settings {
		setting.ApplyValue()
	}
}
