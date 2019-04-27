package models

// Format is the name of the drop down list ...
type Format int

// PDF is the name of the drop down list value ...
func (Format) PDF() Format {
	return 1
}

// TXT is the name of the drop down list value ...
func (Format) TXT() Format {
	return 2
}

// Others is the name of the drop down list value ...
func (Format) Others() Format {
	return 3
}