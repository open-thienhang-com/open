<a href="https://jenkins.io">
  <picture>
    <source width="400" media="(prefers-color-scheme: dark)" srcset="https://www.jenkins.io/images/jenkins-logo-title-dark.svg">
    <img width="400" src="https://www.jenkins.io/images/jenkins-logo-title.svg">
  </picture>
</a>

[![Jenkins Regular Release](https://img.shields.io/endpoint?url=https%3A%2F%2Fwww.jenkins.io%2Fchangelog%2Fbadge.json)](https://www.jenkins.io/changelog)
[![Jenkins LTS Release](https://img.shields.io/endpoint?url=https%3A%2F%2Fwww.jenkins.io%2Fchangelog-stable%2Fbadge.json)](https://www.jenkins.io/changelog-stable)
[![Docker Pulls](https://img.shields.io/docker/pulls/jenkins/jenkins.svg)](https://hub.docker.com/r/jenkins/jenkins/)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3538/badge)](https://bestpractices.coreinfrastructure.org/projects/3538)

<br />
<br />
<br />

# Jenkins cheat sheet

<a href="https://www.edureka.co/blog/wp-content/uploads/2018/10/Jenkins-Cheat-Sheet-Edureka.pdf">
  <picture>
    <source width="" media="(prefers-color-scheme: dark)" srcset="./Jenkins-Cheat-Sheet-Edureka-1.png">
    <img width="" src="./Jenkins-Cheat-Sheet-Edureka-1.png">
  </picture>
</a>

#### Xem thêm tại <a href="https://www.jenkins.io/user-handbook.pdf"> Jenkins User Handbook</a>

<br />
<br />
<br />

# Tìm hiểu thêm

## Giới thiệu chung

Jenkins là một <span style="font-style:italic">công cụ tự động hóa có mã nguồn mở</span> được viết bằng Java kết hợp với nhiều Plugin, Jenkins có mục đích chính là tích hợp liên tục (hay còn gọi là CI- Continuous Integration). Công cụ này được sử dụng để xây dựng và kiểm tra các dự án phần mềm liên tục, giúp các nhà phát triển tích hợp các thay đổi vào dự án một cách dễ dàng hơn.

Jenkins cũng cho phép bạn liên tục cung cấp phần mềm của mình bằng cách tích hợp với một số lượng lớn các công nghệ thử nghiệm hoặc đã triển khai.

Với Jenkins, các tổ chức có thể đẩy nhanh quá trình phát triển phần mềm thông qua tự động hóa. Jenkins tích hợp các vòng đời phát triển của một quy trình bao gồm xây dựng, tài liệu, thử nghiệm, gói, giai đoạn, triển khai, phân tích tĩnh,…

Bên cạnh đó, các Plugin trong Jenkins cho phép tích hợp các giai đoạn DevOps khác nhau. Nếu bạn muốn tích hợp một công cụ cụ thể, bạn cần cài đặt Plugin cho công cụ đó.

### Lịch sử ra đời của phần mềm Jenkins

Dự án Jenkins được bắt đầu vào năm 2004 bởi Kohsuke Kawaguchi, khi ông đang làm việc cho Sun Microsystems. Kohsuke là một nhà phát triển tại Sun và cảm thấy mệt mỏi khi phải hứng chịu các cơn thịnh nộ của nhóm mình mỗi khi code của anh ta thất bại.

Kohsuke đã tạo ra Jenkins như một cách để thực hiện tích hợp liên tục. Sau đó, Kohsuke đã mở nguồn và tạo ra dự án Jenkins. Chẳng bao lâu sau, việc sử dụng Jenkins đã lan rộng khắp thế giới. Bản Jenkins đầu tiên được phát hành vào tháng 2 năm 2005.

Sau khi Oracle mua lại Sun Microsystems, cộng đồng Hudson đã chấp thuận những đề xuất để tạo ra dự án Jenkins. Vào tháng 2 năm 2011, Hudson đã được tách ra thay vì đổi tên thành Jenkins.

Mặc dù Hudson và Jenkins được phát triển độc lập, Jenkins lại có được nhiều dự án và cộng tác viên hơn Hudson. Do đó, Hudson đã không còn được cộng đồng duy trì.

Đến năm 2014, Kawaguchi đã trở thành CTO của CloudBees, một công ty chuyên cung cấp các sản phẩm dựa trên chính nền tảng Jenkins.

### Jenkins hoạt động như thế nào?

Jenkins là một ứng dụng dựa trên máy chủ và đòi hỏi phải có một máy chủ web như Apache Tomcat để chạy trên các nền tảng khác nhau như Windows, Linux, macOS, Unix,..Để sử dụng Jenkins, bạn cần tạo các đường dẫn gồm một loạt các bước mà một máy chủ Jenkins sẽ nhận. Tích hợp liên tục là một công cụ mạnh mẽ bao gồm một bộ công cụ được thiết kế để lưu trữ,
giám sát, biên dịch và kiểm tra mã hoặc các thay đổi mã.

- Máy chủ tích hợp liên tục, ví dụ như: Jenkins, Bamboo, CruiseControl, TeamCity,..
- Công cụ kiểm soát nguồn, ví dụ như: CVS, SVN, GIT,Mercurial, Perforce, ClearCase và các công cụ khác
- Công cụ xây dựng, ví dụ như: Make, ANT, Maven, Ivy, Gradle và các công cụ khác
- Framework kiểm tra tự động hóa, ví dụ như: Selenium, Appium, TestComplete, UFT và những thứ khác

<div align="center">
    <picture>
        <img  src="https://wiki.tino.org/wp-content/uploads/2021/07/word-image-1163.png">
    </picture>
</div>

## CI và CD trong Jenkins

### CI (Continuous Integration)

CI là viết tắt của Continuous Integration, tạm dịch: Tích hợp liên tục. Đây là một quá trình tích hợp các thay đổi code từ nhiều nhà phát triển trong một dự án. Phần mềm được kiểm tra ngay lập tức sau khi code commit. Với mỗi code commit, code sẽ được tiến hành xây dựng và thử nghiệm. Nếu thử nghiệm được thông qua, bản dựng sẽ được thử nghiệm để triển khai. Nếu việc triển khai thành công, mã sẽ được chuyển sang sản xuất.

Quá trình cam kết, xây dựng, kiểm tra và triển khai này là một quá trình liên tục và do đó mới được gọi là tích hợp/triển khai liên tục.

### CD (Continuous Delivery)

CD là viết tắt của Continuous Delivery, tạm dịch: Chuyển giao liên tục. Đây là quy trình triển khai tất cả thay đổi trong quá trình tự động test và deploy các code lên các môi trường staging và production. Ngoài ra, CD còn hỗ trợ tự động hóa phần testing bên cạnh việc sử dụng units test. Mục đích của CD là thử nghiệm phần mềm liên tục để kiểm tra hệ thống trước khi bàn giao cho khách hàng.

<div  align="center">
    <picture>
        <img src="https://wiki.tino.org/wp-content/uploads/2021/07/word-image-1164.png">
    </picture>
</div>

## Ưu điểm và nhược điểm của Jenkins

### Ưu điểm

- Jenkins đang được quản lý rất tích cực. Mỗi tháng, công ty phát hành Jenkins sẽ tổ chức các cuộc họp công khai và lấy ý kiến ​​đóng góp từ cộng đồng để phát triển dự án Jenkins.
- Jenkins đã có khoảng 320 Plugin được xuất bản trong cơ sở dữ liệu Plugin của mình.
- Công cụ Jenkins cũng hỗ trợ kiến ​​trúc đám mây để bạn có thể triển khai Jenkins trên các nền tảng dựa trên đám mây.
- Tích hợp với nhiều nền tảng CI/CD và giúp cho team được đồng bộ hóa
- Rất dễ dàng để tìm ra các lỗi trong Jenkins. Nhà phát triển có thể kiểm tra lỗi và giải quyết chúng nhanh chóng.
- Rút ngắn thời gian bàn giao dự án và linh hoạt trong công việc
- Jenkins hỗ trợ các loại kho mã nguồn khác nhau như SVN, Git, v.v. Nhà phát triển có thể đặt các trình kích hoạt khác nhau sau khi thực hiện các thay đổi trong mã.

### Nhược điểm

- Giao diện của Jenkins đã lỗi thời và không thân thiện với người dùng so với xu hướng hiện tại.
- Việc quản lý bảng điều khiển Jenkins khá khó khăn khi chúng ta có quá nhiều công việc phải thực hiện.
- Việc bảo trì Jenkins không dễ dàng vì phần mềm này chạy trên một máy chủ và yêu cầu một số kỹ năng như quản trị viên máy chủ để giám sát hoạt động của nó.
- Gặp một số khó khăn trong việc cài đặt và cấu hình Jenkins.
- Bạn phải tự mình bảo trì cơ sở hạ tầng.
- Tích hợp liên tục thường xuyên bị hỏng do một số thay đổi nhỏ trong cài đặt.
- Các thay đổi do một nhà phát triển thực hiện sẽ được không hiển thị với một nhà phát triển khác trong nhóm và chỉ có người quản lý mới có quyền truy cập. Điều này làm cho việc theo dõi dự án lớn gặp khó khăn.
