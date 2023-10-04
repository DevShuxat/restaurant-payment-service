CREATE SCHEMA IF NOT EXISTS restaurant_payment;


CREATE TABLE IF NOT EXISTS restaurant_payment.merchant_settings (
  entity_id varchar(36)  PRIMARY KEY,
  cashbox_id varchar(255),
  enabled boolean,
  created_at timestamp,
  updated_at timestamp
);

CREATE TABLE IF NOT EXISTS restaurant_payment.cards (
  id varchar(36) PRIMARY KEY,
  card_token text not null,
  created_at timestamp,
)

CREATE TABLE IF NOT EXISTS restaurant_payment.receipts (
  id varchar(36) PRIMARY,
  order_id varchar(36) not null,
  card_id varchar(36) not null,
  amount integer,
  data jsonb,
  created_at timestamp,
  updated_at timestamp 
)


CREATE INDEX IF NOT EXISTS idx_order_receipt ON restaurant_payment.receipts (order_id);

CREATE INDEX IF NOT EXISTS restaurant_payment.transactions (
  id varchar(36) PRIMARY KEY,
  currency varchar(3),
  status varchar(36),
  create_time integer,
  pay_time integer,
  cancel_time integer,
  card_id varchar(36) not null,
  amount integer,
  account jsonb,
  created_at timestamp,
  updated_at timestamp
);



