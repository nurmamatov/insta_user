CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/proto_gen.sh ${CURRENT_DIR}
	ls genproto/*.pb.go | xargs -n1 -IX bash -c "sed -e '/bool/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"

mig-creat:
	migrate create -ext sql -dir migrations -seq create_user_photo_tables

mig-up:
	migrate -source file:./migrations -database 'postgres://khusniddin:1234@localhost:5432/insta_user?sslmode=disable' up

mig-down:
	migrate -source file:./migrations -database 'postgres://khusniddin:1234@localhost:5432/insta_user?sslmode=disable' down
