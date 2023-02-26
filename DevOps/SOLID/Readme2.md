# S.O.L.I.D - NGUYÊN TẮC THIẾT KẾ HƯỚNG ĐỐI TƯỢNG (OOD)
## 1. SOLID là gì? Sứ mệnh của SOLID ?
- **OOD** (Object-Oriented Design), trong lĩnh vực phát triển phần mềm OOD đóng vai trò quan trọng giúp chúng ta viết code một cách linh hoạt, mở rộng dễ dàng, có thể giúp rút ngắn thời gian bảo trì và dễ dàng tái sử dụng. 
- Nguyên tắc *SOLID* được đưa ra bởi Robert Cecil Martin (ông được gọi là Uncle Bod - ông đã phát triển nhiều nguyên tắc thiết kế phần mềm và là người sáng lập ra Tuyên ngôn Agile, với cuốn sách best-seller Clean Code ) trong bài báo **"Design Principles and Design Patterns"** vào năm 2000 của mình. Những khái niệm này sau đó được xây dựng bởi Michael Feathers, người giới thiệu cho chúng ta từ viết tắt SOLID. Và trong 22 năm qua, nguyên tắc này đã cách mạng hóa thế giới lập trình hướng đối tượng, thay đổi cách chúng ta viết phần mềm. 
+ **SOLID** là từ viết tắt của năm nguyên tắc:
+ **S**ingle Responsibility Principle <i>SRP</i>
+ **O**pen/Closed Principle <i>O/CP</i>
+ **L**iskov’s Substitution Principle <i>LSP</i>
+ **I**nterface Segregation Principle <i>ISP</i>
+ **D**ependency Inversion Principle <i>DIP</i> <br>
*=> SOLID là Nguyên tắc thiết kế lập trình giúp giảm sự phụ thuộc lẫn nhau. Kết hợp chặt chẽ có nghĩa là một nhóm các lớp phụ thuộc nhiều vào nhau mà bạn nên tránh trong code của mình. Khi các lớp ít phụ thuộc lẫn nhau sẽ giảm thiểu những thay đổi trong code của bạn, code sẽ dễ sử dụng hơn, có thể bảo trì, linh hoạt và ổn định.*<br>
## 2. 5 nguyên tắc trong OOD - S.O.L.I.D
**Single Responsibility Principle (SRP)** - Nguyên tắc trách nhiệm duy nhất <br>
    `A class should only have one responsibility. Furthermore, it should only have one reason to change` <br>
Nguyên lý này nói rằng một class chỉ nên có một trách nhiệm, hơn nữa chỉ nên có một lý do để thay đổi.<br>
Một vài lợi ích của nó:
1. Testing: một class có ít trách nhiệm sẽ có ít testcase
2. Lower coupling: ít phụ thuộc hơn
3.  Organization: tổ chức tốt hơn dễ tìm kiếm
+ Ví dụ thực tế trước nhé! <br>
Đường một chiều chỉ làm một việc là đi về một hướng, nhanh gọn an toàn hơn đường 2 chiều, nguy cơ rủi ro và đầy nguy hiểm. :V 
+ Ví dụ với code: <br>
Chúng ta có một class Book đơn giản:

```
    public class Book {

    private String name;
    private String author;
    private String desc;

    //constructor, getters and setters
}
```
Trong đoạn code này, mình có muột class Book lưu 3 thuộc tính là tên, tác giả và mô tả. <br>
Bây giờ thêm một vài phương thức để truy vấn desc
```
    public class Book {

    private String name;
    private String author;
    private String desc;

    //constructor, getters and setters

    // phương thức liên quan đến thuộc tính của Book class
    public String replaceWordInDesc(String word, String repText){
        return desc.replaceAll(word, repText);
    }

    public boolean isWordInDesc(String word){
        return text.contains(word);
    }
}
```
Nhưng ở đây chúng ta cần in ra thông tin sách và đọc, nếu mình thêm phương thức vào trong lớp Book
```
    public class Book {
    //...

    void printDescToConsole(){
        // thực hiện in desc 
    }
}
```
Ở đoạn code trên đã vi phạm nguyên tắc SRP, để khắc phúc thì tạo riêng biệt một class làm nhiệm vụ là in 
```
    public class BookPrinter {

    //in văn bản
    void printDescToConsole(String text){
        //định dạng và in văn bản ra 
    }

    void printTextToAnotherMedium(String text){
        // in với mục đích khác
    }
}

```
Class này giúp giảm bớt nhiệm vụ in ấn của lớp Book, mà ta còn có thể tận dụng lớp BookPrinter để gửi văn bản tới bên thứ 3, có thể là email, nhật ký, ... <br>
**Open/Closed Principle** - Nguyên tắc Mở/Đóng <br>
    ` Classes should be open for extension but closed for modification.`
Các thực thể phần mềm (class, module, func,...) phải mở để mở rộng, nhưng đóng để sửa đổi có nghĩa là bạn có thể mở rộng hành vi của class mà không cần sửa đổi nó. <br>
Ta có class Student
```
    class Student{
    private string name;
    private int age;
    private int student_type;
    int payTuitionFee()
    {
        if( this->student_type == 'foreign' )
            return STANDARD_FEE * 1.3;
        else if ( this->student_type == 'talented' )
            return STANDARD_FEE * 0.8;
        else
            return STANDARD_FEE;
    }
    
    // other functions
    // ...
    }

```
Class này hoàn toàn hoạt động đúng, nhưng khi chúng ta có thêm kiểu student (sinh viên cao học, sinh viên tại chức,...) thì phải vào sửa lại phương thức để đáp ứng hay sao? Ứng dụng nguyên tắc thứ 2!<br>
Thiết kế lại class Student:<br>
```
    define ( STANDARD_FEE, 1000 );
 
    class Student
    {
    private string name;
    private int age;
    
    int payTuitionFee()
    {
        return STANDARD_FEE; 
    }
    
    // other functions
    // ...
    }
```
Mở rộng các lớp student con bằng cách kế thừa:
```
    class AdvancedStudent extends Student
    {
    int payTuitionFee()
    {
        return STANDARD_FEE * 0.8;
    }
    
    // other functions
    // ...
    }
```
Lớp Student thỏa mãn tính đóng với mọi sự thay đổi bên trong, nhưng luôn mở cho sự kế thừa để mở rộng sau này! Perfect! <br>
[Tham khảo để hiểu rõ thêm ở đây -> https://taivublog.com/2017/03/25/solid-la-gi-nguyen-tac-2-dong-va-mo-open-closed-principle-ocp/]
Đi tiếp nào! <br>
**Liskov’s Substitution Principle (LSP)** - Nguyên tắc thay thế Liskov <br>
Nguyên tắc thay thế Liskov là nguyên tắc khó hiểu nhất trong những khái niệm lập trình, được đưa ra bởi Barbara Liskov vào năm 1987, cho rằng: <br>
    `If class A is a subtype of class B, we should be able to replace B with A without disrupting the behavior of our program` <br>
Các lớp con phải được thay thế cho các lớp cơ sở hoặc lớp cha của chúng. Nguyên tắc này đảm bảo rằng bất kỳ lớp nào là con của lớp cha đều có thể sử dụng được thay cho lớp cha của nó mà không có bất kỳ khó khăn nào. <br>
Khó hiểu quá, chúng ta đi phân tích code nè:<br>
Chúng ta có một interface Car với 2 phương thức khởi động và tăng tốc.
```
    public interface Car {

        void turnOnEngine();
        void accelerate();

    }
```
Thực hiện kế thừa interface thêm một vài phương thức mới:
```
    public class MotorCar implements Car {

    private Engine engine;

    //Constructors, getters + setters

    public void turnOnEngine() {
        //turn on the engine!
        engine.on();
    }

    public void accelerate() {
        //move forward!
        engine.powerOn(1000);
    }
}
```
Như đoạn code ở trên chúng ta có lớp con kế thừa được hết phương thức của lớp cha. Khoan! Dừng lại khoảng chừng là 2s, chúng ta có những chiếc xe điện thì sao! Điều này đã làm vi phạm nguyên tắc này! Để khắc phục thì ta phải làm lại mô hình có tính đến trạng thái không có động cơ. <br>
**Interface Segregation Principle (ISP)** - Nguyên tắc phân tách giao diện <br>
    `   Larger interfaces should be split into smaller ones. By doing so, we can ensure that implementing classes only need to be concerned about the methods that are of interest to them.` <br>
Nguyên tắc này áp dụng cho các Interface tương tự nguyên tắc đầu tiên Single Responsibility Principle của class.<br>
+ Một ví dụ thực tế trước nhé! <br>
Nếu bạn đi ăn ở một nhà hàng và bạn ăn chay, khi gọi món thì bạn sẽ nhận được một menu trong đó bao gồm các món mặn, chay,...mix từa lưa. Điều đó khiến bạn mất thời gian lọc và lựa món. Theo nguyên tắc này, chúng ta nên có một menu riêng các món thuần chay, không phải tất cả các món trong nhà hàng. Việc chia nhỏ sẽ giúp giảm tần suất thay đổi cần thiết. 
+ Ta có một interface người chăm sóc nhứng chú gấu: 
```
    public interface BearKeeper {
        void washTheBear();
        void feedTheBear();
        void petTheBear();
}
```
Thật tiếc, nhưng chúng ta không thể đảm nhận hết 3 công việc không mấy dễ dàng này, áp dụng nguyên tắc chia nhỏ giao diện ra nào
```
    public interface BearCleaner {
    void washTheBear();
    }

    public interface BearFeeder {
        void feedTheBear();
    }

    public interface BearPetter {
        void petTheBear();
    }
```
Bằng việc chia nhỏ interface, chúng ta implement interface công việc của chúng ta.
```
    public class BearCarer implements BearCleaner, BearFeeder {

    public void washTheBear() {
        //tắm cho gấu 
    }

    public void feedTheBear() {
        //bón cho gấu 
    }
}
```
```
    public class BravePerson implements BearPetter {

    public void petTheBear() {
        //Good luck!
    }
}
```
**Dependency Inversion Principle (DIP)** - Nguyên tắc đảo ngược phụ thuộc<br>
    ` This way, instead of high-level modules depending on low-level modules, both will depend on abstractions.` 
1. Các module cấp cao không nên phụ thuộc vào các modules cấp thấp. Cả 2 nên phụ thuộc vào abstraction.
2. Interface (abstraction) không nên phụ thuộc vào chi tiết, mà ngược lại. ( Các class giao tiếp với nhau thông qua interface, không phải thông qua implementation.)
+ Đi vào hiểu hơn về ví dụ <br>
```public class Windows98Machine {}```
Thêm thuộc tính và hàm khởi tạo cho class <br>
```
    public class Windows98Machine {

    private final StandardKeyboard keyboard;
    private final StandardKeyboard monitor;

    public Windows98Machine() {
        monitor = new Monitor();
        keyboard = new StandardKeyboard();
    }

}
```
class Windows98Machine phụ thuộc vào class StandardKeyboard, nếu ta muốn sửa đổi StandardKeyboard, thêm interface Keyboard tổng quát hơn và sử dụng trong class Windows98Machine <br>
```
    public interface Keyboard { }
```
```
    public class Windows98Machine{

    private final Keyboard keyboard;
    private final Monitor monitor;

    public Windows98Machine(Keyboard keyboard, Monitor monitor) {
        this.keyboard = keyboard;
        this.monitor = monitor;
    }
}
```
```
    public class StandardKeyboard implements Keyboard { }
```
*To be continute* <br>
Dependency Injection và Dependency Inversion <br>
[Tài liệu tham khảo -> https://www.baeldung.com/solid-principles]
Series <br>