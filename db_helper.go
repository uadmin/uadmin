package uadmin

import "strings"

func fixQueryEnclosure(v string) string {
	if Database.Type == "postgres" {
		return strings.ReplaceAll(v, "`", "\"")
	}
	return v
}

func trimEnclosure(v string) string {
	if Database.Type == "postgres" {
		return strings.ReplaceAll(v, "\"", "")
	}
	return strings.ReplaceAll(v, "`", "")
}

func columnEnclosure() string {
	if Database.Type == "postgres" {
		return "\""
	}
	return "`"
}

func getLike(caseSensitive bool) string {
	if Database.Type == "postgres" {
		if caseSensitive {
			return "LIKE"
		}
		return "ILIKE"
	}
	if Database.Type == "mysql" {
		if caseSensitive {
			return "LIKE BINARY"
		}
		return "LIKE"
	}
	if Database.Type == "sqlite" {
		return "LIKE"
	}
	return "LIKE"
}
