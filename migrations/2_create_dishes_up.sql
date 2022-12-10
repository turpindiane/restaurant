create table public.dishes(
    id serial not null primary key,
    title varchar(255) not null,
    price decimal not null,
    dish_course course
);