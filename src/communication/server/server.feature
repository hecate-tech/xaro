
Feature: Player management 
    Scenario: User joins the server
    Given Player and Connection information is valid
        When The user attempts connection
        Then add the client to list of users if not already existant


    
    