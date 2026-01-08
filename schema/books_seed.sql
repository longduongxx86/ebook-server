-- Tạo 1 tài khoản quản lý (admin)
INSERT INTO users (id, username, email, password_hash, full_name, phone, address, avatar_url, role, shop_name, provider, email_verified, created_at, updated_at) VALUES
    (1, 'admin', 'admin@ebookstore.com', '$2a$10$YourHashedPasswordHere', 'Nguyễn Văn Quản Lý', '0987654321', '123 Đường ABC, Quận 1, TP.HCM', 'https://example.com/avatar/admin.jpg', 'admin', 'Ebook Store', 'local', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

-- Tạo 10 tài khoản user (customer)
INSERT INTO users (id, username, email, password_hash, full_name, phone, address, avatar_url, role, provider, email_verified, created_at, updated_at) VALUES
(2, 'nguyenvanh', 'nguyenvanh@gmail.com', '$2a$10$HashedPassword1', 'Nguyễn Văn Hải', '0901234567', '456 Lê Lợi, Quận 3, TP.HCM', 'https://example.com/avatar/user1.jpg', 'customer', 'local', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(3, 'tranthimy', 'tranthimy@gmail.com', '$2a$10$HashedPassword2', 'Trần Thị Mỹ', F'0912345678', '789 Nguyễn Huệ, Quận 1, TP.HCM', 'https://example.com/avatar/user2.jpg', 'customer', 'local', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(4, 'levanlong', 'levanlong@yahoo.com', '$2a$10$HashedPassword3', 'Lê Văn Long', '0923456789', '321 Trần Hưng Đạo, Quận 5, TP.HCM', 'https://example.com/avatar/user3.jpg', 'customer', 'local', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(5, 'phamthing', 'phamthing@gmail.com', '$2a$10$HashedPassword4', 'Phạm Thị Ngọc', '0934567890', '654 Cách Mạng Tháng 8, Quận 10, TP.HCM', 'https://example.com/avatar/user4.jpg', 'customer', 'local', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(6, 'hoangminh', 'hoangminh@gmail.com', '$2a$10$HashedPassword5', 'Hoàng Minh Tuấn', '0945678901', '987 Nguyễn Trãi, Quận 1, TP.HCM', 'https://example.com/avatar/user5.jpg', 'customer', 'local', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(7, 'vuthikieu', 'vuthikieu@gmail.com', '$2a$10$HashedPassword6', 'Vũ Thị Kiều', '0956789012', '159 Phạm Văn Đồng, Thủ Đức, TP.HCM', 'https://example.com/avatar/user6.jpg', 'customer', 'local', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(8, 'dangquang', 'dangquang@yahoo.com', '$2a$10$HashedPassword7', 'Đặng Quang Huy', '0967890123', '753 Lý Thường Kiệt, Quận 11, TP.HCM', 'https://example.com/avatar/user7.jpg', 'customer', 'local', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(9, 'buitram', 'buitram@gmail.com', '$2a$10$HashedPassword8', 'Bùi Thị Trâm', '0978901234', '852 Võ Văn Kiệt, Quận 1, TP.HCM', 'https://example.com/avatar/user8.jpg', 'customer', 'local', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(10, 'ngoduc', 'ngoduc@gmail.com', '$2a$10$HashedPassword9', 'Ngô Đức Anh', '0989012345', '951 Điện Biên Phủ, Quận Bình Thạnh, TP.HCM', 'https://example.com/avatar/user9.jpg', 'customer', 'local', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(11, 'dotuan', 'dotuan@gmail.com', '$2a$10$HashedPassword10', 'Đỗ Văn Tuấn', '0990123456', '147 Lê Văn Sỹ, Quận 3, TP.HCM', 'https://example.com/avatar/user10.jpg', 'customer', 'local', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());


-- Tạo 20 danh mục sách
INSERT INTO categories (id, name, slug, description, created_at, updated_at) VALUES
(1, 'Tiểu thuyết', 'tieu-thuyet', 'Các tác phẩm văn học dài, hư cấu với cốt truyện phức tạp', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(2, 'Trinh thám', 'trinh-tham', 'Truyện ly kỳ, hấp dẫn với những vụ án bí ẩn cần được giải quyết', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(3, 'Khoa học viễn tưởng', 'khoa-hoc-vien-tuong', 'Thể loại hư cấu dựa trên khoa học và công nghệ tương lai', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(4, 'Fantasy', 'fantasy', 'Thế giới phép thuật, sinh vật huyền bí và các cuộc phiêu lưu kỳ ảo', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(5, 'Lịch sử', 'lich-su', 'Sách về các sự kiện, nhân vật và thời kỳ lịch sử', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(6, 'Kinh tế', 'kinh-te', 'Sách về tài chính, đầu tư, kinh doanh và quản lý', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(7, 'Self-help', 'self-help', 'Sách phát triển bản thân, kỹ năng sống và tư duy tích cực', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(8, 'Công nghệ thông tin', 'cong-nghe-thong-tin', 'Sách về lập trình, phần mềm, hệ thống máy tính', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(9, 'Y học - Sức khỏe', 'y-hoc-suc-khoe', 'Sách về y học, dinh dưỡng và chăm sóc sức khỏe', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(10, 'Tâm lý học', 'tam-ly-hoc', 'Sách nghiên cứu về hành vi và tâm trí con người', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(11, 'Văn học cổ điển', 'van-hoc-co-dien', 'Các tác phẩm văn học kinh điển từ các tác giả nổi tiếng thế giới', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(12, 'Thơ ca', 'tho-ca', 'Các tập thơ, trường ca và tác phẩm thơ văn', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(13, 'Du lịch', 'du-lich', 'Sách hướng dẫn du lịch, khám phá và trải nghiệm', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(14, 'Nấu ăn', 'nau-an', 'Sách dạy nấu ăn, công thức và kỹ thuật ẩm thực', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(15, 'Thiếu nhi', 'thieu-nhi', 'Sách dành cho trẻ em với nội dung giáo dục và giải trí', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(16, 'Giáo dục', 'giao-duc', 'Sách giáo khoa, tài liệu học tập và phương pháp giảng dạy', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(17, 'Tôn giáo - Tâm linh', 'ton-giao-tam-linh', 'Sách về tôn giáo, triết học và phát triển tâm linh', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(18, 'Nghệ thuật - Thiết kế', 'nghe-thuat-thiet-ke', 'Sách về hội họa, kiến trúc, thiết kế và mỹ thuật', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(19, 'Văn hóa - Xã hội', 'van-hoa-xa-hoi', 'Sách nghiên cứu về văn hóa, phong tục và xã hội', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(20, 'Truyện ngắn', 'truyen-ngan', 'Các tác phẩm văn học ngắn, cô đọng với thông điệp sâu sắc', UNIX_TIMESTAMP(), UNIX_TIMESTAMP());


INSERT INTO books (id, title, author, description, price, cost, stock, slug, image_url, seller_id, isbn, category_id, created_at, updated_at) VALUES
-- Category 1: Tiểu thuyết (5 sách)
(1, 'Những Người Khốn Khổ', 'Victor Hugo', 'Kiệt tác văn học Pháp về cuộc đời Jean Valjean', 150000, 90000, 50, 'nhung-nguoi-khon-kho', 'https://example.com/books/1.jpg', 1, '978-604-1-00001-1', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(2, 'Chiến Tranh Và Hòa Bình', 'Lev Tolstoy', 'Tiểu thuyết sử thi về xã hội Nga thời Napoleon', 180000, 110000, 30, 'chien-tranh-va-hoa-binh', 'https://example.com/books/2.jpg', 1, '978-604-1-00002-2', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(3, 'Bắt Trẻ Đồng Xanh', 'J.D. Salinger', 'Câu chuyện tuổi trẻ nổi loạn của Holden Caulfield', 85000, 50000, 80, 'bat-tre-dong-xanh', 'https://example.com/books/3.jpg', 1, '978-604-1-00003-3', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(4, 'Tuổi Thơ Dữ Dội', 'Phùng Quán', 'Hồi ký về tuổi thơ trong chiến tranh', 95000, 55000, 60, 'tuoi-tho-du-doi', 'https://example.com/books/4.jpg', 1, '978-604-1-00004-4', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(5, 'Đồi Gió Hú', 'Emily Brontë', 'Câu chuyện tình yêu đầy bi kịch trên đồi hoang', 120000, 70000, 40, 'doi-gio-hu', 'https://example.com/books/5.jpg', 1, '978-604-1-00005-5', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Category 2: Trinh thám (5 sách)
(6, 'Án Mạng Trên Sông Nile', 'Agatha Christie', 'Vụ án bí ẩn trên chuyến du thuyền sông Nile', 110000, 65000, 45, 'an-mang-tren-song-nile', 'https://example.com/books/6.jpg', 1, '978-604-1-00006-6', 2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(7, 'Sherlock Holmes - Toàn Tập', 'Arthur Conan Doyle', 'Tuyển tập các vụ án của thám tử lừng danh Sherlock Holmes', 200000, 120000, 25, 'sherlock-holmes-toan-tap', 'https://example.com/books/7.jpg', 1, '978-604-1-00007-7', 2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(8, 'Mật Mã Da Vinci', 'Dan Brown', 'Cuộc truy tìm bí mật qua các kiệt tác nghệ thuật', 135000, 80000, 55, 'mat-ma-da-vinci', 'https://example.com/books/8.jpg', 1, '978-604-1-00008-8', 2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(9, 'Phía Sau Nghi Can X', 'Higashino Keigo', 'Truyện trinh thám Nhật Bản với kết thúc bất ngờ', 99000, 60000, 65, 'phia-sau-nghi-can-x', 'https://example.com/books/9.jpg', 1, '978-604-1-00009-9', 2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(10, 'Biên Niên Ký Sát Nhân', 'Jeffery Deaver', 'Hành trình truy bắt kẻ giết người hàng loạt', 125000, 75000, 35, 'bien-nien-ky-sat-nhan', 'https://example.com/books/10.jpg', 1, '978-604-1-00010-5', 2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Category 3: Khoa học viễn tưởng (5 sách)
(11, 'Dune - Xứ Cát', 'Frank Herbert', 'Thế giới viễn tưởng về hành tinh sa mạc Arrakis', 160000, 95000, 40, 'dune-xu-cat', 'https://example.com/books/11.jpg', 1, '978-604-1-00011-6', 3, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(12, '1984', 'George Orwell', 'Xã hội dystopia với sự kiểm soát toàn diện', 105000, 60000, 70, '1984', 'https://example.com/books/12.jpg', 1, '978-604-1-00012-7', 3, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(13, 'Người Máy Có Mơ Về Cừu Điện Không?', 'Philip K. Dick', 'Câu hỏi về ý thức và bản chất con người trong tương lai', 115000, 70000, 50, 'nguoi-may-co-mo-ve-cuu-dien-khong', 'https://example.com/books/13.jpg', 1, '978-604-1-00013-8', 3, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(14, 'Cuộc Chiến Giữa Các Thế Giới', 'H.G. Wells', 'Người ngoài hành tinh xâm lược Trái Đất', 98000, 58000, 60, 'cuoc-chien-giua-cac-the-gioi', 'https://example.com/books/14.jpg', 1, '978-604-1-00014-9', 3, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(15, 'Foundation - Nền Móng', 'Isaac Asimov', 'Vũ trụ tương lai với psychohistory và Đế chế Thiên Hà', 175000, 105000, 30, 'foundation-nen-mong', 'https://example.com/books/15.jpg', 1, '978-604-1-00015-0', 3, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Category 4: Fantasy (5 sách)
(16, 'Chúa Tể Những Chiếc Nhẫn', 'J.R.R. Tolkien', 'Hành trình tiêu diệt Chiếc Nhẫn Chúa', 220000, 130000, 25, 'chua-te-nhung-chiec-nhan', 'https://example.com/books/16.jpg', 1, '978-604-1-00016-1', 4, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(17, 'Harry Potter Và Hòn Đá Phù Thủy', 'J.K. Rowling', 'Câu chuyện về cậu bé phù thủy Harry Potter', 140000, 85000, 75, 'harry-potter-va-hon-da-phu-thuy', 'https://example.com/books/17.jpg', 1, '978-604-1-00017-2', 4, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(18, 'Trò Chơi Vương Quyền', 'George R.R. Martin', 'Cuộc chiến ngai sắt tại Westeros', 195000, 115000, 35, 'tro-choi-vuong-quyen', 'https://example.com/books/18.jpg', 1, '978-604-1-00018-3', 4, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(19, 'Tên Gió', 'Patrick Rothfuss', 'Câu chuyện về Kvothe, nhạc sĩ, sinh viên, kẻ trộm', 165000, 98000, 45, 'ten-gio', 'https://example.com/books/19.jpg', 1, '978-604-1-00019-4', 4, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(20, 'Thần Thoại Bắc Âu', 'Neil Gaiman', 'Câu chuyện về các vị thần Bắc Âu', 125000, 75000, 55, 'than-thoai-bac-au', 'https://example.com/books/20.jpg', 1, '978-604-1-00020-7', 4, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Category 5: Lịch sử (5 sách)
(21, 'Sapiens: Lược Sử Loài Người', 'Yuval Noah Harari', 'Lịch sử tiến hóa của loài người từ thời đồ đá', 155000, 92000, 50, 'sapiens-luoc-su-loai-nguoi', 'https://example.com/books/21.jpg', 1, '978-604-1-00021-8', 5, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(22, 'Lịch Sử Việt Nam', 'Phan Huy Lê', 'Toàn cảnh lịch sử Việt Nam từ thời dựng nước', 185000, 110000, 40, 'lich-su-viet-nam', 'https://example.com/books/22.jpg', 1, '978-604-1-00022-9', 5, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(23, 'Đại Việt Sử Ký Toàn Thư', 'Ngô Sĩ Liên', 'Bộ quốc sử chính thống của Việt Nam', 250000, 150000, 20, 'dai-viet-su-ky-toan-thu', 'https://example.com/books/23.jpg', 1, '978-604-1-00023-0', 5, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(24, 'Lịch Sử Thế Giới', 'E.H. Gombrich', 'Lịch sử thế giới viết cho giới trẻ', 135000, 80000, 60, 'lich-su-the-gioi', 'https://example.com/books/24.jpg', 1, '978-604-1-00024-1', 5, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(25, 'Napoleon Bonaparte', 'Andrew Roberts', 'Tiểu sử toàn diện về Napoleon', 175000, 105000, 30, 'napoleon-bonaparte', 'https://example.com/books/25.jpg', 1, '978-604-1-00025-2', 5, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Category 6: Kinh tế (5 sách)
(26, 'Kinh Tế Học Vui Vẻ', 'Steven Levitt', 'Kinh tế học ứng dụng trong đời sống', 120000, 72000, 55, 'kinh-te-hoc-vui-ve', 'https://example.com/books/26.jpg', 1, '978-604-1-00026-3', 6, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(27, 'Cha Giàu Cha Nghèo', 'Robert Kiyosaki', 'Tư duy tài chính để đạt tự do tài chính', 95000, 57000, 80, 'cha-giau-cha-ngheo', 'https://example.com/books/27.jpg', 1, '978-604-1-00027-4', 6, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(28, 'Nhà Đầu Tư Thông Minh', 'Benjamin Graham', 'Cẩm nang đầu tư giá trị kinh điển', 145000, 87000, 45, 'nha-dau-tu-thong-minh', 'https://example.com/books/28.jpg', 1, '978-604-1-00028-5', 6, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(29, 'Tư Bản Thế Kỷ 21', 'Thomas Piketty', 'Phân tích bất bình đẳng trong chủ nghĩa tư bản', 195000, 117000, 35, 'tu-ban-the-ky-21', 'https://example.com/books/29.jpg', 1, '978-604-1-00029-6', 6, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(30, 'Bí Mật Tư Duy Triệu Phú', 'T. Harv Eker', 'Công thức tư duy của những người giàu có', 110000, 66000, 65, 'bi-mat-tu-duy-trieu-phu', 'https://example.com/books/30.jpg', 1, '978-604-1-00030-2', 6, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Category 7: Self-help (5 sách)
(31, 'Đắc Nhân Tâm', 'Dale Carnegie', 'Nghệ thuật thu phục lòng người', 85000, 51000, 90, 'dac-nhan-tam', 'https://example.com/books/31.jpg', 1, '978-604-1-00031-3', 7, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(32, '7 Thói Quen Hiệu Quả', 'Stephen Covey', '7 thói quen của người thành đạt', 135000, 81000, 60, '7-thoi-quen-hieu-qua', 'https://example.com/books/32.jpg', 1, '978-604-1-00032-4', 7, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(33, 'Người Giỏi Không Bởi Học Nhiều', 'Alpha Books', 'Phương pháp học tập thông minh', 99000, 59400, 70, 'nguoi-gioi-khong-boi-hoc-nhieu', 'https://example.com/books/33.jpg', 1, '978-604-1-00033-5', 7, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(34, 'Đời Ngắn Đừng Ngủ Dài', 'Robin Sharma', 'Bài học về tận dụng thời gian', 105000, 63000, 65, 'doi-ngan-dung-ngu-dai', 'https://example.com/books/34.jpg', 1, '978-604-1-00034-6', 7, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(35, 'Atomic Habits', 'James Clear', 'Xây dựng thói quen nhỏ để đạt thành công lớn', 145000, 87000, 50, 'atomic-habits', 'https://example.com/books/35.jpg', 1, '978-604-1-00035-7', 7, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Category 8: Công nghệ thông tin (5 sách)
(36, 'Clean Code', 'Robert C. Martin', 'Nghệ thuật viết code sạch', 165000, 99000, 40, 'clean-code', 'https://example.com/books/36.jpg', 1, '978-604-1-00036-8', 8, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(37, 'Design Patterns', 'Gang of Four', 'Các mẫu thiết kế phần mềm', 185000, 111000, 35, 'design-patterns', 'https://example.com/books/37.jpg', 1, '978-604-1-00037-9', 8, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(38, 'Python Crash Course', 'Eric Matthes', 'Học Python từ cơ bản đến nâng cao', 155000, 93000, 45, 'python-crash-course', 'https://example.com/books/38.jpg', 1, '978-604-1-00038-0', 8, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(39, 'JavaScript: The Good Parts', 'Douglas Crockford', 'Tinh hoa của JavaScript', 140000, 84000, 50, 'javascript-the-good-parts', 'https://example.com/books/39.jpg', 1, '978-604-1-00039-1', 8, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(40, 'AI - Trí Tuệ Nhân Tạo', 'Stuart Russell', 'Giới thiệu về trí tuệ nhân tạo', 195000, 117000, 30, 'ai-tri-tue-nhan-tao', 'https://example.com/books/40.jpg', 1, '978-604-1-00040-7', 8, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Category 9: Y học - Sức khỏe (5 sách)
(41, 'Nhân Tố Enzyme', 'Hiromi Shinya', 'Phương pháp sống khỏe bằng enzyme', 115000, 69000, 55, 'nhan-to-enzyme', 'https://example.com/books/41.jpg', 1, '978-604-1-00041-8', 9, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(42, 'Bí Mật Dinh Dưỡng', 'Colin Campbell', 'Mối liên hệ giữa dinh dưỡng và sức khỏe', 135000, 81000, 45, 'bi-mat-dinh-duong', 'https://example.com/books/42.jpg', 1, '978-604-1-00042-9', 9, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(43, 'Giải Phẫu Học', 'Frank H. Netter', 'Atlas giải phẫu người chi tiết', 225000, 135000, 25, 'giai-phau-hoc', 'https://example.com/books/43.jpg', 1, '978-604-1-00043-0', 9, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(44, 'Y Học Cổ Truyền', 'Đỗ Tất Lợi', 'Cây thuốc và vị thuốc Việt Nam', 145000, 87000, 40, 'y-hoc-co-truyen', 'https://example.com/books/44.jpg', 1, '978-604-1-00044-1', 9, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(45, 'Mindfulness - Sống Trọn Từng Khoảnh Khắc', 'Jon Kabat-Zinn', 'Thực hành chánh niệm cho sức khỏe tinh thần', 125000, 75000, 50, 'mindfulness-song-tron-tung-khoanh-khac', 'https://example.com/books/45.jpg', 1, '978-604-1-00045-2', 9, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Category 10: Tâm lý học (5 sách)
(46, 'Tâm Lý Học Đám Đông', 'Gustave Le Bon', 'Phân tích hành vi của đám đông', 110000, 66000, 60, 'tam-ly-hoc-dam-dong', 'https://example.com/books/46.jpg', 1, '978-604-1-00046-3', 10, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(47, 'Phi Lý Trí', 'Dan Ariely', 'Hành vi phi lý trí trong quyết định', 130000, 78000, 50, 'phi-ly-tri', 'https://example.com/books/47.jpg', 1, '978-604-1-00047-4', 10, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(48, 'Tâm Lý Học Tội Phạm', 'John E. Douglas', 'Phân tích tâm lý kẻ phạm tội', 150000, 90000, 40, 'tam-ly-hoc-toi-pham', 'https://example.com/books/48.jpg', 1, '978-604-1-00048-5', 10, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(49, 'Đọc Vị Bất Kỳ Ai', 'David J. Lieberman', 'Nghệ thuật thấu hiểu người khác', 105000, 63000, 65, 'doc-vi-bat-ky-ai', 'https://example.com/books/49.jpg', 1, '978-604-1-00049-6', 10, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(50, 'Tâm Lý Học Hành Vi', 'B.F. Skinner', 'Lý thuyết về hành vi và điều kiện hóa', 140000, 84000, 45, 'tam-ly-hoc-hanh-vi', 'https://example.com/books/50.jpg', 1, '978-604-1-00050-2', 10, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
-- Tiếp tục tạo 50 cuốn sách (category 11-20)
INSERT INTO books (id, title, author, description, price, cost, stock, slug, image_url, seller_id, isbn, category_id, created_at, updated_at) VALUES
-- Category 11: Văn học cổ điển (5 sách)
(51, 'Chiếc Lá Cuối Cùng', 'O. Henry', 'Truyện ngắn cảm động về tình yêu cuộc sống', 75000, 45000, 85, 'chiec-la-cuoi-cung', 'https://example.com/books/51.jpg', 1, '978-604-1-00051-3', 11, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(52, 'Ông Già Và Biển Cả', 'Ernest Hemingway', 'Câu chuyện về ý chí con người trước thiên nhiên', 95000, 57000, 60, 'ong-gia-va-bien-ca', 'https://example.com/books/52.jpg', 1, '978-604-1-00052-4', 11, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(53, 'Kiêu Hãnh Và Định Kiến', 'Jane Austen', 'Tình yêu và địa vị xã hội trong thế kỷ 19', 120000, 72000, 55, 'kieu-hanh-va-dinh-kien', 'https://example.com/books/53.jpg', 1, '978-604-1-00053-5', 11, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(54, 'Giết Con Chim Nhại', 'Harper Lee', 'Câu chuyện về phân biệt chủng tộc và công lý', 135000, 81000, 50, 'giet-con-chim-nhai', 'https://example.com/books/54.jpg', 1, '978-604-1-00054-6', 11, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(55, 'Bá Tước Monte Cristo', 'Alexandre Dumas', 'Hành trình trả thù của Edmond Dantès', 185000, 111000, 40, 'ba-tuoc-monte-cristo', 'https://example.com/books/55.jpg', 1, '978-604-1-00055-7', 11, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Category 12: Thơ ca (5 sách)
(56, 'Truyện Kiều', 'Nguyễn Du', 'Kiệt tác thơ Nôm của đại thi hào Nguyễn Du', 110000, 66000, 65, 'truyen-kieu', 'https://example.com/books/56.jpg', 1, '978-604-1-00056-8', 12, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(57, 'Nhật Ký Trong Tù', 'Hồ Chí Minh', 'Tập thơ viết trong thời gian bị giam cầm', 85000, 51000, 75, 'nhat-ky-trong-tu', 'https://example.com/books/57.jpg', 1, '978-604-1-00057-9', 12, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(58, 'Lời Và Kỷ Vật', 'Xuân Diệu', 'Tuyển tập thơ tình của ông hoàng thơ tình Việt Nam', 95000, 57000, 70, 'loi-va-ky-vat', 'https://example.com/books/58.jpg', 1, '978-604-1-00058-0', 12, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(59, 'Tuyển Tập Thơ Hàn Mặc Tử', 'Hàn Mặc Tử', 'Thơ của thi sĩ tài hoa bạc mệnh', 105000, 63000, 60, 'tuyen-tap-tho-han-mac-tu', 'https://example.com/books/59.jpg', 1, '978-604-1-00059-1', 12, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(60, 'Thơ Đường', 'Nhiều tác giả', 'Tuyển tập thơ Đường nổi tiếng Trung Quốc', 125000, 75000, 50, 'tho-duong', 'https://example.com/books/60.jpg', 1, '978-604-1-00060-7', 12, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Category 13: Du lịch (5 sách)
(61, 'Hành Trình Về Phương Đông', 'Baird T. Spalding', 'Chuyến du hành khám phá tâm linh phương Đông', 130000, 78000, 55, 'hanh-trinh-ve-phuong-dong', 'https://example.com/books/61.jpg', 1, '978-604-1-00061-8', 13, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(62, 'Việt Nam - Đi Và Trải Nghiệm', 'Lonely Planet', 'Cẩm nang du lịch Việt Nam đầy đủ', 145000, 87000, 45, 'viet-nam-di-va-trai-nghiem', 'https://example.com/books/62.jpg', 1, '978-604-1-00062-9', 13, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(63, 'Châu Âu Tự Túc', 'Phan Hồng Thủy', 'Kinh nghiệm du lịch châu Âu tiết kiệm', 115000, 69000, 60, 'chau-au-tu-tuc', 'https://example.com/books/63.jpg', 1, '978-604-1-00063-0', 13, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(64, 'Nhật Bản - Văn Hóa Và Con Người', 'Alex Kerr', 'Khám phá văn hóa độc đáo của Nhật Bản', 135000, 81000, 50, 'nhat-ban-van-hoa-va-con-nguoi', 'https://example.com/books/64.jpg', 1, '978-604-1-00064-1', 13, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(65, 'Himalaya - Nóc Nhà Thế Giới', 'Heinrich Harrer', 'Hành trình chinh phục dãy Himalaya', 155000, 93000, 40, 'himalaya-noc-nha-the-gioi', 'https://example.com/books/65.jpg', 1, '978-604-1-00065-2', 13, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Category 14: Nấu ăn (5 sách)
(66, 'Nghệ Thuật Ẩm Thực Việt', 'Nguyễn Dzoãn Cẩm Vân', 'Công thức nấu ăn truyền thống Việt Nam', 125000, 75000, 55, 'nghe-thuat-am-thuc-viet', 'https://example.com/books/66.jpg', 1, '978-604-1-00066-3', 14, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(67, 'Baking - Từ Cơ Bản Đến Nâng Cao', 'Paul Hollywood', 'Kỹ thuật làm bánh chuyên nghiệp', 145000, 87000, 45, 'baking-tu-co-ban-den-nang-cao', 'https://example.com/books/67.jpg', 1, '978-604-1-00067-4', 14, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(68, 'Món Ăn Đường Phố Châu Á', 'David Thompson', 'Công thức các món ăn đường phố nổi tiếng', 115000, 69000, 60, 'mon-an-duong-pho-chau-a', 'https://example.com/books/68.jpg', 1, '978-604-1-00068-5', 14, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(69, 'Sushi - Nghệ Thuật Ẩm Thực Nhật', 'Jiro Ono', 'Kỹ thuật làm sushi chuyên nghiệp', 165000, 99000, 35, 'sushi-nghe-thuat-am-thuc-nhat', 'https://example.com/books/69.jpg', 1, '978-604-1-00069-6', 14, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(70, 'Chay Tịnh - Món Ăn Thanh Nhẹ', 'Minh Chay', 'Công thức nấu ăn chay bổ dưỡng', 105000, 63000, 65, 'chay-tinh-mon-an-thanh-nhe', 'https://example.com/books/70.jpg', 1, '978-604-1-00070-2', 14, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Category 15: Thiếu nhi (5 sách)
(71, 'Dế Mèn Phiêu Lưu Ký', 'Tô Hoài', 'Cuộc phiêu lưu của chú dế mèn', 85000, 51000, 80, 'de-men-phieu-luu-ky', 'https://example.com/books/71.jpg', 1, '978-604-1-00071-3', 15, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(72, 'Harry Potter Và Phòng Chứa Bí Mật', 'J.K. Rowling', 'Phần 2 của series Harry Potter', 140000, 84000, 70, 'harry-potter-va-phong-chua-bi-mat', 'https://example.com/books/72.jpg', 1, '978-604-1-00072-4', 15, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(73, 'Hoàng Tử Bé', 'Antoine de Saint-Exupéry', 'Câu chuyện triết lý sâu sắc cho mọi lứa tuổi', 95000, 57000, 75, 'hoang-tu-be', 'https://example.com/books/73.jpg', 1, '978-604-1-00073-5', 15, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(74, 'Truyện Cổ Grimm', 'Anh em Grimm', 'Tuyển tập truyện cổ tích nổi tiếng', 115000, 69000, 65, 'truyen-co-grimm', 'https://example.com/books/74.jpg', 1, '978-604-1-00074-6', 15, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(75, 'Alice Ở Xứ Sở Thần Tiên', 'Lewis Carroll', 'Cuộc phiêu lưu kỳ diệu của Alice', 105000, 63000, 70, 'alice-o-xu-so-than-tien', 'https://example.com/books/75.jpg', 1, '978-604-1-00075-7', 15, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Category 16: Giáo dục (5 sách)
(76, 'Phương Pháp Montessori', 'Maria Montessori', 'Phương pháp giáo dục sớm cho trẻ', 135000, 81000, 55, 'phuong-phap-montessori', 'https://example.com/books/76.jpg', 1, '978-604-1-00076-8', 16, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(77, 'Dạy Con Làm Giàu', 'Robert Kiyosaki', 'Giáo dục tài chính cho trẻ em', 115000, 69000, 60, 'day-con-lam-giau', 'https://example.com/books/77.jpg', 1, '978-604-1-00077-9', 16, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(78, 'Tư Duy Phản Biện', 'Richard Paul', 'Rèn luyện kỹ năng tư duy phản biện', 125000, 75000, 50, 'tu-duy-phan-bien', 'https://example.com/books/78.jpg', 1, '978-604-1-00078-0', 16, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(79, 'Nghệ Thuật Dạy Học', 'Ken Bain', 'Phương pháp giảng dạy hiệu quả', 145000, 87000, 45, 'nghe-thuat-day-hoc', 'https://example.com/books/79.jpg', 1, '978-604-1-00079-1', 16, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(80, 'Học Cách Học', 'Barbara Oakley', 'Khoa học về học tập hiệu quả', 130000, 78000, 55, 'hoc-cach-hoc', 'https://example.com/books/80.jpg', 1, '978-604-1-00080-7', 16, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Category 17: Tôn giáo - Tâm linh (5 sách)
(81, 'Kinh Thánh', 'Nhiều tác giả', 'Bộ sách thánh của Kitô giáo', 175000, 105000, 40, 'kinh-thanh', 'https://example.com/books/81.jpg', 1, '978-604-1-00081-8', 17, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(82, 'Kinh Phật Cho Người Tại Gia', 'Thích Nhất Hạnh', 'Lời Phật dạy áp dụng trong đời sống', 125000, 75000, 60, 'kinh-phat-cho-nguoi-tai-gia', 'https://example.com/books/82.jpg', 1, '978-604-1-00082-9', 17, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(83, 'Đạo Đức Kinh', 'Lão Tử', 'Triết lý Đạo gia Trung Hoa', 110000, 66000, 65, 'dao-duc-kinh', 'https://example.com/books/83.jpg', 1, '978-604-1-00083-0', 17, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(84, 'Thiền Trong Đời Sống', 'S.N. Goenka', 'Thực hành thiền Vipassana', 135000, 81000, 50, 'thien-trong-doi-song', 'https://example.com/books/84.jpg', 1, '978-604-1-00084-1', 17, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(85, 'Hành Trình Tâm Linh', 'Eckhart Tolle', 'Khám phá bản ngã và ý thức', 145000, 87000, 45, 'hanh-trinh-tam-linh', 'https://example.com/books/85.jpg', 1, '978-604-1-00085-2', 17, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Category 18: Nghệ thuật - Thiết kế (5 sách)
(86, 'Nghệ Thuật Xem Tranh', 'John Berger', 'Hướng dẫn thưởng thức và phân tích tranh', 155000, 93000, 45, 'nghe-thuat-xem-tranh', 'https://example.com/books/86.jpg', 1, '978-604-1-00086-3', 18, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(87, 'Nguyên Lý Thiết Kế', 'William Lidwell', 'Các nguyên tắc cơ bản trong thiết kế', 165000, 99000, 40, 'nguyen-ly-thiet-ke', 'https://example.com/books/87.jpg', 1, '978-604-1-00087-4', 18, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(88, 'Lịch Sử Mỹ Thuật', 'E.H. Gombrich', 'Tổng quan về lịch sử mỹ thuật thế giới', 185000, 111000, 35, 'lich-su-my-thuat', 'https://example.com/books/88.jpg', 1, '978-604-1-00088-5', 18, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(89, 'Nhiếp Ảnh Cơ Bản', 'Bryan Peterson', 'Kỹ thuật nhiếp ảnh cho người mới bắt đầu', 135000, 81000, 55, 'nhiep-anh-co-ban', 'https://example.com/books/89.jpg', 1, '978-604-1-00089-6', 18, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(90, 'Kiến Trúc Hiện Đại', 'Le Corbusier', 'Lý thuyết và thực hành kiến trúc hiện đại', 195000, 117000, 30, 'kien-truc-hien-dai', 'https://example.com/books/90.jpg', 1, '978-604-1-00090-2', 18, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Category 19: Văn hóa - Xã hội (5 sách)
(91, 'Văn Minh Phương Tây', 'Will Durant', 'Lịch sử văn minh phương Tây', 175000, 105000, 40, 'van-minh-phuong-tay', 'https://example.com/books/91.jpg', 1, '978-604-1-00091-3', 19, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(92, 'Xã Hội Học Đại Cương', 'Anthony Giddens', 'Giới thiệu về xã hội học', 145000, 87000, 50, 'xa-hoi-hoc-dai-cuong', 'https://example.com/books/92.jpg', 1, '978-604-1-00092-4', 19, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(93, 'Văn Hóa Việt Nam', 'Trần Ngọc Thêm', 'Đặc trưng văn hóa Việt Nam', 155000, 93000, 45, 'van-hoa-viet-nam', 'https://example.com/books/93.jpg', 1, '978-604-1-00093-5', 19, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(94, 'Chủ Nghĩa Tư Bản Và Tự Do', 'Milton Friedman', 'Kinh tế học và tự do cá nhân', 165000, 99000, 40, 'chu-nghia-tu-ban-va-tu-do', 'https://example.com/books/94.jpg', 1, '978-604-1-00094-6', 19, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(95, 'Toàn Cầu Hóa Và Ảnh Hưởng', 'Joseph Stiglitz', 'Phân tích về toàn cầu hóa', 150000, 90000, 45, 'toan-cau-hoa-va-anh-huong', 'https://example.com/books/95.jpg', 1, '978-604-1-00095-7', 19, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Category 20: Truyện ngắn (5 sách)
(96, 'Truyện Ngắn Nam Cao', 'Nam Cao', 'Tuyển tập truyện ngắn của nhà văn Nam Cao', 115000, 69000, 60, 'truyen-ngan-nam-cao', 'https://example.com/books/96.jpg', 1, '978-604-1-00096-8', 20, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(97, 'Truyện Ngắn Thạch Lam', 'Thạch Lam', 'Tuyển tập truyện ngắn đặc sắc của Thạch Lam', 105000, 63000, 65, 'truyen-ngan-thach-lam', 'https://example.com/books/97.jpg', 1, '978-604-1-00097-9', 20, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(98, 'Tuyển Tập Truyện Ngắn Thế Giới', 'Nhiều tác giả', 'Tuyển tập truyện ngắn hay nhất thế giới', 135000, 81000, 50, 'tuyen-tap-truyen-ngan-the-gioi', 'https://example.com/books/98.jpg', 1, '978-604-1-00098-0', 20, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(99, 'Truyện Ngắn Chekhov', 'Anton Chekhov', 'Tuyển tập truyện ngắn của đại văn hào Nga', 125000, 75000, 55, 'truyen-ngan-chekhov', 'https://example.com/books/99.jpg', 1, '978-604-1-00099-1', 20, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(100, 'Truyện Ngắn O. Henry', 'O. Henry', 'Truyện ngắn với kết thúc bất ngờ đặc trưng', 110000, 66000, 60, 'truyen-ngan-o-henry', 'https://example.com/books/100.jpg', 1, '978-604-1-00100-4', 20, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

-- Tạo 100 review mẫu (mỗi sách có khoảng 1-3 review từ các user khác nhau)
INSERT INTO reviews (id, book_id, user_id, rating, comment, approved, created_at, updated_at) VALUES
-- Reviews cho sách 1-10
(1, 1, 2, 5, 'Tuyệt vời! Một kiệt tác văn học không thể bỏ qua. Câu chuyện cảm động sâu sắc.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(2, 1, 3, 4, 'Sách hay nhưng hơi dài, cần kiên nhẫn để đọc hết.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(3, 2, 4, 5, 'Một tác phẩm đồ sộ về cả nội dung và tầm vóc. Đáng đọc!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(4, 3, 5, 5, 'Rất thích cách viết chân thực và gần gũi. Phù hợp với giới trẻ.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(5, 4, 6, 4, 'Hồi ký xúc động về một thời đã qua. Đọc mà thấy thương những số phận.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(6, 5, 7, 5, 'Tình yêu đẹp nhưng đầy bi kịch. Câu chuyện ám ảnh mãi không quên.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(7, 6, 8, 4, 'Vụ án ly kỳ, hấp dẫn. Agatha Christie quả là bậc thầy trinh thám.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(8, 7, 9, 5, 'Tuyển tập đầy đủ nhất về Sherlock Holmes. Rất đáng sưu tầm.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(9, 8, 10, 5, 'Cuốn sách khiến mình say mê từ đầu đến cuối. Kết thúc bất ngờ!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(10, 9, 2, 4, 'Truyện trinh thám Nhật Bản rất khác biệt. Thích cách xây dựng tâm lý nhân vật.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Reviews cho sách 11-20
(11, 11, 3, 5, 'Thế giới viễn tưởng được xây dựng công phu. Tuyệt vời!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(12, 12, 4, 5, 'Đáng sợ nhưng chân thực. Những gì Orwell mô tả đang dần thành hiện thực.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(13, 13, 5, 4, 'Câu hỏi triết lý sâu sắc về ý thức và công nghệ.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(14, 14, 6, 3, 'Tác phẩm kinh điển nhưng có phần hơi cũ so với hiện tại.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(15, 15, 7, 5, 'Asimov thật thiên tài! Tầm nhìn vượt thời đại.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(16, 16, 8, 5, 'Bộ sách fantasy hay nhất mọi thời đại. Đã đọc đi đọc lại nhiều lần.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(17, 17, 9, 5, 'Tuổi thơ của bao thế hệ. Harry Potter mãi trong tim!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(18, 18, 10, 4, 'Cốt truyện phức tạp, nhiều nhân vật. Cần tập trung khi đọc.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(19, 19, 2, 5, 'Văn phong đẹp như thơ. Câu chuyện về Kvothe thật phi thường.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(20, 20, 3, 4, 'Tuyển tập thần thoại hấp dẫn. Giúp hiểu thêm về văn hóa Bắc Âu.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Reviews cho sách 21-30
(21, 21, 4, 5, 'Thay đổi hoàn toàn cách nhìn về lịch sử loài người. Xuất sắc!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(22, 22, 5, 4, 'Tài liệu tham khảo tốt cho người nghiên cứu lịch sử Việt Nam.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(23, 23, 6, 5, 'Bảo vật quốc gia. Nên có trong tủ sách mỗi gia đình Việt.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(24, 24, 7, 4, 'Viết cho giới trẻ nên dễ hiểu, không khô khan.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(25, 25, 8, 5, 'Tiểu sử toàn diện về một thiên tài quân sự.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(26, 26, 9, 4, 'Kinh tế học trở nên thú vị hơn bao giờ hết.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(27, 27, 10, 5, 'Thay đổi tư duy tài chính hoàn toàn. Rất đáng đọc!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(28, 28, 2, 4, 'Sách kinh điển về đầu tư giá trị. Hơi khó với người mới.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(29, 29, 3, 5, 'Phân tích sâu sắc về bất bình đẳng trong xã hội hiện đại.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(30, 30, 4, 3, 'Nội dung khá chung chung, không có nhiều cái mới.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Reviews cho sách 31-40
(31, 31, 5, 5, 'Sách self-help kinh điển. Ai cũng nên đọc ít nhất một lần.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(32, 32, 6, 4, '7 thói quen thực sự hữu ích cho công việc và cuộc sống.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(33, 33, 7, 3, 'Nội dung ổn nhưng không có gì đặc biệt so với sách cùng loại.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(34, 34, 8, 4, 'Truyền cảm hứng sống tích cực. Đọc xong thấy có động lực hơn.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(35, 35, 9, 5, 'Cách tiếp cận khoa học về thói quen. Rất thực tế và hiệu quả.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(36, 36, 10, 5, 'Sách gối đầu giường của mọi lập trình viên. Phải đọc!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(37, 37, 2, 4, 'Tài liệu tham khảo quan trọng cho thiết kế phần mềm.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(38, 38, 3, 5, 'Học Python tốt nhất cho người mới bắt đầu. Bài tập thực hành phong phú.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(39, 39, 4, 4, 'Tập trung vào những phần tốt nhất của JavaScript.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(40, 40, 5, 5, 'Giới thiệu tổng quan về AI dễ hiểu cho người không chuyên.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Reviews cho sách 41-50
(41, 41, 6, 4, 'Phương pháp sống khỏe thú vị. Đáng để thử nghiệm.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(42, 42, 7, 5, 'Nghiên cứu công phu về dinh dưỡng. Thay đổi hoàn toàn chế độ ăn của tôi.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(43, 43, 8, 5, 'Atlas giải phẫu chi tiết nhất. Rất hữu ích cho sinh viên y.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(44, 44, 9, 4, 'Tài liệu quý về y học cổ truyền Việt Nam.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(45, 45, 10, 5, 'Thực hành chánh niệm đã giúp tôi giảm stress đáng kể.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(46, 46, 2, 4, 'Phân tích thú vị về tâm lý đám đông. Vẫn còn giá trị sau hơn 100 năm.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(47, 47, 3, 5, 'Mở mắt về những quyết định phi lý trí trong cuộc sống.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(48, 48, 4, 4, 'Góc nhìn độc đáo về tâm lý tội phạm. Hấp dẫn như phim trinh thám.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(49, 49, 5, 3, 'Kỹ năng hữu ích nhưng đôi khi áp dụng không đúng ngữ cảnh.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(50, 50, 6, 4, 'Lý thuyết hành vi được trình bày khoa học và dễ hiểu.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Reviews cho sách 51-60
(51, 51, 7, 5, 'Truyện ngắn kinh điển. Kết thúc làm tôi rơi nước mắt.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(52, 52, 8, 5, 'Câu chuyện về ý chí con người. Hemingway viết quá xuất sắc!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(53, 53, 9, 4, 'Tình yêu thời Victoria lãng mạn nhưng cũng đầy định kiến.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(54, 54, 10, 5, 'Sách hay nhất về công lý và lòng nhân ái. Nên đọc trong trường học.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(55, 55, 2, 5, 'Cuộc trả thù hoàn hảo nhất trong lịch sử văn học.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(56, 56, 3, 5, 'Kiệt tác của dân tộc. Mỗi lần đọc lại thấy mới.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(57, 57, 4, 5, 'Thơ hay, ý chí kiên cường của người chiến sĩ cách mạng.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(58, 58, 5, 4, 'Thơ tình Xuân Diệu mãi trẻ trung và say đắm.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(59, 59, 6, 5, 'Thơ Hàn Mặc Tử - một tài năng bạc mệnh.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(60, 60, 7, 4, 'Tuyển tập thơ Đường đa dạng. Thích phần chú thích chi tiết.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Reviews cho sách 61-70
(61, 61, 8, 5, 'Hành trình khám phá tâm linh đầy cảm hứng.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(62, 62, 9, 4, 'Cẩm nang du lịch Việt Nam đầy đủ, cập nhật.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(63, 63, 10, 5, 'Tiết kiệm được nhiều chi phí nhờ cuốn sách này. Rất thực tế!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(64, 64, 2, 4, 'Hiểu sâu hơn về văn hóa Nhật Bản trước khi đi du lịch.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(65, 65, 3, 5, 'Cuộc phiêu lưu ngoạn mục đến nóc nhà thế giới.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(66, 66, 4, 4, 'Công thức nấu ăn Việt đa dạng, dễ thực hiện.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(67, 67, 5, 5, 'Từ người mới đã có thể làm được bánh ngon nhờ sách này.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(68, 68, 6, 4, 'Món ăn đường phố châu Á hấp dẫn. Thích phần ảnh minh họa.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(69, 69, 7, 5, 'Nghệ thuật sushi được trình bày tinh tế. Đẹp như một cuốn artbook.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(70, 70, 8, 4, 'Món chay thanh đạm nhưng không kém phần hấp dẫn.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Reviews cho sách 71-80
(71, 71, 9, 5, 'Tuổi thơ ùa về với Dế Mèn. Sách thiếu nhi hay nhất Việt Nam!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(72, 72, 10, 5, 'Phần tiếp theo không làm thất vọng. Harry Potter càng đọc càng hay.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(73, 73, 2, 5, 'Cuốn sách cho mọi lứa tuổi. Mỗi lần đọc có một cảm nhận khác.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(74, 74, 3, 4, 'Truyện cổ Grimm nguyên bản, có phần hơi kinh dị so với phiên bản Disney.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(75, 75, 4, 5, 'Thế giới thần tiên kỳ diệu. Ước gì được đến xứ sở đó một lần.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(76, 76, 5, 4, 'Phương pháp giáo dục tiến bộ. Áp dụng được nhiều cho con nhỏ.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(77, 77, 6, 3, 'Nội dung lặp lại nhiều so với các sách khác của tác giả.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(78, 78, 7, 5, 'Kỹ năng cần thiết trong thời đại thông tin. Rèn luyện được tư duy phản biện.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(79, 79, 8, 4, 'Giáo viên nên đọc để cải thiện phương pháp giảng dạy.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(80, 80, 9, 5, 'Khoa học về học tập được trình bày dễ hiểu. Hiệu quả thực sự!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Reviews cho sách 81-90
(81, 81, 10, 5, 'Kinh thánh là nền tảng của văn minh phương Tây. Nên đọc để hiểu văn hóa.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(82, 82, 2, 4, 'Lời Phật dạy ứng dụng thiết thực trong đời sống hiện đại.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(83, 83, 3, 5, 'Triết lý Đạo gia sâu sắc. Càng đọc càng thấm.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(84, 84, 4, 4, 'Thiền giúp tâm trí thanh tịnh. Sách hướng dẫn chi tiết.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(85, 85, 5, 5, 'Hành trình khám phá bản thân đầy ý nghĩa.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(86, 86, 6, 4, 'Học được cách thưởng thức nghệ thuật đúng cách.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(87, 87, 7, 5, 'Nguyên tắc thiết kế cơ bản nhưng cực kỳ quan trọng.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(88, 88, 8, 5, 'Lịch sử mỹ thuật toàn diện. Ảnh minh họa chất lượng cao.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(89, 89, 9, 4, 'Học nhiếp ảnh cơ bản tốt. Phù hợp cho người mới bắt đầu.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(90, 90, 10, 5, 'Kiến trúc hiện đại được giải thích rõ ràng qua các công trình tiêu biểu.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Reviews cho sách 91-100
(91, 91, 2, 5, 'Bộ sách đồ sộ về lịch sử văn minh. Đọc để mở mang tầm mắt.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(92, 92, 3, 4, 'Giáo trình xã hội học dễ tiếp cận cho người không chuyên.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(93, 93, 4, 5, 'Hiểu sâu về văn hóa Việt qua phân tích khoa học.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(94, 94, 5, 4, 'Lập luận chặt chẽ về tự do kinh tế. Góc nhìn thú vị.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(95, 95, 6, 5, 'Phân tích sắc sảo về mặt trái của toàn cầu hóa.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(96, 96, 7, 5, 'Truyện ngắn Nam Cao - hiện thực phê phán sâu sắc.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(97, 97, 8, 4, 'Văn phong tinh tế, giàu cảm xúc. Thạch Lam đúng là bậc thầy truyện ngắn.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(98, 98, 9, 5, 'Tuyển tập đa dạng, giới thiệu được nhiều tác giả nổi tiếng thế giới.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(99, 99, 10, 5, 'Chekhov - bậc thầy truyện ngắn. Mỗi truyện là một kiệt tác nhỏ.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(100, 100, 2, 4, 'Kết thúc bất ngờ - thương hiệu của O. Henry. Rất giải trí!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());


-- Tạo thêm 100 review (tổng 200 review)
INSERT INTO reviews (id, book_id, user_id, rating, comment, approved, created_at, updated_at) VALUES
-- Thêm reviews cho sách 1-10
(101, 1, 4, 4, 'Tác phẩm kinh điển nhưng cần kiên nhẫn đọc. Đáng giá thời gian!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(102, 1, 5, 5, 'Mỗi nhân vật đều để lại ấn tượng sâu sắc. Jean Valjean là hình mẫu đạo đức.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(103, 2, 6, 4, 'Đồ sộ về cả độ dày và chiều sâu. Đọc để hiểu về lịch sử Nga.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(104, 3, 7, 5, 'Holden Caulfield giống hệt tôi thời tuổi teen. Rất chân thực!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(105, 4, 8, 4, 'Hồi ký xúc động. Đọc mà nhớ về tuổi thơ của ông bà mình.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(106, 5, 9, 5, 'Tình yêu của Heathcliff và Cathy mãi là bi kịch đẹp nhất văn học.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(107, 6, 10, 4, 'Poirot thông minh tuyệt đỉnh. Kết thúc làm tôi bất ngờ hoàn toàn.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(108, 7, 2, 5, 'Đã sưu tầm bộ này. Sherlock Holmes mãi là thần tượng!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(109, 8, 3, 4, 'Kết hợp hoàn hảo giữa nghệ thuật, lịch sử và trinh thám.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(110, 9, 4, 5, 'Yukawa Manabu - giáo sư vật lý thám tử độc đáo nhất!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Thêm reviews cho sách 11-20
(111, 11, 5, 4, 'Thế giới Dune phức tạp nhưng hấp dẫn. Muốn đọc tiếp các phần sau.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(112, 12, 6, 5, 'Big Brother đang theo dõi! Đáng sợ nhưng cần thiết để cảnh giác.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(113, 13, 7, 4, 'Câu hỏi triết học sâu sắc về ranh giới giữa người và máy.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(114, 14, 8, 3, 'Tác phẩm đầu của thể loại nhưng hơi đơn giản so với tiêu chuẩn hiện tại.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(115, 15, 9, 5, 'Psychohistory - ý tưởng thiên tài của Asimov!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(116, 16, 10, 5, 'Middle-earth sống động như thật. Tolkien quả là thiên tài xây dựng thế giới.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(117, 17, 2, 5, 'Mở đầu hoàn hảo cho series huyền thoại. Hogwarts là nhà!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(118, 18, 3, 4, 'Cốt truyện đa tuyến phức tạp. Đọc xong phải xem biểu đồ gia tộc mới hiểu.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(119, 19, 4, 5, 'Văn phong đẹp như nhạc. Đang chờ phần tiếp theo!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(120, 20, 5, 4, 'Thần thoại Bắc Âu thú vị hơn nhiều so với tưởng tượng.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Thêm reviews cho sách 21-30
(121, 21, 6, 5, 'Đọc xong thay đổi cách nhìn về loài người và xã hội.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(122, 22, 7, 4, 'Tài liệu tham khảo tốt nhưng hơi khô khan với người không chuyên.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(123, 23, 8, 5, 'Di sản quý giá của dân tộc. Nên được số hóa để bảo tồn.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(124, 24, 9, 4, 'Lịch sử thế giới được kể như một câu chuyện. Dễ nhớ!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(125, 25, 10, 5, 'Tiểu sử toàn diện nhất về Napoleon từ trước đến nay.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(126, 26, 2, 4, 'Kinh tế học trở nên gần gũi với cuộc sống hàng ngày.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(127, 27, 3, 5, 'Tư duy làm giàu khác biệt. Đã áp dụng và thấy hiệu quả!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(128, 28, 4, 4, 'Bài học về đầu tư giá trị vẫn còn nguyên giá trị.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(129, 29, 5, 5, 'Nghiên cứu đồ sộ về bất bình đẳng. Dữ liệu thuyết phục.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(130, 30, 6, 3, 'Nội dung tương tự nhiều sách self-help khác.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Thêm reviews cho sách 31-40
(131, 31, 7, 5, 'Đắc Nhân Tâm - sách gối đầu giường của người thành công.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(132, 32, 8, 4, '7 thói quen giúp cải thiện hiệu suất công việc đáng kể.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(133, 33, 9, 3, 'Tiêu đề hay nhưng nội dung không đặc sắc.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(134, 34, 10, 4, 'Truyền động lực sống tích cực mỗi ngày.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(135, 35, 2, 5, 'Cách tiếp cận khoa học về thói quen. Thay đổi được nhiều thói quen xấu.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(136, 36, 3, 5, 'Uncle Bob thật sự hiểu về nghề lập trình. Mỗi lập trình viên nên đọc!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(137, 37, 4, 4, 'Tài liệu tham khảo không thể thiếu khi thiết kế phần mềm.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(138, 38, 5, 5, 'Học Python nhanh nhất với cuốn sách này. Project thực tế rất hữu ích.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(139, 39, 6, 4, 'Hiểu sâu về JavaScript qua những phần tốt nhất của ngôn ngữ.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(140, 40, 7, 5, 'Giới thiệu AI toàn diện cho người mới bắt đầu. Rất dễ hiểu!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Thêm reviews cho sách 41-50
(141, 41, 8, 4, 'Phương pháp enzyme giúp cải thiện tiêu hóa đáng kể.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(142, 42, 9, 5, 'Nghiên cứu khoa học nghiêm túc về dinh dưỡng. Đáng tin cậy!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(143, 43, 10, 5, 'Netter s Atlas - không thể thiếu cho sinh viên y khoa!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(144, 44, 2, 4, 'Bảo tồn tri thức y học cổ truyền quý giá của Việt Nam.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(145, 45, 3, 5, 'Mindfulness thay đổi cuộc sống của tôi hoàn toàn.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(146, 46, 4, 4, 'Phân tích tâm lý đám đông vẫn còn nguyên giá trị trong thời đại mạng xã hội.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(147, 47, 5, 5, 'Hiểu được tại sao mình luôn đưa ra quyết định phi lý.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(148, 48, 6, 4, 'Góc nhìn độc đáo từ một cựu FBI profiler.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(149, 49, 7, 3, 'Kỹ năng hữu ích nhưng cần thực hành nhiều mới thành thạo.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(150, 50, 8, 4, 'Lý thuyết hành vi ứng dụng được trong nhiều lĩnh vực.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Thêm reviews cho sách 51-60
(151, 51, 9, 5, 'Câu chuyện nhỏ nhưng ý nghĩa lớn. O. Henry thật tài tình!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(152, 52, 10, 5, 'Hemingway viết về ý chí con người không ai bằng.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(153, 53, 2, 4, 'Tình yêu thời Victoria với tất cả quy tắc và ràng buộc.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(154, 54, 3, 5, 'Atticus Finch - người cha, luật sư lý tưởng của mọi thời đại.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(155, 55, 4, 5, 'Cuộc trả thù hoàn hảo, kế hoạch tỉ mỉ đến từng chi tiết.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(156, 56, 5, 5, 'Truyện Kiều - tinh hoa văn học Việt. Đọc bao lần cũng không chán.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(157, 57, 6, 5, 'Thơ Bác giản dị mà sâu sắc. Thể hiện tinh thần lạc quan cách mạng.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(158, 58, 7, 4, 'Thơ tình Xuân Diệu vẫn làm say lòng bao thế hệ.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(159, 59, 8, 5, 'Thiên tài thơ ca Hàn Mặc Tử. Mỗi bài thơ là một kiệt tác.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(160, 60, 9, 4, 'Thơ Đường hàm súc, ý tại ngôn ngoại. Cần đọc kỹ mới cảm nhận hết.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Thêm reviews cho sách 61-70
(161, 61, 10, 5, 'Hành trình tâm linh thay đổi nhận thức về cuộc sống.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(162, 62, 2, 4, 'Cẩm nang du lịch Việt đầy đủ, nhiều thông tin hữu ích.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(163, 63, 3, 5, 'Tiết kiệm được 30% chi phí du lịch châu Âu nhờ sách này!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(164, 64, 4, 4, 'Hiểu văn hóa Nhật trước khi đi giúp trải nghiệm phong phú hơn.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(165, 65, 5, 5, 'Hành trình chinh phục Himalaya đầy cảm hứng.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(166, 66, 6, 4, 'Nấu được nhiều món Việt ngon nhờ sách này. Công thức dễ làm.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(167, 67, 7, 5, 'Từ không biết gì đến làm được bánh đẹp cho gia đình.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(168, 68, 8, 4, 'Món ăn đường phố châu Á đa dạng, hương vị đặc trưng.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(169, 69, 9, 5, 'Nghệ thuật sushi được nâng tầm thành triết lý sống.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(170, 70, 10, 4, 'Ăn chay không còn nhàm chán nhờ những công thức sáng tạo.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Thêm reviews cho sách 71-80
(171, 71, 2, 5, 'Tuổi thơ với Dế Mèn phiêu lưu ký. Đọc cho con nghe mỗi tối.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(172, 72, 3, 5, 'Phòng chứa bí mật hấp dẫn không kém phần đầu. Rất thích!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(173, 73, 4, 5, 'Hoàng tử bé - cuốn sách nhỏ cho trẻ con, bài học lớn cho người lớn.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(174, 74, 5, 4, 'Truyện cổ Grimm nguyên bản có phần tàn bạo nhưng chân thực.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(175, 75, 6, 5, 'Alice đưa tôi vào thế giới tưởng tượng kỳ diệu. Thích nhân vật Mad Hatter!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(176, 76, 7, 4, 'Phương pháp Montessori giúp trẻ phát triển tự nhiên, toàn diện.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(177, 77, 8, 3, 'Nội dung lặp lại, không có nhiều điểm mới so với các sách khác.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(178, 78, 9, 5, 'Tư duy phản biện - kỹ năng sống còn trong thời đại thông tin.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(179, 79, 10, 4, 'Giáo viên nên đọc để đổi mới phương pháp giảng dạy.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(180, 80, 2, 5, 'Học cách học hiệu quả đã giúp tôi tiết kiệm rất nhiều thời gian.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Thêm reviews cho sách 81-90
(181, 81, 3, 5, 'Kinh thánh ảnh hưởng sâu sắc đến văn hóa phương Tây.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(182, 82, 4, 4, 'Lời Phật dạy ứng dụng được trong công việc và gia đình.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(183, 83, 5, 5, 'Đạo Đức Kinh - triết lý sống thuận tự nhiên.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(184, 84, 6, 4, 'Thiền giúp tôi giảm căng thẳng, sống chậm lại.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(185, 85, 7, 5, 'Hành trình tìm lại chính mình qua những trang sách.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(186, 86, 8, 4, 'Học được cách phân tích và cảm thụ tranh nghệ thuật.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(187, 87, 9, 5, 'Nguyên tắc thiết kế cơ bản nhưng cực kỳ quan trọng cho designer.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(188, 88, 10, 5, 'Lịch sử mỹ thuật từ cổ đại đến hiện đại. Ảnh chất lượng cao.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(189, 89, 2, 4, 'Học nhiếp ảnh cơ bản tốt cho người mới bắt đầu.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(190, 90, 3, 5, 'Kiến trúc hiện đại qua các công trình tiêu biểu.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Thêm reviews cho sách 91-100
(191, 91, 4, 5, 'Lịch sử văn minh phương Tây được kể sinh động.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(192, 92, 5, 4, 'Xã hội học giúp hiểu về cấu trúc và vận hành xã hội.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(193, 93, 6, 5, 'Hiểu sâu về văn hóa Việt qua nghiên cứu khoa học.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(194, 94, 7, 4, 'Tự do kinh tế và vai trò của nhà nước.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(195, 95, 8, 5, 'Phân tích sâu về tác động của toàn cầu hóa.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(196, 96, 9, 5, 'Truyện ngắn Nam Cao - hiện thực phê phán sắc sảo.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(197, 97, 10, 4, 'Văn phong Thạch Lam tinh tế, giàu cảm xúc.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(198, 98, 2, 5, 'Tuyển tập đa dạng các tác giả truyện ngắn thế giới.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(199, 99, 3, 5, 'Chekhov - bậc thầy truyện ngắn với cái kết mở.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(200, 100, 4, 4, 'Kết thúc bất ngờ - đặc trưng của O. Henry.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());


-- Tạo thêm 100 review (tập trung vào một số sách nổi tiếng, bỏ qua một số sách)
INSERT INTO reviews (id, book_id, user_id, rating, comment, approved, created_at, updated_at) VALUES
-- Tập trung nhiều review cho sách 1 (Những Người Khốn Khổ) - thêm 5 review
(201, 1, 6, 5, 'Kiệt tác vượt thời gian. Đọc lại lần thứ 3 vẫn thấy mới.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(202, 1, 7, 4, 'Dài nhưng đáng đọc. Các nhân vật phụ cũng rất đặc sắc.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(203, 1, 8, 5, 'Jean Valjean - hình mẫu về sự chuộc lỗi và lòng nhân ái.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(204, 1, 9, 4, 'Phần mô tả cuộc sống người nghèo ở Paris rất chân thực.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(205, 1, 10, 5, 'Tác phẩm lớn nhất của Victor Hugo. Xứng đáng đọc nhiều lần.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Tập trung nhiều review cho sách 17 (Harry Potter 1) - thêm 6 review
(206, 17, 2, 5, 'Đưa tôi vào thế giới phép thuật từ trang đầu tiên.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(207, 17, 3, 5, 'Món quà tuổi thơ tuyệt vời nhất. Cảm ơn J.K. Rowling!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(208, 17, 4, 4, 'Khởi đầu hoàn hảo cho series huyền thoại.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(209, 17, 5, 5, 'Hogwarts - ngôi nhà mà ai cũng muốn thuộc về.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(210, 17, 6, 5, 'Harry, Ron, Hermione - bộ ba không thể quên.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(211, 17, 7, 5, 'Đọc cho con nghe mỗi tối. Cả nhà đều thích!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Tập trung nhiều review cho sách 27 (Cha Giàu Cha Nghèo) - thêm 4 review
(212, 27, 8, 5, 'Thay đổi hoàn toàn tư duy về tiền bạc.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(213, 27, 9, 4, 'Bài học về tài sản và tiêu sản rất giá trị.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(214, 27, 10, 5, 'Đã mua thêm 5 cuốn tặng bạn bè. Ai cũng cần đọc!', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(215, 27, 2, 3, 'Ý tưởng hay nhưng lặp đi lặp lại nhiều.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Tập trung nhiều review cho sách 31 (Đắc Nhân Tâm) - thêm 5 review
(216, 31, 3, 5, 'Sách gối đầu giường cho người làm kinh doanh.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(217, 31, 4, 5, 'Đọc xong cải thiện được nhiều mối quan hệ.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(218, 31, 5, 4, 'Nguyên tắc vàng trong giao tiếp ứng xử.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(219, 31, 6, 5, 'Tái bản nhiều lần nhưng vẫn luôn bán chạy.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(220, 31, 7, 4, 'Bài học về sự chân thành trong các mối quan hệ.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Tập trung nhiều review cho sách 56 (Truyện Kiều) - thêm 4 review
(221, 56, 8, 5, 'Tinh hoa văn học dân tộc. Thuộc lòng nhiều đoạn.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(222, 56, 9, 5, 'Nguyễn Du - đại thi hào của dân tộc Việt.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(223, 56, 10, 4, 'Câu chuyện về thân phận người phụ nữ xưa.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(224, 56, 2, 5, 'Ngôn ngữ thơ điêu luyện, hình ảnh đẹp.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Tập trung nhiều review cho sách 71 (Dế Mèn) - thêm 3 review
(225, 71, 3, 5, 'Tuổi thơ với Dế Mèn, Cào Cào, Xén Tóc.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(226, 71, 4, 4, 'Bài học về tình bạn và lòng dũng cảm.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(227, 71, 5, 5, 'Tô Hoài viết cho thiếu nhi hay nhất Việt Nam.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Tập trung nhiều review cho sách 96 (Truyện ngắn Nam Cao) - thêm 3 review
(228, 96, 6, 5, 'Chí Phèo, Lão Hạc - những nhân vật bất hủ.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(229, 96, 7, 4, 'Hiện thực phê phán sâu sắc của Nam Cao.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(230, 96, 8, 5, 'Truyện ngắn đỉnh cao của văn học Việt Nam.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Phân bổ review cho các sách khác (mỗi sách 1-2 review)
(231, 2, 9, 4, 'Tiểu thuyết sử thi đồ sộ của Tolstoy.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(232, 3, 10, 5, 'Holden Caulfield - biểu tượng của tuổi trẻ nổi loạn.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(233, 4, 2, 4, 'Hồi ký xúc động về một thời khó khăn.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(234, 5, 3, 5, 'Tình yêu đầy bi kịch trên đồi hoang.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(235, 6, 4, 4, 'Agatha Christie - nữ hoàng trinh thám.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

(236, 12, 5, 5, '1984 - cảnh báo về xã hội kiểm soát.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(237, 16, 6, 5, 'Tolkien - ông tổ của fantasy hiện đại.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(238, 21, 7, 4, 'Sapiens - góc nhìn mới về lịch sử loài người.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(239, 32, 8, 4, '7 thói quen thay đổi cuộc đời.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(240, 36, 9, 5, 'Clean Code - kinh thánh của lập trình viên.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

(241, 41, 10, 4, 'Phương pháp enzyme cải thiện sức khỏe.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(242, 46, 2, 3, 'Tâm lý đám đông - nghiên cứu cổ điển.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(243, 52, 3, 5, 'Hemingway - phong cách văn chương độc đáo.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(244, 61, 4, 4, 'Hành trình tâm linh đầy cảm hứng.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(245, 66, 5, 5, 'Nấu ăn Việt ngon đúng điệu.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

(246, 72, 6, 5, 'Harry Potter phần 2 không làm thất vọng.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(247, 76, 7, 4, 'Montessori - phương pháp giáo dục tiến bộ.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(248, 81, 8, 5, 'Kinh thánh - nền tảng văn hóa phương Tây.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(249, 86, 9, 4, 'Học cách thưởng thức nghệ thuật.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(250, 91, 10, 5, 'Lịch sử văn minh phương Tây hấp dẫn.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Thêm review cho một số sách ít được review (sách 8, 13, 18, 23, 28, 33, 38, 43, 48, 53)
(251, 8, 2, 4, 'Mật mã Da Vinci - kết hợp lịch sử và trinh thám.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(252, 13, 3, 5, 'Câu hỏi triết học sâu sắc về AI.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(253, 18, 4, 4, 'Game of Thrones - cuộc chiến ngai sắt.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(254, 23, 5, 5, 'Đại Việt sử ký - báu vật quốc gia.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(255, 28, 6, 4, 'Benjamin Graham - cha đẻ đầu tư giá trị.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

(256, 33, 7, 3, 'Sách self-help bình thường.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(257, 38, 8, 5, 'Học Python hiệu quả với cuốn sách này.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(258, 43, 9, 5, 'Atlas giải phẫu Netter - không thể thiếu.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(259, 48, 10, 4, 'Tâm lý tội phạm hấp dẫn.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(260, 53, 2, 4, 'Jane Austen - nữ văn sĩ tài năng.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Thêm review cho một số sách khác (sách 58, 63, 68, 73, 78, 83, 88, 93, 98)
(261, 58, 3, 5, 'Thơ tình Xuân Diệu - mãi trẻ trung.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(262, 63, 4, 4, 'Du lịch châu Âu tiết kiệm.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(263, 68, 5, 5, 'Món ăn đường phố châu Á hấp dẫn.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(264, 73, 6, 5, 'Hoàng tử bé - cho mọi lứa tuổi.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(265, 78, 7, 4, 'Tư duy phản biện cần thiết.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

(266, 83, 8, 5, 'Đạo Đức Kinh - triết lý sống thuận tự nhiên.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(267, 88, 9, 5, 'Lịch sử mỹ thuật toàn diện.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(268, 93, 10, 4, 'Văn hóa Việt Nam đa dạng.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(269, 98, 2, 5, 'Truyện ngắn thế giới tuyển chọn hay.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(270, 99, 3, 5, 'Chekhov - bậc thầy truyện ngắn.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Thêm một số review 1 sao, 2 sao để đa dạng rating
(271, 14, 4, 2, 'Tác phẩm cũ, không còn hấp dẫn.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(272, 30, 5, 1, 'Nội dung sáo rỗng, không có giá trị.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(273, 33, 6, 2, 'Tiêu đề hay nhưng nội dung nghèo nàn.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(274, 49, 7, 1, 'Không thực tế, khó áp dụng.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(275, 77, 8, 2, 'Lặp lại nội dung các sách khác.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Thêm một số review chưa approved (đang chờ duyệt)
(276, 1, 9, 5, 'Tuyệt vời! Đã đọc 3 lần.', 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(277, 17, 10, 5, 'Harry Potter sống mãi!', 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(278, 27, 2, 4, 'Sách hay về tài chính.', 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(279, 31, 3, 5, 'Gối đầu giường!', 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(280, 56, 4, 5, 'Kiệt tác dân tộc.', 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

-- Review cuối cùng
(281, 100, 5, 4, 'O. Henry - ông vua kết thúc bất ngờ.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(282, 95, 6, 5, 'Phân tích sâu về toàn cầu hóa.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(283, 90, 7, 4, 'Kiến trúc hiện đại thú vị.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(284, 85, 8, 5, 'Hành trình tâm linh ý nghĩa.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(285, 80, 9, 4, 'Học cách học hiệu quả.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

(286, 75, 10, 5, 'Alice - hành trình kỳ diệu.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(287, 70, 2, 4, 'Món chay ngon và bổ dưỡng.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(288, 65, 3, 5, 'Himalaya - thách thức và cảm hứng.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(289, 60, 4, 4, 'Thơ Đường hàm súc.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(290, 55, 5, 5, 'Bá tước Monte Cristo - trả thù hoàn hảo.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

(291, 50, 6, 3, 'Tâm lý học hành vi - khá khô khan.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(292, 45, 7, 5, 'Mindfulness thay đổi cuộc sống.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(293, 40, 8, 4, 'Giới thiệu AI dễ hiểu.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(294, 35, 9, 5, 'Atomic Habits - khoa học về thói quen.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(295, 25, 10, 4, 'Tiểu sử Napoleon chi tiết.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),

(296, 20, 2, 3, 'Thần thoại Bắc Âu - đọc cho biết.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(297, 15, 3, 5, 'Foundation - tầm nhìn vượt thời đại.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(298, 10, 4, 4, 'Trinh thám hấp dẫn.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(299, 5, 5, 5, 'Đồi gió hú - tình yêu và hận thù.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(300, 1, 6, 5, 'Victor Hugo - thiên tài văn chương.', 1, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());


-- Cập nhật average_rating và review_count cho các sách
UPDATE books b
SET average_rating = (
    SELECT AVG(rating)
    FROM reviews r
    WHERE r.book_id = b.id AND r.approved = 1
),
    review_count = (
        SELECT COUNT(*)
        FROM reviews r
        WHERE r.book_id = b.id AND r.approved = 1
    ) where  b.id >-1;