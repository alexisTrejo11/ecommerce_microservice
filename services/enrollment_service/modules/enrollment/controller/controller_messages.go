package controller

const (
	MsgUnauthorized                = "unauthorized"
	MsgEnrollmentNotFound          = "enrollment_not_found"
	MsgInvalidEnrollmentID         = "invalid_enrollment_id"
	MsgInvalidUserID               = "invalid_user_id"
	MsgInvalidCourseID             = "invalid_course_id"
	MsgUserEnrollmentRetrieved     = "User Enrollment Successfully Retrieved"
	MsgEnrollmentRetrieved         = "Enrollment Successfully Retrieved"
	MsgEnrollmentsRetrieved        = "Enrollments Successfully Retrieved"
	MsgNoEnrollmentsFoundForCourse = "No Enrollments Found for this Course"
	MsgCourseEnrollmentRetrieved   = "Course Enrollment Successfully Retrieved"

	MsgCourseCompleted     = "Course Successfully Completed"
	MsgEnrollmentCancelled = "Enrollment Successfully Cancelled"
	MsgUserEnrolled        = "User Successfully Enrolled"

	KeyGetMyEnrollments             = "get_my_enrollments"
	KeyGetEnrollmentByID            = "get_enrollments_by_id"
	KeyGetEnrollmentByUserAndCourse = "get_enrollment_by_user_and_course"
	KeyGetCourseEnrollments         = "get_course_enrollments"

	KeyCompleteCourse     = "complete_course"
	KeyCancelEnrollment   = "cancel_enrollment"
	KeyEnrollUserInCourse = "enroll_user_in_course"
)
