package messages

const (
	SuccessMessage                           = "success"
	FailedMessage                            = "failed"
	TokenIsNotProvidedMessage                = "token is not provided, please requests with token"
	TokenIsNotValidMessage                   = "token is not valid"
	ApiKeyIsNotProvidedMessage               = "apikey is not provided, please requests with apikey"
	ApiKeyIsNotValidMessage                  = "apikey is not valid"
	CredentialIsNotMatchMessage              = "credential is not match"
	AccountIsNotVerifiedMessage              = "your account is not verified yet"
	AccountIsNotBiasaRole                    = "your account cannot access the biasa page"
	AccountIsNotAdminRole                    = "your account cannot access the admin page"
	TokenIsExpiredMessage                    = "token is expired, please re-login or revoke token with refresh token"
	JWTSecretKeyIsMissingMessage             = "please provide the jwt secret key in the environment"
	JWTMiddlewareIsMissingMessage            = "please add the jwt middleware new first, before proceeding any advanced jwt middleware"
	InterfaceConversionErrorMessage          = "interface conversion error"
	PasswordAndConfirmPasswordNotSameMessage = "password and confirm password doesn't match"
	FailedBindQuery                          = "failed query binding"
)
