use spocon;

create table users
(
  id varchar(255) not null,
  name varchar(255) not null,
  email varchar(255) not null unique,
  team_id varchar(255) not null,
  created_at TIMESTAMP not null default CURRENT_TIMESTAMP,
  updated_at TIMESTAMP not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
  primary key (id),
  foreign key (team_id) references teams(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
