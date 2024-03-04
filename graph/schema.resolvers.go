package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"github.com/ishanz23/go-turso-starter-api/graph/model"
)

// CreateLocation is the resolver for the createLocation field.
func (r *mutationResolver) CreateLocation(ctx context.Context, input model.NewLocation) (*model.Location, error) {
	stmt, err := r.DB.Prepare("INSERT INTO location(name, state, description, lat, long, altitude, coverUrl) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(input.Name, input.State, input.Description, input.Lat, input.Long, input.Altitude, input.CoverURL)
	if err != nil {
		return nil, err
	}
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	newLocation := model.Location{
		ID:          int(lastInsertedID),
		Name:        &input.Name,
		State:       &input.State,
		Description: input.Description,
		Lat:         &input.Lat,
		Long:        &input.Long,
		Altitude:    &input.Altitude,
		CoverURL:    &input.CoverURL,
	}
	return &newLocation, nil
}

// CreateHomestay is the resolver for the createHomestay field.
func (r *mutationResolver) CreateHomestay(ctx context.Context, input model.NewHomestay) (*model.Homestay, error) {
	stmt, err := r.DB.Prepare("INSERT INTO homestay(name, address, locationId) VALUES(?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(input.Name, input.Address, input.LocationID)
	if err != nil {
		return nil, err
	}
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	newHomestay := model.Homestay{
		ID:      int(lastInsertedID),
		Name:    &input.Name,
		Address: &input.Address,
	}
	return &newHomestay, nil
}

// CreateRoom is the resolver for the createRoom field.
func (r *mutationResolver) CreateRoom(ctx context.Context, input model.NewRoom) (*model.Room, error) {
	stmt, err := r.DB.Prepare("INSERT INTO room(name, category, baseOccupancy, extraOccupancy, toiletAttached, balconyAttached, kitchenAttached, airConditioned, recommended, isDorm, homestayId) VALUES(? ,?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	// generate default values for extraOccupancy, toiletAttached, balconyAttached, kitchenAttached, airConditioned, recommended, isDorm
	if input.ExtraOccupancy == nil {
		extraOccupancy := 0
		input.ExtraOccupancy = &extraOccupancy
	}
	if input.ToiletAttached == nil {
		toiletAttached := true
		input.ToiletAttached = &toiletAttached
	}
	if input.BalconyAttached == nil {
		balconyAttached := false
		input.BalconyAttached = &balconyAttached
	}
	if input.KitchenAttached == nil {
		kitchenAttached := false
		input.KitchenAttached = &kitchenAttached
	}
	if input.AirConditioned == nil {
		airConditioned := false
		input.AirConditioned = &airConditioned
	}
	if input.Recommended == nil {
		recommended := false
		input.Recommended = &recommended
	}
	if input.IsDorm == nil {
		isDorm := false
		input.IsDorm = &isDorm
	}

	fmt.Println(input)
	// insert into the database
	result, err := stmt.Exec(input.Name, input.Category, input.BaseOccupancy, input.ExtraOccupancy, input.ToiletAttached, input.BalconyAttached, input.KitchenAttached, input.AirConditioned, input.Recommended, input.IsDorm, input.HomestayID)
	if err != nil {
		return nil, err
	}
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	newRoom := model.Room{
		ID:              int(lastInsertedID),
		Name:            input.Name,
		Category:        input.Category,
		BaseOccupancy:   &input.BaseOccupancy,
		ExtraOccupancy:  input.ExtraOccupancy,
		ToiletAttached:  input.ToiletAttached,
		BalconyAttached: input.BalconyAttached,
		KitchenAttached: input.KitchenAttached,
		AirConditioned:  input.AirConditioned,
		Recommended:     input.Recommended,
		IsDorm:          input.IsDorm,
		Homestay:        &model.Homestay{},
	}
	return &newRoom, nil
}

// Locations is the resolver for the locations field.
func (r *queryResolver) Locations(ctx context.Context) ([]*model.Location, error) {
	rows, err := r.DB.Query("SELECT id, name, state, description, lat, long, altitude, coverUrl from location")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	locations := []*model.Location{}
	for rows.Next() {
		location := model.Location{}
		fmt.Println("Location is: ", rows)
		if err := rows.Scan(&location.ID, &location.Name, &location.State, &location.Description, &location.Lat, &location.Long, &location.Altitude, &location.CoverURL); err != nil {
			return nil, err
		}
		locations = append(locations, &location)
	}
	return locations, nil
}

// Homestays is the resolver for the homestays field.
func (r *queryResolver) Homestays(ctx context.Context, locationID *int) ([]*model.Homestay, error) {
	rows, err := r.DB.Query("SELECT homestay.id, homestay.name, homestay.address, location.name, location.description, location.altitude, location.coverUrl, location.lat, location.long, location.state from homestay INNER JOIN location ON homestay.locationId=location.id where location.id = ?", locationID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	homestays := []*model.Homestay{}
	for rows.Next() {
		homestay := model.Homestay{Location: &model.Location{ID: 0, Name: nil, Lat: nil, Long: nil, State: nil, Altitude: nil, Description: nil, CoverURL: nil, Homestays: nil}}
		location := model.Location{}
		if err := rows.Scan(&homestay.ID, &homestay.Name, &homestay.Address, &location.Name, &location.Description, &location.Altitude, &location.CoverURL, &location.Lat, &location.Long, &location.State); err != nil {
			return nil, err
		}
		homestay.Location = &location
		homestays = append(homestays, &homestay)
	}
	return homestays, nil
}

// Homestay is the resolver for the homestay field.
func (r *queryResolver) Homestay(ctx context.Context, homestayID int) (*model.Homestay, error) {
	var homestay model.Homestay

	// Query for Homestay
	err := r.DB.QueryRow("SELECT id, name, address FROM homestay WHERE id = ?", homestayID).Scan(
		&homestay.ID,
		&homestay.Name,
		&homestay.Address,
	)
	if err != nil {
		return nil, err
	}

	// Query for associated Rooms
	rows, err := r.DB.Query("SELECT id, name, category, toiletAttached FROM room WHERE homestayId = ?", homestayID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Populate the rooms in the homestay
	for rows.Next() {
		var room model.Room
		if err := rows.Scan(&room.ID, &room.Name, &room.Category, &room.ToiletAttached); err != nil {
			return nil, err
		}
		homestay.Rooms = append(homestay.Rooms, &room)
	}

	return &homestay, nil
}

// Rooms is the resolver for the rooms field.
func (r *queryResolver) Rooms(ctx context.Context) ([]*model.Room, error) {
	panic(fmt.Errorf("not implemented: Rooms - rooms"))
}

// Room is the resolver for the room field.
func (r *queryResolver) Room(ctx context.Context, homestayID int) ([]*model.Room, error) {
	panic(fmt.Errorf("not implemented: Room - room"))
}

// AskMeAnything is the resolver for the askMeAnything field.
func (r *queryResolver) AskMeAnything(ctx context.Context, prompt *string) (string, error) {
	// r.GenAi.StartChat().SendMessage(ctx, prompt)
	promptObj := []genai.Part{genai.Text(*prompt)}
	response, err := r.GenAi.GenerateContent(ctx, promptObj...)
	fmt.Println("Response is: ", response)
	fmt.Println("err is: ", err)
	return response.PromptFeedback.BlockReason.String(), err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
