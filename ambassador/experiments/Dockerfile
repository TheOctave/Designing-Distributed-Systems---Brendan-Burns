FROM node:24.7-alpine3.21

WORKDIR /app

COPY server.js .
COPY package.json .

RUN npm install

EXPOSE 3000

CMD ["node", "server.js"]