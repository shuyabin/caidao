> 菜鸟级golang菜刀

因为不想装java的环境... 那么大.. 所以就写了一个菜鸟级的golang菜刀小工具

> 安装

```
下载代码后 执行

go get -u github.com/astaxie/beego
go get -u github.com/mattn/go-sqlite3 ( 注意这里需要gcc windwos的话  http://tdm-gcc.tdragon.net/download )
go get -u github.com/tidwall/gjson

go run main.go
或者
go build 
打开网页 127.0.0.1:8080

```
> 截图

![转码](/path/to/img.jpg)

> 或者直接使用?

```
到当前目中cmd 执行 caidao.exe   (windows)
```

> 上传见不得人的小文件

```
.. 这里就不讨论怎么能通过x洞上传小文件
上传文件 目前我就写了php的版本 cus.php (主要文件) file.php (依赖库) conso.php (终端)

终端 账号密码  root  root

如果不改文件的话。。添加url的时候 密码 Cknife
```

![image](https://github.com/shuyabin/caidao/blob/master/images/shouye.png)
![image](https://github.com/shuyabin/caidao/blob/master/images/add.png)
![image](https://github.com/shuyabin/caidao/blob/master/images/zhongduan.png)
![image](https://github.com/shuyabin/caidao/blob/master/images/zhongduan1.png)


> 后面需要做的

```
1. 终端目录可配置
2. 中午文件上传
3. 大文件上传处理
4. 客户端eval处理
5. 目录树优化( 目前的js只能是吧所有文件一次性递归循环出来)
6. 客户端化 webssh 终端 并且支持windows
```