package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"gopkg.in/redis.v5"
	"net/http"
	"strconv"
	"strings"
	"github.com/qjw/go-wx-sdk/mp"
	"github.com/qjw/go-wx-sdk/corp"
	"github.com/qjw/go-wx-sdk/mch"
	"github.com/qjw/go-wx-sdk/cache"
	"github.com/qjw/go-wx-sdk/small"
	"github.com/qjw/go-wx-sdk-app/api/gateway"
	"github.com/qjw/go-wx-sdk-app/doc"
	"github.com/qjw/go-wx-sdk-app/static"
	"github.com/qjw/go-wx-sdk-app/const_var"
	apicorp "github.com/qjw/go-wx-sdk-app/api/corp"
	apimch "github.com/qjw/go-wx-sdk-app/api/mch"
	apimp "github.com/qjw/go-wx-sdk-app/api/mp"
	apism "github.com/qjw/go-wx-sdk-app/api/small"
	"github.com/qjw/kelly"
	"github.com/qjw/kelly/middleware"
	"github.com/qjw/kelly/middleware/swagger"
	"runtime"
	"reflect"
)

func GetMpFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "mp_token",
			Usage:  "公众号token",
			Value:  "abcdefg",
		},
		cli.StringFlag{
			Name:   "mp_appid",
			Usage:  "公众号app id",
			Value:  "123456798123456",
		},
		cli.StringFlag{
			Name:   "mp_secret",
			Usage:  "公众号app secret",
			Value:  "123456789123456",
		},
		cli.StringFlag{
			Name:   "mp_encoding_aes_key",
			Usage:  "公众号消息加密的key",
			Value:  "abcdefg",
		},
		cli.StringFlag{
			Name:   "mp_verify",
			Usage:  "设置jssdk安全域名的授权",
			Value:  "",
		},
	}
}


func GetCorpFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "corp_id",
			Usage:  "企业号corpid",
			Value:  "123456789123456",
		},
		cli.StringFlag{
			Name:   "corp_secret",
			Usage:  "企业号应用secret",
			Value:  "123456789123456",
		},
		cli.StringFlag{
			Name:   "corp_tag",
			Usage:  "企业号应用tag，单个企业号必须唯一",
			Value:  "abc",
		},
	}
}

func GetCorpAgentFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "corp_agent_token",
			Usage:  "企业号 测试应用token",
			Value:  "abcdefg",
		},
		cli.StringFlag{
			Name:   "corp_agent_encoding_aes_key",
			Usage:  "企业号 测试应用 消息加密的key",
			Value:  "abcdefg",
		},
		cli.Int64Flag{
			Name:   "corp_agentid",
			Usage:  "企业号 测试应用ID",
			Value:  0,
		},
	}
}


func GetCorpKfFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "corp_kf_token",
			Usage:  "企业号 客服token",
			Value:  "abcdefg",
		},
		cli.StringFlag{
			Name:   "corp_kf_encoding_aes_key",
			Usage:  "企业号 客服 消息加密的key",
			Value:  "abcdefg",
		},
	}
}

func GetWxPayFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "wxpay_appid",
			Usage:  "微信支付appid",
			Value:  "abcdefg",
		},
		cli.StringFlag{
			Name:   "wxpay_mchid",
			Usage:  "微信支付商户号",
			Value:  "abcdefg",
		},
		cli.StringFlag{
			Name:   "wxpay_apikey",
			Usage:  "微信支付api key，用于签名",
			Value:  "abcdefg",
		},
	}
}

func GetSmallFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "sm_token",
			Usage:  "小程序token",
			Value:  "abcdefg",
		},
		cli.StringFlag{
			Name:   "sm_appid",
			Usage:  "小程序app id",
			Value:  "123456798123456",
		},
		cli.StringFlag{
			Name:   "sm_secret",
			Usage:  "小程序app secret",
			Value:  "123456789123456",
		},
		cli.StringFlag{
			Name:   "sm_encoding_aes_key",
			Usage:  "小程序消息加密的key",
			Value:  "abcdefg",
		},
	}
}

func combineSliceArray(first []cli.Flag, lst ...[]cli.Flag) []cli.Flag {

	for _, item := range lst {
		first = append(first, item...)
	}
	return first
}

var serverCmd = cli.Command{
	Name:  "server",
	Usage: "启动账户服务器",
	Action: func(c *cli.Context) {
		if err := server(c); err != nil {
			log.Fatal(err)
		}
	},
	Flags: combineSliceArray([]cli.Flag{
		cli.StringFlag{
			Name:   "host",
			Usage:  "服务器主机",
			Value:  "localhost",
		},
		cli.IntFlag{
			Name:   "port",
			Usage:  "服务器端口",
			Value:  13577,
		},
		cli.StringFlag{
			Name:   "redis_host",
			Usage:  "redis服务器主机",
			Value:  "localhost",
		},
		cli.IntFlag{
			Name:   "redis_port",
			Usage:  "redis服务器端口",
			Value:  6379,
		},
		cli.IntFlag{
			Name:   "redis_db",
			Usage:  "redis服务器数据库",
			Value:  9,
		},
		cli.StringFlag{
			Name:   "redis_password",
			Usage:  "redis服务器密码",
			Value:  "",
		},
		cli.BoolFlag{
			Name:   "corp_disable",
			Usage:  "是否禁用企业号接口",
		},
		cli.BoolFlag{
			Name:   "mch_disable",
			Usage:  "是否禁用商户接口",
		},
		cli.BoolFlag{
			Name:   "mp_disable",
			Usage:  "是否禁用公众号接口",
		},
		cli.BoolFlag{
			Name:   "sm_disable",
			Usage:  "是否禁用小程序接口",
		},
	},
		GetMpFlags(),
		GetCorpFlags(),
		GetCorpAgentFlags(),
		GetCorpKfFlags(),
		GetWxPayFlags(),
		GetSmallFlags(),
	),
}

var redisClient *redis.Client = nil

/**
初始化Redis
*/
func InitRedis(c *cli.Context) {
	log.Print("start to init redis")
	redisClient = redis.NewClient(&redis.Options{
		Addr:     c.String("redis_host") + ":" + strconv.Itoa(c.Int("redis_port")),
		Password: c.String("redis_password"),
		DB:       c.Int("redis_db"),
	})
	if err := redisClient.Ping().Err(); err != nil {
		log.Fatal("failed to connect redis")
	}
	log.Print("init redis success")
}

func NewWxJsMiddleware(verify string, prefix string) kelly.HandlerFunc {
	verify_file := ""
	verifies := strings.Split(verify, ":")
	if len(verify) > 0 {
		if len(verifies) != 2 {
			panic("invalid verify")
		} else {
			verify_file = verifies[0]
			verify = verifies[1]
		}
	}

	return func(c *kelly.Context) {
		uri := c.Request().URL.RequestURI()
		if strings.HasPrefix(uri, prefix) &&
			"/"+verify_file == uri {
			c.WriteString(http.StatusOK, verify)
		} else {
			c.InvokeNext()
		}
	}
}

func initMp(c *cli.Context, grouter kelly.Router) {
	context := mp.NewContext(
		&mp.Config{
			Token:          c.String("mp_token"),
			AppID:          c.String("mp_appid"),
			AppSecret:      c.String("mp_secret"),
			EncodingAESKey: c.String("mp_encoding_aes_key"),
		},
		cache.NewCache(redisClient),
	)

	// js安全域名注册
	grouter.Use(NewWxJsMiddleware(c.String("mp_verify"), "/MP_verify"))

	apimp.InitializeApiRoutes(grouter, context)
	gateway.InitializeApiRoutes(grouter.Group("/gateway"), context)

	fs := http.FileServer(static.AssetFS())
	grouter.GET("/html/*filepath",
		func(c *kelly.Context) {
			fs.ServeHTTP(c, c.Request())
		},
	)
}

func initCorp(c *cli.Context, grouter kelly.Router) {
	agentConfig := &corp.AgentConfig{
		AgentID:        c.Int64("corp_agentid"),
		Token:          c.String("corp_agent_token"),
		EncodingAESKey: c.String("corp_agent_encoding_aes_key"),
	}
	kfConfig := &corp.KfConfig{
		Token:          c.String("corp_kf_token"),
		EncodingAESKey: c.String("corp_kf_encoding_aes_key"),
	}
	context := corp.NewContext(
		&corp.Config{
			CorpID:     c.String("corp_id"),
			CorpSecret: c.String("corp_secret"),
			Tag:        c.String("corp_tag"),
		},
		cache.NewCache(redisClient),
	)

	apicorp.InitializeApiRoutes(grouter, context, agentConfig)

	gateway.InitializeAgentApiRoutes(grouter.Group("/agent_gateway"),
		context,
		agentConfig,
		kfConfig,
	)
}

func initMch(c *cli.Context, grouter kelly.Router) {
	mchConfig := &mch.Config{
		AppID:c.String("wxpay_appid"),
		MchID:c.String("wxpay_mchid"),
		ApiKey:c.String("wxpay_apikey"),
	}
	context := mch.NewContext(mchConfig)
	apimch.InitializeApiRoutes(grouter, context)
}

func initSM(c *cli.Context, grouter kelly.Router) {
	context := small.NewContext(
		&small.Config{
			Token:          c.String("sm_token"),
			AppID:          c.String("sm_appid"),
			AppSecret:      c.String("sm_secret"),
			EncodingAESKey: c.String("sm_encoding_aes_key"),
		},
		cache.NewCache(redisClient),
	)

	// go的gc不会释放对象smApi
	smApi := small.NewSmallApi(context)
	gateway.InitializeSmApiRoutes(grouter.Group("/gateway"), context, smApi)
	apism.InitializeSmallApiRoutes(grouter,context,smApi)
}

func server(c *cli.Context) error {
	InitRedis(c)

	middlewares := []kelly.HandlerFunc{middleware.Version(const_var.Version)}
	if c.GlobalBool("debug") {
		middlewares = append(middlewares,middleware.Cors(&middleware.CorsConfig{
			AllowAllOrigins: true,
			AllowHeaders:    []string{"X-APP-ID", "X-APP-HASH", "Origin", "Content-Length", "Content-Type"},
			AllowMethods:    []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		}))
	}
	r := kelly.NewClassic(middlewares...)
	r.OPTIONS("/*path", func(c *kelly.Context) {
		c.ResponseStatusOK()
	})

	// 增加全局的endpoint钩子
	r.GlobalAnnotation(func(c *kelly.AnnotationContext) {
		handle := c.HandlerFunc()
		name := runtime.FuncForPC(reflect.ValueOf(handle).Pointer()).Name()
		log.Printf("register [%7s|%2d]%s%s ---- %s",
			c.Method(), c.HandleCnt(), c.R().Path(), c.Path(), name)
	})


	// swagger之后再设置认证
	grouter := r.Group(const_var.ApiPrefix)

	// swagger
	swagger.InitializeApiRoutes(r,
		&swagger.Config{
			BasePath:const_var.ApiPrefix,
			Title: "微信SDK测试工具",
			Description: "微信SDK测试工具",
			DocVersion: "0.1",
			SwaggerUiUrl: "http://swagger.qiujinwu.com",
			SwaggerUrlPrefix: "doc",
			Debug: c.GlobalBool("debug"),
		},
		func(key string) ([]byte, error) { return doc.Asset(key) },
	)

	//--------------------------------------------------------------------------------------------------------------
	// 公众号
	if !c.Bool("mp_disable"){
		initMp(c, grouter)
	}
	//--------------------------------------------------------------------------------------------------------------
	// 企业号
	if !c.Bool("corp_disable"){
		initCorp(c, grouter)
	}
	//--------------------------------------------------------------------------------------------------------------
	// 微信支付
	if !c.Bool("mch_disable"){
		initMch(c, grouter)
	}
	//--------------------------------------------------------------------------------------------------------------
	// 小程序
	if !c.Bool("sm_disable"){
		initSM(c, grouter)
	}

	r.Run(c.String("host") + ":" + strconv.Itoa(c.Int("port")))
	return nil
}
