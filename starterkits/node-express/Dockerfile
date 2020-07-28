ARG NODE_VERSION=12
ARG BASE_VERSION=buster
FROM srijanlabs/node:${NODE_VERSION}-${BASE_VERSION}-dev as builder

WORKDIR /app

COPY . .

RUN npm i

FROM srijanlabs/node:${NODE_VERSION}-buster

COPY . /app
COPY --from=builder --chown=continua /app/node_modules /app/

CMD [ "node", "index.js" ]
