FROM golang:1.19

WORKDIR /app
COPY .env .
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]
