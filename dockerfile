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
COPY --from=build /app .
COPY --from=build /app/configs .

# Expose port 8080 for the container
EXPOSE 9000

# Run the executable
CMD ["./app"]
