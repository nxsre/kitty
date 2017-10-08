### 安装方法
#### 1. 安装源码包

```go
go get -v -x github.com/soopsio/kitty
```
#### 2. 导入数据库

```ps
#powershell
cd $env:GOPATH/src/github.com/soopsio/kitty
mysql.exe -uroot -e "create database if not exists kitty;grant all on kitty.* to kitty@'localhost' identified by '123456'";
mysql.exe -uroot kitty -e "source .\db\kitty.sql"
bee run
```

#### 打开浏览器

```
http://127.0.0.1:8080/
```


### 修复的bug

#### 1. 表结构问题

###### 修复 Error 1054: Unknown column 'job_id' in 'field list'

```
ALTER TABLE `kitty`.`job_snapshot`   
  ADD COLUMN `job_id` INT(11) NOT NULL COMMENT '任务ID' AFTER `id`;
```

  
#### 2. 包导入问题
```
package kitty/app/controller: unrecognized import path "kitty/app/controller" (import path does not begin with hostname)
package kitty/app/job: unrecognized import path "kitty/app/job" (import path does not begin with hostname)
package kitty/app/service: unrecognized import path "kitty/app/service" (import path does not begin with hostname)
```

### TODO

#### 1. 任务快照历史 JobSnapshotHistory

#### 2. 任务修改历史 JobInfoHistory

#### 3. 用户鉴权