package uadmin

// FieldType for field type in builder field
type FieldType int

// String for string type
func (FieldType) String() FieldType {
	return 1
}

// Int for int type
func (FieldType) Int() FieldType {
	return 2
}

// Float for float type
func (FieldType) Float() FieldType {
	return 3
}

// Multilingual for multiminual type
func (FieldType) Multilingual() FieldType {
	return 4
}

// Email for email type
func (FieldType) Email() FieldType {
	return 5
}

// Boolean for boolean type
func (FieldType) Boolean() FieldType {
	return 6
}

// DateTime for time.Time type
func (FieldType) DateTime() FieldType {
	return 7
}

// DateTimePtr for time.Time pointer type
func (FieldType) DateTimePtr() FieldType {
	return 8
}

// ForeignKey for fk type
func (FieldType) ForeignKey() FieldType {
	return 9
}

// M2M for m2m type
func (FieldType) M2M() FieldType {
	return 10
}

// ProgressBar for progressbar type
func (FieldType) ProgressBar() FieldType {
	return 11
}

// HTML for html type
func (FieldType) HTML() FieldType {
	return 12
}

// StaticList for list type
func (FieldType) StaticList() FieldType {
	return 13
}

// File for file type
func (FieldType) File() FieldType {
	return 14
}

// Image for image type
func (FieldType) Image() FieldType {
	return 15
}

// Money for money type
func (FieldType) Money() FieldType {
	return 16
}

// Code for code type
func (FieldType) Code() FieldType {
	return 17
}

// Link for link type
func (FieldType) Link() FieldType {
	return 18
}

// Password for password type
func (FieldType) Password() FieldType {
	return 19
}
