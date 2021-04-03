## MYSQL 기본
* 행 == 튜플 == 레코드
* 열 == 필드 == 속성
* degree: 열의 수
* cardinality: 행의 수

```
## mac brew mysql start & stop
$ brew services start mysql
$ brew services stop mysql
```

```
$ mysql -uroot -S ./tmp/mysql.sock

$ SHOW DATABASES;
$ CREATE DATEBASE mytest;
$ USE mytest;

$ SHOW tables;
$ CREATE TABLE car(color VARCHAR(20) PRIMARY KEY, mileage INT, speed INT);
$ CREATE TABLE book(id INT PRIMARY KEY AUTO_INCREMENT, author VARCHAR(100));

$ ALTER TABLE car ADD no INT DEFAULT 7;
$ ALTER TABLE car DROP no;

$ SELECT color FROM car WHERE color LIKE 'h%aa%';

$ SELECT color, speed FROM car where no IN (SELECT num FROM job LIMIT 5);

$ SELECT color, speed FROM car where no IN (7, 11) ORDER BY speed DESC;

## inner join
$ SELECT car.color, job.money FROM car JOIN job ON car.no = job.num ORDER BY color; # Ansi 표준
$ SELECT car.color, job.money FROM car, job WHERE car.no = job.num ORDER BY color;

## GROUP BY, HAVING 과 집계함수 COUNT
$ SELECT speed, count(no) FROM car GROUP BY speed;
$ SELECT speed, count(no) FROM car GROUP BY speed HAVING count(no) > 1;

$ SELECT color FROM car LIMIT 0,2;
$ SELECT color FROM car LIMIT 2,2;
$ SELECT color FROM car LIMIT 4,2;
```
 
### DML
 
#### SELECT
* ```SELECT * FROM 금지```
* FROM 절은 3depth 이상 들어가지 않도록 한다.
* ALIAS는 직관적으로 명시한다.
* SELECT 절에 있는 subquery는 최대한 자제한다.

#### SELECT 쿼리의 처리 순서
1. FROM : 해당 데이터가 있는곳에서
2. WHERE : 조건에 맞는 데이터를 찾고
3. GROUP BY : 원하는 데이터로 가공하여
4. SELECT : 필요한 컬럼만 뽑아내서
5. ORDER BY : 정렬하고
6. LIMIT  : 원하는 수만큼 추출
  
#### INSERT
```
INSERT INTO table_name VALUES (?,?,?);
INSERT INTO table_name(col1, col2) VALUES (?, ?);
INSERT INTO table_name SELECT.. FROM ..;
REPLACE INTO table_name VALUES(?,?,?,?); # override
INSERT INTO table_name ON DUPLICATE KEY UPDATE ...; # 중복시 PK만 냅두고 변경
```

* 대량의 INSERT 작업이 있는 batch job은 주기적으로 commit을 수행한다.
* INSERT INTO table SELECT .. FROM table 의 구문에서 SELECT 성능을 확인한다.
* REPLACE는 데이터를 삭제 후 INSERT
* INSERT ON DUPLICATE KEY UPDATE는 원하는 컬럼의 데이터를 수정

#### UPDATE
```
UPDATE table_name SET col1=? WHERE col2=?;
UPDATE table_name SET col1=? WHERE col2=? ORDER BY .. LIMIT 10;
UPDATE table1, table2 SET table1.a=? WHERE table1.b=table2.b;
```

* WHERE절이 있는지 확인한다.
* UPDATE 대상이 맞는지 다시 확인한다. (SELECT로 한번 더)
* 불안하면 start transaction; update ..... ; commit; or rollback;
* 대량 작업은 DBA와 함께 진행한다.
* UPDATE, DELETE시 반드시 WHERE 절에는 PK만 작성하도록 한다. (인덱스를 조건으로 걸어도 LOCK 경합이 발생할 수 있음)

#### DELETE
```
DELETE FROM tablename WHERE col1=?;
DELETE FROM tablename WHERE col2=? ORDER BY... LIMIT 10;
```

* UPDATE와 동일

* TRUNCATE vs DELETE FROM
  * TRUNCATE TABLE: drop + create 와 동일(디스크 공간 반납), auto_increment 값 초기화
  * DELETE FROM TABLE: 실제 디스크 공간 반납X, auto_increment 값 유지


#### JOIN
 JOIN은 다른 테이블의 열을 가져와서 "열을 늘리는" 집합 연산이다. (행은 UNION)

* 내부조인(Inner Join)
* 외부조인(Outer Join)
* 동등, 비동등, 자연, 크로스, 카티션, 셀프 조인 등 있음

```
JOIN 작성법: 아래 쿼리들 방식의 성능은 동일

## ASNSI 표준 <- 권장
SELECT f.title, f.id
FROM film as f
JOIN film_actor ON f.id = film_actor.id
WHERE f.title = ?

## where 절에 표시
SELECT f.title, f.id
FROM film as f, film_actor
WHERE f.id = film_actor.id
AND f.title = ?
```

##### Nested Loop Join (NL Join)
* mysql은 8.0 이하 모든 버전에선 NL Join을 사용한다.
* 순차적으로 처리된다.
* 먼저 드라이브 되는 선행 테이블의 크기가 적으면 유리하다.(이중 포문으로 생각하자)
* 인덱스가 모두 사용되는건 아니다.
* 쿼리 검수 잘 받자..

##### Inner Join
 두 개 이상의 테이블을 내부 결합
 
 ```
 mysql> select * from book;
+----+--------+
| id | author |
+----+--------+
|  1 | tuyy1  |
|  2 | tuyy2  |
|  3 | tuyy3  |
|  4 | tuyy4  |
|  5 | tuyy5  |
|  6 | tuyy6  |
|  7 | tuyy7  |
|  8 | tuyy8  |
+----+--------+
8 rows in set (0.00 sec)

mysql> select * from color;
+----+--------+------+
| id | name   | no   |
+----+--------+------+
|  1 | red001 |    1 |
|  2 | red002 |    1 |
|  3 | red003 |    1 |
|  4 | red004 |    2 |
|  5 | red005 |    2 |
|  6 | red006 |    3 |
|  7 | red007 |    3 |
+----+--------+------+
7 rows in set (0.00 sec)

mysql> select book.id, book.author, color.id as color_id, color.name, color.no from book join color on book.id = color.no;
+----+--------+----------+--------+------+
| id | author | color_id | name   | no   |
+----+--------+----------+--------+------+
|  1 | tuyy1  |        1 | red001 |    1 |
|  1 | tuyy1  |        2 | red002 |    1 |
|  1 | tuyy1  |        3 | red003 |    1 |
|  2 | tuyy2  |        4 | red004 |    2 |
|  2 | tuyy2  |        5 | red005 |    2 |
|  3 | tuyy3  |        6 | red006 |    3 |
|  3 | tuyy3  |        7 | red007 |    3 |
+----+--------+----------+--------+------+
 ```
 
 

##### Outer Join
 두 개 이상의 테이블을 외부 결합, 결과가 없는 경우 null을 반환, left/right join 존재
 
 ```
 mysql> select book.id, book.author, color.id as color_id, color.name, color.no from book left outer join color on book.id = color.no;
+----+--------+----------+--------+------+
| id | author | color_id | name   | no   |
+----+--------+----------+--------+------+
|  1 | tuyy1  |        3 | red003 |    1 |
|  1 | tuyy1  |        2 | red002 |    1 |
|  1 | tuyy1  |        1 | red001 |    1 |
|  2 | tuyy2  |        5 | red005 |    2 |
|  2 | tuyy2  |        4 | red004 |    2 |
|  3 | tuyy3  |        7 | red007 |    3 |
|  3 | tuyy3  |        6 | red006 |    3 |
|  4 | tuyy4  |     NULL | NULL   | NULL |
|  5 | tuyy5  |     NULL | NULL   | NULL |
|  6 | tuyy6  |     NULL | NULL   | NULL |
|  7 | tuyy7  |     NULL | NULL   | NULL |
|  8 | tuyy8  |     NULL | NULL   | NULL |
+----+--------+----------+--------+------+

mysql> select book.id, book.author, color.id as color_id, color.name, color.no from book right outer join color on book.id = color.no;
+------+--------+----------+--------+------+
| id   | author | color_id | name   | no   |
+------+--------+----------+--------+------+
|    1 | tuyy1  |        1 | red001 |    1 |
|    1 | tuyy1  |        2 | red002 |    1 |
|    1 | tuyy1  |        3 | red003 |    1 |
|    2 | tuyy2  |        4 | red004 |    2 |
|    2 | tuyy2  |        5 | red005 |    2 |
|    3 | tuyy3  |        6 | red006 |    3 |
|    3 | tuyy3  |        7 | red007 |    3 |
+------+--------+----------+--------+------+
 ```

##### Cross Join
 두 개 테이블을 가지고 가능한 모든 조합을 만든다. (잘 안씀)

```
mysql> select book.id, book.author, color.id as color_id, color.name, color.no from book cross join color on book.id = color.no;
+----+--------+----------+--------+------+
| id | author | color_id | name   | no   |
+----+--------+----------+--------+------+
|  1 | tuyy1  |        1 | red001 |    1 |
|  1 | tuyy1  |        2 | red002 |    1 |
|  1 | tuyy1  |        3 | red003 |    1 |
|  2 | tuyy2  |        4 | red004 |    2 |
|  2 | tuyy2  |        5 | red005 |    2 |
|  3 | tuyy3  |        6 | red006 |    3 |
|  3 | tuyy3  |        7 | red007 |    3 |
+----+--------+----------+--------+------+
```

##### Equi Join
 둘은 비교 대상이 아니며 구분 기준이 전혀 다르다.

```
## inner join
$ SELECT TS.store_id FROM Goods AS TS INNER JOIN MyGoods AS S ON TS.goods_id > S.goods_id;

## equi join
$ SELECT TS.store_id FROM Goods AS TS LEFT OUTER JOIN MyGoods AS S ON TS.goods_id = S.goods_id;
```

#### HINT
 쿼리 튜닝의 한 방식, 올바른 실행 계획을 세우도록 명시적으로 옵션을 지정하는 방법이다. 그러나 잘못 쓰면 문제가 있을 수 있음
 
* STRAIGHT_JOIN : 조인 순서 보장하는 힌트
* USE INDEX (PRIMARY) : 인덱스 사용을 보장하는 힌트

#### UNION & UNION ALL
 다른 테이블의 행을 가지고 와서 "행을 늘리는" 집합 연산이다. 두 개 이상의 조회 결과를 합칠 수 있다. (단, 컬럼이 타입까지 모두 동일해야함)
 
```
# 중복을 제거
SELECT .. FROM table UNION SELECT ... FROM table2;
 
# 중복을 허용 (SUM, MAX, AVG 연산시 에러 발생하므로 주의)
SELECT .. FROM table UNION ALL SELECT ... FROM table2;
```

#### INDEX
MYSQL의 인덱스는 B-Tree 구조이다. 튜닝의 중요하고 쉬운? 부분이다.

* 인덱스를 건 컬럼 순서가 중요하다.

##### 인덱스를 타지 못하는 경우
```
# LIKE '% %'
SELECT * FROM test WHERE keyword LIKE '%가%';

# 좌변이 가공
SELECT * FROM emp WHERE SUBSTR(emp_nm,1,1) = 'M';

# 부정연산자를 사용하는 경우
SELECT * FROM emp WHERE emp_nm != '정인구';

# 암시적인 형변환 (emp_no는 int형인데 조건은 스트링.. 암시적으로 형 변환됨)
SELECT * FROM emp WHERE emp_no = '11617';

# NULL, NOT NULL (mysql은 is null은 인덱스 활용 가능)
SELECT * FROM emp WHERE emp_no IS NOT NULL;
```
