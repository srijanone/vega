ARG NODE_VERSION=12
FROM srijanlabs/node:${NODE_VERSION}-10-dev as builder

WORKDIR /app

COPY . .

RUN npm i \
    && npm rebuild \
    && npm run build

FROM srijanlabs/nginx:10

COPY --from=builder --chown=continua /app/build/ /usr/share/nginx/html
