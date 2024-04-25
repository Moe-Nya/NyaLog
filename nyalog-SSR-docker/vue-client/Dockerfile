FROM node:18.17.1 AS builder

WORKDIR /app

COPY . /app

RUN npm config set registry https://registry.npmmirror.com

RUN npm i

EXPOSE 5173

CMD ["npm", "run", "dev"]