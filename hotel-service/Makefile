include .env



MIGRATION_DIR=migrations
DB_URL=$(DB_USER):$(DB_PASSWORD)@tcp($(DB_ADDRESS))/hotel

migrate_create:
	goose -dir $(MIGRATION_DIR) create $(name) sql
migrate-up:
	goose -dir $(MIGRATION_DIR) mysql "$(DB_URL)" up
migrate-down:
	goose -dir $(MIGRATION_DIR) mysql "$(DB_URL)" down
migrate-reset:
	goose -dir $(MIGRATION_DIR) mysql "$(DB_URL)" reset
migrate-redo:
	goose -dir $(MIGRATION_DIR) mysql "$(DB_URL)" redo
migrate-status:
	goose -dir $(MIGRATION_DIR) mysql "$(DB_URL)" status
migrate-to:
	goose -dir $(MIGRATION_DIR) mysql "$(DB_URL)" up-to $(version)
migrate-down-to:
	goose -dir $(MIGRATION_DIR) mysql "$(DB_URL)" down-to $(version)
migrate-force:
	goose -dir $(MIGRATION_DIR) mysql "$(DB_URL)" force $(version)
migrate-help:
	goose -h