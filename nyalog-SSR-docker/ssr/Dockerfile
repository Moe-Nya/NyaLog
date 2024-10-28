FROM node:18.17.1 AS builder

# 设置无交互模式
ENV DEBIAN_FRONTEND=noninteractive

# 更新并安装必要的依赖
RUN apt-get update && \
    apt-get install -y \
    wget \
    gnupg \
    ca-certificates \
    fonts-liberation \
    libappindicator3-1 \
    libasound2 \
    libatk-bridge2.0-0 \
    libatk1.0-0 \
    libcups2 \
    libdrm2 \
    libxcomposite1 \
    libxdamage1 \
    libxrandr2 \
    libgbm1 \
    xdg-utils \
    libu2f-udev \
    libvulkan1 \
    --no-install-recommends

# 下载并安装 Google Chrome 的 .deb 文件
RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb && \
    apt-get install -y ./google-chrome-stable_current_amd64.deb && \
    rm google-chrome-stable_current_amd64.deb

WORKDIR /app

COPY . /app

RUN npm config set registry https://registry.npmmirror.com

RUN npm i

EXPOSE 1500

CMD ["node", "service.js"]