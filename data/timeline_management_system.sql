-- phpMyAdmin SQL Dump
-- version 4.7.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 11, 2017 at 05:22 PM
-- Server version: 10.1.28-MariaDB
-- PHP Version: 5.6.32

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `timeline_management_system`
--

-- --------------------------------------------------------

--
-- Table structure for table `divisi`
--

CREATE TABLE `divisi` (
  `ID_Divisi` varchar(6) NOT NULL,
  `Nama_Divisi` varchar(50) NOT NULL,
  `Ketua_Divisi` varchar(50) NOT NULL,
  `Bidang` varchar(40) NOT NULL,
  `Ketua_Bidang` varchar(40) NOT NULL,
  `Deskripsi_Divisi` text NOT NULL,
  `E-mail_Ketua_Divisi` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `divisi`
--

INSERT INTO `divisi` (`ID_Divisi`, `Nama_Divisi`, `Ketua_Divisi`, `Bidang`, `Ketua_Bidang`, `Deskripsi_Divisi`, `E-mail_Ketua_Divisi`) VALUES
('ACR030', 'Acara', 'Rezha Kusuma A', 'Acara', 'Rezha Kusuma A', 'Ketua Bidang Acara ', 'rezha.kusuma@gmail.com'),
('AKT020', 'Akomtrans', 'Irene Edria Devina', 'Operasional', 'Alif Bhaskoro', 'Akomtrans Arkavidia', 'ireneedriadr@gmail.com'),
('BND005', 'Bendahara 1', 'Tessa Angela', 'Kesekjenan', 'Pratamamia Agung Prihatmaja', 'Bendahara 1 Acara Arkavidia', 'emmanuellatessaangela@gmail.com'),
('BND006', 'Bendahara 2', 'Iftitakhul Zakiah', 'Kesekjenan', 'Pratamamia Agung Prihatmaja', 'Bendahara 2 Acara Arkavidia', 'Iftitakhulzakiyah@gmail.com'),
('CAM026', 'Marketing dan Creative', 'Joshua Atmadja', 'Marketing dan Creative', 'Joshua Atmadja', 'Marketing dan Creative Acara Arkavidia', 'joshuatmadja@gmail.com'),
('CTF022', 'Capture The Flag', 'Kevin Supendi', 'Lomba', 'Muhammad Isham Azmansyah F', 'Capture The Flag Acara Arkavidia', '13514094@std.stei.itb.ac.id'),
('DKR018', 'Dekorasi', 'Audry Nyonata', 'Operasional', 'Alif Bhaskoro', 'Dekorasi Acara Arkavidia', 'audrynyonata@gmail.com'),
('DLG014', 'Delegasi', 'Adya Naufal Fikri', 'Eksternal', 'Wiega Sonora', 'Delegasi Arkavidia', '13515130@std.stei.itb.ac.id'),
('DNS010', 'Danus', 'Mikhael Artur Darmakesuma', 'Fundraising', 'Theo Adhitya Sani W.', 'Danus Arkavidia', 'mikhael.artur.d@gmail.com '),
('DSN027', 'Desain', 'Atika Azzahra', 'Marketing dan Creative', 'Joshua Atmadja', 'Desain Acara Arkavidia', 'atikazzahraa@gmail.com '),
('EKS011', 'Eksternal', 'Wiega Sonora', 'Eksternal', 'Wiega Sonora', 'Ketua Bidang Eksternal', 'wiegasonora@gmail.com'),
('EXP032', 'Expo', 'Arief Septian Nurhada', 'Acara', 'Rezha Kusuma A', 'Expo Acara Arkavidia', 'ariefseptian7@gmail.com'),
('FRS008', 'Fundraising', 'Theo Adhitya Sani W.', 'Fundraising', 'Theo Adhitya Sani W.', 'Ketua Bidang Fundraising', 'theo_asw@yahoo.com'),
('HCK025', 'Hackaton', 'Azis Adi Kuncoro', 'Lomba', 'Muhammad Isham Azmansyah F', 'Hackaton Arkavidia', 'azisadikuncoro@gmail.com'),
('ITT029', 'IT', 'Muhammad Reza Ramadhan', 'Marketing dan Creative', 'Joshua Atmadja', 'IT Arkavidia', 'rezaramadhan.m@gmail.com '),
('KTW001', 'Ketua', 'Praditya Raudi', 'Ketua', 'Praditya Raudi', 'Ketua Acara Arkavidia', 'raudiradz@gmail.com'),
('LGS016', 'Logistik', 'Martin Lutta Putra', 'Operasional', 'Alif Bhaskoro', 'Logistik Arkavidia', 'sukaberkhayal@gmail.com'),
('LMB021', 'Lomba', 'Muhammad Isham Azmansyah F.', 'Lomba', 'Muhammad Isham Azmansyah F', 'Ketua Bidang Lomba', 'i@adimaja.com'),
('LOF013', 'LO', 'Rizki Ihza Parama', 'Eksternal', 'Wiega Sonora', 'LO Arkavidia', 'rizki.ihza@gmail.com'),
('MSD007', 'MSDM', 'Fahmi Satria Aji', 'Kesekjenan', 'Pratamamia Agung Prihatmaja', 'MSDM Arkavidia', 'teemeezajee@gmail.com '),
('OPS015', 'Operasional', 'Alif Bhaskoro', 'Operasional', 'Alif Bhaskoro', 'Operasional Acara Arkavidia', 'Alifbhaskoro@gmail.com'),
('PBD028', 'Publikasi dan Dokumentasi', 'Aya Aurora Rimbamorani', 'Marketing dan Creative', 'Joshua Atmadja', 'Publikasi dan Dokumentasi Arkavidia', 'aya.aurora25@gmail.com'),
('PCS023', 'Programming Contest', 'Jauhar Arifin', 'Lomba', 'Muhammad Isham Azmansyah F', 'Programming Contest Acara Arkavidia', 'jauhararifin10@gmail.com'),
('PEG033', 'Pre Event dan Gala Dinner', 'Adrian Hartanto P', 'Acara', 'Rezha Kusuma A', 'Pre Event dan Gala Dinner Acara Arkavidia', 'adrianhp100797@gmail.com'),
('PRD017', 'Produksi', 'Kevin Liem', 'Operasional', 'Alif Bhaskoro', 'Produksi Acara Arkavidia', 'kevin18.liem@gmail.com'),
('PRL012', 'Public Relations', 'Radiyya Dwisaputra', 'Eksternal', 'Wiega Sonora', 'Public Relations Acara Arkavidia', 'radiyyasaputra@gmail.com'),
('PRZ019', 'Perizinan', 'Kharis Isriyanto', 'Operasional', 'Alif Bhaskoro', 'Perizinan Acara Arkavidia', 'kharisisriyant@gmail.com'),
('SKJ002', 'Kesekjenan', 'Pratamamia Agung Prihatmaja', 'Kesekjenan', 'Pratamamia Agung Prihatmaja', 'Kesekjenan Arkavidia', 'apratamamia@gmail.com'),
('SKR003', 'Sekretaris 1', 'Aliyah Sausan Huwel', 'Kesekjenan', 'Pratamamia Agung Prihatmaja', 'Sekretaris 1 Arkavidia', 'Sausanhuwel@gmail.com'),
('SKR004', 'Sekretaris 2', 'Steffi Indrayani', 'Kesekjenan', 'Pratamamia Agung Prihatmaja', 'Sekretaris 2 Acara Arkavidia', 'steffiindrayani16@gmail.com'),
('SMR031', 'Seminar', 'Teo Wijayarto', 'Acara', 'Rezha Kusuma A', 'Seminar Acara Arkavidia', '18215004@std.stei.itb.ac.id'),
('SPS009', 'Sponsorship', 'Rizky Faramita', 'Fundraising', 'Theo Adhitya Sani W.', 'Sponsorship Acara Arkavidia', 'rizky.faramitha@gmail.com'),
('TCV024', 'Technovation ', 'Bethea Zia Davida', 'Lomba', 'Muhammad Isham Azmansyah F', 'Technovation Arkavidia', 'betheazia17@gmail.com');

-- --------------------------------------------------------

--
-- Table structure for table `memiliki`
--

CREATE TABLE `memiliki` (
  `ID_Divisi` varchar(6) NOT NULL,
  `ID_Pekerjaan` varchar(7) NOT NULL,
  `Nama_Divisi` varchar(50) NOT NULL,
  `Tanggal_Tenggat` date NOT NULL,
  `Penanggung_Jawab` varchar(50) NOT NULL,
  `Deskripsi_Pekerjaan` text NOT NULL,
  `Kontak` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `memiliki`
--

INSERT INTO `memiliki` (`ID_Divisi`, `ID_Pekerjaan`, `Nama_Divisi`, `Tanggal_Tenggat`, `Penanggung_Jawab`, `Deskripsi_Pekerjaan`, `Kontak`) VALUES
('MSD007', 'PKJ001', 'MSDM', '2018-01-31', 'Fahmi Satria Aji', 'Rapat Koordinasi Panitia', 'teemeezajee@gmail.com '),
('SKR003', 'PKJ002', 'Sekretaris 1', '2017-12-21', 'Aliyah Sausan Huwel', 'Penyebaran Proposal ke Instansi Terkait', 'Sausanhuwel@gmail.com'),
('SKJ002', 'PKJ003', 'Kesekjenan', '2017-12-17', 'Pratamamia Agung Prihatmaja', 'Follow Up Proposal', 'apratamamia@gmail.com'),
('PBD028', 'PKJ004', 'Publikasi dan Dokumentasi', '2017-02-11', 'Aya Aurora Rimbamorani', 'Publikasi Media Sosial', 'aya.aurora25@gmail.com'),
('SPS009', 'PKJ005', 'Sponsorship', '2017-12-16', 'Rizky Faramita', 'Follow Up Proposal Sponsor', 'rizky.faramitha@gmail.com'),
('DLG014', 'PKJ006', 'Delegasi', '2017-12-19', 'Adya Naufal Fikri', 'Penyebaran Invitation Letter Delegasi', '13515130@std.stei.itb.ac.id'),
('EXP032', 'PKJ007', 'Expo', '2017-12-21', 'Arief Septian Nurhada', 'Opening Registration Untuk Workshop', 'ariefseptian7@gmail.com'),
('LMB021', 'PKJ008', 'Lomba', '2017-12-24', 'Muhammad Isham Azmansyah F', 'Opening Registration Lomba', 'i@adimaja.com'),
('SMR031', 'PKJ009', 'Seminar', '2017-12-31', 'Teo Wijayarto', 'Opening Registration Seminar', '18215004@std.stei.itb.ac.id'),
('CAM026', 'PKJ010', 'Marketing dan Creative', '2018-01-01', 'Joshua Atmadja', 'Release Teaser Arkavidia', 'joshuatmadja@gmail.com'),
('DKR018', 'PKJ011', 'Dekorasi', '2018-01-12', 'Audry Nyonata', 'Dekorasi Pensuasanaan di ITB', 'audrynyonata@gmail.com'),
('KTW001', 'PKJ012', 'Ketua ', '2018-01-17', 'Praditya Raudi', 'Technical Meeting dengan Instansti Terlibat', 'raudiradz@gmail.com'),
('PEG033', 'PKJ013', 'Pre Event dan Gala Dinner', '2018-01-24', 'Adrian Hartanto Pramudita', 'Workshop Startup Camp', 'adrianhp100797@gmail.com'),
('LMB021', 'PKJ014', 'Lomba', '2018-02-02', 'Muhammad Isham Azmansyah F', 'Pengumuman Finalis Lomba', 'i@adimaja.com'),
('KTW001', 'PKJ015', 'Ketua ', '2018-02-09', 'Praditya Raudi', 'Show Time', 'raudiradz@gmail.com');

-- --------------------------------------------------------

--
-- Table structure for table `pekerjaan`
--

CREATE TABLE `pekerjaan` (
  `ID_Pekerjaan` varchar(10) NOT NULL,
  `Nama_Pekerjaan` varchar(50) NOT NULL,
  `Tanggal_Mulai` date NOT NULL,
  `Tanggal_Selesai` date NOT NULL,
  `Penanggung_Jawab` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `pekerjaan`
--

INSERT INTO `pekerjaan` (`ID_Pekerjaan`, `Nama_Pekerjaan`, `Tanggal_Mulai`, `Tanggal_Selesai`, `Penanggung_Jawab`) VALUES
('PKJ001', 'Rapat Koordinasi Panitia', '2017-12-14', '2018-01-31', 'Fahmi Satria Aji'),
('PKJ002', 'Penyebaran Proposal ke Instansi Terkait', '2017-12-14', '2017-12-21', 'Aliyah Sausan Huwel'),
('PKJ003', 'Follow Up Proposal', '2017-12-14', '2017-12-17', 'Pratamamia Agung Prihatmaja'),
('PKJ004', 'Publikasi Media Sosial', '2017-12-14', '2017-02-11', 'Aya Aurora Rimbamorani'),
('PKJ005', 'Follow Up Proposal Sponsor', '2017-12-14', '2017-12-16', 'Rizky Faramita'),
('PKJ006', 'Penyebaran Invitation Letter Delegasi', '2017-12-14', '2017-12-19', 'Adya Naufal Fikri'),
('PKJ007', 'Opening Registration Untuk Workshop', '2017-12-14', '2017-12-21', 'Arief Septian Nurhada'),
('PKJ008', 'Opening Registration Lomba', '2017-12-14', '2017-12-24', 'Muhammad Isham Azmansyah F'),
('PKJ009', 'Opening Registration Seminar', '2017-12-14', '2017-12-31', 'Teo Wijayarto'),
('PKJ010', 'Release Teaser Arkavidia', '2017-12-14', '2018-01-01', 'Joshua Atmadja'),
('PKJ011', 'Dekorasi Pensuasanaan di ITB', '2017-12-14', '2018-01-12', 'Audry Nyonata'),
('PKJ012', 'Technical Meeting dengan Instansti Terlibat', '2017-12-14', '2018-01-17', 'Praditya Raudi'),
('PKJ013', 'Workshop Startup Camp', '2017-12-14', '2018-01-24', 'Adrian Hartanto Pramudita'),
('PKJ014', 'Pengumuman Finalis Lomba', '2017-12-14', '2018-02-02', 'Muhammad Isham Azmansyah F'),
('PKJ015', 'Show Time', '2017-12-14', '2018-02-09', 'Praditya Raudi');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `divisi`
--
ALTER TABLE `divisi`
  ADD PRIMARY KEY (`ID_Divisi`);

--
-- Indexes for table `memiliki`
--
ALTER TABLE `memiliki`
  ADD KEY `ID_Divisi` (`ID_Divisi`),
  ADD KEY `ID_Pekerjaan` (`ID_Pekerjaan`);

--
-- Indexes for table `pekerjaan`
--
ALTER TABLE `pekerjaan`
  ADD PRIMARY KEY (`ID_Pekerjaan`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
