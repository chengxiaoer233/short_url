<div align="center">

# 短连接系统实现
</div>


<div align="center">

![](https://img.shields.io/github/languages/code-size/chengxiaoer233/short_url?label=CodeSize)
![](https://img.shields.io/github/stars/chengxiaoer233/short_url?label=GitHub)
![](https://img.shields.io/github/watchers/chengxiaoer233/short_url?label=Watch)
[![Go Report Card](https://goreportcard.com/badge/github.com/chengxiaoer233/short_url)](https://goreportcard.com/report/github.com/chengxiaoer233/short_url)
[![LICENSE](https://img.shields.io/badge/license-MIT-green)](https://mit-license.org/)
</div>


<div align="center">

<img  src="https://my-source666.obs.cn-south-1.myhuaweicloud.com/myBlog/golang-jixiangwu-image.png" width="600" height="350"/>

</div>

## 1：需求评估： 
 
  * 输入输出  
    + 1：输入一个长链接地址，返回端链接地址
    + 2：多次输入同一个长连接地址，返回相同的短连接地址
    + 3：输入短连接地址，能够访问长连接资源

## 2：系统约束和要求拆解
  + 1：生成的短地址是否需要设置过期时间？设置多久？
  + 2：长连接生成的短连接必须是唯一的，不能存在重复问题
  + 3：是否需要支持自定义短连接的域名？
  + 4：生成的短连接的path长度是否有限制？组成元素是否有限制
  + 5：读写qps？
  + 6：系统允许延时是多少？
  + 7：需要消耗的存储空间是多少？
  + 8：系统可靠性？几个九？
  + 9：安全性：别人能否爬取接口
  
## 3: 简单方案设计
  * 转换接口（写接口，长连接转换为短连接）
    + 1：判断输入的长连接是否已经存储过，存在则直接返回数据
    + 2：不存在则根据一定规则生成短连接，并持久化到db中
    
  * 读接口（短连接转换为长连接）
    + 1：获取短域名，判断是否已经存在，存在则返回，不存在则报错
    + 2：存在则获取后进行302重定向

## 4：需求拆解处理

   * （1）过期时间的设置和处理问题   
    
        + 1：延时删除：业务再次访问的时候，去执行判断和删除动作  
            * 优点   
               + 1：性能损失最小，不用实时计算，实现简单方便  
             * 缺点   
               + 1：如果业务没有继续访问，就不会再次判断，会浪费存储空间  
                
       + 2：定时删除：每个url一个定时器，时间过了就删除  
            * 优点   
                + 1：存储空间的利用率最高  
            * 缺点  
                + 1：每个url需要开启一个定时器，浪费cpu
  
       + 3：轮询删除：启动一个后台线程，不断的扫描所有元素，折中方案
           
           * 总结： 我们这里使用的url不设置过期时间，默认一直有效    
    
           
   * （2）短连接是否需要支持用户自定义域名
       + 1：只要用户的域名已经备案了就可以自定义，同时将自定义域名解析到短连服务器，才可以解析。
   
   * （2）长连接变短连接，如何保证唯一性？ 
        * 1：生成时  
            + 1：判断长连接是否存在（布隆过滤器），不存在则可以进行生成动作，存在（假阳）先查询Redis（查询长连接key），存在返回，
                    不存在则查MySQL，存在返回，不存在，则返回第一步，长连接地址固定加个字符串，继续流程。 
            
            + 2： 唯一ID的生成算法
                + UUID：生成后截取固定长度，可能会冲突，没法保证唯一性
                + 哈希：murmurhash等，存在一定概率的冲突，在写入小时不会冲突
                + 自增ID：可以生成唯一ID，一般利用数据库生成，效果差
        
        * 总结：     
            + 1：生成前我们需要判断是否重生成（布隆过滤器 + Redis, redis存kv，两份（长短相互映射））
            + 2：              