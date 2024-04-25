import MessageAPI from '../components/message.vue'

function errmsg(code) {
    switch (code) {
        case 1001:
            window.$message.error('Token不存在');
            break;
        case 1002:
            window.$message.error('Token非法');
            break;
        case 1003:
            window.$message.error('Token解析失败');
            break;
        case 1004:
            window.$message.error('Token产生失败');
            break;
        case 3001:
            window.$message.error('2FA生成失败');
            break;
        case 3002:
            window.$message.error('2FA验证码错误');
            break;
        case 4001:
            window.$message.error('二维码生成失败');
            break;
        case 2001:
            window.$message.error('用户信息错误');
            break;
        case 2002:
            window.$message.error('用户已存在');
            break;
        case 2003:
            window.$message.error('用户不存在');
            break;
        case 2004:
            window.$message.error('用户未验证');
            break;
        case 2005:
            window.$message.error('用户检查失败');
            break;
        case 2006:
            window.$message.error('用户登录失败');
            break;
        case 2007:
            window.$message.error('用户UID错误');
            break;
        case 2008:
            window.$message.error('密码错误');
            break;
        case 2009:
            window.$message.error('IP限制');
            break;
        case 2010:
            window.$message.error('重置密码失败');
            break;
        case 2011:
            window.$message.error('查询用户失败');
            break;
        case 2012:
            window.$message.error('用户信息更改失败');
            break;
        case 5001:
            window.$message.error('邮件验证码发送失败');
            break;
        case 5002:
            window.$message.error('验证邮件验证码失败');
            break;
        case 6001:
            window.$message.error('博客信息未设置');
            break;
        case 6002:
            window.$message.error('博客信息设置失败');
            break;
        case 6003:
            window.$message.error('查询博客信息失败');
            break;
        case 7001:
            window.$message.error('文章更新失败');
            break;
        case 7002:
            window.$message.error('文章查询失败');
            break;
        case 7003:
            window.$message.error('文章删除失败');
            break;
        case 8001:
            window.$message.error('分类创建失败');
            break;
        case 8002:
            window.$message.error('分类查询失败');
            break;
        case 8003:
            window.$message.error('分类更新失败');
            break;
        case 8004:
            window.$message.error('分类删除失败');
            break;
        case 9001:
            window.$message.error('页面创建失败');
            break;
        case 9002:
            window.$message.error('页面查询失败');
            break;
        case 9003:
            window.$message.error('页面编辑失败');
            break;
        case 9004:
            window.$message.error('页面删除失败');
            break;
        case 10001:
            window.$message.error('创建登录信息失败');
            break;
        case 10002:
            window.$message.error('登录信息失败');
            break;
        case 10003:
            window.$message.error('发送评论失败');
            break;
        case 10004:
            window.$message.error('查询评论失败');
            break;
        case 10005:
            window.$message.error('删除评论失败');
            break;
        case 11001:
            window.$message.error('新增findme失败');
            break;
        case 11002:
            window.$message.error('查询findme失败');
            break;
        case 11003:
            window.$message.error('修改findme失败');
            break;
        case 11004:
            window.$message.error('删除findme失败');
            break;
        case 12001:
            window.$message.error('创建导航栏元素失败');
            break;
        case 12002:
            window.$message.error('查询导航栏元素失败');
            break;
        case 12003:
            window.$message.error('编辑导航栏元素失败');
            break;
        case 12004:
            window.$message.error('删除导航栏元素失败');
            break;
        case 13001:
            window.$message.error('创建友链失败');
            break;
        case 13002:
            window.$message.error('查询友链失败');
            break;
        case 13003:
            window.$message.error('修改友链失败');
            break;
        case 13004:
            window.$message.error('删除友链失败');
            break;
        default:
            window.$message.error('未知错误');
            break;
    }
}

export default errmsg