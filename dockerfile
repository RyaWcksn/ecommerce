# Build stage
FROM golang AS build

# Set the working directory to /app
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Runtime stage
FROM alpine:latest

# Install any required dependencies
RUN apk --no-cache add ca-certificates

# Set the working directory to /app
WORKDIR /app

# Copy the executable from the build stage
COPY configs/config.yaml /app/configs/config.yaml
COPY --from=build /app/app .

# Expose port 9000 for the container
EXPOSE 9000

# Run the executable
CMD ["./app"]
