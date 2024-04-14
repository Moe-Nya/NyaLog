<div align=center><img src="https://raw.githubusercontent.com/Moe-Nya/NyaLog/main/staticimg/NyaLog.PNG" style="height: 300px; width:320px;"></div>

# NyaLog

> 🐱 MoeNya's lovely blog.

## 介绍

NyaLog是一个由Golang的Gin框架和Vue.js制作的前后端分离、自适应响应式博客。作者一直以来都有一个跳出博客框架的想法，这次将它付诸现实了~

目前NyaLog还存在一些缺陷：如果导航栏、FindMe还不能支持自排序（在开发的时候没有想到），在未来将会对这些使用上影响体验的地方进行修改。

[开发文档](https://github.com/Moe-Nya/NyaLog/blob/main/document/readme.md)

## 使用

### 下载

在package中下载对应版本。

### Docker运行

#### MySql配置

请自行在系统中安装并且运行好MySql，并且为nyalog创建一个数据库。

这里以Ubuntu为例，安装MySQL：

```bash
sudo apt-get install mysql-server
```

接下来初始化MySQL：

```bash
sudo mysql_secure_installation
```

选择：1. 不进行密码强校验 2.删除匿名用户 3. 删除测试数据库。

检查数据库状态：

```bash
systemctl status mysql.service
```

此时数据库的用户名root，密码为空，直接进入数据库修改密码：

```bash
sudo mysql -uroot -p
```

当提示输入密码的时候直接回车。

修改密码：

```mysql
ALTER USER ‘root’@‘localhost’ IDENTIFIED BY ‘密码’ PASSWORD EXPIRE NEVER;
```

最后进入数据库为NyaLog创建一个数据库。

#### 后端

首先选择一个你喜欢的路径，创建后端文件夹`yourPath/nyalog-backend`。

将文件包中的`gin-server`复制到`nyalog-backend`中，打开`gin-server/gin-blog-server/config/config.ini`，按照注释代码填写好相关信息。

```ini
[System]
# 后端访问地址
Domain = https://yourIP:youPort
# 密码最小长度
PasswordMinLen = 10
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
AppMode = release
# 运行端口
HttpPort = :1400
# 用于生成用户登录的token，随便输入即可
JwtKey = randon
# 管理员用户一天内最大登录次数
LoginNum = 10

[DataBase]
# 数据库类型
Db = mysql
# 数据库地址
DbHost = localhost
# 数据库端口
DbPort = 3306
# 数据库用户名
DbUser = username
# 数据库密码
DbPassword = userpassword
# 数据库名称
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

填写好信息后回到`gin-server`中，建立名为`Dockerfile`的文本文件，并且写入以下内容：

```dockerfile
FROM golang:1.21.6

RUN go env -w GO111MODULE=on
# 如果你所在的地区是中国大陆，不妨使用中国大陆镜像源
# RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

WORKDIR /nyalog
COPY . /nyalog

RUN go build .

EXPOSE 1300

ENTRYPOINT ["./NyaLog"]
```

随后建立Docker镜像：

```bash
sudo docker build -t nyalogbackend:v1.0.1 .
```

输入`sudo docker images`可以查看镜像是否建立成功。

接着运行容器：

```bash
sudo docker run -ti -d --name nyalogbackend -p 1300:1300 nyalogbackend:v1.0.1
```

输入`sudo docker ps`可以查看容器是否已经运行起来。

#### 前端

前端因为需要使用者自己去设置一些东西，比如index.html中的站点信息，github和后端接口，所以没有办法直接打包，需要自行下载node.js和npm。

在路径中创建`yourPath/nyalog-frontend`文件夹，将下载包中的`gin-server`复制到`nyalog-frontend`中，在`index.html`中按照自己的SEO习惯进行设置，将站点icon的静态图片资源放到`public/img`中；在`config.json`中配置好GitHub第三方登录权限的`client_id`和后端接口入口`api_url`”：

```html
<!doctype html>
<html lang="en">
  <head>
    <script type="module" src="src/main.js"></script>
    <meta charset="UTF-8" />
    <link rel="icon" type="image/svg+xml" href="img/siteicon.png" />
    <link rel="stylesheet" href="/vue-client/public/static/css/scrollbar.css" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>your title</title>
  </head>
  <body>
    <div id="app"></div>
  </body>
</html>

```

```json
{
    "client_id": "your client_id",
    "api_url": "http://localhost:8080/api/v1"
}
```

确认好设置无误，并且保证Node.js和npm安装完成：

<span style="color: red">**Node.js版本不低于v18.17.1！**</span>

执行打包命令：

```bash
npm run build
```

打包后会新生成一个`dist`文件夹，在该文件夹同级目录下新建`Dockerfile`文本文件：

```dockerfile
FROM nginx

EXPOSE 80

COPY /dist /usr/share/nginx/html

ENTRYPOINT nginx -g "daemon off;"
```

建立前端镜像：

```bash
sudo docker build -t nyalogfrontend:v1.0.1 .
```

使用`sudo docker images`可以查看镜像是否创建成功。

创建容器：

```bash
sudo docker run -ti -d --name nyalogfrontend -p 1400:80 nyalogfrontend:v1.0.1
```

使用`sudo docker ps`可以查看容器是否成功创建并运行。