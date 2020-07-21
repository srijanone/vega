ARG NODE_VERSION=12
FROM srijanlabs/node:${NODE_VERSION}-10-dev as builder

WORKDIR /app

COPY . .

RUN npm i 

FROM srijanlabs/node:${NODE_VERSION}-10

COPY . /app
COPY --from=builder --chown=continua /app/node_modules /app/

CMD [ "node", "app.js" ]
