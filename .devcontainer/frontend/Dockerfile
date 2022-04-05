FROM node:16-alpine as build-stage

RUN apk add --update &&  apk add git && apk add bash

RUN mkdir -p /frontend/app

WORKDIR /frontend/app

RUN npm install -g @vue/cli

CMD npm run serve

# RUN npm install

# RUN npm create vite@latest