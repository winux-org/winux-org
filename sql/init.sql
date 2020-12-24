create database ideas;

\connect ideas;

create table ideas (
  id bigserial, -- generated always as identity,
  idea varchar(1024) not null,
  email varchar(128) not null,

  primary key(id)
);

create table ups (
  idea_id serial8 not null,

  constraint fk_idea_id
    foreign key(idea_id)
      references ideas(id)
);


insert into ideas (idea, email) values
  ('coś 34', 'artur.gurgul@gmail.com'),
  ('sdfuiob sdoiuf hsdiob usdfhbiu ', 'artur.gurgul@gmail.com'),
  ('afiuovh adsiuob hsdaiubv hsdaiub', 'artur.gurgul@gmail.com'),
  ('adojfjh nbakisu vfhbiaudk vfhi', 'artur.gurgul@gmail.com');
