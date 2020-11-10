create table temperature.temperature
(
	ID int auto_increment
		primary key,
	Temp float null,
	Time timestamp null,
	Location char(255) null
);

