FROM srijanlabs/node:12.16.0-1.x-dev

USER continua

WORKDIR /var/www/app

COPY package* ./


RUN npm install

COPY . .

CMD [ "server.js" ]
