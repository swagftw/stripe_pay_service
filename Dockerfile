# build iamge
FROM golang:1.18-alpine AS BUILD

RUN apk --no-cache add ca-certificates

# change directory to the root of the project
WORKDIR /app

# copy project to the container
COPY . .

# download deps
RUN go mod download && go mod vendor

# build the project
RUN CGO_ENABLED=0 go build -o /api_server ./cmd/api/main.go

# deploy or run
FROM scratch

WORKDIR /

# copy ca cert
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# copy api server and config files
COPY --from=build /api_server /api_server
COPY --from=build /app/utl/config/config.local.yaml /utl/config/
COPY --from=build /app/.env /

# will expose port
EXPOSE 8080

# run the api server
CMD ["/api_server"]


