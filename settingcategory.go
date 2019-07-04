package uadmin

// SettingCategory is a category for system settings
type SettingCategory struct {
	Model
	Name string
	Icon string `uadmin:"image"`
}
