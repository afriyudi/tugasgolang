# tugasgolang

buat table :
CREATE TABLE `employee1` (
  `id` int(6) UNSIGNED NOT NULL,
  `task` varchar(80) NOT NULL,
  `assignee` varchar(30) NOT NULL,
  `dateline` date NOT NULL,
  `mark` varchar(1) DEFAULT 'N'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
