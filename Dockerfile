FROM golang:1.16-alpine3.13 as Builder

WORKDIR /usr/build

COPY . .

RUN go install src/main.go
RUN go build src/main.go


FROM alpine:latest  
RUN apk --no-cache add ca-certificates bash

WORKDIR /usr/app

COPY .env .env
COPY ./wait-for-it.sh wait-for-it.sh
RUN chmod +x wait-for-it.sh
COPY --from=Builder /usr/build/main .

ENTRYPOINT [ "./main" ]