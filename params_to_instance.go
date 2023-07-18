package uadmin

// func setParams(params map[string]string, m reflect.Value, schema ModelSchema) (reflect.Value, error) {
// 	paramMap := map[string]interface{}{}
// 	for k, v := range params {
// 		key := k
// 		if key == "" {
// 			continue
// 		}
// 		if key[0] == '_' {
// 			key = key[1:]
// 		}
// 		f := schema.FieldByColumnName(key)
// 		if f != nil {
// 			key = f.Name
// 		}
// 		paramMap[key] = v

// 		// fix value for numbers
// 		if f.Type == cNUMBER {
// 			if strings.HasPrefix(f.TypeName, "float") {
// 				paramMap[key], _ = strconv.ParseFloat(v, 64)
// 			} else if strings.HasPrefix(f.TypeName, "uint") {
// 				paramMap[key], _ = strconv.ParseUint(v, 10, 64)
// 			} else if strings.HasPrefix(f.TypeName, "int") {
// 				paramMap[key], _ = strconv.ParseInt(v, 10, 64)
// 			}
// 		} else if f.Type == cBOOL {
// 			if paramMap[key] == "true" || paramMap[key] == "1" {
// 				paramMap[key] = true
// 			} else {
// 				paramMap[key] = false
// 			}
// 		} else if f.Type == cLIST {
// 			paramMap[key], _ = strconv.ParseInt(v, 10, 64)
// 		} else if f.Type == cDATE {

// 		}
// 	}

// 	buf, _ := json.Marshal(params)
// 	var err error
// 	if m.Kind() == reflect.Pointer {
// 		err = json.Unmarshal(buf, m.Interface())
// 	} else if m.Kind() == reflect.Struct {
// 		err = json.Unmarshal(buf, m.Addr().Interface())
// 	}

// 	return m, err
// }

// func parseDate(v string) interface{} {
// 	if v == "" || v == "null" {
// 		return nil
// 	}
// 	dt, err := time.Parse("2006-05-04T15:02:01Z", v)
// 	if err != nil {
// 		return dt
// 	}
// }
