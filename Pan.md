网盘文档

目前实现的功能

1、文件上传 下载

2、文件分享

3、登录注册

4、文件权限管理

5、一次性快传

6、加密分享链接

//TO DO

1、断点续传

2、token的复杂处理

## API接口介绍

## 用户注册

**简要描述：**

- 用户注册接口

**请求URL：**

- localhost:8080/user/signup

**请求方式：**

- POST

**参数：**

| 参数名   | 必选 | 类型   | 说明   |
| -------- | ---- | ------ | ------ |
| username | 是   | string | 用户名 |
| password | 是   | string | 密码   |



请求示例

就是 Body的

![1566658580958](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566658580958.png)



返回示例

![1566658695476](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566658695476.png)

## 用户登录

**简要描述：**

- 用户注册接口

**请求URL：**

- localhost:8080/user/signin

**请求方式：**

- POST

**参数：**

| 参数名   | 必选 | 类型   | 说明   |
| -------- | ---- | ------ | ------ |
| username | 是   | string | 用户名 |
| password | 是   | string | 密码   |



请求示例

![1566658751889](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566658751889.png)

返回示例：成功

![1566658778160](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566658778160.png)

​                   失败

![1566658945858](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566658945858.png)

## 用户个人信息

**简要描述：**

- 文件隔离，展示用户个人信息及文件

**请求URL：**

- localhost:8080/user/info

**请求方式：**

- POST

**参数：**

| 参数名   | 必选 | 类型   | 说明     |
| -------- | ---- | ------ | -------- |
| username | 是   | string | 用户名   |
| token    |      |        | 未完善好 |

请求示例

![1566659890015](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566659890015.png)

## 文件上传

**简要描述：**

- 简单的文件上传

**请求URL：**

- http://localhost:8080/file/upload

**请求方式：**

- POST

**参数：**

| 参数名   | 必选 | 类型   | 说明               |
| -------- | ---- | ------ | ------------------ |
| file     | 是   | file   | 拉取文件           |
| username | 是   | string | 上传的用户名       |
| token    |      |        | 未完善好，暂时不要 |

请求示例

![1566660005247](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566660005247.png)

返回示例  成功

![1566660039383](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566660039383.png)

## 文件分享

**简要描述：**

- 简单的文件分享，这里的分享就是自己选取文件，前端给他filehash 他拿到这个file的基本云数据。

  然后分享方式就把这个链接分享给别人就行

**请求URL：**

- http://localhost:8080/file/meta

**请求方式：**

- POST

**参数：**

| 参数名   | 必选 | 类型 | 说明             |
| -------- | ---- | ---- | ---------------- |
| filehash | 是   | file | file进行sha1加密 |



![1566660169657](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566660169657.png)

## 文件下载

**简要描述：**

- 简单的文件下载

**请求URL：**

- http://localhost:8080/file/download

**请求方式：**

- POST

**参数：**

| 参数名   | 必选 | 类型 | 说明             |
| -------- | ---- | ---- | ---------------- |
| filehash | 是   | file | file进行sha1加密 |

这个到浏览器里操作，通过os.Read拿到文件

## 文件更新

**简要描述：**

- 简单的文件更新，只能更新文件名

**请求URL：**

- http://localhost:8080/file/update

**请求方式：**

- POST

![1566662186888](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566662186888.png)

## 文件删除

**简要描述：**

- 简单的文件删除，根据filehash删除文件

**请求URL：**

- http://localhost:8080/file/delete

**请求方式：**

- POST
- ![1566662371947](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566662371947.png)

## 文件快传

**简要描述：**

- 文件快传，如果服务端的文件表中有hash相同的文件，则可以快传

**请求URL：**

- http://localhost:8080/file/fastupload

**请求方式：**

![1566662468392](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566662468392.png)

如果没有的话

![1566662490635](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566662490635.png)



### 表的设计

![1566662583876](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566662583876.png)

具体的创表文件在项目目录下

![1566662623933](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566662623933.png)

### 文件的加密处理

主要就是把上传的文件变成一个文件元的结构体，进行sha1的加密

![1566662686353](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566662686353.png)

### 统一鉴权和Token的处理

使用MD5加密

![1566662737145](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566662737145.png)

![1566662751748](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566662751748.png)

统一鉴权拦截

![1566662801068](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566662801068.png)

![1566662822640](C:\Users\win10\AppData\Roaming\Typora\typora-user-images\1566662822640.png)

### 今晚继续搞还没做好的Docker



已经搞了很久了，服务器上能跑dockerfile但是应该是gopath路径问题 run不起来，比较遗憾还是没有生成镜像，

项目里给了一个简单的dockerfile

mysql的docker镜像直接pull能跑，但也没有装进我的项目里去

整体上还是一个只能本地运行的项目

### TO DO

之前的大作业，由于自己项目原因被获准写稍微简单点，结果由于自己偷懒以为这样可以少做点工作划水了。展示的时候感觉自己特别丢人。这次的考核没有时间和其他的差异，也不知道自己和别人比起来写的怎么样

就我自己而言，最难受的还是docker还没跑起来，我感觉这是我应该做到的

昨天已经熬夜挺晚了，今天不想熬夜太晚了，争取把Docker弄好睡觉

