CREATE TABLE post
( 
   `post_id`      INT NOT NULL AUTO_INCREMENT,
   `title`        VARCHAR(255) NOT NULL,
   `subtitle`     VARCHAR(255) NOT NULL,
   `publish_date` VARCHAR(255) NOT NULL,
   `author`       VARCHAR(255) NOT NULL,
   `author_url`   VARCHAR(255) NOT NULL DEFAULT "", 
   `featured`     TINYINT(1) DEFAULT 0,
   `image_url`   VARCHAR(255) NOT NULL  DEFAULT "",
   `content`      TEXT NOT NULL,

   `heroImg1Path`   VARCHAR(255) NOT NULL,
   `heroImg2Path`   VARCHAR(255) NOT NULL,
   `authorImgPath`   VARCHAR(255) NOT NULL,
   PRIMARY KEY (`post_id`)
) ENGINE = InnoDB
CHARACTER SET = utf8mb4
COLLATE utf8mb4_unicode_ci
;

/*title VARCHAR(255)
subtitle VARCHAR(255)
author VARCHAR(255)
author_url VARCHAR(255)
publish_date VARCHAR(255)
image_url VARCHAR(255)
featured TINYINT(1) DEFAULT 0*/


CREATE TABLE user 
( 
   `user_id`      INT NOT NULL AUTO_INCREMENT,
   `email`        VARCHAR(255) NOT NULL,
   `password`     VARCHAR(255) NOT NULL,
   PRIMARY KEY (`user_id`)
) ENGINE = InnoDB
CHARACTER SET = utf8mb4
COLLATE utf8mb4_unicode_ci
;

INSERT INTO user (email, password)
VALUES ('adminTest@test.com', 'randomPassword');