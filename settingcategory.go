package uadmin

type SettingCategory struct {
	Model
	Name string
	Icon string `uadmin:"image"`
}
