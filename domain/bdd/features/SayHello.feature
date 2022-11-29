Feature: Creation of Shapes

  Scenario: create a rectangle
    Given length of 2 and width of 4
    When I create a rectangle
    Then it area is "8"
