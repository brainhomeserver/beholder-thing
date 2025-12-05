Feature: Backend health endpoint

  Scenario: Service reports healthy
    When I send a GET request to "/healthz"
    Then the response code should be 200
    And the header "Content-Type" should be "application/json"
    And the response body should be:
      """
      {"status":"ok"}
      """
