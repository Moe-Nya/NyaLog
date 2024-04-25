const puppeteer = require('./node_modules/puppeteer');
const MAX_WSE = 2; //启动几个浏览器 
let WSE_LIST = []; //存储browserWSEndpoint列表
//负载均衡
(async () => {
	for (var i = 0; i < MAX_WSE; i++) {
		const browser = await puppeteer.launch({
            //无头模式
			headless: true,
            //参数
			args: [
				'--disable-gpu',
				'--disable-dev-shm-usage',
				'--disable-setuid-sandbox',
				'--no-first-run',
				'--no-sandbox',
				'--no-zygote',
				'--single-process'
			],
			//一般不需要配置这条，除非启动一直报错找不到谷歌浏览器
			//executablePath:'C:/Program //Files/Google/Chrome/Application/chrome.exe'
		});
		let browserWSEndpoint = await browser.wsEndpoint();
		WSE_LIST.push(browserWSEndpoint);
	}
})();

module.exports = WSE_LIST