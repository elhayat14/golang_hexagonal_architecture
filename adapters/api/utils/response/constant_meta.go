package response

var DefaultMeta = Meta{
	Code:    0,
	Message: "Success",
	Status:  true,
}
var DefaultMetaError = Meta{
	Code:    99,
	Message: "General Error",
	Status:  false,
}
var ConstantMeta = map[string]Meta{
	//400
	"created": {Code: 0, Message: "Success create", Status: true},
	"update":  {Code: 0, Message: "Success update", Status: true},
	//db
	"data_not_found": {Code: 1, Message: "Not Found", Status: false},
	"not_valid":      {Code: 2, Message: "Not Valid", Status: false},
	"error_add_data": {Code: 3, Message: "Failed To Add", Status: false},
}
