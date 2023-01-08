Feature: Creation of Shapes

  @Done
  Scenario: create a rectangle
    Given length of 2 and width of 4
    When I create a rectangle
    Then it area is "8"

  @Done
  Scenario: create a circle
    Given radius of 2
    When I create a circle
    Then it area is "12.56"

  @Done
  Scenario: stretch a rectangle
    Given an existing "rectangle"
    When I stretch it by 2
    Then it area is "32"

  @Done
  Scenario: stretch a circle
    Given an existing "circle"
    When I stretch it by 2
    Then it area is "50.24"
