default:
	@echo "=============building Local API============="
	docker build -f src/gopoll/Dockerfile -t gopoll .

up: default
	@echo "=============starting api locally============="
	docker-compose up -d && docker-compose logs -f

logs:
	docker-compose logs -f

down:
	docker-compose down

test:
	go test -v -cover ./...

clean: down
	@echo "=============cleaning up============="		

	docker system prune -f
	docker volume prune -f
