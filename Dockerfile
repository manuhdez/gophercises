# Start from golang base image
FROM golang:1.17-alpine as builder

WORKDIR /app

# Copy go.mod and go.sum to /build
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# Build the binary
ARG project_dir
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -a -v -o server $project_dir

# Start a new stage from a busybox base image
FROM busybox:latest

WORKDIR /dist

# Copy the build artifacts from the previous stage
COPY --from=builder /app/server .

# Copy the data folder from the previous stage
ARG data_dir
COPY  --from=builder /app/$data_dir ./data

# Run the binary
CMD ["./server"]
