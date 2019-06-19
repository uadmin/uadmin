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

## [0.1.1] -  2019-06-20

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
