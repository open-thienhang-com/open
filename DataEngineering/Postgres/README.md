## PostgreSQL là lựa chọn tốt cho Data Engineering?

SQL là yếu tố cần thiết để thành công trong bất kỳ công việc nào tập trung vào dữ liệu, đặc biệt là kỹ thuật dữ liệu. Những lý do PostgreSQL là lựa chọn tốt nhất như:

- Truy vấn song song: PostgreSQL giúp bạn có thể chạy các truy vấn song song. Đây là khi sức mạnh của CPU được tận dụng để cho phép chạy nhiều truy vấn cùng một lúc. Điều này đặc biệt quan trọng trong khoa học dữ liệu, nơi thường có một truy vấn chung.
- Hỗ trợ cú pháp SQL đầy đủ: PostgreSQL hỗ trợ nhiều cú pháp SQL và nhấn mạnh vào việc tuân thủ tiêu chuẩn SQL. Do đó, nó hỗ trợ các hàm cửa sổ, kế thừa bảng và các biểu thức chung.
- Hỗ trợ dữ liệu mở rộng: PostgreSQL hỗ trợ cấu trúc dữ liệu NoSQL như JSON và XML.
- Phân vùng khai báo: Đây là khi các bảng được chia thành các phân đoạn khác nhau gọi là phân vùng (partition). Ví dụ, bạn có thể tạo một phân vùng khác nhau cho từng mã vùng cho các tập dữ liệu lớn, được phân phối theo địa lý.

## Execution plan

Execution plan giống như một bản danh sách những việc phải làm để hoàn thành một câu lệnh truy vấn. Mỗi câu truy vấn trong PostgreSQL đều có một execution plan. Chúng ta hoàn toàn không thể truy cập vào quá trình chi tiết của một execution plan. Một plan có thể hiệu quả hoặc không tùy thuộc vào rất nhiêu yếu tố như số lượng record, điều kiện join, vv.

PostgresSQL hỗ trợ câu lệnh Explain để chúng ta xem một quá trình thực thi một cách chi tiết:

- Một câu truy vấn như sau:

```sh
SELECT last_name FROM employees where salary >= 50000;
```

- Chúng ta có thể xem chi tiết quá trình xử lí của câu truy vấn như sau:

```sh
EXPLAIN SELECT last_name FROM employees where salary >= 50000;
                      QUERY PLAN
Seq Scan on employees  (cost=0.00..16.50 rows=173 width=118)
   Filter: (salary >= 50000)
```

We can both execute the query and inspect the path/time it took with:

```sh
EXPLAIN ANALYZE SELECT last_name FROM employees where salary >= 50000
                                        QUERY PLAN
--------------------------------------------------------------------------------------------------------

 Seq Scan on employees  (cost=0.00..16.50 rows=173 width=118) (actual time=0.018..0.018 rows=0 loops=1)
   Filter: (salary >= 50000)
 Total runtime: 0.053 ms
```

Nguồn tham khảo:
<https://www.postgresguide.com/performance/explain/>
