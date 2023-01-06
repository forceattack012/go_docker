# Create Network 
    docker network create todo-network

# Docker PostgreSQL
    docker run --rm --network todo-network --network-alias pg-todo  --name todo-postgres -e POSTGRES_PASSWORD=todopassword -dp 5439:5432 postgres

# build docker todo-api 
     docker build -t todo-api .
# run docker todo-api
    docker run -dp 8080:8080 --network todo-network todo-api 
    --network is config network for communicate sql and app
# docker log todo-api 
    docker logs <docker-id>

# dsn for docker in golang app 
    host=pg-todo user=postgres password=todopassword dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Bangkok 
    host is network-alias from PostgreSQL we config
    port is use internal pg's port

# run todo-app with env 
    docker run -dp 8080:8080 --env-file docker.env --network todo-network todo-api 

# run docker-compose
    docker compose up -d
# down docker-comnpose
    docker-compose down 