FROM golang:1.26 AS build
WORKDIR /src
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -o /tiny-api .

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /tiny-api /tiny-api
EXPOSE 8080
ENTRYPOINT ["/tiny-api"]
