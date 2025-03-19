package mapper

import (
	"time"

	certificate "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/model"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
)

func ToCertificateDTO(cert certificate.Certificate) dtos.CertificateDTO {
	return dtos.CertificateDTO{
		ID:             cert.ID,
		EnrollmentID:   cert.EnrollmentID,
		IssuedAt:       cert.IssuedAt,
		CertificateURL: cert.CertificateURL,
		ExpiresAt:      toTimePointer(cert.ExpiresAt),
	}
}

func ToCertificate(certDTO dtos.CertificateDTO) certificate.Certificate {
	return certificate.Certificate{
		ID:             certDTO.ID,
		EnrollmentID:   certDTO.EnrollmentID,
		IssuedAt:       certDTO.IssuedAt,
		CertificateURL: certDTO.CertificateURL,
		ExpiresAt:      fromTimePointer(certDTO.ExpiresAt),
	}
}

func ToCertificateDTOs(certificates []certificate.Certificate) []dtos.CertificateDTO {
	var dtos []dtos.CertificateDTO
	for _, cert := range certificates {
		dtos = append(dtos, ToCertificateDTO(cert))
	}
	return dtos
}

func ToCertificates(certDTOs []dtos.CertificateDTO) []certificate.Certificate {
	var certificates []certificate.Certificate
	for _, certDTO := range certDTOs {
		certificates = append(certificates, ToCertificate(certDTO))
	}
	return certificates
}

func fromTimePointer(t *time.Time) time.Time {
	if t != nil {
		return *t
	}
	return time.Time{}
}

func toTimePointer(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t
}
