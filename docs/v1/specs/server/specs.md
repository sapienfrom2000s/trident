Specs:

1. Github Webhook handling
- Success — valid payload, returns normalized event; saves event to db and creates a job in db as well.
- Failure — invalid payload

2. Webhook secret validation (GitHub)
- Success — valid HMAC-SHA256 signature matches payload
- Failure — invalid signature (tampered payload or wrong secret)
- Failure — missing X-Hub-Signature-256 header

2. Job(exposed via REST)
	1. Create
	2. Stop(by User or Process)
	3. Retry

3. Process supervision
- Kill a stuck process when last heartbeat is older than 30s (SIGKILL)
- Respawn the job after killing (max 2 retries)

4. Heartbeat endpoint
- Success — known job id, timestamp recorded
- Failure — unknown job id
