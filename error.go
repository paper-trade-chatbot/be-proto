package models

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrCode int32

const (
	ErrCode_AccountRegistered         ErrCode = 1001
	ErrCode_AccountNotRegistered      ErrCode = 1002
	ErrCode_PhoneRegistered           ErrCode = 1003
	ErrCode_PasswordFormatMismatch    ErrCode = 1004
	ErrCode_PasswordMismatch          ErrCode = 1005
	ErrCode_SocialTokenMismatch       ErrCode = 1006
	ErrCode_PhoneVerifiedMismatch     ErrCode = 1007
	ErrCode_InviteCodeInvalid         ErrCode = 1008
	ErrCode_MailVerifiedMismatch      ErrCode = 1009
	ErrCode_PhoneNotRegistered        ErrCode = 1010
	ErrCode_MailNotRegistered         ErrCode = 1011
	ErrCode_MailRegistered            ErrCode = 1012
	ErrCode_InvalidCountry            ErrCode = 1013
	ErrCode_InvalidBank               ErrCode = 1014
	ErrCode_InvalidBankAccount        ErrCode = 1015
	ErrCode_InvalidIdNumber           ErrCode = 1016
	ErrCode_IdNumberRegistered        ErrCode = 1017
	ErrCode_AccountOrPasswordMismatch ErrCode = 1018
	ErrCode_BalanceNotEnough          ErrCode = 2001
	ErrCode_StreamEnded               ErrCode = 3001
	ErrCode_UserNotInRoom             ErrCode = 3002
	ErrCode_UserNoAuth                ErrCode = 4001
	ErrCode_RecordNotFound            ErrCode = 5001
	ErrCode_RequestParamsError        ErrCode = 5002
	ErrCode_NoPermissionToEdit        ErrCode = 5003
	ErrCode_NoPermissionToDelete      ErrCode = 5004
	ErrCode_ServicePanicError         ErrCode = 5005
	ErrCode_NotImplemented            ErrCode = 5006
	ErrCode_NotTradingTime            ErrCode = 6001
	ErrCode_PositionNotEnough         ErrCode = 6002
	ErrCode_LeverageNotTheSame        ErrCode = 6003
	ErrCode_NoQuote                   ErrCode = 6004
	ErrCode_NotEnabled                ErrCode = 6005
	ErrCode_FracShareBelowOne         ErrCode = 6006
	// ErrCode_StopLossPriceLTEGainProfitPrice ErrCode = 6007
	// ErrCode_StopLossPriceGTEGainProfitPrice ErrCode = 6008
	ErrCode_MarketPriceLTEStopLossPrice       ErrCode = 6007
	ErrCode_MarketPriceGTEStopLossPrice       ErrCode = 6008
	ErrCode_MarketPriceGTEGainProfitPrice     ErrCode = 6009
	ErrCode_MarketPriceLTEGainProfitPrice     ErrCode = 6010
	ErrCode_OrderPriceLTEStopLossPrice        ErrCode = 6011
	ErrCode_OrderPriceGTEStopLossPrice        ErrCode = 6012
	ErrCode_OrderPriceGTEGainProfitPrice      ErrCode = 6013
	ErrCode_OrderPriceLTEGainProfitPrice      ErrCode = 6014
	ErrCode_ClosePositionHaveStopLossFlag     ErrCode = 6015
	ErrCode_OrderHaveStopLossFlagWithPosition ErrCode = 6016
	ErrCode_NotUnsettledOrder                 ErrCode = 6017
	ErrCode_TradeAmountLessThanMinimum        ErrCode = 6018
	ErrCode_OverTimeInterval                  ErrCode = 6019
	ErrCode_ReadOnly                          ErrCode = 6020
	ErrCode_GetSecretError                    ErrCode = 7001
	ErrCode_DepositError                      ErrCode = 7002
	ErrCode_DepositNotificationError          ErrCode = 7003
	ErrCode_WithdrawBalanceNotEnough          ErrCode = 7004
	ErrCode_EditExchangeRateError             ErrCode = 7005
	ErrCode_EditWithdrawError                 ErrCode = 7006
	ErrCode_AddExchangeRateError              ErrCode = 7007
	ErrCode_DeleteExchangeRateError           ErrCode = 7008
	ErrCode_WithdrawLossTooMuch               ErrCode = 7009
	ErrCode_WithdrawGuaranteeNotEnough        ErrCode = 7010
	ErrCode_ExceedMaximumDeposit              ErrCode = 7011
	ErrCode_LessMinimumDeposit                ErrCode = 7012

	ErrCode_APIRequestTooMany            ErrCode = 8001
	ErrCode_InvalidUploadArgument        ErrCode = 8002
	ErrCode_InvalidDashboardArgument     ErrCode = 8003
	ErrCode_ExceedMaximumReferralUsers   ErrCode = 9001
	ErrCode_ExceedMaximumAttendance      ErrCode = 9002
	ErrCode_UserPromotionReviewing       ErrCode = 9003
	ErrCode_EventNotGoing                ErrCode = 9004
	ErrCode_LessThanMinimumDepositAmount ErrCode = 9005
	ErrCode_ExceedEventIntervalDays      ErrCode = 9006
	ErrCode_NotFirstDeposit              ErrCode = 9007
	ErrCode_LessThanMinimumTradeAmount   ErrCode = 9008
	ErrCode_NotFirstTrade                ErrCode = 9009
	ErrCode_AttendedEvent                ErrCode = 9010
	ErrCode_EventNeedApply               ErrCode = 9011
	ErrCode_BudgetNotEnough              ErrCode = 9012
	ErrCode_NoPostPermission             ErrCode = 10001
	ErrCode_NoCommentPermission          ErrCode = 10002
	ErrCode_InvalidPostID                ErrCode = 10003
	ErrCode_InvalidCommentID             ErrCode = 10004
)

var (
	// Account
	ErrAccountRegistered         = status.Error(codes.Code(ErrCode_AccountRegistered), "account registered")
	ErrAccountNotRegistered      = status.Error(codes.Code(ErrCode_AccountNotRegistered), "account not registered")
	ErrPhoneRegistered           = status.Error(codes.Code(ErrCode_PhoneRegistered), "phone number registered")
	ErrPasswordFormatMismatch    = status.Error(codes.Code(ErrCode_PasswordFormatMismatch), "password format mismatch")
	ErrPasswordMismatch          = status.Error(codes.Code(ErrCode_PasswordMismatch), "password mismatch")
	ErrSocialTokenMismatch       = status.Error(codes.Code(ErrCode_SocialTokenMismatch), "social token mismatch")
	ErrPhoneVerifiedMismatch     = status.Error(codes.Code(ErrCode_PhoneVerifiedMismatch), "phone verified mismatch")
	ErrMailVerifiedMismatch      = status.Error(codes.Code(ErrCode_MailVerifiedMismatch), "mail verified mismatch")
	ErrInviteCodeInvalid         = status.Error(codes.Code(ErrCode_InviteCodeInvalid), "invite code invalid")
	ErrPhoneNotRegistered        = status.Error(codes.Code(ErrCode_PhoneNotRegistered), "phone number not registered")
	ErrMailNotRegistered         = status.Error(codes.Code(ErrCode_MailNotRegistered), "mail not registered")
	ErrMailRegistered            = status.Error(codes.Code(ErrCode_MailRegistered), "mail registered")
	ErrInvalidCountry            = status.Error(codes.Code(ErrCode_InvalidCountry), "invalid country")
	ErrInvalidBank               = status.Error(codes.Code(ErrCode_InvalidBank), "invalid bank")
	ErrInvalidBankAccount        = status.Error(codes.Code(ErrCode_InvalidBankAccount), "invalid bank account")
	ErrInvalidIdNumber           = status.Error(codes.Code(ErrCode_InvalidIdNumber), "invalid id number")
	ErrIdNumberRegistered        = status.Error(codes.Code(ErrCode_IdNumberRegistered), "id number registered")
	ErrAccountOrPasswordMismatch = status.Error(codes.Code(ErrCode_AccountOrPasswordMismatch), "account or password mismatch")

	// Wallet
	ErrBalanceNotEnough = status.Error(codes.Code(ErrCode_BalanceNotEnough), "balance not enough")

	// Live
	ErrStreamEnded   = status.Error(codes.Code(ErrCode_StreamEnded), "stream ended")
	ErrUserNotInRoom = status.Error(codes.Code(ErrCode_UserNotInRoom), "user not in room")

	// Auth
	ErrRoleNoAuth         = status.Error(codes.Code(ErrCode_UserNoAuth), "this role has no auth")
	ErrAccountNoAuth      = status.Error(codes.Code(ErrCode_UserNoAuth), "this account has no auth")
	ErrAccountNotVerified = status.Error(codes.Code(ErrCode_UserNoAuth), "this account not verified")

	// Other
	ErrDBRecordNotFound       = status.Error(codes.Code(ErrCode_RecordNotFound), "database record not found")
	ErrRedisRecordNotFound    = status.Error(codes.Code(ErrCode_RecordNotFound), "redis record not found")
	ErrRequestParamsError     = status.Error(codes.Code(ErrCode_RequestParamsError), "request param error")
	ErrNotImplemented         = status.Error(codes.Code(ErrCode_NotImplemented), "not implemented error")
	ErrDBNotImplemented       = status.Error(codes.Code(ErrCode_NotImplemented), "database not implemented")
	ErrServiceNotImplemented  = status.Error(codes.Code(ErrCode_NotImplemented), "service not implemented")
	ErrFunctionNotImplemented = status.Error(codes.Code(ErrCode_NotImplemented), "function not implemented")
	ErrCaseNotImplemented     = status.Error(codes.Code(ErrCode_NotImplemented), "function in this case not implemented")
	ErrNoPermissionToEdit     = status.Error(codes.Code(ErrCode_NoPermissionToEdit), "no permission to edit")
	ErrNoPermissionToDelete   = status.Error(codes.Code(ErrCode_NoPermissionToDelete), "no permission to delete")
	ErrServicePanicError      = status.Error(codes.Code(ErrCode_ServicePanicError), "service panic error")

	// Order
	ErrNotTradingTime     = status.Error(codes.Code(ErrCode_NotTradingTime), "not trading time")
	ErrFracShareBelowOne  = status.Error(codes.Code(ErrCode_FracShareBelowOne), "fractional share below one")
	ErrPositionNotEnough  = status.Error(codes.Code(ErrCode_PositionNotEnough), "position not enough")
	ErrLeverageNotTheSame = status.Error(codes.Code(ErrCode_LeverageNotTheSame), "leverage not the same")
	ErrNoQuote            = status.Error(codes.Code(ErrCode_NoQuote), "there is no quote")
	ErrNotEnabled         = status.Error(codes.Code(ErrCode_NotEnabled), "product not enabled")
	ErrOverTimeInterval   = status.Error(codes.Code(ErrCode_OverTimeInterval), "over time interval")
	ErrReadOnly           = status.Error(codes.Code(ErrCode_ReadOnly), "read only")
	// ErrStopLossPriceLTEGainProfitPrice = status.Error(codes.Code(ErrCode_StopLossPriceLTEGainProfitPrice), "stopLossPrice equals or less than gainProfitPrice")
	// ErrStopLossPriceGTEGainProfitPrice = status.Error(codes.Code(ErrCode_StopLossPriceGTEGainProfitPrice), "stopLossPrice equals or greater than gainProfitPrice")
	ErrMarketPriceLTEStopLossPrice       = status.Error(codes.Code(ErrCode_MarketPriceLTEStopLossPrice), "marketPrice equals or less than stopLossPrice")
	ErrMarketPriceGTEStopLossPrice       = status.Error(codes.Code(ErrCode_MarketPriceGTEStopLossPrice), "marketPrice equals or greater than stopLossPrice")
	ErrMarketPriceGTEGainProfictPrice    = status.Error(codes.Code(ErrCode_MarketPriceGTEGainProfitPrice), "marketPrice equals or greater than gainProfitPrice")
	ErrMarketPriceLTEGainProfictPrice    = status.Error(codes.Code(ErrCode_MarketPriceLTEGainProfitPrice), "marketPrice equals or less than gainProfitPrice")
	ErrOrderPriceLTEStopLossPrice        = status.Error(codes.Code(ErrCode_OrderPriceLTEStopLossPrice), "orderPrice equals or less than stopLossPrice")
	ErrOrderPriceGTEStopLossPrice        = status.Error(codes.Code(ErrCode_OrderPriceGTEStopLossPrice), "orderPrice equals or greater than stopLossPrice")
	ErrOrderPriceGTEGainProfictPrice     = status.Error(codes.Code(ErrCode_OrderPriceGTEGainProfitPrice), "orderPrice equals or greater than gainProfitPrice")
	ErrOrderPriceLTEGainProfictPrice     = status.Error(codes.Code(ErrCode_OrderPriceLTEGainProfitPrice), "orderPrice equals or less than gainProfitPrice")
	ErrClosePositionHaveStopLossFlag     = status.Error(codes.Code(ErrCode_ClosePositionHaveStopLossFlag), "close position can not have stop loss flag")
	ErrOrderHaveStopLossFlagWithPosition = status.Error(codes.Code(ErrCode_OrderHaveStopLossFlagWithPosition), "can not have stopLossFlag order holding position")
	ErrNotUnsettledOrder                 = status.Error(codes.Code(ErrCode_NotUnsettledOrder), "not unsettled order")
	ErrTradeAmountLessThanMinimum        = status.Error(codes.Code(ErrCode_TradeAmountLessThanMinimum), "trade amount less than minimum")

	// Payment
	ErrGetSecretError             = status.Error(codes.Code(ErrCode_GetSecretError), "get secret error")
	ErrDeposit                    = status.Error(codes.Code(ErrCode_DepositError), "deposit error")
	ErrDepositNotify              = status.Error(codes.Code(ErrCode_DepositNotificationError), "deposit notify error")
	ErrWithdrawBalanceNotEnough   = status.Error(codes.Code(ErrCode_WithdrawBalanceNotEnough), "withdraw balance not enough")
	ErrEditExchangeRate           = status.Error(codes.Code(ErrCode_EditExchangeRateError), "edit exchange rate error")
	ErrEditWithdraw               = status.Error(codes.Code(ErrCode_EditWithdrawError), "edit withdraw error")
	ErrAddExchangeRate            = status.Error(codes.Code(ErrCode_AddExchangeRateError), "add exchange rate error")
	ErrDeleteExchangeRate         = status.Error(codes.Code(ErrCode_DeleteExchangeRateError), "delete exchange rate error")
	ErrWithdrawLossTooMuch        = status.Error(codes.Code(ErrCode_WithdrawLossTooMuch), "withdraw loss too much")
	ErrWithdrawGuaranteeNotEnough = status.Error(codes.Code(ErrCode_WithdrawGuaranteeNotEnough), "withdraw guarantee not enough")
	ErrExceedMaximumDeposit       = status.Error(codes.Code(ErrCode_ExceedMaximumDeposit), "exceed maximum deposit amount")
	ErrLessMinimumDeposit         = status.Error(codes.Code(ErrCode_LessMinimumDeposit), "less than minimum deposit amount")

	// General
	ErrFolderNotFound             = status.Error(codes.Code(ErrCode_InvalidUploadArgument), "upload to not existed folder")
	ErrFilenameAlreadyExist       = status.Error(codes.Code(ErrCode_InvalidUploadArgument), "upload filename already exist")
	ErrEmptyFilename              = status.Error(codes.Code(ErrCode_InvalidUploadArgument), "upload filename empty")
	ErrInvalidDashboardArgument   = status.Error(codes.Code(ErrCode_InvalidDashboardArgument), "invalid dashboard argument")
	ErrInvalidDashboardSearchType = status.Error(codes.Code(ErrCode_InvalidDashboardArgument), "invalid dashboard search type")
	ErrInvalidDashboardDate       = status.Error(codes.Code(ErrCode_InvalidDashboardArgument), "invalid dashboard date")

	// Promotion
	ErrExceedMaximumReferralUsers   = status.Error(codes.Code(ErrCode_ExceedMaximumReferralUsers), "exceed maximum referral users")
	ErrExceedMaximumAttendance      = status.Error(codes.Code(ErrCode_ExceedMaximumAttendance), "exceed maximum maximum attendance")
	ErrUserPromotionReviewing       = status.Error(codes.Code(ErrCode_UserPromotionReviewing), "user promotion reviewing")
	ErrEventNotGoing                = status.Error(codes.Code(ErrCode_EventNotGoing), "event not going")
	ErrLessThanMinimumDepositAmount = status.Error(codes.Code(ErrCode_LessThanMinimumDepositAmount), "less than minimum deposit amount")
	ErrExceedEventIntervalDay       = status.Error(codes.Code(ErrCode_ExceedEventIntervalDays), "exceed interval days")
	ErrNotFirstDeposit              = status.Error(codes.Code(ErrCode_NotFirstDeposit), "not first deposit")
	ErrLessThanMinimumTradeAmount   = status.Error(codes.Code(ErrCode_LessThanMinimumTradeAmount), "less than minimum trade amount")
	ErrNotFirstTrade                = status.Error(codes.Code(ErrCode_NotFirstTrade), "not first trade")
	ErrAttendedEvent                = status.Error(codes.Code(ErrCode_AttendedEvent), "attended event")
	ErrEventNeedApply               = status.Error(codes.Code(ErrCode_EventNeedApply), "event needs apply")
	ErrBudgetNotEnough              = status.Error(codes.Code(ErrCode_BudgetNotEnough), "budget not enough")

	// media
	ErrMediaUnknown                = status.Error(codes.Unknown, "media service unknown error")
	ErrNoPostPermission            = status.Error(codes.Code(ErrCode_NoPostPermission), "no permission to post")
	ErrNoCommentPermission         = status.Error(codes.Code(ErrCode_NoCommentPermission), "no permission to comment")
	ErrPostNotFound                = status.Error(codes.Code(ErrCode_InvalidPostID), "post not found")
	ErrUnableToCommentPost         = status.Error(codes.Code(ErrCode_InvalidPostID), "unable to comment to a child comment")
	ErrCommentNotFound             = status.Error(codes.Code(ErrCode_InvalidCommentID), "comment not found")
	ErrUnableToCommentThisComment  = status.Error(codes.Code(ErrCode_InvalidCommentID), "unable to comment to this comment")
	ErrUnableToCommentChildComment = status.Error(codes.Code(ErrCode_InvalidCommentID), "unable to comment to a child comment")
)
