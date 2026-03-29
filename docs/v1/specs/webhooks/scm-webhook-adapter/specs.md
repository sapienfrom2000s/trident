Add generic adapter to make sure that multiple vendors are supported. We normalize the fields at our end
and then use it. The struct structure might look something like the following:

```
NormalizedEvent {
  repo_url      string(mandatory)
  repo_name     string(mandatory)
  commit_sha    string(mandatory)
  branch        string(mandatory)
  triggered_by  string
  event_type    push | pull_request(mandatory)
  provider      github | gitlab
  metadata      map[string]string
}
```

We will expose different webhook API's for different providers.
POST /webhook/github → GitHubAdapter
POST /webhook/gitlab → GitLabAdapter

All the adapters should implement Parse and Validate.

### Specs

1. Interface contract — Parse
  1. Check if fields are being set
  2. Check if metadata field is being set properly
2. Interface contract - Validate
  1. Mandatory fields are allowed
3. Unsupported event type — record but do not trigger
4. Error response format
5. Verify signatures - if the source is what it says it is
