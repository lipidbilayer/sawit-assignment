package repository

import (
	"context"
	"log"
	"strconv"
)

func (r *Repository) GetEstateById(ctx context.Context, id string) (output GetEstateByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT id, length, width FROM estates WHERE id = $1", id).Scan(&output.Id, &output.Length, &output.Width)
	if err != nil {
		return
	}
	return
}

func (r *Repository) InsertEstate(ctx context.Context, input InsertEstateInput) (id string, err error) {
	err = r.Db.QueryRowContext(ctx, "INSERT INTO estates (width, length) VALUES ($1, $2) RETURNING id", input.Width, input.Length).Scan(&id)
	return
}

func (r *Repository) InsertEstateObject(ctx context.Context, estateId string, input InsertEstateObjectInput) (id string, err error) {
	err = r.Db.QueryRowContext(ctx, "INSERT INTO estate_objects (estate_id, x_location, y_location, height) VALUES ($1, $2, $3, $4) RETURNING id", estateId, input.X, input.Y, input.Height).Scan(&id)
	return
}

func (r *Repository) GetEstateStats(ctx context.Context, id string) (output GetEstateStatsOutput, err error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT x_location, y_location, height FROM estate_objects WHERE estate_id = $1", id)
	if err != nil {
		return
	}

	results := []EstateObject{}
	defer rows.Close()
	for rows.Next() {
		item := EstateObject{}
		err = rows.Scan(&item.XLocation, &item.YLocation, &item.Height)
		if err != nil {
			log.Println(err)
		}
		results = append(results, item)
	}

	min, max, median := CalculateStats(results)

	return GetEstateStatsOutput{Count: len(results), Max: max, Min: min, Median: median}, nil
}

func (r *Repository) GetDronePlanByEstateId(ctx context.Context, id string) (distance int, err error) {
	query := `SELECT e.width, e.length, eo.x_location, eo.y_location, eo.height 
	FROM estates AS e 
	LEFT JOIN estate_objects AS eo ON e.id = eo.estate_id 
	WHERE estate_id = $1`
	rows, err := r.Db.QueryContext(ctx, query, id)
	if err != nil {
		return
	}

	objectMap := map[string]EstateObject{}
	estate := Estate{}
	defer rows.Close()
	for rows.Next() {
		object := EstateObject{}
		err = rows.Scan(&estate.Width, &estate.Length, &object.XLocation, &object.YLocation, &object.Height)
		if err != nil {
			log.Println(err)
		}
		coordinate := strconv.Itoa(object.XLocation) + "," + strconv.Itoa(object.YLocation)
		objectMap[coordinate] = object
	}
	distance = SumDroneTravelDistance(estate, objectMap)
	return distance, nil
}
