package controller_test

import (
	"context"
	"errors"
	"net/http/httptest"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	controller "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/controller"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
	appErr "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/error"
	mocks "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/tests/mock"
	"github.com/gofiber/fiber/v2"
)

// MockEnrollmentService es un mock del servicio de inscripciones
type MockEnrollmentService struct {
	mock.Mock
}

func (m *MockEnrollmentService) GetEnrollmentByID(ctx context.Context, id uuid.UUID) (*dtos.EnrollmentDTO, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*dtos.EnrollmentDTO), args.Error(1)
}

func (m *MockEnrollmentService) GetEnrollmentByUserAndCourse(ctx context.Context, userID, courseID uuid.UUID) (*dtos.EnrollmentDTO, error) {
	args := m.Called(ctx, userID, courseID)
	return args.Get(0).(*dtos.EnrollmentDTO), args.Error(1)
}

func (m *MockEnrollmentService) GetUserEnrollments(ctx context.Context, userID uuid.UUID, page, limit int) ([]dtos.EnrollmentDTO, int64, error) {
	args := m.Called(ctx, userID, page, limit)
	return args.Get(0).([]dtos.EnrollmentDTO), args.Get(1).(int64), args.Error(2)
}

func (m *MockEnrollmentService) GetCourseEnrollments(ctx context.Context, courseID uuid.UUID, page, limit int) ([]dtos.EnrollmentDTO, int64, error) {
	args := m.Called(ctx, courseID, page, limit)
	return args.Get(0).([]dtos.EnrollmentDTO), args.Get(1).(int64), args.Error(2)
}

func (m *MockEnrollmentService) EnrollUserInCourse(ctx context.Context, userID, courseID uuid.UUID) (*dtos.EnrollmentDTO, error) {
	args := m.Called(ctx, userID, courseID)
	return args.Get(0).(*dtos.EnrollmentDTO), args.Error(1)
}

func (m *MockEnrollmentService) CancelEnrollment(ctx context.Context, userID, enrollmentID uuid.UUID) error {
	args := m.Called(ctx, userID, enrollmentID)
	return args.Error(0)
}

func (m *MockEnrollmentService) MarkEnrollmentComplete(ctx context.Context, enrollmentID uuid.UUID) error {
	args := m.Called(ctx, enrollmentID)
	return args.Error(0)
}

func (m *MockEnrollmentService) IsUserEnrolledInCourse(ctx context.Context, userID, courseID uuid.UUID) bool {
	args := m.Called(ctx, userID, courseID)
	return args.Bool(0)
}

// TestSuite estructura para las pruebas
type EnrollmentQueryControllerTestSuite struct {
	suite.Suite
	app             *fiber.App
	mock            *MockEnrollmentService
	jwtMgr          *mocks.MockJWTManager
	mockProgress    *mocks.MockProgressService
	mockCertificate *mocks.MockCertificateService
}

func (suite *EnrollmentQueryControllerTestSuite) SetupTest() {
	suite.mock = new(MockEnrollmentService)
	suite.jwtMgr = new(mocks.MockJWTManager)
	suite.app = fiber.New()

	suite.mock = new(MockEnrollmentService)
	suite.mockProgress = new(mocks.MockProgressService)
	suite.mockCertificate = new(mocks.MockCertificateService)

}

func (suite *EnrollmentQueryControllerTestSuite) TestGetMyEnrollments_Success() {
	userID := uuid.New()
	expectedEnrollments := []dtos.EnrollmentDTO{
		{ID: uuid.New(), UserID: userID, CourseID: uuid.New()},
	}

	suite.mock.On("GetUserEnrollments", mock.Anything, userID, 1, 10).Return(expectedEnrollments, int64(1), nil)

	req := httptest.NewRequest("GET", "/enrollments/me", nil)
	req.Header.Set("Authorization", "Bearer valid_token")

	resp, _ := suite.app.Test(req, -1)
	defer resp.Body.Close()

	suite.Equal(fiber.StatusOK, resp.StatusCode)
}

func (suite *EnrollmentQueryControllerTestSuite) TestGetMyEnrollments_Unauthorized() {
	req := httptest.NewRequest("GET", "/enrollments/me", nil)
	req.Header.Set("Authorization", "Bearer invalid_token")

	suite.jwtMgr.On("VerifyToken", mock.Anything).Return(uuid.Nil, errors.New("invalid token"))

	resp, _ := suite.app.Test(req, -1)
	defer resp.Body.Close()

	suite.Equal(fiber.StatusUnauthorized, resp.StatusCode)
}

func (suite *EnrollmentQueryControllerTestSuite) TestGetEnrollmentByID_Success() {
	enrollmentID := uuid.New()
	expectedEnrollment := &dtos.EnrollmentDTO{ID: enrollmentID}

	suite.mock.On("GetEnrollmentByID", mock.Anything, enrollmentID).Return(expectedEnrollment, nil)

	req := httptest.NewRequest("GET", "/v1/api/enrollments/"+enrollmentID.String(), nil)
	req.Header.Set("Authorization", "Bearer valid_token")

	resp, _ := suite.app.Test(req, -1)
	defer resp.Body.Close()

	suite.Equal(fiber.StatusOK, resp.StatusCode)
}

func (suite *EnrollmentQueryControllerTestSuite) TestGetEnrollmentByID_NotFound() {
	enrollmentID := uuid.New()
	suite.mock.On("GetEnrollmentByID", mock.Anything, enrollmentID).Return((*dtos.EnrollmentDTO)(nil), appErr.ErrEnrollmentNotFoundDB)

	req := httptest.NewRequest("GET", "/v1/api/enrollments/"+enrollmentID.String(), nil)
	req.Header.Set("Authorization", "Bearer valid_token")

	resp, _ := suite.app.Test(req, -1)
	defer resp.Body.Close()

	suite.Equal(fiber.StatusNotFound, resp.StatusCode)
}

func (suite *EnrollmentQueryControllerTestSuite) TestGetCourseEnrollments_Success() {
	courseID := uuid.New()
	expectedEnrollments := []dtos.EnrollmentDTO{
		{ID: uuid.New(), CourseID: courseID},
	}

	suite.mock.On("GetCourseEnrollments", mock.Anything, courseID, 1, 10).Return(expectedEnrollments, int64(1), nil)

	req := httptest.NewRequest("GET", "/v1/api/courses/"+courseID.String()+"/enrollments", nil)
	req.Header.Set("Authorization", "Bearer valid_token")

	resp, _ := suite.app.Test(req, -1)
	defer resp.Body.Close()

	suite.Equal(fiber.StatusOK, resp.StatusCode)
}

func (suite *EnrollmentQueryControllerTestSuite) TestGetCourseEnrollments_Empty() {
	courseID := uuid.New()
	suite.mock.On("GetCourseEnrollments", mock.Anything, courseID, 1, 10).Return([]dtos.EnrollmentDTO{}, int64(0), nil)

	req := httptest.NewRequest("GET", "/v1/api/courses/"+courseID.String()+"/enrollments", nil)
	req.Header.Set("Authorization", "Bearer valid_token")

	resp, _ := suite.app.Test(req, -1)
	defer resp.Body.Close()

	suite.Equal(fiber.StatusOK, resp.StatusCode)
}

func (suite *EnrollmentQueryControllerTestSuite) TestGetCourseEnrollments_InvalidCourseID() {
	req := httptest.NewRequest("GET", "/v1/api/courses/invalid-uuid/enrollments", nil)
	req.Header.Set("Authorization", "Bearer valid_token")

	resp, _ := suite.app.Test(req, -1)
	defer resp.Body.Close()

	suite.Equal(fiber.StatusBadRequest, resp.StatusCode)
}

type EnrollmentCommandControllerTestSuite struct {
	suite.Suite
	app             *fiber.App
	mockEnrollment  *MockEnrollmentService
	mockProgress    *mocks.MockProgressService
	mockCertificate *mocks.MockCertificateService
}

func (suite *EnrollmentCommandControllerTestSuite) SetupTest() {
	suite.mockEnrollment = new(MockEnrollmentService)
	suite.mockProgress = new(mocks.MockProgressService)
	suite.mockCertificate = new(mocks.MockCertificateService)

	suite.app = fiber.New()

	controller := controller.NewEnrollmentComandController(
		suite.mockEnrollment,
		suite.mockCertificate,
		suite.mockProgress,
	)

	suite.app.Post("/v1/api/users/:user_id/courses/:course_id/enroll", controller.EnrollUserInCourse)
	suite.app.Delete("/v1/api/users/:user_id/enrollments/:enrollent_id", controller.CancellMyEnrollment)
	suite.app.Post("/v1/api/enrollments/:enrollent_id/complete", controller.CompleteMyCourse)
}

func (suite *EnrollmentCommandControllerTestSuite) TestCompleteMyCourse_Success() {
	enrollmentID := uuid.New()
	certificate := &dtos.CertificateDTO{ID: uuid.New(), EnrollmentID: enrollmentID}

	suite.mockEnrollment.On("MarkEnrollmentComplete", mock.Anything, enrollmentID).Return(nil)
	suite.mockCertificate.On("GenerateCertificate", mock.Anything, enrollmentID).Return(certificate, nil)

	req := httptest.NewRequest("POST", "/v1/api/enrollments/"+enrollmentID.String()+"/complete", nil)

	resp, _ := suite.app.Test(req, -1)
	defer resp.Body.Close()

	suite.Equal(fiber.StatusOK, resp.StatusCode)
}

func (suite *EnrollmentCommandControllerTestSuite) TestCompleteMyCourse_InvalidUUID() {
	req := httptest.NewRequest("POST", "/v1/api/enrollments/invalid-uuid/complete", nil)

	resp, _ := suite.app.Test(req, -1)
	defer resp.Body.Close()

	suite.Equal(fiber.StatusBadRequest, resp.StatusCode)
}

func (suite *EnrollmentCommandControllerTestSuite) TestCancellMyEnrollment_Success() {
	userID := uuid.New()
	enrollmentID := uuid.New()

	suite.mockEnrollment.On("CancelEnrollment", mock.Anything, userID, enrollmentID).Return(nil)

	req := httptest.NewRequest("DELETE", "/v1/api/users/"+userID.String()+"/enrollments/"+enrollmentID.String(), nil)

	resp, _ := suite.app.Test(req, -1)
	defer resp.Body.Close()

	suite.Equal(fiber.StatusOK, resp.StatusCode)
}

func (suite *EnrollmentCommandControllerTestSuite) TestCancellMyEnrollment_InvalidUUID() {
	req := httptest.NewRequest("DELETE", "/v1/api/users/invalid-uuid/enrollments/another-invalid-uuid", nil)

	resp, _ := suite.app.Test(req, -1)
	defer resp.Body.Close()

	suite.Equal(fiber.StatusBadRequest, resp.StatusCode)
}

func (suite *EnrollmentCommandControllerTestSuite) TestEnrollUserInCourse_Success() {
	userID := uuid.New()
	courseID := uuid.New()
	enrollment := &dtos.EnrollmentDTO{ID: uuid.New(), UserID: userID, CourseID: courseID}

	suite.mockEnrollment.On("EnrollUserInCourse", mock.Anything, userID, courseID).Return(enrollment, nil)
	suite.mockProgress.On("CreateCourseTrackRecord", mock.Anything, enrollment.ID).Return(nil)

	req := httptest.NewRequest("POST", "/v1/api/users/"+userID.String()+"/courses/"+courseID.String()+"/enroll", nil)

	resp, _ := suite.app.Test(req, -1)
	defer resp.Body.Close()

	suite.Equal(fiber.StatusCreated, resp.StatusCode)
}

func (suite *EnrollmentCommandControllerTestSuite) TestEnrollUserInCourse_AlreadyEnrolled() {
	userID := uuid.New()
	courseID := uuid.New()

	suite.mockEnrollment.On("IsUserEnrolledInCourse", mock.Anything, userID, courseID).Return(true)

	req := httptest.NewRequest("POST", "/v1/api/users/"+userID.String()+"/courses/"+courseID.String()+"/enroll", nil)

	resp, _ := suite.app.Test(req, -1)
	defer resp.Body.Close()

	suite.Equal(fiber.StatusConflict, resp.StatusCode)
}
