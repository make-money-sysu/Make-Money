create table user
(
	id int unsigned not null primary key,
	password varchar(40),
	real_name varchar(40),
	nick_name varchar(40),
	age smallint unsigned,
	gender char(1),
	head_picture blob,
	balance float default 0.0,
	profession varchar(40),
	grade varchar(20),
);
create table survey
(
	id int unsigned not null auto_increment primary key auto_increment,
	publisher_id int unsigned,
	name varchar(40),
	content varchar(1000),
	foreign key(publisher_id) references user(id)
);
create table do_survey
(
	survey_id int unsigned not null,
	recipient_id int unsigned not null,
	primary key(survey_id,recipient_id),
	foreign key(survey_id) references survey(id),
	foreign key(recipient_id) references user(id)
);
create table friends
(
	user1_id int unsigned,
	user2_id int unsigned,
	accepted boolean,
	primary key(user1_id,user2_id),
	foreign key(user1_id) references user(id),
	foreign key(user2_id) references user(id)
);