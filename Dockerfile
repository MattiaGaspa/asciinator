FROM golang:1.24.4-alpine3.22

WORKDIR /asciinator
COPY . .
RUN go mod download
ENV GIN_MODE=release

ENTRYPOINT ["go", "run", "."]