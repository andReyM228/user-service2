-- +migrate Up
create extension
    if not exists "uuid-ossp";

-- +migrate StatementBegin
create
or replace function trigger_set_updated_at() returns trigger as
$$
begin
    NEW.updated_at = now();

return NEW;

end
$$ language plpgsql;

-- +migrate StatementEnd
-- +migrate Down
drop function
    if exists trigger_set_updated_at cascade;