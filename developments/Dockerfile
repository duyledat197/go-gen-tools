#* build:
FROM golang:1.20.0 as build-stage

WORKDIR /app
COPY .. /app
RUN CGO_ENABLED=0 && go build -o go-gen-tools ./cmd/srv/...

#* production: 
FROM alpine:3.14 as production-stage
COPY --from=build-stage /app/go-gen-tools /bin
RUN chmod +x /bin/go-gen-tools
