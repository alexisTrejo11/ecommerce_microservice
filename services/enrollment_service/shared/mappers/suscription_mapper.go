package mapper

import (
	suscription "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/suscription/model"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
)

func ToSubscriptionDTO(sub suscription.Subscription) dtos.SubscriptionDTO {
	return dtos.SubscriptionDTO{
		ID:        sub.GetID(),
		UserID:    sub.GetUserID(),
		PlanName:  sub.GetPlanName(),
		StartDate: sub.GetStartDate(),
		EndDate:   sub.GetEndDate(),
		PaymentID: sub.GetPaymentID(),
		Type:      sub.GetType(),
		Status:    sub.GetStatus(),
	}
}

func ToSubscription(subDTO dtos.SubscriptionInsertDTO) suscription.Subscription {
	return *suscription.NewSubscription(
		subDTO.UserID,
		subDTO.PlanName,
		subDTO.PaymentID,
		subDTO.Status,
		subDTO.Type,
	)
}
