DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT primary key,
  `username` varchar(50) DEFAULT '' COMMENT '用户账号',
  `password` varchar(50) DEFAULT '' COMMENT '用户密码',
  `display_name` varchar(50) DEFAULT '' COMMENT '用户显示名称',
  `dept_id` int not null comment '部门id',
  `dept_name` varchar(50) DEFAULT '' COMMENT '部门名称',
  `role_id` int not null comment '角色id',
  `role_name` varchar(50) DEFAULT '' COMMENT '角色名称',
  `create_at` datetime default CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime default CURRENT_TIMESTAMP COMMENT '修改时间',
  `is_deleted` int default '0' COMMENT '删除标志位，0代表正常，1代表已删除',
  unique key `UK_username_dept_role` (`username`,`dept_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment '用户信息表';