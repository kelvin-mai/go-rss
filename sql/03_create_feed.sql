create table feeds(
  id text not null primary key default nanoid(),
  created_at timestamptz not null default now(),
  updated_at timestamptz,
  name text not null,
  url text not null
);

create trigger feeds_updated_at
  before update on feeds
  for each row
  execute procedure moddatetime (updated_at);

create table user_feed (
  id text not null primary key default nanoid(),
  created_at timestamptz not null default now(),
  updated_at timestamptz,
  user_id text not null references users(id),
  feed_id text not null references feeds(id)
);

create trigger user_feed_updated_at
  before update on feeds
  for each row
  execute procedure moddatetime (updated_at);