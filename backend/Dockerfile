FROM golang:1.20-alpine AS dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the Go application source code into the container's working directory
COPY . .
RUN chmod +x main.go

# Build the Go application
ARG APP_NAME
ARG SERVER_HOST
ARG SERVER_PORT
ARG DB_USERNAME
ARG DB_PASSWORD
ARG DB_HOST
ARG DB_PORT
ARG DB_NAME
RUN go build -o main -ldflags="-X 'main.appName=${APP_NAME}' -X 'main.serverHost=${SERVER_HOST}' -X 'main.serverPort=${SERVER_PORT}' -X 'main.dbUsername=${DB_USERNAME}' -X 'main.dbPassword=${DB_PASSWORD}' -X 'main.dbHost=${DB_HOST}' -X 'main.dbPort=${DB_PORT}' -X 'main.dbName=${DB_NAME}'"

# Create a new stage for the final image
FROM gcr.io/distroless/static-debian11 as prod
WORKDIR /app

COPY --from=dev /app .
ENV APP_NAME=${APP_NAME}
ENV SERVER_HOST=${SERVER_HOST}
ENV SERVER_PORT=${SERVER_PORT}
ENV DB_USERNAME=${DB_USERNAME}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_NAME=${DB_NAME}
USER root
CMD ["/app/main","serve"]