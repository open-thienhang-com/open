package service

// func CreateBook(b entity.Book) (string, error) {
// 	result, err := entity.BooksCollection.InsertOne(Ctx, b)
// 	if err != nil {
// 		return "0", err
// 	}
// 	return fmt.Sprintf("%v", result.InsertedID), err
// }

// func GetBook(id string) (Book, error) {
// 	var b Book
// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return b, err
// 	}

// 	err = BooksCollection.
// 		FindOne(Ctx, bson.D{{"_id", objectId}}).
// 		Decode(&b)
// 	if err != nil {
// 		return b, err
// 	}
// 	return b, nil
// }

// func GetBooks() ([]Book, error) {
// 	var book Book
// 	var books []Book

// 	cursor, err := BooksCollection.Find(Ctx, bson.D{})
// 	if err != nil {
// 		defer cursor.Close(Ctx)
// 		return books, err
// 	}

// 	for cursor.Next(Ctx) {
// 		err := cursor.Decode(&book)
// 		if err != nil {
// 			return books, err
// 		}
// 		books = append(books, book)
// 	}

// 	return books, nil
// }

// func UpdateBook(id primitive.ObjectID, pageCount int) error {
// 	filter := bson.D{{"_id", id}}
// 	update := bson.D{{"$set", bson.D{{"page_count", pageCount}}}}
// 	_, err := BooksCollection.UpdateOne(
// 		Ctx,
// 		filter,
// 		update,
// 	)
// 	return err
// }
