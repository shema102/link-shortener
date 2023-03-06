# stage 1 - build the react app
FROM node:19-alpine as node

WORKDIR /app

COPY client/package*.json ./
COPY client/yarn.lock ./

RUN yarn install

COPY client/ ./

RUN yarn build

# stage 2 - build the server
FROM golang:1.20-alpine as go

WORKDIR /server

COPY server/go.mod ./

RUN go mod download

COPY server/ ./

RUN go build -o app

# final stage
FROM alpine:latest

WORKDIR /app

COPY --from=node /app/dist ./client/dist

COPY --from=go /server/app ./server/app

EXPOSE 8080

CMD ["./server/app"]
