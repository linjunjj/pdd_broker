## pdd



web协议以及前端

判断登陆请求包

POST /earth/api/merchant/checkLoginType HTTP/1.1
Referer: Android
User-Agent: androidandroid Mozilla/5.0 (Linux; Android 4.4.2; SM-G955F Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36  pddmt_android_version/2.2.3 pddmt_android_build/2025 pddmt_android_channel/yybandroid Mozilla/5.0 (Linux; Android 4.4.2; SM-G955F Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36  pddmt_android_version/2.2.3 pddmt_android_build/2025 pddmt_android_channel/yyb
ETag: TDh2Hsue
PASSID: 
Content-Type: application/json;charset=UTF-8
Content-Length: 29
Host: mms.pinduoduo.com
Connection: close
Cookie: api_uid=rBQH01u/V4xeZiVSDExjAg==; PASS_ID=bapp_1-9FlAvk91RzX8d8U/xkgvx6eXfKXbcHiToeiSyQQsz5MAfM1Z49/3lTo5sq1llcrz9ULGUuzfzNn9Pe/g3ZjiCg_188639248_17085632

{"username":"pdd18863924870"}

店铺查询

GET /earth/api/merchant/queryMerchantInfo HTTP/1.1
Referer: Android
User-Agent: androidandroid Mozilla/5.0 (Linux; Android 4.4.2; SM-G955F Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36  pddmt_android_version/2.2.3 pddmt_android_build/2025 pddmt_android_channel/yybandroid Mozilla/5.0 (Linux; Android 4.4.2; SM-G955F Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36  pddmt_android_version/2.2.3 pddmt_android_build/2025 pddmt_android_channel/yyb
ETag: TDh2Hsue
Content-Type: application/json;charset=UTF-8
PASSID: bapp_1-qfuJ+XmbTsdk4onr3Nlf/vnfRj7VDZ9AdJVTpPQmaKzxq1NxX8DiwLGNqwnl6AO0ThltKGgfpcWRXNWobxcyxg_188639248_17085632
Host: mms.pinduoduo.com
Connection: close
Cookie: api_uid=rBQH01u/V4xeZiVSDExjAg==; PASS_ID=bapp_1-qfuJ+XmbTsdk4onr3Nlf/vnfRj7VDZ9AdJVTpPQmaKzxq1NxX8DiwLGNqwnl6AO0ThltKGgfpcWRXNWobxcyxg_188639248_17085632

用户信息查询

POST /earth/api/user/userinfo HTTP/1.1
Referer: Android
User-Agent: androidandroid Mozilla/5.0 (Linux; Android 4.4.2; SM-G955F Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36  pddmt_android_version/2.2.3 pddmt_android_build/2025 pddmt_android_channel/yybandroid Mozilla/5.0 (Linux; Android 4.4.2; SM-G955F Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36  pddmt_android_version/2.2.3 pddmt_android_build/2025 pddmt_android_channel/yyb
ETag: TDh2Hsue
PASSID: bapp_1-qfuJ+XmbTsdk4onr3Nlf/vnfRj7VDZ9AdJVTpPQmaKzxq1NxX8DiwLGNqwnl6AO0ThltKGgfpcWRXNWobxcyxg_188639248_17085632
Content-Type: application/json;charset=UTF-8
Content-Length: 2
Host: mms.pinduoduo.com
Connection: close
Cookie: api_uid=rBQH01u/V4xeZiVSDExjAg==; PASS_ID=bapp_1-qfuJ+XmbTsdk4onr3Nlf/vnfRj7VDZ9AdJVTpPQmaKzxq1NxX8DiwLGNqwnl6AO0ThltKGgfpcWRXNWobxcyxg_188639248_17085632

{}


货物信息提醒

POST /sydney/api/dailyMallGoods/redDot HTTP/1.1
Referer: Android
User-Agent: androidandroid Mozilla/5.0 (Linux; Android 4.4.2; SM-G955F Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36  pddmt_android_version/2.2.3 pddmt_android_build/2025 pddmt_android_channel/yybandroid Mozilla/5.0 (Linux; Android 4.4.2; SM-G955F Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36  pddmt_android_version/2.2.3 pddmt_android_build/2025 pddmt_android_channel/yyb
ETag: TDh2Hsue
PASSID: bapp_1-qfuJ+XmbTsdk4onr3Nlf/vnfRj7VDZ9AdJVTpPQmaKzxq1NxX8DiwLGNqwnl6AO0ThltKGgfpcWRXNWobxcyxg_188639248_17085632
Content-Type: application/json;charset=UTF-8
Content-Length: 2
Host: mms.pinduoduo.com
Connection: close
Cookie: api_uid=rBQH01u/V4xeZiVSDExjAg==; PASS_ID=bapp_1-qfuJ+XmbTsdk4onr3Nlf/vnfRj7VDZ9AdJVTpPQmaKzxq1NxX8DiwLGNqwnl6AO0ThltKGgfpcWRXNWobxcyxg_188639248_17085632

{}

获取websocket聊天接口
POST /chats/getToken HTTP/1.1
User-Agent: android Mozilla/5.0 (Linux; Android 4.4.2; SM-G955F Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36  pddmt_android_version/2.2.3 pddmt_android_build/2025 pddmt_android_channel/yyb
Content-Type: application/json; charset=utf-8
Content-Length: 20
Host: mms.pinduoduo.com
Connection: close
Cookie: api_uid=rBQH01u/V4xeZiVSDExjAg==; PASS_ID=bapp_1-qfuJ+XmbTsdk4onr3Nlf/vnfRj7VDZ9AdJVTpPQmaKzxq1NxX8DiwLGNqwnl6AO0ThltKGgfpcWRXNWobxcyxg_188639248_17085632

{"client":"android"}


实时获取用户聊天信息

GET /chats/userinfo/realtime HTTP/1.1
User-Agent: android Mozilla/5.0 (Linux; Android 4.4.2; SM-G955F Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36  pddmt_android_version/2.2.3 pddmt_android_build/2025 pddmt_android_channel/yyb
Host: mms.pinduoduo.com
Connection: close
Cookie: api_uid=rBQH01u/V4xeZiVSDExjAg==; PASS_ID=bapp_1-qfuJ+XmbTsdk4onr3Nlf/vnfRj7VDZ9AdJVTpPQmaKzxq1NxX8DiwLGNqwnl6AO0ThltKGgfpcWRXNWobxcyxg_188639248_17085632


类目订单信息

GET /chats/latest/groupSystemMessages HTTP/1.1
User-Agent: android Mozilla/5.0 (Linux; Android 4.4.2; SM-G955F Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36  pddmt_android_version/2.2.3 pddmt_android_build/2025 pddmt_android_channel/yyb
Host: mms.pinduoduo.com
Connection: close
Cookie: api_uid=rBQH01u/V4xeZiVSDExjAg==; PASS_ID=bapp_1-qfuJ+XmbTsdk4onr3Nlf/vnfRj7VDZ9AdJVTpPQmaKzxq1NxX8DiwLGNqwnl6AO0ThltKGgfpcWRXNWobxcyxg_188639248_17085632


查询显示商品列表

POST /vodka/v2/mms/official/query/display/mall/goodsList HTTP/1.1
Referer: Android
User-Agent: androidandroid Mozilla/5.0 (Linux; Android 4.4.2; SM-G955F Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36  pddmt_android_version/2.2.3 pddmt_android_build/2025 pddmt_android_channel/yybandroid Mozilla/5.0 (Linux; Android 4.4.2; SM-G955F Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36  pddmt_android_version/2.2.3 pddmt_android_build/2025 pddmt_android_channel/yyb
ETag: TDh2Hsue
PASSID: bapp_1-qfuJ+XmbTsdk4onr3Nlf/vnfRj7VDZ9AdJVTpPQmaKzxq1NxX8DiwLGNqwnl6AO0ThltKGgfpcWRXNWobxcyxg_188639248_17085632
Content-Type: application/json;charset=UTF-8
Content-Length: 68
Host: mms.pinduoduo.com
Connection: close
Cookie: api_uid=rBQH01u/V4xeZiVSDExjAg==; PASS_ID=bapp_1-qfuJ+XmbTsdk4onr3Nlf/vnfRj7VDZ9AdJVTpPQmaKzxq1NxX8DiwLGNqwnl6AO0ThltKGgfpcWRXNWobxcyxg_188639248_17085632

{"is_onsale":1,"sold_out":0,"page":1,"mall_id":"188639248","size":1}


日经营报表

POST /sydney/api/dailyMallGoods/dailyReport HTTP/1.1
Referer: Android
User-Agent: androidandroid Mozilla/5.0 (Linux; Android 4.4.2; SM-G955F Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36  pddmt_android_version/2.2.3 pddmt_android_build/2025 pddmt_android_channel/yybandroid Mozilla/5.0 (Linux; Android 4.4.2; SM-G955F Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36  pddmt_android_version/2.2.3 pddmt_android_build/2025 pddmt_android_channel/yyb
ETag: TDh2Hsue
PASSID: bapp_1-qfuJ+XmbTsdk4onr3Nlf/vnfRj7VDZ9AdJVTpPQmaKzxq1NxX8DiwLGNqwnl6AO0ThltKGgfpcWRXNWobxcyxg_188639248_17085632
Content-Type: application/json;charset=UTF-8
Content-Length: 26
Host: mms.pinduoduo.com
Connection: close
Cookie: api_uid=rBQH01u/V4xeZiVSDExjAg==; PASS_ID=bapp_1-qfuJ+XmbTsdk4onr3Nlf/vnfRj7VDZ9AdJVTpPQmaKzxq1NxX8DiwLGNqwnl6AO0ThltKGgfpcWRXNWobxcyxg_188639248_17085632

{"queryDate":"2018-10-20"}
