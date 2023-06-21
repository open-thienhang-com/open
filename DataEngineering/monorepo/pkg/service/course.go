package service

import (
	"context"
	"fmt"
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

func (s *Service) CreateLession(uuid string, l *entity.Lession) error {
	courseUUID, err := primitive.ObjectIDFromHex(uuid)
	if err != nil {
		return err
	}
	//
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
	//
	session, err := s.client.StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(context.Background())

	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		if err = session.StartTransaction(txnOpts); err != nil {
			return err
		}
		lr, err := s.lessionDB.InsertOne(
			sessionContext, l,
		)
		if err != nil {
			return err
		}
		id := lr.InsertedID
		l.ID = id.(primitive.ObjectID)
		filter := bson.M{"_id": courseUUID}
		log.Error(l)
		result, err := s.courseDB.UpdateOne(sessionContext,
			filter,
			bson.M{
				"$push": bson.M{
					"lessions": bson.M{"$each": []entity.Lession{*l}},
				},
			},
		)
		if err != nil {
			return err
		}
		if err = session.CommitTransaction(sessionContext); err != nil {
			return err
		}
		fmt.Println(result)
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

func (s *Service) CreateCourse(account *entity.Account, c entity.Course) (entity.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := s.courseDB.InsertOne(ctx, c)
	id := result.InsertedID
	c.ID = id.(primitive.ObjectID)
	return c, err
}

func (s *Service) UpdateCourse(c entity.Course) (entity.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": c.ID}
	r, err := s.courseDB.UpdateOne(ctx, filter, bson.M{
		"$set": c,
	})
	log.Error(r)
	log.Error(err)
	return c, err
}

func (s *Service) GetCourse(uuid string) (result *entity.Course, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objectId, err := primitive.ObjectIDFromHex(uuid)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	filter := bson.M{"_id": objectId}
	err = s.courseDB.FindOne(ctx, filter).Decode(&result)
	log.Error(result)
	return result, err
}

func (s *Service) GetCourses() ([]entity.Course, error) {
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// _, err := us.data.UpdateByID(ctx, c.ID, c)
	return nil, nil
}
func (s *Service) GetLessions(sToken string) []entity.Lession {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var Lessions []entity.Lession

	cursor, _ := s.lessionDB.Find(ctx, bson.M{})
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var Lession entity.Lession
		cursor.Decode(&Lession)
		Lessions = append(Lessions, Lession)
	}

	return Lessions
}
