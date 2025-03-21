package controller

const (
	// Operation names
	opGetCourseProgress    = "get_my_course_progress"
	opMarkLessonComplete   = "mark_lesson_complete"
	opMarkLessonIncomplete = "mark_lesson_incomplete"

	// Error messages
	errUnauthorized    = "unauthorized"
	errInvalidLessonID = "invalid_lesson_id"

	// Success messages
	msgCourseProgressRetrieved = "User course progress successfully retrieved"
	msgLessonCompleted         = "Lesson successfully marked as completed"
	msgLessonIncompleted       = "Lesson successfully marked as incomplete"
)
