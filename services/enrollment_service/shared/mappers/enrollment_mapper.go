package mapper

import (
	enrollment "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/model"
	progress "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/model"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
)

// ToEnrollmentDTO convierte un modelo Enrollment a EnrollmentDTO
func ToEnrollmentDTO(enrollment enrollment.Enrollment) dtos.EnrollmentDTO {
	var completedLessonsDTO []dtos.CompletedLessonDTO
	for _, lesson := range enrollment.CompletedLessons {
		completedLessonsDTO = append(completedLessonsDTO, ToCompletedLessonDTO(lesson))
	}

	return dtos.EnrollmentDTO{
		ID:               enrollment.ID,
		UserID:           enrollment.UserID,
		CourseID:         enrollment.CourseID,
		EnrollmentDate:   enrollment.EnrollmentDate,
		CompletionDate:   toTimePointer(enrollment.CompletionDate),
		CompletionStatus: enrollment.CompletionStatus,
		Progress:         enrollment.Progress,
		LastAccessedAt:   toTimePointer(enrollment.LastAccessedAt),
		CompletedLessons: completedLessonsDTO,
	}
}

// ToEnrollment convierte un EnrollmentDTO a modelo Enrollment
func ToEnrollment(enrollmentDTO dtos.EnrollmentDTO) enrollment.Enrollment {
	var completedLessons []progress.CompletedLesson
	/*

		for _, lessonDTO := range enrollmentDTO.CompletedLessons {
			completedLessons = append(completedLessons, ToCompletedLesson(lessonDTO))
		}
	*/

	return enrollment.Enrollment{
		ID:       enrollmentDTO.ID,
		UserID:   enrollmentDTO.UserID,
		CourseID: enrollmentDTO.CourseID,
		//EnrollmentDate:   enrollmentdtos.EnrollmentDate,
		CompletionDate:   fromTimePointer(enrollmentDTO.CompletionDate),
		CompletionStatus: enrollmentDTO.CompletionStatus,
		Progress:         enrollmentDTO.Progress,
		LastAccessedAt:   fromTimePointer(enrollmentDTO.LastAccessedAt),
		CompletedLessons: completedLessons,
	}
}

func ToEnrollmentDTOs(enrollments []enrollment.Enrollment) []dtos.EnrollmentDTO {
	dtos := make([]dtos.EnrollmentDTO, len(enrollments))
	for i, enrollment := range enrollments {
		dtos[i] = ToEnrollmentDTO(enrollment)
	}
	return dtos
}

func ToEnrollments(enrollmentDTOs []dtos.EnrollmentDTO) []enrollment.Enrollment {
	var enrollments []enrollment.Enrollment
	for _, enrollmentDTO := range enrollmentDTOs {
		enrollments = append(enrollments, ToEnrollment(enrollmentDTO))
	}
	return enrollments
}
