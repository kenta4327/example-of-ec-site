db-migrate-dry-run:
	docker run -v ${PWD}/schema:/schema --rm \
	--net=example-of-ec-site_default \
	--platform linux/amd64 \
	diverse/mysqldef \
	-u root \
	-p hoge \
	-h db \
	--file=/schema/tables.sql \
	--dry-run \
	sample

db-migrate:
	docker run -v ${PWD}/schema:/schema --rm \
	--net=example-of-ec-site_default \
	--platform linux/amd64 \
	diverse/mysqldef \
	-u root \
	-p hoge \
	-h db \
	--file=/schema/tables.sql \
	sample
