# Bài 1

Công ty Sendo có kế hoạch tổ chức 1 sự kiện lớn và gửi giấy mời đến toàn bộ shop, sau đó, các shop đồng ý đến sự kiện sẽ báo lại thông tin ngày, cũng như thời gian bắt đầu, thời gian kết thúc mà shop đó có thể tham gia được sự kiện. Bạn hãy viết giúp một chương trình để tìm ra ngày, cũng như thời gian bắt đầu để làm sao có số lượng shop có thể tham gia là lớn nhất

## Input

Đầu vào là một file văn bản tên Input.txt, dòng đầu tiên ghi số CASE là số lượng testcase sẽ có (1 ≤ CASE ≤ 100). Dòng tiếp theo là N, là số lượng shop đồng ý đến sự kiện (0 ≤ N ≤ 5.000.000) của testcase đầu tiên. Trong N dòng tiếp theo, mỗi dòng ghi 3 giá trị DATE (dạng yymmdd, yy là 19 hoặc 20), START, END (0 ≤ START < END ≤ 23, là các số nguyên) các giá trị cách nhau 1 dấu cách. File tiếp tục lặp lại cho đến hết số lượng testcase.

## Output

Đầu ra là một file văn bản tên Output.txt có CASE dòng. Mỗi dòng ghi ra ba giá trị NUMBER, DATE, START cách nhau 1 dấu cách, trong trường hợp không tìm thấy ngày/giờ thích hợp thì ghi ra ba số -1 -1 -1. Trong đó NUMBER chính là số shop lớn nhất có thể tham dự. Trong trường hợp có 2 ngày/giờ thích hợp có cùng số lượng shop có thể tham gia lớn nhất và bằng nhau thì chọn ngày/giờ nhỏ nhất

## Ví dụ

### Input.txt
```
2
3
200627 5 8
200627 6 7
200626 1 2
5
200115 8 9
200116 8 9
200116 1 2
200115 6 7
191213 0 1
```

### Output.txt
```
2 200627 6
1 191213 0
```


# Bài 2

Vào ngày 13/13 công ty Sendo sẽ có event bán hàng siêu đặc biệt. Sendo sẽ bán N sản phẩm (0 ≤ N ≤ 1.000) với giá lần lượt là A1, A2, … AN. (1 ≤ Ai ≤ 1.000.000). Người mua hàng bắt buộc cần mua đủ cả N sản phẩm nhưng có thể chia ra mua nhiều lần. Nếu mỗi lần mua, khách hàng mua M sản phẩm (M ≤ N) mà M >= S với S là một giá trị cho trước, thì trong lần mua đó khách hàng không phải trả tiền cho sản phẩm có giá nhỏ nhất trong lần mua đó, trường hợp có 2 sản phẩm cùng có giá nhỏ nhất trong lần mua đó thì cũng chỉ không phải trả cho 1 sản phẩm. (2 ≤ S ≤ 100)

Yêu cầu: Tìm số tiền nhỏ nhất mà người mua có thể dùng để mua đủ N sản phẩm

## Ví dụ

Có N=3 sản phẩm với giá lần lượt là 10, 20, 30, cho trước S=2
Khách hàng có thể có các cách mua sau:

### Cách 1
- Mua 1 lần cả 3 sản phẩm, lúc này, khách hàng không phải trả tiền cho món hàng giá nhỏ nhất là 10

→ Khách hàng chỉ phải trả 20+30=50

### Cách 2
- Mua 1 lần 1 sản phẩm giá 10
- Và 1 lần 2 sản phẩm giá 20 & 30, lúc này, khách hàng không phải trả cho món 20

→ Khách hàng phải trả tổng là 10+30=40

### Cách 3
- Mua 3 lần, mỗi lần 1 sản phẩm

→ Khách hàng phải trả tổng là 10+20+30=60

## Input

Đầu vào là một file văn bản tên Input.txt, dòng đầu tiên ghi số CASE là số lượng testcase sẽ có (1 ≤ CASE ≤ 100). Dòng tiếp theo luôn có hai số N, S cách nhau 1 dấu cách. Trên N dòng tiếp theo, mỗi dòng ghi giá của sản phẩm thứ i. File tiếp tục lặp lại cho đến hết số lượng testcase.

## Output

Đầu ra là một file văn bản tên Output.txt có CASE dòng. Mỗi dòng ghi giá trị tiền nhỏ nhất mà người mua cần bỏ ra để mua đủ sản phẩm. Trong trường hợp không có kết quả thì ghi ra giá trị -1

## Ví dụ

### Input.txt
```
2
3 2
10
20
30
1 5
50
```

### Output.txt
```
40
50
```
