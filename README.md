# go-wx-sdk的测试工具，使用swagger ui对API进行测试

``` bash
king@king:~/tmp/test$ /tmp/wechat server -h
NAME:
   wechat server - 启动账户服务器

USAGE:
   wechat server [command options] [arguments...]

OPTIONS:
   --host value                         服务器主机 (default: "localhost")
   --port value                         服务器端口 (default: 13577)
   --redis_host value                   redis服务器主机 (default: "localhost")
   --redis_port value                   redis服务器端口 (default: 6379)
   --redis_db value                     redis服务器数据库 (default: 9)
   --redis_password value               redis服务器密码
   --corp_disable                       是否禁用企业号接口
   --mch_disable                        是否禁用商户接口
   --mp_disable                         是否禁用公众号接口
   --mp_token value                     公众号token (default: "abcdefg")
   --mp_appid value                     公众号app id (default: "123456798123456")
   --mp_secret value                    公众号app secret (default: "123456789123456")
   --mp_encoding_aes_key value          公众号消息加密的key (default: "abcdefg")
   --mp_verify value                    设置jssdk安全域名的授权
   --corp_id value                      企业号corpid (default: "123456789123456")
   --corp_secret value                  企业号应用secret (default: "123456789123456")
   --corp_tag value                     企业号应用tag，单个企业号必须唯一 (default: "abc")
   --corp_agent_token value             企业号 测试应用token (default: "abcdefg")
   --corp_agent_encoding_aes_key value  企业号 测试应用 消息加密的key (default: "abcdefg")
   --corp_agentid value                 企业号 测试应用ID (default: 0)
   --corp_kf_token value                企业号 客服token (default: "abcdefg")
   --corp_kf_encoding_aes_key value     企业号 客服 消息加密的key (default: "abcdefg")
   --wxpay_appid value                  微信支付appid (default: "abcdefg")
   --wxpay_mchid value                  微信支付商户号 (default: "abcdefg")
   --wxpay_apikey value                 微信支付api key，用于签名 (default: "abcdefg")
```

修改swagger文档或者前端资源需要调用如下命令重新打包
``` bash
go generate github.com/qjw/go-wx-sdk-app/doc/
go generate github.com/qjw/go-wx-sdk-app/static
```