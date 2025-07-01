FROM golang:1.24.4-alpine3.22
LABEL authors="gasmat"

WORKDIR /asciinator
COPY . .
ENV GIN_MODE=release

ENTRYPOINT ["go", "run", "./src/"]