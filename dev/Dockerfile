FROM golang:1.14.7-alpine3.12 AS build

WORKDIR /usr/local/phifd
COPY . .
RUN go build dev/main.go


FROM alpine:3.12

ARG ID

WORKDIR /usr/local/phifd
COPY --from=build /usr/local/phifd/main .
ENV ID $ID
CMD ./main -id $ID -port 4444
