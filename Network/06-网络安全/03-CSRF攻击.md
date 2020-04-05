---
author: "kuonz"
draft: false
title: "CSRF攻击"
date: 2020-04-05
categories: ["网络安全"]
---
  
### CSRF攻击是什么

`CSRF` 是 `Cross-Site Request Forgery` 的缩写（也可缩写为 `XSRF`），中文翻译为跨站请求伪造，也被称为 `one click attack` 或 `session riding`

`CSRF` 攻击指的是攻击者伪造出一个请求，受害者进行该请求时，服务器会误认为这是合法用户做出的操作，所以不会进行阻止和拦截，导致受害者无意间做出了损害利益的操作

因为大多数网站的权限管理都是通过 Cookie + Session 完成的，所以受害者进行请求时，会将浏览器上保存的 Cookie 一同发送给浏览器，导致伪造请求成功发送



### CSRF特点

`CSRF` 破坏力依赖于受害者的权限，如果受害者是管理员，则 `CSRF` 危害更大

`CSRF` 攻击难度大，不流行，但也很难防范，称为**沉睡的巨人**



### CSRF例子

1. `Alice` 登录了一个金融网站 `mybank.com`

2. `Bob` 知道 `mybank.com`，且发现网站的转账功能有 `CSRF` 漏洞

3. 于是 `Bob` 在 `myblog.com` 上发表了一条 `blog`，这个 `blog` 支持图片自定义功能

4. 于是 `Bob` 插入了这么一行代码

   ```HTML
   <img src="http://mybank.com/transferMoney.jsp?to=Bob&cash=3000 />
   ```

5. `Alice` 浏览到该博客，浏览器请求图片，自动请求了图片的地址，又由于 `Alice` 的浏览器上存有`mybank.com` 的会话 `Cookie`，于是不知不觉中 `Alice` 向 `Bob` 账户转账了300元，而她自己毫不知情



### 解决方法

#### 操作验证

* 对于敏感重要操作添加确认操作
* 二次验证
* 重新认证

#### Referer验证

`Referer` 指的是页面请求来源，指只接受本站的请求，服务器才做响应，如果不是，就拦截不处理

#### Token验证

* 用户登录时，服务器生成一个不可预知的 `CSRF Token`，该 `Token` 放在用户 `Session` 中，并传送给浏览器`Cookie`
* 在任意一个需要保护的页面中，增加一个隐藏的字段来存放该 `Token`，对于需要保护的 `URL`，增加一个参数来存放该 `Token`，或者把 `Token` 隐藏在 `http` 的请求首部中
* 提交此请求时，在服务端检查提交的 `Token` 与用户 `Session` 中的 `Token` 是否一致，如果一致则继续处理，如果不一致则返回一个错误信息给用户
* 当用户退出或失效时，则从 `Session` 删除该 `Token`



### XSS和CSRF区别

* `XSS` 是利用合法用户获取其信息，而 `CSRF` 是伪造成合法用户发起请求

* `CSRF` 相当于是 `XSS` 的基础版，`CSRF` 能做到的 `XSS` 都可以做到
* `XSS` 主要指向客户端，而 `CSRF` 针对服务端