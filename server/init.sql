create table user
(
	id int unsigned not null primary key,
	password varchar(40) not null,
	real_name varchar(40) not null,
	nick_name varchar(40) not null,
	age smallint unsigned not null,
	gender char(1) not null,
	head_picture blob not null,
	balance float default 0.0,
	profession varchar(40) not null,
	grade varchar(20) not null,
	phone varchar(20) not null,
	email varchar(40) not null
);

create table survey
(
	id int unsigned not null auto_increment primary key auto_increment,
	publisher_id int unsigned not null,
	title varchar(40) not null,
	content varchar(1000) not null,
	state int not null default 0,
	checked int not null default 0,
	create_time datetime,
	foreign key(publisher_id) references user(id)
);

create table do_survey
(
	survey_id int unsigned not null,
	recipient_id int unsigned not null,
	content varchar(1000) not null,
	primary key(survey_id,recipient_id),
	foreign key(survey_id) references survey(id),
	foreign key(recipient_id) references user(id)
);
create table friends
(
	fid	int unsigned not null auto_increment primary key,
	user1_id int unsigned not null,
	user2_id int unsigned not null,
	accepted boolean not null,
	-- primary key(user1_id,user2_id),
	foreign key(user1_id) references user(id),
	foreign key(user2_id) references user(id)
);

create table package
(
	id int unsigned not null auto_increment primary key,
	owner_id int unsigned not null,
	receiver_id int unsigned,
	create_time datetime,
	reward float not null,
	-- 0为刚发布，1为已被接单，2为确认送达
	state int unsigned default 0,
	note varchar(200) default "",
	foreign key(owner_id) references user(id),
	foreign key(receiver_id) references user(id)
)


create table msg
(
	mid int not null auto_increment primary key,
	fromid int unsigned not null,
	toid int unsigned not null,
	create_time datetime,
	-- 0为系统消息 ，10为未查看，11为已查看，但未知悉，12为已查看，已知悉，13为已撤回
	state int unsigned default 10,
	content varchar(140) not null,
	foreign key(fromid) references user(id),
	foreign key(toid) references user(id)
)