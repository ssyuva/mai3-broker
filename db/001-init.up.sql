create table perpetuals
(
  id SERIAL PRIMARY KEY,
  perpetual_address text not null,
  governor_address text not null,
  share_token text not null,
  collateral_symbol text not null,
  operator_address text not null,
  oracle_address text not null,
  collateral_address text not null,
  is_published boolean not null default true,
  block_number bigint not null
);

create unique index idx_perpetual_address on perpetuals (perpetual_address);
create index idx_perpetuals_height on perpetuals (block_number);

-- orders table
create table orders
(
  id SERIAL PRIMARY KEY,
  order_hash text not null,
  trader_address text not null,
  type integer not null,
  price numeric(32,18) not null,
  amount numeric(32,18) not null,
  version integer not null,
  expires_at timestamp,
  salt bigint not null,
  is_close_only boolean not null DEFAULT FALSE,
  chain_id bigint not null,
  perpetual_address text not null,
  broker_address text not null,
  referrer_address text not null,
  relayer_address text not null,
  status text not null,
  available_amount numeric(32,18) not null,
  confirmed_amount numeric(32,18) not null,
  filled_price numeric(32,18) not null,
  canceled_amount numeric(32,18) not null,
  pending_amount numeric(32,18) not null,
  gas_fee_amount numeric(32,18) not null,
  signature text not null,
  cancels_json json not null,
  updated_at timestamp,
  created_at timestamp
);

create unique index idx_order_hash on orders (order_hash);
create index idx_perpetual_address_status on orders (perpetual_address, status); -- where perpetual_address, pending
create index idx_perpetual_address_trader_address on orders (trader_address, perpetual_address, status, created_at); -- where trader_address, perpetual_address, pending
create index idx_perpetual_trader_multistatus on orders (trader_address, created_at, status, perpetual_address); -- where trader_address, status in (...), perpetual_address
create index idx_trader_status on orders (trader_address, status, created_at); -- where trader_address, status in (...) without market_id

create table match_transactions
(
  id text PRIMARY KEY,
  perpetual_address text not null,
  broker_address text not null,
  match_json text not null,
  status text not null,
  block_confirmed bool not null,
  block_number int,
  transaction_hash text,
  created_at timestamp not null,
  executed_at timestamp
);

create index idx_match_transactions_status_perpetual_address_created_at on match_transactions(status, perpetual_address, created_at);
create index idx_match_transactions_block_number on match_transactions(block_number);


create table watchers
(
  id SERIAL PRIMARY KEY,
  synced_block_number bigint not null,
  initial_block_number bigint not null
);

create table synced_blocks
(
  watcher_id integer not null,
  block_number bigint not null,
  block_hash text not null,
  parent_hash text not null,
  block_time timestamp not null,
  PRIMARY KEY(watcher_id, block_number)
);

create table broker_nonces
(
  address text not null PRIMARY KEY,
  nonce integer not null,
  updated_at timestamp not null
);

create table launch_transactions
(
  id SERIAL PRIMARY KEY,
  tx_id text not null,
  from_address text,
  to_address text,
  type integer,
  inputs bytea,
  block_number bigint,
  block_hash text,
  block_time bigint,
  transaction_hash text,
  nonce bigint,
  gas_price bigint,
  gas_limit bigint,
  gas_used bigint,
  status integer,
  value numeric(78,18),
  commit_time timestamp,
  update_time timestamp
);

create index idx_transactions_tx_id on launch_transactions (tx_id);

create table kv_stores
(
  key text not null PRIMARY KEY,
  category text not null,
  value bytea,
  update_time timestamp
);