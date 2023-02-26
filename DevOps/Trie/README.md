# Cây tìm kiếm nhị phân là gì?
Cây tìm kiếm nhị phân có tên tiếng anh là Binary Search Tree (BST), là một trong những cấu trúc dữ liệu cơ bản bên cạnh queue, stack, linked-list, array.
Cây tìm kiếm nhị phân là 1 dạng đồ thị nhưng các nút (node) của cây phải có những tính chất sau:
- Mỗi node chỉ có thể có tối đa 2 node con
- Giá trị của node con bên trái nhỏ hơn node cha của nó
-  Giá trị của node con phải lớn hơn node cha của nó
Tính chất cũng phỉa đúng với các node của của 2 node con trên, nói cách khác giá trị tất cả các con bên trái của 1 node phải nhỏ hơn giá trị của node đó và giá trị của tất cả các con bên phải của node đó phải lớn hơn giá trị của nó. Sự so sánh giá trị ở trên là so sánh toán học, so sánh chuỗi kí tự,...

[![](https://codelearn.io/Media/Default/Users/TrisDo/BST/sth.png)](https://codelearn.io/Media/Default/Users/TrisDo/BST/sth.png)

Đặc biệt trong 1 cây tìm kiếm nhị phân không cho phép 2 giá trị trùng nhau. Chính quy luật và cách sắp xếp như trên cấu trúc BST đã giúp sắp xếp dữ liệu theo một cách có trật tự, từ đó giúp người sử dụng dễ dàng hơn trong việc tổ chức dữ liệu cũng như việc tìm kiếm.
Sau đây là một vài khái niệm trong cây tìm kiếm nhị phân:
 - Root node (nút gốc): node đầu tiên của cây
 - Leaf node (nút lá): Node không có node con trái phải
 - Internal node: Những node không phải nút gốc cũng không phải nút lá
 - Level: như hình minh hoạ trên chúng ta có 2 cây với 3 tầng.
 Về cơ bản chúng ta có 3 loại cây nhị phân:

 [![](https://codelearn.io/Media/Default/Users/TrisDo/BST/sth2.png)](https://codelearn.io/Media/Default/Users/TrisDo/BST/sth2.png)

 Full binary tree: Những node không phải nút lá đều có 2 con trái và phải
 Complete binary tree: Tất cả các tầng đều chứa đầy nodes ngoại trừ tầng cuối có thể đầy hoặc không nhưng các node tầng cuối phải được xếp lần lượt từ trái đến phải.
 Perfect binary tree: tất cả nodes đều có 2 con và các nút lá ở cùng một level.
 Với cây tìm kiếm nhị phân chúng ta có những tác vụ cơ bản sau:
  	 - Search: Tìm kiếm
   	-  Insert: Thêm 1 node
   	- Remove: Xoá 1 note
   	- Traversal: Duyệt cây với 3 loại cơ bản: pre-oder traversal, In-order traversal, post-order traversal.
Bên trên là những khái niệm cơ bản về cấu trúc dữ liệu cây nói chung và cây tìm kiếm nhị phân nói riêng, sau đây chúng ta đến với cách hoạt động của các tác vụ trong cây tìm kiếm nhị phân.

# Cây tìm kiếm nhị phân dùng để làm gì?

## Search
Để bắt đầu hoạt động tìm kiếm một giá trị x cho trước, chúng ta bắt đầu từ root. Nếu giá trị x nhỏ hơn giá trị của root thì chuyển đến so sánh với node con trái của root và ngược lại.
Tiếp tục quá trình xét như trên với các node tiếp theo đến khi tìm được, còn nếu đến nút lá mà so sánh x không bằng giá trị nút lá thì xác nhận không tìm thấy.

## Insert
Khi muốn thêm một giá trị x vào cây nhị phân, ta bắt đầu tìm kiếm từ nút gốc (root), nếu giá trị x nhỏ hơn giá trị nút gốc thì tìm vị trí trống của cây con bên trái nút gốc, nếu x lớn hơn giá trj nút gốc ta tìm vị trí trống của cây con bên phải nút gốc. Trường hợp tìm được giá trị của 1 node trong cây bằng với x thì xác nhận x đã tồn tại trong cây.

 [![](https://codelearn.io/Media/Default/Users/TrisDo/BST/sth5.png)](https://codelearn.io/Media/Default/Users/TrisDo/BST/sth5.png)

 Với ví dụ trong ảnh trên, ta cần thêm 4 vào trong cây nhị phân cho trước. Bắt đầu từ nút gốc là 6, vì 4 < 6 nên tìm vị trí trống phía cây con bên trái của 6, tiếp theo 4 > 3 và có một vị trí trống phía bên phải cảu nút 3 vây nên đó là vị trí phù hợp để thêm 4 vào trong cây.
 
## Remove
Trong hoạt động xoá 1 node của cây nhị phân chúng ta sẽ găp phải 3 trường hợp sau:
  	- Node cần xoá chỉ có 1 node con (trái hoặc phải)
	- Node cần xoá không có node con
	- Node cần xoá có cả 2 node
Với trường hợp đầu tiên node cần xoá có 1 node con, ta chỉ cần thay vị trí của node con đó với node cần xoá

[![](https://codelearn.io/Media/Default/Users/TrisDo/BST/sth6.png)](https://codelearn.io/Media/Default/Users/TrisDo/BST/sth6.png)

Với trường hợp node cần xoá không có node con thì chúng ta đơn giản chỉ cần xoá vị trí node đó trong cây.

[![](https://codelearn.io/Media/Default/Users/TrisDo/BST/sth7.png)](https://codelearn.io/Media/Default/Users/TrisDo/BST/sth7.png)

Trường hợp cuối cùng node cần xoá có cả 2 node. Với trường hợp này việc của ta cần làm là tìm được 1 node thế (successor) để lắp vào vị trí của node cần xoá, nói cách khác node thế phải có tính chất bé hơn tất cả node bên trái của node cần xoá và lớn hơn tất cả các node bên phải của node cần xoá.

[![](https://codelearn.io/Media/Default/Users/TrisDo/BST/Annotation%202020-08-12%20201326.jpg)](https://codelearn.io/Media/Default/Users/TrisDo/BST/Annotation%202020-08-12%20201326.jpg)

Với cây nhị phân trong hình, khi ta muốn xáo node 5 thì node 6 chính là node thay thế cho node 5, node 6 còn được gọi là left-most tree, tức node trái cùng của 1 cây. Sau khi thay thế node 5 bằng node 6 ta cần xoá node 6 ở vị trí cũ đi, khi này ta quay trở lại trường hợp xoá 1 node có 1 node con tại vì node 6 cũ có 1 node con là 7

[![](https://codelearn.io/Media/Default/Users/TrisDo/BST/sth9.png)](https://codelearn.io/Media/Default/Users/TrisDo/BST/sth9.png)

Node thế (successor) có thể là left-most của cây con bên phải hoặc right-most tree của cây con bên trái. Với left-most tree được định nghĩa là con trái cùng hay giá trị nhỏ nhất trong cây nhị phân, right-most là con phải cùng, hay giá trị lớn nhất trong cây nhị phân.

## Pre-order traversal
Với cách duyệt này ta sẽ đi qua node cha trước sau đến node con trái rồi đến node con phải

[![](https://codelearn.io/Media/Default/Users/TrisDo/BST/sth10.png)](https://codelearn.io/Media/Default/Users/TrisDo/BST/sth10.png)

Với cây trên thứ tự các code sau khi duyệt là: 6,3,1,10,9,12

## In-order traversal
Ta duyệt lần lượt node con trái, node con trái sau đến node cha rồi đến node con phải.

[![](https://codelearn.io/Media/Default/Users/TrisDo/BST/sth10.png)](https://codelearn.io/Media/Default/Users/TrisDo/BST/sth10.png)

Với cây trên thứ tự các node sau khi duyệt là 1,3,6,9,10,12.

## Post-order traversal
Ta duyệt lần lượt node con trái, node con phải rồi đến node cha

[![](https://codelearn.io/Media/Default/Users/TrisDo/BST/sth10.png)](https://codelearn.io/Media/Default/Users/TrisDo/BST/sth10.png)

Với cây trên thứ tự code sau khi duyệt là 1,3,9,12,10,6

#Trie
Trie là một cấu trúc dữ liệu dùng để quản lý một tập hợp các xâu. Trie cho phép:
Trie là một cấu trúc dữ liệu dùng để quản lý một tập hợp các xâu. Trie cho phép:
- Thêm một xâu vào tập hợp.
- Xoá một xâu khỏi tập hợp.
- Kiểm tra một xâu có tồn tại trong tập hợp hay không.

# Cấu trúc
Trie gồm một gốc không chứa thông tin, trên mỗi cạnh lưu một ký tự, mỗi nút và đường đi từ gốc đến nút đó thể hiện 1 xâu, gồm các ký tự thuộc cạnh trên đường đi đó.

[![](https://vnoi.info/wiki/uploads/trie.png)](https://vnoi.info/wiki/uploads/trie.png)

Trong hình vẽ trên, nút 1 là nút gốc, nút 7 thể hiện có 1 xâu là 'bg', nút 8 thể hiện có 1 xâu là 'db', nút 9 thể hiện 1 xâu là 'dc', nút 10 thể hiện 1 xâu là 'acd', nút 5 thể hiện là có 1 xâu là 'abc'.
Tuy nhiên, như các bạn có thể thấy, đối với một số nút, chẳng hạn như nút 4, ta không biết nó là thể hiện kết thúc 1 xây hay chỉ là 1 phần của đường đi từ nút 1 đến nút 9. Vì vậy, khi cài đặt, thông thường, tại nút U ta cần lưu thêm thông tin nú U có là kết thúc của 1 xâu hay không, hoặc nút U là kết thúc của bao nhiêu xâu, tuỳ theo yêu cầu bài toán.

# Một số ưu điểm của trie
- Cài đặt đơn giản, dễ nhớ
- Tiết kiệm bộ nhớ: Khi số lượng khoá lớn và các khoá có độ dài nhỏ, thông thường trie tiết kiệm bộ nhớ do các phần đầu giống nhau của các khoá chỉ được lưu 1 lần. Ưu điểm này có ứng dụng rất lớn, chẳng hạn trong từ điển.
- Thao tác tìm kiếm: O(m) với m là độ dài khoá. Với Binary search tree(cân bằng): là O(logN). Khi số lượng khoá cần tìm lớn và độ dài mỗi khoá tương đối nhỏ, logN xấp xỉ m, và như các bạn đã biết, để cài được Binary search tree cân bằng không phải là việc đơn giản. Hơn nữa, các thao tác trên trie rất đơn giản và thường chạy nhanh hơn trên thực tế.
- Dựa vào tính chất của cây trie, có thể thực hiện một số liên quan đến thứ tự từ điển như sắp xếp, tìm một khoá có thứ tự từ điển nhỏ nhất và lớn hơn một khoá cho trước,...; và một số thao tác liên quan đến tiền tố, hậu tố.

# Cài đặt
```go
package main

import "fmt"

const (
	//ALBHABET_SIZE total characters in english alphabet
	ALBHABET_SIZE = 26
)

type trieNode struct {
	children  [ALBHABET_SIZE]*trieNode
	isWordEnd bool
}

type trie struct {
	root *trieNode
}

func initTrie() *trie {
	return &trie{
		root: &trieNode{},
	}
}

func (t *trie) insert(word string) {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			current.children[index] = &trieNode{}
		}
		current = current.children[index]
	}
	current.isWordEnd = true
}

func (t *trie) find(word string) bool {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	if current.isWordEnd {
		return true
	}
	return false
}

func main() {
	trie := initTrie()
	words := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses"}
	for i := 0; i < len(words); i++ {
		trie.insert(words[i])
	}
	wordsToFind := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses", "rosess", "ans", "san"}
	for i := 0; i < len(wordsToFind); i++ {
		found := trie.find(wordsToFind[i])
		if found {
			fmt.Printf("Word \"%s\" found in trie\n", wordsToFind[i])
		} else {
			fmt.Printf("Word \"%s\" not found in trie\n", wordsToFind[i])
		}
	}
}

```