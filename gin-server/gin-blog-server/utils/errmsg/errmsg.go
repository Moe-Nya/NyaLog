package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// JWT中间件
	TokenNotExist       = 1001
	TokenInvalid        = 1002
	TokenParseFailed    = 1003
	TokenGenerateFailed = 1004

	// 2FA
	GenerateSecretFailed = 3001
	CodeError            = 3002

	// QRcode
	GenerateQRFailed = 4001

	// 管理员用户
	UserInfoError       = 2001
	UserExist           = 2002
	UserNotExist        = 2003
	UserNotValidate     = 2004
	CheckUserFailed     = 2005
	LoginError          = 2006
	UidError            = 2007
	PasswordError       = 2008
	CantLogin           = 2009
	ResetPasswordFailed = 2010
	QueryUserFailed     = 2011

	// 邮件发送
	SendEmailFailed   = 5001
	ValidateCodeError = 5002

	// 博客设置
	BlogNotSet    = 6001
	BlogSetFailed = 6002

	// 文章管理
	ArticleUpdateFailed = 7001
	ArticleQueryFailed  = 7002
	ArticleDeleteFailed = 7003

	// 文章分类管理
	CatCreateFailed = 8001
	CatQueryFailed  = 8002
	CatUpdateFailed = 8003
	CatDeleteFailed = 8004

	// 页面管理
	PageCreateFailed = 9001
	PageQueryFailed  = 9002
	PageEditFailed   = 9003
	PageDeleteFailed = 9004

	// 评论管理
	CreateQueryFailed   = 10001
	GetQueryFailed      = 10002
	SendCommentFailed   = 10003
	SelectCommentFailed = 10004
	DeleteCommentFailed = 1005

	// findme 管理
	CreateFindMeFailed = 11001
	FindMeQueryFailed  = 11002
	FindMeModifyFailed = 11003
	DeleteFindMeFailed = 11004

	// 导航栏管理
	CreateNavFailed = 12001
	NavQueryFailed  = 12002
	NavEditFailed   = 12003
	NavDeleteFailed = 12004

	// 友链管理
	CreateFriendLinkFailed = 13001
	FriendLinkQueryFailed  = 13002
	FriendLinkModifyFailed = 13003
	FriendLinkDeleteFailed = 13004
)

var codeMsg = map[int]string{
	SUCCESS: "ok",
	ERROR:   "error",

	// JWT中间件
	TokenNotExist:       "token not exist",
	TokenInvalid:        "token invalid",
	TokenParseFailed:    "failed to parse claims",
	TokenGenerateFailed: "token generate failed",

	// 2FA
	GenerateSecretFailed: "generate secret failed",
	CodeError:            "code error",

	// QRcode
	GenerateQRFailed: "generate qrcode failed",

	// 管理员用户
	UserInfoError:       "user information error",
	UserExist:           "user exist",
	UserNotExist:        "user not exist",
	UserNotValidate:     "user not validate",
	CheckUserFailed:     "check user failed",
	LoginError:          "an error occurred while logging in",
	UidError:            "uid error",
	PasswordError:       "password error",
	CantLogin:           "this ipdress can't login agin",
	ResetPasswordFailed: "reset password failed",
	QueryUserFailed:     "query user failed",

	// 邮件发送
	SendEmailFailed:   "email send failed",
	ValidateCodeError: "validate coede error",

	// 博客设置
	BlogNotSet:    "blog not set",
	BlogSetFailed: "blog set failed",

	// 文章管理
	ArticleUpdateFailed: "article update failed",
	ArticleQueryFailed:  "article query failed",
	ArticleDeleteFailed: "article delete failed",

	// 文章分类管理
	CatCreateFailed: "category create failed",
	CatQueryFailed:  "category query failed",
	CatUpdateFailed: "category update failed",
	CatDeleteFailed: "category delete failed",

	// 页面管理
	PageCreateFailed: "page create failed",
	PageQueryFailed:  "page query failed",
	PageEditFailed:   "page edit failed",
	PageDeleteFailed: "page delete failed",

	// 评论管理
	CreateQueryFailed:   "create query failed",
	GetQueryFailed:      "get query failed",
	SendCommentFailed:   "send comment failed",
	SelectCommentFailed: "select comment failed",
	DeleteCommentFailed: "delete comment failed",

	// findme 管理
	CreateFindMeFailed: "create findme failed",
	FindMeQueryFailed:  "findme query failed",
	FindMeModifyFailed: "findme modify failed",
	DeleteFindMeFailed: "delete findme failed",
	CreateNavFailed:    "create navigation failed",
	NavQueryFailed:     "navigation query failed",
	NavEditFailed:      "navigation edit failed",
	NavDeleteFailed:    "navigation delete failed",

	// 友链管理
	CreateFriendLinkFailed: "create friend link failed",
	FriendLinkQueryFailed:  "friend link query failed",
	FriendLinkModifyFailed: "friend link modify failed",
	FriendLinkDeleteFailed: "friend link delete failed",
}

func GetErrorMsg(code int) string {
	return codeMsg[code]
}
