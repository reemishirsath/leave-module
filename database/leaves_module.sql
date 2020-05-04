
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";
--
-- Database: `leaves_module`
--

-- --------------------------------------------------------

--
-- Table structure for table `leaves`
--

CREATE TABLE `leaves` (
  `id` int(10) UNSIGNED NOT NULL,
  `user_id` int(10) UNSIGNED NOT NULL,
  `status` enum('APPLIED','REJECTED','APPROVED') NOT NULL DEFAULT 'APPLIED',
  `leave_reason` varchar(500) NOT NULL,
  `feedback` varchar(500) NOT NULL,
  `from_date` date NOT NULL,
  `to_date` date NOT NULL,
  `leave_days` tinyint(3) UNSIGNED NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `leaves`
--

INSERT INTO `leaves` (`id`, `user_id`, `status`, `leave_reason`, `feedback`, `from_date`, `to_date`, `leave_days`, `created_at`) VALUES
(1, 2, 'APPROVED', 'Personal Leave', 'This is Approved Feedback', '2020-01-31', '2020-02-05', 5, '2020-02-04 03:40:25'),
(2, 2, 'REJECTED', 'Personal Leave', 'This is Reject Feedback', '2020-01-31', '2020-02-05', 5, '2020-02-04 04:01:38'),
(3, 2, 'APPLIED', 'Personal Leave', '', '2020-01-31', '2020-02-05', 5, '2020-02-04 04:09:27'),
(4, 2, 'APPLIED', 'Personal Leave', '', '2020-01-31', '2020-02-05', 5, '2020-02-04 04:13:29');

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` int(5) UNSIGNED NOT NULL,
  `role_name` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `role_name`) VALUES
(1, 'Admin'),
(2, 'User');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(20) NOT NULL,
  `email` varchar(50) NOT NULL,
  `password` varchar(100) NOT NULL,
  `role_id` int(5) UNSIGNED NOT NULL,
  `total_leaves` tinyint(3) NOT NULL DEFAULT '21',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`, `role_id`, `total_leaves`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'admin', 'admin@scalent.io', '$2a$04$FddCOm6PXR7Rs524ISWcIev1iotTj76BvK1FKjzken4DlHpd5fm8S', 1, 100, '2020-02-03 07:23:20', NULL, NULL),
(2, 'user1', 'user@scalent.io', '$2a$04$C092gPZaSfkuvL57jCdhy.4fcTIzWPvT0BJWcTvAObRTJiyYNcCtO', 2, 100, '2020-02-03 07:23:20', NULL, NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `leaves`
--
ALTER TABLE `leaves`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `role_id_key` (`role_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `leaves`
--
ALTER TABLE `leaves`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
--
-- AUTO_INCREMENT for table `roles`
--
ALTER TABLE `roles`
  MODIFY `id` int(5) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- Constraints for dumped tables
--

--
-- Constraints for table `leaves`
--
ALTER TABLE `leaves`
  ADD CONSTRAINT `user_id_key` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE;

--
-- Constraints for table `users`
--
ALTER TABLE `users`
  ADD CONSTRAINT `role_id_key` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON UPDATE CASCADE;
