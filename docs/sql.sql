-- noinspection SqlNoDataSourceInspectionForFile

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
    `title` varchar(100) DEFAULT '' COMMENT '文章标题',
    `desc` varchar(255) DEFAULT '' COMMENT '简述',
    `content` text COMMENT '内容',
    `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0',
    `is_del` int(10) unsigned DEFAULT '0' COMMENT '是否删除，0为未删除，1为已删除',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态0为禁用，状态1为启用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) DEFAULT '' COMMENT '标签名称',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del` int(10) unsigned DEFAULT '0' COMMENT '是否删除，0为未删除，1为已删除',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';


-- ----------------------------
-- Table structure for blog_article_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_article_tag`;
CREATE TABLE `blog_article_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `article_id` int(11) NOT NULL COMMENT '文章ID',
    `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '标签ID',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签关联表';



-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth` (
                             `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                             `app_key` varchar(50) DEFAULT '' COMMENT 'key',
                             `app_secret` varchar(50) DEFAULT '' COMMENT 'secret',
                             `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
                             `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
                             `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
                             `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
                             `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
                             `is_del` int(10) unsigned DEFAULT '0' COMMENT '是否删除，0为未删除，1为已删除',
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT="认证管理";

INSERT INTO `blog_auth` (`id`, `app_key`, `app_secret`, `created_on`, `created_by`, `modified_on`, `modified_by`,
`deleted_on`) VALUES ('1', 'eddycjy', 'go-programming-tour-book', 0, 'eddycjy', 0, '', 0);
