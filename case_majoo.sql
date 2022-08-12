-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 06 Agu 2022 pada 05.17
-- Versi server: 10.4.13-MariaDB
-- Versi PHP: 7.3.20

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `case_majoo`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `merchants`
--

CREATE TABLE `merchants` (
  `id` bigint(20) NOT NULL,
  `user_id` int(40) NOT NULL,
  `merchant_name` varchar(40) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `created_by` bigint(20) NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_by` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `merchants`
--

INSERT INTO `merchants` (`id`, `user_id`, `merchant_name`, `created_at`, `created_by`, `updated_at`, `updated_by`) VALUES
(1, 1, 'merchant 1', '2022-08-05 03:17:21', 1, '2022-08-05 03:17:21', 1),
(2, 2, 'Merchant 2', '2022-08-05 03:17:21', 2, '2022-08-05 03:17:21', 2);

-- --------------------------------------------------------

--
-- Struktur dari tabel `outlets`
--

CREATE TABLE `outlets` (
  `id` bigint(20) NOT NULL,
  `merchant_id` bigint(20) NOT NULL,
  `outlet_name` varchar(40) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `created_by` bigint(20) NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_by` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `outlets`
--

INSERT INTO `outlets` (`id`, `merchant_id`, `outlet_name`, `created_at`, `created_by`, `updated_at`, `updated_by`) VALUES
(1, 1, 'Outlet 1', '2022-08-05 03:17:21', 1, '2022-08-05 03:17:21', 1),
(2, 2, 'Outlet 1', '2022-08-05 03:17:21', 2, '2022-08-05 03:17:21', 2),
(3, 1, 'Outlet 2', '2022-08-05 03:17:21', 1, '2022-08-05 03:17:21', 1);

-- --------------------------------------------------------

--
-- Struktur dari tabel `transactions`
--

CREATE TABLE `transactions` (
  `id` bigint(20) NOT NULL,
  `merchant_id` bigint(20) NOT NULL,
  `outlet_id` bigint(20) NOT NULL,
  `bill_total` double NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `created_by` bigint(20) NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_by` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `transactions`
--

INSERT INTO `transactions` (`id`, `merchant_id`, `outlet_id`, `bill_total`, `created_at`, `created_by`, `updated_at`, `updated_by`) VALUES
(1, 1, 1, 2000, '2021-11-01 05:30:04', 1, '2021-11-01 05:30:04', 1),
(2, 1, 1, 2500, '2021-11-01 10:20:14', 1, '2021-11-01 10:20:14', 1),
(3, 1, 1, 4000, '2021-11-02 05:30:04', 1, '2021-11-02 05:30:04', 1),
(4, 1, 1, 1000, '2021-11-04 05:30:04', 1, '2021-11-04 05:30:04', 1),
(5, 1, 1, 7000, '2021-11-05 09:59:30', 1, '2021-11-05 09:59:30', 1),
(6, 1, 3, 2000, '2021-11-02 11:30:04', 1, '2021-11-02 11:30:04', 1),
(7, 1, 3, 2500, '2021-11-03 10:20:14', 1, '2021-11-03 10:20:14', 1),
(8, 1, 3, 4000, '2021-11-04 05:30:04', 1, '2021-11-04 05:30:04', 1),
(9, 1, 3, 1000, '2021-11-04 05:31:04', 1, '2021-11-04 05:31:04', 1),
(10, 1, 3, 7000, '2021-11-05 09:59:30', 1, '2021-11-05 09:59:30', 1),
(11, 2, 2, 2000, '2021-11-01 11:30:04', 2, '2021-11-01 11:30:04', 2),
(12, 2, 2, 2500, '2021-11-02 10:20:14', 2, '2021-11-02 10:20:14', 2),
(13, 2, 2, 4000, '2021-11-03 05:30:04', 2, '2021-11-03 05:30:04', 2),
(14, 2, 2, 1000, '2021-11-04 05:31:04', 2, '2021-11-04 05:31:04', 2),
(15, 2, 2, 7000, '2021-11-05 09:59:30', 2, '2021-11-05 09:59:30', 2),
(16, 2, 2, 2000, '2021-11-05 11:30:04', 2, '2021-11-05 11:30:04', 2),
(17, 2, 2, 2500, '2021-11-06 10:20:14', 2, '2021-11-06 10:20:14', 2),
(18, 2, 2, 4000, '2021-11-07 05:30:04', 2, '2021-11-07 05:30:04', 2),
(19, 2, 2, 1000, '2021-11-08 05:31:04', 2, '2021-11-08 05:31:04', 2),
(20, 2, 2, 7000, '2021-11-09 09:59:30', 2, '2021-11-09 09:59:30', 2),
(21, 2, 2, 1000, '2021-11-10 05:31:04', 2, '2021-11-10 05:31:04', 2),
(22, 2, 2, 7000, '2021-11-11 09:59:30', 2, '2021-11-11 09:59:30', 2);

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `user_name` varchar(45) DEFAULT NULL,
  `password` varchar(225) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `created_by` bigint(20) NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_by` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `user_name`, `password`, `created_at`, `created_by`, `updated_at`, `updated_by`) VALUES
(1, 'Admin 1', 'admin1', 'e00cf25ad42683b3df678c61f42c6bda', '2022-08-05 03:17:20', 1, '2022-08-05 03:17:20', 1),
(2, 'Admin 2', 'admin2', 'c84258e9c39059a89ab77d846ddab909', '2022-08-05 03:17:20', 2, '2022-08-05 03:17:20', 2);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `merchants`
--
ALTER TABLE `merchants`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `outlets`
--
ALTER TABLE `outlets`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `merchants`
--
ALTER TABLE `merchants`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `outlets`
--
ALTER TABLE `outlets`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT untuk tabel `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
