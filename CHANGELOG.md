# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.1] -  2019-05-11

### Added
	-	Global Variable:
		- DefaultMediaPermission [0644]: The default permission for files uploaded through uAdmin
### Changed
### Deprecated
### Removed
### Fixed
	- Delete(): Checks for ID == 0 where it doesn't delete all records when ID = 0 is passed
### Security

## [0.2.0] -  2019-06-20

### Added
	- Global Variable:
		- Settings model which can be used system wide to keep application settings.
		- GetSetting function to read settings from the settings model.
		- User.GetAccess function can evaluate the user's permission to a model using user level and group level and return a UserPermission instance.
		- ErrorHandleFunc: A function that can be passed for handling errors system wide that gets called when Trail is called.
		- AllowedIPs: is a list of allowed IPs to access uAdmin interfrace
		- BlockedIPs: is a list of blocked IPs from accessing uAdmin interfrace
		- RestrictSessionIP: restricts access to the system if the IP of the user changes after login
		- RetainMediaVersions: is to allow the system to keep files uploaded even after they are changed
### Changed
	- Forgot password now sends HTML emails.
	- GetString function can return the name of a static list item.
### Deprecated
### Removed
### Fixed
  - Bug fix for read only fields for new and edit.
	- Model names in dashboard are generated using proper plural function instead of just adding "s"
### Security

## [0.3.0] -  2019-07-02

### Added
	- Approval System: By adding `approval` tag to any field, the field will required a special permission to edit. If the user does not have this permission, the edit will generate an Approval record that can be approved by a user with access to the new Approval model.
	- Upload files using drag and drop into the field.
	- Image and File fields now allow a new tag called `webcam` which adds web can access directly from the field.
	- Model specific themes: You can use `ModelSchema.FormTheme` and `ModelSchema.ListTheme` to choose a theme for a model.
	- Settings automatically adds all uAdmin global variables as settings.
	- Added .gitignore which should have been there all along!!
### Changed
	- cropImageHandler now returns JSON with a status letting the for success and failure.
	- Improved performance for saving settings page.
### Deprecated
### Removed
### Fixed
	- Fixed filters for list view. Now you can apply multiple filters. Also now you can filter by foreign key.
	- Removed all `fmt.Println` and `log.Println` for printing errors and replaced it with `Trail`.
	- Export to excel date time type now takes time zone into account.
	- Fixed bug with method fields names rendering with missing letters.
	- Settings URL takes into account `RootURL`
	- Edit logs will only be saved if the form didn't have errors.
	- Fixed bug with `required` file and image fields where you had to choose a file everytime you save even if there was a files saved in the record.
	- Fixed bug with `required` foreign key and list type where it was not working before.
### Security
	- revertLogHandler required the requires to be authenticated and the user to have access to edit the model and have read access to logs.

## [0.4.0] -  2019-07-02

### Added
  - Implemented request rate limits to protect from DDoS
  - Implemented AB/Testing system
  - Implemented dAPI which is an API to access model data
  - Implemented CacheSessions and CachePermissions for direct in memory access to sessions and permissions
  - Implemented a Metrics system using the following function: `SetMetric`, `IncrementMetric`, `TimeMetric` and `NewMetric`
  - Trail can log to syslog
  - HTTP requests can be logged to syslog
  - Added `uadmin.Handler(func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request)` to enable syslog for HTTP requests
  - Add OptimizeSQLQuery mode
  - Handle ctrl+F in Home and List view to focus on the search field instead of the native search
  - Added `stringer` meta tag
  - Added two API end points `/api/get_models` and `/api/get_fields?m={MODEL_NAME}`
### Changed
  - Changed `Login(*http.Request, string, string) (*User, bool)` to `Login(*http.Request, string, string) (*Session, bool)`
  - Changed `Login2FA(*http.Request, string, string, string) *User` to `Login2FA(*http.Request, string, string, string) *Session`
  - Changed `HTMLContext` to `RenderHTML` that has support for templates functions
  - URL filter place holders are capital letters `{USERNAME}`, `{USERID}` and `{NOW}`
  - Search is allowed for `list_exclude` fields.
  - Changed the search API path to `/api/search/`
  - Make static handler a public function `uadmin.StaticHandler(http.ResponseWriter, *http.Request)`
  - Added three new level to `Trail` for compatibility with syslog which are `Critical`, `Alert` and `Emergency`
  - If database doesn't exist for mysql, uAdmin will try to create a new database.
### Deprecated
  - `User.HasPermission` will be private starting `0.6.0`
  - `UserGroup.HasPermission` will be private starting `0.6.0`
### Removed
### Fixed
  - Fixed image crop modal conflict in list view with delete modal and add it to form and inlines
  - Fixed FK in approvals
  - Remove required from fields with pending approval
  - Support filtering/searching by NULL value for `time.Time` pointer
  - Fixed filtering by FK
### Security
  - Restrict access to inlines based on user model permissions
  - Search API escapes HTML results
