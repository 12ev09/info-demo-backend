build:
	docker-compose build

up:
	docker-compose up -d

migrate_up:
	docker-compose exec app bash -c "sql-migrate up"

migrate_down:
	docker-compose exec app bash -c "sql-migrate down"

migrate_status:
	docker-compose exec app bash -c "sql-migrate status"

migrate_create:
	docker-compose exec app bash -c "sql-migrate new create_items"

logs:
	docker-compose logs -f　app
