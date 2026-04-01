Specs:

1. Intercept webhooks
2. Forward request to adapter
3. Spawn agent
  1. Spawn a new process
  2. Retry mechanism
  3. Kill a stuck process
  4. Handle explicit error from agent
5. Intercept heartbeat from agent via rest
6. Intercept job spawning via rest

  
Marksman Loop:

1. Kill agent forcefully(-9) when last heartbeat was older than 30s. Ask server to respawn the job(max 2).
