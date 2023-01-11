# Create Network 
    docker network create todo-network

# Docker PostgreSQL
    docker run --rm --network todo-network --network-alias pg-todo  --name todo-postgres -e POSTGRES_PASSWORD=todopassword -dp 5439:5432 postgres

# build docker todo-api 
     docker build --no-cache -t todo-api .
     docker build -f dockerfile.kube -t todo-api:kube . 
# run docker todo-api
   #    [mount] 
        docker run -dp 8080:8080 --mount type=bind,src="${pwd}/compose_dev_config",target=/config --network todo-network todo-api
    --network is config network for communicate sql and app
# docker log todo-api 
    docker logs <docker-id>

# dsn for docker in golang app 
    host=pg-todo user=postgres password=todopassword dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Bangkok 
    host is network-alias from PostgreSQL we config
    port is use internal pg's port

# run todo-app with env 
    docker run -dp 8080:8080 -e env=local --network todo-network todo-api
    docker run -dp 8080:8080 -e env=dev --network todo-network todo-api  

# run docker-compose
    docker compose up -d
# down docker-comnpose
    docker-compose down 

# add file yml in container when config's wrong must new build not ok
    --

# create secret 
    kubectl create secret generic ${name} --from-env-file={file_env}
    kubectl create secret generic todo-env --from-env-file=kube.env

#   delete secret
    kubectl delete secret todo-env

#   kubectl deploy 
    kubectl apply -f .\todo-db-deployment.yml   (db)
    kubectl apply -f .\todo-api-deployment.yml  (api)

#   kubectl delete deployment 
    kubectl delete deployment todo-api

# kubectl log view logs of app 
    kubectl logs ${pod_id}

# kubectl delete service 
    kubectl delete svc ${name1} ${name2}
    kubectl delete svc example-service hello-node todo-api-service todo-api-svc

# type load baland Comunicating between api and db using internal IP 
    example db's ip : 10.1.2.3, port 841
        api connect db by ip 10.1.2.3:841

# install ingress-controller for expose by ingress
    kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.5.1/deploy/static/provider/cloud/deploy.yaml

    kubectl get pods --namespace=ingress-nginx

# fix ingress 
    https://stackoverflow.com/questions/66998567/kubernetes-ingress-not-working-on-docker-for-desktop-mac

# create config map 
    kubectl create configmap todo-api-config --from-file=kube/config.yaml 
    kubectl create -f .\todo-api-configmap.yaml 

# update configMap 
    kubectl replace -f .\todo-api-configmap.yaml 
     kubectl replace -f kubernetes-configmap.yaml

# kube get pod name from app in deployment
    kubectl get pods -l app=db-postgres -o custom-columns=:metadata.name

# kube get internal ip pod 
    kubectl get pod db-postgres-59bdfcfc8f-6v2cz --template '{{.status.podIP}}'



# in container if expose 80 but you not chaNGE PORT IN APP it's route to port 8080 you change it your app run start port