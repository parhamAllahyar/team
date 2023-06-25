
FROM golang:1.20.5-alpine3.18

WORKDIR /app

COPY . .
RUN go mod tidy

WORKDIR ./cmd



ENTRYPOINT  ["go" , "run" , "main.go"]





# FROM golang:1.20.5-alpine3.18 as builder

# # Declaring env variables
# ENV GOOS linux
# ENV CGO_ENABLED 0

# # Setting up the work directory
# WORKDIR /app

# # Installing dependencies
# COPY go.mod  ./
# RUN go mod download

# # Copying all the files in our project
# COPY . .

# # Building our application
# RUN go build  app


# FROM alpine:3.16 as production


# RUN apk add --no-cache ca-certificates


# COPY --from=builder app .

# # Starting our application
# CMD ./app

# # Exposing server port
# EXPOSE 4011
