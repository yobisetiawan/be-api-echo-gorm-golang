package database

import (
	"log"
)

func DBWipe() {
	if DB == nil {
		log.Fatalf("DB is not initialized")
	}

	query := `
				DO $$ 
				DECLARE
					r RECORD;
				BEGIN
					-- Disable referential integrity checks
					EXECUTE 'SET session_replication_role = replica';

					-- Loop over all tables and drop them
					FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') 
					LOOP
						EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
					END LOOP;

					-- Enable referential integrity checks
					EXECUTE 'SET session_replication_role = DEFAULT';
				END $$;
			`

	ress := DB.Exec(query)
	if ress.Error != nil {
		log.Fatalf("Failed to drop all tables: %v", ress.Error)
	}

	log.Println("All tables dropped successfully.")
}

func DBMigrate(action string) {
	DBAutoMigrate()
}
