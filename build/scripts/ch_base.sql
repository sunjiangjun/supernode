CREATE
    DATABASE IF NOT EXISTS base;

CREATE TABLE IF NOT EXISTS base.address
(

    `token`
        String,

    `address`
        String,

    `tx_type`
        String,

    `block_chain`
        Int64,

    `id`
        Int64
) ENGINE = ReplacingMergeTree ORDER BY id SETTINGS index_granularity = 8192;


CREATE TABLE IF NOT EXISTS base.sub_filter
(

    `token`
        String,

    `tx_code`
        String,
    `params`
        String,

    `block_chain`
        Int64,

    `id`
        Int64
) ENGINE = ReplacingMergeTree ORDER BY id SETTINGS index_granularity = 8192;


CREATE TABLE IF NOT EXISTS base.token
(


    `token`
        String,

    `email`
        String,

    `id`
        Int64
) ENGINE = ReplacingMergeTree ORDER BY id SETTINGS index_granularity = 8192;
