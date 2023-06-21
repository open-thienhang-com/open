package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"api_thienhang_com/pkg/entity"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Service) InitResidential() {
	// files, err := ioutil.ReadDir("raw")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, file := range files {
	// 	fmt.Println(file.Name(), file.IsDir())
	// 	file := utils.LoadJsonFile("raw/" + file.Name())
	// 	// logrus.Info(string(file))
	// 	var data entity.Province

	// 	_ = json.Unmarshal(file, &data)
	// 	// LOAD PROVINCE
	// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	// 	defer cancel()
	// 	// transformData := make([]interface{}, len(data))
	// 	for _, v := range data.Districts {
	// 		v.ID = primitive.NewObjectID()
	// 		for _, m := range v.Wards {
	// 			m.ID = primitive.NewObjectID()

	// 		}
	// 	}
	// 	s.provincelDB.InsertOne(ctx, data)
	// }

	return
}

func (s *Service) AddResidential(account *entity.Account, c *entity.Address) (*entity.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := s.residentialDB.InsertOne(ctx, c)
	id := result.InsertedID
	c.ID = id.(primitive.ObjectID)
	return c, err
}

func (s *Service) UpdateResidential(account *entity.Account, c *entity.Address) (*entity.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": c.ID}
	r, err := s.residentialDB.UpdateOne(ctx, filter, bson.M{
		"$set": c,
	})
	log.Error(r)
	log.Error(err)
	return c, err
}

func (s *Service) GetResidential(account *entity.Account, uuid string) (result *entity.Address, err error) {
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

func (s *Service) GetWard(account *entity.Account, uuid string) (*entity.Ward, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objectId, err := primitive.ObjectIDFromHex(uuid)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	filter := bson.M{"_id": objectId}
	var result entity.Ward
	err = s.wardlDB.FindOne(ctx, filter).Decode(&result)
	return &result, err
}

func (s *Service) GetDistrict(account *entity.Account, uuid string) (*entity.District, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objectId, err := primitive.ObjectIDFromHex(uuid)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	filter := bson.M{"_id": objectId}
	var result entity.District
	err = s.districtlDB.FindOne(ctx, filter).Decode(&result)
	log.Error(result)
	return &result, err
}
func (s *Service) GetProvinces(account *entity.Account) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	projection := bson.D{
		{"code", 1},
		{"district.name", 1},
		{"district.ward.name", 1},
	}
	filter := bson.D{}

	cursor, err := s.provincelDB.Find(ctx, filter, options.Find().SetProjection(projection))
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		return nil, errors.New("Không tìm thấy địa chỉ")
	} else if err != nil {
		log.Error(err)
		return nil, errors.New("Lỗi không xác định")
	}
	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	return &episodes, err
}

func (s *Service) GetDistricts(account *entity.Account, province string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	projection := bson.D{
		{"code", 1},
		{"district.name", 1},
		{"district.ward.name", 1},
	}
	filter := bson.D{
		{"code", province},
	}

	cursor, err := s.provincelDB.Find(ctx, filter, options.Find().SetProjection(projection))
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		return nil, errors.New("Không tìm thấy địa chỉ")
	} else if err != nil {
		log.Error(err)
		return nil, errors.New("Lỗi không xác định")
	}
	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	return &episodes, err
}

func (s *Service) GetWards(account *entity.Account, province, district string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	projection := bson.D{
		{"code", 1},
		{"district.name.$", 1},
		{"district.ward.name", 1},
	}
	filter := bson.D{
		{"code", province},
		{"district",
			bson.D{
				{"$elemMatch", bson.D{
					{"name", district},
				}},
			},
		},
	}

	cursor, err := s.provincelDB.Find(ctx, filter, options.Find().SetProjection(projection))
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		return nil, errors.New("Không tìm thấy địa chỉ")
	} else if err != nil {
		log.Error(err)
		return nil, errors.New("Lỗi không xác định")
	}
	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	return &episodes, err
}

func (s *Service) GetBuildings(account *entity.Account, province, district, ward string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Warn("XXXXXX")
	log.Warn(province, district, ward)
	projection := bson.D{
		{"code", 1},
		// {"district.name", 1},
		{"district.ward.name.$", 1},
		// {"district.ward.building", 1},
		// {"district.ward.name", bson.D{
		// 	{"$slice", -1},
		// }},
	}
	filter := bson.D{
		{"code", province},
		{"district.name", district},
		{"district.ward.name", ward},
	}

	cursor, err := s.provincelDB.Find(ctx, filter, options.Find().SetProjection(projection))
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		return nil, errors.New("không tìm thấy địa chỉ")
	} else if err != nil {
		log.Error(err)
		return nil, errors.New("lỗi không xác định")
	}
	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	return &episodes, err
}
func (s *Service) GetProvince(account *entity.Account, province, district, ward, building string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	projection := bson.D{
		{"code", 1},
		{"district.name.$", 1},
		// {"district.ward.name", 1},
	}
	// var result entity.Province
	// https://www.mongodb.com/docs/manual/tutorial/project-fields-from-query-results/
	filter := bson.D{
		{"code", province},
		{"district",
			bson.D{
				{"$elemMatch", bson.D{
					{"name", "Bình Chánh"},
				}},
			},
		},
	}

	cursor, err := s.provincelDB.Find(ctx, filter, options.Find().SetProjection(projection))
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodes)
	return &episodes, err
}

func (s *Service) GetAddress(account *entity.Account, uuid string) (result *entity.Address, err error) {
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
