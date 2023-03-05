# Kubernetes

![](https://longvan.net/hinhanh/tintuc/gioi-thieu-ve-kubernetes.jpg)

## 1. Kubernetes là gì?
Để dễ hiểu Kubernetes là gì thì chúng ta hãy liên tưởng đến 1 dàn nhạc.
    
### Orchestra - dàn nhạc
![](https://img.freepik.com/free-vector/conductor-musicians-standing-theater-stage-flat-illustration-cartoon-people-playing-violin-cello-harp_74855-10532.jpg)
Một dàn nhạc sẽ bao gồm:
- Conductor - nhạc trưởng: Có vai trò chỉ huy dàn nhạc, giữ nhịp độ, điều tiết các nhạc cụ, ...
- Composer - nhà soạn nhạc: Điều tiết nhịp độ, giúp các nhạc cụ vào đúng thứ tự, chỉ dẫn về sắc thái mà các nhạc cụ cần phải biểu đạt. 
- Nhạc công: Thực hiện các nhạc cụ theo đúng yêu cầu.

### Kubernetes
- Là một Container Orchestration Tool, là một tool open source được phát triển bởi Google.
- Kubernetes đóng vai trò như là một nhạc trưởng của các container hay là service (đóng vai trò như những nhạc công)
- Kubernetes đảm bảo cho các service hoạt động trơn tru theo chỉ dẫn từ 1 file config. File này có đuôi là yaml hoặc là yml. File config này đóng vai trò như là Composer - nhà soạn nhạc vậy.

## 2. Tại sao Kubernetes lại được tạo ra?
Trước tim chúng ta cần tìm hiểu về các kiến trúc sau:

### Monolith Architecture
- Trong kiến trúc Monolith thì tất cả các chức năng bao gồm các thư viện của bên thứ 3 đều được gộp chung vào một deployment.
- Với các ứng dụng nhỏ thì kiến trúc này hoàn toàn ok, không vấn đề gì. Tuy nhiên, việc deploy sẽ mất rất nhiều thời gian do các module phải đi cùng nhau hoặc phải cùng công nghệ.
- Nếu mỗi module do một team phát triển thì điều này sẽ gây phức tạp hơn.
- Khi muốn mở rộng hoặc scale chúng ta sẽ phải update tất cả các module điều này sẽ gây mất thời gian và hệ thống sẽ gặp rất nhiều rủi ro.
- Vì vậy, kiến trúc Microservice được tạo ra để khắc phục điều này.

### Microservice Architecture
- Trong kiến trúc này thì một chức năng lớn được chia thành nhiều chức năng nhỏ hơn (được gọi là các service).
- Việc scale hệ thống sẽ dễ dàng hơn vì chỉ cần update các service nào liên quan mà không phải update toàn hệ thống.
- Dễ dàng cho phát triển bởi nhiều team với nhau và có thể sử dụng nhiều công nghệ.
- Tuy nhiên, có một vấn đề trong kiến trúc này là một service cần phải được chạy trên một machine. (Một machine ở đây có thể là một PC, Laptop, Host,...) Vì thế nếu trong hệ thống của ta có nhiều service thì số lượng machine sẽ phải tăng lên -> Dẫn đến rất tốn kém.

![Kiến trúc monolith và microservice](https://www.openlegacy.com/hubfs/Picture1.webp "Kiến trúc monolith và microservice")

### Docker
- Để giải quyết vấn đề trên, chúng ta có khái niệm container. Các service thay vì được chứa vào machine sẽ được chứa vào container.
- Các ứng dụng quản lý container được gọi là Containerized Application.
- Docker là một trong những ứng dụng quản lý các container. Được vận hành bởi hệ điều hành ảo hoá, được sử dụng để đóng gói và chạy các phần mềm trong máy. Các containers này tách biệt với nhau và đóng gói những ứng dụng, những dữ liệu, thư viện, cơ cấu tổ chức thư mục riêng của nó, và những ứng dụng hay dữ liệu này được liên kết chặt chẽ với nhau thông qua kênh riêng biệt.
![](https://www.docker.com/wp-content/uploads/2021/11/docker-containerized-appliction-blue-border_2.png)
- Một machine sẽ cài được 1 Containerized Application.
- Trong ứng dụng Containerized sẽ chứa nhiều container và mỗi container sẽ quản lý một hoặc nhiều service.
- Đặt vấn đề:
    - Giả sử một container trong docker được nâng cấp thì container này phải được tắt đi và bật container mới nâng cấp lên.
    - Thì khoảng thời gian tính từ lúc tắt container cũ đi và bật container mới lên được gọi là downtime cần phải phải giảm thiểu xuống thấp nhất có thể. Và làm sao để container này thông báo cho các container khác biết là nó đang update?
    - Và chính vì Docker hoạt động dựa trên việc tổ chức thông tin, quản lý chúng bằng cách "đóng gói" chúng thì cần một hệ thống giúp chúng ta chỉnh sửa, hỗ trợ việc sử dụng, quản lý các ứng dụng đó.
    - Những vấn đề này sẽ được giải quyết bằng Kubernetes.

### Kubernetes
- Kubernetes API nằm trên master node thông qua K8s Control Plane sẽ giao tiếp với các node ở phía slave node, mỗi slave node gồm nhiều pod(pod là đơn vị nhỏ nhất trong kubernetes, quản lý nhiều hoặc một container).
- Nhiều node trên slave node sẽ được nhóm lại thành các cluster. Bên trong các cluster thì các container sẽ được giao tiếp với nhau thông qua địa chỉ IP (địa chỉ ip này được tạo ra từ virtual network bên trong cluster).
- Lưu ý, địa chỉ ip này chúng ta không thể ping, connect được từ máy tính cài kubernetes mà cần có những config proxy để giải quyết được vấn đề đó.
- Giả sử các lập trình viên muốn tuỳ biến cluster theo các yêu cầu nào đấy, tất cả các yêu cầu này phải nằm trong file có đuôi là yml, yaml sau đó gửi file này lên cho kubernetes thực hiện.
- Vậy để sử dụng Kubernetes cluster cho ứng dụng doanh nghiệp của mình thì chúng ta phải biến các module trong ứng dụng thành container.
- Giả sử ứng dụng của chúng ta có NodeJS container, Java spring container, MySQL container, SQL Server container. Như vậy host là nơi chứa các container và mỗi host cần phải có 1 địa chỉ ip. Địa chỉ ip này là nội bộ và nằm trong network của mỗi cluster. Nếu 1 container bị crack thì pod sẽ bị restart và 1 pod khác sẽ được tái sinh, sẽ có nhiều node giống nhau ta gọi đó là replicas.

![](https://www.opsramp.com/wp-content/uploads/2022/07/Kubernetes-Architecture.png)

