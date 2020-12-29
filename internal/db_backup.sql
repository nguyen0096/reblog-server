--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1
-- Dumped by pg_dump version 13.1

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

--
-- Name: rb_core; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA rb_core;


ALTER SCHEMA rb_core OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: user; Type: TABLE; Schema: rb_core; Owner: postgres
--

CREATE TABLE rb_core."user" (
    id integer NOT NULL,
    username character varying(100) NOT NULL,
    first_name character varying(100),
    last_name character varying(100),
    address character varying(200)
);


ALTER TABLE rb_core."user" OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: rb_core; Owner: postgres
--

CREATE SEQUENCE rb_core.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE rb_core.user_id_seq OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: rb_core; Owner: postgres
--

ALTER SEQUENCE rb_core.user_id_seq OWNED BY rb_core."user".id;


--
-- Name: user id; Type: DEFAULT; Schema: rb_core; Owner: postgres
--

ALTER TABLE ONLY rb_core."user" ALTER COLUMN id SET DEFAULT nextval('rb_core.user_id_seq'::regclass);


--
-- Data for Name: user; Type: TABLE DATA; Schema: rb_core; Owner: postgres
--

COPY rb_core."user" (id, username, first_name, last_name, address) FROM stdin;
1	nguyen01	Nguyen	Test 01	\N
2	tu01	Tu	Test 01	\N
3	nguyen02	Nguyen	Test 02	\N
4	nguyen03	Nguyen	Test 03	\N
5	nguyen04	Nguyen	Test 04	\N
6	nguyen05	Nguyen	Test 05	\N
\.


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: rb_core; Owner: postgres
--

SELECT pg_catalog.setval('rb_core.user_id_seq', 6, true);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: rb_core; Owner: postgres
--

ALTER TABLE ONLY rb_core."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: SCHEMA rb_core; Type: ACL; Schema: -; Owner: postgres
--

GRANT ALL ON SCHEMA rb_core TO reblog;


--
-- Name: TABLE "user"; Type: ACL; Schema: rb_core; Owner: postgres
--

GRANT ALL ON TABLE rb_core."user" TO reblog;


--
-- PostgreSQL database dump complete
--

