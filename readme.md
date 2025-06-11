# BÀI TEST TUYỂN DỤNG GOLANG DEVELOPER - ĐẶNG THANH LÂM 

## ⚙️ Kiến trúc hệ thống

* Hệ thống gồm 3 node chạy trên Docker Compose:

  * `node1`: Leader – nhận TX, propose block, chờ vote
  * `node2`: Follower – sync và vote block
  * `node3`: Follower – sync và vote block
* Mỗi node được mount 1 volume riêng để lưu `chain.json` (data persist)
* Các node kết nối qua `bridge network` tên là `chainnet`


![Sơ đồ hệ thống](./des.png)

## 💪 Docker Compose Setup

```bash
docker-compose up --build
```
* Chạy đã hẹ hẹ hẹ
## ↻ Quy trình hoạt động

1. Gửi giao dịch:

   ```bash
   curl -X POST http://localhost:2201/submit-tx \
     -H "Content-Type: application/json" \
     -d @sendtx.json
   ```

2. Leader tạo block mới:

   ```bash
   curl http://localhost:2201/propose-block
   ```

3. Block được broadcast tới follower → vote → nếu ≥ 2 accept → block được commit.



## 📦 Volume Mapping

| Node  | Volume      | Port |
| ----- | ----------- | ---- |
| node1 | node1\_data | 2201 |
| node2 | node2\_data | 2202 |
| node3 | node3\_data | 2203 |


**Hệ thống có khả năng chống chịu lỗi và tự phục hồi**:

  * Mỗi node lưu `chain.json` trên volume riêng biệt → khi container chết vẫn khôi phục được dữ liệu. 
  * Nếu follower khởi động mà không có dữ liệu → tự động sync lại chain từ leader.
  * Leader và follower giữ kết nối mạng nội bộ (`chainnet-bridge`) nên vẫn hoạt động được nếu tạm mất kết nối ngoài.
  * Log keep-alive định kỳ giúp giám sát sự cố sớm.

* Volume mount (persist chain)

* Auto sync block khi khởi động lại node

* `depends_on` đảm bảo leader lên trước

* Keep-alive log mỗi 10s để giám sát

##
##
##



-- -- -- -- -

## Thông tin ứng viên
#### Đặng Thanh Lâm 
* Zalo/Phone: 0359001647
* Linkedin: tlamabc
* Github: tlamabc
* Mail: dangthanhlam1312@gmail.com
-- -- -- -- -