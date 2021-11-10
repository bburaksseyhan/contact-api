FROM golang:1.16-alpine as build-env
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN  go build -o /contact-api github.com/bburaksseyhan/contact-api/src/cmd/api   

FROM alpine:3.14

RUN apk update \
    && apk upgrade\
    && apk add --no-cache tzdata curl

#RUN apk --no-cache add bash
ENV TZ Europe/Istanbul

WORKDIR /app
COPY --from=build-env /contact-api .
COPY --from=build-env /app/src/cmd/api /app/

EXPOSE 80
CMD [ "./contact-api" ]