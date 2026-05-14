-- DROP SCHEMA public;

CREATE SCHEMA public AUTHORIZATION postgres;
-- public.movies definition

-- Drop table

-- DROP TABLE public.movies;

CREATE TABLE public.movies (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"name" varchar(255) NOT NULL,
	duration_minutes int8 NOT NULL,
	description text NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT movies_pkey PRIMARY KEY (id)
);

-- Permissions

ALTER TABLE public.movies OWNER TO postgres;
GRANT ALL ON TABLE public.movies TO postgres;


-- public.regions definition

-- Drop table

-- DROP TABLE public.regions;

CREATE TABLE public.regions (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"name" varchar(255) NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT regions_pkey PRIMARY KEY (id)
);

-- Permissions

ALTER TABLE public.regions OWNER TO postgres;
GRANT ALL ON TABLE public.regions TO postgres;


-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"name" varchar(255) NOT NULL,
	email varchar(100) NOT NULL,
	"password" varchar(255) NOT NULL,
	"role" text DEFAULT 'CUSTOMER'::text NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);
CREATE UNIQUE INDEX user_deleted_at ON public.users USING btree (email) WHERE (deleted_at IS NULL);

-- Permissions

ALTER TABLE public.users OWNER TO postgres;
GRANT ALL ON TABLE public.users TO postgres;


-- public.cities definition

-- Drop table

-- DROP TABLE public.cities;

CREATE TABLE public.cities (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	region_id uuid NOT NULL,
	"name" varchar(255) NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT cities_pkey PRIMARY KEY (id),
	CONSTRAINT fk_cities_region FOREIGN KEY (region_id) REFERENCES public.regions(id) ON DELETE RESTRICT
);
CREATE INDEX idx_region_id ON public.cities USING btree (region_id);

-- Permissions

ALTER TABLE public.cities OWNER TO postgres;
GRANT ALL ON TABLE public.cities TO postgres;


-- public.cinemas definition

-- Drop table

-- DROP TABLE public.cinemas;

CREATE TABLE public.cinemas (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	city_id uuid NOT NULL,
	"name" varchar(255) NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT cinemas_pkey PRIMARY KEY (id),
	CONSTRAINT fk_cinemas_city FOREIGN KEY (city_id) REFERENCES public.cities(id) ON DELETE RESTRICT
);
CREATE INDEX idx_cty_id ON public.cinemas USING btree (city_id);

-- Permissions

ALTER TABLE public.cinemas OWNER TO postgres;
GRANT ALL ON TABLE public.cinemas TO postgres;


-- public.studios definition

-- Drop table

-- DROP TABLE public.studios;

CREATE TABLE public.studios (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	cinema_id uuid NOT NULL,
	"name" varchar(255) NOT NULL,
	capacity int8 NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT studios_pkey PRIMARY KEY (id),
	CONSTRAINT fk_cinemas_studios FOREIGN KEY (cinema_id) REFERENCES public.cinemas(id)
);
CREATE INDEX idx_cinema_id ON public.studios USING btree (cinema_id);

-- Permissions

ALTER TABLE public.studios OWNER TO postgres;
GRANT ALL ON TABLE public.studios TO postgres;


-- public.schedules definition

-- Drop table

-- DROP TABLE public.schedules;

CREATE TABLE public.schedules (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	movie_id uuid NOT NULL,
	studio_id uuid NOT NULL,
	show_time timestamptz NOT NULL,
	end_time timestamptz NOT NULL,
	price numeric(12, 2) NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT schedules_pkey PRIMARY KEY (id),
	CONSTRAINT fk_movies_schedules FOREIGN KEY (movie_id) REFERENCES public.movies(id),
	CONSTRAINT fk_schedules_studio FOREIGN KEY (studio_id) REFERENCES public.studios(id) ON DELETE RESTRICT
);
CREATE INDEX idx_movie_id ON public.schedules USING btree (movie_id);
CREATE INDEX idx_schedules_deleted_at ON public.schedules USING btree (deleted_at);
CREATE INDEX idx_studio_id ON public.schedules USING btree (studio_id);

-- Permissions

ALTER TABLE public.schedules OWNER TO postgres;
GRANT ALL ON TABLE public.schedules TO postgres;


-- public.seats definition

-- Drop table

-- DROP TABLE public.seats;

CREATE TABLE public.seats (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	studio_id uuid NOT NULL,
	seat_row text NOT NULL,
	seat_number int8 NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT seats_pkey PRIMARY KEY (id),
	CONSTRAINT fk_studios_seats FOREIGN KEY (studio_id) REFERENCES public.studios(id)
);

-- Permissions

ALTER TABLE public.seats OWNER TO postgres;
GRANT ALL ON TABLE public.seats TO postgres;


-- public.bookings definition

-- Drop table

-- DROP TABLE public.bookings;

CREATE TABLE public.bookings (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	user_id uuid NOT NULL,
	schedule_id uuid NOT NULL,
	booking_code text NOT NULL,
	total_amount numeric(12, 2) NOT NULL,
	status varchar(50) NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT bookings_pkey PRIMARY KEY (id),
	CONSTRAINT fk_bookings_schedule FOREIGN KEY (schedule_id) REFERENCES public.schedules(id) ON DELETE RESTRICT,
	CONSTRAINT fk_users_bookings FOREIGN KEY (user_id) REFERENCES public.users(id)
);
CREATE UNIQUE INDEX idx_bookings_booking_code ON public.bookings USING btree (booking_code);
CREATE INDEX idx_schedule_id ON public.bookings USING btree (schedule_id);
CREATE INDEX idx_user_id ON public.bookings USING btree (user_id);

-- Permissions

ALTER TABLE public.bookings OWNER TO postgres;
GRANT ALL ON TABLE public.bookings TO postgres;


-- public.payments definition

-- Drop table

-- DROP TABLE public.payments;

CREATE TABLE public.payments (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	booking_id uuid NOT NULL,
	payment_method varchar(50) NOT NULL,
	final_amount numeric(20, 2) NULL,
	status varchar(50) NOT NULL,
	paid_at timestamptz NOT NULL,
	CONSTRAINT payments_pkey PRIMARY KEY (id),
	CONSTRAINT fk_bookings_payment FOREIGN KEY (booking_id) REFERENCES public.bookings(id)
);

-- Permissions

ALTER TABLE public.payments OWNER TO postgres;
GRANT ALL ON TABLE public.payments TO postgres;


-- public.refunds definition

-- Drop table

-- DROP TABLE public.refunds;

CREATE TABLE public.refunds (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	booking_id uuid NOT NULL,
	amount numeric(20, 2) NULL,
	status varchar(50) NOT NULL,
	refunded_at timestamptz NOT NULL,
	CONSTRAINT refunds_pkey PRIMARY KEY (id),
	CONSTRAINT fk_bookings_refund FOREIGN KEY (booking_id) REFERENCES public.bookings(id)
);

-- Permissions

ALTER TABLE public.refunds OWNER TO postgres;
GRANT ALL ON TABLE public.refunds TO postgres;


-- public.seat_locks definition

-- Drop table

-- DROP TABLE public.seat_locks;

CREATE TABLE public.seat_locks (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	user_id uuid NOT NULL,
	schedule_id uuid NOT NULL,
	seat_id uuid NOT NULL,
	status varchar(30) NOT NULL,
	expired_at timestamptz NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT seat_locks_pkey PRIMARY KEY (id),
	CONSTRAINT fk_seat_locks_schedule FOREIGN KEY (schedule_id) REFERENCES public.schedules(id) ON DELETE CASCADE,
	CONSTRAINT fk_seat_locks_seat FOREIGN KEY (seat_id) REFERENCES public.seats(id) ON DELETE CASCADE,
	CONSTRAINT fk_seat_locks_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX idx_seat_locks_schedule_id ON public.seat_locks USING btree (schedule_id);
CREATE UNIQUE INDEX idx_seat_locks_seat_id ON public.seat_locks USING btree (seat_id);
CREATE INDEX idx_seat_locks_user_id ON public.seat_locks USING btree (user_id);

-- Permissions

ALTER TABLE public.seat_locks OWNER TO postgres;
GRANT ALL ON TABLE public.seat_locks TO postgres;


-- public.booking_seats definition

-- Drop table

-- DROP TABLE public.booking_seats;

CREATE TABLE public.booking_seats (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	booking_id uuid NOT NULL,
	schedule_id uuid NOT NULL,
	seat_id uuid NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT booking_seats_pkey PRIMARY KEY (id),
	CONSTRAINT fk_booking_seats_schedule FOREIGN KEY (schedule_id) REFERENCES public.schedules(id) ON DELETE RESTRICT,
	CONSTRAINT fk_booking_seats_seat FOREIGN KEY (seat_id) REFERENCES public.seats(id) ON DELETE RESTRICT,
	CONSTRAINT fk_bookings_booking_seats FOREIGN KEY (booking_id) REFERENCES public.bookings(id)
);
CREATE INDEX idx_booking_id ON public.booking_seats USING btree (booking_id);
CREATE INDEX "uniqueIndex" ON public.booking_seats USING btree (schedule_id, seat_id);

-- Permissions

ALTER TABLE public.booking_seats OWNER TO postgres;
GRANT ALL ON TABLE public.booking_seats TO postgres;



-- DROP FUNCTION public.uuid_generate_v1();

CREATE OR REPLACE FUNCTION public.uuid_generate_v1()
 RETURNS uuid
 LANGUAGE c
 PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_generate_v1$function$
;

-- Permissions

ALTER FUNCTION public.uuid_generate_v1() OWNER TO postgres;
GRANT ALL ON FUNCTION public.uuid_generate_v1() TO postgres;

-- DROP FUNCTION public.uuid_generate_v1mc();

CREATE OR REPLACE FUNCTION public.uuid_generate_v1mc()
 RETURNS uuid
 LANGUAGE c
 PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_generate_v1mc$function$
;

-- Permissions

ALTER FUNCTION public.uuid_generate_v1mc() OWNER TO postgres;
GRANT ALL ON FUNCTION public.uuid_generate_v1mc() TO postgres;

-- DROP FUNCTION public.uuid_generate_v3(uuid, text);

CREATE OR REPLACE FUNCTION public.uuid_generate_v3(namespace uuid, name text)
 RETURNS uuid
 LANGUAGE c
 IMMUTABLE PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_generate_v3$function$
;

-- Permissions

ALTER FUNCTION public.uuid_generate_v3(uuid, text) OWNER TO postgres;
GRANT ALL ON FUNCTION public.uuid_generate_v3(uuid, text) TO postgres;

-- DROP FUNCTION public.uuid_generate_v4();

CREATE OR REPLACE FUNCTION public.uuid_generate_v4()
 RETURNS uuid
 LANGUAGE c
 PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_generate_v4$function$
;

-- Permissions

ALTER FUNCTION public.uuid_generate_v4() OWNER TO postgres;
GRANT ALL ON FUNCTION public.uuid_generate_v4() TO postgres;

-- DROP FUNCTION public.uuid_generate_v5(uuid, text);

CREATE OR REPLACE FUNCTION public.uuid_generate_v5(namespace uuid, name text)
 RETURNS uuid
 LANGUAGE c
 IMMUTABLE PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_generate_v5$function$
;

-- Permissions

ALTER FUNCTION public.uuid_generate_v5(uuid, text) OWNER TO postgres;
GRANT ALL ON FUNCTION public.uuid_generate_v5(uuid, text) TO postgres;

-- DROP FUNCTION public.uuid_nil();

CREATE OR REPLACE FUNCTION public.uuid_nil()
 RETURNS uuid
 LANGUAGE c
 IMMUTABLE PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_nil$function$
;

-- Permissions

ALTER FUNCTION public.uuid_nil() OWNER TO postgres;
GRANT ALL ON FUNCTION public.uuid_nil() TO postgres;

-- DROP FUNCTION public.uuid_ns_dns();

CREATE OR REPLACE FUNCTION public.uuid_ns_dns()
 RETURNS uuid
 LANGUAGE c
 IMMUTABLE PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_ns_dns$function$
;

-- Permissions

ALTER FUNCTION public.uuid_ns_dns() OWNER TO postgres;
GRANT ALL ON FUNCTION public.uuid_ns_dns() TO postgres;

-- DROP FUNCTION public.uuid_ns_oid();

CREATE OR REPLACE FUNCTION public.uuid_ns_oid()
 RETURNS uuid
 LANGUAGE c
 IMMUTABLE PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_ns_oid$function$
;

-- Permissions

ALTER FUNCTION public.uuid_ns_oid() OWNER TO postgres;
GRANT ALL ON FUNCTION public.uuid_ns_oid() TO postgres;

-- DROP FUNCTION public.uuid_ns_url();

CREATE OR REPLACE FUNCTION public.uuid_ns_url()
 RETURNS uuid
 LANGUAGE c
 IMMUTABLE PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_ns_url$function$
;

-- Permissions

ALTER FUNCTION public.uuid_ns_url() OWNER TO postgres;
GRANT ALL ON FUNCTION public.uuid_ns_url() TO postgres;

-- DROP FUNCTION public.uuid_ns_x500();

CREATE OR REPLACE FUNCTION public.uuid_ns_x500()
 RETURNS uuid
 LANGUAGE c
 IMMUTABLE PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_ns_x500$function$
;

-- Permissions

ALTER FUNCTION public.uuid_ns_x500() OWNER TO postgres;
GRANT ALL ON FUNCTION public.uuid_ns_x500() TO postgres;


-- Permissions

GRANT ALL ON SCHEMA public TO postgres;