Quick Reference
===============

Many-to-many relationship (m2m)
-------------------------------

.. code-block:: go

    type Model struct {
        uadmin.Model
        Field     []Field   `uadmin:"list_exclude"`
        FieldList string    `uadmin:"read_only"`
    }

    // FieldSave ...
    func (m *Model) FieldSave() {
        // Initializes the fieldList as empty string
        fieldList := ""

        // This process will get the name of the Field, store into the
        // fieldList and if the index value is not equal to the number of 
        // Field, it will insert the comma symbol at the end of the word.
        for x, key := range m.Field {
            fieldList += key.Name
            if x != len(m.Field)-1 {
                fieldList += ", "
            }
        }

        // Store the catList variable to the FieldList field in the model.
        m.FieldList = fieldList

        // Override save
        uadmin.Save(m)
    }

    // Save ...
    func (i *Item) Save() {
        if i.ID == 0 {
            i.FieldSave()
        }

        i.FieldSave()
    }

Overriding Save Function
------------------------

.. code-block:: go

    func (m *Model)Save(){
        //business logic
        uadmin.Save(m)
    }

Validation
----------

.. code-block:: go


    func (v Validate) Validate() (ret map[string]string) {
        ret = map[string]string{}
        if v.Name != "test" {
            ret["Name"] = "Error name not found"
        }
        return
    }