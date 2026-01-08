FROM golang:1.24 AS golang

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o /main main.go

FROM gcr.io/distroless/static-debian12

COPY --from=golang /main .

# EXPOSE 1488

CMD ["/main"]

