--
-- PostgreSQL database dump
--

-- Dumped from database version 10.10
-- Dumped by pg_dump version 11.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id uuid NOT NULL,
    block_number integer NOT NULL,
    "timestamp" timestamp without time zone NOT NULL,
    hash character varying(255) NOT NULL,
    nonce integer NOT NULL,
    block_hash character varying(255) NOT NULL,
    form_addr character varying(255) NOT NULL,
    to_addr character varying(255) NOT NULL,
    value character varying(255) NOT NULL,
    gas character varying(255) NOT NULL,
    gas_price character varying(255) NOT NULL,
    is_error integer NOT NULL,
    input character varying(255) NOT NULL,
    cumulative_gas_used character varying(255) NOT NULL,
    gas_used character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    hide boolean DEFAULT false NOT NULL,
    message character varying(255) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- Name: watch_addresses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.watch_addresses (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    address character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.watch_addresses OWNER TO postgres;

--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: watch_addresses watch_addresses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.watch_addresses
    ADD CONSTRAINT watch_addresses_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: transactions_hash_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX transactions_hash_idx ON public.transactions USING btree (hash);


--
-- PostgreSQL database dump complete
--

