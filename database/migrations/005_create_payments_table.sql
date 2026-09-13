CREATE TABLE IF NOT EXISTS payments (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  external_payment_id VARCHAR(128) UNIQUE NOT NULL,
  site_id UUID REFERENCES sites(id),
  amount NUMERIC(14,2) NOT NULL,
  currency VARCHAR(8) NOT NULL,
  gateway_id VARCHAR(128),
  status VARCHAR(50) NOT NULL DEFAULT 'pending',
  metadata JSONB,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
