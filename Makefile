connect-to-db:
	psql "sslmode=verify-ca sslrootcert=config/server-ca.pem sslcert=config/client-cert.pem sslkey=config/client-key.pem hostaddr=34.146.103.238 port=5432 user=postgres password=[[hQ%Kz?]DI%Tss,"

deploy-dev1:
	git checkout dev1
	git fetch origin
	git reset --hard origin/main
	git merge ${branch}
	git log -n 2
	git push -f origin dev1
	git checkout ${branch}

# docker立ち上げ + protoからGoコードを自動生成
.PHONY: proto
proto:
	docker compose -f docker-compose.gen.yml up -d
	make docker_proto
	docker compose -f docker-compose.gen.yml down

# protoからGoコードを自動生成
.PHONY: docker_proto
docker_proto:
	docker container exec bottlist_gen sh -c '(cd /app && make generate_proto)'

.PHONY: generate_proto
generate_proto:
	sh scripts/gen_proto.sh

.PHONY: install
install:
	go clean -cache
	go install github.com/gogo/protobuf/protoc-gen-gofast@v1.3.2
	go install github.com/envoyproxy/protoc-gen-validate@v0.6.1