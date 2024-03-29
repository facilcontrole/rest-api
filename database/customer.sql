
CREATE TABLE customer
(
  id uuid NOT NULL,
  name name,
  company_id uuid NOT NULL,
  created_at timestamp without time zone NOT NULL DEFAULT now(),
  updated_at timestamp without time zone,
  items jsonb DEFAULT '{}'::jsonb,
  CONSTRAINT customer_pkey PRIMARY KEY (id)
)