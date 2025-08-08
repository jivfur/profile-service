package testhelper

import (
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/jivfur/profile-service/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type ProfileOption func(*model.Profile)

// NewFakeProfile creates a default fake profile and applies options
var Genders = []string{
	"MALE", "FEMALE", "TRANS MAN", "TRANS WOMAN", "NON BINARY", "QUEER",
}
var SexualOrientation = []string{"GAY", "STRAIGHT", "BISEXUAL", "ASEXUAL", "DEMISEXUAL", "PANSEXUAL", "QUEER"}
var SexualPosition = []string{"TOP", "VERS TOP", "VERSATILE", "VERS BOTTOM", "BOTTOM", "SIDE"}

func NewFakeProfile(opts ...ProfileOption) *model.Profile {
	gofakeit.Seed(0)
	UUID := gofakeit.UUID()
	bio := gofakeit.Paragraph(1, 3, 5, "\n")
	city := gofakeit.City()
	photoURL := UUID + ".jpg"
	dob := time.Now().AddDate(-20, 0, 0)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	p := &model.Profile{
		ID:                UUID,
		Username:          gofakeit.Username(),
		Email:             gofakeit.Email(),
		PasswordHash:      string(hashedPassword),
		Name:              gofakeit.Name(),
		Bio:               &bio,
		Gender:            Genders[rand.Intn(len(Genders))],
		SexualOrientation: SexualOrientation[rand.Intn(len(SexualOrientation))],
		SexualPosition:    SexualPosition[rand.Intn(len(SexualPosition))],
		PhotoURL:          &photoURL, // Simulate a photo URL
		Location:          &city,
		DateOfBirth:       &dob, // 20 years ago
		Hobbies:           createListOfWords(3),
		Interests:         createListOfWords(5),
		EmailVerified:     gofakeit.Bool(),
		FaceVerified:      gofakeit.Bool(),
		VerificationScore: gofakeit.Float32Range(0, 100),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func createListOfWords(n int) *string {
	var hobbies []string
	for i := 0; i < n; i++ {
		hobbies = append(hobbies, gofakeit.Hobby()) // or any other faker func
	}
	joined := strings.Join(hobbies, ",")
	return &joined
}

// WithNo sets the named field to its zero value
func WithNo(fieldName string) ProfileOption {
	return func(p *model.Profile) {
		v := reflect.ValueOf(p).Elem()
		f := v.FieldByName(fieldName)
		if f.IsValid() && f.CanSet() {
			f.Set(reflect.Zero(f.Type()))
		}
	}
}

// WithCustom sets the named field to the given value
func WithCustom(fieldName string, value interface{}) ProfileOption {
	return func(p *model.Profile) {
		v := reflect.ValueOf(p).Elem()
		f := v.FieldByName(fieldName)
		if f.IsValid() && f.CanSet() {
			val := reflect.ValueOf(value)
			// Handle pointer fields by indirecting
			if f.Kind() == reflect.Ptr && val.Type().AssignableTo(f.Type().Elem()) {
				ptr := reflect.New(f.Type().Elem())
				ptr.Elem().Set(val)
				f.Set(ptr)
			} else if val.Type().AssignableTo(f.Type()) {
				f.Set(val)
			}
		}
	}
}
