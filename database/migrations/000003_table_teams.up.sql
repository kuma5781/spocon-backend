use spocon;

create table teams
(
  id varchar(255) not null,
  uuid char(36) not null unique,
  name varchar(255) not null,
  sport_id varchar(255) not null,
  grade_id varchar(255) not null,
  reception_status enum('on','off') default 'on',
  icon_path varchar(255) not null,
  description varchar(1023) not null,
  address_state varchar(255) not null,
  address_city varchar(255) not null,
  created_at TIMESTAMP not null default CURRENT_TIMESTAMP,
  updated_at TIMESTAMP not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
  primary key (id),
  foreign key (sport_id) references sports(id),
  foreign key (grade_id) references grades(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
