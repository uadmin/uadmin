package uadmin

type FieldType int

func (FieldType) String() FieldType {
	return 1
}

func (FieldType) Int() FieldType {
	return 2
}

func (FieldType) Float() FieldType {
	return 3
}

func (FieldType) Multilingual() FieldType {
	return 4
}

func (FieldType) Email() FieldType {
	return 5
}

func (FieldType) Boolean() FieldType {
	return 6
}

func (FieldType) DateTime() FieldType {
	return 7
}

func (FieldType) DateTimePtr() FieldType {
	return 8
}

func (FieldType) ForeignKey() FieldType {
	return 9
}

func (FieldType) M2M() FieldType {
	return 10
}

func (FieldType) ProgressBar() FieldType {
	return 11
}

func (FieldType) HTML() FieldType {
	return 12
}

func (FieldType) StaticList() FieldType {
	return 13
}

func (FieldType) File() FieldType {
	return 14
}

func (FieldType) Image() FieldType {
	return 15
}

func (FieldType) Money() FieldType {
	return 16
}

func (FieldType) Code() FieldType {
	return 17
}

func (FieldType) Link() FieldType {
	return 18
}

func (FieldType) Password() FieldType {
	return 19
}
