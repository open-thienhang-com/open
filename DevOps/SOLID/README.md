# SOLID: 5 nguyên lý trong thiết kế hướng đối tượng

![](https://res.cloudinary.com/practicaldev/image/fetch/s--sWoCO4Jb--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/925xo3xgmvhw8l1xnt6t.png)

## SOLID ra đời như thế nào?
**Lập trình hướng đối tượng (object oriented programming – OOP**) là một trong những mô hình lập trình được sử dụng nhiều nhất. Các tính chất đặc biệt khiến việc hướng đối tượng trở nên hiệu quả đó là:
- **Tính trừu tượng (abstraction)**: Tạo ra các lớp trừu tượng mô hình hoá các đối tượng trong thế giới thực.
- **Tính đóng gói (Encapsulation)**: Các thực thể của lớp trừu tượng có các giá trị thuộc tính riêng biệt.
- **Tính kế thừa (Inheritance)**: Các đối tượng có thể dễ dàng kế thừa và mở rộng lẫn nhau.
- **Tính đa hình (Polymorphism)**: Có thể thực hiện một hành động đơn theo nhiều cách thức khác nhau tuỳ theo loại đối tượng cụ thể đang được gọi.

## SOLID là gì?
**SOLID** là viết tắt của 5 chữ cái đầu trong 5 nguyên tắc thiết kế hướng đối tượng. Giúp cho lập trình viên viết ra những đoạn code dễ đọc, dễ hiểu, dễ maintain. Nó được đưa ra bởi Robert C. Martin và Michael Feathers. 5 nguyên tắc đó bao gồm:
- **S**ingle responsibility priciple (SRP)
- **O**pen/Closed principle (OCP)
- **L**iskov substitution principe (LSP)
- **I**nterface segregation principle (ISP)
- **D**ependency inversion principle (DIP)

## Single responsibility priciple
**Nội dung:**
`A class should have only a single responsibility.`

Nguyên lý này có ý nghĩa là một class chỉ nên giữ một trách nhiệm duy nhất. Một class có quá nhiều chức năng sẽ trở nên cồng kềnh và trở nên khó đọc, khó maintain. Mà đối với ngành IT việc requirement thay đổi, cần thêm sửa chức năng là rất bình thường, nên việc code trong sáng, dễ đọc dễ hiểu là rất cần thiết.

**Ví dụ**
Hình dung rằng nhân viên của một công ty phần mềm cần phải làm 1 trong 3 việc sau đây: lập trình phần mềm (developer), kiểm tra phần mềm (tester), bán phần mềm (salesman). Mỗi nhân viên sẽ có một chức vụ và dựa vào chức vụ sẽ làm công việc tương ứng. Khi đó bạn có nên thiết kế lớp “Employee” với thuộc tính “position” và 3 phương thức developSoftware(), testSoftware() và saleSoftware() không?

Câu trả lời là KHÔNG. Thử hình dung nếu có thêm một chức vụ nữa là quản lí nhân sự, ta sẽ phải sửa lại lớp “Employee”, thêm phương thức mới vào sao? Nếu có thêm 10 chức vụ nữa thì sao? Khi đó các đối tượng được tạo ra sẽ dư thừa rất nhiều phương thức: Developer thì đâu cần dùng hàm testSoftware() và saleSoftware() đúng không nào, lỡ may dùng lầm phương thức cũng sẽ gây hậu quả khôn lường.

## Open/Closed principle
**Nội dung**
`Software entities (classes, modules, functions, etc.) should be open for extension, but closed for modification.`

Theo nguyên lý này, mỗi khi ta muốn thêm chức năng cho chương trình, chúng ta nên viết class mới mở rộng class cũ (bằng cách kế thừa hoặc sở hữu class cũ) chứ không nên sửa đổi class cũ. Việc này dẫn đến tình trạng phát sinh nhiều class, nhưng chúng ta sẽ không cần phải test lại các class cũ nữa, mà chỉ tập trung vào test các class mới, nơi chứa các chức năng mới.

**Ví dụ**
Ta có một module như sau:
```go
type AreaCalculator struct {
	shapes string
}

func (a *AreaCalculator) area(){
	if a.shapes == "Square"{
		/*Tính diện tích hình vuông*/
	} else if a.shapes == "Circle"{
		/*Tính diện tích hình tròn*/
	}
}
```
Dựa vào module trên ta sẽ xem xét kịch bản nếu như ta bổ sung thêm hình tam giác hay hình chữ nhật thì ta phải sửa bên trong module trên nhiều lần. Điều này sẽ dẫn tới vi phạm nguyên lý **Open/Closed principle**. Do đó giải pháp hay hơn chính là ta sẽ tạo một interface với phương thức **area()** và khi muốn tính diện tích của một hình mới ta chỉ việt gọi một struct mới và gọi phương thức **area()**:
```go
type Shape interface {
    area() float64
}

type Circle struct {
}

type Square struct {
}
func (c *Circle) area() float64 {
    /*Tính diện tích hình tròn*/
}

func (r *Square) area() float64 {
	/*Tính diện tích hình vuông*/
}
```


## Liskov substitution principle
**Nội dung**
`Objects in a program should be replaceable with instances of their subtypes without altering the correctness of that program`.

Nguyên lý này có thể hiểu là các đối tượng của class cha có thể được thay thế bởi các đối tượng của các class con mà không làm thay đổi tính đúng đắn của chương trình.


## Interface segregation principle
**Nội dung**
`Many client-specific interfaces are better than one general-purpose interface.`

Nguyên lý này rất dễ hiểu. Hãy tưởng tượng chúng ta có 1 interface lớn, khoảng 100 methods. Việc implements sẽ rất vất vả vì các class impliment interface này sẽ bắt buộc phải phải thực thi toàn bộ các method của interface. Ngoài ra còn có thể dư thừa vì 1 class không cần dùng hết 100 method. Khi tách interface ra thành nhiều interface nhỏ, gồm các method liên quan tới nhau, việc implement và quản lý sẽ dễ hơn.

## Dependency inversion principle
**Nội dung**
`Depend on abstractions, not on concretions.`

Có thể hiểu nguyên lí này như sau: những thành phần trong 1 chương trình chỉ nên phụ thuộc vào những cái trừu tượng (abstraction). Những thành phần trừu tượng không nên phụ thuộc vào các thành phần mang tính cụ thể mà nên ngược lại.

Những cái trừu tượng (abstraction) là những cái ít thay đổi và biến động, nó tập hợp những đặc tính chung nhất của những cái cụ thể. Những cái cụ thể dù khác nhau thế nào đi nữa đều tuân theo các quy tắc chung mà cái trừu tượng đã định ra. Việc phụ thuộc vào cái trừu tượng sẽ giúp chương trình linh động và thích ứng tốt với các sự thay đổi diễn ra liên tục.


## Tổng kết
**SOLID** là 5 nguyên tắc cơ bản trong việc thiết kế phần mềm. Nó giúp chúng ta tổ chức sắp xếp các function, method, class một cách chính xác hơn. Làm sao để kết nối các thành phần, module với nhau.

#### Rõ ràng, dễ hiểu
Teamwork là điều không thể tránh trong lập trình. Áp dụng SOLID vào công việc bạn sẽ tạo ra các hàm tốt, dễ hiểu hơn. Giúp cho bạn và đồng nghiệp đọc hiểu code của nhau tốt hơn.

#### Dễ thay đổi
**SOLID** giúp tạo ra các module, class rõ ràng, mạch lạc, mang tính độc lập cao. Do vậy khi có sự yêu cầu thay đổi, mở rộng từ khách hàng, ta cũng không tốn quá nhiều công sức để thực hiện việc thay đổi.

#### Tái sử dụng
**SOLID** khiến các lập trình viên suy nghĩ nhiều hơn về cách viết phần mềm, do vậy code viết ra sẽ mạch lạc, dễ hiểu, dễ sử dụng.
