venv:
	python3 -m venv venv && source venv/bin/activate

install-pgvector:
	docker pull pgvector/pgvector:pg16

db:
	docker run --name pgvector-database -e POSTGRES_HOST_AUTH_METHOD=trust -p 5433:5432 -d pgvector/pgvector:pg16

install:
	pip3 install -r requirements.txt

createdb:
	bash scripts/db.sh --create

migrateup:
	bash scripts/db.sh --up

migratedown:
	bash scripts/db.sh --down

server:
	uvicorn src.app.main:app --port 7777 --host 0.0.0.0

build:
	bash scripts/build.sh

.PHONY: venv createdb migrateup migratedown server build install install-pgvector db