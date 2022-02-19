package repository

import (
	"context"
	"fmt"
	"sopes-practica1/database"
	m "sopes-practica1/models"

	"go.mongodb.org/mongo-driver/bson"
)

var collection = database.GetCollection("historial")
var ctx = context.Background()

func Create(his m.Historial) error {
	var err error

	_, err = collection.InsertOne(ctx, his)

	if err != nil {
		return err
	}

	return nil
}

func Read() (m.Historiales, error) {
	var hiss m.Historiales

	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {

		var historial m.Historial
		err = cur.Decode(&historial)
		fmt.Print(historial)

		if err != nil {
			return nil, err
		}

		hiss = append(hiss, &historial)
	}

	return hiss, nil
}
