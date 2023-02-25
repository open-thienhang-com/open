Trong bài viết này mình sẽ tập trung chủ yếu vào database PostgresSQL.

# Queries

## Công cụ

<https://dbeaver.io/download/>

## Dữ liệu mẫu

Chạy file .sql trong thư mục `example01/demo`

## Các lệnh chính

| Câu lệnh  | Giá trị |
| ------------- | ------------- |
| FROM  | Source Data  |
| WHERE  | Row Filter  |
| GROUP BY  | Grouping  |
| SELECT  | Return Expressions  |
| ORDER BY  | Presentation Order  |
| OFFSET FETCH  | Paging  |

### Single Data Source queries

---
Chọn các records từ một nguồn dữ liệu duy nhất `<Data Source>` có thể là: table, view, function hoặc một subquery ta dùng lệnh sau:

```sh
SELECT X FROM  <Data Source>
```

Ví dụ:

```sh
> SELECT * FROM Staff;
```

Kết quả trả về gồm 9 records có trong bảng staff

| ID  | Email | Hire_Date |
| ------------- | ------------- | ------------- |
| 1  | ashley.flores@animalshelter.com  |2016-01-01  |
| 2  | dennis.hill@animalshelter.com  |2018-10-07  |
| ...  | ...  |...  |
| 9  | wayne.carter@animalshelter.com  |2018-04-02  |

```sh
> SELECT 'thienhang' as Fact
 FROM Staff;
```

Kết quả trả về gồm 9 records chỉ với giá trị thienhang tương ứng của cột Fact

| ID  | Fact |
| ------------- | ------------- |
| 1  | thienhang  |
| 2  | thienhang  |
| ...  | ...  |...  |
| 9  | thienhang  |
Bảng 1.2

```sh
> SELECT *, 'thienhang' as Fact
 FROM Staff;
```

Kết quả trả về gồm 9 records bao bọc lại bảng 1.1

| ID  | Email | Hire_Date | Fact |
| ------------- | ------------- | ------------- | ------------- |
| 1  | ashley.flores@animalshelter.com  |2016-01-01  |thienhang  |
| 2  | dennis.hill@animalshelter.com  |2018-10-07  |thienhang  |
| ...  | ...  |...  |
| 9  | wayne.carter@animalshelter.com  |2018-04-02  |thienhang  |

### Dual source query processing

Khi có nhiều hơn một nguồn dữ liệu, ví dụ cần kết hợp hai bảng lại để lấy thông tin chi tiết của một nhân viên

STEP 1: Cartesian Product

```sh
A CROSS JOIN B
```

| Table A |
| :---:   |
| 1 |
| 2 |

| Table B |
| :---:   |
| 3 |
| 4 |

| Table A | Table B    |
| :---:   | :---: |
| 1 | 3   |
| 1 | 4   |
| 2 | 3   |
| 2 | 4   |

STEP 2: Qualification

```sh
A INNER JOIN B
ON
'A.VALUE = B.VALUE'
```
