venv:
	python3 -m venv venv && source venv/bin/activate

migrateup:
	bash scripts/db.sh --up

migratedown:
	bash scripts/db.sh --down

server:
	uvicorn src.app.main:app --port 7777 --host 0.0.0.0

build:
	bash scripts/build.sh

.PHONY: venv migrateup migratedown server build