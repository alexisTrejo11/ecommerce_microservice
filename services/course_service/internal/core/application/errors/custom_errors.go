package customErrors

import "fmt"

type DomainError struct {
	Code    string
	Message string
	Err     error
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s, Details: %v", e.Code, e.Message, e.Err)
}

func NewDomainError(code string, message string, err error) *DomainError {
	return &DomainError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

var (
	ErrCourseNameRequired     = NewDomainError("COURSE_INVALID_INPUT", "Course domain: Course name is required", nil)
	ErrCourseInvalidLanguage  = NewDomainError("COURSE_INVALID_LANGUAGE", "Course domain: The provided language is not valid", nil)
	ErrCourseAlreadyPublished = NewDomainError("COURSE_ALREADY_PUBLISHED", "Course domain: The course has already been published", nil)

	ErrModuleNotFound           = NewDomainError("MODULE_NOT_FOUND", "Module domain: The module does not exist in the course", nil)
	ErrModuleTitleInvalid       = NewDomainError("MODULE_INVALID_TITLE", "Module domain: The module title must be between 3 and 100 characters", nil)
	ErrModuleOrderInvalid       = NewDomainError("MODULE_INVALID_ORDER", "Module domain: The module order must be a non-negative number", nil)
	ErrModuleMaxLessonsExceeded = NewDomainError("MODULE_MAX_LESSONS_EXCEEDED", "Module domain: A module cannot have more than 50 lessons", nil)

	ErrLessonTitleRequired        = NewDomainError("LESSON_INVALID_TITLE", "Lesson domain: The lesson title is required", nil)
	ErrLessonTitleTooLong         = NewDomainError("LESSON_INVALID_TITLE", "Lesson domain: The lesson title exceeds the maximum allowed length", nil)
	ErrLessonInvalidDuration      = NewDomainError("LESSON_INVALID_DURATION", "Lesson domain: The lesson duration is invalid", nil)
	ErrLessonMaxResourcesExceeded = NewDomainError("LESSON_RESOURCE_LIMIT_EXCEEDED", "Lesson domain: The maximum resource limit per lesson has been reached", nil)

	ErrResourceTitleRequired = NewDomainError("RESOURCE_TITLE_REQUIRED", "Resource domain: The resource title is required", nil)
	ErrResourceURLRequired   = NewDomainError("RESOURCE_URL_REQUIRED", "Resource domain: The resource URL is required", nil)
	ErrResourceInvalidType   = NewDomainError("RESOURCE_INVALID_TYPE", "Resource domain: The resource type is invalid", nil)

	// DB
	ErrNotFoundDB         = NewDomainError("NOT_FOUND", "The requested entity was not found", nil)
	ErrDB                 = NewDomainError("DATABASE_ERROR", "An error occurred while accessing the database", nil)
	ErrInvalidOperationDB = NewDomainError("INVALID_OPERATION", "The requested operation is invalid", nil)

	ErrCourseNotFoundDB   = NewDomainError("COURSE_NOT_FOUND", "The requested course was not found", nil)
	ErrModuleNotFoundDB   = NewDomainError("MODULE_NOT_FOUND", "The requested module was not found", nil)
	ErrLessonNotFoundDB   = NewDomainError("LESSON_NOT_FOUND", "The requested lesson was not found", nil)
	ErrResourceNotFoundDB = NewDomainError("RESOURCE_NOT_FOUND", "The requested resource was not found", nil)
	ErrLessonFetchErrorDB = NewDomainError("LESSON_FETCH_ERROR", "An error occurred while fetching lessons for the module", nil)
)
