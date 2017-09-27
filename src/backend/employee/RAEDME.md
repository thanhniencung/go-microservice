- Tạo mới một nhân viên 
- Chỉnh sửa thông tin nhân viên 
- Xoá thông tin nhân viên 
- Lấy tất cả thông tin có trong database 


CREATE TABLE `hr-management`.`employee` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `firstName` VARCHAR(100) NOT NULL,
  `lastName` VARCHAR(100) NOT NULL,
  `joinDate` DATETIME NOT NULL,
  `teamId` INT NOT NULL,
  `avatar` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`));