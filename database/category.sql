
CREATE TABLE category
(
  id uuid NOT NULL,
  name name,
  created_at timestamp without time zone NOT NULL DEFAULT now(),
  updated_at timestamp without time zone,
  items jsonb DEFAULT '{}'::jsonb,
  CONSTRAINT category_pkey PRIMARY KEY (id)
)