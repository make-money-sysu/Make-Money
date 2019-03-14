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
	name varchar(40) not null,
	content varchar(1000) not null,
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
	user1_id int unsigned not null,
	user2_id int unsigned not null,
	accepted boolean not null,
	primary key(user1_id,user2_id),
	foreign key(user1_id) references user(id),
	foreign key(user2_id) references user(id)
);