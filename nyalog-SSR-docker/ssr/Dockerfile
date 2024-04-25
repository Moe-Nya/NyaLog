FROM node:18.17.1 AS builder

# 安装chrome相关依赖
RUN apt-get update && apt-get install -y \
    wget \
    gnupg \
    ca-certificates \
    && wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - \
    && echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google-chrome.list \
    && apt-get update \
    && apt-get install -y google-chrome-stable

WORKDIR /app

COPY . /app

RUN npm config set registry https://registry.npmmirror.com

RUN npm i

EXPOSE 1500

CMD ["node", "service.js"]