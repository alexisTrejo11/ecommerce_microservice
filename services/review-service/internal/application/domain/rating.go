package domain

import (
	"math"
	"time"

	"github.com/google/uuid"
)

type Review struct {
	id         uuid.UUID
	userID     uuid.UUID
	courseID   uuid.UUID
	rating     int
	comment    string
	createdAt  time.Time
	updatedAt  time.Time
	isApproved bool
}

func NewReview(userID, courseID uuid.UUID, rating int, comment string) *Review {
	return &Review{
		id:         uuid.New(),
		userID:     userID,
		courseID:   courseID,
		rating:     rating,
		comment:    comment,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
		isApproved: true,
	}
}

func CalculateRating(reviews []Review) float64 {
	numberOfReviews := len(reviews)
	if numberOfReviews == 0 {
		return 0
	}

	sum := 0
	for _, review := range reviews {
		sum += review.rating
	}

	average := float64(sum) / float64(numberOfReviews)
	rating := math.Round(average*100) / 100

	return rating
}

func (r *Review) SetRating(rating int) {
	r.rating = rating
	r.updatedAt = time.Now()
}

func (r *Review) SetComment(comment string) {
	r.comment = comment
	r.updatedAt = time.Now()
}

func (r *Review) Approve() {
	r.isApproved = true
	r.updatedAt = time.Now()
}

func (r *Review) Reject() {
	r.isApproved = false
	r.updatedAt = time.Now()
}

type User struct {
	Id          uuid.UUID
	Name        string
	Enrollments []Course
}

type Course struct {
	Id     uuid.UUID
	Name   string
	Rating int
}

func (r *Review) GetID() uuid.UUID        { return r.id }
func (r *Review) GetUserID() uuid.UUID    { return r.userID }
func (r *Review) GetCourseID() uuid.UUID  { return r.courseID }
func (r *Review) GetRating() int          { return r.rating }
func (r *Review) GetComment() string      { return r.comment }
func (r *Review) GetCreatedAt() time.Time { return r.createdAt }
func (r *Review) GetUpdatedAt() time.Time { return r.updatedAt }
func (r *Review) IsApproved() bool        { return r.isApproved }
