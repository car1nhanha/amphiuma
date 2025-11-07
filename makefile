.PHONY: up down is_running create_table_repo create_table_file

up:
	@docker compose up -d

down:
	@docker compose down

is_running:
	@docker ps

list_tables:
	@docker exec -it localstack awslocal dynamodb list-tables

create_table_repo:
	@aws dynamodb create-table \
		--table-name repos \
		--attribute-definitions AttributeName=owner_login,AttributeType=S \
		--key-schema AttributeName=owner_login,KeyType=HASH \
		--billing-mode PAY_PER_REQUEST \
		--endpoint-url http://localhost:4566 \
		--region us-east-1

create_table_file:
	@aws dynamodb create-table \
		--table-name files \
		--attribute-definitions AttributeName=owner_login,AttributeType=S \
		--key-schema AttributeName=owner_login,KeyType=HASH \
		--billing-mode PAY_PER_REQUEST \
		--endpoint-url http://localhost:4566 \
		--region us-east-1
