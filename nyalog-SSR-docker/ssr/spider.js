const puppeteer = require('./node_modules/puppeteer');//由于目录不一致，所以使用的是绝对路径
const WSE_LIST = require('./puppeteer-pool.js'); //这里注意文件的路径和文件名
const spider = async (url) => {
	let tmp = Math.floor(Math.random() * WSE_LIST.length);
	//随机获取浏览器
	let browserWSEndpoint = WSE_LIST[tmp];
	//连接
	const browser = await puppeteer.connect({
		browserWSEndpoint
	});
	
	//打开一个标签页
	var page = await browser.newPage();
	
	// Intercept network requests.
	await page.setRequestInterception(true);

    page.on('request', req => {
      // Ignore requests for resources that don't produce DOM
      // (images, stylesheets, media).
      const whitelist = ['document', 'script', 'xhr', 'fetch'];
      if (!whitelist.includes(req.resourceType())) {
        return req.abort();
      }
      
      // Pass through all other requests.
      req.continue();
	page.setJavaScriptEnabled(true);
    });
	
	//打开网页
	await page.goto(url, {
		timeout: 20000, //连接超时时间，单位ms
		waitUntil: 'networkidle0' //网络空闲说明已加载完毕
	});
	
	//获取渲染好的页面源码。不建议使用await page.content();evaluate获取页面，因为在我测试中发现，页面还没有完全加载。就获取到了。页面源码不完整。也就是动态路由没有加载。vue路由也配置了history模式
	let html = await page.content(() => {
		return document.getElementsByTagName('html')[0].outerHTML;
	});
	await page.close();
	return html;
}

module.exports = spider;