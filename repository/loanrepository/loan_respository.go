package loanrepository

import (
	"context"
	"loans/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoanRepository struct {
	LoanCollection *mongo.Collection
}

func NewLoanRepository(database *mongo.Database) *LoanRepository {
	return &LoanRepository{
		LoanCollection: database.Collection("loans"),
	}
}

func (l *LoanRepository) InsertLoan(loan *domain.Loan) error {
	_, err := l.LoanCollection.InsertOne(context.Background(), loan)
	return err
}

func (l *LoanRepository) GetLoanByID(id primitive.ObjectID) (*domain.Loan, error) {
	loan := &domain.Loan{}
	filter := bson.M{"_id": id}

	err := l.LoanCollection.FindOne(context.Background(), filter).Decode(loan)
	return loan, err
}

func (l *LoanRepository) GetLoanByUserID(userID primitive.ObjectID) (*domain.Loan, error) {
	loan := &domain.Loan{}
	filter := bson.M{"user_id": userID}

	err := l.LoanCollection.FindOne(context.Background(), filter).Decode(loan)
	return loan, err
}

func (l *LoanRepository) UpdateLoan(id primitive.ObjectID, loanData bson.M) error {
	_, err := l.LoanCollection.UpdateOne(context.Background(), id, loanData)
	return err
}

func (l *LoanRepository) DeleteLoan(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := l.LoanCollection.DeleteOne(context.Background(), filter)
	return err
}

func (l *LoanRepository) GetLoans(page, limit int) ([]*domain.Loan, error) {
	// Calculate the number of documents to skip
	skip := (page - 1) * limit

	// Define an empty slice to hold the results
	var loans []*domain.Loan

	// Create the find options
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))

	// Execute the query
	cursor, err := l.LoanCollection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	// Defer closing the cursor
	defer cursor.Close(context.TODO())

	// Iterate over the cursor and decode each document
	for cursor.Next(context.TODO()) {
		loan := &domain.Loan{}

		err := cursor.Decode(loan)
		if err != nil {
			return nil, err
		}

		loans = append(loans, loan)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return loans, nil
}
