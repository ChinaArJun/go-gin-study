create table `blog_tag`  (
    `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name` varchar(100) DEFAULT '' COMMENT '标签名称',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `created_at` timestamp  NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
    `deleted_on` int(10) DEFAULT '0' COMMENT '类型 0默认 1删除',
    `state` int(10) unsigned default '0' COMMENT '状态 0禁用 1启用',
    PRIMARY KEY (`id`),
    KEY `state` (`state`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT '标签表';

CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text,
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
	`created_at`  timestamp NULL DEFAULT NULL COMMENT '创建时间',
	`modified_at` timestamp NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  PRIMARY KEY (`id`),
	KEY `state` (`state`) USING BTREE,
	KEY `tag_id` (`tag_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT '文章表';

CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT '认证表';

INSERT INTO `blog`.`blog_auth` (`id`, `username`, `password`) VALUES (null, 'test', 'test123456');
