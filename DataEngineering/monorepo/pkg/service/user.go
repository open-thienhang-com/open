package service

import (
	"context"
	"errors"
	"time"

	"api_thienhang_com/pkg/entity"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func (s *Service) GetAccount(sToken string) (data *entity.Account, rError error) {
	info, err := s.gg.VerifyIDToken(context.Background(), sToken)
	if err != nil {
		return nil, err
	}
	log.Error(info)
	email := info.Claims["email"]
	phone := info.Claims["phone_number"]
	if email == nil {
		email = ""
	}

	if phone == nil {
		phone = ""
	}

	if email == nil && phone == nil {
		return nil, errors.New("opps. Can not sync")
	}
	return &entity.Account{
		UID:         info.UID,
		Email:       email.(string),
		PhoneNumber: phone.(string),
	}, nil
}

func (s *Service) CheckUser(acc *entity.Account) (u *entity.User, err error) {
	err = s.userDB.FindOne(context.Background(), bson.M{
		"email": acc.Email,
	}).Decode(&u)

	if err == nil {
		return u, nil
	}

	u = &entity.User{
		Email:       acc.Email,
		PhoneNumber: acc.PhoneNumber,
		LivesIn:     []primitive.ObjectID{},
		// Educations:     []entity.Education{},
		// Experiences:    []entity.Experience{},
		// Skills:         []entity.Skill{},
		// References:     []entity.Reference{},
		// Awards:         []entity.Award{},
		// Qualifications: []entity.Qualification{},
		// Pages:          []primitive.ObjectID{},
		// Courses:        []primitive.ObjectID{},
		Updated_at: time.Now(),
		Created_at: time.Now(),
	}
	result, err := s.userDB.InsertOne(context.Background(), u)
	if err != nil {
		return nil, err
	}
	id := result.InsertedID
	u.ID = id.(primitive.ObjectID)
	return u, nil
}

func (s *Service) UpdateUser(account *entity.Account, user *entity.User) (userNew *entity.User, err error) {
	filter := bson.M{"email": account.Email}
	_, err = s.userDB.UpdateOne(
		context.Background(),
		filter,
		bson.M{
			"$set": bson.M{
				"phone":     user.PhoneNumber,
				"firstname": user.FirstName,
				"lastname":  user.LastName,
				"about":     user.About,
				"address":   user.Addresses,
				"livesin":   user.LivesIn,
				// "occupation": user.Occupation,
				"status":     user.Status,
				"gender":     user.Gender,
				"website":    user.Website,
				"dob":        user.DOB,
				"updated_at": time.Now(),
			},
		},
	)
	if err != nil {
		return nil, err
	}

	err = s.userDB.FindOne(context.Background(), bson.M{
		"email": account.Email,
	}).Decode(&userNew)

	if err != nil {
		return nil, err
	}

	return userNew, nil
}

// EDUCATION
func (s *Service) AddEducation(account *entity.Account, educations []entity.Education) error {
	if len(educations) == 0 {
		return errors.New("empty array")
	}
	for i := 0; i < len(educations); i++ {
		if educations[i].ID == primitive.NilObjectID {
			educations[i].ID = primitive.NewObjectID()
		}
	}

	log.Error(educations)
	// *** TRANSACTION ****
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
	//
	session, err := s.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	//
	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		if err = session.StartTransaction(txnOpts); err != nil {
			log.Error(err)
			return err
		}

		filter := bson.M{
			"email": account.Email,
		}
		log.Error(account.Email)
		log.Error(educations)
		opts := options.Update().SetUpsert(true)
		result, err := s.userDB.UpdateOne(sessionContext,
			filter,
			bson.M{
				"$addToSet": bson.M{
					"educations": bson.M{"$each": educations},
				},
			},
			opts,
		)
		if err != nil {
			log.Error(err)
			return err
		}
		if err = session.CommitTransaction(sessionContext); err != nil {
			log.Error(err)
			return err
		}
		log.Error(result)
		return nil
	})
	if err != nil {
		if abortErr := session.AbortTransaction(context.Background()); abortErr != nil {
			panic(abortErr)
		}
		panic(err)
	}
	return nil
}
func (s *Service) UpdateEducation(account *entity.Account, education entity.Education) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "email", Value: account.Email}}

	arrayFilters := options.ArrayFilters{Filters: bson.A{bson.M{"x._id": education.ID}}}
	upsert := true
	opts := options.UpdateOptions{
		ArrayFilters: &arrayFilters,
		Upsert:       &upsert,
	}

	update := bson.M{
		"$set": bson.M{
			"educations.$[x].from":       education.From,
			"educations.$[x].to":         education.To,
			"educations.$[x].title":      education.Description,
			"educations.$[x].major":      education.Major,
			"educations.$[x].grade":      education.Grade,
			"educations.$[x].updated_at": time.Now(),
		},
	}

	result, err := s.userDB.UpdateOne(
		ctx,
		filter,
		update,
		&opts,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Error(result)
	return nil
}
func (s *Service) DeleteEducation(account *entity.Account, education entity.Education) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "email", Value: account.Email}}

	update := bson.M{
		"$pull": bson.M{
			"educations": bson.M{"_id": education.ID},
		},
	}

	log.Error(education)
	result, err := s.userDB.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Error(result)
	return nil
}

// EXPERIENCE
func (s *Service) AddExperience(account *entity.Account, experiences []entity.Experience) error {
	if len(experiences) == 0 {
		return errors.New("empty array")
	}
	for i := 0; i < len(experiences); i++ {
		if experiences[i].ID == primitive.NilObjectID {
			experiences[i].ID = primitive.NewObjectID()
		}
	}

	// *** TRANSACTION ****
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
	//
	session, err := s.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	//
	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		if err = session.StartTransaction(txnOpts); err != nil {
			log.Error(err)
			return err
		}

		filter := bson.M{
			"email": account.Email,
		}
		opts := options.Update().SetUpsert(true)
		result, err := s.userDB.UpdateOne(sessionContext,
			filter,
			bson.M{
				"$addToSet": bson.M{
					"experiences": bson.M{"$each": experiences},
				},
			},
			opts,
		)
		if err != nil {
			log.Error(err)
			return err
		}
		if err = session.CommitTransaction(sessionContext); err != nil {
			log.Error(err)
			return err
		}
		log.Error(result)
		return nil
	})
	if err != nil {
		if abortErr := session.AbortTransaction(context.Background()); abortErr != nil {
			panic(abortErr)
		}
		panic(err)
	}
	return nil
}
func (s *Service) UpdateExperience(account *entity.Account, experience entity.Experience) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "email", Value: account.Email}}

	arrayFilters := options.ArrayFilters{Filters: bson.A{bson.M{"x._id": experience.ID}}}
	upsert := true
	opts := options.UpdateOptions{
		ArrayFilters: &arrayFilters,
		Upsert:       &upsert,
	}

	update := bson.M{
		"$set": bson.M{
			"experiences.$[x].from":           experience.From,
			"experiences.$[x].to":             experience.To,
			"experiences.$[x].title":          experience.Title,
			"experiences.$[x].company":        experience.Company,
			"experiences.$[x].responsibility": experience.Responsibility,
			"experiences.$[x].updated_at":     time.Now(),
		},
	}

	result, err := s.userDB.UpdateOne(
		ctx,
		filter,
		update,
		&opts,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Error(result)
	return nil
}
func (s *Service) DeleteExperience(account *entity.Account, experience entity.Experience) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "email", Value: account.Email}}

	update := bson.M{
		"$pull": bson.M{
			"experiences": bson.M{"_id": experience.ID},
		},
	}
	result, err := s.userDB.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Error(result)
	return nil
}

// SKILL
func (s *Service) AddSkill(account *entity.Account, skills []entity.Skill) error {
	if len(skills) == 0 {
		return errors.New("empty array")
	}
	for i := 0; i < len(skills); i++ {
		if skills[i].ID == primitive.NilObjectID {
			skills[i].ID = primitive.NewObjectID()
		}
	}

	// *** TRANSACTION ****
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
	//
	session, err := s.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	//
	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		if err = session.StartTransaction(txnOpts); err != nil {
			log.Error(err)
			return err
		}

		filter := bson.M{
			"email": account.Email,
		}
		opts := options.Update().SetUpsert(true)
		result, err := s.userDB.UpdateOne(sessionContext,
			filter,
			bson.M{
				"$addToSet": bson.M{
					"skills": bson.M{"$each": skills},
				},
			},
			opts,
		)
		if err != nil {
			log.Error(err)
			return err
		}
		if err = session.CommitTransaction(sessionContext); err != nil {
			log.Error(err)
			return err
		}
		log.Error(result)
		return nil
	})
	if err != nil {
		if abortErr := session.AbortTransaction(context.Background()); abortErr != nil {
			panic(abortErr)
		}
		panic(err)
	}
	return nil
}
func (s *Service) UpdateSkill(account *entity.Account, skill entity.Skill) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "email", Value: account.Email}}

	arrayFilters := options.ArrayFilters{Filters: bson.A{bson.M{"x._id": skill.ID}}}
	upsert := true
	opts := options.UpdateOptions{
		ArrayFilters: &arrayFilters,
		Upsert:       &upsert,
	}

	update := bson.M{
		"$set": bson.M{
			"skills.$[x].name":       skill.Name,
			"skills.$[x].level":      skill.Level,
			"skills.$[x].updated_at": time.Now(),
		},
	}

	result, err := s.userDB.UpdateOne(
		ctx,
		filter,
		update,
		&opts,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Error(result)
	return nil
}
func (s *Service) DeleteSkill(account *entity.Account, skill entity.Skill) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "email", Value: account.Email}}

	update := bson.M{
		"$pull": bson.M{
			"skills": bson.M{"_id": skill.ID},
		},
	}
	result, err := s.userDB.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Error(result)
	return nil
}

// REFERENCE
func (s *Service) AddReference(account *entity.Account, references []entity.Reference) error {
	if len(references) == 0 {
		return errors.New("empty array")
	}
	for i := 0; i < len(references); i++ {
		if references[i].ID == primitive.NilObjectID {
			references[i].ID = primitive.NewObjectID()
		}
	}

	// *** TRANSACTION ****
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
	//
	session, err := s.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	//
	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		if err = session.StartTransaction(txnOpts); err != nil {
			log.Error(err)
			return err
		}

		filter := bson.M{
			"email": account.Email,
		}
		opts := options.Update().SetUpsert(true)
		result, err := s.userDB.UpdateOne(sessionContext,
			filter,
			bson.M{
				"$addToSet": bson.M{
					"references": bson.M{"$each": references},
				},
			},
			opts,
		)
		if err != nil {
			log.Error(err)
			return err
		}
		if err = session.CommitTransaction(sessionContext); err != nil {
			log.Error(err)
			return err
		}
		log.Error(result)
		return nil
	})
	if err != nil {
		if abortErr := session.AbortTransaction(context.Background()); abortErr != nil {
			panic(abortErr)
		}
		panic(err)
	}
	return nil
}
func (s *Service) UpdateReference(account *entity.Account, reference entity.Reference) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "email", Value: account.Email}}

	arrayFilters := options.ArrayFilters{Filters: bson.A{bson.M{"x._id": reference.ID}}}
	upsert := true
	opts := options.UpdateOptions{
		ArrayFilters: &arrayFilters,
		Upsert:       &upsert,
	}

	update := bson.M{
		"$set": bson.M{
			"references.$[x].name":       reference.Name,
			"references.$[x].email":      reference.Email,
			"references.$[x].updated_at": reference.Updated_at,
		},
	}

	result, err := s.userDB.UpdateOne(
		ctx,
		filter,
		update,
		&opts,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Error(result)
	return nil
}
func (s *Service) DeleteReference(account *entity.Account, reference entity.Reference) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "email", Value: account.Email}}

	update := bson.M{
		"$pull": bson.M{
			"references": bson.M{"_id": reference.ID},
		},
	}
	result, err := s.userDB.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Error(result)
	return nil
}

// AWARD
func (s *Service) AddAward(account *entity.Account, awards []entity.Award) error {
	if len(awards) == 0 {
		return errors.New("empty array")
	}
	for i := 0; i < len(awards); i++ {
		if awards[i].ID == primitive.NilObjectID {
			awards[i].ID = primitive.NewObjectID()
		}
	}

	// *** TRANSACTION ****
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
	//
	session, err := s.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	//
	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		if err = session.StartTransaction(txnOpts); err != nil {
			log.Error(err)
			return err
		}

		filter := bson.M{
			"email": account.Email,
		}
		opts := options.Update().SetUpsert(true)
		result, err := s.userDB.UpdateOne(sessionContext,
			filter,
			bson.M{
				"$addToSet": bson.M{
					"awards": bson.M{"$each": awards},
				},
			},
			opts,
		)
		if err != nil {
			log.Error(err)
			return err
		}
		if err = session.CommitTransaction(sessionContext); err != nil {
			log.Error(err)
			return err
		}
		log.Error(result)
		return nil
	})
	if err != nil {
		if abortErr := session.AbortTransaction(context.Background()); abortErr != nil {
			panic(abortErr)
		}
		panic(err)
	}
	return nil
}
func (s *Service) UpdateAward(account *entity.Account, award entity.Award) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "email", Value: account.Email}}

	arrayFilters := options.ArrayFilters{Filters: bson.A{bson.M{"x._id": award.ID}}}
	upsert := true
	opts := options.UpdateOptions{
		ArrayFilters: &arrayFilters,
		Upsert:       &upsert,
	}

	update := bson.M{
		"$set": bson.M{
			"awards.$[x].name": award.Name,
		},
	}

	result, err := s.userDB.UpdateOne(
		ctx,
		filter,
		update,
		&opts,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Error(result)
	return nil
}
func (s *Service) DeleteAward(account *entity.Account, award entity.Award) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "email", Value: account.Email}}

	update := bson.M{
		"$pull": bson.M{
			"awards": bson.M{"_id": award.ID},
		},
	}
	result, err := s.userDB.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Error(result)
	return nil
}

// QUALIFICATION
func (s *Service) AddQualification(account *entity.Account, qualifications []entity.Qualification) error {
	if len(qualifications) == 0 {
		return errors.New("empty array")
	}
	for i := 0; i < len(qualifications); i++ {
		if qualifications[i].ID == primitive.NilObjectID {
			qualifications[i].ID = primitive.NewObjectID()
		}
	}

	// *** TRANSACTION ****
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
	//
	session, err := s.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	//
	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		if err = session.StartTransaction(txnOpts); err != nil {
			log.Error(err)
			return err
		}

		filter := bson.M{
			"email": account.Email,
		}
		opts := options.Update().SetUpsert(true)
		result, err := s.userDB.UpdateOne(sessionContext,
			filter,
			bson.M{
				"$addToSet": bson.M{
					"qualifications": bson.M{"$each": qualifications},
				},
			},
			opts,
		)
		if err != nil {
			log.Error(err)
			return err
		}
		if err = session.CommitTransaction(sessionContext); err != nil {
			log.Error(err)
			return err
		}
		log.Error(result)
		return nil
	})
	if err != nil {
		if abortErr := session.AbortTransaction(context.Background()); abortErr != nil {
			panic(abortErr)
		}
		panic(err)
	}
	return nil
}
func (s *Service) UpdateQualification(account *entity.Account, qualification entity.Qualification) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "email", Value: account.Email}}

	arrayFilters := options.ArrayFilters{Filters: bson.A{bson.M{"x._id": qualification.ID}}}
	upsert := true
	opts := options.UpdateOptions{
		ArrayFilters: &arrayFilters,
		Upsert:       &upsert,
	}

	update := bson.M{
		"$set": bson.M{
			"qualifications.$[x].name": qualification.Name,
		},
	}

	result, err := s.userDB.UpdateOne(
		ctx,
		filter,
		update,
		&opts,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Error(result)
	return nil
}
func (s *Service) DeleteQualification(account *entity.Account, qualification entity.Qualification) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "email", Value: account.Email}}

	update := bson.M{
		"$pull": bson.M{
			"qualifications": bson.M{"_id": qualification.ID},
		},
	}
	result, err := s.userDB.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Error(result)
	return nil
}
