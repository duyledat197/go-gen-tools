FROM duyledat197/golang-proto:0.1 as build-stage

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 && go build -o app-exe ./cmd/srv/...


# production: 
FROM golang:1.20.0 as production-stage

COPY --from=build-stage /app/app-exe /bin

RUN chmod +x /bin/app-exe

ENTRYPOINT [ "/bin/app-exe", "start" ]

