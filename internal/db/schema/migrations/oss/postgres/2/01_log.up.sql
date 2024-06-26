-- Copyright (c) HashiCorp, Inc.
-- SPDX-License-Identifier: MPL-2.0

begin;

-- management of the log_migation table is moved into
-- the schema.Manager prior to migrations running
-- since it interacts with the boundary_schema_version table

-- log_migration entries represent logs generated during migrations
-- create table log_migration(
--   id bigint generated always as identity primary key,
--   migration_version bigint not null, -- cannot declare FK since the table is truncated during runtime
--   create_time wt_timestamp,
--   entry text not null
-- );
-- comment on table log_migration is 
-- 'log_migration entries are logging output from databaes migrations';
-- 
-- -- log_migration triggers
-- create trigger 
--   default_create_time_column
-- before
-- insert on log_migration
--   for each row execute procedure default_create_time();
-- 
-- create trigger
--   immutable_columns
-- before
-- update on log_migration
--   for each row execute procedure immutable_columns('id', 'migration_version', 'create_time', 'entry');
-- 
-- 
-- -- log_migration_version() defines a function to be used in a "before update"
-- -- trigger for log_migrations entries.  Its intent: set the log_migration
-- -- version column to the current migration version.  
-- create or replace function
--   log_migration_version()
--   returns trigger
-- as $$
--   declare current_version bigint;
--   begin
--     select max(version) from boundary_schema_version into current_version;
--     new.migration_version = current_version;
--     return new;
--   end;
-- $$ language plpgsql;
-- comment on function log_migration_version() is
-- 'log_migration_version will set the log_migration entries to the current migration version';    
-- 
-- create trigger
--   migration_version_column
-- before
-- insert on log_migration
--   for each row execute procedure log_migration_version();
-- 
commit;
