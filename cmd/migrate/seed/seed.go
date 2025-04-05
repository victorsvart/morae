package seed

import (
	"context"
	"database/sql"
	"log"
	"morae/internal/store/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func SeedPostgres(db *sql.DB) {
	var exists bool

	queryExists := "SELECT EXISTS (SELECT 1 FROM mre_users WHERE email_address = 'admin@godmode.com')"
	err := db.QueryRow(queryExists).Scan(&exists)
	if err != nil {
		log.Fatalf("failed to check if admin user exists: %v", err)
	}

	if exists {
		log.Println("Admin user already created")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123qwe"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to hash password: %v", err)
	}

	query := `
		INSERT INTO mre_users (full_name, email_address, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (email_address) DO NOTHING
	`

	now := time.Now()

	_, err = db.Exec(query, "admin admin", "admin@godmode.com", string(hashedPassword), now, now)
	if err != nil {
		log.Fatalf("failed to insert admin user: %v", err)
	}

	log.Println("Admin user seeded successfully.")
}

func SeedMongoDb(mongoDb *mongo.Database) {
	roomsCollection := mongoDb.Collection("rooms")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := roomsCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Println("Failed to count documents:", err)
		return
	}

	if count > 0 {
		log.Println("Rooms collection already seeded.")
		return
	}

  // []any needed otherwise insertmany bitches for some reason
	rooms := []any{
		mongodb.RoomDocument{ID: primitive.NewObjectID(), OwnerId: 1, FullAddress: "123 Main St, Springfield", Street: "Main St", Number: 123, District: "Downtown", State: "Illinois"},
		mongodb.RoomDocument{ID: primitive.NewObjectID(), OwnerId: 1, FullAddress: "456 Elm St, Metropolis", Street: "Elm St", Number: 456, District: "Uptown", State: "New York"},
		mongodb.RoomDocument{ID: primitive.NewObjectID(), OwnerId: 1, FullAddress: "789 Oak Ave, Gotham", Street: "Oak Ave", Number: 789, District: "Midtown", State: "New Jersey"},
		mongodb.RoomDocument{ID: primitive.NewObjectID(), OwnerId: 1, FullAddress: "321 Pine Rd, Star City", Street: "Pine Rd", Number: 321, District: "Old Town", State: "California"},
		mongodb.RoomDocument{ID: primitive.NewObjectID(), OwnerId: 1, FullAddress: "654 Cedar Blvd, Central City", Street: "Cedar Blvd", Number: 654, District: "East Side", State: "Nevada"},
		mongodb.RoomDocument{ID: primitive.NewObjectID(), OwnerId: 1, FullAddress: "987 Birch Ln, Coast City", Street: "Birch Ln", Number: 987, District: "West End", State: "Oregon"},
		mongodb.RoomDocument{ID: primitive.NewObjectID(), OwnerId: 1, FullAddress: "135 Maple Dr, Blüdhaven", Street: "Maple Dr", Number: 135, District: "Harbor", State: "Texas"},
		mongodb.RoomDocument{ID: primitive.NewObjectID(), OwnerId: 1, FullAddress: "246 Ash St, Keystone City", Street: "Ash St", Number: 246, District: "Northside", State: "Florida"},
		mongodb.RoomDocument{ID: primitive.NewObjectID(), OwnerId: 1, FullAddress: "579 Walnut Way, Smallville", Street: "Walnut Way", Number: 579, District: "South Park", State: "Kansas"},
		mongodb.RoomDocument{ID: primitive.NewObjectID(), OwnerId: 1, FullAddress: "864 Poplar Ct, Riverdale", Street: "Poplar Ct", Number: 864, District: "Civic Center", State: "Georgia"},
	}

	_, err = roomsCollection.InsertMany(ctx, rooms)
	if err != nil {
		log.Println("Failed to seed rooms:", err)
		return
	}

  log.Println("MONGODB -  Seeded room documents")
}
