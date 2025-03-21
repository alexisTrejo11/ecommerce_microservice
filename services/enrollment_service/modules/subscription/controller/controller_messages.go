package controller

const (
	MsgInvalidSubscriptionID = "invalid_subscription_id"
	MsgInvalidParam          = "invalid_param"
	MsgInvalidRequestBody    = "invalid_request_body"
	MsgInvalidRequestData    = "invalid_request_data"
	MsgUnauthorized          = "unauthorized"
	MsgSubscriptionDeleted   = "Subscription Successfully Deleted."
	MsgSubscriptionUpdated   = "Subscription Successfully Updated."
	MsgSubscriptionCreated   = "Subscription Successfully Created"
	MsgSubscriptionRetrieved = "Subscription Successfully Retrieved"
	MsgSubscriptionCancelled = "Subscription Successfully Cancelled."

	KeyDeleteSubscription       = "delete_subscription"
	KeyChangeMySubscriptionType = "change_my_subscription_type"
	KeyCreateSubscription       = "create_subscription"
	KeyGetMySubscription        = "get_my_subscription"
	KeyCancelMySubscription     = "cancel_my_subscription"
)
