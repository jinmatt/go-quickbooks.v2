package consts

// Global constants
const (
	QBFaultType            = "AUTHENTICATION"
	ContextTypeXML         = "text/xml"
	QBAccountIncomeType    = "Income"
	QBItemServiceType      = "Service"
	QBSalesItemLineDetail  = "SalesItemLineDetail"
	QBPaymentIncomeTxnType = "Invoice"

	//QB fault types
	QBValidationFault           = "ValidationFault"
	QBSystemFault               = "SystemFault"
	QBAuthenticationFault       = "AuthenticationFault"
	QBAuthorizationFault        = "AuthorizationFault"
	QBAuthenticationFaultCode   = "110"
	QBAuthorizationFaultMessage = "Authorization has failed"

	//QB Authorization Code error types
	QBAuthorizationCodeFailure        = "AuthorizationFailure"
	QBAuthorizationCodeFailureCode    = "120"
	QBAuthorizationFailureCodeMessage = "Invalid grant"
)
