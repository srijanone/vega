ARG NODE_VERSION=12
ARG BASE_VERSION=buster
FROM srijanlabs/node:${NODE_VERSION}-${BASE_VERSION}-dev as builder

WORKDIR /app

COPY . .

RUN npm i \
    && npm rebuild \
    && npm run build

FROM srijanlabs/nginx:10

COPY --from=builder --chown=continua /app/build/ /usr/share/nginx/html
