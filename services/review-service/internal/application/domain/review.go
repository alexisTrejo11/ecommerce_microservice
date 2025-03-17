package domain

import (
	"math"
	"time"

	"github.com/google/uuid"
)

var MAX_COMENT_CHAR_LENGTH = 1000

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

func NewReview(userID, courseID uuid.UUID, rating int, comment string) (*Review, error) {
	review := &Review{
		id:         uuid.New(),
		userID:     userID,
		courseID:   courseID,
		rating:     rating,
		comment:    comment,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
		isApproved: true,
	}

	if err := review.validateComment(); err != nil {
		return nil, err
	}
	if err := review.validateRating(); err != nil {
		return nil, err
	}

	return review, nil
}

func NewReviewFromDTO(
	id uuid.UUID,
	userID uuid.UUID,
	courseID uuid.UUID,
	rating int,
	comment string,
	createdAt time.Time,
	updatedAt time.Time,
	isApproved bool,
) *Review {
	return &Review{
		id:         id,
		userID:     userID,
		courseID:   courseID,
		rating:     rating,
		comment:    comment,
		createdAt:  createdAt,
		updatedAt:  updatedAt,
		isApproved: isApproved,
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

func (r *Review) Update(id, userID, courseID uuid.UUID, rating int, comment string) error {
	r.id = id
	r.userID = userID
	r.rating = rating
	r.courseID = courseID
	r.comment = comment

	if err := r.validateComment(); err != nil {
		return err
	}
	if err := r.validateRating(); err != nil {
		return err
	}

	return nil
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

func (r *Review) validateRating() error {
	if r.rating < 1 || r.rating > 5 {
		return ErrRatingValue
	}
	return nil
}

func (r *Review) validateComment() error {
	if len(r.comment) > MAX_COMENT_CHAR_LENGTH {
		return ErrCommentLenght
	}
	return nil
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
