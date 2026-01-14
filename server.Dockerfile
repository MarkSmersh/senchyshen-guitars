FROM golang:1.24 AS golang

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o /main main.go

FROM gcr.io/distroless/static-debian12

COPY --from=golang /main .

ENV DB_URL=postgres://ss:445552@localhost:5432/senchyshen_guitars
ENV GIN_MODE=release

VOLUME ["/assets"]

EXPOSE 3000

CMD ["/main"]
