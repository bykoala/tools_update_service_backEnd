use update_data;
create table updateinfo(
                           `id` int primary key auto_increment comment '更新信息表id',
                           `version` varchar(10) not null default 'v1.0.0' comment '版本号',
                           `content` varchar(1000) not null default '突然想跟你发生一下关系' comment '更新内容',
                           `forced` tinyint(1) not null default 0 comment '是否强制更细，0:否；1:是',
                           `status` tinyint(1) not null default 0 comment '是否更新，0:不更新；1:更新',
                           `url` varchar(200) not null default 'tongzhi.wang' comment '下载链接',
                           `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间'
);

select * from updateinfo order by id desc limit 1;