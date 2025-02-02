Feature: Docker Container
  In order to test Shipyard creates containers correctly using variables
  I should apply a blueprint which defines a simple container setup
  and test the resources are created correctly

Scenario: Single Container with Shipyard Variables
  Given the following environment variables are set
    | key            | value                 |
    | BAH            | bah                   |
  And the following shipyard variables are set
    | key            | value                 |
    | something      | set by test           |
  And I have a running blueprint
  Then the following resources should be running
    | name                      | type      |
    | onprem                    | network   |
    | consul                    | container |
    | envoy                     | sidecar   |
    | consul-container-http     | container_ingress   |
  And the info "{.Config.Env}" for the running "container" called "consul" should contain "something=set by test"
  And the info "{.Config.Env}" for the running "container" called "consul" should contain "foo=bah"