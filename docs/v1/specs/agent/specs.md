Specs:

1. Intercept internal request from server to run a job.
2. Validate the fields
  1. Job id is must
3. Spawns a new process.
4. Executes commands from .aegis.yml in pwd
5. Sends heartbeat to server every 5 sec
6. Send success/failure status
