package uadmin

import "strings"

func fixQueryEnclosure(v string) string {
	if Database.Type == "postgres" {
		return strings.ReplaceAll(v, "`", "\"")
	}
	return v
}

func columnEnclosure() string {
	if Database.Type == "postgres" {
		return "\""
	}
	return "`"
}
