FROM node:24.7-alpine3.22

WORKDIR /app

COPY ./cloned-repo/ .

RUN apk add --no-cache openssh-client git

EXPOSE 8099

CMD ["node", "--watch", "server.js"]