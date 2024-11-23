package cli

import (
	"be_api/app/database"
	"be_api/app/models"
	"fmt"
	"log"
	"os/exec"

	_ "github.com/golang-migrate/migrate/source/file" // Add this line
)

func CliRunTask(args []string) {

	if args[0] == "migrate" {
		if args[1] == "create" {
			cmd := exec.Command("migrate", "create", "-ext", "sql", "-dir", "migrations", args[2])

			err := cmd.Run()
			if err != nil {
				log.Fatalf("Error creating migration: %v", err)
			}

			log.Println("Migration created successfully!")
		} else {
			database.DBMigrate(args[1])
		}

	}

	if args[0] == "seed" {
		if args[1] == "dev" {
			database.DBWipe()
			database.DBMigrate("up")

			pCat := make([]models.ProductCategory, 300000)

			for i := 0; i < 300000; i++ {
				pCat[i] = models.ProductCategory{
					Title: fmt.Sprintf("Test%d", i+1),
				}
			}

			database.DB.CreateInBatches(pCat, 200)

		}

		if args[1] == "prod" {
			database.DBWipe()
			database.DBMigrate("up")
		}
	}

	fmt.Println("RUN CliRunTask ", args, len(args), "Arguments")
}
