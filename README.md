admin Tags

	read_only:"true"
	email:"true"
	hidden:"true"
	html:"true"
	fk:"ModelName"
	list:"true"
	list_filter:"true"
	search:"true"
	dontCache:"true"
  	required:"true"
  	help:"true"
  	pattern:"true"
  	pattern_msg:"Message"
	max:"int"
	min:"int"
	link:"true"
	file:"true"
	dependsOn:""
	linkerObj:""
	linkerParentField:""
	linkerChildField:""
	childObj:""
	listExclude:"true"
	categorical_filter:"true"
	password:"true"
	image:"true"
	
	
	Registering Models:
	eadmin.Register("/link/",&Models{})
	
	
	Adding Inlines :
	orderFK := map[string]string{}
	orderFK["grouppermission"] = "user_group_id"
	RegisterInlines("usergroup", orderFK, &GroupPermission{})
	
	
	Using Many to Many Field:
	`gorm:"many2many:bill_subscriptions"`
	
	
	
	overriding Save Function
	
	func (m* Model)Save(){
	    //business logic
	    eadmin.Save(m)
	}
	
	func (v Validate) Validate() (ret map[string]string) {
	    ret = map[string]string{}
	    if v.Name != "test" {
		    ret["Name"] = "Error name not found"
	    }
	    return
    }