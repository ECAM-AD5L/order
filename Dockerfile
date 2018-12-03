#
# BUILD STAGE
# 
FROM golang:latest as build_base
WORKDIR /order

# Force the go compiler to use modules
ENV GO111MODULE=on

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

#This is the ‘magic’ step that will download all the dependencies that are specified in 
# the go.mod and go.sum file.
# Because of how the layer caching system works in Docker, the  go mod download 
# command will _ only_ be re-run when the go.mod or go.sum file change 
# (or when we add another docker instruction this line)
RUN go mod download

# This image builds the app
FROM build_base AS app_builder

# Here we copy the rest of the source code
COPY . .
# And compile the project
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -installsuffix cgo .

# 
# CERTS Stage
#
FROM alpine:latest as certs

# Install the CA certificates
RUN apk --update add ca-certificates

#
# PRODUCTION STAGE
# 
FROM scratch AS prod
# Copy the CA certificate from the certs stage
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# Copy the binary built during the build stage
COPY --from=app_builder /order/order .
CMD ["./order"]
