package uadmin

import (
	"database/sql"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

func parseCustomDBSchema(rows *sql.Rows) interface{} {
	m := []map[string]interface{}{}

	// Parse the data into an array
	columns, _ := rows.Columns()

	//var current interface{}
	for rows.Next() {
		vals := makeResultReceiver(len(columns))
		rows.Scan(vals...)
		row := map[string]interface{}{}
		for i := range columns {
			row[columns[i]] = getDBValue(vals[i])
		}
		m = append(m, row)
	}
	return m
}

func getDBValue(p interface{}) interface{} {
	i := p.(*interface{})
	switch v := (*i).(type) {
	case int:
		return v
	case string:
		return v
	case uint:
		return v
	case []uint8:
		return string(v)
	default:
		return v
	}
}

func getURLArgs(r *http.Request) map[string]string {
	params := map[string]string{}

	// First Parse GET params
	getParams := strings.Split(r.URL.RawQuery, "&")
	for _, param := range getParams {
		paramParts := strings.SplitN(param, "=", 2)
		if len(paramParts) != 2 {
			continue
		}

		// Skip session
		if paramParts[0] == "session" {
			continue
		}

		params[paramParts[0]] = paramParts[1]
	}

	// Parse post parameters
	for k, v := range r.PostForm {
		if len(v) < 1 {
			continue
		}
		// Skip session
		if k == "session" {
			continue
		}
		params[k] = v[0]
	}

	return params
}

func makeResultReceiver(length int) []interface{} {
	result := make([]interface{}, 0, length)
	for i := 0; i < length; i++ {
		var current interface{}
		current = struct{}{}
		result = append(result, &current)
	}
	return result
}

func getFilters(params map[string]string, tableName string, schema *ModelSchema) (query string, args []interface{}) {
	qParts := []string{}
	args = []interface{}{}

	for k, v := range params {
		// TODO: Explain the condition with a URL to specs
		// Anything that starts with a '$' or '_' in assigning a field name
		// will be skipped. '$' is used as a special parameter such as limit,
		// offset, and order while '_' is used for writing data (Add/Edit).
		// This loop executes either by assigning an exact field name
		// or by using $or parameter as a special case. It compares which of
		// the assigned parameters matches a specified criteria.
		// http://example.com/api/d/modelname/read/?f=1
		// http://example.com/api/d/modelname/read/?$or=f=1|f=2
		if (k[0] == '$' && k != "$or" && k != "$q") || k[0] == '_' {
			continue
		}
		if k == "$or" {
			orQList := strings.Split(v, "|")
			orQParts := []string{}
			orArgs := []interface{}{}
			for _, orQ := range orQList {
				andQParts := []string{}
				andArgs := []interface{}{}

				andQList := strings.Split(orQ, "+")

				for _, andQ := range andQList {
					andParts := strings.SplitN(andQ, "=", 2)
					if len(andParts) != 2 {
						continue
					}
					andQParts = append(andQParts, getQueryOperator(andParts[0], tableName))
					andArgs = append(andArgs, getQueryArg(andParts[0], andParts[1])...)
				}
				orQParts = append(orQParts, "("+strings.Join(andQParts, " AND ")+")")
				orArgs = append(orArgs, andArgs...)
			}
			qParts = append(qParts, "("+strings.Join(orQParts, " OR ")+")")
			args = append(args, orArgs...)
		} else if k == "$q" {
			orQParts := []string{}
			orArgs := []interface{}{}
			for _, field := range schema.Fields {
				if field.Searchable {
					// TODO: Supprt non-string types
					if field.TypeName == "string" {
						orQParts = append(orQParts, getQueryOperator(field.ColumnName+"__icontains", tableName))
						orArgs = append(orArgs, getQueryArg(field.ColumnName+"__icontains", v)...)
					} else if field.Type == "number" {
						orQParts = append(orQParts, getQueryOperator(field.ColumnName+"__contains", tableName))
						orArgs = append(orArgs, getQueryArg(field.ColumnName+"__contains", v)...)
					}
				}
			}
			if len(orQParts) != 0 {
				qParts = append(qParts, "("+strings.Join(orQParts, " OR ")+")")
				args = append(args, orArgs...)
			}
		} else {
			qParts = append(qParts, getQueryOperator(k, tableName))
			args = append(args, getQueryArg(k, v)...)
		}
	}
	query = strings.Join(qParts, " AND ")

	if !strings.Contains(query, "deleted_at") {
		if val, ok := params["$deleted"]; ok {
			if val == "0" {
				qParts = append(qParts, tableName+".deleted_at IS NULL")
			}
		} else {
			qParts = append(qParts, tableName+".deleted_at IS NULL")
		}
		query = strings.Join(qParts, " AND ")
	}

	return query, args
}

func getQueryOperator(v string, tableName string) string {
	// Determine if the query is negated
	n := len(v) > 0 && v[0] == '!'
	nTerm := ""
	if n {
		nTerm = " NOT"
		v = v[1:]
	}

	// add table name
	if !strings.Contains(v, ".") {
		v = "`" + tableName + "`.`" + v
	} else {
		vParts := strings.SplitN(v, ".", 2)
		v = "`" + vParts[0] + "`.`" + vParts[1]
	}

	if strings.HasSuffix(v, "__gt") {
		if n {
			return strings.TrimSuffix(v, "__gt") + "` <= ?"
		} else {
			return strings.TrimSuffix(v, "__gt") + "` > ?"
		}
	}
	if strings.HasSuffix(v, "__gte") {
		if n {
			return strings.TrimSuffix(v, "__gte") + "` < ?"
		} else {
			return strings.TrimSuffix(v, "__gte") + "` >= ?"
		}
	}
	if strings.HasSuffix(v, "__lt") {
		if n {
			return strings.TrimSuffix(v, "__lt") + "` >= ?"
		} else {
			return strings.TrimSuffix(v, "__lt") + "` < ?"
		}
	}
	if strings.HasSuffix(v, "__lte") {
		if n {
			return strings.TrimSuffix(v, "__lte") + "` > ?"
		} else {
			return strings.TrimSuffix(v, "__lte") + "` <= ?"
		}
	}
	if strings.HasSuffix(v, "__in") {
		return strings.TrimSuffix(v, "__in") + nTerm + "` IN (?)"
	}
	if strings.HasSuffix(v, "__is") {
		return strings.TrimSuffix(v, "__is") + "` IS" + nTerm + " ?"
	}
	if strings.HasSuffix(v, "__contains") {
		if Database.Type == "mysql" {
			return strings.TrimSuffix(v, "__contains") + "`" + nTerm + " LIKE BINARY ?"
		} else if Database.Type == "sqlite" {
			return strings.TrimSuffix(v, "__contains") + "`" + nTerm + " LIKE ?"
		}
	}
	if strings.HasSuffix(v, "__between") {
		return strings.TrimSuffix(v, "__between") + "`" + nTerm + " BETWEEN ? AND ?"
	}
	if strings.HasSuffix(v, "__startswith") {
		if Database.Type == "mysql" {
			return strings.TrimSuffix(v, "__startswith") + "`" + nTerm + " LIKE BINARY ?"
		} else if Database.Type == "sqlite" {
			return strings.TrimSuffix(v, "__startswith") + "`" + nTerm + " LIKE ?"
		}
	}
	if strings.HasSuffix(v, "__endswith") {
		if Database.Type == "mysql" {
			return strings.TrimSuffix(v, "__endswith") + "`" + nTerm + " LIKE BINARY ?"
		} else if Database.Type == "sqlite" {
			return strings.TrimSuffix(v, "__endswith") + "`" + nTerm + " LIKE ?"
		}
	}
	if strings.HasSuffix(v, "__re") {
		return strings.TrimSuffix(v, "__re") + nTerm + " REGEXP ?"
	}
	if strings.HasSuffix(v, "__icontains") {
		return "UPPER(" + strings.TrimSuffix(v, "__icontains") + "`)" + nTerm + " LIKE UPPER(?)"
	}
	if strings.HasSuffix(v, "__istartswith") {
		return "UPPER(" + strings.TrimSuffix(v, "__istartswith") + "`)" + nTerm + " LIKE UPPER(?)"
	}
	if strings.HasSuffix(v, "__iendswith") {
		return "UPPER(" + strings.TrimSuffix(v, "__iendswith") + "`)" + nTerm + " LIKE UPPER(?)"
	}
	if n {
		return v + "` <> ?"
	}
	return v + "` = ?"
}

func getQueryArg(k, v string) []interface{} {
	var err error
	v, err = url.QueryUnescape(v)
	if err != nil {
		Trail(WARNING, "getQueryArg url.QueryUnescape unable to unescape value. %s", err)
		return []interface{}{v}
	}
	if strings.HasSuffix(k, "__in") {
		return []interface{}{strings.Split(v, ",")}
	}
	if strings.HasSuffix(k, "__is") {
		if strings.ToUpper(v) == "NULL" {
			return []interface{}{interface{}(nil)}
		} else {
			return []interface{}{v}
		}
	}
	if strings.HasSuffix(k, "__contains") {
		return []interface{}{"%" + v + "%"}
	}
	if strings.HasSuffix(k, "__between") {
		splittedValue := strings.Split(v, ",")
		from := splittedValue[0]
		to := splittedValue[1]
		return []interface{}{from, to}
	}
	if strings.HasSuffix(k, "__startswith") {
		return []interface{}{v + "%"}
	}
	if strings.HasSuffix(k, "__endswith") {
		return []interface{}{"%" + v}
	}
	if strings.HasSuffix(k, "__icontains") {
		return []interface{}{"%" + v + "%"}
	}
	if strings.HasSuffix(k, "__istartswith") {
		return []interface{}{v + "%"}
	}
	if strings.HasSuffix(k, "__iendswith") {
		return []interface{}{"%" + v}
	}
	return []interface{}{v}
}

func getQueryFields(params map[string]string, tableName string) (string, bool) {
	//customSchema := false

	fieldRaw, customSchema := params["$f"]
	if fieldRaw == "" {
		return "", customSchema
	}

	fieldParts := strings.Split(fieldRaw, ",")
	fieldArray := []string{}

	for _, field := range fieldParts {
		// Check for SQL injection
		if strings.Contains(field, " ") || strings.Contains(field, ";") {
			continue
		}

		// Check for aggregation function
		if strings.Contains(field, "__") {
			//customSchema = true
			fieldParts := strings.SplitN(field, "__", 2)

			//add table name
			if !strings.Contains(fieldParts[0], ".") {
				fieldParts[0] = "`" + tableName + "`.`" + fieldParts[0]
			} else {
				fieldNameParts := strings.Split(fieldParts[0], ".")
				fieldParts[0] = "`" + fieldNameParts[0] + "`.`" + fieldNameParts[1] + "`"
			}

			switch fieldParts[1] {
			case "sum":
				field = "SUM(" + fieldParts[0] + ") AS " + strings.Replace(field, ".", "__", -1)
			case "min":
				field = "MIN(" + fieldParts[0] + ") AS " + strings.Replace(field, ".", "__", -1)
			case "max":
				field = "MAX(" + fieldParts[0] + ") AS " + strings.Replace(field, ".", "__", -1)
			case "avg":
				field = "AVG(" + fieldParts[0] + ") AS " + strings.Replace(field, ".", "__", -1)
			case "count":
				field = "COUNT(" + fieldParts[0] + ") AS " + strings.Replace(field, ".", "__", -1)
			}
		} else {
			//add table name
			if !strings.Contains(field, ".") {
				field = "`" + tableName + "`.`" + field + "`"
			} else {
				fieldNameParts := strings.Split(field, ".")
				if fieldNameParts[0] != tableName {
					customSchema = true
				}
				field = "`" + fieldNameParts[0] + "`.`" + fieldNameParts[1] + "`" + " AS " + strings.Replace(field, ".", "__", -1)
			}
		}
		fieldArray = append(fieldArray, field)
	}

	return strings.Join(fieldArray, ", "), customSchema
}

func getQueryGroupBy(params map[string]string) string {
	groupByRaw, _ := params["$groupby"]
	if groupByRaw == "" {
		return ""
	}

	groupByParts := strings.Split(groupByRaw, ",")
	groupByArray := []string{}

	for _, field := range groupByParts {
		// Check for SQL injection
		if strings.Contains(field, " ") || strings.Contains(field, ";") {
			continue
		}

		groupByArray = append(groupByArray, field)
	}

	return strings.Join(groupByArray, ", ")
}

func getQueryOrder(params map[string]string) string {
	orderRaw := params["$order"]
	if orderRaw == "" {
		return ""
	}

	orderParts := strings.Split(orderRaw, ",")
	orderArray := []string{}
	for _, part := range orderParts {
		if len(part) < 2 {
			continue
		}
		if part[0] == '-' {
			orderArray = append(orderArray, part[1:]+" desc")
		} else {
			orderArray = append(orderArray, part)
		}
	}

	return strings.Join(orderArray, ", ")
}

func getQueryLimit(params map[string]string) string {
	limitRaw := params["$limit"]
	return limitRaw
}

func getQueryOffset(params map[string]string) string {
	offsetRaw := params["$offset"]
	return offsetRaw
}

func getQueryJoin(params map[string]string, tableName string) string {
	// $join syntax
	// {} required
	// [] optional
	// $join={to_table_name}__[join_method]__{[from_table.]from_column}__[[to_table.]to_column]
	joinType := map[string]string{
		"inner": "INNER JOIN",
		"outer": "FULL OUTER JOIN",
		"left":  "LEFT JOIN",
		"right": "RIGHT JOIN",
	}
	join := params["$join"]
	if join == "" {
		return ""
	}
	joinList := strings.Split(join, ",")
	joinFinalList := []string{}

	joinTmpl := "{JOIN_METHOD} {TO_TABLE} ON {FROM_TABLE}.{FROM_COLUMN}={TO_TABLE}.{TO_COLUMN}"

	for _, j := range joinList {
		jParts := strings.Split(j, "__")

		// Sanity Check
		if len(jParts) < 2 {
			continue
		}

		joinMethod := "INNER JOIN"
		fromTable := tableName
		toTable := jParts[0]
		fromColumn := ""
		toColumn := "id"
		index := 1

		// Check if the first part is a join method
		for k, v := range joinType {
			if k == jParts[index] {
				joinMethod = v
				index++
				break
			}
		}

		// Sanity Check
		if index == 2 && len(jParts) < 3 {
			continue
		}

		// Get from table/column
		if strings.Contains(jParts[index], ".") {
			fromParts := strings.Split(jParts[index], ".")
			fromTable = fromParts[0]
			fromColumn = fromParts[1]
		} else {
			fromColumn = jParts[index]
		}
		index++

		// Check if there is a to table/column
		if len(jParts) >= (index + 1) {
			if strings.Contains(jParts[index], ".") {
				toParts := strings.Split(jParts[index], ".")
				toTable = toParts[0]
				toColumn = toParts[1]
			} else {
				toColumn = jParts[index]
			}
		}
		joinStm := strings.Replace(joinTmpl, "{JOIN_METHOD}", joinMethod, -1)
		joinStm = strings.Replace(joinStm, "{TO_TABLE}", toTable, -1)
		joinStm = strings.Replace(joinStm, "{TO_COLUMN}", toColumn, -1)
		joinStm = strings.Replace(joinStm, "{FROM_TABLE}", fromTable, -1)
		joinStm = strings.Replace(joinStm, "{FROM_COLUMN}", fromColumn, -1)
		joinFinalList = append(joinFinalList, joinStm)
	}
	return strings.Join(joinFinalList, " ")
}

func getQueryM2M(params map[string]string, m interface{}, customSchema bool, modelName string) error {
	// $m2m=0
	// $m2m=$fill $m2m=1
	// $m2m=$id
	// $m2m=categories__fill
	// $m2m=categories__id
	// $m2m=categories__id,components__fill
	// $m2m=categories__id,components__id
	m2m := ""
	fillType := ""
	if params["$m2m"] == "" || params["$m2m"] == "fill" || params["$m2m"] == "1" {
		m2m = "1"
		fillType = "*"
	} else if params["$m2m"] == "id" {
		m2m = "1"
		fillType = "id"
	} else {
		m2m = params["$m2m"]
	}

	// Check if M2M is not required
	// TODO: Implement M2M for custom schema
	if m2m == "0" || customSchema {
		return nil
	}

	// Create a list of M2M
	// SELECT `cards`.*  FROM `cards` INNER JOIN `customer_card` ON `customer_card`.`table2_id`=`cards`.`id` WHERE `customer_card`.`table1_id` = 1
	// SELECT `cards`.id FROM `cards` INNER JOIN `customer_card` ON `customer_card`.`table2_id`=`cards`.`id` WHERE `customer_card`.`table1_id` = 1
	m2mTmpl := "SELECT `{TABLE_NAME}`.{FIELDS} FROM `{TABLE_NAME}` INNER JOIN `{M2M_TABLE_NAME}` ON `{M2M_TABLE_NAME}`.table2_id=`{TABLE_NAME}`.id WHERE `{M2M_TABLE_NAME}`.table1_id=? AND `{TABLE_NAME}`.deleted_at IS NULL"
	m2mStmt := map[string]string{}
	m2mModelName := map[string]string{}

	s := Schema[modelName]

	//table1 := s.TableName
	table2 := ""

	if m2m == "1" {
		for _, f := range s.Fields {
			if f.Type == cM2M {
				table2 = Schema[strings.ToLower(f.TypeName)].TableName
				m2mTable := s.ModelName + "_" + Schema[strings.ToLower(f.TypeName)].ModelName
				m2mStmt[f.Name] = m2mTmpl
				m2mStmt[f.Name] = strings.Replace(m2mStmt[f.Name], "{TABLE_NAME}", table2, -1)
				m2mStmt[f.Name] = strings.Replace(m2mStmt[f.Name], "{FIELDS}", fillType, -1)
				m2mStmt[f.Name] = strings.Replace(m2mStmt[f.Name], "{M2M_TABLE_NAME}", m2mTable, -1)
				m2mModelName[f.Name] = Schema[strings.ToLower(f.TypeName)].ModelName
			}
		}
	} else {
		m2mList := strings.Split(m2m, ",")
		for _, part := range m2mList {
			m2mParts := strings.Split(part, "__")
			if len(m2mParts) != 2 {
				//Trail(WARNING, "DAPI: invalid m2m syntax: %s", part)
				continue
			}
			if m2mParts[1] != "fill" && m2mParts[1] != "id" {
				continue
			} else if m2mParts[1] == "fill" {
				m2mParts[1] = "*"
			}

			f := Schema[modelName].FieldByName(m2mParts[0])
			table2Schema := Schema[strings.ToLower(f.TypeName)]
			table2 = table2Schema.TableName
			m2mTable := s.ModelName + "_" + table2Schema.ModelName
			m2mStmt[f.Name] = m2mTmpl
			m2mStmt[f.Name] = strings.Replace(m2mStmt[f.Name], "{TABLE_NAME}", table2, -1)
			m2mStmt[f.Name] = strings.Replace(m2mStmt[f.Name], "{FIELDS}", m2mParts[1], -1)
			m2mStmt[f.Name] = strings.Replace(m2mStmt[f.Name], "{M2M_TABLE_NAME}", m2mTable, -1)
			m2mModelName[f.Name] = Schema[strings.ToLower(f.TypeName)].ModelName
		}
	}

	mValue := reflect.ValueOf(m)
	for i := 0; i < mValue.Elem().Len(); i++ {
		for k, v := range m2mStmt {
			tempList, _ := NewModelArray(m2mModelName[k], true)
			GetDB().Raw(v, GetID(mValue.Elem().Index(i))).Scan(tempList.Interface())
			mValue.Elem().Index(i).FieldByName(k).Set(tempList.Elem())
		}
	}

	return nil
}

func returnDAPIJSON(w http.ResponseWriter, r *http.Request, a map[string]interface{}, params map[string]string, command string, model interface{}) error {
	if params["$stat"] == "1" {
		start := r.Context().Value(CKey("start"))

		if start != nil {
			sTime := start.(time.Time)
			a["stats"] = map[string]interface{}{
				"qtime": time.Now().Sub(sTime).String(),
			}
		}
	}

	if model != nil {
		if command == "read" {
			if postQuery, ok := model.(APIPostQueryReader); ok {
				if !postQuery.APIPostQueryRead(w, r, a) {
					return nil
				}
			}
		}
		if command == "add" {
			Trail(DEBUG, "pre2 %s, %#v", command, model)
			if postQuery, ok := model.(APIPostQueryAdder); ok {
				Trail(DEBUG, "pre3 %s, %#v", command, model)
				if !postQuery.APIPostQueryAdd(w, r, a) {
					Trail(DEBUG, "pre4 %s, %#v", command, model)
					return nil
				}
			}
		}
		if command == "edit" {
			if postQuery, ok := model.(APIPostQueryEditor); ok {
				if !postQuery.APIPostQueryEdit(w, r, a) {
					return nil
				}
			}
		}
		if command == "delete" {
			if postQuery, ok := model.(APIPostQueryDeleter); ok {
				if !postQuery.APIPostQueryDelete(w, r, a) {
					return nil
				}
			}
		}
		if command == "schema" {
			if postQuery, ok := model.(APIPostQuerySchemer); ok {
				if !postQuery.APIPostQuerySchema(w, r, a) {
					return nil
				}
			}
		}
	}

	ReturnJSON(w, r, a)
	return nil
}

type APILogReader interface {
	APILogRead(*http.Request) bool
}

type APILogEditor interface {
	APILogEdit(*http.Request) bool
}

type APILogAdder interface {
	APILogAdd(*http.Request) bool
}

type APILogDeleter interface {
	APILogDelete(*http.Request) bool
}

type APILogSchemer interface {
	APILogSchema(*http.Request) bool
}

type APIPublicReader interface {
	APIPublicRead(*http.Request) bool
}

type APIPublicEditor interface {
	APIPublicEdit(*http.Request) bool
}

type APIPublicAdder interface {
	APIPublicAdd(*http.Request) bool
}

type APIPublicDeleter interface {
	APIPublicDelete(*http.Request) bool
}

type APIPublicSchemer interface {
	APIPublicSchema(*http.Request) bool
}

type APIDisabledReader interface {
	APIDisabledRead(*http.Request) bool
}

type APIDisabledEditor interface {
	APIDisabledEdit(*http.Request) bool
}

type APIDisabledAdder interface {
	APIDisabledAdd(*http.Request) bool
}

type APIDisabledDeleter interface {
	APIDisabledDelete(*http.Request) bool
}

type APIDisabledSchemer interface {
	APIDisabledSchema(*http.Request) bool
}

type APIPreQueryReader interface {
	APIPreQueryRead(http.ResponseWriter, *http.Request) bool
}

type APIPostQueryReader interface {
	APIPostQueryRead(http.ResponseWriter, *http.Request, map[string]interface{}) bool
}

type APIPreQueryAdder interface {
	APIPreQueryAdd(http.ResponseWriter, *http.Request) bool
}

type APIPostQueryAdder interface {
	APIPostQueryAdd(http.ResponseWriter, *http.Request, map[string]interface{}) bool
}

type APIPreQueryEditor interface {
	APIPreQueryEdit(http.ResponseWriter, *http.Request) bool
}

type APIPostQueryEditor interface {
	APIPostQueryEdit(http.ResponseWriter, *http.Request, map[string]interface{}) bool
}

type APIPreQueryDeleter interface {
	APIPreQueryDelete(http.ResponseWriter, *http.Request) bool
}

type APIPostQueryDeleter interface {
	APIPostQueryDelete(http.ResponseWriter, *http.Request, map[string]interface{}) bool
}

type APIPreQuerySchemer interface {
	APIPreQuerySchema(http.ResponseWriter, *http.Request) bool
}

type APIPostQuerySchemer interface {
	APIPostQuerySchema(http.ResponseWriter, *http.Request, map[string]interface{}) bool
}
