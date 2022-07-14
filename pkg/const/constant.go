package constant

const (
	CoinTransactionStateCreated    = "created"
	CoinTransactionStateWait       = "wait"
	CoinTransactionStatePaying     = "paying"
	CoinTransactionStateSuccessful = "successful"
	CoinTransactionStateFail       = "fail"
	CoinTransactionStateRejected   = "rejected"

	WithdrawTypeBenefit            = "benefit"
	WithdrawTypeCommission         = "commission"
	WithdrawTypeUserPaymentBalance = "user-payment-balance"

	TransactionForCompatible      = "compatible"
	TransactionForWithdraw        = "withdraw"
	TransactionForWarmTransfer    = "warm-transfer"
	TransactionForCollecting      = "collecting"
	TransactionForPlatformBenefit = "platform-benefit"
	TransactionForUserBenefit     = "user-benefit"
)
