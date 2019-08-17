CREATE TABLE TASKS (
    discord VARCHAR(20) NOT NULL,
    id INTEGER NOT NULL,
    name VARCHAR(100) NOT NULL,
    status INTEGER NOT NULL,
    data_create timestamp default current_timestamp,
    data_update timestamp,
    PRIMARY KEY(discord, id)
);