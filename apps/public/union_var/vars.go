package union_var

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	TRACE_ID_KEY           = "trace_id_key"
	X_TRACE_ID_HEADER      = "X-Trace-ID"
	JWT_TYPE_ACCESS_TOKEN  = "access"
	JWT_TYPE_REFRESH_TOKEN = "refresh"
	JWT_CONTEXT_KEY        = "jwt_context_key"
	JWT_ROLE_ADMIN         = "jwt_role_admin"
	JWT_ROLE_GUEST         = "jwt_role_guest"
	JWT_ROLE_USER          = "jwt_role_user"
	MINIO_EXPIRE_TIME      = 3 * time.Hour
	MINIO_ON_CONTINUE_TIME = 2 * time.Hour
)

// jwt相关
var (
	SigningMethod = jwt.SigningMethodHS256 // 签名方法
	AccessSecret  = []byte("redrock_video_jwt_access_secret")
	RefreshSecret = []byte("redrock_video_jwt_refresh_secret")
	Issuer        = "redrock.video.auth.gateway"
	AccessTTL     = 24 * time.Hour     // 访问令牌有效期
	RefreshTTL    = 7 * 24 * time.Hour // 刷新令牌有效期
)
