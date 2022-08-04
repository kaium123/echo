package utility

// var validate = validator.New()

// func CheckJsonBody(req model.Product) []string {
// 	var str []string
// 	err := validate.Struct(req)

// 	if err != nil {

// 		for _, err := range err.(validator.ValidationErrors) {
// 			if err.Kind() == reflect.String {
// 				if err.Value() != "" {
// 					tmp := "Invalid "
// 					tmp += err.StructField()
// 					str = append(str, tmp)
// 				}
// 			}

// 			if err.Kind() == reflect.Int32 {
// 				val := err.Value()
// 				var zero int32 = 0
// 				if val != zero {
// 					tmp := "Invalid "
// 					tmp += err.StructField()
// 					str = append(str, tmp)
// 				}
// 			}
// 		}
// 	}

// 	return str
// }

// func GetErrorFeilds(err error) []string {
// 	var str []string
// 	for _, err := range err.(validator.ValidationErrors) {
// 		s := "Invalid "
// 		s += err.StructField()
// 		str = append(str, s)
// 	}
// 	return str
// }
