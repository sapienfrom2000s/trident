Specs:

1. Implement store method
2. Use timout
3. Schema
  1. Jobs table - created_at, scheduled_at, started_at, ended_at, total_time, status. Some of them are mandatory and some are optional.
  2. Agents table - job_id, status, last_heartbeat_at, spawned_at, completed_at, heartbeat_count, pid
  3. Events table - commit_hash, repo, author, branch, provider
