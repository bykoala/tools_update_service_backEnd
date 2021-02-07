use update_data;

CREATE TABLE `updateinfo` (
                              `id` int NOT NULL AUTO_INCREMENT COMMENT '更新信息表id',
                              `version` varchar(10) NOT NULL DEFAULT 'v1.0.0' COMMENT '版本号',
                              `content` varchar(1000) NOT NULL DEFAULT '突然想跟你发生一下关系' COMMENT '更新内容',
                              `forced` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否强制更细，0:否；1:是',
                              `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否更新，0:不更新；1:更新',
                              `url` varchar(200) NOT NULL DEFAULT 'tongzhi.wang' COMMENT '下载链接',
                              `size` float NOT NULL DEFAULT '0',
                              `md5` varchar(128) NOT NULL DEFAULT '' COMMENT '文件md5值',
                              `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

create table updateinfo(
                           `id` int primary key auto_increment comment '更新信息表id',
                           `version` varchar(10) not null default 'v1.0.0' comment '版本号',
                           `content` varchar(1000) not null default '终于可以跟你发生一下关系了' comment '更新内容',
                           `forced` tinyint(1) not null default 0 comment '是否强制更细，0:否；1:是',
                           `status` tinyint(1) not null default 0 comment '是否更新，0:不更新；1:更新',
                           `url` varchar(200) not null default 'tongzhi.wang' comment '下载链接',
                           `size` float not null default 0.00 comment '文件大小',
                           `create_time` timestamp Not NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间'
)CHARACTER SET utf8 COLLATE utf8_general_ci;

select * from updateinfo order by id desc limit 1;

create table feedback(
    `id` bigint primary key auto_increment comment '反馈表id',
    `postId` bigint  not null unique comment '反馈帖子id',
    `title`        varchar(128)                        not null comment '标题',
    `content`      varchar(8192)                       not null comment '内容',
    `author_id`    bigint                              not null comment '作者的用户id',
    `classification` tinyint not null default 0 comment '帖子分类，0:建议；1:反馈bug',
    `status`       tinyint   default 1                 not null comment '帖子状态',
    `create_time`  timestamp default CURRENT_TIMESTAMP not null comment '创建时间',
    `update_time`  timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
)