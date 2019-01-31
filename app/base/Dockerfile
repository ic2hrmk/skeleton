# Build:            docker build -t skeleton-base .
# Run gateway:      docker run --rm -p=8081:8081 skeleton-base
# Run service:      docker run --rm -p=10001:10001 skeleton-base

FROM golang:1.10

WORKDIR /go/src/github.com/ic2hrmk/skeleton/app/base
COPY . .

RUN go build -o base-service main.go && mv base-service /go/bin/

CMD ["port-service","--env=.env","--address=:8081"]
