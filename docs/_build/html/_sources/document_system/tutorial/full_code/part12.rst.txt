Document System Tutorial Part 12 - Custom Count function (Full Source Code)
===========================================================================

.. code-block:: go

	// Count !
	func (d Document) Count(a interface{}, query interface{}, args ...interface{}) int {
		// Converts the query into a string
		Q := fmt.Sprint(query)

		// Checks whether the string contains a query and a UserID
		if strings.Contains(Q, "user_id = ?") {
			// Split the query part by part
			qParts := strings.Split(Q, " AND ")

			// Initialize tempArgs as an interface and tempQuery as a
			// string
			tempArgs := []interface{}{}
			tempQuery := []string{}

			// Loop the query every part
			for i := range qParts {
				// Checks whether the specific query part is not
				// equal to the UserID value
				if qParts[i] != "user_id = ?" {
					// Append the arguments into the tempArgs
					// variable
					tempArgs = append(tempArgs, args[i])

					// Append the specific query part into the
					// tempQuery variable
					tempQuery = append(tempQuery, qParts[i])
				}
			}
			// Concatenate the query to create a single string
			query = strings.Join(tempQuery, " AND ")

			// Assign tempArgs object into the args variable
			args = tempArgs
		}

		// Return the a, query, and args... inside the Count function
		// parameters
		return uadmin.Count(a, query, args...)
	}
