FROM golang:1.17-buster AS build


# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

ENV GOARCH=amd64

ENV GO111MODULE=off
RUN go build -o /go/bin/app cmd/main.go


FROM gcr.io/distroless/base-debian11
COPY --from=build /go/bin/app /app

# This container exposes port 8080 to the outside world
EXPOSE 8080
USER nonroot:nonroot
# Run the binary program produced by `go install`
CMD ["/app"]