# Build stage
FROM golang:1.18-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary Go mod files to cache dependencies
COPY go.mod go.sum ./

# Download and cache Go dependencies
RUN go mod download

# Copy the entire project directory to the container
COPY . .

# Build the Go application with optimized flags
RUN go build -ldflags="-s -w" -o /app/onlineboutique ./cmd/...

# Final stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/onlineboutique .

# Set environment variables
ENV CART_SERVICE_ADDR="cart:8081" \
    PRODUCT_CATALOG_SERVICE_ADDR="productcatalog:8082" \
    CURRENCY_SERVICE_ADDR="currency:8083" \
    PAYMENT_SERVICE_ADDR="payment:8084" \
    SHIPPING_SERVICE_ADDR="shipping:8085" \
    EMAIL_SERVICE_ADDR="email:8086" \
    CHECKOUT_SERVICE_ADDR="checkout:8087" \
    RECOMMENDATION_SERVICE_ADDR="recommendation:8088" \
    AD_SERVICE_ADDR="ad:8089" \
    SHOPPING_ASSISTANT_SERVICE_ADDR="shoppingassistant:80"

# Specify the command to run the application
CMD ["./onlineboutique"]