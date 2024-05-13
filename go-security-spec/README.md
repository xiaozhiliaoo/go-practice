# Go安全规范

## cos安全规范

### 用于前端直传 COS 的临时密钥安全指引

https://cloud.tencent.com/document/product/436/40265

权限最小原则：
资源超范围限定
操作超范围限定
资源和操作超范围限定：通过多个 statement 的形式组合
越权获取临时密钥：

## 规范

```text
切片长度校验：防止程序panic
nil指针判断：结构体，结构体Unmarshal
整数安全：防止程序panic
make分配长度校验
禁止SetFinalizer和循环引用同时使用
禁止重复释放Channel
确保每个协程都能退出
不使用unsafe包
不使用slice作为函数入参
文件路径穿越检查
文件访问权限控制
shell命令执行检查
网络通信使用TLS方式
TLS启动证书验证
控制敏感信息访问
控制敏感信息输出
控制敏感信息存储
异常处理和日志记录
不使用硬编码密码和密钥
密钥存储安全
不使用弱密码算法
使用regexp进行正则表达式匹配
按类型进行数据校验
SQL语句默认使用预编译并绑定变量
资源请求过滤验证
模版渲染过滤验证
跨站资源共享CORS限制请求来源
设置正确的HTTP响应包类型
添加安全响应头
外部输入拼接到HTTP响应头中需要过滤
外部输入拼接到response页面前进行编码处理
安全维护Session信息
CSRF防护
默认鉴权
禁止在闭包中直接调用循环变量
禁止并发写map
确保并发安全
依赖库来源安全
```

## 学习笔记

shellescape检查shell

Strict-Transport-Security参数 强制客户端始终通过HTTPS

&tls.Config{InsecureSkipVerify: false} 控制客户端是否跳过服务器证书的验证

crypto/des，crypto/md5，crypto/sha1，crypto/rc4 弱密码算法，不安全的。crypto/aes、crypto/sha256是安全的

回溯引用(Backreferences)和前后查找(lookaround)，先行断言-向前查找(lookahead)和后行断言-向后查找(lookbehind)

https://pkg.go.dev/regexp

sec-api：安全地做HTTP请求

xxe（XML External Entity Injection）漏洞

https://github.com/go-playground/validator validate.Var(val, "gte=1,lte=100") 限制必须是1-100的正整数

Access-Control-Allow-Origin：同源策略进行保护

X-Content-Type-Options: nosniff 表示浏览器不应该尝试猜测资源的MIME类型，而是应该严格遵循服务器提供的内容类型。这有助于防止浏览器执行恶意脚本，
因为攻击者可能会尝试伪装脚本的MIME类型，以便在不受信任的上下文中执行

X-Frame-Options：HTTP响应头，用于防止网页被嵌入到其他网站的<iframe>、<frame>或<object>标签中，从而防止点击劫持攻击（Clickjacking）
DENY：表示网页不能被嵌入到任何其他网站的<iframe>、<frame>或<object>标签中。
SAMEORIGIN：表示网页只能被嵌入到同源的网站的<iframe>、<frame>或<object>标签中。
ALLOW-FROM uri：表示网页只能被嵌入到指定URI的网站的<iframe>、<frame>或<object>标签中。

CSRF：Cross-Site Request Forgery 跨站请求伪造
陌生链接不要随便点 https://blog.poetries.top/browser-working-principle/guide/part6/lesson34.html

CORS：Cross-Origin Resource Sharing 跨站资源共享

https://github.com/gorilla/csrf csrf_toke值 csrf.Protect 了确保CSRF保护的有效性，你需要在表单中包含CSRF令牌

tRPC bkn插件，框架级CSRF漏洞对抗机制

防御CSRF：Token、验证码机制、浏览器SameSite Cookies功能

CSRF防御最好是：referer检查 + csrf_token

白名单

最小权限原则

Horus 开源组件风险检测（依赖包检测）

## 问题

重复释放channel，为什么会造成DOS攻击？

