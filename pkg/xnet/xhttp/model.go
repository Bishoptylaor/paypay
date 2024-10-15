package xhttp

const (
	ResTypeJSON = "json"
	ResTypeXML  = "xml"

	TypeJSON              = "json"
	TypeXML               = "xml"
	TypeFormData          = "form-data"
	TypeMultipartFormData = "multipart-form-data"
	// TypeMultipartRelated  = "multipart-related"
)

var (
	_ReqContentTypeMap = map[string]string{
		TypeJSON:              "application/json",
		TypeXML:               "application/xml",
		TypeFormData:          "application/x-www-form-urlencoded",
		TypeMultipartFormData: "multipart/form-data",
		// TypeMultipartRelated:  "multipart/related",
	}

	_ResContentTypeMap = map[string]string{
		ResTypeJSON: "application/json",
		ResTypeXML:  "application/xml",
	}
)
