package service

import (
	"context"
	"os"
	"time"

	"api_thienhang_com/pkg/entity"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/option"
)

type IService interface {
	// ** ACCOUNT & USER ***
	GetAccount(idToken string) (data *entity.Account, rError error)

	CheckUser(user *entity.Account) (*entity.User, error)
	UpdateUser(account *entity.Account, user *entity.User) (*entity.User, error)

	// EDUCATION
	AddEducation(account *entity.Account, educations []entity.Education) error
	UpdateEducation(account *entity.Account, education entity.Education) error
	DeleteEducation(account *entity.Account, education entity.Education) error

	//
	AddResidential(account *entity.Account, address *entity.Address) (*entity.Address, error)
	UpdateResidential(account *entity.Account, address *entity.Address) (*entity.Address, error)
	GetResidential(account *entity.Account, uuid string) (*entity.Address, error)

	InitResidential()

	GetProvinces(account *entity.Account) (interface{}, error)
	GetDistricts(account *entity.Account, province string) (interface{}, error)
	GetWards(account *entity.Account, province, district string) (interface{}, error)
	GetBuildings(account *entity.Account, province, district, ward string) (interface{}, error)

	GetProvince(account *entity.Account, province, district, ward, building string) (interface{}, error)
	// GetProvince(account *entity.Account, province, district, ward, building string) (interface{}, error)
	// GetProvince(account *entity.Account, province, district, ward, building string) (interface{}, error)
	// GetProvince(account *entity.Account, province, district, ward, building string) (interface{}, error)

	GetAddress(account *entity.Account, uuid string) (*entity.Address, error)
	// EXPERIENCE
	AddExperience(account *entity.Account, experiences []entity.Experience) error
	UpdateExperience(account *entity.Account, experience entity.Experience) error
	DeleteExperience(account *entity.Account, experience entity.Experience) error

	// SKILL
	AddSkill(account *entity.Account, skills []entity.Skill) error
	UpdateSkill(account *entity.Account, skill entity.Skill) error
	DeleteSkill(account *entity.Account, skill entity.Skill) error

	// REFERENCE
	AddReference(account *entity.Account, references []entity.Reference) error
	UpdateReference(account *entity.Account, reference entity.Reference) error
	DeleteReference(account *entity.Account, reference entity.Reference) error

	// AWARD
	AddAward(account *entity.Account, awards []entity.Award) error
	UpdateAward(account *entity.Account, award entity.Award) error
	DeleteAward(account *entity.Account, award entity.Award) error

	// QUALIFICATION
	AddQualification(account *entity.Account, qualifications []entity.Qualification) error
	UpdateQualification(account *entity.Account, qualification entity.Qualification) error
	DeleteQualification(account *entity.Account, qualification entity.Qualification) error

	// ************************************************************************
	// COURSE
	GetCourse(ID string) (result *entity.Course, err error)
	CreateCourse(account *entity.Account, c entity.Course) (entity.Course, error)
	UpdateCourse(c entity.Course) (entity.Course, error)

	// LESSION
	CreateLession(uuid string, p *entity.Lession) error
}

type Service struct {
	client *mongo.Client
	ctx    context.Context

	gg *auth.Client

	userDB        *mongo.Collection
	lessionDB     *mongo.Collection
	courseDB      *mongo.Collection
	bookDB        *mongo.Collection
	residentialDB *mongo.Collection

	wardlDB     *mongo.Collection
	districtlDB *mongo.Collection
	provincelDB *mongo.Collection
}

func Init() IService {
	// Connect to Mongo
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Error(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Error(err)
	}

	database := client.Database(os.Getenv("DB_NAME"))

	opt := option.WithCredentialsFile("../../cmd/config/gomatchingdotorg-firebase-adminsdk-jju15-228db9c08e.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Error("error initializing app: %v", err)
	}
	gg, err := app.Auth(context.Background())
	if err != nil {
		log.Error(err)
	}

	//
	userDB := database.Collection("user")

	lessionDB := database.Collection("lession")
	courseDB := database.Collection("course")
	bookDB := database.Collection("book")
	resDB := database.Collection("address")
	wardDB := database.Collection("ward")
	provinceDB := database.Collection("province")
	districtDB := database.Collection("district")

	return &Service{
		client,
		ctx,
		gg,
		userDB,
		lessionDB,
		courseDB,
		bookDB,
		resDB,
		wardDB,
		districtDB,
		provinceDB,
	}
}

func (s *Service) Destroy() {
	defer s.client.Disconnect(s.ctx)
}
