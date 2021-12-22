package response

import (
	"fmt"
	customError "golang_hexagonal_architecture/utils/error"
)

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
type Response struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

func NewResponse(data interface{}, meta Meta) Response {
	return Response{
		Data: data,
		Meta: meta,
	}
}
func MetaFromError(err interface{}) Meta {
	if dataErr, ok := err.(*customError.Error); ok {
		meta := DefaultMetaError
		if data, ok := ConstantMeta[dataErr.Error.Error()]; ok {
			meta = data
			newMessage := fmt.Sprintf("%s %s", dataErr.Name, meta.Message)
			meta.Message = newMessage
		}
		return meta
	}
	if dataErr, ok := err.(error); ok {
		meta := DefaultMetaError
		if data, ok := ConstantMeta[dataErr.Error()]; ok {
			meta = data
		} else {
			meta.Message = dataErr.Error()
		}
		return meta
	}
	return DefaultMetaError
}
