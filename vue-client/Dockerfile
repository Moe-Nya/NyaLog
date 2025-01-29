FROM node:18.17.1 AS builder

WORKDIR /app

COPY . /app

RUN npm config set registry https://registry.npmmirror.com

RUN npm install -g npm@9.6.7
RUN npm i
RUN npm run build

FROM nginx:latest

EXPOSE 80

WORKDIR /app

COPY --from=builder /app/dist /usr/share/nginx/html
COPY --from=builder /app/nginx /etc/nginx/conf.d

# 注释以使用官方镜像启动命令
# CMD ["nginx", "-g", "daemon off;"]