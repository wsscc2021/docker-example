# Build stage
FROM golang:1.17-alpine AS build

WORKDIR /source

COPY src/ ./

RUN go mod tidy \
 && go mod download \
 && go build -o ./main

# Runtime stage
FROM golang:1.17-alpine

WORKDIR /app

COPY --from=build /source/main ./

RUN chmod +x ./main

ENTRYPOINT ["./main"]