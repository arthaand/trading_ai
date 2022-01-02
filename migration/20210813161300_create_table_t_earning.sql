-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.t_earning (
	id uuid NOT NULL,
	t_transaction_id uuid NOT NULL,
	earning_code varchar(40) NULL,
	transaction_id varchar(40) NULL,
	requestor_institution_id varchar(40) NULL,
	reward_type int4 NULL,
	reward_value int8 NULL,
	request_at timestamp NULL,
	account_no varchar(40) NULL,
	transfer_id uuid NULL,
	status_cd varchar(3) NOT NULL,
	status_msg varchar(100) NULL,
	created_at timestamp NOT NULL,
	created_by varchar(40) NOT NULL,
	updated_at timestamp NULL,
	updated_by varchar(40) NULL,
	CONSTRAINT t_earning_pk PRIMARY KEY (id)
);
ALTER TABLE public.t_earning ADD CONSTRAINT t_earning_fk FOREIGN KEY (t_transaction_id) REFERENCES public.t_transaction(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.t_earning;
-- +goose StatementEnd
