CREATE TABLE movies.movie
(
    id           varchar(100) NOT NULL,
    title        varchar(100) NOT NULL,
    `cast`       varchar(250) NULL,
    release_date DATE         NULL,
    genre        varchar(100) NULL,
    director     varchar(100) NULL,
    CONSTRAINT movie_PK PRIMARY KEY (id)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_0900_ai_ci;


CREATE TABLE movies.`user`
(
    id       varchar(100) NOT NULL,
    username varchar(100) NOT NULL,
    password varchar(100) NOT NULL,
    CONSTRAINT user_PK PRIMARY KEY (id),
    CONSTRAINT user_UN UNIQUE KEY (username)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_0900_ai_ci;


CREATE TABLE movies.wish_list
(
    user_id  varchar(100) NOT NULL,
    movie_id varchar(100) NOT NULL,
    comment  varchar(300) NULL,
    CONSTRAINT wish_list_PK PRIMARY KEY (user_id, movie_id),
    CONSTRAINT wish_list_FK FOREIGN KEY (user_id) REFERENCES movies.`user` (id) ON DELETE CASCADE,
    CONSTRAINT wish_list_FK_1 FOREIGN KEY (movie_id) REFERENCES movies.movie (id)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_0900_ai_ci;
