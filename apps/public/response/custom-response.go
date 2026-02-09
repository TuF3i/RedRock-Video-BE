package response

var (
	OperationSuccess        = Response{Status: 10200, Info: "Operation Success"}           // OperationSuccess 执行成功
	EmptyJWTString          = Response{Status: 10001, Info: "Empty JWT String"}            // EmptyJWTString 空JWT字符串
	JWTNotRegisteredInRedis = Response{Status: 10002, Info: "JWT Not Registered In Redis"} // JWTNotRegisteredInRedis JWT未注册或已过期
	ValidateRequestFail     = Response{Status: 10003, Info: "Validate Request Fail"}       // ValidateRequestFail 校验请求失败
	EmptyRVID               = Response{Status: 10004, Info: "Empty RVID"}                  // EmptyRVID 空RVID
	YouDoNotHaveAccess      = Response{Status: 10005, Info: "You Do Not Have Access"}      // YouDoNotHaveAccess 你没有权限
)
