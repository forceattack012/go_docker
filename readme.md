#Create Network 
    docker network create todo-network

#Docker PostgreSQL
    docker run --rm --network todo-network --name todo-postgres -e POSTGRES_PASSWORD=todopassword -dp 5439:5432 postgres