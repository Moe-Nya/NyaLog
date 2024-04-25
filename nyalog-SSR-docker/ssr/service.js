const express = require('./node_modules/express');
var app = express();
var spider = require("./spider.js");
var minify = require('html-minifier').minify;
app.get('*', async (req, res) => {
	let url = "你的地址，如http://127.0.0.1:1080" + req.originalUrl;
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
app.listen(1500, () => {
	console.log('服务已启动！');
});