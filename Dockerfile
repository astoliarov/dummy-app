FROM golang:1.14-alpine as build

RUN mkdir /build & apk add --update make
WORKDIR /build


COPY . .
RUN make build

FROM alpine:latest as production
RUN mkdir /server
WORKDIR /server
COPY --from=build /build/dummy-app /server
CMD ["/server/dummy-app"]
