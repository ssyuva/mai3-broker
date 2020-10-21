create table perpetuals
(
  id SERIAL PRIMARY KEY,
  perpetual_address text not null,
  symbol text not null,
  collateral_token_symbol text not null,
  price_tick numeric(32,18) not null,
  price_decimals integer not null,
  price_symbol text not null,
  contract_size_symbol text not null,
  amount_decimals integer not null,
  contract_type text not null,
  is_published boolean not null default false,
  perpetual_type text not null,
  broker_address text not null
);

create unique index idx_perpetual_address on perpetuals (perpetual_address);

-- orders table
create table orders
(
  id SERIAL PRIMARY KEY,
  order_hash text not null,
  trader_address text not null,
  type text not null,
  side text not null,
  price numeric(32,18) not null,
  amount numeric(32,18) not null,
  version text not null,
  expires_at timestamp,
  salt bigint not null,
  chain_id bigint not null,
  perpetual_address text not null,  
  status text not null,
  available_amount numeric(32,18) not null,
  confirmed_amount numeric(32,18) not null,
  confirmed_volume numeric(32,18) not null,
  filled_price numeric(32,18) not null,
  canceled_amount numeric(32,18) not null,
  pending_amount numeric(32,18) not null,
  pending_volume numeric(32,18) not null,
  gas_fee_amount numeric(32,18) not null,
  signature text not null,
  cancels_json json not null,
  updated_at timestamp,
  created_at timestamp
);

create unique index idx_order_hash on orders (order_hash);
create unique index idx_order_signature on orders (signature);
create index idx_perpetual_address_status on orders (perpetual_address, status); -- where perpetual_address, pending
create index idx_perpetual_address_trader_address on orders (trader_address, perpetual_address, status, created_at); -- where trader_address, perpetual_address, pending
create index idx_perpetual_trader_multistatus on orders (trader_address, created_at, status, perpetual_address); -- where trader_address, status in (...), perpetual_address
create index idx_trader_status on orders (trader_address, status, created_at); -- where trader_address, status in (...) without market_id

create table match_transactions
(
  id text PRIMARY KEY,
  perpetual_address text not null,
  match_json text not null,
  status text not null,
  created_at timestamp not null,
  block_confirmed bool not null,
  block_number int,
  transaction_hash text,
  sent_at timestamp,
  executed_at timestamp
);

create index idx_match_transactions_status_perpetual_address_created_at on match_transactions(status, perpetual_address, created_at);
create index idx_match_transactions_block_number on match_transactions(block_number);


create table watchers
(
  id SERIAL PRIMARY KEY,
  synced_block_number integer not null,
  initial_block_number integer not null
);

create table synced_blocks
(
  watcher_id integer not null,
  block_number integer not null,
  block_hash text not null,
  sync_at timestamp not null,
  PRIMARY KEY(watcher_id, block_number)
);