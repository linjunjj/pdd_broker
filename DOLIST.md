##客服集成系统
---
+ websocket 协议对接
 +  消息发送入库
 +  消息用户识别
 +  消息内容过滤拦截
 +  消息连接逻辑处理
 +  HTTP控制管理界面
 +  UI界面

##商品管理系统
---
+ 1.系统登陆绑定
  + 1.2 用户分组权限
     + 1.2.1 销售
     + 1.2.2 客服
     + 1.2.3 售后
     + 1.2.4 总后台
  + 用户绑定表列
+ 接口对接功能封装
 + 扫码登陆接口封装
 + 获取商品ID数据
 + 上传/修改/删除商品
 + 库存(Skuid)控制管理
 + 销售统计
 + 物流管理


##数据库表
---
+ sys_user
	+ id,username,password,token,time,tag,wslink
+ sys_sgb
	+ id,pdd_id,other\_id,catid,skuid,title,pic,text,amount,money,time
+ sys_order
	+ id,pdd_id,order\_id,goods\_id,dizi,address,phone,name,money,time
+ user_data
	+ id,user,pass,token,cookie,header,ip,sys_user\_id,proxy\_id
+ user_address
	+ id,user,address1,address2,address3
+ user_info
	+ id,user,mallinfo,phone,company,
+ user_sum
	+ id,time,money1,money2,money3,order_sum,vistor\_sum,percent,conversion
+ crawl_taobao
	+ ...
+ crawl_jd
	+ ...
+ taobao_catid
	+ id,name,level
+ pdd_catid
	+ id,name,level
+ chat_log
	+ id,text,time,token,user,to_user,from\_user
+ proxy_ip
	+ proxy_id,host,port,type

##客服集成系统
---
+ 业务流程
	+ 拼多多服务器 > 后台控制端 > 服务器逻辑层 > 自封装协议 > 返回客户
+ 自定义简化封装协议
+ 控制器方法
	+ 认证账号信息 > 查询获取token 表 > 并发开启长连维持 > 入库排序 | 接受数据发送| 


