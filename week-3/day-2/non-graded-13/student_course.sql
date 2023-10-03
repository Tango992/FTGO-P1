-- DDL
CREATE TABLE IF NOT EXISTS major (
    major_id INT AUTO_INCREMENT NOT NULL,
    major_name VARCHAR(100) NOT NULL,
    PRIMARY KEY (major_id)
);

CREATE TABLE IF NOT EXISTS students (
    student_id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    major_id INT NOT NULL,
    year_of_study INT NOT NULL,
    PRIMARY KEY (student_id),
    FOREIGN KEY (major_id) REFERENCES major(major_id)
);

CREATE TABLE IF NOT EXISTS instructor (
    instructor_id INT AUTO_INCREMENT NOT NULL,
    instructor_name VARCHAR(100) NOT NULL,
    PRIMARY KEY (instructor_id)
);

CREATE TABLE IF NOT EXISTS courses (
    course_id INT AUTO_INCREMENT NOT NULL,
    title VARCHAR(100) NOT NULL,
    instructor_id INT NOT NULL,
    schedule DATETIME NOT NULL,
    credit_hours INT NOT NULL,
    PRIMARY KEY (course_id),
    FOREIGN KEY (instructor_id) REFERENCES instructor(instructor_id)
);

CREATE TABLE IF NOT EXISTS student_courses (
    student_id INT NOT NULL,
    course_id INT NOT NULL,
    preferred_schedule DATETIME NOT NULL,
    FOREIGN KEY (student_id) REFERENCES students(student_id),
    FOREIGN KEY (course_id) REFERENCES courses(course_id)
);

-- DML
INSERT INTO instructor (instructor_name) 
VALUES ('Budi'), ('John'), ('Jane'), ('Bar'), ('Baz');

INSERT INTO major (major_name)
VALUES ('Communications'), ('Computer Science'), ('History');

INSERT INTO students (name, email, major_id, year_of_study) 
VALUES
    ('Ana', 'ana@mail.com', 1, 2020),
    ('Ina', 'ina@mail.com', 1, 2021),
    ('Jack', 'jack@mail.com', 2, 2020),
    ('Elon', 'elon@mail.com', 3, 2022),
    ('Mark', 'mark@mail.com', 2, 2020);

INSERT INTO courses (title, instructor_id, schedule, credit_hours)
VALUES
    ('Introduction to communications', 1, '2023-10-01 10:00:00', 4),
    ('Introduction to computer', 2, '2023-10-02 08:00:00', 3),
    ('Introduction to history', 3, '2023-09-29 14:00:00', 4),
    ('Communication skills', 4, '2023-09-28 09:00:00', 3),
    ('Nationality', 5, '2023-09-26 16:00:00', 2);

INSERT INTO student_courses (student_id, course_id, preferred_schedule)
VALUES 
    (1, 1, '2023-10-01 10:00:00'),
    (1, 5, '2023-10-01 12:00:00'),
    (1, 4, '2023-09-29 15:00:00'),
    (2, 1, '2023-09-28 10:00:00'),
    (2, 5, '2023-09-26 16:00:00'),
    (3, 2, '2023-10-02 08:00:00'),
    (3, 5, '2023-10-02 15:00:00'),
    (4, 3, '2023-09-29 14:00:00'),
    (4, 5, '2023-09-26 15:00:00'),
    (5, 2, '2023-10-02 08:00:00'),
    (5, 4, '2023-09-28 09:00:00');

-- Retrieve the list of all students enrolled in a specific course. 
SELECT students.name, courses.title 
FROM student_courses
JOIN students ON students.student_id = student_courses.student_id
JOIN courses ON courses.course_id = student_courses.course_id
WHERE student_courses.course_id = 1;

-- Find all the courses a particular student has registered for. 
SELECT students.name, courses.title 
FROM student_courses
JOIN students ON students.student_id = student_courses.student_id
JOIN courses ON courses.course_id = student_courses.course_id
WHERE student_courses.student_id = 1;

-- Get the preferred schedule of a student for a specific course.
SELECT students.name, courses.title, student_courses.preferred_schedule 
FROM student_courses
JOIN students ON students.student_id = student_courses.student_id
JOIN courses ON courses.course_id = student_courses.course_id
WHERE student_courses.student_id = 5 AND student_courses.course_id = 2;