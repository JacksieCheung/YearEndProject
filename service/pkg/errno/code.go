package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrQuery            = &Errno{Code: 10003, Message: "Error occurred while getting url queries."}
	ErrPathParam        = &Errno{Code: 10004, Message: "Error occurred while getting path param."}

	ErrDatabase = &Errno{Code: 20002, Message: "Database error."}

	ErrAuthFailed = &Errno{Code: 30001, Message: "Wrong number or password"}
	ErrAuthParam  = &Errno{Code: 30002, Message: "Create param failed."}

	ErrAtoi     = &Errno{Code: 40001, Message: "Atoi error."}
	ErrDecoding = &Errno{Code: 40002, Message: "Base64 decoding error."}
)
