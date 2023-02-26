# MESSAGE BROKER

## MESSAGE BROKER LÀ GÌ?

Message broker (hay còn gọi là integration broker hoặc interface engine) là một module trung gian trung chuyển message từ người gửi đến người nhận. Nó là một mô hình kiến trúc (architentural pattern) để kiểm tra, trung chuyển và điều hướng message; làm trung gian giữa các ứng dụng với nhau, tối giản hóa giao tiếp giữa các ứng dụng đó và để tăng hiệu quả tối đa cho việc tách ra các khối nhỏ hơn. Nhiệm vụ chính của một Message broker là tiếp nhận những message từ các ứng dụng và thực hiện một thao tác nào đó.

Hiện tại có rất nhiều các message broker software có thể kể đến như: Amazon Web Services (AWS) Simple Queue Service (SQS), Apache Kafka, Apache ActiveMQ. Nhưng phổ biến nhất trong số những cái tên kể trên đó là RabbitMQ!

## PHÂN LOẠI GIỮA CÁC HỆ THỐNG MESSAGE QUEUE

Nếu bạn là một backend-developer chắc hẳn bạn không còn xa lạ gì với những hệ thống message-queue. Hầu như project nào của mình cũng có sự xuất hiện của message-queue và việc khó khăn nhất là lựa chọn nên dùng cái nào và không nên dùng cái nào.

Thử tưởng tượng một ngày sếp của bạn muốn tích hợp một hệ thống message queue hoặc bạn cảm thấy nên sử dụng một hệ thống message queue để giải quyết bài toán mà team đang gặp phải. Bạn bắt đầu tìm kiếm và nhận ra rằng có quá nhiều hệ thống message queue tồn tại. Mình có thể liệt kê một số loạt mà mình biết dưới đây :

- RabitMQ
- ActiveMQ
- Kafka
- SQS
- ZeroMQ
- MSMQ
- IronMQ
- Kinesis
- RocketMQ
- Redis Pub/sub

Trong quá khứ mình đã có không ít lần lựa trọn sai và mọi sai lầm đều phải trả giá không ít thì nhiều.

Sau những sai lầm đó thì mình nhận ra một điều, đó là mặc dù cùng là message queue nhưng lại được chia làm 2 loại có mục đich sử dụng và những tính năng liên quan hoàn toàn khác nhau.

Mình tạm chia thành 2 loại như sau:

|  Message Base  | Data Pipeline  |
| ------------ | ------------ |
| RabitMQ  |  	Kafka |
| ActiveMQ  | 	Kinesis  |
|  SQS | 	RocketMQ  |
|  ZeroMQ |   |
| MSMQ  |   |
| IronMQ  |   |
| Redis Pub/sub  |   |

![](https://topdev.vn/blog/wp-content/uploads/2021/01/16-1.png)

Dựa vào bảng trên, ta có thể thấy được sự khác nhau cơ bản giữa 2 loại, cũng như cách sử dụng trong tưng bài toán cụ thể.

Đối với loại “message base”: là những loại message queue truyền thống, thích hợp làm hệ thống trao đổi message giữa các service. Việc đảm bảo mỗi consumer đều nhận được message và duy nhất một lần là quan trọng nhất.

Đối với loại “data-pipeline”, có cách lưu trữ message cũng như truyền tải message đến consumer hoàn toán khác với hệ thống message queue truyền thống. Việc đảm bảo mỗi consumer đều phải nhận được message và duy nhất một lần không phải là ưu tiên số một, mà thay vào đó là khả năng lưu trũ message vả tốc độ truyền tải message. Khi có message mới, consumer sẽ lựa chọn số lượng message mà mình muốn lấy, chính vì thế mà cùng một message consumer có thể nhận đi nhận lại nhiều lần. Những hệ thống sử dụng message queue loại này thường là hệ thống Event Sourcing, hoặc hệ thống đồng bộ dữ liệu từ những database khác nhau như Debezium.

Khi các bạn lựa chọn message queue cho hệ thống của mình, các bạn nên xác định rõ mục địch của hệ thống messague queue để xem mình cần loại trong hai loại trên. Việc xác định được loại message queue nào mình cần sẽ giúp các bạn giảm bớt thời gian tìm hiểu cũng như tìm được chính sác cái mà mình cần.

Đôi khi chúng ta cũng thấy một số hệ thống sẽ sử dụng nhiều loại message queue, thường sẽ là 1 của “message base” và 1 của “data pipeline” để tận dụng tối đa ưu điểm của từng loại vào giải quyết bài toán cụ thể.

# KAFKA

## Kafka là gì?

Kafka là gì? Là hệ thống message pub/sub phân tán (distributed messaging system). Bên public dữ liệu được gọi là producer, bên subcribe nhận dữ liệu theo topic được gọi là consumer. Kafka có khả năng truyền một lượng lớn message theo thời gian thực, trong trường hợp bên nhận chưa nhận message vẫn được lưu trên một hàng đợi và cả trên ổ đĩa đảm bảo an toàn. Đồng thời nó cũng được replicate trong cluster giúp phòng tránh mất dữ liệu.

![](https://topdev.vn/blog/wp-content/uploads/2019/05/kafka-simple.png)

## Các khái niệm cơ bản

Kafka là gì? – Có thể hiểu là một hệ thống logging để lưu lại các trạng thái của hệ thống đề phòng tránh mất thông tin.

Định nghĩa trên được giải thích bằng các khái niệm sau:

- PRODUCER: Kafka lưu, phân loại message theo topic, sử dụng producer để publish message vào các topic. Dữ liệu được gửi đển partition của topic lưu trữ trên Broker.
- CONSUMER: Kafka sử dụng consumer để subscribe vào topic, các consumer được định danh bằng các group name. Nhiều consumer có thể cùng đọc một topic.
- TOPIC: Dữ liệu truyền trong Kafka theo topic, khi cần truyền dữ liệu cho các ứng dụng khác nhau thì sẽ tạo ra cá topic khác nhau.
- PARTITION: Đây là nơi dữ liệu cho một topic được lưu trữ. Một topic có thể có một hay nhiều partition. Trên mỗi partition thì dữ liệu lưu trữ cố định và được gán cho một ID gọi là offset. Trong một Kafka cluster thì một partition có thể replicate (sao chép) ra nhiều bản. Trong đó có một bản leader chịu trách nhiệm đọc ghi dữ liệu và các bản còn lại gọi là follower. Khi bản leader bị lỗi thì sẽ có một bản follower lên làm leader thay thế. Nếu muốn dùng nhiều consumer đọc song song dữ liệu của một topic thì topic đó cần phải có nhiều partition.
- BROKER: Kafka cluster là một set các server, mỗi một set này được gọi là 1 broker
- ZOOKEEPER: được dùng để quản lý và bố trí các broker.

![](https://topdev.vn/blog/wp-content/uploads/2019/05/kafka-structure.png)

## Một vài use case cho kafka:

- Sử dụng như một hệ thống message queue thay thế cho ActiveMQ hay RabbitMQ
- Website Activity Monitoring: theo dõi hoạt động của website
- Stream Processing: Kafka là một hệ thống rất thích hợp cho việc xử lý dòng dữ liệu trong thời gian thực. Khi dữ liệu của một topic được thêm mới ngay lập tức được ghi vào hệ thống và truyền đến cho bên nhận. Ngoài ra Kafka  còn là một hệ thống có đặc tính duribility dữ liệu có thể được lưu trữ an toàn cho đến khi bên nhận sẵn sàng nhận nó.
- Log Aggregation: tổng hợp log
- Metrics Collection: thu thập dữ liệu, tracking hành động người dùng như các thông số như page view, search action của user sẽ được publish vào một topic và sẽ được xử lý sau
- Event-Sourcing: Lưu lại trạng thái của hệ thống để có thể tái hiện trong trường hợp system bị down.

