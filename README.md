# GitHub Hook Handler
GitHub has the capability to send notifications of various events to arbitrary endpoints, the [webhook's url](https://developer.github.com/webhooks/). These events can be useful for triggering automated testing or deploys or even user management or anything related to the code or other entities managed in GitHub.

This project aims to make consuming and taking action on these hooks as easy as possible.

## Goals
Initially this project is being built as part of Refinery29's 2018 summer hackathon, and so the scope must necessarily be limited. This is the list of my goals in general priority order to accomplish.

-  [x] Expose a server which listens for and consumes a simple json blob *(Done prior to hackathon)*
-  [ ] Consume the pull request labeled event
-  [ ] Consume the pull request unlabeled event
-  [ ] Trigger http call on event
-  [ ] On a 'test' labeled event, run tests
-  [ ] Run the tests in a container
-  [ ] Make the event to trigger tests configurable
-  [ ] Run the tests from a config file in the repo




