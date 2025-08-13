docker compose down -v To Remove your containers
docker-compose up --build (it runs the script to populate and starts the container)
docker exec -it postgres psql -U postgres To connect to the database made by the docker compose

now go to the backend/service/storage/cmd

go run main.go

thats all
