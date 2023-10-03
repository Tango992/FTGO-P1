-- DDL
CREATE TABLE IF NOT EXISTS Student (
    student_id INT AUTO_INCREMENT NOT NULL,
    student_name VARCHAR(100) NOT NULL,
    PRIMARY KEY (student_id)
);

CREATE TABLE IF NOT EXISTS Professor (
    professor_id INT AUTO_INCREMENT NOT NULL,
    professor_name VARCHAR(100) NOT NULL,
    PRIMARY KEY (professor_id)
);

CREATE TABLE IF NOT EXISTS Course (
    course_id INT AUTO_INCREMENT NOT NULL,
    course_title VARCHAR(200) NOT NULL,
    max_capacity INT NOT NULL,
    PRIMARY KEY (course_id)
);

CREATE TABLE IF NOT EXISTS Enrollment (
    enrollment_id INT AUTO_INCREMENT NOT NULL,
    student_id INT NOT NULL,
    course_id INT NOT NULL,
    enrollment_date DATE NOT NULL,
    PRIMARY KEY (enrollment_id),
    FOREIGN KEY (student_id) REFERENCES Student(student_id),
    FOREIGN KEY (course_id) REFERENCES Course(course_id)
);

CREATE TABLE IF NOT EXISTS TeachingAssignment (
    assignment_id INT AUTO_INCREMENT NOT NULL,
    professor_id INT NOT NULL,
    course_id INT NOT NULL,
    start_date DATE NOT NULL,
    PRIMARY KEY (assignment_id),
    FOREIGN KEY (professor_id) REFERENCES Professor(professor_id),
    FOREIGN KEY (course_id) REFERENCES Course(course_id)
);

-- DML
INSERT INTO Student (student_name)
VALUES ('Jack'), ('Elon'), ('Mark');

INSERT INTO Professor (professor_name)
VALUES ('Budi'), ('John'), ('Jane'), ('Bar'), ('Baz');

INSERT INTO Course (course_title, max_capacity)
VALUES 
    ('Introduction to communications', 30),
    ('Introduction to computer', 30),
    ('Introduction to history', 40),
    ('Communication skills', 35),
    ('Nationality', 40),
    ('Lorem ipsum', 30);

INSERT INTO TeachingAssignment (professor_id, course_id, start_date)
VALUES
    (1, 5, '2023-09-09'),
    (1, 1, '2023-09-10'),
    (2, 1, '2023-09-15'),
    (2, 2, '2023-09-11'),
    (3, 3, '2023-09-12'),
    (4, 4, '2023-09-13'),
    (5, 5, '2023-09-14');

INSERT INTO Enrollment (student_id, course_id, enrollment_date)
VALUES
    (1, 1, '2023-09-10'),
    (1, 2, '2023-09-10'),
    (2, 3, '2023-09-10'),
    (2, 4, '2023-09-10'),
    (3, 5, '2023-09-10'),
    (3, 1, '2023-09-10');

-- List of students enrolled in a specific course
SELECT Student.student_name, Course.course_title
FROM Enrollment
JOIN Student ON Student.student_id = Enrollment.student_id
JOIN Course ON Course.course_id = Enrollment.course_id
WHERE Enrollment.course_id = 1;

-- List of courses taught by a specific professor
SELECT Professor.professor_name, Course.course_title
FROM TeachingAssignment
JOIN Professor ON Professor.professor_id = TeachingAssignment.professor_id
JOIN Course ON Course.course_id = TeachingAssignment.course_id
WHERE TeachingAssignment.professor_id = 1;

-- List of professor teaching a specific course
SELECT Professor.professor_name, Course.course_title
FROM TeachingAssignment
JOIN Professor ON Professor.professor_id = TeachingAssignment.professor_id
JOIN Course ON Course.course_id = TeachingAssignment.course_id
WHERE TeachingAssignment.course_id = 1;