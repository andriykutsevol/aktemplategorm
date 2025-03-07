-- phpMyAdmin SQL Dump
-- version 5.2.2
-- https://www.phpmyadmin.net/
--
-- Хост: db_service:3306
-- Время создания: Мар 07 2025 г., 10:19
-- Версия сервера: 8.4.4
-- Версия PHP: 8.2.27

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `mus`
--

--
-- Дамп данных таблицы `Psp`
--

INSERT INTO `Psp` (`ID`, `PspCode`, `PspShortName`, `PspCountryCode`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, `CreatedBy`, `UpdatedBy`, `DeletedBy`) VALUES
(1, 'CG-MTN-MTN', 'CG-MTN-MTN_1 Short name', 'CM', '2025-03-07 08:40:08', '2025-03-07 08:40:08', NULL, 0xde95ca9d489846f4af79b1e70326b4c1, 0xde95ca9d489846f4af79b1e70326b4c1, NULL),
(2, 'CG-MTN-MTN', 'CG-MTN-MTN_1 Short name', 'CM', '2025-03-07 08:40:19', '2025-03-07 08:40:19', NULL, 0xde95ca9d489846f4af79b1e70326b4c1, 0xde95ca9d489846f4af79b1e70326b4c1, NULL),
(3, 'CG-MTN-MTN', 'CG-MTN-MTN_1 Short name', 'CM', '2025-03-07 08:40:24', '2025-03-07 08:40:24', NULL, 0xde95ca9d489846f4af79b1e70326b4c1, 0xde95ca9d489846f4af79b1e70326b4c1, NULL);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
