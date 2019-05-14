Quick Reference
===============

Overriding Save Function
------------------------

.. code-block:: go

    func (m *Model) Save() {
        // business logic
        uadmin.Save(m)
    }

String Function for Drop Down List
----------------------------------

.. code-block:: go

    func (m *Model) String() string {
        var name string
        _obj := *(m)
        t := reflect.TypeOf(_obj)
        for i := 0; i < t.NumMethod(); i++ {
            m := t.Method(i).Name
            if m != "String" {
                tmp := reflect.ValueOf(_obj).MethodByName(m).Call([]reflect.Value{})
                if fmt.Sprint(tmp[0]) == fmt.Sprint(_obj) {
                    name = m
                    break
                }
            }
        }
        return name
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
    