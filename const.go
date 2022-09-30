package models

type InterestCode int32

const (
	InterestCode_Wallet         InterestCode = 1
	InterestCode_LeverageMargin InterestCode = 2
	InterestCode_ForexOvernight InterestCode = 3
)

type ProductType int

const (
	ProductFutures ProductType = 1
	ProductStock   ProductType = 2
	ProductCrypto  ProductType = 3
	ProductForex   ProductType = 4
)

type WalletRecordAction string

const (
	WalletRecordInterest       WalletRecordAction = "interest"
	WalletRecordFee            WalletRecordAction = "fee"
	WalletRecordOrderNew       WalletRecordAction = "order_new"
	WalletRecordOrderClose     WalletRecordAction = "order_close"
	WalletRecordStockDividends WalletRecordAction = "stock_dividends"
)

type StockDividendsStatus string

const (
	StockDividendsSettled      StockDividendsStatus = "settled"
	StockDividendsUnsettled    StockDividendsStatus = "unsettled"
	StockDividendsInsufficient StockDividendsStatus = "insufficient"
	StockDividendsFailed       StockDividendsStatus = "failed"
)

type RiskNotifyLevel int32

const (
	RiskNotifyLevelWarning     RiskNotifyLevel = 1
	RiskNotifyLevelReadyToStop RiskNotifyLevel = 2
	RiskNotifyLevelStopLoss    RiskNotifyLevel = 3
)

type PaymentAction string

const (
	PaymentActionWithdraw PaymentAction = "payment_withdraw"
	PaymentActionDeposit  PaymentAction = "payment_deposit"
)

type VerifyStatus string

const (
	VerifyStatusUnverified VerifyStatus = "unverified"
	VerifyStatusPending    VerifyStatus = "pending"
	VerifyStatusVerified   VerifyStatus = "verified"
	VerifyStatusFailed     VerifyStatus = "failed"
)

var VerifyStatusEnum = map[string]int32{
	"unverified": 1,
	"pending":    2,
	"verified":   3,
	"failed":     4,
}

type RoleCode string

const (
	RoleCodeMember RoleCode = "member"
)

type Status string

const (
	StatusEnabled Status = "enabled"
	StatusDerived Status = "derived"
)

const (
	RiskLevel = "RiskLevel"
)

type StopLossTakeProfit int32

const (
	StopLossTakeProfitStopLoss   StopLossTakeProfit = 1
	StopLossTakeProfitTakeProfit StopLossTakeProfit = 2
)

type DashboardSearchType int32

const (
	DashboardSearchTypeNone DashboardSearchType = iota
	DashboardSearchTypeDaily
	DashboardSearchTypeWeekly
	DashboardSearchTypeMonthly
	DashboardSearchTypeUnlimited
)
