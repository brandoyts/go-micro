protogen:
	protoc \
		--go_out=pkg/ \
		--go-grpc_out=pkg/ \
		-I proto \
		proto/booking.proto \
		proto/user.proto \
		proto/auth.proto

start:
	docker compose down --volumes && docker compose up -d --build

stop:
	docker compose down --volumes