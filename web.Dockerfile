FROM node:24 AS node

WORKDIR /app

COPY /web .

RUN npm i

RUN npm run build

FROM gcr.io/distroless/nodejs24-debian12

COPY --from=node /app/build .

ENV PUBLIC_SERVER="http://localhost:1488"

CMD ["index.js"]

EXPOSE 3000
