FROM golang:1.13 

WORKDIR /src

COPY . . 

RUN go mod tidy
RUN go mod download

EXPOSE 5000

CMD python3 main.py

