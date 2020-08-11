FROM golang:1.13 

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download

COPY . . 

EXPOSE 5000

CMD go run main.go

