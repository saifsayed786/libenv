package libenv

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/golobby/container/v3"
	"github.com/joho/godotenv"
)

var ENV_MAP = make(map[string]string)

func (env *Env) Init() error {
	env.STAGE = os.Getenv("STAGE")
	if env.STAGE == "" {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println("No .Env file found")
		}
		fmt.Println(".Env File Loaded")
	}

	env.JWT_ACCESS_TOKEN_TIME_TO_EXPIRE_IN_DAY = 1
	env.JWT_REFRESH_TOKEN_TIME_TO_EXPIRE_IN_DAYS = 30
	env.MAX_OTP_ATTEMPT = 3
	env.MAX_PASSWORD_ATTEMPT = 3
	env.MAX_MPIN_ATTEMPT = 3
	env.MAX_WATCHLIST_Limit = 5
	env.ANDROIDMINVERSION = "0.1.20"
	env.ANDROIDMAXVERSION = "0.1.21"

	env.IOSMINVERSION = "0.1.23"
	env.IOSMAXVERSION = "0.1.24"
	env.IPO_LOGIN_TICKER_SECS = 10
	env.IPO_HEARTBEAT_TICKER_MIN = 55
	env.IPO_LOGIN_RETRIES_MAX_COUNT = 3
	env.IPO_DB_CHANNEL_SIZE = 1000
	env.MUTF_DB_CHANNEL_SIZE = 1000

	//
	env.DI_DB_PAPERTRADE = "papertrade"
	env.DI_DB_JWTAUTH = "jwtauth"

	env.DI_DB_FACETRADE = "facetrade"
	env.DI_DB_IPO = "ipo"
	env.DI_DB_PAYMENT = "payment"
	env.DI_DB_STOCKAL = "stockal"

	env.DI_DB_ADMINMASTER = "adminmaster"
	env.DI_DB_EDIS = "edis"
	env.DI_DB_CONTRACTS = "contractmaster"
	env.DI_DB_CLINTMASTER = "clientmaster"
	env.DI_DB_ADVISORY = "advisory"
	env.DI_DB_ZEROLOGS = "zerologs"
	env.DI_DB_ACCORD = "accord"
	env.DI_DB_WATCHLIST = "watchlist"
	env.DI_DB_SWITCHACCOUNT = "switchaccount"
	env.DI_DB_SettingService = "settingservice"
	env.DI_DB_TRACKER = "tracker"
	env.DI_DB_MUTUALFUND = "mutualfund_dev"
	env.DI_DB_WHATSAPP = "whatsappdb"
	env.DI_DB_NOTIS = "notis"

	env.IS_FINVU = false

	env.STAGE = os.Getenv("STAGE")
	if env.STAGE == "" {
		env.STAGE = "local"
	}
	if env.STAGE == "" {
		env.DEBUG = *(flag.Bool("debug", true, "sets log level to debug"))
	} else {
		env.DEBUG = *(flag.Bool("debug", false, "sets log level to debug"))
	}

	env.HOSTTYPE = os.Getenv("HOSTTYPE")
	if env.HOSTTYPE == "" {
		env.HOSTTYPE = "AWS"
	}

	env.CLUSTER = os.Getenv("CLUSTER")
	if env.CLUSTER == "" {
		env.CLUSTER = "nuuu-trade-cluster"
	}

	env.EXPOSEPORT = os.Getenv("EXPOSEPORT")
	if env.EXPOSEPORT == "" {
		env.EXPOSEPORT = "8087"
	}

	env.APPLABEL = os.Getenv("APPLABEL")

	env.VOLUMNE_PATH = os.Getenv("VOLUMNE_PATH")
	if env.VOLUMNE_PATH == "" {
		env.VOLUMNE_PATH = "D:/GitHub/KubeVolumeData/"
	}

	env.MSQL_USERNAME = os.Getenv("MSQL_USERNAME")
	if env.MSQL_USERNAME == "" {
		env.MSQL_USERNAME = "root"
	}

	env.JWT_SECRET = os.Getenv("JWT_SECRET")
	if env.JWT_SECRET == "" {
		env.JWT_SECRET = "this is a Sample Secret"
	}

	env.MSQL_USERNAME = os.Getenv("MSQL_USERNAME")
	if env.MSQL_USERNAME == "" {
		env.MSQL_USERNAME = "root"
	}

	env.MSQL_PASSWORD = os.Getenv("MSQL_PASSWORD")
	if env.MSQL_PASSWORD == "" {
		env.MSQL_PASSWORD = "Aa1!"
	}
	env.MSQL_SERVICENAME = os.Getenv("MSQL_SERVICENAME")
	if env.MSQL_SERVICENAME == "" {
		env.MSQL_SERVICENAME = "localhost"
	}
	env.MSQL_PORT = os.Getenv("MSQL_PORT")
	if env.MSQL_PORT == "" {
		env.MSQL_PORT = "3306"
	}
	env.MSQL_Client_DBNAME = os.Getenv("MSQL_Client_DBNAME")
	if env.MSQL_Client_DBNAME == "" {
		env.MSQL_Client_DBNAME = "clientmaster"
	}

	env.MSQL_contractmaster_DBNAME = os.Getenv("MSQL_contractmaster_DBNAME")
	if env.MSQL_contractmaster_DBNAME == "" {
		env.MSQL_contractmaster_DBNAME = "contractmaster"
	}

	env.MSQL_edis_DBNAME = os.Getenv("MSQL_edis_DBNAME")
	if env.MSQL_edis_DBNAME == "" {
		env.MSQL_edis_DBNAME = "edis"
	}

	env.MSQL_facetrade_DBNAME = os.Getenv("MSQL_facetrade_DBNAME")
	if env.MSQL_facetrade_DBNAME == "" {
		env.MSQL_facetrade_DBNAME = "facetrade"
	}

	env.MSQL_tracker_DBNAME = os.Getenv("MSQL_tracker_DBNAME")
	if env.MSQL_tracker_DBNAME == "" {
		env.MSQL_tracker_DBNAME = "tracker"
	}

	env.MSQL_ADMIN_MASTER_DBNAME = os.Getenv("MSQL_ADMIN_MASTER_DBNAME")
	if env.MSQL_ADMIN_MASTER_DBNAME == "" {
		env.MSQL_ADMIN_MASTER_DBNAME = "adminmaster"
	}

	env.MSQL_IFSC_DBNAME = os.Getenv("MSQL_IFSC_DBNAME")
	if env.MSQL_IFSC_DBNAME == "" {
		env.MSQL_IFSC_DBNAME = "IFSC"
	}

	env.MSQL_papertrade_DBNAME = os.Getenv("MSQL_papertrade_DBNAME")
	if env.MSQL_papertrade_DBNAME == "" {
		env.MSQL_papertrade_DBNAME = "papertrade"
	}

	env.MSQL_IPO_DBNAME = os.Getenv("MSQL_ipo_DBNAME")
	if env.MSQL_IPO_DBNAME == "" {
		env.MSQL_IPO_DBNAME = "ipo"
	}
	env.RAZORPAY_KEY_TPV = os.Getenv("RAZORPAY_KEY_TPV")
	if env.RAZORPAY_KEY_TPV == "" {
		env.RAZORPAY_KEY_TPV = "rzp_live_ijRw6ljftEJ2iX"
	}

	env.RAZORPAY_SECRET_TPV = os.Getenv("RAZORPAY_SECRET_TPV")
	if env.RAZORPAY_SECRET_TPV == "" {
		env.RAZORPAY_SECRET_TPV = "j2JEpyiEeeh6JQHHF6x5ujRq"
	}

	env.RAZORPAY_NAME_TPV = os.Getenv("RAZORPAY_NAME_TPV")
	if env.RAZORPAY_NAME_TPV == "" {
		env.RAZORPAY_NAME_TPV = "Nu Investors"
	}

	env.RAZORPAY_KEY_NON_TPV = os.Getenv("RAZORPAY_KEY_NON_TPV")
	if env.RAZORPAY_KEY_NON_TPV == "" {
		env.RAZORPAY_KEY_NON_TPV = "rzp_live_q8gUCOxfHIbCkb"
	}

	env.RAZORPAY_SECRET_NON_TPV = os.Getenv("RAZORPAY_SECRET_NON_TPV")
	if env.RAZORPAY_SECRET_NON_TPV == "" {
		env.RAZORPAY_SECRET_NON_TPV = "nxE9uMF93I2le22E19Lvbaqg"
	}

	env.RAZORPAY_NAME_NON_TPV = os.Getenv("RAZORPAY_NAME_NON_TPV")
	if env.RAZORPAY_NAME_NON_TPV == "" {
		env.RAZORPAY_NAME_NON_TPV = "MANGAL KESHAV"
	}

	env.MSQL_payment_DBNAME = os.Getenv("MSQL_payment_DBNAME")
	if env.MSQL_payment_DBNAME == "" {
		env.MSQL_payment_DBNAME = "payment"
	}

	env.MSQL_jwtauth_DBNAME = os.Getenv("MSQL_jwtauth_DBNAME")
	if env.MSQL_jwtauth_DBNAME == "" {
		env.MSQL_jwtauth_DBNAME = "jwtauth"
	}

	env.MSQL_smartsearch_DBNAME = os.Getenv("MSQL_smartsearch_DBNAME")
	if env.MSQL_smartsearch_DBNAME == "" {
		env.MSQL_smartsearch_DBNAME = "smarrtsearch"
	}

	env.MSQL_auth_DBNAME = os.Getenv("MSQL_auth_DBNAME")
	if env.MSQL_auth_DBNAME == "" {
		env.MSQL_auth_DBNAME = "auth"
	}

	env.MSQL_gowsproxy_DBNAME = os.Getenv("MSQL_gowsproxy_DBNAME")
	if env.MSQL_gowsproxy_DBNAME == "" {
		env.MSQL_gowsproxy_DBNAME = "gowsproxy"
	}

	env.MSQL_stocksearch_DBNAME = os.Getenv("MSQL_stocksearch_DBNAME")
	if env.MSQL_stocksearch_DBNAME == "" {
		env.MSQL_stocksearch_DBNAME = "stocksearch"
	}

	env.MSQL_techexcel_DBNAME = os.Getenv("MSQL_techexcel_DBNAME")
	if env.MSQL_techexcel_DBNAME == "" {
		env.MSQL_techexcel_DBNAME = "techexcel"
	}

	env.MSQL_advisory_DBNAME = os.Getenv("MSQL_advisory_DBNAME")
	if env.MSQL_advisory_DBNAME == "" {
		env.MSQL_advisory_DBNAME = "advisory"
	}

	env.MSQL_ACCORD_DBNAME = os.Getenv("MSQL_ACCORD_DBNAME")
	if env.MSQL_ACCORD_DBNAME == "" {
		env.MSQL_ACCORD_DBNAME = "accord_dev"
	}

	env.MSQL_STOCKAL_DBNAME = os.Getenv("MSQL_STOCKAL_DBNAME")
	if env.MSQL_STOCKAL_DBNAME == "" {
		env.MSQL_STOCKAL_DBNAME = "stockal_dev"
	}

	env.MSQL_watchlist_DBNAME = os.Getenv("MSQL_watchlist_DBNAME")
	if env.MSQL_watchlist_DBNAME == "" {
		env.MSQL_watchlist_DBNAME = "watchlist"
	}

	env.MSQL_switchaccount_DBNAME = os.Getenv("MSQL_switchaccount_DBNAME")
	if env.MSQL_switchaccount_DBNAME == "" {
		env.MSQL_switchaccount_DBNAME = "switchaccount_dev"
	}

	env.MSQL_SettingService_DBNAME = os.Getenv("MSQL_SettingService_DBNAME")
	if env.MSQL_SettingService_DBNAME == "" {
		env.MSQL_SettingService_DBNAME = "settingservice"
	}

	env.MSQL_MUTUALFUND_DBNAME = os.Getenv("MSQL_MUTUALFUND_DBNAME")
	if env.MSQL_MUTUALFUND_DBNAME == "" {
		env.MSQL_MUTUALFUND_DBNAME = "mutualfund_dev"
	}
	env.FIREBASE_SERVICEKEY_PATH = os.Getenv("FIREBASE_SERVICEKEY_PATH")
	if env.FIREBASE_SERVICEKEY_PATH == "" {
		env.FIREBASE_SERVICEKEY_PATH = env.VOLUMNE_PATH + "firebase/config/gofacetrade-firebase-adminsdk-p4ds3-4bfec7d3e3.json"
	}
	// gofacetrade	--> Project settings --> Cloud Messaging --> Web Push certificates
	// public key are visible
	// private key can be obtain from MouseOver on key --> click on floatingicon --> Get public key

	///https://firebase.google.com/docs/cloud-messaging/js/client#configure_web_credentials_in_your_app
	// 	Generate a new key pair
	// Open the Cloud Messaging tab of the Firebase console Settings pane and scroll to the Web configuration section.
	// In the Web Push certificates tab, click Generate Key Pair. The console displays a notice that the key pair was generated, and displays the public key string and date added.
	env.FIREBASE_WEB_PUBLICK_EY = os.Getenv("FIREBASE_WEB_PUBLICK_EY")
	if env.FIREBASE_WEB_PUBLICK_EY == "" {
		env.FIREBASE_WEB_PUBLICK_EY = "BKF4QwKnKLHXw-gKFzeeC3Qs3KRLfnM7tYvbjHC-TbhG4klfqhzxazfaMkHcaQOwCdbC2NcB7CIEtPkYKPkAiII"
	}

	env.FIREBASE_WEB_PRIVATE_KEY = os.Getenv("FIREBASE_WEB_PRIVATE_KEY")
	if env.FIREBASE_WEB_PRIVATE_KEY == "" {
		env.FIREBASE_WEB_PRIVATE_KEY = "mBCpCfMHL5fAA5UCxl_a95PdjvGQmHHq8FmFBZDIWMM"
	}

	env.AWS_S3_REGION = os.Getenv("AWS_S3_REGION")
	if env.AWS_S3_REGION == "" {
		env.AWS_S3_REGION = "ap-south-1"
	}

	env.AWS_S3_BUCKET = os.Getenv("AWS_S3_BUCKET")
	if env.AWS_S3_BUCKET == "" {
		env.AWS_S3_BUCKET = "homesite"
	}

	env.LOG_PATH = os.Getenv("LOGS_PATH")
	if env.LOG_PATH == "" {
		env.LOG_PATH = os.TempDir()
	}

	env.AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
	if env.AWS_ACCESS_KEY_ID == "" {
		env.AWS_ACCESS_KEY_ID = "AKIARL4MKAYFZLASH6CC"
		os.Setenv("AWS_ACCESS_KEY_ID", env.AWS_ACCESS_KEY_ID)
	}

	env.AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
	if env.AWS_SECRET_ACCESS_KEY == "" {
		env.AWS_SECRET_ACCESS_KEY = "FIU22r4wBUHXEBTMAZRNVZQUMsoTsSgtDqK67gUW"
		os.Setenv("AWS_SECRET_ACCESS_KEY", env.AWS_SECRET_ACCESS_KEY)
	}

	env.MY_NODE_NAME = os.Getenv("MY_NODE_NAME")
	if env.MY_NODE_NAME == "" {
		env.MY_NODE_NAME = "node"
	}

	env.MY_POD_NAME = os.Getenv("MY_POD_NAME")
	if env.MY_POD_NAME == "" {
		env.MY_POD_NAME = "pod"
	}

	env.MY_POD_NAMESPACE = os.Getenv("MY_POD_NAMESPACE")
	if env.MY_POD_NAMESPACE == "" {
		env.MY_POD_NAMESPACE = "ns"
	}

	env.MY_POD_IP = os.Getenv("MY_POD_IP")
	if env.MY_POD_IP == "" {
		shost, _ := os.Hostname()
		env.MY_POD_IP = shost
	}

	env.MY_POD_SERVICE_ACCOUNT = os.Getenv("MY_POD_SERVICE_ACCOUNT")
	if env.MY_POD_SERVICE_ACCOUNT == "" {
		env.MY_POD_SERVICE_ACCOUNT = "localservice"
	}

	env.SERVER_TYPE = os.Getenv("SERVER_TYPE")
	if env.SERVER_TYPE == "" {
		env.SERVER_TYPE = "local"
	}

	env.REDIS_SERVICE = os.Getenv("REDIS_SERVICE")
	if env.REDIS_SERVICE == "" {
		env.REDIS_SERVICE = "localhost"
	}
	env.REDIS_PORT = os.Getenv("REDIS_PORT")
	if env.REDIS_PORT == "" {
		env.REDIS_PORT = "6379"
	}

	env.REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	if env.REDIS_PASSWORD == "" {
		env.REDIS_PASSWORD = "Mksl@1234"
	}

	env.ROUTE_SERVER = os.Getenv("ROUTE_SERVER")
	if env.ROUTE_SERVER == "" {
		env.ROUTE_SERVER = "uathsauth.hypertrade.in"
	}

	env.PATH_REPLACE = os.Getenv("PATH_REPLACE")
	if env.PATH_REPLACE == "" {
		env.PATH_REPLACE = ""
	}

	env.SMTPApiKey = os.Getenv("SMTPApiKey")
	if env.SMTPApiKey == "" {
		env.SMTPApiKey = "18669d6588975de17f0a3d61a0702ffa3c851ce5ef737e1c8f27bc1c6a233ea8"
	}

	env.SMTPUSERNAME = os.Getenv("SMTPUSERNAME")
	if env.SMTPUSERNAME == "" {
		env.SMTPUSERNAME = "mangalkeshav"
	}

	env.DGIOAUTHKEY = os.Getenv("DGIOAUTHKEY")
	if env.DGIOAUTHKEY == "" {
		env.DGIOAUTHKEY = "Basic QUkzVEIxOVJHUk1aVFhMU1VKSlBJUVVONExXR0xQSEk6VDlSTExHMVJNNEFTNFZKRkdIWVpMV01CN1BNUEI3Nk4="
	}

	env.FROM = os.Getenv("FROM")
	if env.FROM == "" {
		env.FROM = "query@mangalkeshav.com"
	}

	env.ROUTE_SERVER = os.Getenv("ROUTE_SERVER")
	if env.ROUTE_SERVER == "" {
		env.ROUTE_SERVER = "uathsint.hypertrade.in"
	}

	env.HYPERSYNC_HOLDING_URL = os.Getenv("HYPERSYNC_HOLDING_URL")
	if env.HYPERSYNC_HOLDING_URL == "" {
		env.HYPERSYNC_HOLDING_URL = "uathsint.hypertrade.in/quick/user/holdings"
	}
	env.FILE_USERDATA_PATH = os.Getenv("FILE_USERDATA_PATH")
	if env.FILE_USERDATA_PATH == "" {
		env.FILE_USERDATA_PATH = "/home/ratnadipm/go/src/src/workspace/gotechexcel/"
	}

	env.FILE_TEMPLATE_PATH = os.Getenv("FILE_TEMPLATE_PATH")
	if env.FILE_TEMPLATE_PATH == "" {
		env.FILE_TEMPLATE_PATH = "/home/ratnadipm/go/src/src/workspace/gotechexcel/"
	}
	env.FILE_POA = os.Getenv("FILE_POA")
	if env.FILE_POA == "" {
		env.FILE_POA = "/gotechexcel/POA_File/"
	}

	env.PROMETHEUS_COUNTER, _ = strconv.ParseBool(os.Getenv("PROMETHEUS_COUNTER"))
	env.PROMETHEUS_GAUGE, _ = strconv.ParseBool(os.Getenv("PROMETHEUS_GAUGE"))
	env.PROMETHEUS_HISTOGRAM, _ = strconv.ParseBool(os.Getenv("PROMETHEUS_HISTOGRAM"))
	env.PROMETHEUS_SUMMARY, _ = strconv.ParseBool(os.Getenv("PROMETHEUS_SUMMARY"))
	env.PROMETHEUS_PORT = 8090
	env.PROMETHEUS_HOOK, _ = strconv.ParseBool(os.Getenv("PROMETHEUS_HOOK"))

	env.JMETER_PATH = os.Getenv("JMETER_PATH")
	if env.JMETER_PATH == "" {
		env.JMETER_PATH = "./jmeter/bin/jmeter"
	}

	env.JMETER_GITHUB_PATH = os.Getenv("JMETER_GITHUB_PATH")
	if env.JMETER_GITHUB_PATH == "" {
		env.JMETER_GITHUB_PATH = "jmetertest"
	}

	env.REPORT_PATH = os.Getenv("REPORT_PATH")

	env.KUB_WEB_ADDRESS = os.Getenv("KUB_WEB_ADDRESS")
	if env.KUB_WEB_ADDRESS == "" {
		env.KUB_WEB_ADDRESS = "https://www.nuuu.com"
	}

	env.SECRET = os.Getenv("SECRET")
	if env.SECRET == "" {
		env.SECRET = "THIS IS USED TO SIGN AND VERIFY JWT TOKENS, REPLACE IT WITH YOUR OWN SECRET, IT CAN BE ANY STRING"
	}

	env.JWTKEY = os.Getenv("JWTKEY")
	if env.JWTKEY == "" {
		env.JWTKEY = "5daU6I2TAyE14eyskWuTxhjSttsXamJJ"
	}
	env.JWTISSUER = os.Getenv("JWTISSUER")
	if env.JWTISSUER == "" {
		env.JWTISSUER = "mirae.in"
	}

	env.JWTAUDIENCE = os.Getenv("JWTAUDIENCE")
	if env.JWTAUDIENCE == "" {
		env.JWTAUDIENCE = "http://localhost:5208"
	}

	env.JWT_GITHUB_JWT = os.Getenv("JWT_GITHUB_JWT")
	if env.JWT_GITHUB_JWT == "" {
		env.JWT_GITHUB_JWT = "YGlXq+v07DmDzg3amouf/Q=="
	}

	env.GOOGLE_REDIRECT_URL = os.Getenv("GOOGLE_REDIRECT_URL")
	if env.GOOGLE_REDIRECT_URL == "" {
		env.GOOGLE_REDIRECT_URL = "http://localhost:8087/sign/auth/"
	}

	env.GOOGLE_SECRET_PATH = os.Getenv("GOOGLE_SECRET_PATH")
	if env.GOOGLE_SECRET_PATH == "" {
		env.GOOGLE_SECRET_PATH = "./googlesecret.json"
	}

	env.GOOGLE_LOGIN_PAGE = os.Getenv("GOOGLE_LOGIN_PAGE")
	if env.GOOGLE_LOGIN_PAGE == "" {
		env.GOOGLE_LOGIN_PAGE = "login.html"
	}

	env.GOOGLE_REDIRECT_PAGE = os.Getenv("GOOGLE_REDIRECT_PAGE")
	if env.GOOGLE_REDIRECT_PAGE == "" {
		env.GOOGLE_REDIRECT_PAGE = "redirect.html"
	}

	env.GOOGLE_MOBILE_PAGE = os.Getenv("GOOGLE_MOBILE_PAGE")
	if env.GOOGLE_MOBILE_PAGE == "" {
		env.GOOGLE_MOBILE_PAGE = "mobile.html"
	}

	env.GOOGLE_CLIENTID = os.Getenv("GOOGLE_CLIENTID")
	if env.GOOGLE_CLIENTID == "" {
		env.GOOGLE_CLIENTID = "9165696624-uutli61kp3slm1g435a4mkj9sd890iv0.apps.googleusercontent.com"
	}
	env.GOOGLE_CLIENTSECRET = os.Getenv("GOOGLE_CLIENTSECRET")
	if env.GOOGLE_CLIENTSECRET == "" {
		env.GOOGLE_CLIENTSECRET = "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"
	}

	env.GOOGLE_MAIN_PAGE = os.Getenv("GOOGLE_MAIN_PAGE")
	if env.GOOGLE_MAIN_PAGE == "" {
		env.GOOGLE_MAIN_PAGE = "http://localhost:8087/trade"
	}

	env.GOOGLE_MPIN_PAGE = os.Getenv("GOOGLE_MPIN_PAGE")
	if env.GOOGLE_MPIN_PAGE == "" {
		env.GOOGLE_MPIN_PAGE = "mpin.html"
	}

	env.HS_PUBLIC_KEY = os.Getenv("HS_PUBLIC_KEY")
	if env.HS_PUBLIC_KEY == "" {
		env.HS_PUBLIC_KEY = "LS0tLS1CRUdJTiBSU0EgUFVCTElDIEtFWS0tLS0tCk1JSUJDZ0tDQVFFQXpJZ3dadG9iQ1hVdXRST0JpVEV1VmVOeEJrUUJpZkJtMTlLdE1oUTkwRE9mdDlpOWg4bXUKbGFTK2xZeXhjRDgvMGorUHRHNm9zQ2pya2dnTDc5M0JydU9OaUp0bnk1L05wSFpRVDlPdlNqS0FvQ052UGNNdQpvbEVOTGFNQ25jdUxWdmtSaEdlVTJGeDJsYm1qdGFQK2dPUlVyTUovQm9ucm1MZ1B4SkRmVXpWUHBWMXNyTXpBClByNjJ0ZEoyU2h4c0hWcmVhV01sU1RubkJxVWxYSitUVWpZU1hEOGVjcmhBd0NES3c3Z2dCRUhVdTcwdnNydFEKREFYRnVlS0dHdnpuSXVnekErRGtlaXY2UXA0NWt3SmlnVVBYMkZIaFRDU1EzZmkvd04wNC9GSEdONTZOVmJ1RApMVUJhOUZTSXB6UGJWNkVXVjhPZ1FGcnh0Yzh4a2ZPV1Z3SURBUUFCCi0tLS0tRU5EIFJTQSBQVUJMSUMgS0VZLS0tLS0K"
	}

	env.HS_AUTH_KEY = os.Getenv("HS_AUTH_KEY")
	if env.HS_AUTH_KEY == "" {
		env.HS_AUTH_KEY = "UY60MWb5bnHTQ1eD"
	}

	env.SLACKURL_ACCORDCRONJOB = os.Getenv("SLACKURL_ACCORDCRONJOB")
	if env.SLACKURL_ACCORDCRONJOB == "" {
		env.SLACKURL_ACCORDCRONJOB = "https://hooks.slack.com/services/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"
	}

	env.NUUU_MAIN_WEBURL = os.Getenv("NUUU_MAIN_WEBURL")
	if env.NUUU_MAIN_WEBURL == "" {
		env.NUUU_MAIN_WEBURL = "apnacrypto.com"
	}

	env.GOOGLE_SETMPIN_PAGE = os.Getenv("GOOGLE_SETMPIN_PAGE")
	if env.GOOGLE_SETMPIN_PAGE == "" {
		env.GOOGLE_SETMPIN_PAGE = "setmpin.html"
	}

	env.GOOGLE_REGISTER_PAGE = os.Getenv("GOOGLE_REGISTER_PAGE")
	if env.GOOGLE_REGISTER_PAGE == "" {
		env.GOOGLE_REGISTER_PAGE = "register.html"
	}

	env.GOOGLE_FORGOT_PAGE = os.Getenv("GOOGLE_FORGOT_PAGE")
	if env.GOOGLE_FORGOT_PAGE == "" {
		env.GOOGLE_FORGOT_PAGE = "forgotpassword.html"
	}

	env.GOOGLE_UNBLOCK_PAGE = os.Getenv("GOOGLE_UNBLOCK_PAGE")
	if env.GOOGLE_UNBLOCK_PAGE == "" {
		env.GOOGLE_UNBLOCK_PAGE = "unblockuser.html"
	}

	env.GOOGLE_SETPASSWORD_PAGE = os.Getenv("GOOGLE_SETPASSWORD_PAGE")
	if env.GOOGLE_SETPASSWORD_PAGE == "" {
		env.GOOGLE_SETPASSWORD_PAGE = "setpassword.html"
	}

	env.STOCKPAGE_ACCORD_SEARCH_URL = os.Getenv("STOCKPAGE_ACCORD_SEARCH_URL")
	if env.STOCKPAGE_ACCORD_SEARCH_URL == "" {
		env.STOCKPAGE_ACCORD_SEARCH_URL = "/accord/search"
	}

	env.ACCORD_SERVICE_URL = os.Getenv("ACCORD_SERVICE_URL")
	if env.ACCORD_SERVICE_URL == "" {
		env.ACCORD_SERVICE_URL = "http://accord-production-service:8085"
	}

	env.GITHUB_STOCKPAGES_URL = os.Getenv("GITHUB_STOCKPAGES_URL")
	if env.GITHUB_STOCKPAGES_URL == "" {
		env.GITHUB_STOCKPAGES_URL = "https://github.com/TecXLab/stockpages"
	}
	env.GITHUB_STOCKPAGES_DIR = os.Getenv("GITHUB_STOCKPAGES_DIR")
	if env.GITHUB_STOCKPAGES_DIR == "" {
		env.GITHUB_STOCKPAGES_DIR = "/logs-production/production/webpages/stockpages"
	}

	env.GITHUB_USERNAME = os.Getenv("GITHUB_USERNAME")
	if env.GITHUB_USERNAME == "" {
		env.GITHUB_USERNAME = "samirtecx"
	}
	env.GITHUB_PASSWORD = os.Getenv("GITHUB_PASSWORD")
	if env.GITHUB_PASSWORD == "" {
		env.GITHUB_PASSWORD = "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"
	}

	env.GITHUB_STOCKPAGES_CLEAN_CACHEFILE_URL = os.Getenv("GITHUB_STOCKPAGES_CLEAN_CACHEFILE_URL")
	if env.GITHUB_STOCKPAGES_CLEAN_CACHEFILE_URL == "" {
		env.GITHUB_STOCKPAGES_CLEAN_CACHEFILE_URL = "http://stockpagesnet-production-service:8085/clearfilecache"
	}

	env.SLACK_STOCKPAGES_URL = os.Getenv("SLACK_STOCKPAGES_URL")
	if env.SLACK_STOCKPAGES_URL == "" {
		env.SLACK_STOCKPAGES_URL = "https://hooks.slack.com/services/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"
	}

	env.SLACK_HELM_UPDATE_URL = os.Getenv("SLACK_HELM_UPDATE_URL")
	if env.SLACK_HELM_UPDATE_URL == "" {
		env.SLACK_HELM_UPDATE_URL = "https://hooks.slack.com/services/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"
	}

	env.SMS_EMAIL_SERVICE_URL = os.Getenv("SMS_EMAIL_SERVICE_URL")
	if env.SMS_EMAIL_SERVICE_URL == "" {
		env.SMS_EMAIL_SERVICE_URL = "http://smsemailservice-dev-service:8089"
	}

	env.GO_JWT_SERVICE_URL = os.Getenv("GO_JWT_SERVICE_URL")
	if env.GO_JWT_SERVICE_URL == "" {
		env.GO_JWT_SERVICE_URL = "http://gojwtservice-dev-service:8085"
	}

	env.ACCORD_KEY = os.Getenv("ACCORD_KEY")
	if env.ACCORD_KEY == "" {
		env.ACCORD_KEY = "J1yHW6kYHOBy0IuzePKWnGFWTv1J1ztD"
	}

	env.SlackURL = os.Getenv("SlackURL")
	if env.SlackURL == "" {
		env.SlackURL = "https://hooks.slack.com/services/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"
	}
	env.INDEX_MASTER_URL = os.Getenv("INDEX_MASTER_URL")
	if env.INDEX_MASTER_URL == "" {
		env.INDEX_MASTER_URL = "https://nuuu.com/contractmaster/getindexmaster"
	}
	env.BASEPATH = os.Getenv("BASEPATH")
	if env.BASEPATH == "" {
		env.BASEPATH = "chartdataservice"
	}
	env.SETTLEMENT_FILE_PATH = os.Getenv("SETTLEMENT_FILE_PATH")
	if env.SETTLEMENT_FILE_PATH == "" {
		env.SETTLEMENT_FILE_PATH = "settlementfile"
	}
	env.EQPATH = os.Getenv("EQPATH")
	if env.EQPATH == "" {
		env.EQPATH = "nse1"
	}
	env.FOPATH = os.Getenv("FOPATH")
	if env.FOPATH == "" {
		env.FOPATH = "nsefo1"
	}
	env.CDPATH = os.Getenv("CDPATH")
	if env.CDPATH == "" {
		env.CDPATH = "nsecd1"
	}
	env.RAWPATH = os.Getenv("RAWPATH")
	if env.RAWPATH == "" {
		env.RAWPATH = "rawdata"
	}
	env.NSETimeFormat = os.Getenv("NSETimeFormat")
	if env.NSETimeFormat == "" {
		env.NSETimeFormat = "15:04"
	}
	env.NSEEQ_Start_Time = os.Getenv("NSEEQ_Start_Time")
	if env.NSEEQ_Start_Time == "" {
		env.NSEEQ_Start_Time = "09:15"
	}
	env.NSEEQ_End_Time = os.Getenv("NSEEQ_End_Time")
	if env.NSEEQ_End_Time == "" {
		env.NSEEQ_End_Time = "15:30"
	}
	env.NSEFO_Start_Time = os.Getenv("NSEFO_Start_Time")
	if env.NSEFO_Start_Time == "" {
		env.NSEFO_Start_Time = "09:15"
	}
	env.NSEFO_End_Time = os.Getenv("NSEFO_End_Time")
	if env.NSEFO_End_Time == "" {
		env.NSEFO_End_Time = "15:30"
	}
	env.NSECD_Start_Time = os.Getenv("NSECD_Start_Time")
	if env.NSECD_Start_Time == "" {
		env.NSECD_Start_Time = "09:00"
	}
	env.NSECD_End_Time = os.Getenv("NSECD_End_Time")
	if env.NSECD_End_Time == "" {
		env.NSECD_End_Time = "17:00"
	}
	env.NSEEQ_Min_Time = os.Getenv("NSEEQ_Min_Time")
	if env.NSEEQ_Min_Time == "" {
		env.NSEEQ_Min_Time = "915"
	}
	env.NSEEQ_MAX_Time = os.Getenv("NSEEQ_MAX_Time")
	if env.NSEEQ_MAX_Time == "" {
		env.NSEEQ_MAX_Time = "1540"
	}
	env.NSEFO_Min_Time = os.Getenv("NSEFO_Min_Time")
	if env.NSEFO_Min_Time == "" {
		env.NSEFO_Min_Time = "915"
	}
	env.NSEFO_MAX_Time = os.Getenv("NSEFO_MAX_Time")
	if env.NSEFO_MAX_Time == "" {
		env.NSEFO_MAX_Time = "1540"
	}
	env.NSECD_Min_Time = os.Getenv("NSECD_Min_Time")
	if env.NSECD_Min_Time == "" {
		env.NSECD_Min_Time = "900"
	}
	env.NSECD_MAX_Time = os.Getenv("NSECD_MAX_Time")
	if env.NSECD_MAX_Time == "" {
		env.NSECD_MAX_Time = "1700"
	}
	env.TYPESENSE_KEY = os.Getenv("TYPESENSE_KEY")
	if env.TYPESENSE_KEY == "" {
		env.TYPESENSE_KEY = "X-TYPESENSE-API-KEY"
	}

	env.TYPESENSE_VALUE = os.Getenv("TYPESENSE_VALUE")
	if env.TYPESENSE_VALUE == "" {
		env.TYPESENSE_VALUE = "Hu52dwsas2AdxdE"
	}

	env.TYPESENSE_BASE_URL = os.Getenv("TYPESENSE_BASE_URL")
	if env.TYPESENSE_BASE_URL == "" {
		env.TYPESENSE_BASE_URL = "https://www.nuuu.com"
	}

	env.CONTRACTMASTER_TYPESENSE_COLLECTION = os.Getenv("CONTRACTMASTER_TYPESENSE_COLLECTION")
	if env.CONTRACTMASTER_TYPESENSE_COLLECTION == "" {
		env.CONTRACTMASTER_TYPESENSE_COLLECTION = "stocksearch2"
	}

	env.CONTRACTMASTER_TYPESENSE_QUERY_BY = os.Getenv("CONTRACTMASTER_TYPESENSE_QUERY_BY")
	if env.CONTRACTMASTER_TYPESENSE_QUERY_BY == "" {
		env.CONTRACTMASTER_TYPESENSE_QUERY_BY = "cnam,exmnt,strikprc,optyp,seris,expry"
	}

	env.CONTRACTMASTER_TYPESENSE_GROUP_BY = os.Getenv("CONTRACTMASTER_TYPESENSE_GROUP_BY")
	if env.CONTRACTMASTER_TYPESENSE_GROUP_BY == "" {
		env.CONTRACTMASTER_TYPESENSE_GROUP_BY = "seris"
	}

	env.TYPESENSE_GROUP_LIMIT = os.Getenv("TYPESENSE_GROUP_LIMIT")
	if env.TYPESENSE_GROUP_LIMIT == "" {
		env.TYPESENSE_GROUP_LIMIT = "99"
	}

	env.TYPESENSE_EXHAUSTIVE_SEARCH = os.Getenv("TYPESENSE_EXHAUSTIVE_SEARCH")
	if env.TYPESENSE_EXHAUSTIVE_SEARCH == "" {
		env.TYPESENSE_EXHAUSTIVE_SEARCH = "true"
	}

	env.CONTRACTMASTER_TYPESENSE_PERPAGE = os.Getenv("CONTRACTMASTER_TYPESENSE_PERPAGE")
	if env.CONTRACTMASTER_TYPESENSE_PERPAGE == "" {
		env.CONTRACTMASTER_TYPESENSE_PERPAGE = "250"
	}

	env.CONTRACTMASTER_TYPESENSE_SORT_BY = os.Getenv("CONTRACTMASTER_TYPESENSE_SORT_BY")
	if env.CONTRACTMASTER_TYPESENSE_SORT_BY == "" {
		env.CONTRACTMASTER_TYPESENSE_SORT_BY = "priorityno:asc,nexpry:asc"
	}

	env.MUTUALFUND_TYPESENSE_COLLECTION = os.Getenv("MUTUALFUND_TYPESENSE_COLLECTION")
	if env.MUTUALFUND_TYPESENSE_COLLECTION == "" {
		env.MUTUALFUND_TYPESENSE_COLLECTION = "greenwaremf2"
	}

	env.MUTUALFUND_TYPESENSE_QUERY_BY = os.Getenv("MUTUALFUND_TYPESENSE_QUERY_BY")
	if env.MUTUALFUND_TYPESENSE_QUERY_BY == "" {
		env.MUTUALFUND_TYPESENSE_QUERY_BY = "SchemeName,AMCName"
	}

	env.MUTUALFUND_TYPESENSE_FACET_QUERY_BY = os.Getenv("MUTUALFUND_TYPESENSE_FACET_QUERY_BY")
	if env.MUTUALFUND_TYPESENSE_FACET_QUERY_BY == "" {
		env.MUTUALFUND_TYPESENSE_FACET_QUERY_BY = "AMCName"
	}

	env.MUTUALFUND_TYPESENSE_QUERY_BY_SCHCODE = os.Getenv("MUTUALFUND_TYPESENSE_QUERY_BY_SCHCODE")
	if env.MUTUALFUND_TYPESENSE_QUERY_BY_SCHCODE == "" {
		env.MUTUALFUND_TYPESENSE_QUERY_BY_SCHCODE = "SchCode"
	}

	env.MUTUALFUND_TYPESENSE_SORT_BY = os.Getenv("MUTUALFUND_TYPESENSE_SORT_BY")
	if env.MUTUALFUND_TYPESENSE_SORT_BY == "" {
		env.MUTUALFUND_TYPESENSE_SORT_BY = "Return3Year:desc"
	}

	env.ACCORD_TYPESENSE_COLLECTION = os.Getenv("ACCORD_TYPESENSE_COLLECTION")
	if env.ACCORD_TYPESENSE_COLLECTION == "" {
		env.ACCORD_TYPESENSE_COLLECTION = "companymasters"
	}

	env.ACCORD_TYPESENSE_QUERY_BY = os.Getenv("ACCORD_TYPESENSE_QUERY_BY")
	if env.ACCORD_TYPESENSE_QUERY_BY == "" {
		env.ACCORD_TYPESENSE_QUERY_BY = "symbl,cnm"
	}

	env.ACCORD_TYPESENSE_SORT_BY = os.Getenv("ACCORD_TYPESENSE_SORT_BY")
	if env.ACCORD_TYPESENSE_SORT_BY == "" {
		env.ACCORD_TYPESENSE_SORT_BY = "exid:asc"
	}

	env.ACCORD_TYPESENSE_PERPAGE = os.Getenv("ACCORD_TYPESENSE_PERPAGE")
	if env.ACCORD_TYPESENSE_PERPAGE == "" {
		env.ACCORD_TYPESENSE_PERPAGE = "250"
	}

	env.NAMESPACE_WEBHOOK = os.Getenv("NAMESPACE_WEBHOOK")
	if env.NAMESPACE_WEBHOOK == "" {
		env.NAMESPACE_WEBHOOK = "production"
	}

	env.IS_RAW_FILE_WRITING_ENABLE = os.Getenv("IS_RAW_FILE_WRITING_ENABLE")
	if env.IS_RAW_FILE_WRITING_ENABLE == "" {
		env.IS_RAW_FILE_WRITING_ENABLE = "0"
	}

	env.DPId = os.Getenv("DPId")
	if env.DPId == "" {
		env.DPId = "96400"
	}

	env.BOID = os.Getenv("BOID")
	if env.BOID == "" {
		env.BOID = "1209640000000101"
	}

	env.CMID = os.Getenv("CMID")
	if env.CMID == "" {
		env.CMID = "M70032"
	}

	env.TMID = os.Getenv("TMID")
	if env.TMID == "" {
		env.TMID = ""
	}

	env.NumOfDays = os.Getenv("NumOfDays")
	if env.NumOfDays == "" {
		env.NumOfDays = "30"
	}

	env.ExID = os.Getenv("ExID")
	if env.ExID == "" {
		env.ExID = "12"
	}

	env.ReturnURL = os.Getenv("ReturnURL")
	if env.ReturnURL == "" {
		env.ReturnURL = "https://foocut.com/edis/edisresponse"
	}

	env.ReqIdentifier = os.Getenv("ReqIdentifier")
	if env.ReqIdentifier == "" {
		env.ReqIdentifier = "S"
	}

	env.ReqType = os.Getenv("ReqType")
	if env.ReqType == "" {
		env.ReqType = "D"
	}

	env.Version = os.Getenv("Version")
	if env.Version == "" {
		env.Version = ""
	}

	env.TB_PARTNERID_KEY = os.Getenv("TB_PARTNERID_KEY")
	if env.TB_PARTNERID_KEY == "" {
		env.TB_PARTNERID_KEY = "632d500e1ee67adb72810a04"
	}
	env.TB_SECRET_KEY = os.Getenv("TB_SECRET_KEY")
	if env.TB_SECRET_KEY == "" {
		env.TB_SECRET_KEY = "KLMNGLKSHDEF"
	}
	env.TB_SECRETIV_KEY = os.Getenv("TB_SECRETIV_KEY")
	if env.TB_SECRETIV_KEY == "" {
		env.TB_SECRETIV_KEY = "ABLAODEWFHXC"
	}

	env.CASHFREE_CLIENT_ID = os.Getenv("CASHFREE_CLIENT_ID")
	if env.CASHFREE_CLIENT_ID == "" {
		env.CASHFREE_CLIENT_ID = "TEST2472452660208b1a127b260307542742"
	}

	env.CASHFREE_CLIENT_SECRET = os.Getenv("CASHFREE_CLIENT_SECRET")
	if env.CASHFREE_CLIENT_SECRET == "" {
		env.CASHFREE_CLIENT_SECRET = "TEST4f26a25f639fb235e9e36f666238eb95facc8926"
	}

	//
	env.WA_BASE_URL = os.Getenv("WA_BASE_URL")
	if env.WA_BASE_URL == "" {
		env.WA_BASE_URL = "https://www.nuuu.com"
	}
	env.WA_EKYC_MESSAGE_URL = os.Getenv("WA_EKYC_MESSAGE_URL")
	if env.WA_EKYC_MESSAGE_URL == "" {
		env.WA_EKYC_MESSAGE_URL = "/kycapi/WhatsappTextMessage"
	}
	env.WA_GET_TOKEN_URL = os.Getenv("WA_GET_TOKEN_URL")
	if env.WA_GET_TOKEN_URL == "" {
		env.WA_GET_TOKEN_URL = "/kycapi/GetToken"
	}
	env.WA_EKYCID = os.Getenv("WA_EKYCID")
	if env.WA_EKYCID == "" {
		env.WA_EKYCID = "104356255786722"
	}
	env.WA_WHATSAPP_URL = os.Getenv("WA_WHATSAPP_URL")
	if env.WA_WHATSAPP_URL == "" {
		env.WA_WHATSAPP_URL = "https://graph.facebook.com/v14.0/102822289275903/messages"
	}

	env.WA_WHATSAPP_TOKEN = os.Getenv("WA_WHATSAPP_TOKEN")
	if env.WA_WHATSAPP_TOKEN == "" {
		env.WA_WHATSAPP_TOKEN = ""
	}

	env.LEX_BotID = os.Getenv("LEX_BotID")
	if env.LEX_BotID == "" {
		env.LEX_BotID = "GCBLMNS6UL"
	}

	env.LEX_BotAlias = os.Getenv("LEX_BotAlias")
	if env.LEX_BotAlias == "" {
		env.LEX_BotAlias = "BankbotAlias"
	}

	env.LEX_BotAliasID = os.Getenv("LEX_BotAliasID")
	if env.LEX_BotAliasID == "" {
		env.LEX_BotAliasID = "TSTALIASID"
	}

	env.LEX_LocalID = os.Getenv("LEX_LocalID")
	if env.LEX_LocalID == "" {
		env.LEX_LocalID = "en_US"
	}
	env.GOOGLE_TOTP_PAGE = os.Getenv("GOOGLE_TOTP_PAGE")
	if env.GOOGLE_TOTP_PAGE == "" {
		env.GOOGLE_TOTP_PAGE = "totptoken.html"
	}

	env.IPO_MEMBER = os.Getenv("IPO_MEMBER")
	if env.IPO_MEMBER == "" {
		env.IPO_MEMBER = "90251"
	}
	env.IPO_USER_ID = os.Getenv("IPO_USER_ID")
	if env.IPO_USER_ID == "" {
		env.IPO_USER_ID = "ADMIN90251"
	}
	env.IPO_MEMBER_PASSWORD = os.Getenv("IPO_MEMBER_PASSWORD")
	if env.IPO_MEMBER_PASSWORD == "" {
		env.IPO_MEMBER_PASSWORD = "Mksl@123456789"
	}
	env.BROKERID = os.Getenv("BROKERID")
	if env.BROKERID == "" {
		env.BROKERID = "NIT"
	}
	env.SLACKURL_EDIS = os.Getenv("SLACKURL_EDIS")
	if env.SLACKURL_EDIS == "" {
		env.SLACKURL_EDIS = "https://hooks.slack.com/services/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"
	}

	env.EDISDONE_PAGE = os.Getenv("EDISDONE_PAGE")
	if env.EDISDONE_PAGE == "" {
		env.EDISDONE_PAGE = "edisdone.html"
	}
	env.BLOCK_USER_PAGE = os.Getenv("BLOCK_USER_PAGE")
	if env.BLOCK_USER_PAGE == "" {
		env.BLOCK_USER_PAGE = "login-error.html"
	}
	env.GOOGLE_FORGOTMPIN_PAGE = os.Getenv("GOOGLE_FORGOTMPIN_PAGE")
	if env.GOOGLE_FORGOTMPIN_PAGE == "" {
		env.GOOGLE_FORGOTMPIN_PAGE = "forgotmpin.html"
	}
	env.GW_URL = os.Getenv("GW_URL")
	if env.GW_URL == "" {
		env.GW_URL = "http://192.168.5.170:8087/api/sso/v1/for-acc-tkn"
	}
	env.GW_KEY = os.Getenv("GW_KEY")
	if env.GW_KEY == "" {
		env.GW_KEY = "uim#zYtj$lM4C@rD3h"
	}
	env.NUUU_JWT_URL = os.Getenv("NUUU_JWT_URL")
	if env.NUUU_JWT_URL == "" {
		env.NUUU_JWT_URL = "https://www.nuuu.com/sign/getjwttoken?mobileno="
	}
	env.LEX_LAMBDA_PROD_SERVER = os.Getenv("LEX_LAMBDA_PROD_SERVER")
	if env.LEX_LAMBDA_PROD_SERVER == "" {
		env.LEX_LAMBDA_PROD_SERVER = "https://www.nuuu.com"
	}

	env.IPO_API_URL = os.Getenv("IPO_API_URL")
	if env.IPO_API_URL == "" {
		env.IPO_API_URL = "https://uat-ipo.nseindia.com/eipo"
	}

	env.STOCKAL_APIKEY = os.Getenv("STOCKAL_APIKEY")
	if env.STOCKAL_APIKEY == "" {
		env.STOCKAL_APIKEY = "5O6fJAFZ3Vh7uD0qZ005"
	}
	env.STOCKAL_ACCESSKEY = os.Getenv("STOCKAL_ACCESSKEY")
	if env.STOCKAL_ACCESSKEY == "" {
		env.STOCKAL_ACCESSKEY = "dJUbYBeWue5u9DtVyoaUtHGwy0m5uH"
	}
	env.PLAN_SUBSCRIPTION_ID = os.Getenv("PLAN_SUBSCRIPTION_ID")
	if env.PLAN_SUBSCRIPTION_ID == "" {
		env.PLAN_SUBSCRIPTION_ID = "9JRF6HD7N4NDRD0S"
	}
	env.STOCKAL_URL = os.Getenv("STOCKAL_URL")
	if env.STOCKAL_URL == "" {
		env.STOCKAL_URL = "https://uatclientapi-v2.stockal.com/client/v2"
	}
	env.STOCKALFEED_URL = os.Getenv("STOCKALFEED_URL")
	if env.STOCKALFEED_URL == "" {
		env.STOCKALFEED_URL = "https://uat-prices.stockal.com/price/current-price"
	}
	env.STOCKALOAUTHTOKEN_URL = os.Getenv("STOCKALOAUTHTOKEN_URL")
	if env.STOCKALOAUTHTOKEN_URL == "" {
		env.STOCKALOAUTHTOKEN_URL = "https://uat-auth.stockal.com/oauth/token"
	}

	env.MSQL_Whatsapp_DBNAME = os.Getenv("MSQL_Whatsapp_DBNAME")
	if env.MSQL_Whatsapp_DBNAME == "" {
		env.MSQL_Whatsapp_DBNAME = "whatsappdb"
	}
	env.EDISFAILURE_PAGE = os.Getenv("EDISFAILURE_PAGE")
	if env.EDISFAILURE_PAGE == "" {
		env.EDISFAILURE_PAGE = "edisFailure.html"
	}
	env.SIDECAR_IP = os.Getenv("SIDECAR_IP")
	if env.SIDECAR_IP == "" {
		env.SIDECAR_IP = "127.0.0.1"
	}
	env.EXPOSEPORT_SIDECAR = os.Getenv("EXPOSEPORT_SIDECAR")
	if env.EXPOSEPORT_SIDECAR == "" {
		env.EXPOSEPORT_SIDECAR = "8086"
	}
	env.EXPOSEPORT_SENDER_PORT = os.Getenv("EXPOSEPORT_SENDER_PORT")
	if env.EXPOSEPORT_SENDER_PORT == "" {
		env.EXPOSEPORT_SENDER_PORT = "8085"
	}
	env.TOPIC_PUB = os.Getenv("TOPIC_PUB")
	if env.TOPIC_PUB == "" {
		env.TOPIC_PUB = "1"
	}
	env.DI_DB_WHATSAPP = os.Getenv("DI_DB_WHATSAPP")
	if env.DI_DB_WHATSAPP == "" {
		env.DI_DB_WHATSAPP = "whatsappdb"
	}

	env.WA_FINVUID = os.Getenv("WA_FINVUID")
	if env.WA_FINVUID == "" {
		env.WA_FINVUID = "104356255786722"
	}

	env.IPO_SLACK_URL = os.Getenv("IPO_SLACK_URL")
	if env.IPO_SLACK_URL == "" {
		env.IPO_SLACK_URL = "https://hooks.slack.com/services/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"
	}

	env.STOCKAL_PUBLIC_PEM_PATH = os.Getenv("STOCKAL_PUBLIC_PEM_PATH")
	if env.STOCKAL_PUBLIC_PEM_PATH == "" {
		env.STOCKAL_PUBLIC_PEM_PATH = "prodnuuuprivatekey.pem"
	}

	env.STOCKAL_PARTNER_SSO_URL = os.Getenv("STOCKAL_PARTNER_SSO_URL")
	if env.STOCKAL_PARTNER_SSO_URL == "" {
		env.STOCKAL_PARTNER_SSO_URL = "https://uatclientapi-v2.stockal.com/v2/partnersso/login"
	}

	env.STOCKAL_JWT_AUD = os.Getenv("STOCKAL_JWT_AUD")
	if env.STOCKAL_JWT_AUD == "" {
		env.STOCKAL_JWT_AUD = "stockal.com"
	}

	env.STOCKAL_JWT_IIS = os.Getenv("STOCKAL_JWT_IIS")
	if env.STOCKAL_JWT_IIS == "" {
		env.STOCKAL_JWT_IIS = "NUUU"
	}
	env.HYPERSYNCLOGIN_URL = os.Getenv("HYPERSYNCLOGIN_URL")
	if env.HYPERSYNCLOGIN_URL == "" {
		env.HYPERSYNCLOGIN_URL = "http://192.168.5.170:8086/system/login"
	}
	env.STOCKAL_REDIRECT_URL = os.Getenv("STOCKAL_REDIRECT_URL")
	if env.STOCKAL_REDIRECT_URL == "" {
		env.STOCKAL_REDIRECT_URL = "https://nuuu.stockal.com/sso-signin?code={authcode}"
	}
	env.HS_PUBLIC_KEY_UAT = os.Getenv("HS_PUBLIC_KEY_UAT")
	if env.HS_PUBLIC_KEY_UAT == "" {
		env.HS_PUBLIC_KEY_UAT = "LS0tLS1CRUdJTiBSU0EgUFVCTElDIEtFWS0tLS0tCk1JSUJDZ0tDQVFFQXJzQWtKZWpxalFYOStvV0FWUmQ5L2VpNllpVmdwZjRlT040OElvOG1uN3J0VW9OK0FHaE4KRUVYMVdNS2luM1VoditzdTRXU2J0QWVvaVRKcmVhTWpMZkNkSkNtb2krNXdrbDhpUDJsUkszeGhCdXc1MDhncgpvWFkwbEU0QkZWVzVxd3cwSXVFSitRNkZYV1pPSHRMcmExMnRUMzYwTmgrTU5MSVNwcG5WU3hGbUhIQjVXd2dFCmdYSzl4VEVhTUtnZ3BOQTZNK3UwMVptZmN4enlZbWpMTWlLb29xQmxLdXJRd0pkdGt2cUJDaGU1N1czd2IyY0MKamRJV1UrL2tUbFRIT0ZiZUVpZDhkT1B6Uk94ckdBRTRINE1KTUdIRGFRZDF2TUNGaW9idzRZcW9yckdCb0VteQpVWnFrL2hIMWpJa2RlYUNPRlRjMlRoZ010SCszMytrejVRSURBUUFCCi0tLS0tRU5EIFJTQSBQVUJMSUMgS0VZLS0tLS0K"
	}

	env.HS_AUTH_KEY_UAT = os.Getenv("HS_AUTH_KEY_UAT")
	if env.HS_AUTH_KEY_UAT == "" {
		env.HS_AUTH_KEY_UAT = "Ki9gY2E39q6OoRDV"
	}
	env.TB_MODULE_NAME = os.Getenv("TB_MODULE_NAME")
	if env.TB_MODULE_NAME == "" {
		env.TB_MODULE_NAME = "TPA/WB/PD"
	}

	env.TEAM_CICD = os.Getenv("TEAM_CICD")
	if env.TEAM_CICD == "" {
		env.TEAM_CICD = "https://miraeassetcm.webhook.office.com/webhookb2/181d9dfc-d3a8-4199-b8a1-c21cc722a27d@8525e18d-2cce-43ce-8106-64d085abd87a/JenkinsCI/c907078cddd1435e9378aff5972d066e/92d05e0e-2dca-49be-94e5-39017dc880f2"
	}

	env.TEAM_EKS = os.Getenv("TEAM_EKS")
	if env.TEAM_EKS == "" {
		env.TEAM_EKS = "https://miraeassetcm.webhook.office.com/webhookb2/181d9dfc-d3a8-4199-b8a1-c21cc722a27d@8525e18d-2cce-43ce-8106-64d085abd87a/JenkinsCI/c907078cddd1435e9378aff5972d066e/92d05e0e-2dca-49be-94e5-39017dc880f2"
	}

	env.CICD_USER = os.Getenv("CICD_USER")
	if env.CICD_USER == "" {
		env.CICD_USER = "d68bT]HsT[z3I1d"
	}

	env.CICD_ADMIN = os.Getenv("CICD_ADMIN")
	if env.CICD_ADMIN == "" {
		env.CICD_ADMIN = "euM3cZkv1}[MerM"
	}

	env.VERSION = os.Getenv("VERSION")
	if env.VERSION == "" {
		env.VERSION = "V1"
	}

	env.NETCORE_SMS_HOST = os.Getenv("NETCORE_SMS_HOST")
	if env.NETCORE_SMS_HOST == "" {
		env.NETCORE_SMS_HOST = "http://bulkpush.mytoday.com/BulkSms/JsonSingleApi"
	}

	env.NETCORE_SMS_USERNAME = os.Getenv("NETCORE_SMS_USERNAME")
	if env.NETCORE_SMS_USERNAME == "" {
		env.NETCORE_SMS_USERNAME = "9619452478"
	}

	env.NETCORE_SMS_PASSWORD = os.Getenv("NETCORE_SMS_PASSWORD")
	if env.NETCORE_SMS_PASSWORD == "" {
		env.NETCORE_SMS_PASSWORD = "Macm@1234"
	}

	env.NETCORE_EMAIL_API_KEY = os.Getenv("NETCORE_EMAIL_API_KEY")
	if env.NETCORE_EMAIL_API_KEY == "" {
		env.NETCORE_EMAIL_API_KEY = "efde364d7c0a1d805563bbb43b88c357"
	}

	env.NETCORE_EMAIL_HOST = os.Getenv("NETCORE_EMAIL_HOST")
	if env.NETCORE_EMAIL_HOST == "" {
		env.NETCORE_EMAIL_HOST = "https://emailapi.netcorecloud.net/v5.1/mail/send"
	}

	env.NETCORE_EMAIL_CONTENT_TYPE = os.Getenv("NETCORE_EMAIL_CONTENT_TYPE")
	if env.NETCORE_EMAIL_CONTENT_TYPE == "" {
		env.NETCORE_EMAIL_CONTENT_TYPE = "html"
	}

	env.NETCORE_SMS_FEEDID = 383551

	env.LOGIN_OTP_DIGITS = 6
	env.FORGOTPWD_OTP_DIGITS = 3
	env.OTPDIGIT = 6
	env.OTP_TIMEOUT = 180
	//

	env.RUPEESEED_IV = os.Getenv("RUPEESEED_IV")
	if env.RUPEESEED_IV == "" {
		env.RUPEESEED_IV = "320ef7705d1030f0a1a55b3dcf676cb8"
	}

	env.RUPEESEED_SALT = os.Getenv("RUPEESEED_SALT")
	if env.RUPEESEED_SALT == "" {
		env.RUPEESEED_SALT = "498960e491150a0fc0f21822a147fd62"
	}

	env.RUPEESEED_PASSTEXT = os.Getenv("RUPEESEED_PASSTEXT")
	if env.RUPEESEED_PASSTEXT == "" {
		env.RUPEESEED_PASSTEXT = "RUPEEWEBSITE"
	}

	env.RUPEESEED_ENDPOINT = os.Getenv("RUPEESEED_ENDPOINT")
	if env.RUPEESEED_ENDPOINT == "" {
		env.RUPEESEED_ENDPOINT = "https://trade.mstock.com"
	}

	env.SetMap()
	container.Singleton(func() *Env {
		return env
	})
	return nil
}

func (env *Env) SetMap() {
	val := reflect.ValueOf(env).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		f := valueField.Interface()
		val := reflect.ValueOf(f)
		ENV_MAP[typeField.Name] = val.String()
	}
}

func (env *Env) GetValue(key string) string {
	value, exists := ENV_MAP[key]
	if exists {
		return value
	}
	return ""
}
