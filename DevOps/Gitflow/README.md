# Gitflow Workflow
1. Gitflow <br>
   1.1 Tổng quát
    - Gitflow là một mô hình phân nhánh git, theo mô hình này developer tạo ra một nhánh tính năng (feature) riêng và merge vào nhánh chính khi tính năng được hoàn tất.
    - Gitflow Workflow phù hợp với mô hình phát triển phần mềm có thời gian release có chu kỳ như Scrum. Mô hình chỉ định vai trò cụ thể của các branches khác nhau và thời điểm mà chúng cần tương tác và có các branches đặc biệt cho việc preparing, maintaining và release. Ngoài ra chúng ta còn có thể cấu hình CI/CD trên mỗi Git repository để hệ thống tự động deloy khi có sự thay đổi về code (cụ thể là nhánh được dùng để build trong lúc setup CI/CD)<br>
   1.2. Mô hình
    - Các nhánh trong git-flow:
    Git-flow gồm có 2 nhánh chính là Master và Develop và 3 nhánh phụ gồm: Feature, Release, HotFix. 
    [![alt text](https://images.viblo.asia/84f47fd1-a009-4beb-8957-26395fe1023d.png)]()
    1. . Master branch
    Branch này có sẵn trong git và là branch chứa mã nguồn khởi tạo của ứng dụng và các version sẵn sàng để release chính thức cho người dùng sử dụng. Thưởng branch này cấu hình cho quản lý tương tác.
    2. Hotfix là các nhánh bảo trì được sử dụng để nhanh chóng sửa lỗi trên các bản phát hành. Các nhánh này gần giống với release và feature  trừ việc nó dựa trên master (không phải develop), là branch duy nhất fork từ master. Sau khi sửa xong, nhánh này nên được merge vào cả hai nhánh main và develop (hoặc nhánh release hiện tại) và phải được gắn thẻ với số phiên bản hiện tại
    3. Release
    Khi tất cả feature cần thiết đã hoàn thành để có thể release cho người dùng branch này sẽ được tạo, tên branch là tên của release version và đồng nghĩa với việc bắt đầu một vòng phát triển mới, không thêm tính năng mà tập trung vào fix bug, tạo document, sau khi hoàn tất merge lên master.
    4. Develop
    Được khởi tạo từ master branches để lưu lại tất cả lịch sử thay đổi của source. Nhánh này dùng để merge code của tất cả các branchs feature.
    5. Feature
    Được base trên branchs Develop. Mỗi khi phát triển một feature mới cần tạo một branch để viết code cho từng feature, cú pháp (feature_name-feature). Sau khi feature hoàn thành, dev tạo merge request đến branchs develop để teamlead review và merge lại vào branch dev.
2. Workflow và các lệnh thường sử dụng
[![Alt text](https://pbs.twimg.com/media/FZE95b-XkAUWWEK.jpg)](Workflow)
- Tạo một repository mới (trường hợp là người tạo project)
> git init
- Trường hợp đã có project, thì chỉ cần clone repository về local
> git clone [url]
> cd [name-project]
- Tạo một nhánh mới với feature tương ứng và thực hiện code chức năng tại nhánh này
> git checkout -b new-branch-name
- Nếu muốn nhảy qua nhanh khác
> git checkout branch-name
- Kiểm tra xem có những file nào đã thay đổi, file được thêm mới hoặc xóa đi.
> git status
- Sau khi kiểm tra và hoàn tất, thêm các thay đổi vào chúng sẽ được chuẩn bị và sẵn dàng dể được thêm vào kho lưu trữ
> git add .  *[add các file thay đổi, thêm file và xóa file*
> git add *      *[add tất cả các file trong project]
> git add -u     [*add các file được sửa đổi, không xóa các file*]
- Để ghi lại các thay đổi file hay thư mục vào repository thì thực hiện commit 
> git commit -m “your-name: content commit”
- Thực hiện đẩy code từ local lên remote
> git push origin branch-name
- Kiểm tra thông tin thay đổi giữa 2 nhánh
> git diff
- Kiểm tra thay đổi giữa 2 nhánh
> git diff first-branch second- branch
- Quay về một thời điểm commit trước đó và xóa lịch sử commit trước đó
> git reset –hard commit-id
- Xóa file
> git rm file-name
- Liệt kê tất cả các nhánh hiện tại ở local
> git branch
- xóa nhánh ở local
> git branch -d  branch-name
- Merge lịch sử của nhánh chỉ định vào nhánh hiện tại
> Git merge branch name
- Đang làm việc ở nhánh này mà có nhánh khác cần sử dụng, commna tạm thời lưu trữ tất cả các tệp được theo dõi đã sửa đổi
> git stash
- Sau khi làm vc ở nhánh vừa sửa thì quay lại
> git stash pop
- Xem danh sách các stash````
> git stash list
- Loại bỏ thay đổi được lưu gần nhất
> git stash drop
- Git patch:  là công cụ của git để“patch" những thay đổi mong muốn lại thành 1 file .patch thường được sử dụng khi muốn gửi file thay đổi mà chưa muốn commit lên source code.
Tạo patch file từ commit
> Git format-patch HEAD~n
n là số commits mới nhất muốn tạo patch
- Apply file patch
> git am < filename.patch
git am --signoff <filename.patch>
- Xem trước file patch sẽ thay đổi gì 
> git apply –stat file.patch
- Kiểm tra file patch có apply đc không
> git apply –check file.path
- git fetch: được sử dụng để tải xuống các nội dung từ Remote repository mà không làm thay đổi trạng thái của Local repository (các dữ liệu như commit, các file, refs)
- Xóa nhánh
> git checkout <another-branch>
> git push -d origin <branchname>   [delete remote branch]
> git branch -d <branchname>        [delete local branch]
- Đổi tên nhánh
  + Đi tới nhánh cần đổi:
 > git checkout <old-name>
  + Đổi tên nhánh ở local:
 > git branch -m <new_name>
  + Push nhánh mới lên:
 > git push origin -u <new_name>
  + Xóa tên nhánh cũ ở remote:
 > git push origin --delete <old_name>