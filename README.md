# 后端技术方案

## 用户登录/注册

表设计：
首先创建 movie 数据库
```sql
CREATE DATABASE IF NOT EXISTS movie DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
```

创建 user 表
```sql
CREATE TABLE `user` (
`uid` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
`userName` varchar(32) NOT NULL COMMENT '用户名',
`password` varchar(32) NOT NULL COMMENT '密码',
`role` tinyint(4) unsigned NOT NULL COMMENT '角色',
`createTime` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
`updateTime` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '更新时间',
PRIMARY KEY (`uid`),
UNIQUE KEY `userName` (`userName`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```
role: 1-guest 2-user 3-admin

用户注册时，向表内写入用户信息
```sql
INSERT INTO user 
(
    userName,
    password,
    role,
    createTime,
    updateTime
)
VALUES
(
    'guest',
    '',
    1,
    1658762844,
    1658762844
),
(
    'crotaliu',
    '12345678',
    3,
    1658762844,
    1658762844
);
```

前端需要校验用户名是否重复
```sql
SELECT uid FROM user WHERE userName='crotaliu';
```

用户登录时，获取用户输入的用户名，查询 user 表，对比 password 是否相同
```sql
SELECT password FROM user WHERE userName='crotaliu';
```

如果密码一致，则将 userName 登录时间 随机 uuid，使用 base64 加密生成 token，记录到 token 表中，并返回到登录成功的 res 中；

创建 token 表
```sql
CREATE TABLE `token` (
`token` varchar(255) NOT NULL COMMENT '用户token',
`userName` varchar(32) NOT NULL COMMENT '用户名',
`loginTime` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '登录时间',
PRIMARY KEY (`token`),
UNIQUE KEY `userName` (`userName`),
KEY `index_loginTime` (`loginTime`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

```sql
INSERT INTO token 
(
    token,
    userName,
    loginTime
)
VALUES
(
    'dXNlck5hbWU9Y3JvdGFsaXU7bG9naW5UaW1lPTE2NTkyMDAyMDI7dXVpZD0wYWYzOWUzYS1lNDE2LTRlYzAtYTExYi01MDliNTAxOTA0YTY=',
    'crotaliu',
    1658762844
);
```

原文：userName=crotaliu;loginTime=1658762844;uuid=38f60189-a8bd-4eee-95bd-cd20534475f1

token：dXNlck5hbWU9Y3JvdGFsaXU7bG9naW5UaW1lPTE2NTkyMDAyMDI7dXVpZD0wYWYzOWUzYS1lNDE2LTRlYzAtYTExYi01MDliNTAxOTA0YTY=

查找密码正确的用户名，删除同名的 token
```sql
DELETE FROM token WHERE userName='crotaliu';
```

前端将 token 写入 localStorage 中，每次请求在 header 中加入 x-token 请求头，值为 token；

获取到 token 后，解密出用户名，从而获取用户角色；
```sql
SELECT role FROM user WHERE userName='crotaliu';
```

```sql
SELECT token FROM token WHERE token='dXNlck5hbWU9Y3JvdGFsaXU7bG9naW5UaW1lPTE2NTkyMDAyMDI7dXVpZD0wYWYzOWUzYS1lNDE2LTRlYzAtYTExYi01MDliNTAxOTA0YTY=';
```
校验 token，每次接口请求前，先校验 token 是否在表中，token 错误的逻辑：
- 用户接口，注册、登录、检查用户名重复接口，忽略 token
- 客户端接口，电影列表、数据源、点赞、收藏、电影详情接口，token 错误时，返回 http status code 200，响应体中返回 { code: 301, result: null, msg: 'token过期，请重新登录' }，刷新当前页面
- 管理端接口，电影添加、删除、上线/下线、编辑、数据源删除、更新、添加接口，token 错误时，返回 http status code 403，跳转回首页

定期清理 token，设置定时任务，每小时执行 1 次
```sql
DELETE FROM token WHERE loginTime >= now() - 7 * 24 * 60 * 60 * 1000;
```

## 客户端

### 电影列表

表设计：
创建 list 表
```sql
CREATE TABLE `list` (
`mid` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '电影ID',
`mURL` varchar(255) NOT NULL COMMENT '电影URL',
`mName` varchar(255) NOT NULL COMMENT '电影名',
`mPoster` varchar(255) NOT NULL COMMENT '电影海报',
`mTypeID` tinyint(4) unsigned NOT NULL COMMENT '电影类型ID',
`mTypeName` varchar(10) NOT NULL COMMENT '电影类型',
`mDoubanScore` float(2,1) NOT NULL COMMENT '豆瓣评分',
`mDirector` varchar(50) NOT NULL COMMENT '电影导演',
`mStarring` varchar(255) NOT NULL COMMENT '电影主演',
`mCountryID` tinyint(4) unsigned NOT NULL COMMENT '电影制片国家/地区ID',
`mCountryName` varchar(10) NOT NULL COMMENT '电影制片国家/地区',
`mLanguageID` tinyint(4) unsigned NOT NULL COMMENT '电影语言ID',
`mLanguageName` varchar(10) NOT NULL COMMENT '电影语言',
`mDateYear` smallint(6) unsigned NOT NULL COMMENT '电影上映年份',
`mDate` varchar(10) NOT NULL COMMENT '电影上映日期',
`mDesc` varchar(1024) NOT NULL COMMENT '电影简介',
`mStatus` tinyint(4) unsigned NOT NULL COMMENT '电影状态',
`createTime` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
`updateTime` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '更新时间',
PRIMARY KEY (`mid`),
KEY `index_name` (`mName`),
KEY `index_mTypeID` (`mTypeID`),
KEY `index_mDouBanScore` (`mDouBanScore`),
KEY `index_mCountryID` (`mCountryID`),
KEY `index_mLanguageID` (`mLanguageID`),
KEY `index_mDateYear` (`mDateYear`),
KEY `index_mStatus` (`mStatus`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```
status: 1-上线 2-下线

// 创建电影时，上映日期取年份，获取到上映年份

#### 获取电影列表

```sql
SELECT
    baseTable.mid,
    baseTable.mName,
    baseTable.mPoster,
    baseTable.mTypeName,
    baseTable.mDouBanScore,
    baseTable.mDirector,
    baseTable.mStarring,
    baseTable.mCountryName,
    baseTable.mLanguageName,
    baseTable.mDateYear,
    baseTable.mDate,
    viewTable.mViews,
    likeTable.mLikes,
    collectTable.mCollects
FROM list AS baseTable

LEFT JOIN (SELECT mid, COUNT(*) AS mViews FROM record WHERE type = 3 GROUP BY mid) AS viewTable
ON baseTable.mid = viewTable.mid

LEFT JOIN (SELECT mid, COUNT(*) AS mLikes FROM record WHERE type = 1 GROUP BY mid) AS likeTable
ON baseTable.mid = likeTable.mid

LEFT JOIN (SELECT mid, COUNT(*) AS mCollects FROM record WHERE type = 2 GROUP BY mid) AS collectTable
ON baseTable.mid = collectTable.mid

WHERE baseTable.mid >= (SELECT mid FROM list WHERE mStatus = 1 AND mTypeID = 1 LIMIT (pageNo - 1) * pageSize,1)
LIMIT pageSize;
```

#### 计算电影总数

```sql
SELECT count(*) FROM list WHERE mStatus=1 AND mTypeID = 1;
```

#### 电影详情

```sql
SELECT
    mid,
    mUrl,
    mName,
    mPoster,
    mTypeName,
    mDoubanScore,
    mDirector,
    mStarring,
    mCountryName,
    mLanguageName,
    mDate,
    mDesc
FROM list WHERE mid = 1;
```

### 收藏/点赞/播放

创建点赞/收藏/播放表
```sql
CREATE TABLE `record` (
`userName` varchar(10) NOT NULL COMMENT '用户名',
`mid` int(10) NOT NULL COMMENT '电影ID',
`type` tinyint(4) unsigned NOT NULL COMMENT '点赞/收藏/播放',
`createTime` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
UNIQUE KEY `userName_mid_type` (`userName`, `mid`, `type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```
type: 1-点赞 2-收藏 3-播放

#### 播放榜

```sql
SELECT
    baseTable.mid,
    baseTable.mName,
    totalTable.mTotal,
FROM list AS baseTable

RIGHT JOIN (SELECT mid, COUNT(*) AS mTotal FROM record WHERE type = 3 GROUP BY mid ORDER BY mTotal DESC LIMIT 5) AS totalTable
ON baseTable.mid = totalTable.mid
WHERE mStatus = 1
```

#### 收藏榜

```sql
SELECT
    baseTable.mid,
    baseTable.mName,
    totalTable.mTotal,
FROM list AS baseTable

RIGHT JOIN (SELECT mid, COUNT(*) AS mTotal FROM record WHERE type = 2 GROUP BY mid ORDER BY mTotal DESC LIMIT 5) AS totalTable
ON baseTable.mid = totalTable.mid
WHERE mStatus = 1
```

#### 点赞榜

```sql
SELECT
    baseTable.mid,
    baseTable.mName,
    totalTable.mTotal,
FROM list AS baseTable

RIGHT JOIN (SELECT mid, COUNT(*) AS mTotal FROM record WHERE type = 1 GROUP BY mid ORDER BY mTotal DESC LIMIT 5) AS totalTable
ON baseTable.mid = totalTable.mid
WHERE mStatus = 1
```

#### 点赞

```sql
-- 查询
SELECT mid FROM record WHERE userName='crotaliu' AND mid = 1 AND type = 1;

-- 记录
INSERT INTO record 
(
    userName,
    mid,
    type,
    createTime
)
VALUES
(
    'crotaliu',
    1,
    1,
    1658762844
);
-- 取消
DELETE FROM record WHERE userName='crotaliu' AND mid = 1 AND type = 1;
```

#### 收藏

```sql
-- 查询
SELECT mid FROM record WHERE userName='crotaliu' AND mid = 1 AND type = 2;

-- 记录
INSERT INTO record 
(
    userName,
    mid,
    type,
    createTime
)
VALUES
(
    'crotaliu',
    1,
    2,
    1658762844
);
-- 取消
DELETE FROM record WHERE userName='crotaliu' AND mid = 1 AND type = 2;
```

#### 播放

```sql
-- 查询
SELECT mid FROM record WHERE userName='crotaliu' AND mid = 1 AND type = 1;

-- 记录
INSERT INTO record 
(
    userName,
    mid,
    type,
    createTime
)
VALUES
(
    'crotaliu',
    1,
    3,
    1658762844
);
-- 取消
DELETE FROM record WHERE userName='crotaliu' AND mid = 1 AND type = 3;
```

## 管理端

### 电影添加

```sql
INSERT INTO list 
(
    mURL,
    mName,
    mPoster,
    mTypeID,
    mTypeName,
    mDouBanScore,
    mDirector,
    mStarring,
    mCountryID,
    mCountryName,
    mLanguageID,
    mLanguageName,
    mDateYear,
    mDate,
    mDesc,
    mStatus
)
VALUES
(
    'https://www.baidu.com/',
    '阿甘正传',
    'https://www.baidu.com/',
    1,
    '喜剧',
    9.5,
    '罗伯特·泽米基斯',
    '埃尔维斯·普雷斯利|汤姆·汉克斯|莎莉·菲尔德|库尔特·拉塞尔',
    1,
    '美国',
    1,
    '英语',
    1994,
    '1994-07-01',
    '阿甘是个智商只有75的低能儿。在学校里为了躲避别的孩子的欺侮，听从一个朋友珍妮的话而开始“跑”。他跑着躲避别人的捉弄。在中学时，他为了躲避别人而跑进了一所学校的橄榄球场，就这样跑进了大学。阿甘被破格录取，并成了橄榄球巨星，受到了肯尼迪总统的接见。',
    1
);
```

### 电影更新

```sql
UPDATE list
SET
mURL='https://www.baidu.com/',
mName='阿甘正传',
updateTime=1658762844
WHERE mid=1;
```

### 电影上/下线

```sql
UPDATE list
SET
mStatus=1,
updateTime=1658762844
WHERE mid=1;
```

### 电影删除

```sql
DELETE FROM list WHERE mid=1;
```

### 属性建表

```sql
CREATE TABLE `prop` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '属性ID',
`label` varchar(10) NOT NULL COMMENT '属性名',
`value` int(10) NOT NULL COMMENT '属性值',
`createTime` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
`updateTime` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间'
PRIMARY KEY (`id`),
KEY `index_value` (`value`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 属性添加

```sql
INSERT INTO prop 
(
    label,
    value,
    createTime,
    updateTime
)
VALUES
(
    '喜剧',
    1,
    1658762844,
    1658762844
);
```

### 属性更新

```sql
UPDATE prop
SET
label='科幻',
value=2,
updateTime=1658762844
WHERE id=1;
```

### 属性删除

```sql
DELETE FROM prop WHERE id=1;
```
