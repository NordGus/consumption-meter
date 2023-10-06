ARG GO_VERSION=1.21.1
ARG ALPINE-VERSION=3.18
ARG NODE_VERSION=20.8
ARG DEBIAN_VERSION=12-slim

FROM node:${NODE_VERSION}-alpine${ALPINE_VERSION} as client-builder
WORKDIR /home/root/app
COPY . .
RUN yarn install && yarn build

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as go-builder
WORKDIR /usr/src
COPY . .
COPY --from=client-builder /home/root/app/dist /dist
RUN go build -o build/rom-app

FROM debian:${DEBIAN_VERSION}
RUN apt update -y && apt full-upgrade -y
COPY --from=go-builder /usr/src/build/rom-app /usr/local/bin
RUN chmod +x /usr/local/bin/rom-app
EXPOSE 4269
CMD ["rom-app", "--env=production"]
