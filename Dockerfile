ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

# The floating golang:1 tag lags a Go release by days, and the image sets
# GOTOOLCHAIN=local, so on release day "go mod download" fails against a
# go.mod that requires the new minor. "auto" lets the build fetch the
# toolchain go.mod asks for; it goes back to a no-op once the base image
# catches up.
ENV GOTOOLCHAIN=auto

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app .


FROM debian:bookworm

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
