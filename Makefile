run-docker-compose:
	docker-compose up -d
down-docker-compose:
	docker-compose down

get-db-ip-pod:
	kubectl get pod db-postgres-59bdfcfc8f-6v2cz --template '{{.status.podIP}}'