-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.t_transaction (
	id uuid NOT NULL,
	institution_id varchar(40) NOT NULL,
	invoice_no varchar(100) NOT NULL,
	merchant_id varchar(40) NOT NULL,
	store_id varchar(40) NOT NULL,
	staff_id varchar(40) NOT NULL,
	staff_name varchar(100) NOT NULL,
	transaction_time timestamp NOT NULL,
	sub_total_amount int8 NOT NULL,
	service_charge int8 NOT NULL,
	tax int8 NOT NULL,
	total_amount int8 NOT NULL,
	created_at timestamp NOT NULL,
	created_by varchar(40) NOT NULL,
	updated_at timestamp NULL,
	updated_by varchar(40) NULL,
	CONSTRAINT t_institutionid_invoiceno_un UNIQUE (institution_id, invoice_no),
	CONSTRAINT t_transaction_pk PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.t_transaction;
-- +goose StatementEnd
