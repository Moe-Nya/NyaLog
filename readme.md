<div align=center><img src="https://raw.githubusercontent.com/Moe-Nya/NyaLog/main/staticimg/NyaLog.PNG" style="height: 300px; width:320px;"></div>

# NyaLog

> 🐱 MoeNya's lovely blog.

## 介绍

NyaLog是一个由Golang的Gin框架和Vue.js制作的前后端分离、自适应响应式博客。作者一直以来都有一个跳出博客框架的想法，这次将它付诸现实了~

目前NyaLog还存在一些缺陷：如果导航栏、FindMe还不能支持自排序（在开发的时候没有想到），在未来将会对这些使用上影响体验的地方进行修改。

[开发文档](https://github.com/Moe-Nya/NyaLog/blob/main/document/readme.md)
<div align=center><img src="https://raw.githubusercontent.com/Moe-Nya/NyaLog/main/staticimg/main page pc.png"></div>

<div align=center><img src="https://raw.githubusercontent.com/Moe-Nya/NyaLog/main/staticimg/main page phone.png"></div>

## 未来更新计划

1. ~~文章、页面图片懒加载~~
2. ~~动态Meta~~
3. ~~尝试SSR（server side rendering）~~(但是没能够实现一键化、自动化。后面会专门写一篇文章来说明如何做SSR)
4. 在文章、页面中引入H5播放器及Aplayer和Dplayer
5. 制作RSS
6. 导航栏自定义排序

## 部署NyaLog

### 下载

在package中下载NyaLog。

### 安装Docker和Docker compose

[docker安装文档](https://docs.docker.com/engine/install/)

### 使用Docker compose运行

##### 配置docker-compose.yml

```yml
services:
  nyalog-mysql:
    container_name: nyalog-mysql
    image: mysql
    command: --default-authentication-plugin=mysql_native_password
    restart: always
    volumes:
      - ./mysql/data:/var/lib/mysql
    environment:
      MYSQL_DATABASE: nyalog
      MYSQL_USER: nyalog
      # 可自定义数据库用户密码
      MYSQL_PASSWORD: nyalog
      MYSQL_ROOT_PASSWORD: nyalog

  nyalog-server:
    container_name: nyalog-server
    depends_on:
      - nyalog-mysql
    build:
      context: ./gin-server
      dockerfile: Dockerfile
    volumes:
      - ./gin-server/gin-blog-server/config:/app/config
    environment:
      - CONFIG_PATH=/app/config/config.ini
      - MYSQL_HOST=nyalog-mysql
    ports:
    # 可自定义后端端口(两个1300都改)
      - "1300:1300"

  nyalog-web:
    container_name: nyalog-web
    build:
      context: ./vue-client
      dockerfile: Dockerfile
    ports:
    # 可自定义前端端口(注意！80端口不要动，动后面的1400)
      - "1400:80"
```

修改`#`注释项下面的内容，其它地方如果对docker不熟悉不建议修改。

##### 后端配置

如果你修改了`nyalog-server`中的端口，那么需要在`gin-server/gin-blog-server`目录下打开Dockerfile修改注释项后的端口：

```dockerfile
FROM golang:1.21.6 AS builder

WORKDIR /app

ENV GOPROXY https://goproxy.cn,direct

# 修改这里
EXPOSE 1300

COPY . /app

RUN go mod download
RUN CGO_ENABLED=0 go build -v=0 -a -trimpath -ldflags "-s -w -extldflags '-static'" -o dist ./main.go

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=builder /app/dist /app/server
CMD sh -c "sleep 5 && /app/server"
```

接着打开`gin-server/gin-blog-server/config/config.ini`按照提示进行修改：

```ini
[System]
# 后端访问地址
Domain = http://localhost:8080
# 密码最小长度
PasswordMinLen = 6
# 管理员用户入口
AdminEntrance = admin
# 设置GPT或者阿里千问的地址和key
GPTURL = 
GPTKey = 
QWURL = 
QWKey =
# 设置GitHub的id和secret
GitHubID = 
GitHUbSecret = 

[Server]
# debug 开发模式，release 生产模式
AppMode = debug
# 运行端口
HttpPort = :1300
# 用于生成用户登录的token，随便输入即可
JwtKey = nyalog
# 管理员用户一天内最大登录次数
LoginNum = 10

[DataBase]
# 数据库类型
Db = mysql
# 数据库地址
DbHost = nyalog-mysql
# 数据库端口
DbPort = 3306
# 数据库用户名-和yml中设置的一致
DbUser = nyalog
# 数据库密码-和yml中设置的一致
DbPassword = nyalog
# 数据库名称-和yml中设置的一致
DbName = nyalog

[Email]
# 邮箱服务器地址
Smtphost = 
# 邮箱服务器端口
Smtpport = 
# 邮箱服务器用户名和密码
Emailusername = 
Emailpassword = 
```

##### 前端配置

如果你在yml中修改了`nyalog-web`中的端口地址，那么需要打开`vue-client`下的Dockerfile修改注释下的端口：

```dockerfile
FROM node:18.17.1 AS builder

WORKDIR /app

COPY . /app

RUN npm config set registry https://registry.npmmirror.com

RUN npm i
RUN npm run build

FROM nginx:latest

# 修改这里
EXPOSE 80

WORKDIR /app

COPY --from=builder /app/dist /usr/share/nginx/html
COPY --from=builder /app/nginx /etc/nginx/conf.d

# 注释以使用官方镜像启动命令
# CMD ["nginx", "-g", "daemon off;"]
```

接着打开当前目录下的`config.json`设置github的client-id、api_url、sitetitle、sitedescription：

```json
{
    "domain": "设置你的域名 如https://moenya.cat",
    "client_id": "设置你的github client-id",
    "api_url": "/api/v1",
    "sitetitle": "主站标题",
    "sitedescription": "主站描述"
}
```

然后打开`index.html`修改站点标题、站点描述（和config.json中的一致即可），这样做有利于SEO：

```html
<!doctype html>
<html lang="en">
  <head>
    <script type="module" src="src/main.js"></script>
    <meta charset="UTF-8" />
    <link rel="icon" type="image/svg+xml" href="" />
    <link rel="stylesheet" href="/vue-client/public/static/css/scrollbar.css" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>主站标题</title>
  </head>
  <body>
    <div id="app"></div>
  </body>
</html>
```

最后打开`vue-client/nginx/default.conf`，按照注释下的提示进行修改：

```nginx
server {
    listen       80;
    listen  [::]:80;
    server_name  localhost;

    location /api/ {
        # 修改这里的端口(1300的地方)，改为你设置的后端端口即可
        proxy_pass http://nyalog-server:1300/api/;
    }

    location / {
        root   /usr/share/nginx/html;
        index  index.html;
        try_files $uri $uri/ /index.html;
    }
}

```

最后在`nyalog-docker-compose`目录下运行：

```bash
sudo docker compose up -d
```

此时可以访问`你的ip:你的前端端口`判断前端容器是否运行成功。访问`你的ip:你的前端端口/admin`，查看是否自动跳转到`你的ip:你的前端端口/registe`判断后端和数据库容器是否运行成功。

##### 使用web软件进行反向代理

如果能够使用ip访问成功，并且确保了3个容器正常运行，那么可以使用如nginx一类的web软件进行反向代理和https证书配置。

请注意在防火墙里将ip访问的方式给禁止掉，保证使用域名访问并且配置了https，这样才能获得NyaLog的全部内容。

### 关于数据库备份

NyaLog将会在`nyalog-docker-compose`目录下生成一个`mysql`文件夹，可以自行备份。
### SSR部署

如果你不打算使用SSR（服务端渲染）的话可以忽略这个小节的内容，但如果你对针对爬虫优化有需求的话可以按照这个小节的教程部署SSR。

#### 原理

因为所使用的SSR方案无法抓取到打包后运行在Nginx的页面，因此在SSR中其实是部署了SSR本体+NodeJS环境下运行的前端程序。

#### 更改docker compose yml文件

进入`nyalog-SSR-docker`，更改服务端口。

```yml
services:
  nyalog-ssr-web:
    container_name: nyalog-ssr-web
    build:
      context: ./vue-client
      dockerfile: Dockerfile
    ports:
    #更改你想要的端口或者使用默认
        - "1600:5173"
    
  nyalog-ssr:
    container_name: nyalog-ssr
    build:
      context: ./ssr
      dockerfile: Dockerfile
    ports:
    #更改你想要的端口或者使用默认
        - "1500:1500"
```

如果更改了这里的端口，那么需要进入`ssr`和`vue-client`中分别更改`Dcokerfile`中暴露的端口。

#### 修改前端页面配置文件

进入`vue-client`打开`config.json`，配置好title和description。（GitHub的客户ID在这里可以不配置）

#### 修改SSR配置文件

进入`ssr`文件打开`service.js`，替换第6行为`nyalog-ssr-web`中暴露的地址，并且如果你在yml中修改了`nyalog-ssr`的端口，也需要在这个文件中修改对应位置的端口。

```javascript
const express = require('./node_modules/express');
var app = express();
var spider = require("./spider.js");
var minify = require('html-minifier').minify;
app.get('*', async (req, res) => {
	let url = "你的地址，如http://127.0.0.1:1600" + req.originalUrl;
	console.log('请求的完整URL：' + url);
	let content = await spider(url).catch((error) => {
		console.log(error);
		res.send('获取html内容失败');
		return;
	});
	// 通过minify库压缩代码
    content=minify(content,{removeComments: true,collapseWhitespace: true,minifyJS:true, minifyCSS:true});
	res.send(content);
});
// 修改这里的1500为你设置的端口
app.listen(1500, () => {
	console.log('服务已启动！');
});
```

最后在`nyalog-SSR-docker`目录下运行：

```bash
docker compose up -d
```

最后使用Postman等工具访问SSR的地址（注意不要直接用浏览器访问，这样会导致SSR程序退出），查看是否能够获取被渲染的静态页面。

记得使用Nginx进行反代，并且设置防火墙让SSR和NodeJS运行的前端项目的IP地址只能被本机访问。这里提供Nginx的反代思路：

```nginx
if ($http_user_agent ~* "Baiduspider|twitterbot|facebookexternalhit|rogerbot|linkedinbot|embedly|quora link preview|showyoubot|outbrain|pinterest|slackbot|vkShare|W3C_Validator|bingbot|Sosospider|Sogou Pic Spider|Googlebot|360Spider|facebookexternalhit|TelegramBot|Discordbot") {
	proxy_pass  你的SSR地址;
}
```

## 使用

在NyaLog的初始页面，导航栏里什么都没有，此时可以按照下图的方式添加三个固定导航栏：

<div align=center><img src="https://raw.githubusercontent.com/Moe-Nya/NyaLog/main/staticimg/Static nav use.png""></div>