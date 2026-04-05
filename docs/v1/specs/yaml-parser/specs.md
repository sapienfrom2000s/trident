Only sequential commands are allowed right now

- some command
- some other command
- ...

From user perspective he should be able to run it in the following
way:

- script.csh

```
# script.csh
#!/bin/csh
echo "Hello World!"
```

Specs:

1. Reject anything other than array.
