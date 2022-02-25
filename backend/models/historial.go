package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Historial struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Numero1   float64            `json:"numero1,omitempty"`
	Numero2   float64            `json:"numero2,omitempty"`
	Operacion string             `json:"operacion,omitempty"`
	Resultado float64            `json:"resultado,omitemty"`
	Fecha     time.Time          `bson:"fecha,omitemty"`
}

type Historiales []*Historial
