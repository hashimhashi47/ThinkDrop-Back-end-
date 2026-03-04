package constants

import "errors"

// role
const (
	User           = "user"
	Admin          = "admin"
	RoleSuperAdmin = "superadmin"
	RoleManager    = "manager"
	RoleSupervisor = "supervisor"

	RoleModerator = "moderator"
	RoleStaff     = "staff"
	RoleCoStaff   = "costaff"
	RoleSupport   = "support"

	RoleAccountant = "accountant"
	RoleAuditor    = "auditor"
)

const (
	Sucess = "Sucess"
	Error  = "Error"
)

var (
	ErrSignupFailed           = "signup failed"
	ErrUserNotFound           = "user not found"
	ErrInvalidPassword        = "invalid password"
	ErrAccessTokenCreateFail  = "failed to create access token"
	ErrRefreshTokenCreateFail = "failed to create refresh token"
	ErrRefreshTokenStoreFail  = "failed to store refresh token"
)

type ReportType string

const (
	ReportBug       ReportType = "bug"
	ReportTechnical ReportType = "technical"
	ReportOther     ReportType = "other"
)

// WalletStatus
const (
	WalletActive   = "active"
	WalletInactive = "inactive"
	WalletBlocked  = "blocked"
)

var (
	SUCESSCODE           = 200
	BADREQUEST           = 400
	UNAUTHORIZED         = 401
	FORBIDDEN            = 403
	NOTFOUND             = 404
	TOOMANYREQUESTS      = 429
	METHODNOTALLOWED     = 405
	INTERNALSERVERERROR  = 500
	NOTIMPLEMENTED       = 501
	BADGATEWAY           = 502
	SERVICEUNAVAILABLE   = 503
	GATEWAYTIMEOUT       = 504
	UNSUPPORTEDMEDIATYPE = 415
	UNPROCESSABLEENTITY  = 422
	STATUSCONFLIT        = 409

	UNAVAILABLE    = "UNAVAILABLE"
	PENDING        = "PENDING"
	COMPLETED      = "COMPLETED"
	APPROVED       = "APPROVED"
	ASSIGNEDSLOT   = "ASSIGNEDSLOT"
	ACCEPTED       = "ACCEPTED"
	REJECTED       = "REJECTED"
	SERVICESTARTED = "SERVICE STARTED"
	SERVICEENDED   = "SERVICE ENDED"
	DELIVERED      = "DELIVERED"
	WAITING        = "WAITING"

	// New page constants
	PLAN_PAGE     = "PLAN_PAGE"
	INVOICE_PAGE  = "INVOICE_PAGE"
	CHECKOUT_PAGE = "CHECKOUT_PAGE"
	PLAN_STATIC   = 7

	// preload
	PRELOADUSER     = "Users"
	PRELOADSTAFF    = "Staff"
	PRELOADSLOT     = "Slot"
	PRELOADBOOKINGS = "Bookings"
	PRELOADBOOKED   = "Bookeds"

	//errors
	ErrTokenNotFound = errors.New("token not found")
	ErrTokenExpired  = errors.New("token expired")
	ErrTokenMismatch = errors.New("token mismatch")
)
