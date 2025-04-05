package seed

import (
	"database/sql"
	"log"
	"time"

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
