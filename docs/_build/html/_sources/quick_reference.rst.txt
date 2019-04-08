Quick Reference
===============

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
    