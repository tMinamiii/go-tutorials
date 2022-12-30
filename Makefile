.PHONY: tools
tools:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
