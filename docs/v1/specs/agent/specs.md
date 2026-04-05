Specs:

1. Intercept internal request from server to run a job. Validates new params.
2. Spawns a new process.
3. Executes commands from .trident.yml in pwd
4. Sends heartbeat to server every 5 sec
5. Send success/failure status
