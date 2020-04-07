FROM srijanlabs/node:12.16.0-1.x-dev

WORKDIR /var/www/app

COPY package.json .

RUN npm install

COPY . .

CMD [ "server.js" ]
