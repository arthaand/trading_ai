-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.t_transaction_dt (
	id uuid NOT NULL,
	t_transaction_id uuid NOT NULL,
	product_id varchar(100) NOT NULL,
	product_name varchar(200) NOT NULL,
	unit_price int8 NOT NULL,
	price int8 NOT NULL,
	quantity int8 NOT NULL,
	created_at timestamp NOT NULL,
	created_by varchar(40) NOT NULL,
	updated_at timestamp NULL,
	updated_by varchar(40) NULL,
	CONSTRAINT t_transaction_dt_pk PRIMARY KEY (id)
);
ALTER TABLE public.t_transaction_dt ADD CONSTRAINT t_transaction_dt_fk FOREIGN KEY (t_transaction_id) REFERENCES public.t_transaction(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.t_transaction_dt;
-- +goose StatementEnd
