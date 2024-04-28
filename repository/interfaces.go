// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	GetEstateById(ctx context.Context, id string) (output GetEstateByIdOutput, err error)
	InsertEstate(ctx context.Context, input InsertEstateInput) (id string, err error)
	InsertEstateObject(ctx context.Context, estateId string, input InsertEstateObjectInput) (id string, err error)
	GetEstateStats(ctx context.Context, id string) (output GetEstateStatsOutput, err error)
	GetDronePlanByEstateId(ctx context.Context, id string) (distance int, err error)
}
