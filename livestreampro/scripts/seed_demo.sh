#!/bin/bash
# Seed the demo environment with example channels, streams and chat messages.
#
# Usage: ./scripts/seed_demo.sh
#
# Requires `psql` to connect to Postgres using environment variables:
#   PGHOST, PGPORT, PGUSER, PGPASSWORD, PGDATABASE
# These default to the Docker Compose values if not set.
# `redis-cli` is optional for seeding chat history.
set -euo pipefail

echo "Seeding demo channels and streams..."

command -v psql >/dev/null || { echo "psql command not found" >&2; exit 1; }

psql <<'SQL'
INSERT INTO channels (id, name, description) VALUES
  (1, 'music', 'Live music sets'),
  (2, 'gaming', 'Daily gaming streams')
ON CONFLICT DO NOTHING;

INSERT INTO streams (id, channel_id, title, start_time) VALUES
  (100, 1, 'Jazz jam session', NOW()),
  (101, 2, 'Speedrun practice', NOW())
ON CONFLICT DO NOTHING;
SQL

echo "Database seeded."

if command -v redis-cli >/dev/null; then
  echo "Seeding sample chat messages..."
  redis-cli LPUSH chat:1 "Welcome to the music channel!" "First track coming up"
  redis-cli LPUSH chat:2 "Who's ready for a PB?" "Welcome gamers!"
fi

echo "Done."

