-- phpMyAdmin SQL Dump
-- version 4.8.3
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2019-01-16 11:34:04
-- 服务器版本： 5.7.24
-- PHP 版本： 5.5.38

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `nkn-data-v1`
--

-- --------------------------------------------------------

--
-- 表的结构 `t_addr`
--

CREATE TABLE `t_addr` (
  `hash` varchar(64) NOT NULL COMMENT '地址哈希的base 58值',
  `first_used_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '地址首次使用时间',
  `first_used_block_height` int(10) UNSIGNED NOT NULL COMMENT '地址首次使用所在的块'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `t_addr_asset`
--

CREATE TABLE `t_addr_asset` (
  `addr` varchar(64) NOT NULL COMMENT '账户地址',
  `asset_id` varchar(64) NOT NULL COMMENT '地址名下拥有的资产id',
  `asset_value` varchar(30) NOT NULL COMMENT '拥有的资产数量'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `t_addr_wallet`
--

CREATE TABLE `t_addr_wallet` (
  `wallet_id` int(10) UNSIGNED NOT NULL COMMENT '钱包id',
  `address` int(11) NOT NULL COMMENT 'utxo地址'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `t_assets`
--

CREATE TABLE `t_assets` (
  `hash` varchar(64) NOT NULL COMMENT '资产id，一个64位哈希值',
  `amount` varchar(30) NOT NULL COMMENT '总量',
  `name` varchar(100) NOT NULL COMMENT '资产名称，最长100字符，超出截断',
  `description` varchar(1000) NOT NULL COMMENT '资产描述，最长1000字符，超出截断',
  `asset_precision` int(10) UNSIGNED NOT NULL COMMENT '资产精度，DNA默认是8位精度',
  `asset_type` int(10) UNSIGNED NOT NULL COMMENT '资产类型',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '注册时间',
  `height` int(10) UNSIGNED NOT NULL COMMENT '注册时的区块高度'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `t_assets_issue_record`
--

CREATE TABLE `t_assets_issue_record` (
  `asset_id` varchar(64) NOT NULL COMMENT '资产id，一个64位哈希值',
  `issue_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '发行时间',
  `issue_to` varchar(64) NOT NULL COMMENT '发行到了哪个地址',
  `value` varchar(30) NOT NULL COMMENT '发行数量',
  `height` int(11) NOT NULL COMMENT '发行时的区块高度'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `t_assets_transfer`
--

CREATE TABLE `t_assets_transfer` (
  `hash` varchar(64) NOT NULL COMMENT '所属交易hash',
  `height` int(10) UNSIGNED NOT NULL COMMENT '交易所在区块高度',
  `height_tx_idx_union` bigint(20) UNSIGNED NOT NULL COMMENT '区块高度、交易序号、转账序号的联合值',
  `from_addr` varchar(64) NOT NULL COMMENT '转出账户',
  `to_addr` varchar(64) NOT NULL COMMENT '转入账户',
  `asset_id` varchar(64) NOT NULL COMMENT '资产id',
  `value` varchar(30) NOT NULL COMMENT '转账数额',
  `fee` varchar(30) NOT NULL COMMENT '转账费用',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '转账时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `t_blocks`
--

CREATE TABLE `t_blocks` (
  `height` int(10) UNSIGNED NOT NULL COMMENT '块高度',
  `hash` varchar(64) NOT NULL COMMENT '本区块哈希',
  `prev_hash` varchar(64) NOT NULL COMMENT '前一个区块的哈希',
  `next_hash` varchar(64) NOT NULL COMMENT '后一个区块的哈希',
  `signature` varchar(128) NOT NULL COMMENT '本区块签名',
  `signer_id` varchar(64) NOT NULL,
  `signer_pk` varchar(64) NOT NULL,
  `state_root` varchar(64) NOT NULL,
  `validator` varchar(100) NOT NULL,
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '本区块时间戳',
  `transaction_root` varchar(64) NOT NULL COMMENT '交易树根',
  `winner_hash` varchar(64) NOT NULL,
  `size` int(10) UNSIGNED NOT NULL COMMENT '本区块大小',
  `transaction_count` int(10) UNSIGNED NOT NULL COMMENT '本区块交易数量'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `t_nodes`
--

CREATE TABLE `t_nodes` (
  `server_addr` varchar(256) NOT NULL COMMENT '节点地址，可以是域名',
  `type` int(10) UNSIGNED NOT NULL COMMENT '节点类型',
  `status` int(10) UNSIGNED NOT NULL COMMENT '节点状态',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `t_pay`
--

CREATE TABLE `t_pay` (
  `id` int(10) UNSIGNED NOT NULL,
  `height` int(10) UNSIGNED NOT NULL,
  `tx_hash` varchar(64) NOT NULL,
  `payer` varchar(64) NOT NULL,
  `payee` varchar(64) NOT NULL,
  `value` varchar(40) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `t_sigchain`
--

CREATE TABLE `t_sigchain` (
  `height` int(11) NOT NULL COMMENT '打包进的区块高度',
  `sig_idx` int(10) UNSIGNED NOT NULL COMMENT '签名索引',
  `id`    varchar(64) NOT NULL,
  `next_pubkey` varchar(100) NOT NULL,
  `tx_hash` varchar(100) NOT NULL COMMENT '所在交易的hash',
  `sig_data` varchar(2000) NOT NULL COMMENT '签名',
  `sig_algo` varchar(16),
  `vrf` varchar(64),
  `proof` varchar(128),
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `t_transactions`
--

CREATE TABLE `t_transactions` (
  `hash` varchar(64) NOT NULL COMMENT '交易的哈希值',
  `height` int(10) UNSIGNED NOT NULL COMMENT '交易所在区块高度',
  `height_idx_union` bigint(20) UNSIGNED NOT NULL COMMENT '所在区块高度和交易在区块内的索引的联合值',
  `tx_type` int(10) UNSIGNED NOT NULL COMMENT '交易类型',

  `attributes` varchar(64) NOT NULL,
  `fee`   int(10) NOT NULL,
  `nonce` int(10) NOT NULL,

  `asset_id` varchar(64) COMMENT '交易的资产id',
  `utxo_input_count` int(10) UNSIGNED COMMENT 'utxo输入个数',
  `utxo_output_count` int(10) UNSIGNED COMMENT 'utxo的输出个数',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '交易时间，和区块时间相同',
  `parse_status` int(10) UNSIGNED NOT NULL COMMENT '交易的解析状态，用于容错'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `t_generate_id`
--

CREATE TABLE `t_generate_id` (
  `height` int(11) NOT NULL COMMENT '打包进的区块高度',
  `tx_hash` varchar(100) NOT NULL COMMENT '所在交易的hash',
  `public_key` varchar(64) NOT NULL,
  `registration_fee` int(10) UNSIGNED NOT NULL,
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- --------------------------------------------------------


-- --------------------------------------------------------

--
-- 转储表的索引
--

--
-- 表的索引 `t_addr`
--
ALTER TABLE `t_addr`
  ADD UNIQUE KEY `hash` (`hash`) USING BTREE,
  ADD KEY `first_used_block_height` (`first_used_block_height`);

--
-- 表的索引 `t_addr_asset`
--
ALTER TABLE `t_addr_asset`
  ADD KEY `addr` (`addr`),
  ADD KEY `asset_id` (`asset_id`);

--
-- 表的索引 `t_assets`
--
ALTER TABLE `t_assets`
  ADD UNIQUE KEY `asset_hash` (`hash`) USING BTREE;

--
-- 表的索引 `t_assets_issue_record`
--
ALTER TABLE `t_assets_issue_record`
  ADD PRIMARY KEY (`asset_id`,`issue_time`,`issue_to`);

--
-- 表的索引 `t_assets_transfer`
--
ALTER TABLE `t_assets_transfer`
  ADD UNIQUE KEY `hash` (`hash`,`from_addr`,`to_addr`) USING BTREE,
  ADD KEY `height_tx_to_addr` (`to_addr`,`height_tx_idx_union`) USING BTREE,
  ADD KEY `height_tx_from_addr` (`from_addr`,`height_tx_idx_union`) USING BTREE;

--
-- 表的索引 `t_blocks`
--
ALTER TABLE `t_blocks`
  ADD PRIMARY KEY (`hash`),
  ADD KEY `height` (`height`);

--
-- 表的索引 `t_pay`
--
ALTER TABLE `t_pay`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `t_sigchain`
--
ALTER TABLE `t_sigchain`
  ADD KEY `height_tx` (`height`,`tx_hash`) USING BTREE,
  ADD KEY `tx_hash` (`tx_hash`),
  ADD KEY `tx` (`tx_hash`),
  ADD KEY `chord_addr` (`id`);

--
-- 表的索引 `t_transactions`
--
ALTER TABLE `t_transactions`
  ADD PRIMARY KEY (`hash`);

--
-- 表的索引 `t_generate_id`
--
ALTER TABLE `t_generate_id`
  ADD KEY `height` (`height`,`tx_hash`) USING BTREE,
  ADD KEY `tx_hash` (`tx_hash`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `t_pay`
--
ALTER TABLE `t_pay`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
