package uadmin

import (
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

type UAdminTests struct{ *testing.T }

func TestRunner(t *testing.T) {
	initialSetup()
	teardownFunction()

	databaseSetup := []struct {
		Name string
		DB   *DBSettings
	}{
		{
			Name: "SQLite",
			DB:   nil,
		},
		{
			Name: "MySQL",
			DB: &DBSettings{
				Type: "mysql",
				Name: "uadmintestdb",
				User: func() string {
					if v := os.Getenv("UADMIN_TEST_MYSQL_USERNAME"); v != "" {
						return v
					}
					return "root"
				}(),
				Password: func() string {
					if v := os.Getenv("UADMIN_TEST_MYSQL_PASSWORD"); v != "" {
						return v
					}
					return ""
				}(),
				Host: func() string {
					if v := os.Getenv("UADMIN_TEST_MYSQL_HOST"); v != "" {
						return v
					}
					return "127.0.0.1"
				}(),
				Port: func() int {
					if v := os.Getenv("UADMIN_TEST_MYSQL_PORT"); v != "" {
						port, err := strconv.Atoi(v)
						if err == nil {
							return port
						}
					}
					return 3306
				}(),
			},
		},
		{
			Name: "Postgres",
			DB: &DBSettings{
				Type: "postgres",
				Name: "uadmintestdb",
				User: func() string {
					if v := os.Getenv("UADMIN_TEST_POSTGRES_USERNAME"); v != "" {
						return v
					}
					return "root"
				}(),
				Password: func() string {
					if v := os.Getenv("UADMIN_TEST_POSTGRES_PASSWORD"); v != "" {
						return v
					}
					return ""
				}(),
				Host: func() string {
					if v := os.Getenv("UADMIN_TEST_POSTGRES_HOST"); v != "" {
						return v
					}
					return "127.0.0.1"
				}(),
				Port: func() int {
					if v := os.Getenv("UADMIN_TEST_POSTGRES_PORT"); v != "" {
						port, err := strconv.Atoi(v)
						if err == nil {
							return port
						}
					}
					return 3306
				}(),
			},
		},
	}
	for _, dbSetup := range databaseSetup {
		Database = dbSetup.DB
		setupFunction()
		uTest := UAdminTests{t}
		t.Run(dbSetup.Name+"=404", func(t *testing.T) {
			uTest.TestPage404Handler()
		})
		t.Run(dbSetup.Name+"=ABTest", func(t *testing.T) {
			uTest.TestABTest()
			uTest.TestLoadModels()
			uTest.TestLoadFields()
		})
		t.Run(dbSetup.Name+"=Admin", func(t *testing.T) {
			uTest.TestIsLocal()
			uTest.TestCommaf()
			uTest.TestPaginationHandler()
			uTest.TestToSnakeCase()
			uTest.TestJSONMarshal()
			uTest.TestReturnJSON()
		})
		t.Run(dbSetup.Name+"=APIHandler", func(t *testing.T) {
			uTest.TestAPIHandler()
		})
		t.Run(dbSetup.Name+"=Approval", func(t *testing.T) {
			uTest.TestApprovalStruct()
		})
		t.Run(dbSetup.Name+"=Auth", func(t *testing.T) {
			uTest.TestGenerateBase64()
			uTest.TestGenerateBase32()
			uTest.TestHashPass()
			uTest.TestIsAuthenticated()
			uTest.TestGetUserFromRequest()
			uTest.TestLogin()
			uTest.TestLogin2FA()
			uTest.TestLogout()
			uTest.TestValidateIP()
			uTest.TestGetSessionByKey()
			uTest.TestGetSession()
		})
		t.Run(dbSetup.Name+"=Crop", func(t *testing.T) {
			uTest.TestCropImageHandler()
		})
		t.Run(dbSetup.Name+"=DAPI", func(t *testing.T) {
			uTest.TestDAPI()
		})
		t.Run(dbSetup.Name+"=DashboardMenu", func(t *testing.T) {
			uTest.TestDashboardMenu()
		})
		t.Run(dbSetup.Name+"=DB", func(t *testing.T) {
			uTest.TestInitializeDB()
			uTest.TestSave()
		})
		t.Run(dbSetup.Name+"=DeleteHandler", func(t *testing.T) {
			uTest.TestProcessDelete()
		})
		t.Run(dbSetup.Name+"=Encrypt", func(t *testing.T) {
			uTest.TestGenerateByteArray()
			uTest.TestEncrypt()
			uTest.TestEncryptRecord()
			uTest.TestEncryptArray()
		})
		t.Run(dbSetup.Name+"=Export", func(t *testing.T) {
			uTest.TestGetFilter()
		})
		t.Run(dbSetup.Name+"=FieldType", func(t *testing.T) {
			uTest.TestFieldType()
		})
		t.Run(dbSetup.Name+"=ForgotPassword", func(t *testing.T) {
			uTest.TestForgotPasswordHandler()
		})
		t.Run(dbSetup.Name+"=FormHandler", func(t *testing.T) {
			uTest.TestFormHandler()
		})
		t.Run(dbSetup.Name+"=GenerateTranslation", func(t *testing.T) {
			uTest.TestSyncCustomTranslation()
			uTest.TestSyncModelTranslation()
		})
		t.Run(dbSetup.Name+"=GetSchema", func(t *testing.T) {
			uTest.TestGetSchema()
		})
		t.Run(dbSetup.Name+"=GroupPermissions", func(t *testing.T) {
			uTest.TestGroupPermission()
		})
		t.Run(dbSetup.Name+"=HomeHamdler", func(t *testing.T) {
			uTest.TestHomeHandler()
		})
		t.Run(dbSetup.Name+"=Language", func(t *testing.T) {
			uTest.TestLanguage()
		})
		t.Run(dbSetup.Name+"=ListHandler", func(t *testing.T) {
			uTest.TestListHandler()
		})
		t.Run(dbSetup.Name+"=LoginHandler", func(t *testing.T) {
			uTest.TestLoginHandler()
		})
		t.Run(dbSetup.Name+"=MainHandler", func(t *testing.T) {
			uTest.TestMainHandler()
		})
		t.Run(dbSetup.Name+"=ProfileHandler", func(t *testing.T) {
			uTest.TestProfileHandler()
		})
		t.Run(dbSetup.Name+"=RevertLogHandler", func(t *testing.T) {
			uTest.TestRevertLogHandler()
		})
		t.Run(dbSetup.Name+"=SendEmail", func(t *testing.T) {
			uTest.TestSendEmail()
		})
		t.Run(dbSetup.Name+"=SettingsHandler", func(t *testing.T) {
			uTest.TestSettingsHandler()
		})

		teardownFunction()
	}

}

type TestModelA struct {
	Model
	Name string
}

type TestModelB struct {
	Model
	Name         string     `uadmin:"help:This is a test help message;search;list_exclude"`
	ItemCount    int        `uadmin:"max:5;min:1;format:%03d;required;read_only:true,edit"`
	Phone        string     `uadmin:"default_value:09;pattern:[0-9+]{7,15};pattern_msg:invalid phone number;encrypt"`
	Active       bool       `uadmin:"hidden;read_only"`
	OtherModel   TestModelA `uadmin:"categorical_filter;filter;read_only:new"`
	OtherModelID uint
	ModelAList   []TestModelA `gorm:"-"`
	Parent       *TestModelB
	ParentID     uint
	Email        string  `uadmin:"email"`
	Greeting     string  `uadmin:"multilingual"`
	Image        string  `uadmin:"image;upload_to:/media/home/me/images/"`
	File         string  `uadmin:"file;upload_to:/media/home/me/files"`
	Secret       string  `uadmin:"password"`
	Description  string  `uadmin:"html"`
	URL          string  `uadmin:"link"`
	Code         string  `uadmin:"code"`
	P1           int     `uadmin:"progress_bar"`
	P2           float64 `uadmin:"progress_bar"`
	P3           float64 `uadmin:"progress_bar:1.0"`
	P4           float64 `uadmin:"progress_bar:1.0:red"`
	P5           float64 `uadmin:"progress_bar:1.0:#f00"`
	P6           float64 `uadmin:"progress_bar:0.3:red,0.7:yellow,1.0:lime"`
	Price        float64 `uadmin:"money"`
	List         testList
}

type TestApproval struct {
	Model
	Name        string     `uadmin:"approval"`
	Start       time.Time  `uadmin:"approval"`
	End         *time.Time `uadmin:"approval"`
	Count       int        `uadmin:"approval"`
	Price       float64    `uadmin:"approval"`
	List        testList   `uadmin:"approval"`
	TestModel   TestModelA `uadmin:"approval"`
	TestModelID uint
	Active      bool `uadmin:"approval"`
}

// Method__List__Form is a method to test method based properties for models
func (TestModelB) Method__List__Form() string {
	return "Value"
}

type testList int

func (testList) A() testList {
	return 1
}

func initialSetup() {
	Port = 5000
	EmailFrom = "uadmin@example.com"
	EmailPassword = "password"
	EmailUsername = "uadmin@example.com"
	EmailSMTPServer = "localhost"
	EmailSMTPServerPort = 2525

	RateLimit = 1000000
	RateLimitBurst = 1000000
	go startEmailServer()

	PasswordAttempts = 1000000
	if !strings.Contains(AllowedHosts, "example.com") {
		AllowedHosts += ",example.com"
	}

	ErrorHandleFunc = func(level int, err string, stack string) {
		if level >= ERROR {
			Trail(DEBUG, stack)
		}
	}
}

func setupFunction() {
	Register(
		TestStruct1{},
		TestModelA{},
		TestModelB{},
		TestApproval{},
	)

	schema := Schema["testmodelb"]
	schema.ListTheme = "default"
	schema.FormTheme = "default"
	Schema["testmodelb"] = schema

	RegisterInlines(TestModelA{}, map[string]string{"TestModelB": "OtherModelID"})

	go StartServer()
	//time.Sleep(time.Second * 10)
	for !dbOK {
		time.Sleep(time.Millisecond * 100)
	}

}

func teardownFunction() {
	// Remove Generated Files
	os.Remove("uadmin.db")
	os.Remove(".key")
	os.Remove(".salt")
	os.Remove(".uproj")
	os.Remove(".bindip")

	// Delete temp media file
	os.RemoveAll("./media")
	os.RemoveAll("./static/i18n")

	// Delete DB state variables
	dbOK = false
	models = map[string]interface{}{}
	ClearDB()
	ServerReady = false
	SiteName = "uAdmin"
	settingsSynched = false
	registered = false
}

// func TestMain(t *testing.M) {
// 	initialSetup()

// 	teardownFunction()
// 	setupFunction()
// 	retCode := t.Run()
// 	teardownFunction()

// 	// test MySQL

// 	Database = &DBSettings{
// 		Type: "mysql",
// 		Name: "uadmintestdb",
// 		User: func() string {
// 			if v := os.Getenv("UADMIN_TEST_MYSQL_USERNAME"); v != "" {
// 				return v
// 			}
// 			return "root"
// 		}(),
// 		Password: func() string {
// 			if v := os.Getenv("UADMIN_TEST_MYSQL_PASSWORD"); v != "" {
// 				return v
// 			}
// 			return ""
// 		}(),
// 		Host: func() string {
// 			if v := os.Getenv("UADMIN_TEST_MYSQL_HOST"); v != "" {
// 				return v
// 			}
// 			return "127.0.0.1"
// 		}(),
// 		Port: func() int {
// 			if v := os.Getenv("UADMIN_TEST_MYSQL_PORT"); v != "" {
// 				port, err := strconv.Atoi(v)
// 				if err == nil {
// 					return port
// 				}
// 			}
// 			return 3306
// 		}(),
// 	}
// 	setupFunction()
// 	retCode += t.Run()
// 	teardownFunction()

// 	os.Exit(retCode)
// }
