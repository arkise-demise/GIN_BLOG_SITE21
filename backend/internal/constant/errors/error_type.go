package errors

import (
	"net/http"

	"github.com/joomcode/errorx"
)

type ErrorType struct {
	StatusCode int
	ErrorType  *errorx.Type
}

var Error = []ErrorType{
	{
		StatusCode: http.StatusBadRequest,
		ErrorType:  ErrInvalidUserInput,
	},
	{
		StatusCode: http.StatusForbidden,
		ErrorType:  ErrAcessError,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrInternalServerError,
	},
	{
		StatusCode: http.StatusBadRequest,
		ErrorType:  ErrAuthClient,
	},
	{
		StatusCode: http.StatusUnauthorized,
		ErrorType:  ErrSSOAuthenticationFailed,
	},
	{
		StatusCode: http.StatusUnauthorized,
		ErrorType:  ErrInvalidAccessToken,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrSSOError,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrUnableToGet,
	},

	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrUnableTocreate,
	},
	{
		StatusCode: http.StatusNotFound,
		ErrorType:  ErrResourceNotFound,
	},

	{
		StatusCode: http.StatusBadRequest,
		ErrorType:  ErrDataAlredyExist,
	},
	{
		StatusCode: http.StatusNotFound,
		ErrorType:  ErrNoRecordFound,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrDBDelError,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrAccountingClient,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrMerchantClient,
	},
	{
		StatusCode: http.StatusBadRequest,
		ErrorType:  ErrEventNotSupported,
	},
	{
		StatusCode: http.StatusBadRequest,
		ErrorType:  ErrKafkaEventNotSupported,
	},
	{
		StatusCode: http.StatusGone,
		ErrorType:  ErrSocketConnectionClosed,
	},
	{
		StatusCode: http.StatusGone,
		ErrorType:  ErrSocketConnectionReset,
	},
	{
		StatusCode: http.StatusGone,
		ErrorType:  ErrSocketConnectionBroken,
	},
	{
		StatusCode: http.StatusRequestTimeout,
		ErrorType:  ErrDeadlineTimedOut,
	},
	{
		StatusCode: http.StatusBadRequest,
		ErrorType:  ErrProgramStatus,
	},
	{
		StatusCode: http.StatusBadRequest,
		ErrorType:  ErrProgramAmount,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrAccountingError,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrSMSSend,
	},
	{
		StatusCode: http.StatusBadGateway,
		ErrorType:  ErrMapsRequest,
	},

	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrMapsResponse,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrUnableToUpdate,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrHTTPRequestPrepareFailed,
	},
	{
		StatusCode: http.StatusForbidden,
		ErrorType:  ErrTripDeviceChange,
	},
	{
		StatusCode: http.StatusBadRequest,
		ErrorType:  ErrSocketReadLimitExceeded,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrOndeRequest,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrNotificationChannel,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrFirebaseClient,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrFirebaseSendNotification,
	},
	{
		StatusCode: http.StatusBadRequest,
		ErrorType:  ErrInsufficientBalance,
	},
	{
		StatusCode: http.StatusBadRequest,
		ErrorType:  ErrBadRequest,
	},
	{
		StatusCode: http.StatusConflict,
		ErrorType:  ErrLoyalityExchange,
	},
	{
		StatusCode: http.StatusInternalServerError,
		ErrorType:  ErrEfloatRequest,
	},
}

// list of error namespaces
var (
	databaseError           = errorx.NewNamespace("database error").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	invalidInput            = errorx.NewNamespace("validation error").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	resourceNotFound        = errorx.NewNamespace("not found").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	unauthorized            = errorx.NewNamespace("unauthorized").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	ineligible              = errorx.NewNamespace("ineligible").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	AccessDenied            = errorx.RegisterTrait("You are not authorized to perform the action")
	Ineligible              = errorx.RegisterTrait("You are not eligible to perform the action")
	serverError             = errorx.NewNamespace("server error")
	authoriztionClientError = errorx.NewNamespace("authorization client error")
	accountingClientError   = errorx.NewNamespace("accounting client error")
	merchantClientError     = errorx.NewNamespace("merchant client error")
	Unauthenticated         = errorx.NewNamespace("user authentication failed")
	ProgramError            = errorx.NewNamespace("program error")
	websocketError          = errorx.NewNamespace("websocket error")
	kafkaError              = errorx.NewNamespace("kafka error")
	mapsError               = errorx.NewNamespace("maps error")
	httpError               = errorx.NewNamespace("http error")
	ondeError               = errorx.NewNamespace("onde request error")
	merchantportalError     = errorx.NewNamespace("merchant portal error")
	notificationError       = errorx.NewNamespace("notification channel error")
	firebaseError           = errorx.NewNamespace("firebase error")
	badRequest              = errorx.NewNamespace("bad request error")
	LoyaltyExhange          = errorx.NewNamespace("loyalty exhange error")
	efloatError             = errorx.NewNamespace("e-float error")
)

// list of errors types in all of the above namespaces

var (
	ErrUnableTocreate          = errorx.NewType(databaseError, "unable to create")
	ErrDataAlredyExist         = errorx.NewType(databaseError, "data alredy exist")
	ErrUnableToGet             = errorx.NewType(databaseError, "unable to get")
	ErrInvalidUserInput        = errorx.NewType(invalidInput, "invalid user input")
	ErrInactiveUserStatus      = errorx.NewType(invalidInput, "Inactive user status")
	ErrTripDeviceChange        = errorx.NewType(invalidInput, "user changed device")
	ErrResourceNotFound        = errorx.NewType(resourceNotFound, "resource not found")
	ErrAcessError              = errorx.NewType(unauthorized, "Unauthorized", AccessDenied)
	ErrIneligibleError         = errorx.NewType(ineligible, "Ineligible", Ineligible)
	ErrInternalServerError     = errorx.NewType(serverError, "internal server error")
	ErrAuthClient              = errorx.NewType(authoriztionClientError, "authorization client error")
	ErrSSOAuthenticationFailed = errorx.NewType(Unauthenticated, "user authentication failed")
	ErrInvalidAccessToken      = errorx.NewType(Unauthenticated, "invalid token").
					ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	ErrSSOError                   = errorx.NewType(serverError, "sso communication failed")
	ErrAccountingError            = errorx.NewType(serverError, "accounting error")
	ErrUnExpectedError            = errorx.NewType(serverError, "unexpected error occurred")
	ErrUnableToUpdate             = errorx.NewType(databaseError, "unable to update")
	ErrDBDelError                 = errorx.NewType(databaseError, "could not delete record")
	ErrNoRecordFound              = errorx.NewType(resourceNotFound, "no record found")
	ErrAccountingClient           = errorx.NewType(accountingClientError, "accounting client error")
	ErrEventNotSupported          = errorx.NewType(websocketError, "event type not supported")
	ErrDeadlineTimedOut           = errorx.NewType(websocketError, "read or write deadline timedout")
	ErrSocketConnectionClosed     = errorx.NewType(websocketError, "socket client connection closed")
	ErrSocketConnectionBroken     = errorx.NewType(websocketError, "broken pipe")
	ErrSocketConnectionReset      = errorx.NewType(websocketError, "connection reset by client")
	ErrProgramStatus              = errorx.NewType(ProgramError, "program status error")
	ErrProgramAmount              = errorx.NewType(ProgramError, "spending limit error")
	ErrSMSSend                    = errorx.NewType(serverError, "couldn't send sms")
	ErrMapsRequest                = errorx.NewType(mapsError, "couldn't make maps request")
	ErrMapsResponse               = errorx.NewType(mapsError, "got error from maps")
	ErrHTTPRequestPrepareFailed   = errorx.NewType(httpError, "couldn't prepare http request")
	ErrSocketReadLimitExceeded    = errorx.NewType(websocketError, "socket read limit exceeded")
	ErrOndeRequest                = errorx.NewType(ondeError, "Onde request error")
	ErrMerchantPortalCashOutError = errorx.NewType(merchantportalError, "merchant portal cash out error")
	ErrMerchantClient             = errorx.NewType(merchantClientError, "merchant client error")
	ErrKafkaEventNotSupported     = errorx.NewType(kafkaError, "event type not supported")
	ErrNotificationChannel        = errorx.NewType(notificationError, "notification channel error")
	ErrFirebaseClient             = errorx.NewType(firebaseError, "firebase client error")
	ErrFirebaseSendNotification   = errorx.NewType(firebaseError, "fire base notification sending error")
	ErrInsufficientBalance        = errorx.NewType(accountingClientError, "account insufficient balance error")
	ErrBadRequest                 = errorx.NewType(badRequest, "bad request error")
	ErrLoyalityExchange           = errorx.NewType(LoyaltyExhange, "sewasew loyalty exhange failed")
	ErrEfloatRequest              = errorx.NewType(efloatError, "e-float request failed")
	ErrMiniRideRequest            = errorx.NewType(serverError, "mini ride request failed")
)
