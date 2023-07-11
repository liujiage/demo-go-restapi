create table user(
		id varchar(50) primary key, 
		name varchar(500), 
		createTime datetime DEFAULT (datetime('now', 'localtime')),
		updateTime datetime
)