这个readme不是说明,是计划.

> 创建备份表

create table xxx as select * from .....

执行完之后,
truncate table xxx 清空本表

或者

更改表名,再创建新表

ALTER  TABLE table_name RENAME TO new_table_name

需要测试索引的问题

> 版本管理表

表,上一版本,当前版本是否完成 ,更改日期

当前版本,没版本表示

如果未更新完成,就做上一版本的数据, 表+当前版本

> 爬百科只管一页即可.不再处理版本号

> 可以物化的数组
