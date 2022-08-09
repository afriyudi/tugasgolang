# tugasgolang

```
1. buat table :
CREATE TABLE `employee1` (
  `id` int(6) UNSIGNED NOT NULL,
  `task` varchar(80) NOT NULL,
  `assignee` varchar(30) NOT NULL,
  `dateline` date NOT NULL,
  `mark` varchar(1) DEFAULT 'N'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

ALTER TABLE `employee1`
  ADD PRIMARY KEY (`id`);

2. insert data
INSERT INTO `employee1` (`id`, `task`, `assignee`, `dateline`, `mark`) VALUES
(1, 'edit Program', 'jack', '0000-00-00', 'Y'),
(2, 'edit Program', 'jack', '0000-00-00', 'N'),
(4, 'buat kopi', 'pok inem', '0000-00-00', 'Y'),
(5, 'bikin pempek', 'om barong', '2022-08-09', 'Y'),
(6, 'Membuat Pindang', 'Bos Besak', '2022-08-10', 'N'),
(7, 'bikin pempek lenggang', 'babeyudi', '2022-08-10', 'Y'),
(8, 'Ngerjain Tugas Golang', 'Mas Letfi', '2022-08-11', 'N');

ALTER TABLE `employee1`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

```
