# OOP - Lập trình hướng đối tượng (Object-Oriented Programming)

## 1. Lợi ích:

    - Cách tiếp cận rất thực tế. Dễ đưa ra ví dụ từ đời sống
    - Dễ bảo trì, mở rộng code

## 2. Class

    - Trong OOP có **_Class_** như 1 khuôn mẫu.
    - Class sẽ có các **_thuộc tính(properties)_** và các **_phương thức (methods)_**

## 3. Tính đóng gói (Encapsulation)

    - public: có thể truy cập từ bất kỳ đâu
    - private: chỉ có thể được truy cập ở bên trong class
    - protected: chỉ có thể được truy cập ở bên trong class hoặc các class kế thừa từ class đó

```
Ví dụ: A Dee có bồ và a Dee giấu nên chỉ có anh Dee biết và MT không biết, Chang không biết và mọi người cũng không (ý nghĩa bồ a Dee là 1 một thuộc tính private, tên a Dee là thuộc tính public)
```

## 4. Tính kế thừa (Inheritance)

```
Ví du: MT kế thừa từ my father làn da đen, nhưng MT lại có giới tánh nữ mà cha lại là giới tánh nam nên ở đây (gọi là override - tái định nghĩa) và MT thích nghe rap cha MT thì không (gọi là con hơn cha nhà có phúc, đùa thôi điều này nghĩa là ngoài những cái kế thừa từ class cha, class con cũng có thể thêm thuộc tính hoặc phương thức riêng)
```

## 5. Tính đa hình (Polymorphism)

    - overloading đa hình khi biên dịch(compile time) : Trong 1 class các phương thức (methods) có cùng tên nhưng kiểu trả về và tham số truyền vào khác nhau (số lượng, kiểu)
    - overriding(đa hình ở thời điểm thực thi(runtime)): Các phương thức được thực hiện ở các lớp con kế thừa từ lớp cha (base class). Nội dung thực hiện bên trong mỗi lớp khác nhau tùy vào logic nghiệp vụ. Chỉ khi nào runtime ta mới biết được đối tượng sẽ sử dụng phương thức nào

```
Ví dụ: Lớp con chó và con mèo kế thừa từ lớp động vật nhưng tiếng kêu của nó sẽ khác nhau
```

## 6. Tính trừu tương (Abstraction)

    - Không cụ thể, tổng quan
    - Tính chất này được thể hiện qua việc sử dụng interface hoặc abstract class.

```
Ví dụ: Máy xay sinh tố, ta không cần biết lúc xay thì máy móc hoạt động thế nào ta chỉ cần biết bấm nút xay là máy xay
```
