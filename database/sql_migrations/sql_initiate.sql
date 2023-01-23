-- +migrate Up
-- +migrate StatementBegin

    
CREATE TABLE Account (
    id SERIAL PRIMARY KEY,
    Username VARCHAR(255) NOT NULL,
    Password VARCHAR(500) NOT NULL,
    CustomerID BIGINT NOT NULL
);
CREATE TABLE customer (
    CustomerID BIGINT PRIMARY KEY,
    Email VARCHAR(50) NOT NULL,
    MobileNumber VARCHAR(255) NOT NULL
);

CREATE TABLE film (
    FilmID SERIAL PRIMARY KEY,
    FilmName VARCHAR(255) NOT NULL,
    FilmStatus VARCHAR(50) NOT NULL,
    FilmCode VARCHAR(255) NOT NULL
);

CREATE TABLE schedule (
    ScheduleID SERIAL PRIMARY KEY,
    FilmCode VARCHAR(255) NOT NULL,
    ScheduleDate DATE NOT NULL,
    StartTime TIME NOT NULL,
    EndTime TIME NOT NULL,
    Price VARCHAR(50) NOT NULL,
    Duration BIGINT NOT NULL,
    StudioID BIGINT NOT NULL
);

CREATE TABLE studio (
    StudioID SERIAL PRIMARY KEY,
    StudioName VARCHAR(255) NOT NULL
);

CREATE TABLE seat (
    SeatID SERIAL PRIMARY KEY,
    NumberSeat VARCHAR(50) NOT NULL
);

CREATE TABLE ticket (
    TicketID SERIAL PRIMARY KEY,
    CustomerID BIGINT NOT NULL,
    ScheduleID BIGINT NOT NULL,
    SeatID BIGINT NOT NULL,
    StudioID BIGINT NOT NULL,
    FilmCode VARCHAR(255) NOT NULL
);

-- +migrate StatementEnd